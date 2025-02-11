//
// Copyright (C) 2018 The Android Open Source Project
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
// source: frameworks/proto_logging/stats/enums/service/procstats_enum.proto

package service

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

type ScreenState int32

const (
	ScreenState_SCREEN_STATE_UNKNOWN ScreenState = 0
	ScreenState_SCREEN_STATE_OFF     ScreenState = 1
	ScreenState_SCREEN_STATE_ON      ScreenState = 2
)

// Enum value maps for ScreenState.
var (
	ScreenState_name = map[int32]string{
		0: "SCREEN_STATE_UNKNOWN",
		1: "SCREEN_STATE_OFF",
		2: "SCREEN_STATE_ON",
	}
	ScreenState_value = map[string]int32{
		"SCREEN_STATE_UNKNOWN": 0,
		"SCREEN_STATE_OFF":     1,
		"SCREEN_STATE_ON":      2,
	}
)

func (x ScreenState) Enum() *ScreenState {
	p := new(ScreenState)
	*p = x
	return p
}

func (x ScreenState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScreenState) Descriptor() protoreflect.EnumDescriptor {
	return file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_enumTypes[0].Descriptor()
}

func (ScreenState) Type() protoreflect.EnumType {
	return &file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_enumTypes[0]
}

func (x ScreenState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ScreenState) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ScreenState(num)
	return nil
}

// Deprecated: Use ScreenState.Descriptor instead.
func (ScreenState) EnumDescriptor() ([]byte, []int) {
	return file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_rawDescGZIP(), []int{0}
}

type MemoryState int32

const (
	MemoryState_MEMORY_STATE_UNKNOWN  MemoryState = 0
	MemoryState_MEMORY_STATE_NORMAL   MemoryState = 1 // normal.
	MemoryState_MEMORY_STATE_MODERATE MemoryState = 2 // moderate memory pressure.
	MemoryState_MEMORY_STATE_LOW      MemoryState = 3 // low memory.
	MemoryState_MEMORY_STATE_CRITICAL MemoryState = 4 // critical memory.
)

// Enum value maps for MemoryState.
var (
	MemoryState_name = map[int32]string{
		0: "MEMORY_STATE_UNKNOWN",
		1: "MEMORY_STATE_NORMAL",
		2: "MEMORY_STATE_MODERATE",
		3: "MEMORY_STATE_LOW",
		4: "MEMORY_STATE_CRITICAL",
	}
	MemoryState_value = map[string]int32{
		"MEMORY_STATE_UNKNOWN":  0,
		"MEMORY_STATE_NORMAL":   1,
		"MEMORY_STATE_MODERATE": 2,
		"MEMORY_STATE_LOW":      3,
		"MEMORY_STATE_CRITICAL": 4,
	}
)

func (x MemoryState) Enum() *MemoryState {
	p := new(MemoryState)
	*p = x
	return p
}

func (x MemoryState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MemoryState) Descriptor() protoreflect.EnumDescriptor {
	return file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_enumTypes[1].Descriptor()
}

func (MemoryState) Type() protoreflect.EnumType {
	return &file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_enumTypes[1]
}

func (x MemoryState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *MemoryState) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = MemoryState(num)
	return nil
}

// Deprecated: Use MemoryState.Descriptor instead.
func (MemoryState) EnumDescriptor() ([]byte, []int) {
	return file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_rawDescGZIP(), []int{1}
}

// this enum list is from frameworks/base/core/java/com/android/internal/app/procstats/ProcessStats.java
// and not frameworks/base/core/java/android/app/ActivityManager.java
// Next id: 18
type ProcessState int32

const (
	ProcessState_PROCESS_STATE_UNKNOWN ProcessState = 0
	// Persistent system process.
	ProcessState_PROCESS_STATE_PERSISTENT ProcessState = 1
	// Top activity; actually any visible activity.
	ProcessState_PROCESS_STATE_TOP ProcessState = 2
	// Process bound by top or a foreground service
	ProcessState_PROCESS_STATE_BOUND_TOP_OR_FGS ProcessState = 15
	// Processing running a foreground service.
	ProcessState_PROCESS_STATE_FGS ProcessState = 16
	// Important foreground process (ime, wallpaper, etc).
	ProcessState_PROCESS_STATE_IMPORTANT_FOREGROUND ProcessState = 3
	// Important background process.
	ProcessState_PROCESS_STATE_IMPORTANT_BACKGROUND ProcessState = 4
	// Performing backup operation.
	ProcessState_PROCESS_STATE_BACKUP ProcessState = 5
	// Background process running a service.
	ProcessState_PROCESS_STATE_SERVICE ProcessState = 6
	// Process not running, but would be if there was enough RAM.
	ProcessState_PROCESS_STATE_SERVICE_RESTARTING ProcessState = 7
	// Process running a receiver.
	ProcessState_PROCESS_STATE_RECEIVER ProcessState = 8
	// Heavy-weight process (currently not used).
	ProcessState_PROCESS_STATE_HEAVY_WEIGHT ProcessState = 9
	// Process hosting home/launcher app when not on top.
	ProcessState_PROCESS_STATE_HOME ProcessState = 10
	// Process hosting the last app the user was in.
	ProcessState_PROCESS_STATE_LAST_ACTIVITY ProcessState = 11
	// Cached process hosting a previous activity.
	ProcessState_PROCESS_STATE_CACHED_ACTIVITY ProcessState = 12
	// Cached process hosting a client activity.
	ProcessState_PROCESS_STATE_CACHED_ACTIVITY_CLIENT ProcessState = 13
	// Cached process that is empty.
	ProcessState_PROCESS_STATE_CACHED_EMPTY ProcessState = 14
	// Frozen process.
	ProcessState_PROCESS_STATE_FROZEN ProcessState = 17
	// Represents all cached states.
	ProcessState_PROCESS_STATE_CACHED ProcessState = 18
	// Process bound by a top service
	ProcessState_PROCESS_STATE_BOUND_TOP ProcessState = 19
	// Process bound by foreground service
	ProcessState_PROCESS_STATE_BOUND_FGS ProcessState = 20
)

// Enum value maps for ProcessState.
var (
	ProcessState_name = map[int32]string{
		0:  "PROCESS_STATE_UNKNOWN",
		1:  "PROCESS_STATE_PERSISTENT",
		2:  "PROCESS_STATE_TOP",
		15: "PROCESS_STATE_BOUND_TOP_OR_FGS",
		16: "PROCESS_STATE_FGS",
		3:  "PROCESS_STATE_IMPORTANT_FOREGROUND",
		4:  "PROCESS_STATE_IMPORTANT_BACKGROUND",
		5:  "PROCESS_STATE_BACKUP",
		6:  "PROCESS_STATE_SERVICE",
		7:  "PROCESS_STATE_SERVICE_RESTARTING",
		8:  "PROCESS_STATE_RECEIVER",
		9:  "PROCESS_STATE_HEAVY_WEIGHT",
		10: "PROCESS_STATE_HOME",
		11: "PROCESS_STATE_LAST_ACTIVITY",
		12: "PROCESS_STATE_CACHED_ACTIVITY",
		13: "PROCESS_STATE_CACHED_ACTIVITY_CLIENT",
		14: "PROCESS_STATE_CACHED_EMPTY",
		17: "PROCESS_STATE_FROZEN",
		18: "PROCESS_STATE_CACHED",
		19: "PROCESS_STATE_BOUND_TOP",
		20: "PROCESS_STATE_BOUND_FGS",
	}
	ProcessState_value = map[string]int32{
		"PROCESS_STATE_UNKNOWN":                0,
		"PROCESS_STATE_PERSISTENT":             1,
		"PROCESS_STATE_TOP":                    2,
		"PROCESS_STATE_BOUND_TOP_OR_FGS":       15,
		"PROCESS_STATE_FGS":                    16,
		"PROCESS_STATE_IMPORTANT_FOREGROUND":   3,
		"PROCESS_STATE_IMPORTANT_BACKGROUND":   4,
		"PROCESS_STATE_BACKUP":                 5,
		"PROCESS_STATE_SERVICE":                6,
		"PROCESS_STATE_SERVICE_RESTARTING":     7,
		"PROCESS_STATE_RECEIVER":               8,
		"PROCESS_STATE_HEAVY_WEIGHT":           9,
		"PROCESS_STATE_HOME":                   10,
		"PROCESS_STATE_LAST_ACTIVITY":          11,
		"PROCESS_STATE_CACHED_ACTIVITY":        12,
		"PROCESS_STATE_CACHED_ACTIVITY_CLIENT": 13,
		"PROCESS_STATE_CACHED_EMPTY":           14,
		"PROCESS_STATE_FROZEN":                 17,
		"PROCESS_STATE_CACHED":                 18,
		"PROCESS_STATE_BOUND_TOP":              19,
		"PROCESS_STATE_BOUND_FGS":              20,
	}
)

func (x ProcessState) Enum() *ProcessState {
	p := new(ProcessState)
	*p = x
	return p
}

func (x ProcessState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcessState) Descriptor() protoreflect.EnumDescriptor {
	return file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_enumTypes[2].Descriptor()
}

func (ProcessState) Type() protoreflect.EnumType {
	return &file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_enumTypes[2]
}

func (x ProcessState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ProcessState) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ProcessState(num)
	return nil
}

// Deprecated: Use ProcessState.Descriptor instead.
func (ProcessState) EnumDescriptor() ([]byte, []int) {
	return file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_rawDescGZIP(), []int{2}
}

type ServiceOperationState int32

const (
	ServiceOperationState_SERVICE_OPERATION_STATE_UNKNOWN    ServiceOperationState = 0
	ServiceOperationState_SERVICE_OPERATION_STATE_RUNNING    ServiceOperationState = 1
	ServiceOperationState_SERVICE_OPERATION_STATE_STARTED    ServiceOperationState = 2
	ServiceOperationState_SERVICE_OPERATION_STATE_FOREGROUND ServiceOperationState = 3
	ServiceOperationState_SERVICE_OPERATION_STATE_BOUND      ServiceOperationState = 4
	ServiceOperationState_SERVICE_OPERATION_STATE_EXECUTING  ServiceOperationState = 5
)

// Enum value maps for ServiceOperationState.
var (
	ServiceOperationState_name = map[int32]string{
		0: "SERVICE_OPERATION_STATE_UNKNOWN",
		1: "SERVICE_OPERATION_STATE_RUNNING",
		2: "SERVICE_OPERATION_STATE_STARTED",
		3: "SERVICE_OPERATION_STATE_FOREGROUND",
		4: "SERVICE_OPERATION_STATE_BOUND",
		5: "SERVICE_OPERATION_STATE_EXECUTING",
	}
	ServiceOperationState_value = map[string]int32{
		"SERVICE_OPERATION_STATE_UNKNOWN":    0,
		"SERVICE_OPERATION_STATE_RUNNING":    1,
		"SERVICE_OPERATION_STATE_STARTED":    2,
		"SERVICE_OPERATION_STATE_FOREGROUND": 3,
		"SERVICE_OPERATION_STATE_BOUND":      4,
		"SERVICE_OPERATION_STATE_EXECUTING":  5,
	}
)

func (x ServiceOperationState) Enum() *ServiceOperationState {
	p := new(ServiceOperationState)
	*p = x
	return p
}

func (x ServiceOperationState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceOperationState) Descriptor() protoreflect.EnumDescriptor {
	return file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_enumTypes[3].Descriptor()
}

func (ServiceOperationState) Type() protoreflect.EnumType {
	return &file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_enumTypes[3]
}

func (x ServiceOperationState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ServiceOperationState) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ServiceOperationState(num)
	return nil
}

// Deprecated: Use ServiceOperationState.Descriptor instead.
func (ServiceOperationState) EnumDescriptor() ([]byte, []int) {
	return file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_rawDescGZIP(), []int{3}
}

// this enum list is from frameworks/base/core/java/com/android/internal/app/procstats/ProcessStats.java
// and not frameworks/base/core/java/android/app/ActivityManager.java
type AggregatedProcessState int32

const (
	AggregatedProcessState_AGGREGATED_PROCESS_STATE_UNKNOWN AggregatedProcessState = 0
	// Persistent system process; PERSISTENT or PERSISTENT_UI in ActivityManager
	AggregatedProcessState_AGGREGATED_PROCESS_STATE_PERSISTENT AggregatedProcessState = 1
	// Top activity; actually any visible activity; TOP or TOP_SLEEPING in ActivityManager
	AggregatedProcessState_AGGREGATED_PROCESS_STATE_TOP AggregatedProcessState = 2
	// Bound top foreground process; BOUND_TOP or BOUND_FOREGROUND_SERVICE in ActivityManager
	AggregatedProcessState_AGGREGATED_PROCESS_STATE_BOUND_TOP_OR_FGS AggregatedProcessState = 3
	// Important foreground process; FOREGROUND_SERVICE in ActivityManager
	AggregatedProcessState_AGGREGATED_PROCESS_STATE_FGS AggregatedProcessState = 4
	// Important foreground process ; IMPORTANT_FOREGROUND in ActivityManager
	AggregatedProcessState_AGGREGATED_PROCESS_STATE_IMPORTANT_FOREGROUND AggregatedProcessState = 5
	// Various background processes; IMPORTANT_BACKGROUND, TRANSIENT_BACKGROUND, BACKUP, SERVICE,
	// HEAVY_WEIGHT in ActivityManager
	AggregatedProcessState_AGGREGATED_PROCESS_STATE_BACKGROUND AggregatedProcessState = 6
	// Process running a receiver; RECEIVER in ActivityManager
	AggregatedProcessState_AGGREGATED_PROCESS_STATE_RECEIVER AggregatedProcessState = 7
	// Various cached processes; HOME, LAST_ACTIVITY, CACHED_ACTIVITY, CACHED_RECENT,
	// CACHED_ACTIVITY_CLIENT, CACHED_EMPTY in ActivityManager
	AggregatedProcessState_AGGREGATED_PROCESS_STATE_CACHED AggregatedProcessState = 8
)

// Enum value maps for AggregatedProcessState.
var (
	AggregatedProcessState_name = map[int32]string{
		0: "AGGREGATED_PROCESS_STATE_UNKNOWN",
		1: "AGGREGATED_PROCESS_STATE_PERSISTENT",
		2: "AGGREGATED_PROCESS_STATE_TOP",
		3: "AGGREGATED_PROCESS_STATE_BOUND_TOP_OR_FGS",
		4: "AGGREGATED_PROCESS_STATE_FGS",
		5: "AGGREGATED_PROCESS_STATE_IMPORTANT_FOREGROUND",
		6: "AGGREGATED_PROCESS_STATE_BACKGROUND",
		7: "AGGREGATED_PROCESS_STATE_RECEIVER",
		8: "AGGREGATED_PROCESS_STATE_CACHED",
	}
	AggregatedProcessState_value = map[string]int32{
		"AGGREGATED_PROCESS_STATE_UNKNOWN":              0,
		"AGGREGATED_PROCESS_STATE_PERSISTENT":           1,
		"AGGREGATED_PROCESS_STATE_TOP":                  2,
		"AGGREGATED_PROCESS_STATE_BOUND_TOP_OR_FGS":     3,
		"AGGREGATED_PROCESS_STATE_FGS":                  4,
		"AGGREGATED_PROCESS_STATE_IMPORTANT_FOREGROUND": 5,
		"AGGREGATED_PROCESS_STATE_BACKGROUND":           6,
		"AGGREGATED_PROCESS_STATE_RECEIVER":             7,
		"AGGREGATED_PROCESS_STATE_CACHED":               8,
	}
)

func (x AggregatedProcessState) Enum() *AggregatedProcessState {
	p := new(AggregatedProcessState)
	*p = x
	return p
}

func (x AggregatedProcessState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AggregatedProcessState) Descriptor() protoreflect.EnumDescriptor {
	return file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_enumTypes[4].Descriptor()
}

func (AggregatedProcessState) Type() protoreflect.EnumType {
	return &file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_enumTypes[4]
}

func (x AggregatedProcessState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *AggregatedProcessState) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = AggregatedProcessState(num)
	return nil
}

// Deprecated: Use AggregatedProcessState.Descriptor instead.
func (AggregatedProcessState) EnumDescriptor() ([]byte, []int) {
	return file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_rawDescGZIP(), []int{4}
}

var File_frameworks_proto_logging_stats_enums_service_procstats_enum_proto protoreflect.FileDescriptor

var file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_rawDesc = []byte{
	0x0a, 0x41, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x63, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x19, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2a, 0x52,
	0x0a, 0x0b, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a,
	0x14, 0x53, 0x43, 0x52, 0x45, 0x45, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x43, 0x52, 0x45, 0x45,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x53, 0x43, 0x52, 0x45, 0x45, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4e,
	0x10, 0x02, 0x2a, 0x8c, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13,
	0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x52,
	0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x52, 0x41, 0x54, 0x45, 0x10, 0x02,
	0x12, 0x14, 0x0a, 0x10, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10,
	0x04, 0x2a, 0x94, 0x05, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1c, 0x0a,
	0x18, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50,
	0x45, 0x52, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x50,
	0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x4f, 0x50,
	0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x54, 0x4f, 0x50, 0x5f, 0x4f, 0x52,
	0x5f, 0x46, 0x47, 0x53, 0x10, 0x0f, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53,
	0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x47, 0x53, 0x10, 0x10, 0x12, 0x26, 0x0a,
	0x22, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49,
	0x4d, 0x50, 0x4f, 0x52, 0x54, 0x41, 0x4e, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x45, 0x47, 0x52, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x26, 0x0a, 0x22, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x41, 0x4e, 0x54,
	0x5f, 0x42, 0x41, 0x43, 0x4b, 0x47, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x04, 0x12, 0x18, 0x0a,
	0x14, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x42,
	0x41, 0x43, 0x4b, 0x55, 0x50, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x43, 0x45,
	0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45,
	0x10, 0x06, 0x12, 0x24, 0x0a, 0x20, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x52, 0x4f, 0x43,
	0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56,
	0x45, 0x52, 0x10, 0x08, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x48, 0x45, 0x41, 0x56, 0x59, 0x5f, 0x57, 0x45, 0x49, 0x47,
	0x48, 0x54, 0x10, 0x09, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x48, 0x4f, 0x4d, 0x45, 0x10, 0x0a, 0x12, 0x1f, 0x0a, 0x1b,
	0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x41,
	0x53, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x10, 0x0b, 0x12, 0x21, 0x0a,
	0x1d, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43,
	0x41, 0x43, 0x48, 0x45, 0x44, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x10, 0x0c,
	0x12, 0x28, 0x0a, 0x24, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x43, 0x41, 0x43, 0x48, 0x45, 0x44, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54,
	0x59, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x0d, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x52,
	0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x43, 0x48,
	0x45, 0x44, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0x0e, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52,
	0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x52, 0x4f, 0x5a,
	0x45, 0x4e, 0x10, 0x11, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x43, 0x48, 0x45, 0x44, 0x10, 0x12, 0x12, 0x1b,
	0x0a, 0x17, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x42, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x54, 0x4f, 0x50, 0x10, 0x13, 0x12, 0x1b, 0x0a, 0x17, 0x50,
	0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x42, 0x4f, 0x55,
	0x4e, 0x44, 0x5f, 0x46, 0x47, 0x53, 0x10, 0x14, 0x2a, 0xf8, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4f, 0x50,
	0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x26, 0x0a, 0x22, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x4f, 0x52,
	0x45, 0x47, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x21, 0x0a, 0x1d, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x04, 0x12, 0x25, 0x0a, 0x21,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x05, 0x2a, 0x82, 0x03, 0x0a, 0x16, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x24,
	0x0a, 0x20, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x4f,
	0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x27, 0x0a, 0x23, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54,
	0x45, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x50, 0x45, 0x52, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x20, 0x0a,
	0x1c, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x43,
	0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x4f, 0x50, 0x10, 0x02, 0x12,
	0x2d, 0x0a, 0x29, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x52,
	0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x42, 0x4f, 0x55, 0x4e,
	0x44, 0x5f, 0x54, 0x4f, 0x50, 0x5f, 0x4f, 0x52, 0x5f, 0x46, 0x47, 0x53, 0x10, 0x03, 0x12, 0x20,
	0x0a, 0x1c, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x4f,
	0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x47, 0x53, 0x10, 0x04,
	0x12, 0x31, 0x0a, 0x2d, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x50,
	0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4d, 0x50,
	0x4f, 0x52, 0x54, 0x41, 0x4e, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x45, 0x47, 0x52, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x05, 0x12, 0x27, 0x0a, 0x23, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x45,
	0x44, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x42, 0x41, 0x43, 0x4b, 0x47, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x06, 0x12, 0x25, 0x0a, 0x21,
	0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45,
	0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45,
	0x52, 0x10, 0x07, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x45,
	0x44, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x43, 0x41, 0x43, 0x48, 0x45, 0x44, 0x10, 0x08, 0x42, 0x15, 0x42, 0x11, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x50, 0x01,
}

var (
	file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_rawDescOnce sync.Once
	file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_rawDescData = file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_rawDesc
)

func file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_rawDescGZIP() []byte {
	file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_rawDescOnce.Do(func() {
		file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_rawDescData)
	})
	return file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_rawDescData
}

var file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_goTypes = []any{
	(ScreenState)(0),            // 0: android.service.procstats.ScreenState
	(MemoryState)(0),            // 1: android.service.procstats.MemoryState
	(ProcessState)(0),           // 2: android.service.procstats.ProcessState
	(ServiceOperationState)(0),  // 3: android.service.procstats.ServiceOperationState
	(AggregatedProcessState)(0), // 4: android.service.procstats.AggregatedProcessState
}
var file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_init() }
func file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_init() {
	if File_frameworks_proto_logging_stats_enums_service_procstats_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_goTypes,
		DependencyIndexes: file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_depIdxs,
		EnumInfos:         file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_enumTypes,
	}.Build()
	File_frameworks_proto_logging_stats_enums_service_procstats_enum_proto = out.File
	file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_rawDesc = nil
	file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_goTypes = nil
	file_frameworks_proto_logging_stats_enums_service_procstats_enum_proto_depIdxs = nil
}
