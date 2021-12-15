// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/udsver_v1/udsver.proto

/*
Package udsver_v1 is a generated protocol buffer package.

It is generated from these files:
	protos/udsver_v1/udsver.proto

It has these top-level messages:
	Request
	Response
*/
package udsver_v1

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Response_Status_Code int32

const (
	// https://cloud.google.com/appengine/docs/admin-api/reference/rpc/google.rpc#google.rpc.Code
	Response_Status_OK                  Response_Status_Code = 0
	Response_Status_CANCELLED           Response_Status_Code = 1
	Response_Status_UNKNOWN             Response_Status_Code = 2
	Response_Status_INVALID_ARGUMENT    Response_Status_Code = 3
	Response_Status_DEADLINE_EXCEEDED   Response_Status_Code = 4
	Response_Status_NOT_FOUND           Response_Status_Code = 5
	Response_Status_ALREADY_EXISTS      Response_Status_Code = 6
	Response_Status_PERMISSION_DENIED   Response_Status_Code = 7
	Response_Status_UNAUTHENTICATED     Response_Status_Code = 8
	Response_Status_RESOURCE_EXHAUSTED  Response_Status_Code = 9
	Response_Status_FAILED_PRECONDITION Response_Status_Code = 10
	Response_Status_ABORTED             Response_Status_Code = 11
	Response_Status_OUT_OF_RANGE        Response_Status_Code = 12
	Response_Status_UNIMPLEMENTED       Response_Status_Code = 13
	Response_Status_INTERNAL            Response_Status_Code = 14
	Response_Status_UNAVAILABLE         Response_Status_Code = 15
	Response_Status_DATA_LOSS           Response_Status_Code = 16
)

var Response_Status_Code_name = map[int32]string{
	0:  "OK",
	1:  "CANCELLED",
	2:  "UNKNOWN",
	3:  "INVALID_ARGUMENT",
	4:  "DEADLINE_EXCEEDED",
	5:  "NOT_FOUND",
	6:  "ALREADY_EXISTS",
	7:  "PERMISSION_DENIED",
	8:  "UNAUTHENTICATED",
	9:  "RESOURCE_EXHAUSTED",
	10: "FAILED_PRECONDITION",
	11: "ABORTED",
	12: "OUT_OF_RANGE",
	13: "UNIMPLEMENTED",
	14: "INTERNAL",
	15: "UNAVAILABLE",
	16: "DATA_LOSS",
}
var Response_Status_Code_value = map[string]int32{
	"OK":                  0,
	"CANCELLED":           1,
	"UNKNOWN":             2,
	"INVALID_ARGUMENT":    3,
	"DEADLINE_EXCEEDED":   4,
	"NOT_FOUND":           5,
	"ALREADY_EXISTS":      6,
	"PERMISSION_DENIED":   7,
	"UNAUTHENTICATED":     8,
	"RESOURCE_EXHAUSTED":  9,
	"FAILED_PRECONDITION": 10,
	"ABORTED":             11,
	"OUT_OF_RANGE":        12,
	"UNIMPLEMENTED":       13,
	"INTERNAL":            14,
	"UNAVAILABLE":         15,
	"DATA_LOSS":           16,
}

func (x Response_Status_Code) String() string {
	return proto.EnumName(Response_Status_Code_name, int32(x))
}
func (Response_Status_Code) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0, 0} }

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Status *Response_Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetStatus() *Response_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// TODO Maybe replace with googleapis
type Response_Status struct {
	Code    Response_Status_Code `protobuf:"varint,1,opt,name=code,enum=udsver.v1.Response_Status_Code" json:"code,omitempty"`
	Message string               `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Response_Status) Reset()                    { *m = Response_Status{} }
func (m *Response_Status) String() string            { return proto.CompactTextString(m) }
func (*Response_Status) ProtoMessage()               {}
func (*Response_Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Status) GetCode() Response_Status_Code {
	if m != nil {
		return m.Code
	}
	return Response_Status_OK
}

func (m *Response_Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "udsver.v1.Request")
	proto.RegisterType((*Response)(nil), "udsver.v1.Response")
	proto.RegisterType((*Response_Status)(nil), "udsver.v1.Response.Status")
	proto.RegisterEnum("udsver.v1.Response_Status_Code", Response_Status_Code_name, Response_Status_Code_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Verify service

type VerifyClient interface {
	Check(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type verifyClient struct {
	cc *grpc.ClientConn
}

func NewVerifyClient(cc *grpc.ClientConn) VerifyClient {
	return &verifyClient{cc}
}

func (c *verifyClient) Check(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/udsver.v1.Verify/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Verify service

type VerifyServer interface {
	Check(context.Context, *Request) (*Response, error)
}

func RegisterVerifyServer(s *grpc.Server, srv VerifyServer) {
	s.RegisterService(&_Verify_serviceDesc, srv)
}

func _Verify_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/udsver.v1.Verify/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyServer).Check(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Verify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "udsver.v1.Verify",
	HandlerType: (*VerifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Verify_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/udsver_v1/udsver.proto",
}

func init() { proto.RegisterFile("protos/udsver_v1/udsver.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xdb, 0x6e, 0x13, 0x4d,
	0x0c, 0xc7, 0xbf, 0x1c, 0xba, 0x49, 0x9c, 0xd3, 0xd4, 0xf9, 0x80, 0x28, 0x52, 0x05, 0xca, 0x15,
	0x57, 0x81, 0xa6, 0x77, 0xdc, 0x4d, 0x77, 0x9c, 0x76, 0xd4, 0x89, 0x27, 0x9a, 0xdd, 0x0d, 0xe5,
	0x6a, 0x15, 0x9a, 0xe1, 0x20, 0xd4, 0x6e, 0xc9, 0x26, 0x91, 0x78, 0x0c, 0x9e, 0x86, 0xf7, 0xe0,
	0x89, 0xd0, 0x6e, 0x53, 0x24, 0x24, 0xb8, 0xb3, 0xff, 0xfa, 0xf9, 0x6f, 0x5b, 0x36, 0x9c, 0xdc,
	0x6f, 0xb2, 0x6d, 0x96, 0xbf, 0xda, 0xad, 0xf3, 0xbd, 0xdf, 0xa4, 0xfb, 0xd3, 0x43, 0x34, 0x29,
	0x75, 0x6c, 0x1d, 0xb2, 0xfd, 0xe9, 0xf8, 0x04, 0x1a, 0xce, 0x7f, 0xdd, 0xf9, 0x7c, 0x8b, 0x08,
	0xf5, 0xbb, 0xd5, 0xad, 0x1f, 0x56, 0x5e, 0x54, 0x5e, 0xb6, 0x5c, 0x19, 0x8f, 0x7f, 0xd6, 0xa0,
	0xe9, 0x7c, 0x7e, 0x9f, 0xdd, 0xe5, 0x1e, 0xa7, 0x10, 0xe4, 0xdb, 0xd5, 0x76, 0x97, 0x97, 0x48,
	0x7b, 0x3a, 0x9a, 0xfc, 0xf6, 0x99, 0x3c, 0x42, 0x93, 0xa8, 0x24, 0xdc, 0x81, 0x1c, 0x7d, 0xaf,
	0x41, 0xf0, 0x20, 0xe1, 0x19, 0xd4, 0x6f, 0xb2, 0xf5, 0x83, 0x7f, 0x6f, 0xfa, 0xfc, 0xdf, 0xc5,
	0x93, 0x30, 0x5b, 0x7b, 0x57, 0xc2, 0x38, 0x84, 0xc6, 0xad, 0xcf, 0xf3, 0xd5, 0x47, 0x3f, 0xac,
	0x96, 0x73, 0x3d, 0xa6, 0xe3, 0x1f, 0x55, 0xa8, 0x17, 0x20, 0x06, 0x50, 0xb5, 0x57, 0xe2, 0x3f,
	0xec, 0x42, 0x2b, 0x94, 0x1c, 0x92, 0x31, 0xa4, 0x44, 0x05, 0xdb, 0xd0, 0x48, 0xf8, 0x8a, 0xed,
	0x5b, 0x16, 0x55, 0xfc, 0x1f, 0x84, 0xe6, 0xa5, 0x34, 0x5a, 0xa5, 0xd2, 0x5d, 0x24, 0x73, 0xe2,
	0x58, 0xd4, 0xf0, 0x09, 0x1c, 0x2b, 0x92, 0xca, 0x68, 0xa6, 0x94, 0xae, 0x43, 0x22, 0x45, 0x4a,
	0xd4, 0x0b, 0x23, 0xb6, 0x71, 0x3a, 0xb3, 0x09, 0x2b, 0x71, 0x84, 0x08, 0x3d, 0x69, 0x1c, 0x49,
	0xf5, 0x2e, 0xa5, 0x6b, 0x1d, 0xc5, 0x91, 0x08, 0x8a, 0xca, 0x05, 0xb9, 0xb9, 0x8e, 0x22, 0x6d,
	0x39, 0x55, 0xc4, 0x9a, 0x94, 0x68, 0xe0, 0x00, 0xfa, 0x09, 0xcb, 0x24, 0xbe, 0x24, 0x8e, 0x75,
	0x28, 0x63, 0x52, 0xa2, 0x89, 0x4f, 0x01, 0x1d, 0x45, 0x36, 0x71, 0x61, 0xd1, 0xe5, 0x52, 0x26,
	0x51, 0xa1, 0xb7, 0xf0, 0x19, 0x0c, 0x66, 0x52, 0x1b, 0x52, 0xe9, 0xc2, 0x51, 0x68, 0x59, 0xe9,
	0x58, 0x5b, 0x16, 0x50, 0x4c, 0x2e, 0xcf, 0xad, 0x2b, 0xa8, 0x36, 0x0a, 0xe8, 0xd8, 0x24, 0x4e,
	0xed, 0x2c, 0x75, 0x92, 0x2f, 0x48, 0x74, 0xf0, 0x18, 0xba, 0x09, 0xeb, 0xf9, 0xc2, 0x50, 0xb1,
	0x06, 0x29, 0xd1, 0xc5, 0x0e, 0x34, 0x35, 0xc7, 0xe4, 0x58, 0x1a, 0xd1, 0xc3, 0x3e, 0xb4, 0x13,
	0x96, 0x4b, 0xa9, 0x8d, 0x3c, 0x37, 0x24, 0xfa, 0xc5, 0x42, 0x4a, 0xc6, 0x32, 0x35, 0x36, 0x8a,
	0x84, 0x98, 0xbe, 0x81, 0x60, 0xe9, 0x37, 0x9f, 0x3f, 0x7c, 0xc3, 0xd7, 0x70, 0x14, 0x7e, 0xf2,
	0x37, 0x5f, 0x10, 0xff, 0xb8, 0x46, 0xf9, 0x0f, 0xa3, 0xc1, 0x5f, 0x2e, 0xf4, 0x3e, 0x28, 0x3f,
	0xe8, 0xec, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xe7, 0x02, 0xa6, 0x62, 0x02, 0x00, 0x00,
}
