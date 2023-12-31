// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: service.proto

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

type ServiceType int32

const (
	ServiceType_ServiceTypeUnknown ServiceType = 0
	// 主服务
	ServiceType_ServiceTypeMain ServiceType = 1
	// 账号服务
	ServiceType_ServiceTypeAccount ServiceType = 2
	// 场景(战斗)服务
	ServiceType_ServiceTypeScene ServiceType = 3
	// 任务服务
	ServiceType_ServiceTypeTask ServiceType = 4
	// 聊天服务
	ServiceType_ServiceTypeChat ServiceType = 5
	// 网关服务
	ServiceType_ServiceTypeAgent ServiceType = 6
	// 服务管理
	ServiceType_ServiceTypeManager ServiceType = 7
	// 后台控制服务
	ServiceType_ServiceTypeController ServiceType = 8
	// 日志服务
	ServiceType_ServiceTypeLog ServiceType = 9
	// 服务类型上限标记必须放在最后
	ServiceType_ServiceTypeLimit ServiceType = 10
)

// Enum value maps for ServiceType.
var (
	ServiceType_name = map[int32]string{
		0:  "ServiceTypeUnknown",
		1:  "ServiceTypeMain",
		2:  "ServiceTypeAccount",
		3:  "ServiceTypeScene",
		4:  "ServiceTypeTask",
		5:  "ServiceTypeChat",
		6:  "ServiceTypeAgent",
		7:  "ServiceTypeManager",
		8:  "ServiceTypeController",
		9:  "ServiceTypeLog",
		10: "ServiceTypeLimit",
	}
	ServiceType_value = map[string]int32{
		"ServiceTypeUnknown":    0,
		"ServiceTypeMain":       1,
		"ServiceTypeAccount":    2,
		"ServiceTypeScene":      3,
		"ServiceTypeTask":       4,
		"ServiceTypeChat":       5,
		"ServiceTypeAgent":      6,
		"ServiceTypeManager":    7,
		"ServiceTypeController": 8,
		"ServiceTypeLog":        9,
		"ServiceTypeLimit":      10,
	}
)

func (x ServiceType) Enum() *ServiceType {
	p := new(ServiceType)
	*p = x
	return p
}

func (x ServiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_service_proto_enumTypes[0].Descriptor()
}

func (ServiceType) Type() protoreflect.EnumType {
	return &file_service_proto_enumTypes[0]
}

func (x ServiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceType.Descriptor instead.
func (ServiceType) EnumDescriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

type SceneServiceSubType int32

const (
	SceneServiceSubType_UnknownSubType SceneServiceSubType = 0
	SceneServiceSubType_World          SceneServiceSubType = 1
	SceneServiceSubType_Home           SceneServiceSubType = 2
	SceneServiceSubType_Dungeon        SceneServiceSubType = 3
)

// Enum value maps for SceneServiceSubType.
var (
	SceneServiceSubType_name = map[int32]string{
		0: "UnknownSubType",
		1: "World",
		2: "Home",
		3: "Dungeon",
	}
	SceneServiceSubType_value = map[string]int32{
		"UnknownSubType": 0,
		"World":          1,
		"Home":           2,
		"Dungeon":        3,
	}
)

func (x SceneServiceSubType) Enum() *SceneServiceSubType {
	p := new(SceneServiceSubType)
	*p = x
	return p
}

func (x SceneServiceSubType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SceneServiceSubType) Descriptor() protoreflect.EnumDescriptor {
	return file_service_proto_enumTypes[1].Descriptor()
}

func (SceneServiceSubType) Type() protoreflect.EnumType {
	return &file_service_proto_enumTypes[1]
}

func (x SceneServiceSubType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SceneServiceSubType.Descriptor instead.
func (SceneServiceSubType) EnumDescriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65,
	0x2a, 0x85, 0x02, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x10, 0x01, 0x12, 0x16, 0x0a,
	0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x10, 0x04,
	0x12, 0x13, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43,
	0x68, 0x61, 0x74, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x10, 0x07, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x10, 0x08, 0x12, 0x12,
	0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x6f, 0x67,
	0x10, 0x09, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x10, 0x0a, 0x2a, 0x4b, 0x0a, 0x13, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x0e, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70,
	0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x48, 0x6f, 0x6d, 0x65, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x10, 0x03, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x47, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_service_proto_goTypes = []interface{}{
	(ServiceType)(0),         // 0: gameMessageCore.ServiceType
	(SceneServiceSubType)(0), // 1: gameMessageCore.SceneServiceSubType
}
var file_service_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		EnumInfos:         file_service_proto_enumTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
