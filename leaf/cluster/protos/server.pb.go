// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: server.proto

package protos

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

type ServerType int32

const (
	ServerType_NULL ServerType = 0
	// 唯一节点
	ServerType_Master ServerType = 1
	// 可部署多个节点
	ServerType_Gate ServerType = 2
	// 非网关的所有节点
	ServerType_Node ServerType = 3
)

// Enum value maps for ServerType.
var (
	ServerType_name = map[int32]string{
		0: "NULL",
		1: "Master",
		2: "Gate",
		3: "Node",
	}
	ServerType_value = map[string]int32{
		"NULL":   0,
		"Master": 1,
		"Gate":   2,
		"Node":   3,
	}
)

func (x ServerType) Enum() *ServerType {
	p := new(ServerType)
	*p = x
	return p
}

func (x ServerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServerType) Descriptor() protoreflect.EnumDescriptor {
	return file_server_proto_enumTypes[0].Descriptor()
}

func (ServerType) Type() protoreflect.EnumType {
	return &file_server_proto_enumTypes[0]
}

func (x ServerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerType.Descriptor instead.
func (ServerType) EnumDescriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      int32      `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Address string     `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	Typ     ServerType `protobuf:"varint,3,opt,name=Typ,proto3,enum=actor.ServerType" json:"Typ,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

func (x *Server) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Server) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Server) GetTyp() ServerType {
	if x != nil {
		return x.Typ
	}
	return ServerType_NULL
}

type ServerMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Route string `protobuf:"bytes,2,opt,name=Route,proto3" json:"Route,omitempty"`
	Body  string `protobuf:"bytes,3,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *ServerMsg) Reset() {
	*x = ServerMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerMsg) ProtoMessage() {}

func (x *ServerMsg) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerMsg.ProtoReflect.Descriptor instead.
func (*ServerMsg) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *ServerMsg) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ServerMsg) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

func (x *ServerMsg) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Register struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server *Server `protobuf:"bytes,1,opt,name=Server,proto3" json:"Server,omitempty"`
}

func (x *Register) Reset() {
	*x = Register{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Register) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Register) ProtoMessage() {}

func (x *Register) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Register.ProtoReflect.Descriptor instead.
func (*Register) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{2}
}

func (x *Register) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

type WatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Servers []*Server `protobuf:"bytes,1,rep,name=Servers,proto3" json:"Servers,omitempty"`
}

func (x *WatchResponse) Reset() {
	*x = WatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResponse) ProtoMessage() {}

func (x *WatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResponse.ProtoReflect.Descriptor instead.
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{3}
}

func (x *WatchResponse) GetServers() []*Server {
	if x != nil {
		return x.Servers
	}
	return nil
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x57, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x03, 0x54, 0x79, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x03, 0x54, 0x79, 0x70, 0x22, 0x45,
	0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x31, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x25, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x0d, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x07, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x2a, 0x36, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x61, 0x74, 0x65, 0x10, 0x02,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x10, 0x03, 0x42, 0x0f, 0x5a, 0x0d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_server_proto_goTypes = []interface{}{
	(ServerType)(0),       // 0: actor.ServerType
	(*Server)(nil),        // 1: actor.Server
	(*ServerMsg)(nil),     // 2: actor.ServerMsg
	(*Register)(nil),      // 3: actor.Register
	(*WatchResponse)(nil), // 4: actor.WatchResponse
}
var file_server_proto_depIdxs = []int32{
	0, // 0: actor.Server.Typ:type_name -> actor.ServerType
	1, // 1: actor.Register.Server:type_name -> actor.Server
	1, // 2: actor.WatchResponse.Servers:type_name -> actor.Server
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerMsg); i {
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
		file_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Register); i {
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
		file_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchResponse); i {
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
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		EnumInfos:         file_server_proto_enumTypes,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}
