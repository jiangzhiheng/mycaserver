// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: pkg/grpc/service.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CertificateServiceClient is the client API for CertificateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertificateServiceClient interface {
	CsrTemplate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CertificateSignRequest, error)
	SignCsr(ctx context.Context, in *CertificateSignRequest, opts ...grpc.CallOption) (*SignResponse, error)
	GetCert(ctx context.Context, in *FileIdentifer, opts ...grpc.CallOption) (*FileStream, error)
	GetKey(ctx context.Context, in *FileIdentifer, opts ...grpc.CallOption) (*FileStream, error)
}

type certificateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateServiceClient(cc grpc.ClientConnInterface) CertificateServiceClient {
	return &certificateServiceClient{cc}
}

func (c *certificateServiceClient) CsrTemplate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CertificateSignRequest, error) {
	out := new(CertificateSignRequest)
	err := c.cc.Invoke(ctx, "/grpc.CertificateService/CsrTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) SignCsr(ctx context.Context, in *CertificateSignRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := c.cc.Invoke(ctx, "/grpc.CertificateService/SignCsr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) GetCert(ctx context.Context, in *FileIdentifer, opts ...grpc.CallOption) (*FileStream, error) {
	out := new(FileStream)
	err := c.cc.Invoke(ctx, "/grpc.CertificateService/GetCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) GetKey(ctx context.Context, in *FileIdentifer, opts ...grpc.CallOption) (*FileStream, error) {
	out := new(FileStream)
	err := c.cc.Invoke(ctx, "/grpc.CertificateService/GetKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateServiceServer is the server API for CertificateService service.
// All implementations must embed UnimplementedCertificateServiceServer
// for forward compatibility
type CertificateServiceServer interface {
	CsrTemplate(context.Context, *emptypb.Empty) (*CertificateSignRequest, error)
	SignCsr(context.Context, *CertificateSignRequest) (*SignResponse, error)
	GetCert(context.Context, *FileIdentifer) (*FileStream, error)
	GetKey(context.Context, *FileIdentifer) (*FileStream, error)
	mustEmbedUnimplementedCertificateServiceServer()
}

// UnimplementedCertificateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCertificateServiceServer struct {
}

func (UnimplementedCertificateServiceServer) CsrTemplate(context.Context, *emptypb.Empty) (*CertificateSignRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CsrTemplate not implemented")
}
func (UnimplementedCertificateServiceServer) SignCsr(context.Context, *CertificateSignRequest) (*SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignCsr not implemented")
}
func (UnimplementedCertificateServiceServer) GetCert(context.Context, *FileIdentifer) (*FileStream, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCert not implemented")
}
func (UnimplementedCertificateServiceServer) GetKey(context.Context, *FileIdentifer) (*FileStream, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKey not implemented")
}
func (UnimplementedCertificateServiceServer) mustEmbedUnimplementedCertificateServiceServer() {}

// UnsafeCertificateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificateServiceServer will
// result in compilation errors.
type UnsafeCertificateServiceServer interface {
	mustEmbedUnimplementedCertificateServiceServer()
}

func RegisterCertificateServiceServer(s grpc.ServiceRegistrar, srv CertificateServiceServer) {
	s.RegisterService(&CertificateService_ServiceDesc, srv)
}

func _CertificateService_CsrTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).CsrTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.CertificateService/CsrTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).CsrTemplate(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_SignCsr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateSignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).SignCsr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.CertificateService/SignCsr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).SignCsr(ctx, req.(*CertificateSignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_GetCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileIdentifer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).GetCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.CertificateService/GetCert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).GetCert(ctx, req.(*FileIdentifer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_GetKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileIdentifer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).GetKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.CertificateService/GetKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).GetKey(ctx, req.(*FileIdentifer))
	}
	return interceptor(ctx, in, info, handler)
}

// CertificateService_ServiceDesc is the grpc.ServiceDesc for CertificateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertificateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.CertificateService",
	HandlerType: (*CertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CsrTemplate",
			Handler:    _CertificateService_CsrTemplate_Handler,
		},
		{
			MethodName: "SignCsr",
			Handler:    _CertificateService_SignCsr_Handler,
		},
		{
			MethodName: "GetCert",
			Handler:    _CertificateService_GetCert_Handler,
		},
		{
			MethodName: "GetKey",
			Handler:    _CertificateService_GetKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/service.proto",
}
