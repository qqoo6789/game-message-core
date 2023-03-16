// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: capture.proto

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

// 捕获状态
type CaptureStatus int32

const (
	CaptureStatus_CaptureStatusUnknown CaptureStatus = 0
	// 进入捕获
	CaptureStatus_CaptureStatusStart CaptureStatus = 1
	// 更新捕获信息
	CaptureStatus_CaptureStatusUpdate CaptureStatus = 2
	// 成功结束
	CaptureStatus_CaptureStatusEndSuccess CaptureStatus = 3
	// 失败结束
	CaptureStatus_CaptureStatusEndFail CaptureStatus = 4
)

// Enum value maps for CaptureStatus.
var (
	CaptureStatus_name = map[int32]string{
		0: "CaptureStatusUnknown",
		1: "CaptureStatusStart",
		2: "CaptureStatusUpdate",
		3: "CaptureStatusEndSuccess",
		4: "CaptureStatusEndFail",
	}
	CaptureStatus_value = map[string]int32{
		"CaptureStatusUnknown":    0,
		"CaptureStatusStart":      1,
		"CaptureStatusUpdate":     2,
		"CaptureStatusEndSuccess": 3,
		"CaptureStatusEndFail":    4,
	}
)

func (x CaptureStatus) Enum() *CaptureStatus {
	p := new(CaptureStatus)
	*p = x
	return p
}

func (x CaptureStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CaptureStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_capture_proto_enumTypes[0].Descriptor()
}

func (CaptureStatus) Type() protoreflect.EnumType {
	return &file_capture_proto_enumTypes[0]
}

func (x CaptureStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CaptureStatus.Descriptor instead.
func (CaptureStatus) EnumDescriptor() ([]byte, []int) {
	return file_capture_proto_rawDescGZIP(), []int{0}
}

// 失败原因
type CaptureFailReason int32

const (
	CaptureFailReason_CaptureFailReasonUnknown CaptureFailReason = 0
	// 主动放弃
	CaptureFailReason_CaptureFailReasonGiveup CaptureFailReason = 1
	// 我方意外死亡
	CaptureFailReason_CaptureFailReasonRoleDie CaptureFailReason = 2
	// 怪物意外死亡
	CaptureFailReason_CaptureFailReasonMonsterDie CaptureFailReason = 3
	// 距离太远
	CaptureFailReason_CaptureFailReasonFarDistance CaptureFailReason = 4
	// 绳子磨损断开
	CaptureFailReason_CaptureFailReasonRopeBreak CaptureFailReason = 5
	// 捕获超时
	CaptureFailReason_CaptureFailReasonTimeout CaptureFailReason = 6
	// item数量不足
	CaptureFailReason_CaptureFailReasonLackOfItem CaptureFailReason = 7
)

// Enum value maps for CaptureFailReason.
var (
	CaptureFailReason_name = map[int32]string{
		0: "CaptureFailReasonUnknown",
		1: "CaptureFailReasonGiveup",
		2: "CaptureFailReasonRoleDie",
		3: "CaptureFailReasonMonsterDie",
		4: "CaptureFailReasonFarDistance",
		5: "CaptureFailReasonRopeBreak",
		6: "CaptureFailReasonTimeout",
		7: "CaptureFailReasonLackOfItem",
	}
	CaptureFailReason_value = map[string]int32{
		"CaptureFailReasonUnknown":     0,
		"CaptureFailReasonGiveup":      1,
		"CaptureFailReasonRoleDie":     2,
		"CaptureFailReasonMonsterDie":  3,
		"CaptureFailReasonFarDistance": 4,
		"CaptureFailReasonRopeBreak":   5,
		"CaptureFailReasonTimeout":     6,
		"CaptureFailReasonLackOfItem":  7,
	}
)

func (x CaptureFailReason) Enum() *CaptureFailReason {
	p := new(CaptureFailReason)
	*p = x
	return p
}

func (x CaptureFailReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CaptureFailReason) Descriptor() protoreflect.EnumDescriptor {
	return file_capture_proto_enumTypes[1].Descriptor()
}

func (CaptureFailReason) Type() protoreflect.EnumType {
	return &file_capture_proto_enumTypes[1]
}

func (x CaptureFailReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CaptureFailReason.Descriptor instead.
func (CaptureFailReason) EnumDescriptor() ([]byte, []int) {
	return file_capture_proto_rawDescGZIP(), []int{1}
}

type CaptureData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaptureId    int64         `protobuf:"varint,1,opt,name=capture_id,json=captureId,proto3" json:"capture_id,omitempty"`             //捕获对象
	CaptureAt    int64         `protobuf:"varint,2,opt,name=capture_at,json=captureAt,proto3" json:"capture_at,omitempty"`             // 捕获开始时间
	Progress     int32         `protobuf:"varint,3,opt,name=progress,proto3" json:"progress,omitempty"`                                // 捕获进度
	RopeCurHp    int32         `protobuf:"varint,4,opt,name=rope_cur_hp,json=ropeCurHp,proto3" json:"rope_cur_hp,omitempty"`           // 捕获绳当前hp
	RopeMaxHp    int32         `protobuf:"varint,5,opt,name=rope_max_hp,json=ropeMaxHp,proto3" json:"rope_max_hp,omitempty"`           // 捕获绳最大hp
	CaptureCurHp float32       `protobuf:"fixed32,6,opt,name=capture_cur_hp,json=captureCurHp,proto3" json:"capture_cur_hp,omitempty"` // 捕获当前hp
	CaptureMaxHp float32       `protobuf:"fixed32,7,opt,name=capture_max_hp,json=captureMaxHp,proto3" json:"capture_max_hp,omitempty"` // 捕获最大hp
	Status       CaptureStatus `protobuf:"varint,8,opt,name=status,proto3,enum=gameMessageCore.CaptureStatus" json:"status,omitempty"`
}

func (x *CaptureData) Reset() {
	*x = CaptureData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capture_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaptureData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptureData) ProtoMessage() {}

func (x *CaptureData) ProtoReflect() protoreflect.Message {
	mi := &file_capture_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptureData.ProtoReflect.Descriptor instead.
func (*CaptureData) Descriptor() ([]byte, []int) {
	return file_capture_proto_rawDescGZIP(), []int{0}
}

func (x *CaptureData) GetCaptureId() int64 {
	if x != nil {
		return x.CaptureId
	}
	return 0
}

func (x *CaptureData) GetCaptureAt() int64 {
	if x != nil {
		return x.CaptureAt
	}
	return 0
}

func (x *CaptureData) GetProgress() int32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *CaptureData) GetRopeCurHp() int32 {
	if x != nil {
		return x.RopeCurHp
	}
	return 0
}

func (x *CaptureData) GetRopeMaxHp() int32 {
	if x != nil {
		return x.RopeMaxHp
	}
	return 0
}

func (x *CaptureData) GetCaptureCurHp() float32 {
	if x != nil {
		return x.CaptureCurHp
	}
	return 0
}

func (x *CaptureData) GetCaptureMaxHp() float32 {
	if x != nil {
		return x.CaptureMaxHp
	}
	return 0
}

func (x *CaptureData) GetStatus() CaptureStatus {
	if x != nil {
		return x.Status
	}
	return CaptureStatus_CaptureStatusUnknown
}

var File_capture_proto protoreflect.FileDescriptor

var file_capture_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65,
	0x22, 0xab, 0x02, 0x0a, 0x0b, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x41, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0b, 0x72, 0x6f,
	0x70, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x5f, 0x68, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x72, 0x6f, 0x70, 0x65, 0x43, 0x75, 0x72, 0x48, 0x70, 0x12, 0x1e, 0x0a, 0x0b, 0x72, 0x6f,
	0x70, 0x65, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x68, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x72, 0x6f, 0x70, 0x65, 0x4d, 0x61, 0x78, 0x48, 0x70, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x61,
	0x70, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x5f, 0x68, 0x70, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x43, 0x75, 0x72, 0x48, 0x70,
	0x12, 0x24, 0x0a, 0x0e, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x61, 0x78, 0x5f,
	0x68, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72,
	0x65, 0x4d, 0x61, 0x78, 0x48, 0x70, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x91,
	0x01, 0x0a, 0x0d, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x14, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x61,
	0x70, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x43,
	0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x64, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x61, 0x70, 0x74,
	0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x64, 0x46, 0x61, 0x69, 0x6c,
	0x10, 0x04, 0x2a, 0x8e, 0x02, 0x0a, 0x11, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x46, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x61, 0x70, 0x74,
	0x75, 0x72, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x55, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72,
	0x65, 0x46, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x47, 0x69, 0x76, 0x65, 0x75,
	0x70, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x46, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x69, 0x65, 0x10,
	0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x46, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x44, 0x69, 0x65,
	0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x46, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x46, 0x61, 0x72, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x46,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x6f, 0x70, 0x65, 0x42, 0x72, 0x65,
	0x61, 0x6b, 0x10, 0x05, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x46,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x10, 0x06, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x46, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4c, 0x61, 0x63, 0x6b, 0x4f, 0x66, 0x49, 0x74, 0x65,
	0x6d, 0x10, 0x07, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x47, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_capture_proto_rawDescOnce sync.Once
	file_capture_proto_rawDescData = file_capture_proto_rawDesc
)

func file_capture_proto_rawDescGZIP() []byte {
	file_capture_proto_rawDescOnce.Do(func() {
		file_capture_proto_rawDescData = protoimpl.X.CompressGZIP(file_capture_proto_rawDescData)
	})
	return file_capture_proto_rawDescData
}

var file_capture_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_capture_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_capture_proto_goTypes = []interface{}{
	(CaptureStatus)(0),     // 0: gameMessageCore.CaptureStatus
	(CaptureFailReason)(0), // 1: gameMessageCore.CaptureFailReason
	(*CaptureData)(nil),    // 2: gameMessageCore.CaptureData
}
var file_capture_proto_depIdxs = []int32{
	0, // 0: gameMessageCore.CaptureData.status:type_name -> gameMessageCore.CaptureStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_capture_proto_init() }
func file_capture_proto_init() {
	if File_capture_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_capture_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaptureData); i {
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
			RawDescriptor: file_capture_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_capture_proto_goTypes,
		DependencyIndexes: file_capture_proto_depIdxs,
		EnumInfos:         file_capture_proto_enumTypes,
		MessageInfos:      file_capture_proto_msgTypes,
	}.Build()
	File_capture_proto = out.File
	file_capture_proto_rawDesc = nil
	file_capture_proto_goTypes = nil
	file_capture_proto_depIdxs = nil
}
