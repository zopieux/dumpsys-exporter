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
// source: frameworks/base/core/proto/android/os/worksource.proto

package os

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

type WorkSourceProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkSourceContents []*WorkSourceProto_WorkSourceContentProto `protobuf:"bytes,1,rep,name=work_source_contents,json=workSourceContents" json:"work_source_contents,omitempty"`
	WorkChains         []*WorkSourceProto_WorkChain              `protobuf:"bytes,2,rep,name=work_chains,json=workChains" json:"work_chains,omitempty"`
}

func (x *WorkSourceProto) Reset() {
	*x = WorkSourceProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frameworks_base_core_proto_android_os_worksource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkSourceProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkSourceProto) ProtoMessage() {}

func (x *WorkSourceProto) ProtoReflect() protoreflect.Message {
	mi := &file_frameworks_base_core_proto_android_os_worksource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkSourceProto.ProtoReflect.Descriptor instead.
func (*WorkSourceProto) Descriptor() ([]byte, []int) {
	return file_frameworks_base_core_proto_android_os_worksource_proto_rawDescGZIP(), []int{0}
}

func (x *WorkSourceProto) GetWorkSourceContents() []*WorkSourceProto_WorkSourceContentProto {
	if x != nil {
		return x.WorkSourceContents
	}
	return nil
}

func (x *WorkSourceProto) GetWorkChains() []*WorkSourceProto_WorkChain {
	if x != nil {
		return x.WorkChains
	}
	return nil
}

type WorkSourceProto_WorkSourceContentProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  *int32  `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (x *WorkSourceProto_WorkSourceContentProto) Reset() {
	*x = WorkSourceProto_WorkSourceContentProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frameworks_base_core_proto_android_os_worksource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkSourceProto_WorkSourceContentProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkSourceProto_WorkSourceContentProto) ProtoMessage() {}

func (x *WorkSourceProto_WorkSourceContentProto) ProtoReflect() protoreflect.Message {
	mi := &file_frameworks_base_core_proto_android_os_worksource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkSourceProto_WorkSourceContentProto.ProtoReflect.Descriptor instead.
func (*WorkSourceProto_WorkSourceContentProto) Descriptor() ([]byte, []int) {
	return file_frameworks_base_core_proto_android_os_worksource_proto_rawDescGZIP(), []int{0, 0}
}

func (x *WorkSourceProto_WorkSourceContentProto) GetUid() int32 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *WorkSourceProto_WorkSourceContentProto) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type WorkSourceProto_WorkChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*WorkSourceProto_WorkSourceContentProto `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (x *WorkSourceProto_WorkChain) Reset() {
	*x = WorkSourceProto_WorkChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frameworks_base_core_proto_android_os_worksource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkSourceProto_WorkChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkSourceProto_WorkChain) ProtoMessage() {}

func (x *WorkSourceProto_WorkChain) ProtoReflect() protoreflect.Message {
	mi := &file_frameworks_base_core_proto_android_os_worksource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkSourceProto_WorkChain.ProtoReflect.Descriptor instead.
func (*WorkSourceProto_WorkChain) Descriptor() ([]byte, []int) {
	return file_frameworks_base_core_proto_android_os_worksource_proto_rawDescGZIP(), []int{0, 1}
}

func (x *WorkSourceProto_WorkChain) GetNodes() []*WorkSourceProto_WorkSourceContentProto {
	if x != nil {
		return x.Nodes
	}
	return nil
}

var File_frameworks_base_core_proto_android_os_worksource_proto protoreflect.FileDescriptor

var file_frameworks_base_core_proto_android_os_worksource_proto_rawDesc = []byte{
	0x0a, 0x36, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x64,
	0x72, 0x6f, 0x69, 0x64, 0x2f, 0x6f, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69,
	0x64, 0x2e, 0x6f, 0x73, 0x1a, 0x30, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x02, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x64, 0x0a, 0x14, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x61, 0x6e, 0x64, 0x72, 0x6f,
	0x69, 0x64, 0x2e, 0x6f, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x12, 0x77, 0x6f,
	0x72, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x46, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x2e,
	0x6f, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x0a, 0x77, 0x6f,
	0x72, 0x6b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x1a, 0x49, 0x0a, 0x16, 0x57, 0x6f, 0x72, 0x6b,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x09, 0x9a, 0x9f, 0xd5, 0x87, 0x03, 0x03,
	0x08, 0xc8, 0x01, 0x1a, 0x60, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x6b, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x12, 0x48, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x2e, 0x6f, 0x73, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x3a, 0x09, 0x9a, 0x9f, 0xd5, 0x87,
	0x03, 0x03, 0x08, 0xc8, 0x01, 0x3a, 0x09, 0x9a, 0x9f, 0xd5, 0x87, 0x03, 0x03, 0x08, 0xc8, 0x01,
	0x42, 0x02, 0x50, 0x01,
}

var (
	file_frameworks_base_core_proto_android_os_worksource_proto_rawDescOnce sync.Once
	file_frameworks_base_core_proto_android_os_worksource_proto_rawDescData = file_frameworks_base_core_proto_android_os_worksource_proto_rawDesc
)

func file_frameworks_base_core_proto_android_os_worksource_proto_rawDescGZIP() []byte {
	file_frameworks_base_core_proto_android_os_worksource_proto_rawDescOnce.Do(func() {
		file_frameworks_base_core_proto_android_os_worksource_proto_rawDescData = protoimpl.X.CompressGZIP(file_frameworks_base_core_proto_android_os_worksource_proto_rawDescData)
	})
	return file_frameworks_base_core_proto_android_os_worksource_proto_rawDescData
}

var file_frameworks_base_core_proto_android_os_worksource_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_frameworks_base_core_proto_android_os_worksource_proto_goTypes = []any{
	(*WorkSourceProto)(nil),                        // 0: android.os.WorkSourceProto
	(*WorkSourceProto_WorkSourceContentProto)(nil), // 1: android.os.WorkSourceProto.WorkSourceContentProto
	(*WorkSourceProto_WorkChain)(nil),              // 2: android.os.WorkSourceProto.WorkChain
}
var file_frameworks_base_core_proto_android_os_worksource_proto_depIdxs = []int32{
	1, // 0: android.os.WorkSourceProto.work_source_contents:type_name -> android.os.WorkSourceProto.WorkSourceContentProto
	2, // 1: android.os.WorkSourceProto.work_chains:type_name -> android.os.WorkSourceProto.WorkChain
	1, // 2: android.os.WorkSourceProto.WorkChain.nodes:type_name -> android.os.WorkSourceProto.WorkSourceContentProto
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_frameworks_base_core_proto_android_os_worksource_proto_init() }
func file_frameworks_base_core_proto_android_os_worksource_proto_init() {
	if File_frameworks_base_core_proto_android_os_worksource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_frameworks_base_core_proto_android_os_worksource_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*WorkSourceProto); i {
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
		file_frameworks_base_core_proto_android_os_worksource_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*WorkSourceProto_WorkSourceContentProto); i {
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
		file_frameworks_base_core_proto_android_os_worksource_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*WorkSourceProto_WorkChain); i {
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
			RawDescriptor: file_frameworks_base_core_proto_android_os_worksource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_frameworks_base_core_proto_android_os_worksource_proto_goTypes,
		DependencyIndexes: file_frameworks_base_core_proto_android_os_worksource_proto_depIdxs,
		MessageInfos:      file_frameworks_base_core_proto_android_os_worksource_proto_msgTypes,
	}.Build()
	File_frameworks_base_core_proto_android_os_worksource_proto = out.File
	file_frameworks_base_core_proto_android_os_worksource_proto_rawDesc = nil
	file_frameworks_base_core_proto_android_os_worksource_proto_goTypes = nil
	file_frameworks_base_core_proto_android_os_worksource_proto_depIdxs = nil
}
