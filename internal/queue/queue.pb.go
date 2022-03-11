// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: queue.proto

package queue

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Recipient_Type int32

const (
	Recipient_EMAIL Recipient_Type = 0
	Recipient_PIPE  Recipient_Type = 1
)

// Enum value maps for Recipient_Type.
var (
	Recipient_Type_name = map[int32]string{
		0: "EMAIL",
		1: "PIPE",
	}
	Recipient_Type_value = map[string]int32{
		"EMAIL": 0,
		"PIPE":  1,
	}
)

func (x Recipient_Type) Enum() *Recipient_Type {
	p := new(Recipient_Type)
	*p = x
	return p
}

func (x Recipient_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Recipient_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_queue_proto_enumTypes[0].Descriptor()
}

func (Recipient_Type) Type() protoreflect.EnumType {
	return &file_queue_proto_enumTypes[0]
}

func (x Recipient_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Recipient_Type.Descriptor instead.
func (Recipient_Type) EnumDescriptor() ([]byte, []int) {
	return file_queue_proto_rawDescGZIP(), []int{1, 0}
}

type Recipient_Status int32

const (
	Recipient_PENDING Recipient_Status = 0
	Recipient_SENT    Recipient_Status = 1
	Recipient_FAILED  Recipient_Status = 2
)

// Enum value maps for Recipient_Status.
var (
	Recipient_Status_name = map[int32]string{
		0: "PENDING",
		1: "SENT",
		2: "FAILED",
	}
	Recipient_Status_value = map[string]int32{
		"PENDING": 0,
		"SENT":    1,
		"FAILED":  2,
	}
)

func (x Recipient_Status) Enum() *Recipient_Status {
	p := new(Recipient_Status)
	*p = x
	return p
}

func (x Recipient_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Recipient_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_queue_proto_enumTypes[1].Descriptor()
}

func (Recipient_Status) Type() protoreflect.EnumType {
	return &file_queue_proto_enumTypes[1]
}

func (x Recipient_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Recipient_Status.Descriptor instead.
func (Recipient_Status) EnumDescriptor() ([]byte, []int) {
	return file_queue_proto_rawDescGZIP(), []int{1, 1}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Message ID. Uniquely identifies this message, it is used for
	// auditing and troubleshooting.
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// The envelope for this message.
	From string       `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To   []string     `protobuf:"bytes,3,rep,name=To,proto3" json:"To,omitempty"`
	Rcpt []*Recipient `protobuf:"bytes,4,rep,name=rcpt,proto3" json:"rcpt,omitempty"`
	Data []byte       `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	// Creation timestamp.
	CreatedAtTs *Timestamp `protobuf:"bytes,6,opt,name=created_at_ts,json=createdAtTs,proto3" json:"created_at_ts,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_queue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_queue_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Message) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Message) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *Message) GetRcpt() []*Recipient {
	if x != nil {
		return x.Rcpt
	}
	return nil
}

func (x *Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Message) GetCreatedAtTs() *Timestamp {
	if x != nil {
		return x.CreatedAtTs
	}
	return nil
}

type Recipient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address to send the message to.
	// This is the final one, after expanding aliases.
	Address            string           `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Type               Recipient_Type   `protobuf:"varint,2,opt,name=type,proto3,enum=queue.Recipient_Type" json:"type,omitempty"`
	Status             Recipient_Status `protobuf:"varint,3,opt,name=status,proto3,enum=queue.Recipient_Status" json:"status,omitempty"`
	LastFailureMessage string           `protobuf:"bytes,4,opt,name=last_failure_message,json=lastFailureMessage,proto3" json:"last_failure_message,omitempty"`
	// Address that this recipient was originally intended to.
	// This is before expanding aliases and only used in very particular
	// cases.
	OriginalAddress string `protobuf:"bytes,5,opt,name=original_address,json=originalAddress,proto3" json:"original_address,omitempty"`
}

func (x *Recipient) Reset() {
	*x = Recipient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recipient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipient) ProtoMessage() {}

func (x *Recipient) ProtoReflect() protoreflect.Message {
	mi := &file_queue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recipient.ProtoReflect.Descriptor instead.
func (*Recipient) Descriptor() ([]byte, []int) {
	return file_queue_proto_rawDescGZIP(), []int{1}
}

func (x *Recipient) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Recipient) GetType() Recipient_Type {
	if x != nil {
		return x.Type
	}
	return Recipient_EMAIL
}

func (x *Recipient) GetStatus() Recipient_Status {
	if x != nil {
		return x.Status
	}
	return Recipient_PENDING
}

func (x *Recipient) GetLastFailureMessage() string {
	if x != nil {
		return x.LastFailureMessage
	}
	return ""
}

func (x *Recipient) GetOriginalAddress() string {
	if x != nil {
		return x.OriginalAddress
	}
	return ""
}

// Timestamp representation, for convenience.
// We used to use the well-known type, but the dependency makes packaging much
// more convoluted and adds very little value, so we now just include it here.
type Timestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Represents seconds of UTC time since Unix epoch.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution.
	Nanos int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *Timestamp) Reset() {
	*x = Timestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_queue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timestamp.ProtoReflect.Descriptor instead.
func (*Timestamp) Descriptor() ([]byte, []int) {
	return file_queue_proto_rawDescGZIP(), []int{2}
}

func (x *Timestamp) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Timestamp) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

var File_queue_proto protoreflect.FileDescriptor

var file_queue_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x02, 0x54, 0x6f, 0x12, 0x24, 0x0a, 0x04, 0x72, 0x63, 0x70, 0x74, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x72, 0x63, 0x70, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x34,
	0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x74, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x54, 0x73, 0x22, 0xa8, 0x02, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x1b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x49, 0x50, 0x45,
	0x10, 0x01, 0x22, 0x2b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07,
	0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x4e,
	0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x22,
	0x3b, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x42, 0x2b, 0x5a, 0x29,
	0x62, 0x6c, 0x69, 0x74, 0x69, 0x72, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x72, 0x2f, 0x67,
	0x6f, 0x2f, 0x63, 0x68, 0x61, 0x73, 0x71, 0x75, 0x69, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_queue_proto_rawDescOnce sync.Once
	file_queue_proto_rawDescData = file_queue_proto_rawDesc
)

func file_queue_proto_rawDescGZIP() []byte {
	file_queue_proto_rawDescOnce.Do(func() {
		file_queue_proto_rawDescData = protoimpl.X.CompressGZIP(file_queue_proto_rawDescData)
	})
	return file_queue_proto_rawDescData
}

var file_queue_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_queue_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_queue_proto_goTypes = []interface{}{
	(Recipient_Type)(0),   // 0: queue.Recipient.Type
	(Recipient_Status)(0), // 1: queue.Recipient.Status
	(*Message)(nil),       // 2: queue.Message
	(*Recipient)(nil),     // 3: queue.Recipient
	(*Timestamp)(nil),     // 4: queue.Timestamp
}
var file_queue_proto_depIdxs = []int32{
	3, // 0: queue.Message.rcpt:type_name -> queue.Recipient
	4, // 1: queue.Message.created_at_ts:type_name -> queue.Timestamp
	0, // 2: queue.Recipient.type:type_name -> queue.Recipient.Type
	1, // 3: queue.Recipient.status:type_name -> queue.Recipient.Status
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_queue_proto_init() }
func file_queue_proto_init() {
	if File_queue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_queue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_queue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recipient); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_queue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timestamp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_queue_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_queue_proto_goTypes,
		DependencyIndexes: file_queue_proto_depIdxs,
		EnumInfos:         file_queue_proto_enumTypes,
		MessageInfos:      file_queue_proto_msgTypes,
	}.Build()
	File_queue_proto = out.File
	file_queue_proto_rawDesc = nil
	file_queue_proto_goTypes = nil
	file_queue_proto_depIdxs = nil
}
