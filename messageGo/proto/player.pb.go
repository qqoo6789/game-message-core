// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: player.proto

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

type PlayerRespawnPoint int32

const (
	PlayerRespawnPoint_PlayerRespawnPointLcoal PlayerRespawnPoint = 0
	PlayerRespawnPoint_PlayerRespawnPointCity  PlayerRespawnPoint = 1
	PlayerRespawnPoint_PlayerRespawnPointArea  PlayerRespawnPoint = 2
)

// Enum value maps for PlayerRespawnPoint.
var (
	PlayerRespawnPoint_name = map[int32]string{
		0: "PlayerRespawnPointLcoal",
		1: "PlayerRespawnPointCity",
		2: "PlayerRespawnPointArea",
	}
	PlayerRespawnPoint_value = map[string]int32{
		"PlayerRespawnPointLcoal": 0,
		"PlayerRespawnPointCity":  1,
		"PlayerRespawnPointArea":  2,
	}
)

func (x PlayerRespawnPoint) Enum() *PlayerRespawnPoint {
	p := new(PlayerRespawnPoint)
	*p = x
	return p
}

func (x PlayerRespawnPoint) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerRespawnPoint) Descriptor() protoreflect.EnumDescriptor {
	return file_player_proto_enumTypes[0].Descriptor()
}

func (PlayerRespawnPoint) Type() protoreflect.EnumType {
	return &file_player_proto_enumTypes[0]
}

func (x PlayerRespawnPoint) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerRespawnPoint.Descriptor instead.
func (PlayerRespawnPoint) EnumDescriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{0}
}

type TickOutType int32

const (
	TickOutType_TickOutTypeUnknown TickOutType = 0
	// 服务关闭
	TickOutType_ServiceClose TickOutType = 1
	// 重复登录
	TickOutType_RepeatSingIn TickOutType = 2
	// 非法攻击速度
	TickOutType_IllegalAtkSpd TickOutType = 3
)

// Enum value maps for TickOutType.
var (
	TickOutType_name = map[int32]string{
		0: "TickOutTypeUnknown",
		1: "ServiceClose",
		2: "RepeatSingIn",
		3: "IllegalAtkSpd",
	}
	TickOutType_value = map[string]int32{
		"TickOutTypeUnknown": 0,
		"ServiceClose":       1,
		"RepeatSingIn":       2,
		"IllegalAtkSpd":      3,
	}
)

func (x TickOutType) Enum() *TickOutType {
	p := new(TickOutType)
	*p = x
	return p
}

func (x TickOutType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TickOutType) Descriptor() protoreflect.EnumDescriptor {
	return file_player_proto_enumTypes[1].Descriptor()
}

func (TickOutType) Type() protoreflect.EnumType {
	return &file_player_proto_enumTypes[1]
}

func (x TickOutType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TickOutType.Descriptor instead.
func (TickOutType) EnumDescriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{1}
}

// 容貌数据
type PlayerFeature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eyebrow int32 `protobuf:"varint,1,opt,name=eyebrow,proto3" json:"eyebrow,omitempty"` // 眉毛
	Mouth   int32 `protobuf:"varint,2,opt,name=mouth,proto3" json:"mouth,omitempty"`     // 嘴型
	Eye     int32 `protobuf:"varint,3,opt,name=eye,proto3" json:"eye,omitempty"`         // 眼镜
	Face    int32 `protobuf:"varint,4,opt,name=face,proto3" json:"face,omitempty"`       // 脸型
	Hair    int32 `protobuf:"varint,5,opt,name=hair,proto3" json:"hair,omitempty"`       // 头发
	Glove   int32 `protobuf:"varint,6,opt,name=glove,proto3" json:"glove,omitempty"`     // 手套
	Clothes int32 `protobuf:"varint,7,opt,name=clothes,proto3" json:"clothes,omitempty"` // 衣服/上衣
	Pants   int32 `protobuf:"varint,8,opt,name=pants,proto3" json:"pants,omitempty"`     // 裤子
}

func (x *PlayerFeature) Reset() {
	*x = PlayerFeature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerFeature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerFeature) ProtoMessage() {}

func (x *PlayerFeature) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerFeature.ProtoReflect.Descriptor instead.
func (*PlayerFeature) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerFeature) GetEyebrow() int32 {
	if x != nil {
		return x.Eyebrow
	}
	return 0
}

func (x *PlayerFeature) GetMouth() int32 {
	if x != nil {
		return x.Mouth
	}
	return 0
}

func (x *PlayerFeature) GetEye() int32 {
	if x != nil {
		return x.Eye
	}
	return 0
}

func (x *PlayerFeature) GetFace() int32 {
	if x != nil {
		return x.Face
	}
	return 0
}

func (x *PlayerFeature) GetHair() int32 {
	if x != nil {
		return x.Hair
	}
	return 0
}

func (x *PlayerFeature) GetGlove() int32 {
	if x != nil {
		return x.Glove
	}
	return 0
}

func (x *PlayerFeature) GetClothes() int32 {
	if x != nil {
		return x.Clothes
	}
	return 0
}

func (x *PlayerFeature) GetPants() int32 {
	if x != nil {
		return x.Pants
	}
	return 0
}

type PlayerBaseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RoleId int32  `protobuf:"varint,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	// 玩家性别
	// other=未知性别
	// male=男性
	// female=女性
	Gender   string         `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	RoleIcon string         `protobuf:"bytes,6,opt,name=role_icon,json=roleIcon,proto3" json:"role_icon,omitempty"`
	Feature  *PlayerFeature `protobuf:"bytes,7,opt,name=feature,proto3" json:"feature,omitempty"`
	// 引导记录字段
	Guide bool `protobuf:"varint,8,opt,name=guide,proto3" json:"guide,omitempty"`
}

func (x *PlayerBaseData) Reset() {
	*x = PlayerBaseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerBaseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerBaseData) ProtoMessage() {}

func (x *PlayerBaseData) ProtoReflect() protoreflect.Message {
	mi := &file_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerBaseData.ProtoReflect.Descriptor instead.
func (*PlayerBaseData) Descriptor() ([]byte, []int) {
	return file_player_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerBaseData) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PlayerBaseData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlayerBaseData) GetRoleId() int32 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *PlayerBaseData) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *PlayerBaseData) GetRoleIcon() string {
	if x != nil {
		return x.RoleIcon
	}
	return ""
}

func (x *PlayerBaseData) GetFeature() *PlayerFeature {
	if x != nil {
		return x.Feature
	}
	return nil
}

func (x *PlayerBaseData) GetGuide() bool {
	if x != nil {
		return x.Guide
	}
	return false
}

var File_player_proto protoreflect.FileDescriptor

var file_player_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x22,
	0xbf, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x79, 0x65, 0x62, 0x72, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x65, 0x79, 0x65, 0x62, 0x72, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x6f, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x75, 0x74,
	0x68, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x79, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x65, 0x79, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x66, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x69, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x61, 0x69, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x67,
	0x6c, 0x6f, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x67, 0x6c, 0x6f, 0x76,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x6f, 0x74, 0x68, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x74, 0x68, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x61, 0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x6e, 0x74,
	0x73, 0x22, 0xdb, 0x01, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x63, 0x6f, 0x6e, 0x12,
	0x38, 0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x72, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x75, 0x69,
	0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2a,
	0x69, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x61, 0x77, 0x6e,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x63, 0x6f, 0x61, 0x6c,
	0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x61, 0x77, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x69, 0x74, 0x79, 0x10, 0x01, 0x12, 0x1a,
	0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x41, 0x72, 0x65, 0x61, 0x10, 0x02, 0x2a, 0x5c, 0x0a, 0x0b, 0x54, 0x69,
	0x63, 0x6b, 0x4f, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x69, 0x63,
	0x6b, 0x4f, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10,
	0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x53, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x6c, 0x6c, 0x65, 0x67, 0x61, 0x6c,
	0x41, 0x74, 0x6b, 0x53, 0x70, 0x64, 0x10, 0x03, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_player_proto_rawDescOnce sync.Once
	file_player_proto_rawDescData = file_player_proto_rawDesc
)

func file_player_proto_rawDescGZIP() []byte {
	file_player_proto_rawDescOnce.Do(func() {
		file_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_player_proto_rawDescData)
	})
	return file_player_proto_rawDescData
}

var file_player_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_player_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_player_proto_goTypes = []interface{}{
	(PlayerRespawnPoint)(0), // 0: gameMessageCore.PlayerRespawnPoint
	(TickOutType)(0),        // 1: gameMessageCore.TickOutType
	(*PlayerFeature)(nil),   // 2: gameMessageCore.PlayerFeature
	(*PlayerBaseData)(nil),  // 3: gameMessageCore.PlayerBaseData
}
var file_player_proto_depIdxs = []int32{
	2, // 0: gameMessageCore.PlayerBaseData.feature:type_name -> gameMessageCore.PlayerFeature
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_player_proto_init() }
func file_player_proto_init() {
	if File_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerFeature); i {
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
		file_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerBaseData); i {
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
			RawDescriptor: file_player_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_player_proto_goTypes,
		DependencyIndexes: file_player_proto_depIdxs,
		EnumInfos:         file_player_proto_enumTypes,
		MessageInfos:      file_player_proto_msgTypes,
	}.Build()
	File_player_proto = out.File
	file_player_proto_rawDesc = nil
	file_player_proto_goTypes = nil
	file_player_proto_depIdxs = nil
}
