// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: wso2/discovery/api/security_scheme.proto

package api

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

// SecurityScheme config model
type SecurityScheme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefinitionName string `protobuf:"bytes,1,opt,name=definitionName,proto3" json:"definitionName,omitempty"` // name used to define security scheme
	Type           string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`                     // type of the security scheme
	Name           string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`                     // name of the security scheme
	In             string `protobuf:"bytes,4,opt,name=in,proto3" json:"in,omitempty"`                         // location of the API key in request
}

func (x *SecurityScheme) Reset() {
	*x = SecurityScheme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_security_scheme_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityScheme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityScheme) ProtoMessage() {}

func (x *SecurityScheme) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_security_scheme_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityScheme.ProtoReflect.Descriptor instead.
func (*SecurityScheme) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_security_scheme_proto_rawDescGZIP(), []int{0}
}

func (x *SecurityScheme) GetDefinitionName() string {
	if x != nil {
		return x.DefinitionName
	}
	return ""
}

func (x *SecurityScheme) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SecurityScheme) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecurityScheme) GetIn() string {
	if x != nil {
		return x.In
	}
	return ""
}

// Represents a single security array item applied at the API level or the API operation level
type SecurityList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScopeList map[string]*Scopes `protobuf:"bytes,1,rep,name=scopeList,proto3" json:"scopeList,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SecurityList) Reset() {
	*x = SecurityList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_security_scheme_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityList) ProtoMessage() {}

func (x *SecurityList) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_security_scheme_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityList.ProtoReflect.Descriptor instead.
func (*SecurityList) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_security_scheme_proto_rawDescGZIP(), []int{1}
}

func (x *SecurityList) GetScopeList() map[string]*Scopes {
	if x != nil {
		return x.ScopeList
	}
	return nil
}

// Holds the list of scopes belongs to a particular security schema
type Scopes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scopes []string `protobuf:"bytes,1,rep,name=scopes,proto3" json:"scopes,omitempty"`
}

func (x *Scopes) Reset() {
	*x = Scopes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_security_scheme_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scopes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scopes) ProtoMessage() {}

func (x *Scopes) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_security_scheme_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scopes.ProtoReflect.Descriptor instead.
func (*Scopes) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_security_scheme_proto_rawDescGZIP(), []int{2}
}

func (x *Scopes) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

var File_wso2_discovery_api_security_scheme_proto protoreflect.FileDescriptor

var file_wso2_discovery_api_security_scheme_proto_rawDesc = []byte{
	0x0a, 0x28, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x77, 0x73, 0x6f, 0x32,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x70,
	0x0a, 0x0e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x6e,
	0x22, 0xb7, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x4d, 0x0a, 0x09, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x58, 0x0a, 0x0e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x42, 0x7d, 0x0a, 0x25,
	0x6f, 0x72, 0x67, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x65, 0x6f, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x13, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_wso2_discovery_api_security_scheme_proto_rawDescOnce sync.Once
	file_wso2_discovery_api_security_scheme_proto_rawDescData = file_wso2_discovery_api_security_scheme_proto_rawDesc
)

func file_wso2_discovery_api_security_scheme_proto_rawDescGZIP() []byte {
	file_wso2_discovery_api_security_scheme_proto_rawDescOnce.Do(func() {
		file_wso2_discovery_api_security_scheme_proto_rawDescData = protoimpl.X.CompressGZIP(file_wso2_discovery_api_security_scheme_proto_rawDescData)
	})
	return file_wso2_discovery_api_security_scheme_proto_rawDescData
}

var file_wso2_discovery_api_security_scheme_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_wso2_discovery_api_security_scheme_proto_goTypes = []interface{}{
	(*SecurityScheme)(nil), // 0: wso2.discovery.api.SecurityScheme
	(*SecurityList)(nil),   // 1: wso2.discovery.api.SecurityList
	(*Scopes)(nil),         // 2: wso2.discovery.api.Scopes
	nil,                    // 3: wso2.discovery.api.SecurityList.ScopeListEntry
}
var file_wso2_discovery_api_security_scheme_proto_depIdxs = []int32{
	3, // 0: wso2.discovery.api.SecurityList.scopeList:type_name -> wso2.discovery.api.SecurityList.ScopeListEntry
	2, // 1: wso2.discovery.api.SecurityList.ScopeListEntry.value:type_name -> wso2.discovery.api.Scopes
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wso2_discovery_api_security_scheme_proto_init() }
func file_wso2_discovery_api_security_scheme_proto_init() {
	if File_wso2_discovery_api_security_scheme_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wso2_discovery_api_security_scheme_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityScheme); i {
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
		file_wso2_discovery_api_security_scheme_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityList); i {
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
		file_wso2_discovery_api_security_scheme_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scopes); i {
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
			RawDescriptor: file_wso2_discovery_api_security_scheme_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wso2_discovery_api_security_scheme_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_api_security_scheme_proto_depIdxs,
		MessageInfos:      file_wso2_discovery_api_security_scheme_proto_msgTypes,
	}.Build()
	File_wso2_discovery_api_security_scheme_proto = out.File
	file_wso2_discovery_api_security_scheme_proto_rawDesc = nil
	file_wso2_discovery_api_security_scheme_proto_goTypes = nil
	file_wso2_discovery_api_security_scheme_proto_depIdxs = nil
}
