// Package imagesource defines various image sources typically registered as cameras in the API.
//
// Some sources are specific to a type of camera while some are general purpose sources that
// act as a component in an image transformation pipeline.
package imagesource

import (
	"bufio"
	"bytes"
	"context"

	// for embedding camera parameters.
	_ "embed"
	"fmt"
	"image"
	"io/ioutil"
	"net/http"

	"github.com/edaniels/golog"
	"github.com/edaniels/gostream"

	// register ppm.
	_ "github.com/lmittmann/ppm"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"go.viam.com/utils"

	"go.viam.com/rdk/component/camera"
	"go.viam.com/rdk/config"
	"go.viam.com/rdk/registry"
	"go.viam.com/rdk/rimage"
	"go.viam.com/rdk/rimage/transform"
	"go.viam.com/rdk/robot"
)

func init() {
	registry.RegisterComponent(camera.Subtype, "single_stream",
		registry.Component{Constructor: func(ctx context.Context, r robot.Robot,
			config config.Component, logger golog.Logger) (interface{}, error) {
			if config.ConvertedAttributes.(*rimage.AttrConfig).Stream == "" {
				return nil, errors.New("camera 'single_stream' needs attribute 'stream' (color, depth, or both)")
			}
			source, err := NewServerSource(config.ConvertedAttributes.(*rimage.AttrConfig), logger)
			if err != nil {
				return nil, err
			}
			return &camera.ImageSource{ImageSource: source}, nil
		}})

	config.RegisterComponentAttributeMapConverter(config.ComponentTypeCamera, "single_stream",
		func(attributes config.AttributeMap) (interface{}, error) {
			var conf rimage.AttrConfig
			return config.TransformAttributeMapToStruct(&conf, attributes)
		},
		&rimage.AttrConfig{})

	config.RegisterComponentAttributeConverter(config.ComponentTypeCamera, "single_stream", "intrinsic",
		func(val interface{}) (interface{}, error) {
			intrinsics := &transform.PinholeCameraIntrinsics{}
			decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: intrinsics})
			if err != nil {
				return nil, err
			}
			err = decoder.Decode(val)
			if err == nil {
				err = intrinsics.CheckValid()
			}
			return intrinsics, err
		})

	config.RegisterComponentAttributeConverter(config.ComponentTypeCamera, "single_stream", "warp",
		func(val interface{}) (interface{}, error) {
			warp := &transform.AlignConfig{}
			err := mapstructure.Decode(val, warp)
			if err == nil {
				err = warp.CheckValid()
			}
			return warp, err
		})

	config.RegisterComponentAttributeConverter(config.ComponentTypeCamera, "single_stream", "intrinsic_extrinsic",
		func(val interface{}) (interface{}, error) {
			matrices := &transform.DepthColorIntrinsicsExtrinsics{}
			decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: matrices})
			if err != nil {
				return nil, err
			}
			err = decoder.Decode(val)
			if err == nil {
				err = matrices.CheckValid()
			}
			return matrices, err
		})

	config.RegisterComponentAttributeConverter(config.ComponentTypeCamera, "single_stream", "homography",
		func(val interface{}) (interface{}, error) {
			homography := &transform.RawPinholeCameraHomography{}
			decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: homography})
			if err != nil {
				return nil, err
			}
			err = decoder.Decode(val)
			if err == nil {
				err = homography.CheckValid()
			}
			return homography, err
		})

	registry.RegisterComponent(camera.Subtype, "dual_stream",
		registry.Component{Constructor: func(ctx context.Context, r robot.Robot,
			config config.Component, logger golog.Logger) (interface{}, error) {
			if (config.ConvertedAttributes.(*rimage.AttrConfig).Color == "") || (config.ConvertedAttributes.(*rimage.AttrConfig).Depth == "") {
				return nil, errors.New("camera 'dual_stream' needs color and depth attributes")
			}
			return &camera.ImageSource{ImageSource: &dualServerSource{
				ColorURL:  config.ConvertedAttributes.(*rimage.AttrConfig).Color,
				DepthURL:  config.ConvertedAttributes.(*rimage.AttrConfig).Depth,
				isAligned: config.ConvertedAttributes.(*rimage.AttrConfig).Aligned,
			}}, nil
		}})

	config.RegisterComponentAttributeMapConverter(config.ComponentTypeCamera, "dual_stream",
		func(attributes config.AttributeMap) (interface{}, error) {
			var conf rimage.AttrConfig
			return config.TransformAttributeMapToStruct(&conf, attributes)
		},
		&rimage.AttrConfig{})

	registry.RegisterComponent(camera.Subtype, "file",
		registry.Component{Constructor: func(ctx context.Context, r robot.Robot,
			config config.Component, logger golog.Logger) (interface{}, error) {
			return &camera.ImageSource{ImageSource: &fileSource{
				config.ConvertedAttributes.(*rimage.AttrConfig).Color,
				config.ConvertedAttributes.(*rimage.AttrConfig).Depth, config.ConvertedAttributes.(*rimage.AttrConfig).Aligned,
			}}, nil
		}})

	config.RegisterComponentAttributeMapConverter(config.ComponentTypeCamera, "file",
		func(attributes config.AttributeMap) (interface{}, error) {
			var conf rimage.AttrConfig
			return config.TransformAttributeMapToStruct(&conf, attributes)
		},
		&rimage.AttrConfig{})
}

// staticSource is a fixed, stored image.
type staticSource struct {
	Img image.Image
}

// Next returns the stored image.
func (ss *staticSource) Next(ctx context.Context) (image.Image, func(), error) {
	return ss.Img, func() {}, nil
}

// fileSource stores the paths to a color and depth image.
type fileSource struct {
	ColorFN   string
	DepthFN   string
	isAligned bool // are color and depth image already aligned
}

// IsAligned returns a bool that is true if the color and depth images are aligned.
func (fs *fileSource) IsAligned() bool {
	return fs.isAligned
}

// Next returns the image stored in the color and depth files as an ImageWithDepth.
func (fs *fileSource) Next(ctx context.Context) (image.Image, func(), error) {
	img, err := rimage.NewImageWithDepth(fs.ColorFN, fs.DepthFN, fs.IsAligned())
	return img, func() {}, err
}

func decodeColor(colorData []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewBuffer(colorData))
	return img, err
}

func decodeDepth(depthData []byte) (*rimage.DepthMap, error) {
	return rimage.ReadDepthMap(bufio.NewReader(bytes.NewReader(depthData)))
}

func decodeBoth(bothData []byte, aligned bool) (*rimage.ImageWithDepth, error) {
	return rimage.ReadBothFromBytes(bothData, aligned)
}

func readyBytesFromURL(ctx context.Context, client http.Client, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		utils.UncheckedError(resp.Body.Close())
	}()
	return ioutil.ReadAll(resp.Body)
}

// dualServerSource stores two URLs, one which points the color source and the other to the
// depth source.
type dualServerSource struct {
	client    http.Client
	ColorURL  string // this is for a generic image
	DepthURL  string // this is for my bizarre custom data format for depth data
	isAligned bool   // are the color and depth image already aligned
}

// IsAligned returns true if the images returned from the two servers are already aligned
// with each other.
func (ds *dualServerSource) IsAligned() bool {
	return ds.isAligned
}

// Next requests the next images from both the color and depth source, and combines them
// together as an ImageWithDepth before returning them.
func (ds *dualServerSource) Next(ctx context.Context) (image.Image, func(), error) {
	colorData, err := readyBytesFromURL(ctx, ds.client, ds.ColorURL)
	if err != nil {
		return nil, nil, errors.Wrap(err, "couldn't ready color url")
	}
	img, err := decodeColor(colorData)
	if err != nil {
		return nil, nil, err
	}

	depthData, err := readyBytesFromURL(ctx, ds.client, ds.DepthURL)
	if err != nil {
		return nil, nil, errors.Wrap(err, "couldn't ready depth url")
	}
	// do this first and make sure ok before creating any mats
	depth, err := decodeDepth(depthData)
	if err != nil {
		return nil, nil, err
	}

	return rimage.MakeImageWithDepth(rimage.ConvertImage(img), depth, ds.IsAligned(), nil), func() {}, nil
}

// Close closes the connection to both servers.
func (ds *dualServerSource) Close() {
	ds.client.CloseIdleConnections()
}

// StreamType specifies what kind of image stream is coming from the camera.
type StreamType string

// The allowed types of streams that can come from an ImageSource.
const (
	ColorStream = StreamType("color")
	DepthStream = StreamType("depth")
	BothStream  = StreamType("both")
)

// serverSource streams the color/depth/both camera data from an external server at a given URL.
type serverSource struct {
	client    http.Client
	URL       string
	host      string
	stream    StreamType // specifies color, depth, or both stream
	isAligned bool       // are the color and depth image already aligned
	camera    rimage.Projector
}

// IsAligned is a bool that returns true if both.gz image is already aligned. If the server is only returning a single stream
// (either color or depth) IsAligned will return false.
func (s *serverSource) IsAligned() bool {
	return s.isAligned
}

// Projector is the  Projector which projects between 3D and 2D representations.
func (s *serverSource) Projector() rimage.Projector {
	return s.camera
}

// Close closes the server connection.
func (s *serverSource) Close() {
	s.client.CloseIdleConnections()
}

// Next returns the next image in the queue from the server.
func (s *serverSource) Next(ctx context.Context) (image.Image, func(), error) {
	var img *rimage.ImageWithDepth
	var err error

	allData, err := readyBytesFromURL(ctx, s.client, s.URL)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "couldn't read url (%s)", s.URL)
	}

	switch s.stream {
	case ColorStream:
		color, err := decodeColor(allData)
		if err != nil {
			return nil, nil, err
		}
		img = rimage.MakeImageWithDepth(rimage.ConvertImage(color), nil, false, s.Projector())
	case DepthStream:
		depth, err := decodeDepth(allData)
		if err != nil {
			return nil, nil, err
		}
		img = rimage.MakeImageWithDepth(rimage.ConvertImage(depth.ToGray16Picture()), depth, true, s.Projector())
	case BothStream:
		img, err = decodeBoth(allData, s.isAligned)
		if err != nil {
			return nil, nil, err
		}
		img.SetProjector(s.Projector())
	default:
		return nil, nil, errors.Errorf("do not know how to decode stream type %q", string(s.stream))
	}

	return img, func() {}, nil
}

// NewServerSource creates the ImageSource that streams color/depth/both data from an external server at a given URL.
func NewServerSource(cfg *rimage.AttrConfig, logger golog.Logger) (gostream.ImageSource, error) {
	_, camera, err := getCameraSystems(cfg, logger)
	if err != nil {
		return nil, err
	}

	return &serverSource{
		URL:       fmt.Sprintf("http://%s:%d/%s", cfg.Host, cfg.Port, cfg.Args),
		host:      cfg.Host,
		stream:    StreamType(cfg.Stream),
		isAligned: cfg.Aligned,
		camera:    camera,
	}, nil
}
