// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: image.proto

package image

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ImageService_UploadImage_FullMethodName = "/image.ImageService/UploadImage"
	ImageService_DeleteImage_FullMethodName = "/image.ImageService/DeleteImage"
)

// ImageServiceClient is the client API for ImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageServiceClient interface {
	UploadImage(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
	DeleteImage(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type imageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageServiceClient(cc grpc.ClientConnInterface) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) UploadImage(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, ImageService_UploadImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) DeleteImage(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, ImageService_DeleteImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServiceServer is the server API for ImageService service.
// All implementations must embed UnimplementedImageServiceServer
// for forward compatibility.
type ImageServiceServer interface {
	UploadImage(context.Context, *UploadRequest) (*UploadResponse, error)
	DeleteImage(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedImageServiceServer()
}

// UnimplementedImageServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedImageServiceServer struct{}

func (UnimplementedImageServiceServer) UploadImage(context.Context, *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (UnimplementedImageServiceServer) DeleteImage(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImage not implemented")
}
func (UnimplementedImageServiceServer) mustEmbedUnimplementedImageServiceServer() {}
func (UnimplementedImageServiceServer) testEmbeddedByValue()                      {}

// UnsafeImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServiceServer will
// result in compilation errors.
type UnsafeImageServiceServer interface {
	mustEmbedUnimplementedImageServiceServer()
}

func RegisterImageServiceServer(s grpc.ServiceRegistrar, srv ImageServiceServer) {
	// If the following call pancis, it indicates UnimplementedImageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ImageService_ServiceDesc, srv)
}

func _ImageService_UploadImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).UploadImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_UploadImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).UploadImage(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_DeleteImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).DeleteImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_DeleteImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).DeleteImage(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageService_ServiceDesc is the grpc.ServiceDesc for ImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "image.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadImage",
			Handler:    _ImageService_UploadImage_Handler,
		},
		{
			MethodName: "DeleteImage",
			Handler:    _ImageService_DeleteImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "image.proto",
}
