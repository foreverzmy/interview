// Code generated by protoc-gen-go. DO NOT EDIT.
// source: topic.proto

package topic

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	options "github.com/infobloxopen/protoc-gen-gorm/options"
	question "github.com/piex/interview/protorepo/question"
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

// Question from public import github.com/piex/interview/protorepo/question/question.proto
type Question = question.Question

// GetQuestionListRequest from public import github.com/piex/interview/protorepo/question/question.proto
type GetQuestionListRequest = question.GetQuestionListRequest

// QuestionList from public import github.com/piex/interview/protorepo/question/question.proto
type QuestionList = question.QuestionList

// CreateQuestionResponse from public import github.com/piex/interview/protorepo/question/question.proto
type CreateQuestionResponse = question.CreateQuestionResponse

// UpdateQuestionResponse from public import github.com/piex/interview/protorepo/question/question.proto
type UpdateQuestionResponse = question.UpdateQuestionResponse

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_7312ad0e4fa171e8, []int{0}
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

type Topic struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Slug                 string               `protobuf:"bytes,4,opt,name=slug,proto3" json:"slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Topic) Reset()         { *m = Topic{} }
func (m *Topic) String() string { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()    {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_7312ad0e4fa171e8, []int{1}
}

func (m *Topic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topic.Unmarshal(m, b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
}
func (m *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(m, src)
}
func (m *Topic) XXX_Size() int {
	return xxx_messageInfo_Topic.Size(m)
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func (m *Topic) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Topic) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Topic) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Topic) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type TopicList struct {
	Total                int64    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Topics               []*Topic `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicList) Reset()         { *m = TopicList{} }
func (m *TopicList) String() string { return proto.CompactTextString(m) }
func (*TopicList) ProtoMessage()    {}
func (*TopicList) Descriptor() ([]byte, []int) {
	return fileDescriptor_7312ad0e4fa171e8, []int{2}
}

func (m *TopicList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicList.Unmarshal(m, b)
}
func (m *TopicList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicList.Marshal(b, m, deterministic)
}
func (m *TopicList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicList.Merge(m, src)
}
func (m *TopicList) XXX_Size() int {
	return xxx_messageInfo_TopicList.Size(m)
}
func (m *TopicList) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicList.DiscardUnknown(m)
}

var xxx_messageInfo_TopicList proto.InternalMessageInfo

func (m *TopicList) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *TopicList) GetTopics() []*Topic {
	if m != nil {
		return m.Topics
	}
	return nil
}

type CreateTopicResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTopicResponse) Reset()         { *m = CreateTopicResponse{} }
func (m *CreateTopicResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTopicResponse) ProtoMessage()    {}
func (*CreateTopicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7312ad0e4fa171e8, []int{3}
}

func (m *CreateTopicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTopicResponse.Unmarshal(m, b)
}
func (m *CreateTopicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTopicResponse.Marshal(b, m, deterministic)
}
func (m *CreateTopicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTopicResponse.Merge(m, src)
}
func (m *CreateTopicResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTopicResponse.Size(m)
}
func (m *CreateTopicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTopicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTopicResponse proto.InternalMessageInfo

func (m *CreateTopicResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetTopicRequest struct {
	QuId                 int64    `protobuf:"varint,1,opt,name=qu_id,json=quId,proto3" json:"qu_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopicRequest) Reset()         { *m = GetTopicRequest{} }
func (m *GetTopicRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopicRequest) ProtoMessage()    {}
func (*GetTopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7312ad0e4fa171e8, []int{4}
}

func (m *GetTopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopicRequest.Unmarshal(m, b)
}
func (m *GetTopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopicRequest.Marshal(b, m, deterministic)
}
func (m *GetTopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopicRequest.Merge(m, src)
}
func (m *GetTopicRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopicRequest.Size(m)
}
func (m *GetTopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopicRequest proto.InternalMessageInfo

func (m *GetTopicRequest) GetQuId() int64 {
	if m != nil {
		return m.QuId
	}
	return 0
}

type GetQusByTopicRequest struct {
	TopicId              int64    `protobuf:"varint,1,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetQusByTopicRequest) Reset()         { *m = GetQusByTopicRequest{} }
func (m *GetQusByTopicRequest) String() string { return proto.CompactTextString(m) }
func (*GetQusByTopicRequest) ProtoMessage()    {}
func (*GetQusByTopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7312ad0e4fa171e8, []int{5}
}

func (m *GetQusByTopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQusByTopicRequest.Unmarshal(m, b)
}
func (m *GetQusByTopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQusByTopicRequest.Marshal(b, m, deterministic)
}
func (m *GetQusByTopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQusByTopicRequest.Merge(m, src)
}
func (m *GetQusByTopicRequest) XXX_Size() int {
	return xxx_messageInfo_GetQusByTopicRequest.Size(m)
}
func (m *GetQusByTopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQusByTopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetQusByTopicRequest proto.InternalMessageInfo

func (m *GetQusByTopicRequest) GetTopicId() int64 {
	if m != nil {
		return m.TopicId
	}
	return 0
}

type GetQusByTopicResponse struct {
	Qus                  []int64  `protobuf:"varint,1,rep,packed,name=qus,proto3" json:"qus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetQusByTopicResponse) Reset()         { *m = GetQusByTopicResponse{} }
func (m *GetQusByTopicResponse) String() string { return proto.CompactTextString(m) }
func (*GetQusByTopicResponse) ProtoMessage()    {}
func (*GetQusByTopicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7312ad0e4fa171e8, []int{6}
}

func (m *GetQusByTopicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQusByTopicResponse.Unmarshal(m, b)
}
func (m *GetQusByTopicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQusByTopicResponse.Marshal(b, m, deterministic)
}
func (m *GetQusByTopicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQusByTopicResponse.Merge(m, src)
}
func (m *GetQusByTopicResponse) XXX_Size() int {
	return xxx_messageInfo_GetQusByTopicResponse.Size(m)
}
func (m *GetQusByTopicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQusByTopicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetQusByTopicResponse proto.InternalMessageInfo

func (m *GetQusByTopicResponse) GetQus() []int64 {
	if m != nil {
		return m.Qus
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "interview.topic.Empty")
	proto.RegisterType((*Topic)(nil), "interview.topic.Topic")
	proto.RegisterType((*TopicList)(nil), "interview.topic.TopicList")
	proto.RegisterType((*CreateTopicResponse)(nil), "interview.topic.CreateTopicResponse")
	proto.RegisterType((*GetTopicRequest)(nil), "interview.topic.GetTopicRequest")
	proto.RegisterType((*GetQusByTopicRequest)(nil), "interview.topic.GetQusByTopicRequest")
	proto.RegisterType((*GetQusByTopicResponse)(nil), "interview.topic.GetQusByTopicResponse")
}

func init() { proto.RegisterFile("topic.proto", fileDescriptor_7312ad0e4fa171e8) }

var fileDescriptor_7312ad0e4fa171e8 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xcd, 0xda, 0xf9, 0x9c, 0xb4, 0x14, 0x6d, 0x0b, 0x4a, 0x2c, 0x24, 0x2c, 0x8b, 0x22, 0x73,
	0xa8, 0x2d, 0xc2, 0xad, 0x5c, 0x20, 0x80, 0x4a, 0x11, 0x48, 0x89, 0xe9, 0xa9, 0x17, 0x94, 0xd8,
	0x5b, 0xb3, 0x52, 0xec, 0xdd, 0x78, 0xd7, 0xd0, 0xfe, 0xb4, 0xe4, 0x2f, 0xf0, 0x07, 0xf8, 0x39,
	0xc8, 0xbb, 0x8e, 0x93, 0x26, 0x41, 0x88, 0xdb, 0xce, 0xe4, 0xbd, 0x37, 0xf3, 0x5e, 0xc6, 0xd0,
	0x95, 0x8c, 0xd3, 0xd0, 0xe3, 0x19, 0x93, 0x0c, 0x1f, 0xd1, 0x54, 0x92, 0xec, 0x07, 0x25, 0x3f,
	0x3d, 0xd5, 0xb6, 0x9e, 0xc6, 0x8c, 0xc5, 0x33, 0xe2, 0xab, 0x9f, 0xa7, 0xf9, 0x8d, 0x2f, 0x69,
	0x42, 0x84, 0x9c, 0x24, 0x5c, 0x33, 0xac, 0xf3, 0x98, 0xca, 0xef, 0xf9, 0xd4, 0x0b, 0x59, 0xe2,
	0xd3, 0xf4, 0x86, 0x4d, 0x67, 0xec, 0x96, 0x71, 0x92, 0x6a, 0x46, 0x78, 0x16, 0x93, 0xf4, 0x2c,
	0x66, 0x59, 0xe2, 0x33, 0x2e, 0x29, 0x4b, 0x85, 0x5f, 0x14, 0x25, 0xf7, 0xf5, 0x06, 0x97, 0x53,
	0x72, 0xeb, 0x57, 0xd3, 0x35, 0x3b, 0x23, 0x9c, 0xf9, 0xf3, 0x9c, 0x88, 0x82, 0x59, 0x3d, 0x34,
	0xd9, 0x69, 0x41, 0xe3, 0x43, 0xc2, 0xe5, 0x9d, 0xf3, 0x1b, 0x41, 0xe3, 0xaa, 0x58, 0x16, 0x5b,
	0x60, 0xd0, 0xa8, 0x87, 0x6c, 0xe4, 0x9a, 0x43, 0x58, 0x2e, 0xfa, 0x4d, 0xa8, 0xbb, 0xe8, 0x23,
	0x0a, 0x0c, 0x1a, 0xe1, 0x21, 0x74, 0xc2, 0x8c, 0x4c, 0x24, 0x89, 0xde, 0xca, 0x9e, 0x61, 0x23,
	0xb7, 0x3b, 0xb0, 0x3c, 0x6d, 0xce, 0x5b, 0x99, 0xf3, 0xae, 0x56, 0xe6, 0x86, 0xed, 0xe5, 0xa2,
	0x5f, 0x07, 0xe3, 0x0d, 0x0a, 0xd6, 0xb4, 0x42, 0x23, 0xe7, 0x51, 0xa9, 0x61, 0xfe, 0x8f, 0x46,
	0x45, 0xc3, 0x4f, 0xa0, 0x2e, 0x66, 0x79, 0xdc, 0xab, 0xdb, 0xc8, 0xed, 0x6c, 0x40, 0x54, 0xf7,
	0xfc, 0x70, 0xb9, 0xe8, 0x77, 0xda, 0xc8, 0x6a, 0xa8, 0xf4, 0x9d, 0x31, 0x74, 0x94, 0xb3, 0xcf,
	0x54, 0x48, 0x7c, 0x02, 0x0d, 0xc9, 0xe4, 0x64, 0xa6, 0x0d, 0x06, 0xba, 0xc0, 0x1e, 0x34, 0x15,
	0x56, 0xf4, 0x0c, 0xdb, 0x74, 0xbb, 0x83, 0xc7, 0xde, 0xd6, 0x5f, 0xe8, 0x29, 0x85, 0xa0, 0x44,
	0x39, 0xa7, 0x70, 0xfc, 0x4e, 0x19, 0xd2, 0x6d, 0x22, 0x38, 0x4b, 0x05, 0xc1, 0x0f, 0xd6, 0xd1,
	0x15, 0x71, 0x39, 0xcf, 0xe1, 0xe8, 0x82, 0xc8, 0x12, 0xa3, 0x92, 0xc7, 0xc7, 0xd0, 0x98, 0xe7,
	0xdf, 0x2a, 0x54, 0x7d, 0x9e, 0x5f, 0x46, 0xce, 0x4b, 0x38, 0xb9, 0x20, 0x72, 0x9c, 0x8b, 0xe1,
	0xdd, 0x3d, 0x70, 0x1f, 0xda, 0x6a, 0xe0, 0x1a, 0xdf, 0x52, 0xf5, 0x65, 0xe4, 0xbc, 0x80, 0x47,
	0x5b, 0x94, 0x72, 0x87, 0x87, 0x60, 0xce, 0x73, 0xd1, 0x43, 0xb6, 0xe9, 0x9a, 0x41, 0xf1, 0x1c,
	0xfc, 0x32, 0xe0, 0x40, 0x61, 0xbe, 0x16, 0x8e, 0x42, 0x82, 0xbf, 0x40, 0x77, 0x63, 0x7b, 0xfc,
	0x17, 0xb3, 0xd6, 0xb3, 0x9d, 0xfe, 0x1e, 0xcf, 0x4e, 0x0d, 0xbf, 0x87, 0x83, 0x95, 0x4b, 0x15,
	0xf1, 0xae, 0x9e, 0x3a, 0x31, 0xcb, 0xda, 0x3f, 0xa7, 0xe0, 0x38, 0x35, 0xfc, 0x09, 0xda, 0x2b,
	0x15, 0x6c, 0xef, 0x20, 0xb7, 0x62, 0xfc, 0x87, 0xd6, 0x35, 0x1c, 0xde, 0x0b, 0x07, 0x9f, 0xee,
	0x13, 0xdc, 0xc9, 0xdb, 0xda, 0x9c, 0x5b, 0x7d, 0x28, 0xe3, 0xf2, 0xa1, 0xb5, 0x87, 0xad, 0x6b,
	0x7d, 0x56, 0xa3, 0xda, 0x08, 0x8d, 0x8c, 0x69, 0x53, 0x9d, 0xec, 0xab, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x4c, 0x96, 0x5b, 0x2c, 0x01, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TopicServiceClient is the client API for TopicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopicServiceClient interface {
	// 创建
	CreateTopic(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*CreateTopicResponse, error)
	// 查找全部 topic
	GetTopicList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TopicList, error)
	// 查询 qu 的 topic
	GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*TopicList, error)
	// 查找有指定 topic 的 qus
	GetQusByTopic(ctx context.Context, in *GetQusByTopicRequest, opts ...grpc.CallOption) (*question.QuestionList, error)
}

type topicServiceClient struct {
	cc *grpc.ClientConn
}

func NewTopicServiceClient(cc *grpc.ClientConn) TopicServiceClient {
	return &topicServiceClient{cc}
}

func (c *topicServiceClient) CreateTopic(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*CreateTopicResponse, error) {
	out := new(CreateTopicResponse)
	err := c.cc.Invoke(ctx, "/interview.topic.TopicService/CreateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) GetTopicList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TopicList, error) {
	out := new(TopicList)
	err := c.cc.Invoke(ctx, "/interview.topic.TopicService/GetTopicList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*TopicList, error) {
	out := new(TopicList)
	err := c.cc.Invoke(ctx, "/interview.topic.TopicService/GetTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) GetQusByTopic(ctx context.Context, in *GetQusByTopicRequest, opts ...grpc.CallOption) (*question.QuestionList, error) {
	out := new(question.QuestionList)
	err := c.cc.Invoke(ctx, "/interview.topic.TopicService/GetQusByTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicServiceServer is the server API for TopicService service.
type TopicServiceServer interface {
	// 创建
	CreateTopic(context.Context, *Topic) (*CreateTopicResponse, error)
	// 查找全部 topic
	GetTopicList(context.Context, *Empty) (*TopicList, error)
	// 查询 qu 的 topic
	GetTopic(context.Context, *GetTopicRequest) (*TopicList, error)
	// 查找有指定 topic 的 qus
	GetQusByTopic(context.Context, *GetQusByTopicRequest) (*question.QuestionList, error)
}

func RegisterTopicServiceServer(s *grpc.Server, srv TopicServiceServer) {
	s.RegisterService(&_TopicService_serviceDesc, srv)
}

func _TopicService_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interview.topic.TopicService/CreateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).CreateTopic(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_GetTopicList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).GetTopicList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interview.topic.TopicService/GetTopicList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).GetTopicList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_GetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).GetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interview.topic.TopicService/GetTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).GetTopic(ctx, req.(*GetTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_GetQusByTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQusByTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).GetQusByTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interview.topic.TopicService/GetQusByTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).GetQusByTopic(ctx, req.(*GetQusByTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopicService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "interview.topic.TopicService",
	HandlerType: (*TopicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTopic",
			Handler:    _TopicService_CreateTopic_Handler,
		},
		{
			MethodName: "GetTopicList",
			Handler:    _TopicService_GetTopicList_Handler,
		},
		{
			MethodName: "GetTopic",
			Handler:    _TopicService_GetTopic_Handler,
		},
		{
			MethodName: "GetQusByTopic",
			Handler:    _TopicService_GetQusByTopic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "topic.proto",
}
