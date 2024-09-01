//
// Copyright (C) 2017 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.4
// source: frameworks/base/core/proto/android/util/common.proto

package util

import (
	_ "github.com/zopieux/dumpsys-exporter/proto/android/frameworks/base/core/proto/android"
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

// *
// Very basic data structure used by aggregated stats.
type AggStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Marked as deprecated in frameworks/base/core/proto/android/util/common.proto.
	Min *int64 `protobuf:"varint,1,opt,name=min" json:"min,omitempty"`
	// Deprecated: Marked as deprecated in frameworks/base/core/proto/android/util/common.proto.
	Average *int64 `protobuf:"varint,2,opt,name=average" json:"average,omitempty"`
	// Deprecated: Marked as deprecated in frameworks/base/core/proto/android/util/common.proto.
	Max *int64 `protobuf:"varint,3,opt,name=max" json:"max,omitempty"`
	// These are all in kilobyte resolution. Can fit in int32, so smaller on the wire than the above
	// int64 fields.
	MeanKb *int32 `protobuf:"varint,4,opt,name=mean_kb,json=meanKb" json:"mean_kb,omitempty"`
	MaxKb  *int32 `protobuf:"varint,5,opt,name=max_kb,json=maxKb" json:"max_kb,omitempty"`
}

func (x *AggStats) Reset() {
	*x = AggStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frameworks_base_core_proto_android_util_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggStats) ProtoMessage() {}

func (x *AggStats) ProtoReflect() protoreflect.Message {
	mi := &file_frameworks_base_core_proto_android_util_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggStats.ProtoReflect.Descriptor instead.
func (*AggStats) Descriptor() ([]byte, []int) {
	return file_frameworks_base_core_proto_android_util_common_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in frameworks/base/core/proto/android/util/common.proto.
func (x *AggStats) GetMin() int64 {
	if x != nil && x.Min != nil {
		return *x.Min
	}
	return 0
}

// Deprecated: Marked as deprecated in frameworks/base/core/proto/android/util/common.proto.
func (x *AggStats) GetAverage() int64 {
	if x != nil && x.Average != nil {
		return *x.Average
	}
	return 0
}

// Deprecated: Marked as deprecated in frameworks/base/core/proto/android/util/common.proto.
func (x *AggStats) GetMax() int64 {
	if x != nil && x.Max != nil {
		return *x.Max
	}
	return 0
}

func (x *AggStats) GetMeanKb() int32 {
	if x != nil && x.MeanKb != nil {
		return *x.MeanKb
	}
	return 0
}

func (x *AggStats) GetMaxKb() int32 {
	if x != nil && x.MaxKb != nil {
		return *x.MaxKb
	}
	return 0
}

// *
// Very basic data structure to represent Duration.
type Duration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartMs *int64 `protobuf:"varint,1,opt,name=start_ms,json=startMs" json:"start_ms,omitempty"`
	EndMs   *int64 `protobuf:"varint,2,opt,name=end_ms,json=endMs" json:"end_ms,omitempty"`
}

func (x *Duration) Reset() {
	*x = Duration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frameworks_base_core_proto_android_util_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Duration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Duration) ProtoMessage() {}

func (x *Duration) ProtoReflect() protoreflect.Message {
	mi := &file_frameworks_base_core_proto_android_util_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Duration.ProtoReflect.Descriptor instead.
func (*Duration) Descriptor() ([]byte, []int) {
	return file_frameworks_base_core_proto_android_util_common_proto_rawDescGZIP(), []int{1}
}

func (x *Duration) GetStartMs() int64 {
	if x != nil && x.StartMs != nil {
		return *x.StartMs
	}
	return 0
}

func (x *Duration) GetEndMs() int64 {
	if x != nil && x.EndMs != nil {
		return *x.EndMs
	}
	return 0
}

var File_frameworks_base_core_proto_android_util_common_proto protoreflect.FileDescriptor

var file_frameworks_base_core_proto_android_util_common_proto_rawDesc = []byte{
	0x0a, 0x34, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x64,
	0x72, 0x6f, 0x69, 0x64, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x2e,
	0x75, 0x74, 0x69, 0x6c, 0x1a, 0x30, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x08, 0x41, 0x67, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x02, 0x18, 0x01, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x07, 0x61, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07,
	0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x18, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x17, 0x0a,
	0x07, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x6b, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6d, 0x65, 0x61, 0x6e, 0x4b, 0x62, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x5f, 0x6b, 0x62,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61, 0x78, 0x4b, 0x62, 0x3a, 0x09, 0x9a,
	0x9f, 0xd5, 0x87, 0x03, 0x03, 0x08, 0xc8, 0x01, 0x22, 0x47, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x73, 0x12,
	0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x3a, 0x09, 0x9a, 0x9f, 0xd5, 0x87, 0x03, 0x03, 0x08, 0xc8,
	0x01, 0x42, 0x02, 0x50, 0x01,
}

var (
	file_frameworks_base_core_proto_android_util_common_proto_rawDescOnce sync.Once
	file_frameworks_base_core_proto_android_util_common_proto_rawDescData = file_frameworks_base_core_proto_android_util_common_proto_rawDesc
)

func file_frameworks_base_core_proto_android_util_common_proto_rawDescGZIP() []byte {
	file_frameworks_base_core_proto_android_util_common_proto_rawDescOnce.Do(func() {
		file_frameworks_base_core_proto_android_util_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_frameworks_base_core_proto_android_util_common_proto_rawDescData)
	})
	return file_frameworks_base_core_proto_android_util_common_proto_rawDescData
}

var file_frameworks_base_core_proto_android_util_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_frameworks_base_core_proto_android_util_common_proto_goTypes = []any{
	(*AggStats)(nil), // 0: android.util.AggStats
	(*Duration)(nil), // 1: android.util.Duration
}
var file_frameworks_base_core_proto_android_util_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_frameworks_base_core_proto_android_util_common_proto_init() }
func file_frameworks_base_core_proto_android_util_common_proto_init() {
	if File_frameworks_base_core_proto_android_util_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_frameworks_base_core_proto_android_util_common_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AggStats); i {
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
		file_frameworks_base_core_proto_android_util_common_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Duration); i {
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
			RawDescriptor: file_frameworks_base_core_proto_android_util_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_frameworks_base_core_proto_android_util_common_proto_goTypes,
		DependencyIndexes: file_frameworks_base_core_proto_android_util_common_proto_depIdxs,
		MessageInfos:      file_frameworks_base_core_proto_android_util_common_proto_msgTypes,
	}.Build()
	File_frameworks_base_core_proto_android_util_common_proto = out.File
	file_frameworks_base_core_proto_android_util_common_proto_rawDesc = nil
	file_frameworks_base_core_proto_android_util_common_proto_goTypes = nil
	file_frameworks_base_core_proto_android_util_common_proto_depIdxs = nil
}
