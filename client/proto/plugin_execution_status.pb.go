// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the PolyForm Free Trial 1.0.0 license
// that can be found in the licenses directory at the root of this repository, also available at
// https://polyformproject.org/wp-content/uploads/2020/05/PolyForm-Free-Trial-1.0.0.txt.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: plugin_execution_status.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Execution status code
type StatusCode int32

const (
	StatusCode_CODE_UNKNOWN StatusCode = 0
	StatusCode_CODE_SUCCESS StatusCode = 1
	StatusCode_CODE_TIMEOUT StatusCode = 2
	StatusCode_CODE_FAILED  StatusCode = 3
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0: "CODE_UNKNOWN",
		1: "CODE_SUCCESS",
		2: "CODE_TIMEOUT",
		3: "CODE_FAILED",
	}
	StatusCode_value = map[string]int32{
		"CODE_UNKNOWN": 0,
		"CODE_SUCCESS": 1,
		"CODE_TIMEOUT": 2,
		"CODE_FAILED":  3,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_plugin_execution_status_proto_enumTypes[0].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_plugin_execution_status_proto_enumTypes[0]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_plugin_execution_status_proto_rawDescGZIP(), []int{0}
}

// Top level request object for plugin execution status
type ExecutionStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   *ExecutionStatus  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExecutionStatusResponse) Reset() {
	*x = ExecutionStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_execution_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionStatusResponse) ProtoMessage() {}

func (x *ExecutionStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_execution_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionStatusResponse.ProtoReflect.Descriptor instead.
func (*ExecutionStatusResponse) Descriptor() ([]byte, []int) {
	return file_plugin_execution_status_proto_rawDescGZIP(), []int{0}
}

func (x *ExecutionStatusResponse) GetStatus() *ExecutionStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ExecutionStatusResponse) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// All data that is sent from plugin execution back to manager
type ExecutionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                            // Id of the task
	Type          string               `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`                                        // Task type
	Code          StatusCode           `protobuf:"varint,3,opt,name=code,proto3,enum=proto.StatusCode" json:"code,omitempty"`                 // Execution status
	Error         string               `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`                                      // Error message if any
	ExecutionTime *durationpb.Duration `protobuf:"bytes,5,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"` // Duration of the execution
	// Types that are assignable to Data:
	//
	//	*ExecutionStatus_ProtoData
	//	*ExecutionStatus_BinaryData
	Data isExecutionStatus_Data `protobuf_oneof:"data"`
}

func (x *ExecutionStatus) Reset() {
	*x = ExecutionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_execution_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionStatus) ProtoMessage() {}

func (x *ExecutionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_execution_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionStatus.ProtoReflect.Descriptor instead.
func (*ExecutionStatus) Descriptor() ([]byte, []int) {
	return file_plugin_execution_status_proto_rawDescGZIP(), []int{1}
}

func (x *ExecutionStatus) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExecutionStatus) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ExecutionStatus) GetCode() StatusCode {
	if x != nil {
		return x.Code
	}
	return StatusCode_CODE_UNKNOWN
}

func (x *ExecutionStatus) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ExecutionStatus) GetExecutionTime() *durationpb.Duration {
	if x != nil {
		return x.ExecutionTime
	}
	return nil
}

func (m *ExecutionStatus) GetData() isExecutionStatus_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *ExecutionStatus) GetProtoData() *anypb.Any {
	if x, ok := x.GetData().(*ExecutionStatus_ProtoData); ok {
		return x.ProtoData
	}
	return nil
}

func (x *ExecutionStatus) GetBinaryData() []byte {
	if x, ok := x.GetData().(*ExecutionStatus_BinaryData); ok {
		return x.BinaryData
	}
	return nil
}

type isExecutionStatus_Data interface {
	isExecutionStatus_Data()
}

type ExecutionStatus_ProtoData struct {
	ProtoData *anypb.Any `protobuf:"bytes,6,opt,name=proto_data,json=protoData,proto3,oneof"` // for executions with proto spec.
}

type ExecutionStatus_BinaryData struct {
	BinaryData []byte `protobuf:"bytes,7,opt,name=binary_data,json=binaryData,proto3,oneof"` // for executions with any unstructured data formats including legacy kryo
}

func (*ExecutionStatus_ProtoData) isExecutionStatus_Data() {}

func (*ExecutionStatus_BinaryData) isExecutionStatus_Data() {}

var File_plugin_execution_status_proto protoreflect.FileDescriptor

var file_plugin_execution_status_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x17, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x48, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x96, 0x02, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x40, 0x0a, 0x0e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x21, 0x0a, 0x0b, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0a, 0x62, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x53, 0x0a,
	0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x43,
	0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10,
	0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x03, 0x42, 0x0e, 0x5a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_execution_status_proto_rawDescOnce sync.Once
	file_plugin_execution_status_proto_rawDescData = file_plugin_execution_status_proto_rawDesc
)

func file_plugin_execution_status_proto_rawDescGZIP() []byte {
	file_plugin_execution_status_proto_rawDescOnce.Do(func() {
		file_plugin_execution_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_execution_status_proto_rawDescData)
	})
	return file_plugin_execution_status_proto_rawDescData
}

var file_plugin_execution_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_plugin_execution_status_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_plugin_execution_status_proto_goTypes = []interface{}{
	(StatusCode)(0),                 // 0: proto.StatusCode
	(*ExecutionStatusResponse)(nil), // 1: proto.ExecutionStatusResponse
	(*ExecutionStatus)(nil),         // 2: proto.ExecutionStatus
	nil,                             // 3: proto.ExecutionStatusResponse.MetadataEntry
	(*durationpb.Duration)(nil),     // 4: google.protobuf.Duration
	(*anypb.Any)(nil),               // 5: google.protobuf.Any
}
var file_plugin_execution_status_proto_depIdxs = []int32{
	2, // 0: proto.ExecutionStatusResponse.status:type_name -> proto.ExecutionStatus
	3, // 1: proto.ExecutionStatusResponse.metadata:type_name -> proto.ExecutionStatusResponse.MetadataEntry
	0, // 2: proto.ExecutionStatus.code:type_name -> proto.StatusCode
	4, // 3: proto.ExecutionStatus.execution_time:type_name -> google.protobuf.Duration
	5, // 4: proto.ExecutionStatus.proto_data:type_name -> google.protobuf.Any
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_plugin_execution_status_proto_init() }
func file_plugin_execution_status_proto_init() {
	if File_plugin_execution_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugin_execution_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionStatusResponse); i {
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
		file_plugin_execution_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionStatus); i {
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
	file_plugin_execution_status_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ExecutionStatus_ProtoData)(nil),
		(*ExecutionStatus_BinaryData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plugin_execution_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugin_execution_status_proto_goTypes,
		DependencyIndexes: file_plugin_execution_status_proto_depIdxs,
		EnumInfos:         file_plugin_execution_status_proto_enumTypes,
		MessageInfos:      file_plugin_execution_status_proto_msgTypes,
	}.Build()
	File_plugin_execution_status_proto = out.File
	file_plugin_execution_status_proto_rawDesc = nil
	file_plugin_execution_status_proto_goTypes = nil
	file_plugin_execution_status_proto_depIdxs = nil
}
