// Code generated by protoc-gen-go. DO NOT EDIT.
// source: automate-gateway/api/iam/v2/tokens.proto

package v2

import (
	context "context"
	fmt "fmt"
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2/request"
	response "github.com/chef/automate/components/automate-gateway/api/iam/v2/response"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("automate-gateway/api/iam/v2/tokens.proto", fileDescriptor_f530e76e6157e08e)
}

var fileDescriptor_f530e76e6157e08e = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4f, 0x6b, 0xd4, 0x40,
	0x18, 0xc6, 0x49, 0x29, 0x4b, 0x9d, 0x4a, 0xb5, 0x53, 0x84, 0x10, 0x14, 0xc2, 0x80, 0x5a, 0x97,
	0x26, 0xb1, 0x2b, 0x5e, 0xf6, 0x22, 0xab, 0x82, 0xa0, 0xa2, 0xa0, 0x15, 0xc1, 0x52, 0x70, 0x9a,
	0xbc, 0xa6, 0x53, 0x93, 0xcc, 0x34, 0x33, 0x69, 0x2d, 0xa5, 0x2a, 0x0a, 0x1e, 0x7a, 0xb3, 0x9e,
	0x04, 0x8f, 0x1e, 0x05, 0x4f, 0x3d, 0xf9, 0x2d, 0xac, 0x1f, 0xc1, 0xab, 0xf8, 0x15, 0x24, 0x13,
	0xb7, 0x49, 0xdc, 0x3f, 0x8d, 0xa7, 0x90, 0x97, 0xe7, 0x9d, 0x79, 0xde, 0x67, 0x7e, 0x33, 0x68,
	0x9e, 0x66, 0x8a, 0xc7, 0x54, 0x81, 0x13, 0x52, 0x05, 0x5b, 0x74, 0xdb, 0xa3, 0x82, 0x79, 0x8c,
	0xc6, 0xde, 0x66, 0xc7, 0x53, 0xfc, 0x39, 0x24, 0xd2, 0x15, 0x29, 0x57, 0x1c, 0x9b, 0xfe, 0x1a,
	0x3c, 0x73, 0xfb, 0x72, 0x97, 0x0a, 0xe6, 0x32, 0x1a, 0xbb, 0x9b, 0x1d, 0xeb, 0x6c, 0xc8, 0x79,
	0x18, 0x81, 0xee, 0xa4, 0x49, 0xc2, 0x15, 0x55, 0x8c, 0xf7, 0xfb, 0xac, 0x05, 0xfd, 0xf1, 0x9d,
	0x10, 0x12, 0x47, 0x6e, 0xd1, 0x30, 0x84, 0xd4, 0xe3, 0x42, 0x2b, 0x86, 0xa8, 0x2f, 0x8f, 0xf3,
	0x93, 0xc2, 0x46, 0x06, 0x52, 0xd5, 0x7c, 0x59, 0x8b, 0xe3, 0x3b, 0xa4, 0xe0, 0x89, 0x84, 0x7a,
	0xcb, 0xd5, 0xb2, 0x25, 0x15, 0xbe, 0x57, 0x31, 0x28, 0x78, 0xc4, 0xfc, 0x6d, 0xdd, 0x3d, 0xe0,
	0xad, 0x73, 0x78, 0x02, 0xb5, 0x96, 0xf4, 0x3a, 0xf8, 0xfb, 0x04, 0x9a, 0xbe, 0x91, 0x02, 0x55,
	0xa0, 0x0b, 0x78, 0xde, 0x1d, 0x95, 0x8e, 0x5b, 0x91, 0x3d, 0x80, 0x0d, 0xeb, 0x52, 0x43, 0xa5,
	0x14, 0xe4, 0x97, 0xb1, 0xdf, 0xfb, 0x6c, 0xa0, 0x56, 0xe1, 0x7d, 0xfd, 0xa3, 0x81, 0x66, 0x5e,
	0x38, 0x3e, 0x0f, 0xc0, 0x91, 0x34, 0x16, 0x11, 0x48, 0xfc, 0xce, 0xe8, 0xbc, 0x35, 0xd0, 0x6b,
	0xa3, 0xfd, 0x12, 0xcd, 0xa0, 0xc9, 0x88, 0x26, 0x21, 0x6e, 0x59, 0x93, 0xb7, 0x1f, 0xde, 0xbf,
	0x87, 0x22, 0xd4, 0x92, 0x3c, 0x4b, 0x7d, 0xc0, 0xab, 0xd6, 0xd3, 0x1d, 0x92, 0xd0, 0x18, 0x48,
	0xd7, 0x26, 0x7a, 0x2d, 0x7b, 0x91, 0x2c, 0xd8, 0x84, 0x05, 0x47, 0x05, 0x47, 0x17, 0xa8, 0xaf,
	0xd8, 0x66, 0xae, 0x52, 0x69, 0x06, 0x0b, 0x36, 0x11, 0x29, 0x5f, 0x07, 0x5f, 0x49, 0xd2, 0xb5,
	0x97, 0x09, 0x50, 0xa9, 0x9c, 0x14, 0x42, 0xc6, 0x93, 0x5c, 0xbb, 0x05, 0xe5, 0xef, 0xca, 0xee,
	0xde, 0x81, 0x79, 0x12, 0x21, 0x46, 0xe3, 0x6e, 0xe1, 0x75, 0xef, 0xc0, 0x9c, 0xc3, 0xb3, 0xe5,
	0x7f, 0xd7, 0xd7, 0xa3, 0xbd, 0x39, 0xfc, 0xf9, 0x61, 0xc2, 0x24, 0x73, 0xf9, 0x09, 0xc9, 0x3a,
	0x64, 0x5d, 0xa3, 0x8d, 0xbf, 0x1a, 0x68, 0xea, 0x16, 0xa8, 0x22, 0xd1, 0xf3, 0xa3, 0x73, 0xea,
	0x6b, 0xf2, 0x38, 0x2f, 0x34, 0x91, 0x49, 0x41, 0x96, 0xf6, 0x7b, 0x53, 0xfd, 0x24, 0xf7, 0x0e,
	0xcc, 0x59, 0x74, 0xaa, 0xe2, 0x6e, 0x87, 0x05, 0xf9, 0x08, 0xa7, 0xf1, 0x4c, 0xa5, 0x18, 0x82,
	0xd2, 0x7e, 0x2d, 0x6c, 0x0e, 0xf1, 0xeb, 0xe5, 0x3d, 0xf8, 0xc7, 0x04, 0x9a, 0x7e, 0x24, 0x82,
	0x26, 0x18, 0x54, 0x64, 0xc7, 0x60, 0x50, 0x53, 0x4a, 0x41, 0x7e, 0x1b, 0xfb, 0xbd, 0x4f, 0x25,
	0x06, 0xef, 0x07, 0x31, 0x78, 0xd5, 0xd9, 0x45, 0x3b, 0xed, 0xed, 0x01, 0x06, 0xc2, 0x23, 0x06,
	0x56, 0xac, 0xe5, 0x92, 0x81, 0x4c, 0x6f, 0x10, 0xd8, 0x05, 0x0b, 0xba, 0xfa, 0x9f, 0xa7, 0x2f,
	0x79, 0xa6, 0xd6, 0x6a, 0xc7, 0x3f, 0x34, 0xd0, 0x3a, 0x03, 0xc5, 0xb6, 0x3a, 0xd3, 0x73, 0xd6,
	0xc8, 0x4c, 0x73, 0x10, 0xbe, 0x19, 0x68, 0xfa, 0x26, 0x44, 0xd0, 0x20, 0xd6, 0x8a, 0xec, 0x98,
	0x58, 0x6b, 0x4a, 0x29, 0xc8, 0xe3, 0x06, 0x44, 0xd4, 0x07, 0x08, 0xf4, 0x0a, 0x05, 0x14, 0xed,
	0xd1, 0x50, 0x7c, 0x31, 0x10, 0xba, 0xcb, 0xa4, 0xfa, 0xfb, 0x54, 0x5c, 0x1c, 0x6d, 0xa9, 0x54,
	0xe5, 0xde, 0xe7, 0x9b, 0x09, 0xa5, 0x20, 0x77, 0xea, 0xd6, 0xff, 0xbd, 0x7a, 0xb3, 0xb8, 0x3a,
	0x4a, 0xc4, 0x64, 0x01, 0xf2, 0x19, 0x3c, 0xec, 0xe2, 0x5d, 0xef, 0x3d, 0xb9, 0x16, 0x32, 0xb5,
	0x96, 0xad, 0xba, 0x3e, 0x8f, 0xbd, 0xdc, 0x82, 0xd7, 0xb7, 0xe0, 0xf9, 0x3c, 0x16, 0x3c, 0x81,
	0x44, 0x49, 0x6f, 0xcc, 0x2b, 0xbb, 0xda, 0xd2, 0xef, 0xe3, 0x95, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xbd, 0x45, 0x57, 0xea, 0x4d, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TokensClient is the client API for Tokens service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokensClient interface {
	//
	//Creates a token
	//
	//Creates a token.
	//Active defaults to true when not specified.
	//Value is auto-generated when not specified.
	//
	//Note that this creates *non-admin* tokens that may then be assigned permissions via policies just like users or teams (unless you have already created policies that encompass all tokens using `tokens:*``).
	//
	//You cannot create admin tokens via the REST API.
	//Admin tokens can only be created by specifying the `--admin` flag to this chef-automate sub-command:
	//```
	//chef-automate iam token create <your-token-name> --admin`
	//```
	//
	//Authorization Action:
	//```
	//iam:tokens:create
	//```
	CreateToken(ctx context.Context, in *request.CreateTokenReq, opts ...grpc.CallOption) (*response.CreateTokenResp, error)
	//
	//Gets a token
	//
	//Returns the details for a token.
	//
	//Authorization Action:
	//```
	//iam:tokens:get
	//```
	GetToken(ctx context.Context, in *request.GetTokenReq, opts ...grpc.CallOption) (*response.GetTokenResp, error)
	//
	//Updates a token
	//
	//This operation overwrites all fields excepting ID, timestamps, and value,
	//including those omitted from the request, so be sure to specify all properties.
	//Properties that you do not include are reset to empty values.
	//
	//Authorization Action:
	//```
	//iam:tokens:update
	//```
	UpdateToken(ctx context.Context, in *request.UpdateTokenReq, opts ...grpc.CallOption) (*response.UpdateTokenResp, error)
	//
	//Deletes a token
	//
	//Deletes a token and remove it from any policies.
	//
	//Authorization Action:
	//```
	//iam:tokens:delete
	//```
	DeleteToken(ctx context.Context, in *request.DeleteTokenReq, opts ...grpc.CallOption) (*response.DeleteTokenResp, error)
	//
	//Lists all tokens
	//
	//Lists all tokens, both admin and non-admin.
	//
	//Authorization Action:
	//```
	//iam:tokens:list
	//```
	ListTokens(ctx context.Context, in *request.ListTokensReq, opts ...grpc.CallOption) (*response.ListTokensResp, error)
}

type tokensClient struct {
	cc grpc.ClientConnInterface
}

func NewTokensClient(cc grpc.ClientConnInterface) TokensClient {
	return &tokensClient{cc}
}

func (c *tokensClient) CreateToken(ctx context.Context, in *request.CreateTokenReq, opts ...grpc.CallOption) (*response.CreateTokenResp, error) {
	out := new(response.CreateTokenResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Tokens/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensClient) GetToken(ctx context.Context, in *request.GetTokenReq, opts ...grpc.CallOption) (*response.GetTokenResp, error) {
	out := new(response.GetTokenResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Tokens/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensClient) UpdateToken(ctx context.Context, in *request.UpdateTokenReq, opts ...grpc.CallOption) (*response.UpdateTokenResp, error) {
	out := new(response.UpdateTokenResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Tokens/UpdateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensClient) DeleteToken(ctx context.Context, in *request.DeleteTokenReq, opts ...grpc.CallOption) (*response.DeleteTokenResp, error) {
	out := new(response.DeleteTokenResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Tokens/DeleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensClient) ListTokens(ctx context.Context, in *request.ListTokensReq, opts ...grpc.CallOption) (*response.ListTokensResp, error) {
	out := new(response.ListTokensResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Tokens/ListTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokensServer is the server API for Tokens service.
type TokensServer interface {
	//
	//Creates a token
	//
	//Creates a token.
	//Active defaults to true when not specified.
	//Value is auto-generated when not specified.
	//
	//Note that this creates *non-admin* tokens that may then be assigned permissions via policies just like users or teams (unless you have already created policies that encompass all tokens using `tokens:*``).
	//
	//You cannot create admin tokens via the REST API.
	//Admin tokens can only be created by specifying the `--admin` flag to this chef-automate sub-command:
	//```
	//chef-automate iam token create <your-token-name> --admin`
	//```
	//
	//Authorization Action:
	//```
	//iam:tokens:create
	//```
	CreateToken(context.Context, *request.CreateTokenReq) (*response.CreateTokenResp, error)
	//
	//Gets a token
	//
	//Returns the details for a token.
	//
	//Authorization Action:
	//```
	//iam:tokens:get
	//```
	GetToken(context.Context, *request.GetTokenReq) (*response.GetTokenResp, error)
	//
	//Updates a token
	//
	//This operation overwrites all fields excepting ID, timestamps, and value,
	//including those omitted from the request, so be sure to specify all properties.
	//Properties that you do not include are reset to empty values.
	//
	//Authorization Action:
	//```
	//iam:tokens:update
	//```
	UpdateToken(context.Context, *request.UpdateTokenReq) (*response.UpdateTokenResp, error)
	//
	//Deletes a token
	//
	//Deletes a token and remove it from any policies.
	//
	//Authorization Action:
	//```
	//iam:tokens:delete
	//```
	DeleteToken(context.Context, *request.DeleteTokenReq) (*response.DeleteTokenResp, error)
	//
	//Lists all tokens
	//
	//Lists all tokens, both admin and non-admin.
	//
	//Authorization Action:
	//```
	//iam:tokens:list
	//```
	ListTokens(context.Context, *request.ListTokensReq) (*response.ListTokensResp, error)
}

// UnimplementedTokensServer can be embedded to have forward compatible implementations.
type UnimplementedTokensServer struct {
}

func (*UnimplementedTokensServer) CreateToken(ctx context.Context, req *request.CreateTokenReq) (*response.CreateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (*UnimplementedTokensServer) GetToken(ctx context.Context, req *request.GetTokenReq) (*response.GetTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (*UnimplementedTokensServer) UpdateToken(ctx context.Context, req *request.UpdateTokenReq) (*response.UpdateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateToken not implemented")
}
func (*UnimplementedTokensServer) DeleteToken(ctx context.Context, req *request.DeleteTokenReq) (*response.DeleteTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}
func (*UnimplementedTokensServer) ListTokens(ctx context.Context, req *request.ListTokensReq) (*response.ListTokensResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTokens not implemented")
}

func RegisterTokensServer(s *grpc.Server, srv TokensServer) {
	s.RegisterService(&_Tokens_serviceDesc, srv)
}

func _Tokens_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Tokens/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).CreateToken(ctx, req.(*request.CreateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tokens_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Tokens/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).GetToken(ctx, req.(*request.GetTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tokens_UpdateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).UpdateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Tokens/UpdateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).UpdateToken(ctx, req.(*request.UpdateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tokens_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Tokens/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).DeleteToken(ctx, req.(*request.DeleteTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tokens_ListTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListTokensReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).ListTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Tokens/ListTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).ListTokens(ctx, req.(*request.ListTokensReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tokens_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.iam.v2.Tokens",
	HandlerType: (*TokensServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _Tokens_CreateToken_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _Tokens_GetToken_Handler,
		},
		{
			MethodName: "UpdateToken",
			Handler:    _Tokens_UpdateToken_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _Tokens_DeleteToken_Handler,
		},
		{
			MethodName: "ListTokens",
			Handler:    _Tokens_ListTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "automate-gateway/api/iam/v2/tokens.proto",
}
