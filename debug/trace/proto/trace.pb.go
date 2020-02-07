// Code generated by protoc-gen-go. DO NOT EDIT.
// source: micro/micro/debug/trace/proto/trace.proto

package go_micro_debug_trace

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Service describes a service running in the micro network.
type Service struct {
	// Service name, e.g. go.micro.service.greeter
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Node                 *Node    `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_6510241a5452b8ec, []int{0}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Service) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Node describes a single instance of a service.
type Node struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_6510241a5452b8ec, []int{1}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// Snapshot is a snapshot of Trace.Read from a particular service when called.
type Snapshot struct {
	// Source of the service where the snapshot was collected from
	Service *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// Trace spans
	Spans                []*Span  `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_6510241a5452b8ec, []int{2}
}

func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Snapshot.Unmarshal(m, b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
}
func (m *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(m, src)
}
func (m *Snapshot) XXX_Size() int {
	return xxx_messageInfo_Snapshot.Size(m)
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Snapshot) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

type Span struct {
	// the trace id
	Trace string `protobuf:"bytes,1,opt,name=trace,proto3" json:"trace,omitempty"`
	// id of the span
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// parent span
	Parent string `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	// name of the resource
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// time of start in nanoseconds
	Started uint64 `protobuf:"varint,5,opt,name=started,proto3" json:"started,omitempty"`
	// duration of the execution in nanoseconds
	Duration uint64 `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
	// associated metadata
	Metadata             map[string]string `protobuf:"bytes,7,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Span) Reset()         { *m = Span{} }
func (m *Span) String() string { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()    {}
func (*Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_6510241a5452b8ec, []int{3}
}

func (m *Span) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span.Unmarshal(m, b)
}
func (m *Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span.Marshal(b, m, deterministic)
}
func (m *Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span.Merge(m, src)
}
func (m *Span) XXX_Size() int {
	return xxx_messageInfo_Span.Size(m)
}
func (m *Span) XXX_DiscardUnknown() {
	xxx_messageInfo_Span.DiscardUnknown(m)
}

var xxx_messageInfo_Span proto.InternalMessageInfo

func (m *Span) GetTrace() string {
	if m != nil {
		return m.Trace
	}
	return ""
}

func (m *Span) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Span) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Span) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Span) GetStarted() uint64 {
	if m != nil {
		return m.Started
	}
	return 0
}

func (m *Span) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Span) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type ReadRequest struct {
	// If set, only return services matching the filter
	Service *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// If false, only the current snapshots will be returned.
	// If true, all historical snapshots in memory will be returned.
	Past bool `protobuf:"varint,2,opt,name=past,proto3" json:"past,omitempty"`
	// Number of traces to return
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6510241a5452b8ec, []int{4}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ReadRequest) GetPast() bool {
	if m != nil {
		return m.Past
	}
	return false
}

func (m *ReadRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ReadResponse struct {
	Spans                []*Span  `protobuf:"bytes,1,rep,name=spans,proto3" json:"spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6510241a5452b8ec, []int{5}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

type WriteRequest struct {
	// If set, only return services matching the filter
	Service *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// snapshot to write
	Stats                *Snapshot `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6510241a5452b8ec, []int{6}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *WriteRequest) GetStats() *Snapshot {
	if m != nil {
		return m.Stats
	}
	return nil
}

type WriteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6510241a5452b8ec, []int{7}
}

func (m *WriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteResponse.Unmarshal(m, b)
}
func (m *WriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteResponse.Marshal(b, m, deterministic)
}
func (m *WriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponse.Merge(m, src)
}
func (m *WriteResponse) XXX_Size() int {
	return xxx_messageInfo_WriteResponse.Size(m)
}
func (m *WriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponse proto.InternalMessageInfo

type StreamRequest struct {
	// If set, only return services matching the filter
	Service *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// If set, only return services matching the namespace
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6510241a5452b8ec, []int{8}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *StreamRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type StreamResponse struct {
	Stats                []*Snapshot `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6510241a5452b8ec, []int{9}
}

func (m *StreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse.Unmarshal(m, b)
}
func (m *StreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse.Marshal(b, m, deterministic)
}
func (m *StreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse.Merge(m, src)
}
func (m *StreamResponse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse.Size(m)
}
func (m *StreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse proto.InternalMessageInfo

func (m *StreamResponse) GetStats() []*Snapshot {
	if m != nil {
		return m.Stats
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "go.micro.debug.trace.Service")
	proto.RegisterType((*Node)(nil), "go.micro.debug.trace.Node")
	proto.RegisterType((*Snapshot)(nil), "go.micro.debug.trace.Snapshot")
	proto.RegisterType((*Span)(nil), "go.micro.debug.trace.Span")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.debug.trace.Span.MetadataEntry")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.debug.trace.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "go.micro.debug.trace.ReadResponse")
	proto.RegisterType((*WriteRequest)(nil), "go.micro.debug.trace.WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "go.micro.debug.trace.WriteResponse")
	proto.RegisterType((*StreamRequest)(nil), "go.micro.debug.trace.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "go.micro.debug.trace.StreamResponse")
}

func init() {
	proto.RegisterFile("micro/micro/debug/trace/proto/trace.proto", fileDescriptor_6510241a5452b8ec)
}

var fileDescriptor_6510241a5452b8ec = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0x26, 0x69, 0xd2, 0x76, 0x5f, 0xb7, 0xab, 0x0c, 0x8b, 0x84, 0xa0, 0x52, 0x47, 0x0f, 0xf5,
	0x92, 0x96, 0x2a, 0x28, 0x7a, 0xf1, 0xa0, 0xde, 0x56, 0x64, 0xaa, 0x78, 0x9e, 0xed, 0x3c, 0x6b,
	0x70, 0x9b, 0xc4, 0x99, 0x49, 0x61, 0x0f, 0xfe, 0x11, 0xfe, 0xc5, 0xca, 0xfc, 0x48, 0x6c, 0xa1,
	0x29, 0x0b, 0xbd, 0x94, 0xf7, 0x65, 0xbe, 0x99, 0xef, 0x7d, 0xdf, 0x7b, 0x14, 0x9e, 0x6f, 0xf2,
	0x95, 0x2c, 0x67, 0xee, 0x57, 0xe0, 0x75, 0xbd, 0x9e, 0x69, 0xc9, 0x57, 0x38, 0xab, 0x64, 0xa9,
	0x4b, 0x57, 0x67, 0xb6, 0x26, 0x97, 0xeb, 0x32, 0xb3, 0xbc, 0xcc, 0xf2, 0x32, 0x7b, 0x46, 0xd7,
	0x30, 0x58, 0xa2, 0xdc, 0xe6, 0x2b, 0x24, 0x04, 0xa2, 0x82, 0x6f, 0x30, 0x09, 0x26, 0xc1, 0xf4,
	0x8c, 0xd9, 0x9a, 0x24, 0x30, 0xd8, 0xa2, 0x54, 0x79, 0x59, 0x24, 0xa1, 0xfd, 0xdc, 0x40, 0x92,
	0x41, 0x54, 0x94, 0x02, 0x93, 0xde, 0x24, 0x98, 0x8e, 0x16, 0x69, 0x76, 0xe8, 0xf5, 0xec, 0x53,
	0x29, 0x90, 0x59, 0x1e, 0x9d, 0x43, 0x64, 0x10, 0xb9, 0x80, 0x30, 0x17, 0x5e, 0x23, 0xcc, 0x85,
	0x51, 0xe0, 0x42, 0x48, 0x54, 0xaa, 0x51, 0xf0, 0x90, 0xd6, 0x30, 0x5c, 0x16, 0xbc, 0x52, 0x3f,
	0x4a, 0x4d, 0x5e, 0xc1, 0x40, 0xb9, 0x36, 0xed, 0xd5, 0xd1, 0xe2, 0xd1, 0x61, 0x41, 0xef, 0x85,
	0x35, 0x6c, 0x32, 0x87, 0x58, 0x55, 0xbc, 0x30, 0x8f, 0xf7, 0xba, 0xfb, 0x5c, 0x56, 0xbc, 0x60,
	0x8e, 0x48, 0xff, 0x84, 0x10, 0x19, 0x4c, 0x2e, 0x21, 0xb6, 0xa7, 0xbe, 0x59, 0x07, 0x7c, 0xff,
	0x61, 0xdb, 0xff, 0x03, 0xe8, 0x57, 0x5c, 0x62, 0xa1, 0x6d, 0x12, 0x67, 0xcc, 0xa3, 0x36, 0xcd,
	0x68, 0x3f, 0x4d, 0xa5, 0xb9, 0xd4, 0x28, 0x92, 0x78, 0x12, 0x4c, 0x23, 0xd6, 0x40, 0x92, 0xc2,
	0x50, 0xd4, 0x92, 0x6b, 0x13, 0x74, 0xdf, 0x1e, 0xb5, 0x98, 0xbc, 0x87, 0xe1, 0x06, 0x35, 0x17,
	0x5c, 0xf3, 0x64, 0x60, 0x5d, 0x4c, 0xbb, 0x5d, 0x64, 0x57, 0x9e, 0xfa, 0xa1, 0xd0, 0xf2, 0x96,
	0xb5, 0x37, 0xd3, 0xb7, 0x30, 0xde, 0x3b, 0x22, 0xf7, 0xa1, 0xf7, 0x13, 0x6f, 0xbd, 0x39, 0x53,
	0x1a, 0xc3, 0x5b, 0x7e, 0x53, 0xa3, 0x77, 0xe7, 0xc0, 0x9b, 0xf0, 0x75, 0x40, 0x2b, 0x18, 0x31,
	0xe4, 0x82, 0xe1, 0xaf, 0x1a, 0xd5, 0x09, 0xd3, 0x20, 0x10, 0x55, 0x5c, 0x69, 0x2b, 0x30, 0x64,
	0xb6, 0x36, 0xaa, 0x37, 0xf9, 0x26, 0x77, 0xf9, 0xf5, 0x98, 0x03, 0xf4, 0x1d, 0x9c, 0x3b, 0x45,
	0x55, 0x95, 0x85, 0xda, 0x99, 0x63, 0x70, 0xd7, 0x39, 0xfe, 0x86, 0xf3, 0x6f, 0x32, 0xd7, 0x78,
	0x72, 0xd3, 0x2f, 0x21, 0x56, 0x9a, 0x6b, 0xb7, 0x9f, 0xa3, 0xc5, 0xe3, 0x8e, 0x6b, 0x7e, 0x55,
	0x99, 0x23, 0xd3, 0x7b, 0x30, 0xf6, 0xf2, 0xce, 0x01, 0xfd, 0x0e, 0xe3, 0xa5, 0x96, 0xc8, 0x37,
	0x27, 0x37, 0xf4, 0x10, 0xce, 0xcc, 0x3a, 0xa9, 0xca, 0x2c, 0xa7, 0x9b, 0xd5, 0xff, 0x0f, 0xf4,
	0x23, 0x5c, 0x34, 0x3a, 0x3e, 0xbb, 0xd6, 0x80, 0xcb, 0xee, 0x6e, 0x06, 0x16, 0x7f, 0x03, 0x88,
	0xbf, 0xd8, 0x95, 0xbf, 0x82, 0xc8, 0xcc, 0x82, 0x3c, 0x39, 0x7c, 0x71, 0x67, 0x33, 0x52, 0x7a,
	0x8c, 0xe2, 0xdb, 0xf9, 0x0c, 0xb1, 0x4d, 0x86, 0x74, 0x90, 0x77, 0xa7, 0x96, 0x3e, 0x3d, 0xca,
	0xf1, 0x2f, 0x7e, 0x85, 0xbe, 0xb3, 0x4c, 0x3a, 0xe8, 0x7b, 0xc1, 0xa7, 0xcf, 0x8e, 0x93, 0xdc,
	0xa3, 0xf3, 0xe0, 0xba, 0x6f, 0xff, 0x38, 0x5f, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x19, 0xff,
	0x81, 0x85, 0x65, 0x05, 0x00, 0x00,
}