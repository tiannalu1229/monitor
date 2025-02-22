// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/sign.proto

package sign

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Sign service

func NewSignEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Sign service

type SignService interface {
	Sign(ctx context.Context, in *SignRequest, opts ...client.CallOption) (*SignResponse, error)
	CheckSign(ctx context.Context, in *CheckSignRequest, opts ...client.CallOption) (*CheckSignResponse, error)
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (Sign_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Sign_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (Sign_BidiStreamService, error)
}

type signService struct {
	c    client.Client
	name string
}

func NewSignService(name string, c client.Client) SignService {
	return &signService{
		c:    c,
		name: name,
	}
}

func (c *signService) Sign(ctx context.Context, in *SignRequest, opts ...client.CallOption) (*SignResponse, error) {
	req := c.c.NewRequest(c.name, "Sign.Sign", in)
	out := new(SignResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signService) CheckSign(ctx context.Context, in *CheckSignRequest, opts ...client.CallOption) (*CheckSignResponse, error) {
	req := c.c.NewRequest(c.name, "Sign.CheckSign", in)
	out := new(CheckSignResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "Sign.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signService) ClientStream(ctx context.Context, opts ...client.CallOption) (Sign_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "Sign.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &signServiceClientStream{stream}, nil
}

type Sign_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*ClientStreamRequest) error
}

type signServiceClientStream struct {
	stream client.Stream
}

func (x *signServiceClientStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *signServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *signServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *signServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *signServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *signServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *signService) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Sign_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "Sign.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &signServiceServerStream{stream}, nil
}

type Sign_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type signServiceServerStream struct {
	stream client.Stream
}

func (x *signServiceServerStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *signServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *signServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *signServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *signServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *signServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *signService) BidiStream(ctx context.Context, opts ...client.CallOption) (Sign_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "Sign.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &signServiceBidiStream{stream}, nil
}

type Sign_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type signServiceBidiStream struct {
	stream client.Stream
}

func (x *signServiceBidiStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *signServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *signServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *signServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *signServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *signServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *signServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Sign service

type SignHandler interface {
	Sign(context.Context, *SignRequest, *SignResponse) error
	CheckSign(context.Context, *CheckSignRequest, *CheckSignResponse) error
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, Sign_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, Sign_ServerStreamStream) error
	BidiStream(context.Context, Sign_BidiStreamStream) error
}

func RegisterSignHandler(s server.Server, hdlr SignHandler, opts ...server.HandlerOption) error {
	type sign interface {
		Sign(ctx context.Context, in *SignRequest, out *SignResponse) error
		CheckSign(ctx context.Context, in *CheckSignRequest, out *CheckSignResponse) error
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
	}
	type Sign struct {
		sign
	}
	h := &signHandler{hdlr}
	return s.Handle(s.NewHandler(&Sign{h}, opts...))
}

type signHandler struct {
	SignHandler
}

func (h *signHandler) Sign(ctx context.Context, in *SignRequest, out *SignResponse) error {
	return h.SignHandler.Sign(ctx, in, out)
}

func (h *signHandler) CheckSign(ctx context.Context, in *CheckSignRequest, out *CheckSignResponse) error {
	return h.SignHandler.CheckSign(ctx, in, out)
}

func (h *signHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.SignHandler.Call(ctx, in, out)
}

func (h *signHandler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.SignHandler.ClientStream(ctx, &signClientStreamStream{stream})
}

type Sign_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type signClientStreamStream struct {
	stream server.Stream
}

func (x *signClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *signClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *signClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *signClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *signClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *signHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.SignHandler.ServerStream(ctx, m, &signServerStreamStream{stream})
}

type Sign_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type signServerStreamStream struct {
	stream server.Stream
}

func (x *signServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *signServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *signServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *signServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *signServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *signHandler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.SignHandler.BidiStream(ctx, &signBidiStreamStream{stream})
}

type Sign_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type signBidiStreamStream struct {
	stream server.Stream
}

func (x *signBidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *signBidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *signBidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *signBidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *signBidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *signBidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
