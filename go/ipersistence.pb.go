// Code generated by protoc-gen-go.
// source: ipersistence.proto
// DO NOT EDIT!

/*
Package ipersistence is a generated protocol buffer package.

It is generated from these files:
	ipersistence.proto

It has these top-level messages:
	FetchSaveRequest
	FetchSaveReply
*/
package ipersistence

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
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

// The request message containing the user's name.
type FetchSaveRequest struct {
	Url    string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	UserID string   `protobuf:"bytes,2,opt,name=userID" json:"userID,omitempty"`
	Tags   []string `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty"`
}

func (m *FetchSaveRequest) Reset()                    { *m = FetchSaveRequest{} }
func (m *FetchSaveRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchSaveRequest) ProtoMessage()               {}
func (*FetchSaveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FetchSaveRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *FetchSaveRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *FetchSaveRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// The response message containing the greetings
type FetchSaveReply struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	ImageID string `protobuf:"bytes,3,opt,name=imageID" json:"imageID,omitempty"`
	Path    string `protobuf:"bytes,4,opt,name=path" json:"path,omitempty"`
}

func (m *FetchSaveReply) Reset()                    { *m = FetchSaveReply{} }
func (m *FetchSaveReply) String() string            { return proto.CompactTextString(m) }
func (*FetchSaveReply) ProtoMessage()               {}
func (*FetchSaveReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FetchSaveReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *FetchSaveReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *FetchSaveReply) GetImageID() string {
	if m != nil {
		return m.ImageID
	}
	return ""
}

func (m *FetchSaveReply) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*FetchSaveRequest)(nil), "ipersistence.FetchSaveRequest")
	proto.RegisterType((*FetchSaveReply)(nil), "ipersistence.FetchSaveReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Saver service

type SaverClient interface {
	FetchSave(ctx context.Context, in *FetchSaveRequest, opts ...grpc.CallOption) (*FetchSaveReply, error)
}

type saverClient struct {
	cc *grpc.ClientConn
}

func NewSaverClient(cc *grpc.ClientConn) SaverClient {
	return &saverClient{cc}
}

func (c *saverClient) FetchSave(ctx context.Context, in *FetchSaveRequest, opts ...grpc.CallOption) (*FetchSaveReply, error) {
	out := new(FetchSaveReply)
	err := grpc.Invoke(ctx, "/ipersistence.Saver/FetchSave", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Saver service

type SaverServer interface {
	FetchSave(context.Context, *FetchSaveRequest) (*FetchSaveReply, error)
}

func RegisterSaverServer(s *grpc.Server, srv SaverServer) {
	s.RegisterService(&_Saver_serviceDesc, srv)
}

func _Saver_FetchSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaverServer).FetchSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipersistence.Saver/FetchSave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaverServer).FetchSave(ctx, req.(*FetchSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Saver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ipersistence.Saver",
	HandlerType: (*SaverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchSave",
			Handler:    _Saver_FetchSave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ipersistence.proto",
}

func init() { proto.RegisterFile("ipersistence.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0xad, 0x9d, 0xd3, 0x5e, 0x44, 0xe6, 0x7d, 0x90, 0x20, 0x2a, 0xa3, 0x4f, 0x7b, 0xea,
	0x83, 0xfa, 0x0b, 0xc6, 0x10, 0x8a, 0x2f, 0xa5, 0xfa, 0x07, 0x62, 0xb8, 0x74, 0x81, 0xd4, 0xc6,
	0xdc, 0x44, 0xdc, 0xbf, 0x97, 0xc4, 0x4e, 0x37, 0x61, 0x6f, 0xe7, 0x3b, 0x87, 0x9c, 0x43, 0x2e,
	0xa0, 0xb6, 0xe4, 0x58, 0xb3, 0xa7, 0x77, 0x45, 0x95, 0x75, 0x83, 0x1f, 0xf0, 0x7c, 0xd7, 0x2b,
	0x1b, 0x98, 0x3d, 0x91, 0x57, 0xeb, 0x17, 0xf9, 0x49, 0x2d, 0x7d, 0x04, 0x62, 0x8f, 0x33, 0xc8,
	0x83, 0x33, 0x22, 0x9b, 0x67, 0x8b, 0xa2, 0x8d, 0x12, 0xaf, 0x60, 0x1a, 0x98, 0x5c, 0xbd, 0x12,
	0xc7, 0xc9, 0x1c, 0x09, 0x11, 0x26, 0x5e, 0x76, 0x2c, 0xf2, 0x79, 0xbe, 0x28, 0xda, 0xa4, 0x4b,
	0x07, 0x17, 0x3b, 0x8d, 0xd6, 0x6c, 0x50, 0xc0, 0x29, 0x07, 0xa5, 0x88, 0x39, 0x75, 0x9e, 0xb5,
	0x5b, 0x8c, 0x49, 0x4f, 0xcc, 0xb2, 0xa3, 0xb1, 0x78, 0x8b, 0x31, 0xd1, 0xbd, 0xec, 0xa8, 0x5e,
	0x89, 0xfc, 0x27, 0x19, 0x31, 0x6e, 0x5a, 0xe9, 0xd7, 0x62, 0x92, 0xec, 0xa4, 0xef, 0x5f, 0xe1,
	0x24, 0xce, 0x39, 0x7c, 0x86, 0xe2, 0x77, 0x1c, 0xef, 0xaa, 0xbd, 0xef, 0xff, 0xff, 0xe7, 0xf5,
	0xcd, 0xc1, 0xdc, 0x9a, 0x4d, 0x79, 0xb4, 0x7c, 0x84, 0x5b, 0x3d, 0x54, 0x9d, 0xb3, 0xaa, 0xa2,
	0x2f, 0xd9, 0x5b, 0x43, 0xbc, 0xf7, 0x62, 0x79, 0x59, 0x37, 0x7f, 0xd4, 0xc4, 0xeb, 0x36, 0xd9,
	0xdb, 0x34, 0x9d, 0xf9, 0xe1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x50, 0xcf, 0xfc, 0x5c, 0x7c, 0x01,
	0x00, 0x00,
}