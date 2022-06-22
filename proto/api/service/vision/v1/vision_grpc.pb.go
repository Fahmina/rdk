// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VisionServiceClient is the client API for VisionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VisionServiceClient interface {
	// GetDetectorNames returns the list of detectors in the registry.
	GetDetectorNames(ctx context.Context, in *GetDetectorNamesRequest, opts ...grpc.CallOption) (*GetDetectorNamesResponse, error)
	// AddDetector adds a new detector to the registry.
	AddDetector(ctx context.Context, in *AddDetectorRequest, opts ...grpc.CallOption) (*AddDetectorResponse, error)
	// GetDetections will return a list of detections in the next image given a camera and a detector
	GetDetections(ctx context.Context, in *GetDetectionsRequest, opts ...grpc.CallOption) (*GetDetectionsResponse, error)
	// GetSegmenterNames returns the list of segmenters in the registry.
	GetSegmenterNames(ctx context.Context, in *GetSegmenterNamesRequest, opts ...grpc.CallOption) (*GetSegmenterNamesResponse, error)
	// GetSegmenterParameters returns the parameter fields needed for the given segmenter.
	GetSegmenterParameters(ctx context.Context, in *GetSegmenterParametersRequest, opts ...grpc.CallOption) (*GetSegmenterParametersResponse, error)
	// GetObjectPointClouds returns all the found objects in a pointcloud from a camera of the underlying robot,
	// as well as the 3-vector center of each of the found objects.
	// A specific MIME type can be requested but may not necessarily be the same one returned.
	GetObjectPointClouds(ctx context.Context, in *GetObjectPointCloudsRequest, opts ...grpc.CallOption) (*GetObjectPointCloudsResponse, error)
}

type visionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVisionServiceClient(cc grpc.ClientConnInterface) VisionServiceClient {
	return &visionServiceClient{cc}
}

func (c *visionServiceClient) GetDetectorNames(ctx context.Context, in *GetDetectorNamesRequest, opts ...grpc.CallOption) (*GetDetectorNamesResponse, error) {
	out := new(GetDetectorNamesResponse)
	err := c.cc.Invoke(ctx, "/proto.api.service.vision.v1.VisionService/GetDetectorNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) AddDetector(ctx context.Context, in *AddDetectorRequest, opts ...grpc.CallOption) (*AddDetectorResponse, error) {
	out := new(AddDetectorResponse)
	err := c.cc.Invoke(ctx, "/proto.api.service.vision.v1.VisionService/AddDetector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) GetDetections(ctx context.Context, in *GetDetectionsRequest, opts ...grpc.CallOption) (*GetDetectionsResponse, error) {
	out := new(GetDetectionsResponse)
	err := c.cc.Invoke(ctx, "/proto.api.service.vision.v1.VisionService/GetDetections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) GetSegmenterNames(ctx context.Context, in *GetSegmenterNamesRequest, opts ...grpc.CallOption) (*GetSegmenterNamesResponse, error) {
	out := new(GetSegmenterNamesResponse)
	err := c.cc.Invoke(ctx, "/proto.api.service.vision.v1.VisionService/GetSegmenterNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) GetSegmenterParameters(ctx context.Context, in *GetSegmenterParametersRequest, opts ...grpc.CallOption) (*GetSegmenterParametersResponse, error) {
	out := new(GetSegmenterParametersResponse)
	err := c.cc.Invoke(ctx, "/proto.api.service.vision.v1.VisionService/GetSegmenterParameters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visionServiceClient) GetObjectPointClouds(ctx context.Context, in *GetObjectPointCloudsRequest, opts ...grpc.CallOption) (*GetObjectPointCloudsResponse, error) {
	out := new(GetObjectPointCloudsResponse)
	err := c.cc.Invoke(ctx, "/proto.api.service.vision.v1.VisionService/GetObjectPointClouds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VisionServiceServer is the server API for VisionService service.
// All implementations must embed UnimplementedVisionServiceServer
// for forward compatibility
type VisionServiceServer interface {
	// GetDetectorNames returns the list of detectors in the registry.
	GetDetectorNames(context.Context, *GetDetectorNamesRequest) (*GetDetectorNamesResponse, error)
	// AddDetector adds a new detector to the registry.
	AddDetector(context.Context, *AddDetectorRequest) (*AddDetectorResponse, error)
	// GetDetections will return a list of detections in the next image given a camera and a detector
	GetDetections(context.Context, *GetDetectionsRequest) (*GetDetectionsResponse, error)
	// GetSegmenterNames returns the list of segmenters in the registry.
	GetSegmenterNames(context.Context, *GetSegmenterNamesRequest) (*GetSegmenterNamesResponse, error)
	// GetSegmenterParameters returns the parameter fields needed for the given segmenter.
	GetSegmenterParameters(context.Context, *GetSegmenterParametersRequest) (*GetSegmenterParametersResponse, error)
	// GetObjectPointClouds returns all the found objects in a pointcloud from a camera of the underlying robot,
	// as well as the 3-vector center of each of the found objects.
	// A specific MIME type can be requested but may not necessarily be the same one returned.
	GetObjectPointClouds(context.Context, *GetObjectPointCloudsRequest) (*GetObjectPointCloudsResponse, error)
	mustEmbedUnimplementedVisionServiceServer()
}

// UnimplementedVisionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVisionServiceServer struct {
}

func (UnimplementedVisionServiceServer) GetDetectorNames(context.Context, *GetDetectorNamesRequest) (*GetDetectorNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetectorNames not implemented")
}
func (UnimplementedVisionServiceServer) AddDetector(context.Context, *AddDetectorRequest) (*AddDetectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDetector not implemented")
}
func (UnimplementedVisionServiceServer) GetDetections(context.Context, *GetDetectionsRequest) (*GetDetectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetections not implemented")
}
func (UnimplementedVisionServiceServer) GetSegmenterNames(context.Context, *GetSegmenterNamesRequest) (*GetSegmenterNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSegmenterNames not implemented")
}
func (UnimplementedVisionServiceServer) GetSegmenterParameters(context.Context, *GetSegmenterParametersRequest) (*GetSegmenterParametersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSegmenterParameters not implemented")
}
func (UnimplementedVisionServiceServer) GetObjectPointClouds(context.Context, *GetObjectPointCloudsRequest) (*GetObjectPointCloudsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectPointClouds not implemented")
}
func (UnimplementedVisionServiceServer) mustEmbedUnimplementedVisionServiceServer() {}

// UnsafeVisionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VisionServiceServer will
// result in compilation errors.
type UnsafeVisionServiceServer interface {
	mustEmbedUnimplementedVisionServiceServer()
}

func RegisterVisionServiceServer(s grpc.ServiceRegistrar, srv VisionServiceServer) {
	s.RegisterService(&VisionService_ServiceDesc, srv)
}

func _VisionService_GetDetectorNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetectorNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).GetDetectorNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.service.vision.v1.VisionService/GetDetectorNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).GetDetectorNames(ctx, req.(*GetDetectorNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_AddDetector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDetectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).AddDetector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.service.vision.v1.VisionService/AddDetector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).AddDetector(ctx, req.(*AddDetectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_GetDetections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).GetDetections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.service.vision.v1.VisionService/GetDetections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).GetDetections(ctx, req.(*GetDetectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_GetSegmenterNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSegmenterNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).GetSegmenterNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.service.vision.v1.VisionService/GetSegmenterNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).GetSegmenterNames(ctx, req.(*GetSegmenterNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_GetSegmenterParameters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSegmenterParametersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).GetSegmenterParameters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.service.vision.v1.VisionService/GetSegmenterParameters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).GetSegmenterParameters(ctx, req.(*GetSegmenterParametersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisionService_GetObjectPointClouds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectPointCloudsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisionServiceServer).GetObjectPointClouds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.service.vision.v1.VisionService/GetObjectPointClouds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisionServiceServer).GetObjectPointClouds(ctx, req.(*GetObjectPointCloudsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VisionService_ServiceDesc is the grpc.ServiceDesc for VisionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VisionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.api.service.vision.v1.VisionService",
	HandlerType: (*VisionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDetectorNames",
			Handler:    _VisionService_GetDetectorNames_Handler,
		},
		{
			MethodName: "AddDetector",
			Handler:    _VisionService_AddDetector_Handler,
		},
		{
			MethodName: "GetDetections",
			Handler:    _VisionService_GetDetections_Handler,
		},
		{
			MethodName: "GetSegmenterNames",
			Handler:    _VisionService_GetSegmenterNames_Handler,
		},
		{
			MethodName: "GetSegmenterParameters",
			Handler:    _VisionService_GetSegmenterParameters_Handler,
		},
		{
			MethodName: "GetObjectPointClouds",
			Handler:    _VisionService_GetObjectPointClouds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api/service/vision/v1/vision.proto",
}