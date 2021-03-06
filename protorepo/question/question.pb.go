// Code generated by protoc-gen-go. DO NOT EDIT.
// source: question.proto

package question

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
	return fileDescriptor_64d310e1b23c85d7, []int{0}
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

type Question struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Title                string               `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Summary              string               `protobuf:"bytes,5,opt,name=summary,proto3" json:"summary,omitempty"`
	Content              string               `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Difficulty           int32                `protobuf:"varint,7,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Question) Reset()         { *m = Question{} }
func (m *Question) String() string { return proto.CompactTextString(m) }
func (*Question) ProtoMessage()    {}
func (*Question) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d310e1b23c85d7, []int{1}
}

func (m *Question) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Question.Unmarshal(m, b)
}
func (m *Question) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Question.Marshal(b, m, deterministic)
}
func (m *Question) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question.Merge(m, src)
}
func (m *Question) XXX_Size() int {
	return xxx_messageInfo_Question.Size(m)
}
func (m *Question) XXX_DiscardUnknown() {
	xxx_messageInfo_Question.DiscardUnknown(m)
}

var xxx_messageInfo_Question proto.InternalMessageInfo

func (m *Question) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Question) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Question) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Question) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Question) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *Question) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Question) GetDifficulty() int32 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

type GetQuestionListRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Keyword              string   `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetQuestionListRequest) Reset()         { *m = GetQuestionListRequest{} }
func (m *GetQuestionListRequest) String() string { return proto.CompactTextString(m) }
func (*GetQuestionListRequest) ProtoMessage()    {}
func (*GetQuestionListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d310e1b23c85d7, []int{2}
}

func (m *GetQuestionListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQuestionListRequest.Unmarshal(m, b)
}
func (m *GetQuestionListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQuestionListRequest.Marshal(b, m, deterministic)
}
func (m *GetQuestionListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQuestionListRequest.Merge(m, src)
}
func (m *GetQuestionListRequest) XXX_Size() int {
	return xxx_messageInfo_GetQuestionListRequest.Size(m)
}
func (m *GetQuestionListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQuestionListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetQuestionListRequest proto.InternalMessageInfo

func (m *GetQuestionListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetQuestionListRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GetQuestionListRequest) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

type QuestionList struct {
	Total                int64       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Questions            []*Question `protobuf:"bytes,2,rep,name=questions,proto3" json:"questions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *QuestionList) Reset()         { *m = QuestionList{} }
func (m *QuestionList) String() string { return proto.CompactTextString(m) }
func (*QuestionList) ProtoMessage()    {}
func (*QuestionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d310e1b23c85d7, []int{3}
}

func (m *QuestionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuestionList.Unmarshal(m, b)
}
func (m *QuestionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuestionList.Marshal(b, m, deterministic)
}
func (m *QuestionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestionList.Merge(m, src)
}
func (m *QuestionList) XXX_Size() int {
	return xxx_messageInfo_QuestionList.Size(m)
}
func (m *QuestionList) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestionList.DiscardUnknown(m)
}

var xxx_messageInfo_QuestionList proto.InternalMessageInfo

func (m *QuestionList) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *QuestionList) GetQuestions() []*Question {
	if m != nil {
		return m.Questions
	}
	return nil
}

type CreateQuestionResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateQuestionResponse) Reset()         { *m = CreateQuestionResponse{} }
func (m *CreateQuestionResponse) String() string { return proto.CompactTextString(m) }
func (*CreateQuestionResponse) ProtoMessage()    {}
func (*CreateQuestionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d310e1b23c85d7, []int{4}
}

func (m *CreateQuestionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateQuestionResponse.Unmarshal(m, b)
}
func (m *CreateQuestionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateQuestionResponse.Marshal(b, m, deterministic)
}
func (m *CreateQuestionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateQuestionResponse.Merge(m, src)
}
func (m *CreateQuestionResponse) XXX_Size() int {
	return xxx_messageInfo_CreateQuestionResponse.Size(m)
}
func (m *CreateQuestionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateQuestionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateQuestionResponse proto.InternalMessageInfo

func (m *CreateQuestionResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "interview.question.Empty")
	proto.RegisterType((*Question)(nil), "interview.question.Question")
	proto.RegisterType((*GetQuestionListRequest)(nil), "interview.question.GetQuestionListRequest")
	proto.RegisterType((*QuestionList)(nil), "interview.question.QuestionList")
	proto.RegisterType((*CreateQuestionResponse)(nil), "interview.question.CreateQuestionResponse")
}

func init() { proto.RegisterFile("question.proto", fileDescriptor_64d310e1b23c85d7) }

var fileDescriptor_64d310e1b23c85d7 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x8f, 0x93, 0x40,
	0x14, 0xef, 0xd0, 0x52, 0xda, 0xd7, 0xa6, 0xab, 0x13, 0xb3, 0x01, 0x62, 0x94, 0x70, 0x70, 0x89,
	0x66, 0xa9, 0xa9, 0xb7, 0x9e, 0xd6, 0xaa, 0xd1, 0x44, 0x0f, 0xbb, 0xa8, 0x97, 0x8d, 0x07, 0x69,
	0x79, 0xc5, 0x89, 0xc0, 0x20, 0x0c, 0xbb, 0xd6, 0x6f, 0xd6, 0x7e, 0x13, 0xbf, 0x86, 0x9f, 0xc0,
	0x30, 0x94, 0xb6, 0xae, 0x4d, 0x75, 0x6f, 0xcc, 0xfc, 0xfe, 0xbc, 0xc7, 0x7b, 0xbf, 0x81, 0xc1,
	0xb7, 0x02, 0x73, 0xc1, 0x78, 0xe2, 0xa6, 0x19, 0x17, 0x9c, 0x52, 0x96, 0x08, 0xcc, 0xae, 0x18,
	0x5e, 0xbb, 0x35, 0x62, 0x3e, 0x0c, 0x39, 0x0f, 0x23, 0x1c, 0x4a, 0xc6, 0xb4, 0x98, 0x0f, 0x05,
	0x8b, 0x31, 0x17, 0x7e, 0x9c, 0x56, 0x22, 0x73, 0x1c, 0x32, 0xf1, 0xa5, 0x98, 0xba, 0x33, 0x1e,
	0x0f, 0x59, 0x32, 0xe7, 0xd3, 0x88, 0x7f, 0xe7, 0x29, 0x26, 0x95, 0x62, 0x76, 0x1a, 0x62, 0x72,
	0x1a, 0xf2, 0x2c, 0x1e, 0xf2, 0xb4, 0x34, 0xcc, 0x87, 0xe5, 0xa1, 0xd2, 0xda, 0x1a, 0xa8, 0xaf,
	0xe2, 0x54, 0x2c, 0xec, 0x9f, 0x0a, 0x74, 0x2e, 0xd6, 0x25, 0xa9, 0x09, 0x0a, 0x0b, 0x74, 0x62,
	0x11, 0xa7, 0x39, 0x81, 0xd5, 0xd2, 0x68, 0x43, 0xcb, 0x21, 0x6f, 0x88, 0xa7, 0xb0, 0x80, 0x4e,
	0xa0, 0x3b, 0xcb, 0xd0, 0x17, 0x18, 0x3c, 0x17, 0xba, 0x62, 0x11, 0xa7, 0x37, 0x32, 0xdd, 0xaa,
	0x45, 0xb7, 0x6e, 0xd1, 0xfd, 0x50, 0xb7, 0x38, 0xe9, 0xac, 0x96, 0x46, 0x0b, 0x94, 0x33, 0xe2,
	0x6d, 0x65, 0xa5, 0x47, 0x91, 0x06, 0x6b, 0x8f, 0xe6, 0x6d, 0x3c, 0x36, 0x32, 0xfa, 0x00, 0x54,
	0xc1, 0x44, 0x84, 0x7a, 0xcb, 0x22, 0x4e, 0x77, 0x87, 0x53, 0x5d, 0x53, 0x1b, 0xb4, 0xbc, 0x88,
	0x63, 0x3f, 0x5b, 0xe8, 0xea, 0x0d, 0x46, 0x0d, 0xd0, 0x47, 0xa0, 0xcd, 0x78, 0x22, 0x30, 0x11,
	0x7a, 0x5b, 0x72, 0xfa, 0xab, 0xa5, 0xd1, 0x81, 0xf6, 0x58, 0x39, 0x39, 0x29, 0x79, 0x6b, 0x90,
	0x3e, 0x01, 0x08, 0xd8, 0x7c, 0xce, 0x66, 0x45, 0x24, 0x16, 0xba, 0x66, 0x11, 0x47, 0x9d, 0xf4,
	0x56, 0x4b, 0x43, 0x03, 0x75, 0x4c, 0x9e, 0x9e, 0x11, 0x6f, 0x07, 0x1e, 0xdf, 0x59, 0x2d, 0x8d,
	0x7e, 0x87, 0x98, 0x9d, 0x7a, 0x83, 0xf6, 0x25, 0x1c, 0xbf, 0x46, 0x51, 0x4f, 0xf7, 0x1d, 0xcb,
	0x85, 0x87, 0x12, 0xa4, 0x14, 0x5a, 0xa9, 0x1f, 0xa2, 0x1c, 0xb5, 0xea, 0xc9, 0xef, 0xf2, 0x2e,
	0x67, 0x3f, 0x50, 0xce, 0x56, 0xf5, 0xe4, 0x37, 0xd5, 0x41, 0xfb, 0x8a, 0x8b, 0x6b, 0x9e, 0x05,
	0x72, 0x5c, 0x5d, 0xaf, 0x3e, 0xda, 0x9f, 0xa1, 0xbf, 0x6b, 0x4c, 0xef, 0x81, 0x2a, 0xb8, 0xf0,
	0xa3, 0x6a, 0x7b, 0x5e, 0x75, 0xa0, 0x63, 0xe8, 0xd6, 0xdd, 0xe4, 0xba, 0x62, 0x35, 0x9d, 0xde,
	0xe8, 0xbe, 0xfb, 0x77, 0xd6, 0xdc, 0xda, 0xca, 0xdb, 0xd2, 0x6d, 0x07, 0x8e, 0x5f, 0xc8, 0xcd,
	0x6d, 0x40, 0xcc, 0x53, 0x9e, 0xe4, 0x48, 0x07, 0xdb, 0x98, 0x94, 0xd1, 0x18, 0xfd, 0x52, 0xe0,
	0xa8, 0x26, 0xbd, 0x2f, 0xad, 0xaf, 0x90, 0x7e, 0x82, 0xc1, 0x9f, 0x6a, 0x7a, 0xb0, 0xb0, 0xf9,
	0x78, 0x1f, 0xba, 0xbf, 0xbe, 0xdd, 0xa0, 0x6f, 0x61, 0xf0, 0x51, 0x26, 0xe2, 0x3f, 0xdd, 0x8d,
	0x7d, 0x68, 0xf5, 0x00, 0x1a, 0xf4, 0x02, 0xee, 0xee, 0xac, 0xe9, 0x25, 0x0a, 0x9f, 0x45, 0xff,
	0xf0, 0x3b, 0x88, 0xda, 0x0d, 0xea, 0xc3, 0xd1, 0x8d, 0xcd, 0xd3, 0xbd, 0x3f, 0xb8, 0x3f, 0x1e,
	0xa6, 0x75, 0xc8, 0xbe, 0x24, 0xda, 0x8d, 0x09, 0x5c, 0x6e, 0x82, 0x76, 0xde, 0x38, 0x27, 0xd3,
	0xb6, 0x7c, 0x40, 0xcf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x12, 0x2b, 0xdd, 0xec, 0x5b, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QuestionServiveClient is the client API for QuestionServive service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuestionServiveClient interface {
	// 创建
	CreateQuestion(ctx context.Context, in *Question, opts ...grpc.CallOption) (*CreateQuestionResponse, error)
	// 更新
	UpdateQuestion(ctx context.Context, in *Question, opts ...grpc.CallOption) (*Empty, error)
	// 查看详情
	GetQuestionDetail(ctx context.Context, in *Question, opts ...grpc.CallOption) (*Question, error)
	// 查看列表
	GetQuestionList(ctx context.Context, in *GetQuestionListRequest, opts ...grpc.CallOption) (*QuestionList, error)
}

type questionServiveClient struct {
	cc *grpc.ClientConn
}

func NewQuestionServiveClient(cc *grpc.ClientConn) QuestionServiveClient {
	return &questionServiveClient{cc}
}

func (c *questionServiveClient) CreateQuestion(ctx context.Context, in *Question, opts ...grpc.CallOption) (*CreateQuestionResponse, error) {
	out := new(CreateQuestionResponse)
	err := c.cc.Invoke(ctx, "/interview.question.QuestionServive/CreateQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiveClient) UpdateQuestion(ctx context.Context, in *Question, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/interview.question.QuestionServive/UpdateQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiveClient) GetQuestionDetail(ctx context.Context, in *Question, opts ...grpc.CallOption) (*Question, error) {
	out := new(Question)
	err := c.cc.Invoke(ctx, "/interview.question.QuestionServive/GetQuestionDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiveClient) GetQuestionList(ctx context.Context, in *GetQuestionListRequest, opts ...grpc.CallOption) (*QuestionList, error) {
	out := new(QuestionList)
	err := c.cc.Invoke(ctx, "/interview.question.QuestionServive/GetQuestionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionServiveServer is the server API for QuestionServive service.
type QuestionServiveServer interface {
	// 创建
	CreateQuestion(context.Context, *Question) (*CreateQuestionResponse, error)
	// 更新
	UpdateQuestion(context.Context, *Question) (*Empty, error)
	// 查看详情
	GetQuestionDetail(context.Context, *Question) (*Question, error)
	// 查看列表
	GetQuestionList(context.Context, *GetQuestionListRequest) (*QuestionList, error)
}

func RegisterQuestionServiveServer(s *grpc.Server, srv QuestionServiveServer) {
	s.RegisterService(&_QuestionServive_serviceDesc, srv)
}

func _QuestionServive_CreateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Question)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiveServer).CreateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interview.question.QuestionServive/CreateQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiveServer).CreateQuestion(ctx, req.(*Question))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionServive_UpdateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Question)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiveServer).UpdateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interview.question.QuestionServive/UpdateQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiveServer).UpdateQuestion(ctx, req.(*Question))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionServive_GetQuestionDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Question)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiveServer).GetQuestionDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interview.question.QuestionServive/GetQuestionDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiveServer).GetQuestionDetail(ctx, req.(*Question))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionServive_GetQuestionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiveServer).GetQuestionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interview.question.QuestionServive/GetQuestionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiveServer).GetQuestionList(ctx, req.(*GetQuestionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QuestionServive_serviceDesc = grpc.ServiceDesc{
	ServiceName: "interview.question.QuestionServive",
	HandlerType: (*QuestionServiveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQuestion",
			Handler:    _QuestionServive_CreateQuestion_Handler,
		},
		{
			MethodName: "UpdateQuestion",
			Handler:    _QuestionServive_UpdateQuestion_Handler,
		},
		{
			MethodName: "GetQuestionDetail",
			Handler:    _QuestionServive_GetQuestionDetail_Handler,
		},
		{
			MethodName: "GetQuestionList",
			Handler:    _QuestionServive_GetQuestionList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "question.proto",
}
