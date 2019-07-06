// Code generated by protoc-gen-go. DO NOT EDIT.
// source: answer.proto

package answer

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	options "github.com/infobloxopen/protoc-gen-gorm/options"
	grpc "google.golang.org/grpc"
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

// Timestamp from public import google/protobuf/timestamp.proto
type Timestamp = timestamp.Timestamp

// GormFileOptions from public import github.com/infobloxopen/protoc-gen-gorm/options/gorm.proto
type GormFileOptions = options.GormFileOptions

// GormMessageOptions from public import github.com/infobloxopen/protoc-gen-gorm/options/gorm.proto
type GormMessageOptions = options.GormMessageOptions

// ExtraField from public import github.com/infobloxopen/protoc-gen-gorm/options/gorm.proto
type ExtraField = options.ExtraField

// GormFieldOptions from public import github.com/infobloxopen/protoc-gen-gorm/options/gorm.proto
type GormFieldOptions = options.GormFieldOptions
type GormFieldOptions_HasOne = options.GormFieldOptions_HasOne
type GormFieldOptions_BelongsTo = options.GormFieldOptions_BelongsTo
type GormFieldOptions_HasMany = options.GormFieldOptions_HasMany
type GormFieldOptions_ManyToMany = options.GormFieldOptions_ManyToMany

// GormTag from public import github.com/infobloxopen/protoc-gen-gorm/options/gorm.proto
type GormTag = options.GormTag

// HasOneOptions from public import github.com/infobloxopen/protoc-gen-gorm/options/gorm.proto
type HasOneOptions = options.HasOneOptions

// BelongsToOptions from public import github.com/infobloxopen/protoc-gen-gorm/options/gorm.proto
type BelongsToOptions = options.BelongsToOptions

// HasManyOptions from public import github.com/infobloxopen/protoc-gen-gorm/options/gorm.proto
type HasManyOptions = options.HasManyOptions

// ManyToManyOptions from public import github.com/infobloxopen/protoc-gen-gorm/options/gorm.proto
type ManyToManyOptions = options.ManyToManyOptions

// AutoServerOptions from public import github.com/infobloxopen/protoc-gen-gorm/options/gorm.proto
type AutoServerOptions = options.AutoServerOptions

// MethodOptions from public import github.com/infobloxopen/protoc-gen-gorm/options/gorm.proto
type MethodOptions = options.MethodOptions

var E_FileOpts = options.E_FileOpts

var E_Opts = options.E_Opts

var E_Field = options.E_Field

var E_Server = options.E_Server

var E_Method = options.E_Method

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f97dbee99679ee, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Answer struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	QuId                 int64                `protobuf:"varint,4,opt,name=quId,proto3" json:"quId,omitempty"`
	Content              string               `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Answer) Reset()         { *m = Answer{} }
func (m *Answer) String() string { return proto.CompactTextString(m) }
func (*Answer) ProtoMessage()    {}
func (*Answer) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f97dbee99679ee, []int{1}
}

func (m *Answer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Answer.Unmarshal(m, b)
}
func (m *Answer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Answer.Marshal(b, m, deterministic)
}
func (m *Answer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Answer.Merge(m, src)
}
func (m *Answer) XXX_Size() int {
	return xxx_messageInfo_Answer.Size(m)
}
func (m *Answer) XXX_DiscardUnknown() {
	xxx_messageInfo_Answer.DiscardUnknown(m)
}

var xxx_messageInfo_Answer proto.InternalMessageInfo

func (m *Answer) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Answer) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Answer) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Answer) GetQuId() int64 {
	if m != nil {
		return m.QuId
	}
	return 0
}

func (m *Answer) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type AnswerList struct {
	Total                int64     `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Answers              []*Answer `protobuf:"bytes,2,rep,name=answers,proto3" json:"answers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AnswerList) Reset()         { *m = AnswerList{} }
func (m *AnswerList) String() string { return proto.CompactTextString(m) }
func (*AnswerList) ProtoMessage()    {}
func (*AnswerList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f97dbee99679ee, []int{2}
}

func (m *AnswerList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnswerList.Unmarshal(m, b)
}
func (m *AnswerList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnswerList.Marshal(b, m, deterministic)
}
func (m *AnswerList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnswerList.Merge(m, src)
}
func (m *AnswerList) XXX_Size() int {
	return xxx_messageInfo_AnswerList.Size(m)
}
func (m *AnswerList) XXX_DiscardUnknown() {
	xxx_messageInfo_AnswerList.DiscardUnknown(m)
}

var xxx_messageInfo_AnswerList proto.InternalMessageInfo

func (m *AnswerList) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *AnswerList) GetAnswers() []*Answer {
	if m != nil {
		return m.Answers
	}
	return nil
}

type CreateAnswerResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAnswerResponse) Reset()         { *m = CreateAnswerResponse{} }
func (m *CreateAnswerResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAnswerResponse) ProtoMessage()    {}
func (*CreateAnswerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f97dbee99679ee, []int{3}
}

func (m *CreateAnswerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAnswerResponse.Unmarshal(m, b)
}
func (m *CreateAnswerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAnswerResponse.Marshal(b, m, deterministic)
}
func (m *CreateAnswerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAnswerResponse.Merge(m, src)
}
func (m *CreateAnswerResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAnswerResponse.Size(m)
}
func (m *CreateAnswerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAnswerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAnswerResponse proto.InternalMessageInfo

func (m *CreateAnswerResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetAnswerListRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	QuId                 int64    `protobuf:"varint,3,opt,name=quId,proto3" json:"quId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAnswerListRequest) Reset()         { *m = GetAnswerListRequest{} }
func (m *GetAnswerListRequest) String() string { return proto.CompactTextString(m) }
func (*GetAnswerListRequest) ProtoMessage()    {}
func (*GetAnswerListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f97dbee99679ee, []int{4}
}

func (m *GetAnswerListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAnswerListRequest.Unmarshal(m, b)
}
func (m *GetAnswerListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAnswerListRequest.Marshal(b, m, deterministic)
}
func (m *GetAnswerListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAnswerListRequest.Merge(m, src)
}
func (m *GetAnswerListRequest) XXX_Size() int {
	return xxx_messageInfo_GetAnswerListRequest.Size(m)
}
func (m *GetAnswerListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAnswerListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAnswerListRequest proto.InternalMessageInfo

func (m *GetAnswerListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetAnswerListRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GetAnswerListRequest) GetQuId() int64 {
	if m != nil {
		return m.QuId
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "interview.answer.Empty")
	proto.RegisterType((*Answer)(nil), "interview.answer.Answer")
	proto.RegisterType((*AnswerList)(nil), "interview.answer.AnswerList")
	proto.RegisterType((*CreateAnswerResponse)(nil), "interview.answer.CreateAnswerResponse")
	proto.RegisterType((*GetAnswerListRequest)(nil), "interview.answer.GetAnswerListRequest")
}

func init() { proto.RegisterFile("answer.proto", fileDescriptor_c6f97dbee99679ee) }

var fileDescriptor_c6f97dbee99679ee = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xcd, 0x3a, 0xb6, 0x93, 0x4e, 0xd3, 0xea, 0xa7, 0x55, 0xa4, 0x9f, 0x63, 0x55, 0x22, 0xf2,
	0x21, 0xf5, 0xa5, 0xb6, 0x14, 0x6e, 0x39, 0xb5, 0xae, 0x10, 0x20, 0x71, 0xa8, 0x96, 0x7f, 0x12,
	0x37, 0xc7, 0x9e, 0x9a, 0x95, 0x62, 0xaf, 0xeb, 0x5d, 0x53, 0xe0, 0x06, 0xdf, 0x2a, 0xf9, 0x74,
	0xc8, 0xbb, 0x09, 0x01, 0x9a, 0x0a, 0x71, 0x9b, 0x9d, 0x99, 0xf7, 0xe6, 0xcd, 0x9b, 0x85, 0x51,
	0x5a, 0xc9, 0x7b, 0x6c, 0xa2, 0xba, 0x11, 0x4a, 0xd0, 0xff, 0x78, 0xa5, 0xb0, 0xf9, 0xc4, 0xf1,
	0x3e, 0x32, 0x79, 0xff, 0x49, 0x21, 0x44, 0xb1, 0xc2, 0x58, 0xd7, 0x97, 0xed, 0x6d, 0xac, 0x78,
	0x89, 0x52, 0xa5, 0x65, 0x6d, 0x20, 0xfe, 0xa2, 0xe0, 0xea, 0x63, 0xbb, 0x8c, 0x32, 0x51, 0xc6,
	0xbc, 0xba, 0x15, 0xcb, 0x95, 0xf8, 0x2c, 0x6a, 0xac, 0x0c, 0x22, 0xbb, 0x28, 0xb0, 0xba, 0x28,
	0x44, 0x53, 0xc6, 0xa2, 0x56, 0x5c, 0x54, 0x32, 0xee, 0x1e, 0x06, 0x1b, 0x0c, 0xc0, 0x79, 0x56,
	0xd6, 0xea, 0x4b, 0xf0, 0xdd, 0x02, 0xf7, 0x4a, 0x0f, 0xa4, 0x3e, 0x58, 0x3c, 0xf7, 0xc8, 0x94,
	0x84, 0xfd, 0x04, 0x36, 0xeb, 0x89, 0x0b, 0x76, 0x48, 0x5e, 0x10, 0x66, 0xf1, 0x9c, 0x26, 0x70,
	0x94, 0x35, 0x98, 0x2a, 0xcc, 0xaf, 0x94, 0x67, 0x4d, 0x49, 0x78, 0x3c, 0xf7, 0x23, 0x23, 0x30,
	0xda, 0x09, 0x8c, 0xde, 0xec, 0x04, 0x26, 0xc3, 0xcd, 0x7a, 0x62, 0x83, 0x75, 0x49, 0xd8, 0x1e,
	0xd6, 0x71, 0xb4, 0x75, 0xbe, 0xe5, 0xe8, 0xff, 0x0b, 0xc7, 0x4f, 0x18, 0x3d, 0x03, 0xfb, 0xae,
	0x7d, 0x99, 0x7b, 0xb6, 0x56, 0xb9, 0x6f, 0xd1, 0x59, 0x3a, 0x83, 0x41, 0x26, 0x2a, 0x85, 0x95,
	0xf2, 0x9c, 0x29, 0x09, 0x8f, 0x92, 0xd1, 0x66, 0x3d, 0x19, 0x82, 0xbb, 0xb0, 0xce, 0xcf, 0x2f,
	0x09, 0xdb, 0x15, 0x17, 0xa7, 0x9b, 0xf5, 0x04, 0x86, 0xc4, 0x77, 0x8d, 0xd5, 0xc1, 0x3b, 0x00,
	0xe3, 0xc1, 0x2b, 0x2e, 0x15, 0x1d, 0x83, 0xa3, 0x84, 0x4a, 0x57, 0xc6, 0x0a, 0x66, 0x1e, 0x74,
	0x0e, 0x03, 0xd3, 0x2d, 0x3d, 0x6b, 0xda, 0x0f, 0x8f, 0xe7, 0x5e, 0xf4, 0xe7, 0xc9, 0x22, 0x43,
	0xc2, 0x76, 0x8d, 0xc1, 0x0c, 0xc6, 0xd7, 0x7a, 0xfd, 0x6d, 0x01, 0x65, 0x2d, 0x2a, 0x89, 0xf4,
	0x74, 0xef, 0x74, 0xe7, 0x6e, 0xc0, 0x60, 0xfc, 0x1c, 0xd5, 0x5e, 0x02, 0xc3, 0xbb, 0x16, 0xa5,
	0xa2, 0x14, 0xec, 0x3a, 0x2d, 0x50, 0x77, 0x3a, 0x4c, 0xc7, 0x5d, 0x4e, 0xf2, 0xaf, 0xa8, 0x8f,
	0xe0, 0x30, 0x1d, 0x77, 0x39, 0xed, 0x4a, 0x5f, 0x33, 0xea, 0x78, 0xfe, 0xcd, 0x82, 0x13, 0xc3,
	0xf8, 0xba, 0x13, 0x99, 0x21, 0x65, 0x30, 0xfa, 0x55, 0x0d, 0x7d, 0x74, 0x01, 0x7f, 0xf6, 0xb0,
	0x72, 0x68, 0x8f, 0xa0, 0x47, 0xaf, 0x61, 0xf4, 0x56, 0x1f, 0xe7, 0xaf, 0x9c, 0xff, 0x3f, 0xac,
	0x98, 0x1f, 0xd8, 0xa3, 0xef, 0xe1, 0xe4, 0xb7, 0xf5, 0xe9, 0x81, 0xf9, 0x87, 0xfc, 0xf1, 0xcf,
	0x1e, 0x9b, 0xd6, 0x35, 0x05, 0xbd, 0x64, 0xf8, 0x61, 0x7b, 0xe1, 0x9b, 0xde, 0x0d, 0x59, 0xba,
	0xfa, 0x93, 0x3d, 0xfd, 0x11, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xdd, 0x68, 0x0c, 0x79, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AnswerServiceClient is the client API for AnswerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnswerServiceClient interface {
	// 创建
	CreateAnswer(ctx context.Context, in *Answer, opts ...grpc.CallOption) (*CreateAnswerResponse, error)
	// 编辑
	UpdateAnswer(ctx context.Context, in *Answer, opts ...grpc.CallOption) (*Empty, error)
	// 获取列表
	GetAnswerList(ctx context.Context, in *GetAnswerListRequest, opts ...grpc.CallOption) (*AnswerList, error)
}

type answerServiceClient struct {
	cc *grpc.ClientConn
}

func NewAnswerServiceClient(cc *grpc.ClientConn) AnswerServiceClient {
	return &answerServiceClient{cc}
}

func (c *answerServiceClient) CreateAnswer(ctx context.Context, in *Answer, opts ...grpc.CallOption) (*CreateAnswerResponse, error) {
	out := new(CreateAnswerResponse)
	err := c.cc.Invoke(ctx, "/interview.answer.AnswerService/CreateAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerServiceClient) UpdateAnswer(ctx context.Context, in *Answer, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/interview.answer.AnswerService/UpdateAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerServiceClient) GetAnswerList(ctx context.Context, in *GetAnswerListRequest, opts ...grpc.CallOption) (*AnswerList, error) {
	out := new(AnswerList)
	err := c.cc.Invoke(ctx, "/interview.answer.AnswerService/GetAnswerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnswerServiceServer is the server API for AnswerService service.
type AnswerServiceServer interface {
	// 创建
	CreateAnswer(context.Context, *Answer) (*CreateAnswerResponse, error)
	// 编辑
	UpdateAnswer(context.Context, *Answer) (*Empty, error)
	// 获取列表
	GetAnswerList(context.Context, *GetAnswerListRequest) (*AnswerList, error)
}

func RegisterAnswerServiceServer(s *grpc.Server, srv AnswerServiceServer) {
	s.RegisterService(&_AnswerService_serviceDesc, srv)
}

func _AnswerService_CreateAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Answer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerServiceServer).CreateAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interview.answer.AnswerService/CreateAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerServiceServer).CreateAnswer(ctx, req.(*Answer))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswerService_UpdateAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Answer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerServiceServer).UpdateAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interview.answer.AnswerService/UpdateAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerServiceServer).UpdateAnswer(ctx, req.(*Answer))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswerService_GetAnswerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnswerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerServiceServer).GetAnswerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interview.answer.AnswerService/GetAnswerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerServiceServer).GetAnswerList(ctx, req.(*GetAnswerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AnswerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "interview.answer.AnswerService",
	HandlerType: (*AnswerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAnswer",
			Handler:    _AnswerService_CreateAnswer_Handler,
		},
		{
			MethodName: "UpdateAnswer",
			Handler:    _AnswerService_UpdateAnswer_Handler,
		},
		{
			MethodName: "GetAnswerList",
			Handler:    _AnswerService_GetAnswerList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "answer.proto",
}
