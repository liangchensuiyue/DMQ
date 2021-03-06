// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: follower.proto

package DMQFollower

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//follower 向 header 注册服务
type FollowerRegistToHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"` // 绑定的ip 地址
	Port    int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`      // 绑定的 端口
}

func (x *FollowerRegistToHeader) Reset() {
	*x = FollowerRegistToHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowerRegistToHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowerRegistToHeader) ProtoMessage() {}

func (x *FollowerRegistToHeader) ProtoReflect() protoreflect.Message {
	mi := &file_follower_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowerRegistToHeader.ProtoReflect.Descriptor instead.
func (*FollowerRegistToHeader) Descriptor() ([]byte, []int) {
	return file_follower_proto_rawDescGZIP(), []int{0}
}

func (x *FollowerRegistToHeader) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *FollowerRegistToHeader) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

// 客户端向 follower 节点注册服务
type ClientRegistToFollower struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodeid    string `protobuf:"bytes,1,opt,name=nodeid,proto3" json:"nodeid,omitempty"`       // 节点id
	Groupname string `protobuf:"bytes,2,opt,name=groupname,proto3" json:"groupname,omitempty"` // 消费者组名
	Topic     string `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`         // 话题名称
	Key       string `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`             // 密钥
}

func (x *ClientRegistToFollower) Reset() {
	*x = ClientRegistToFollower{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRegistToFollower) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRegistToFollower) ProtoMessage() {}

func (x *ClientRegistToFollower) ProtoReflect() protoreflect.Message {
	mi := &file_follower_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRegistToFollower.ProtoReflect.Descriptor instead.
func (*ClientRegistToFollower) Descriptor() ([]byte, []int) {
	return file_follower_proto_rawDescGZIP(), []int{1}
}

func (x *ClientRegistToFollower) GetNodeid() string {
	if x != nil {
		return x.Nodeid
	}
	return ""
}

func (x *ClientRegistToFollower) GetGroupname() string {
	if x != nil {
		return x.Groupname
	}
	return ""
}

func (x *ClientRegistToFollower) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *ClientRegistToFollower) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type MessageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic   string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`     // 数据对应的topic
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 真正的数据
	Length  int64  `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
	Des     string `protobuf:"bytes,4,opt,name=des,proto3" json:"des,omitempty"` // 附加信息
	Key     string `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"` //密钥
}

func (x *MessageData) Reset() {
	*x = MessageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageData) ProtoMessage() {}

func (x *MessageData) ProtoReflect() protoreflect.Message {
	mi := &file_follower_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageData.ProtoReflect.Descriptor instead.
func (*MessageData) Descriptor() ([]byte, []int) {
	return file_follower_proto_rawDescGZIP(), []int{2}
}

func (x *MessageData) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *MessageData) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MessageData) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *MessageData) GetDes() string {
	if x != nil {
		return x.Des
	}
	return ""
}

func (x *MessageData) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// 响应信息
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errno  int32        `protobuf:"varint,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg string       `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Data   *MessageData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_follower_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_follower_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetErrno() int32 {
	if x != nil {
		return x.Errno
	}
	return 0
}

func (x *Response) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

func (x *Response) GetData() *MessageData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_follower_proto protoreflect.FileDescriptor

var file_follower_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x44, 0x4d, 0x51, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x22, 0x46, 0x0a,
	0x16, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x54,
	0x6f, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x76, 0x0a, 0x16, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x79, 0x0a,
	0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x64, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x66, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6e, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72,
	0x72, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6d,
	0x73, 0x67, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x44, 0x4d, 0x51, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0x8f, 0x02, 0x0a, 0x12, 0x44, 0x4d, 0x51, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x19, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x59, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x44, 0x4d, 0x51, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x15,
	0x2e, 0x44, 0x4d, 0x51, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x53, 0x0a, 0x11, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23,
	0x2e, 0x44, 0x4d, 0x51, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x44, 0x4d, 0x51, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x52,
	0x0a, 0x12, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x23, 0x2e, 0x44, 0x4d, 0x51, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x54,
	0x6f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x44, 0x4d, 0x51, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x34, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x6c, 0x69, 0x61, 0x6e, 0x67, 0x63,
	0x68, 0x65, 0x6e, 0x2e, 0x44, 0x4d, 0x51, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x42, 0x10, 0x44, 0x4d, 0x51, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_follower_proto_rawDescOnce sync.Once
	file_follower_proto_rawDescData = file_follower_proto_rawDesc
)

func file_follower_proto_rawDescGZIP() []byte {
	file_follower_proto_rawDescOnce.Do(func() {
		file_follower_proto_rawDescData = protoimpl.X.CompressGZIP(file_follower_proto_rawDescData)
	})
	return file_follower_proto_rawDescData
}

var file_follower_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_follower_proto_goTypes = []interface{}{
	(*FollowerRegistToHeader)(nil), // 0: DMQFollower.FollowerRegistToHeader
	(*ClientRegistToFollower)(nil), // 1: DMQFollower.ClientRegistToFollower
	(*MessageData)(nil),            // 2: DMQFollower.MessageData
	(*Response)(nil),               // 3: DMQFollower.Response
}
var file_follower_proto_depIdxs = []int32{
	2, // 0: DMQFollower.Response.data:type_name -> DMQFollower.MessageData
	2, // 1: DMQFollower.DMQFollowerService.ClientYieldMsgDataRequest:input_type -> DMQFollower.MessageData
	1, // 2: DMQFollower.DMQFollowerService.ClientConsumeData:input_type -> DMQFollower.ClientRegistToFollower
	1, // 3: DMQFollower.DMQFollowerService.ClientCloseChannel:input_type -> DMQFollower.ClientRegistToFollower
	3, // 4: DMQFollower.DMQFollowerService.ClientYieldMsgDataRequest:output_type -> DMQFollower.Response
	3, // 5: DMQFollower.DMQFollowerService.ClientConsumeData:output_type -> DMQFollower.Response
	3, // 6: DMQFollower.DMQFollowerService.ClientCloseChannel:output_type -> DMQFollower.Response
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_follower_proto_init() }
func file_follower_proto_init() {
	if File_follower_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_follower_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowerRegistToHeader); i {
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
		file_follower_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRegistToFollower); i {
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
		file_follower_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageData); i {
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
		file_follower_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_follower_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_follower_proto_goTypes,
		DependencyIndexes: file_follower_proto_depIdxs,
		MessageInfos:      file_follower_proto_msgTypes,
	}.Build()
	File_follower_proto = out.File
	file_follower_proto_rawDesc = nil
	file_follower_proto_goTypes = nil
	file_follower_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DMQFollowerServiceClient is the client API for DMQFollowerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DMQFollowerServiceClient interface {
	// 客户端向 follower 发送数据
	ClientYieldMsgDataRequest(ctx context.Context, opts ...grpc.CallOption) (DMQFollowerService_ClientYieldMsgDataRequestClient, error)
	// 客户端消费数据
	ClientConsumeData(ctx context.Context, in *ClientRegistToFollower, opts ...grpc.CallOption) (DMQFollowerService_ClientConsumeDataClient, error)
	// 客户端关闭管道
	ClientCloseChannel(ctx context.Context, in *ClientRegistToFollower, opts ...grpc.CallOption) (*Response, error)
}

type dMQFollowerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDMQFollowerServiceClient(cc grpc.ClientConnInterface) DMQFollowerServiceClient {
	return &dMQFollowerServiceClient{cc}
}

func (c *dMQFollowerServiceClient) ClientYieldMsgDataRequest(ctx context.Context, opts ...grpc.CallOption) (DMQFollowerService_ClientYieldMsgDataRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DMQFollowerService_serviceDesc.Streams[0], "/DMQFollower.DMQFollowerService/ClientYieldMsgDataRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &dMQFollowerServiceClientYieldMsgDataRequestClient{stream}
	return x, nil
}

type DMQFollowerService_ClientYieldMsgDataRequestClient interface {
	Send(*MessageData) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type dMQFollowerServiceClientYieldMsgDataRequestClient struct {
	grpc.ClientStream
}

func (x *dMQFollowerServiceClientYieldMsgDataRequestClient) Send(m *MessageData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dMQFollowerServiceClientYieldMsgDataRequestClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dMQFollowerServiceClient) ClientConsumeData(ctx context.Context, in *ClientRegistToFollower, opts ...grpc.CallOption) (DMQFollowerService_ClientConsumeDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DMQFollowerService_serviceDesc.Streams[1], "/DMQFollower.DMQFollowerService/ClientConsumeData", opts...)
	if err != nil {
		return nil, err
	}
	x := &dMQFollowerServiceClientConsumeDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DMQFollowerService_ClientConsumeDataClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type dMQFollowerServiceClientConsumeDataClient struct {
	grpc.ClientStream
}

func (x *dMQFollowerServiceClientConsumeDataClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dMQFollowerServiceClient) ClientCloseChannel(ctx context.Context, in *ClientRegistToFollower, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/DMQFollower.DMQFollowerService/ClientCloseChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DMQFollowerServiceServer is the server API for DMQFollowerService service.
type DMQFollowerServiceServer interface {
	// 客户端向 follower 发送数据
	ClientYieldMsgDataRequest(DMQFollowerService_ClientYieldMsgDataRequestServer) error
	// 客户端消费数据
	ClientConsumeData(*ClientRegistToFollower, DMQFollowerService_ClientConsumeDataServer) error
	// 客户端关闭管道
	ClientCloseChannel(context.Context, *ClientRegistToFollower) (*Response, error)
}

// UnimplementedDMQFollowerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDMQFollowerServiceServer struct {
}

func (*UnimplementedDMQFollowerServiceServer) ClientYieldMsgDataRequest(DMQFollowerService_ClientYieldMsgDataRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientYieldMsgDataRequest not implemented")
}
func (*UnimplementedDMQFollowerServiceServer) ClientConsumeData(*ClientRegistToFollower, DMQFollowerService_ClientConsumeDataServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientConsumeData not implemented")
}
func (*UnimplementedDMQFollowerServiceServer) ClientCloseChannel(context.Context, *ClientRegistToFollower) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientCloseChannel not implemented")
}

func RegisterDMQFollowerServiceServer(s *grpc.Server, srv DMQFollowerServiceServer) {
	s.RegisterService(&_DMQFollowerService_serviceDesc, srv)
}

func _DMQFollowerService_ClientYieldMsgDataRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DMQFollowerServiceServer).ClientYieldMsgDataRequest(&dMQFollowerServiceClientYieldMsgDataRequestServer{stream})
}

type DMQFollowerService_ClientYieldMsgDataRequestServer interface {
	SendAndClose(*Response) error
	Recv() (*MessageData, error)
	grpc.ServerStream
}

type dMQFollowerServiceClientYieldMsgDataRequestServer struct {
	grpc.ServerStream
}

func (x *dMQFollowerServiceClientYieldMsgDataRequestServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dMQFollowerServiceClientYieldMsgDataRequestServer) Recv() (*MessageData, error) {
	m := new(MessageData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DMQFollowerService_ClientConsumeData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClientRegistToFollower)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DMQFollowerServiceServer).ClientConsumeData(m, &dMQFollowerServiceClientConsumeDataServer{stream})
}

type DMQFollowerService_ClientConsumeDataServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type dMQFollowerServiceClientConsumeDataServer struct {
	grpc.ServerStream
}

func (x *dMQFollowerServiceClientConsumeDataServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _DMQFollowerService_ClientCloseChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientRegistToFollower)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DMQFollowerServiceServer).ClientCloseChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DMQFollower.DMQFollowerService/ClientCloseChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DMQFollowerServiceServer).ClientCloseChannel(ctx, req.(*ClientRegistToFollower))
	}
	return interceptor(ctx, in, info, handler)
}

var _DMQFollowerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DMQFollower.DMQFollowerService",
	HandlerType: (*DMQFollowerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClientCloseChannel",
			Handler:    _DMQFollowerService_ClientCloseChannel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientYieldMsgDataRequest",
			Handler:       _DMQFollowerService_ClientYieldMsgDataRequest_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ClientConsumeData",
			Handler:       _DMQFollowerService_ClientConsumeData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "follower.proto",
}
