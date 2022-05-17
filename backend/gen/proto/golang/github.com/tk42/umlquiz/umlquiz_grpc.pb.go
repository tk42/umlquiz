// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: umlquiz.proto

package umlquiz

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

// UMLQuizLoginServiceClient is the client API for UMLQuizLoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UMLQuizLoginServiceClient interface {
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error)
}

type uMLQuizLoginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUMLQuizLoginServiceClient(cc grpc.ClientConnInterface) UMLQuizLoginServiceClient {
	return &uMLQuizLoginServiceClient{cc}
}

func (c *uMLQuizLoginServiceClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	out := new(GetTokenResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizLoginService/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UMLQuizLoginServiceServer is the server API for UMLQuizLoginService service.
// All implementations should embed UnimplementedUMLQuizLoginServiceServer
// for forward compatibility
type UMLQuizLoginServiceServer interface {
	GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error)
}

// UnimplementedUMLQuizLoginServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUMLQuizLoginServiceServer struct {
}

func (UnimplementedUMLQuizLoginServiceServer) GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}

// UnsafeUMLQuizLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UMLQuizLoginServiceServer will
// result in compilation errors.
type UnsafeUMLQuizLoginServiceServer interface {
	mustEmbedUnimplementedUMLQuizLoginServiceServer()
}

func RegisterUMLQuizLoginServiceServer(s grpc.ServiceRegistrar, srv UMLQuizLoginServiceServer) {
	s.RegisterService(&UMLQuizLoginService_ServiceDesc, srv)
}

func _UMLQuizLoginService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizLoginServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizLoginService/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizLoginServiceServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UMLQuizLoginService_ServiceDesc is the grpc.ServiceDesc for UMLQuizLoginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UMLQuizLoginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "umlquiz.UMLQuizLoginService",
	HandlerType: (*UMLQuizLoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToken",
			Handler:    _UMLQuizLoginService_GetToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "umlquiz.proto",
}

// UMLQuizHelloServiceClient is the client API for UMLQuizHelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UMLQuizHelloServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type uMLQuizHelloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUMLQuizHelloServiceClient(cc grpc.ClientConnInterface) UMLQuizHelloServiceClient {
	return &uMLQuizHelloServiceClient{cc}
}

func (c *uMLQuizHelloServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizHelloService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UMLQuizHelloServiceServer is the server API for UMLQuizHelloService service.
// All implementations should embed UnimplementedUMLQuizHelloServiceServer
// for forward compatibility
type UMLQuizHelloServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

// UnimplementedUMLQuizHelloServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUMLQuizHelloServiceServer struct {
}

func (UnimplementedUMLQuizHelloServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}

// UnsafeUMLQuizHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UMLQuizHelloServiceServer will
// result in compilation errors.
type UnsafeUMLQuizHelloServiceServer interface {
	mustEmbedUnimplementedUMLQuizHelloServiceServer()
}

func RegisterUMLQuizHelloServiceServer(s grpc.ServiceRegistrar, srv UMLQuizHelloServiceServer) {
	s.RegisterService(&UMLQuizHelloService_ServiceDesc, srv)
}

func _UMLQuizHelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizHelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizHelloService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizHelloServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UMLQuizHelloService_ServiceDesc is the grpc.ServiceDesc for UMLQuizHelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UMLQuizHelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "umlquiz.UMLQuizHelloService",
	HandlerType: (*UMLQuizHelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _UMLQuizHelloService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "umlquiz.proto",
}

// UMLQuizUserServiceClient is the client API for UMLQuizUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UMLQuizUserServiceClient interface {
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	FindUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*ErrorResponse, error)
}

type uMLQuizUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUMLQuizUserServiceClient(cc grpc.ClientConnInterface) UMLQuizUserServiceClient {
	return &uMLQuizUserServiceClient{cc}
}

func (c *uMLQuizUserServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizUserService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uMLQuizUserServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizUserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uMLQuizUserServiceClient) FindUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizUserService/FindUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uMLQuizUserServiceClient) DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*ErrorResponse, error) {
	out := new(ErrorResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizUserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UMLQuizUserServiceServer is the server API for UMLQuizUserService service.
// All implementations should embed UnimplementedUMLQuizUserServiceServer
// for forward compatibility
type UMLQuizUserServiceServer interface {
	AddUser(context.Context, *AddUserRequest) (*UserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserResponse, error)
	FindUser(context.Context, *UserRequest) (*UserResponse, error)
	DeleteUser(context.Context, *UserRequest) (*ErrorResponse, error)
}

// UnimplementedUMLQuizUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUMLQuizUserServiceServer struct {
}

func (UnimplementedUMLQuizUserServiceServer) AddUser(context.Context, *AddUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedUMLQuizUserServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUMLQuizUserServiceServer) FindUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUser not implemented")
}
func (UnimplementedUMLQuizUserServiceServer) DeleteUser(context.Context, *UserRequest) (*ErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

// UnsafeUMLQuizUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UMLQuizUserServiceServer will
// result in compilation errors.
type UnsafeUMLQuizUserServiceServer interface {
	mustEmbedUnimplementedUMLQuizUserServiceServer()
}

func RegisterUMLQuizUserServiceServer(s grpc.ServiceRegistrar, srv UMLQuizUserServiceServer) {
	s.RegisterService(&UMLQuizUserService_ServiceDesc, srv)
}

func _UMLQuizUserService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizUserServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizUserService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizUserServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UMLQuizUserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizUserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizUserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizUserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UMLQuizUserService_FindUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizUserServiceServer).FindUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizUserService/FindUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizUserServiceServer).FindUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UMLQuizUserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizUserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizUserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizUserServiceServer).DeleteUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UMLQuizUserService_ServiceDesc is the grpc.ServiceDesc for UMLQuizUserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UMLQuizUserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "umlquiz.UMLQuizUserService",
	HandlerType: (*UMLQuizUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UMLQuizUserService_AddUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UMLQuizUserService_UpdateUser_Handler,
		},
		{
			MethodName: "FindUser",
			Handler:    _UMLQuizUserService_FindUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UMLQuizUserService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "umlquiz.proto",
}

// UMLQuizQuizServiceClient is the client API for UMLQuizQuizService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UMLQuizQuizServiceClient interface {
	AddQuiz(ctx context.Context, in *AddQuizRequest, opts ...grpc.CallOption) (*QuizResponse, error)
	FindQuiz(ctx context.Context, in *FindQuizRequest, opts ...grpc.CallOption) (*QuizResponse, error)
	UpdateQuiz(ctx context.Context, in *UpdateQuizRequest, opts ...grpc.CallOption) (*QuizResponse, error)
	DeleteQuiz(ctx context.Context, in *DeleteQuizRequest, opts ...grpc.CallOption) (*ErrorResponse, error)
	ListQuizzesAll(ctx context.Context, in *ListQuizzesAllRequest, opts ...grpc.CallOption) (*QuizzesResponse, error)
	ListQuizzesByUser(ctx context.Context, in *ListQuizzesByUserRequest, opts ...grpc.CallOption) (*QuizzesResponse, error)
	SolveQuiz(ctx context.Context, in *SolveQuizRequest, opts ...grpc.CallOption) (*SolveResponse, error)
	LikeQuiz(ctx context.Context, in *LikeQuizRequest, opts ...grpc.CallOption) (*ErrorResponse, error)
}

type uMLQuizQuizServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUMLQuizQuizServiceClient(cc grpc.ClientConnInterface) UMLQuizQuizServiceClient {
	return &uMLQuizQuizServiceClient{cc}
}

func (c *uMLQuizQuizServiceClient) AddQuiz(ctx context.Context, in *AddQuizRequest, opts ...grpc.CallOption) (*QuizResponse, error) {
	out := new(QuizResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizQuizService/AddQuiz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uMLQuizQuizServiceClient) FindQuiz(ctx context.Context, in *FindQuizRequest, opts ...grpc.CallOption) (*QuizResponse, error) {
	out := new(QuizResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizQuizService/FindQuiz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uMLQuizQuizServiceClient) UpdateQuiz(ctx context.Context, in *UpdateQuizRequest, opts ...grpc.CallOption) (*QuizResponse, error) {
	out := new(QuizResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizQuizService/UpdateQuiz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uMLQuizQuizServiceClient) DeleteQuiz(ctx context.Context, in *DeleteQuizRequest, opts ...grpc.CallOption) (*ErrorResponse, error) {
	out := new(ErrorResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizQuizService/DeleteQuiz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uMLQuizQuizServiceClient) ListQuizzesAll(ctx context.Context, in *ListQuizzesAllRequest, opts ...grpc.CallOption) (*QuizzesResponse, error) {
	out := new(QuizzesResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizQuizService/ListQuizzesAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uMLQuizQuizServiceClient) ListQuizzesByUser(ctx context.Context, in *ListQuizzesByUserRequest, opts ...grpc.CallOption) (*QuizzesResponse, error) {
	out := new(QuizzesResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizQuizService/ListQuizzesByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uMLQuizQuizServiceClient) SolveQuiz(ctx context.Context, in *SolveQuizRequest, opts ...grpc.CallOption) (*SolveResponse, error) {
	out := new(SolveResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizQuizService/SolveQuiz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uMLQuizQuizServiceClient) LikeQuiz(ctx context.Context, in *LikeQuizRequest, opts ...grpc.CallOption) (*ErrorResponse, error) {
	out := new(ErrorResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizQuizService/LikeQuiz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UMLQuizQuizServiceServer is the server API for UMLQuizQuizService service.
// All implementations should embed UnimplementedUMLQuizQuizServiceServer
// for forward compatibility
type UMLQuizQuizServiceServer interface {
	AddQuiz(context.Context, *AddQuizRequest) (*QuizResponse, error)
	FindQuiz(context.Context, *FindQuizRequest) (*QuizResponse, error)
	UpdateQuiz(context.Context, *UpdateQuizRequest) (*QuizResponse, error)
	DeleteQuiz(context.Context, *DeleteQuizRequest) (*ErrorResponse, error)
	ListQuizzesAll(context.Context, *ListQuizzesAllRequest) (*QuizzesResponse, error)
	ListQuizzesByUser(context.Context, *ListQuizzesByUserRequest) (*QuizzesResponse, error)
	SolveQuiz(context.Context, *SolveQuizRequest) (*SolveResponse, error)
	LikeQuiz(context.Context, *LikeQuizRequest) (*ErrorResponse, error)
}

// UnimplementedUMLQuizQuizServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUMLQuizQuizServiceServer struct {
}

func (UnimplementedUMLQuizQuizServiceServer) AddQuiz(context.Context, *AddQuizRequest) (*QuizResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddQuiz not implemented")
}
func (UnimplementedUMLQuizQuizServiceServer) FindQuiz(context.Context, *FindQuizRequest) (*QuizResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindQuiz not implemented")
}
func (UnimplementedUMLQuizQuizServiceServer) UpdateQuiz(context.Context, *UpdateQuizRequest) (*QuizResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuiz not implemented")
}
func (UnimplementedUMLQuizQuizServiceServer) DeleteQuiz(context.Context, *DeleteQuizRequest) (*ErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuiz not implemented")
}
func (UnimplementedUMLQuizQuizServiceServer) ListQuizzesAll(context.Context, *ListQuizzesAllRequest) (*QuizzesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuizzesAll not implemented")
}
func (UnimplementedUMLQuizQuizServiceServer) ListQuizzesByUser(context.Context, *ListQuizzesByUserRequest) (*QuizzesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuizzesByUser not implemented")
}
func (UnimplementedUMLQuizQuizServiceServer) SolveQuiz(context.Context, *SolveQuizRequest) (*SolveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolveQuiz not implemented")
}
func (UnimplementedUMLQuizQuizServiceServer) LikeQuiz(context.Context, *LikeQuizRequest) (*ErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeQuiz not implemented")
}

// UnsafeUMLQuizQuizServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UMLQuizQuizServiceServer will
// result in compilation errors.
type UnsafeUMLQuizQuizServiceServer interface {
	mustEmbedUnimplementedUMLQuizQuizServiceServer()
}

func RegisterUMLQuizQuizServiceServer(s grpc.ServiceRegistrar, srv UMLQuizQuizServiceServer) {
	s.RegisterService(&UMLQuizQuizService_ServiceDesc, srv)
}

func _UMLQuizQuizService_AddQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddQuizRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizQuizServiceServer).AddQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizQuizService/AddQuiz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizQuizServiceServer).AddQuiz(ctx, req.(*AddQuizRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UMLQuizQuizService_FindQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindQuizRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizQuizServiceServer).FindQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizQuizService/FindQuiz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizQuizServiceServer).FindQuiz(ctx, req.(*FindQuizRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UMLQuizQuizService_UpdateQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuizRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizQuizServiceServer).UpdateQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizQuizService/UpdateQuiz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizQuizServiceServer).UpdateQuiz(ctx, req.(*UpdateQuizRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UMLQuizQuizService_DeleteQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQuizRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizQuizServiceServer).DeleteQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizQuizService/DeleteQuiz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizQuizServiceServer).DeleteQuiz(ctx, req.(*DeleteQuizRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UMLQuizQuizService_ListQuizzesAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQuizzesAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizQuizServiceServer).ListQuizzesAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizQuizService/ListQuizzesAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizQuizServiceServer).ListQuizzesAll(ctx, req.(*ListQuizzesAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UMLQuizQuizService_ListQuizzesByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQuizzesByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizQuizServiceServer).ListQuizzesByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizQuizService/ListQuizzesByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizQuizServiceServer).ListQuizzesByUser(ctx, req.(*ListQuizzesByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UMLQuizQuizService_SolveQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolveQuizRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizQuizServiceServer).SolveQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizQuizService/SolveQuiz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizQuizServiceServer).SolveQuiz(ctx, req.(*SolveQuizRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UMLQuizQuizService_LikeQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeQuizRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizQuizServiceServer).LikeQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizQuizService/LikeQuiz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizQuizServiceServer).LikeQuiz(ctx, req.(*LikeQuizRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UMLQuizQuizService_ServiceDesc is the grpc.ServiceDesc for UMLQuizQuizService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UMLQuizQuizService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "umlquiz.UMLQuizQuizService",
	HandlerType: (*UMLQuizQuizServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddQuiz",
			Handler:    _UMLQuizQuizService_AddQuiz_Handler,
		},
		{
			MethodName: "FindQuiz",
			Handler:    _UMLQuizQuizService_FindQuiz_Handler,
		},
		{
			MethodName: "UpdateQuiz",
			Handler:    _UMLQuizQuizService_UpdateQuiz_Handler,
		},
		{
			MethodName: "DeleteQuiz",
			Handler:    _UMLQuizQuizService_DeleteQuiz_Handler,
		},
		{
			MethodName: "ListQuizzesAll",
			Handler:    _UMLQuizQuizService_ListQuizzesAll_Handler,
		},
		{
			MethodName: "ListQuizzesByUser",
			Handler:    _UMLQuizQuizService_ListQuizzesByUser_Handler,
		},
		{
			MethodName: "SolveQuiz",
			Handler:    _UMLQuizQuizService_SolveQuiz_Handler,
		},
		{
			MethodName: "LikeQuiz",
			Handler:    _UMLQuizQuizService_LikeQuiz_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "umlquiz.proto",
}

// UMLQuizReportServiceClient is the client API for UMLQuizReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UMLQuizReportServiceClient interface {
	AddReport(ctx context.Context, in *AddReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
	FindReports(ctx context.Context, in *FindReportsRequest, opts ...grpc.CallOption) (*ReportsResponse, error)
	UpdateReport(ctx context.Context, in *UpdateReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
	DeleteReport(ctx context.Context, in *DeleteReportRequest, opts ...grpc.CallOption) (*ErrorResponse, error)
}

type uMLQuizReportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUMLQuizReportServiceClient(cc grpc.ClientConnInterface) UMLQuizReportServiceClient {
	return &uMLQuizReportServiceClient{cc}
}

func (c *uMLQuizReportServiceClient) AddReport(ctx context.Context, in *AddReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizReportService/AddReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uMLQuizReportServiceClient) FindReports(ctx context.Context, in *FindReportsRequest, opts ...grpc.CallOption) (*ReportsResponse, error) {
	out := new(ReportsResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizReportService/FindReports", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uMLQuizReportServiceClient) UpdateReport(ctx context.Context, in *UpdateReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizReportService/UpdateReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uMLQuizReportServiceClient) DeleteReport(ctx context.Context, in *DeleteReportRequest, opts ...grpc.CallOption) (*ErrorResponse, error) {
	out := new(ErrorResponse)
	err := c.cc.Invoke(ctx, "/umlquiz.UMLQuizReportService/DeleteReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UMLQuizReportServiceServer is the server API for UMLQuizReportService service.
// All implementations should embed UnimplementedUMLQuizReportServiceServer
// for forward compatibility
type UMLQuizReportServiceServer interface {
	AddReport(context.Context, *AddReportRequest) (*ReportResponse, error)
	FindReports(context.Context, *FindReportsRequest) (*ReportsResponse, error)
	UpdateReport(context.Context, *UpdateReportRequest) (*ReportResponse, error)
	DeleteReport(context.Context, *DeleteReportRequest) (*ErrorResponse, error)
}

// UnimplementedUMLQuizReportServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUMLQuizReportServiceServer struct {
}

func (UnimplementedUMLQuizReportServiceServer) AddReport(context.Context, *AddReportRequest) (*ReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReport not implemented")
}
func (UnimplementedUMLQuizReportServiceServer) FindReports(context.Context, *FindReportsRequest) (*ReportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindReports not implemented")
}
func (UnimplementedUMLQuizReportServiceServer) UpdateReport(context.Context, *UpdateReportRequest) (*ReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReport not implemented")
}
func (UnimplementedUMLQuizReportServiceServer) DeleteReport(context.Context, *DeleteReportRequest) (*ErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReport not implemented")
}

// UnsafeUMLQuizReportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UMLQuizReportServiceServer will
// result in compilation errors.
type UnsafeUMLQuizReportServiceServer interface {
	mustEmbedUnimplementedUMLQuizReportServiceServer()
}

func RegisterUMLQuizReportServiceServer(s grpc.ServiceRegistrar, srv UMLQuizReportServiceServer) {
	s.RegisterService(&UMLQuizReportService_ServiceDesc, srv)
}

func _UMLQuizReportService_AddReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizReportServiceServer).AddReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizReportService/AddReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizReportServiceServer).AddReport(ctx, req.(*AddReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UMLQuizReportService_FindReports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindReportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizReportServiceServer).FindReports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizReportService/FindReports",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizReportServiceServer).FindReports(ctx, req.(*FindReportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UMLQuizReportService_UpdateReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizReportServiceServer).UpdateReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizReportService/UpdateReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizReportServiceServer).UpdateReport(ctx, req.(*UpdateReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UMLQuizReportService_DeleteReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMLQuizReportServiceServer).DeleteReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umlquiz.UMLQuizReportService/DeleteReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMLQuizReportServiceServer).DeleteReport(ctx, req.(*DeleteReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UMLQuizReportService_ServiceDesc is the grpc.ServiceDesc for UMLQuizReportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UMLQuizReportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "umlquiz.UMLQuizReportService",
	HandlerType: (*UMLQuizReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddReport",
			Handler:    _UMLQuizReportService_AddReport_Handler,
		},
		{
			MethodName: "FindReports",
			Handler:    _UMLQuizReportService_FindReports_Handler,
		},
		{
			MethodName: "UpdateReport",
			Handler:    _UMLQuizReportService_UpdateReport_Handler,
		},
		{
			MethodName: "DeleteReport",
			Handler:    _UMLQuizReportService_DeleteReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "umlquiz.proto",
}
