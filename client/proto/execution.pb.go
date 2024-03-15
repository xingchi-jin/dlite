// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the PolyForm Free Trial 1.0.0 license
// that can be found in the licenses directory at the root of this repository, also available at
// https://polyformproject.org/wp-content/uploads/2020/05/PolyForm-Free-Trial-1.0.0.txt.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: execution.proto

package proto

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

type ShellType int32

const (
	ShellType_SHELL_TYPE_UNSPECIFIED ShellType = 0
	ShellType_SH                     ShellType = 1
	ShellType_BASH                   ShellType = 2
	ShellType_POWERSHELL             ShellType = 3
	ShellType_PWSH                   ShellType = 4
	ShellType_PYTHON                 ShellType = 5
)

// Enum value maps for ShellType.
var (
	ShellType_name = map[int32]string{
		0: "SHELL_TYPE_UNSPECIFIED",
		1: "SH",
		2: "BASH",
		3: "POWERSHELL",
		4: "PWSH",
		5: "PYTHON",
	}
	ShellType_value = map[string]int32{
		"SHELL_TYPE_UNSPECIFIED": 0,
		"SH":                     1,
		"BASH":                   2,
		"POWERSHELL":             3,
		"PWSH":                   4,
		"PYTHON":                 5,
	}
)

func (x ShellType) Enum() *ShellType {
	p := new(ShellType)
	*p = x
	return p
}

func (x ShellType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ShellType) Descriptor() protoreflect.EnumDescriptor {
	return file_execution_proto_enumTypes[0].Descriptor()
}

func (ShellType) Type() protoreflect.EnumType {
	return &file_execution_proto_enumTypes[0]
}

func (x ShellType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ShellType.Descriptor instead.
func (ShellType) EnumDescriptor() ([]byte, []int) {
	return file_execution_proto_rawDescGZIP(), []int{0}
}

type K8SExecution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogKey        string      `protobuf:"bytes,1,opt,name=log_key,json=logKey,proto3" json:"log_key,omitempty"`
	EntryPoint    *Entrypoint `protobuf:"bytes,2,opt,name=entry_point,json=entryPoint,proto3" json:"entry_point,omitempty"`
	CallBackToken string      `protobuf:"bytes,3,opt,name=call_back_token,json=callBackToken,proto3" json:"call_back_token,omitempty"`
	EnvVarOutputs []string    `protobuf:"bytes,4,rep,name=env_var_outputs,json=envVarOutputs,proto3" json:"env_var_outputs,omitempty"`
}

func (x *K8SExecution) Reset() {
	*x = K8SExecution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8SExecution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8SExecution) ProtoMessage() {}

func (x *K8SExecution) ProtoReflect() protoreflect.Message {
	mi := &file_execution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8SExecution.ProtoReflect.Descriptor instead.
func (*K8SExecution) Descriptor() ([]byte, []int) {
	return file_execution_proto_rawDescGZIP(), []int{0}
}

func (x *K8SExecution) GetLogKey() string {
	if x != nil {
		return x.LogKey
	}
	return ""
}

func (x *K8SExecution) GetEntryPoint() *Entrypoint {
	if x != nil {
		return x.EntryPoint
	}
	return nil
}

func (x *K8SExecution) GetCallBackToken() string {
	if x != nil {
		return x.CallBackToken
	}
	return ""
}

func (x *K8SExecution) GetEnvVarOutputs() []string {
	if x != nil {
		return x.EnvVarOutputs
	}
	return nil
}

type Entrypoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command   string    `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	ShellType ShellType `protobuf:"varint,2,opt,name=shell_type,json=shellType,proto3,enum=proto.ShellType" json:"shell_type,omitempty"`
}

func (x *Entrypoint) Reset() {
	*x = Entrypoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entrypoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entrypoint) ProtoMessage() {}

func (x *Entrypoint) ProtoReflect() protoreflect.Message {
	mi := &file_execution_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entrypoint.ProtoReflect.Descriptor instead.
func (*Entrypoint) Descriptor() ([]byte, []int) {
	return file_execution_proto_rawDescGZIP(), []int{1}
}

func (x *Entrypoint) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Entrypoint) GetShellType() ShellType {
	if x != nil {
		return x.ShellType
	}
	return ShellType_SHELL_TYPE_UNSPECIFIED
}

var File_execution_proto protoreflect.FileDescriptor

var file_execution_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x0c, 0x4b, 0x38, 0x73,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x67,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x4b,
	0x65, 0x79, 0x12, 0x32, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x62,
	0x61, 0x63, 0x6b, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26,
	0x0a, 0x0f, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x22, 0x57, 0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x2f,
	0x0a, 0x0a, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x2a,
	0x5f, 0x0a, 0x09, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16,
	0x53, 0x48, 0x45, 0x4c, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x53, 0x48, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x42, 0x41, 0x53, 0x48, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x4f,
	0x57, 0x45, 0x52, 0x53, 0x48, 0x45, 0x4c, 0x4c, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x57,
	0x53, 0x48, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x59, 0x54, 0x48, 0x4f, 0x4e, 0x10, 0x05,
	0x42, 0x0e, 0x5a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_execution_proto_rawDescOnce sync.Once
	file_execution_proto_rawDescData = file_execution_proto_rawDesc
)

func file_execution_proto_rawDescGZIP() []byte {
	file_execution_proto_rawDescOnce.Do(func() {
		file_execution_proto_rawDescData = protoimpl.X.CompressGZIP(file_execution_proto_rawDescData)
	})
	return file_execution_proto_rawDescData
}

var file_execution_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_execution_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_execution_proto_goTypes = []interface{}{
	(ShellType)(0),       // 0: proto.ShellType
	(*K8SExecution)(nil), // 1: proto.K8sExecution
	(*Entrypoint)(nil),   // 2: proto.Entrypoint
}
var file_execution_proto_depIdxs = []int32{
	2, // 0: proto.K8sExecution.entry_point:type_name -> proto.Entrypoint
	0, // 1: proto.Entrypoint.shell_type:type_name -> proto.ShellType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_execution_proto_init() }
func file_execution_proto_init() {
	if File_execution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_execution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*K8SExecution); i {
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
		file_execution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entrypoint); i {
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
			RawDescriptor: file_execution_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_execution_proto_goTypes,
		DependencyIndexes: file_execution_proto_depIdxs,
		EnumInfos:         file_execution_proto_enumTypes,
		MessageInfos:      file_execution_proto_msgTypes,
	}.Build()
	File_execution_proto = out.File
	file_execution_proto_rawDesc = nil
	file_execution_proto_goTypes = nil
	file_execution_proto_depIdxs = nil
}
