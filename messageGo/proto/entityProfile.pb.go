// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: entityProfile.proto

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

type EntityProfileField int32

const (
	EntityProfileField_EntityProfileFieldUnKnown    EntityProfileField = 0
	EntityProfileField_EntityProfileFieldAtt        EntityProfileField = 3
	EntityProfileField_EntityProfileFieldAttSpeed   EntityProfileField = 4
	EntityProfileField_EntityProfileFieldDef        EntityProfileField = 5
	EntityProfileField_EntityProfileFieldHpLimit    EntityProfileField = 6
	EntityProfileField_EntityProfileFieldCritRate   EntityProfileField = 7
	EntityProfileField_EntityProfileFieldCritDamage EntityProfileField = 8
	EntityProfileField_EntityProfileFieldHitRate    EntityProfileField = 9
	EntityProfileField_EntityProfileFieldMissRate   EntityProfileField = 10
	EntityProfileField_EntityProfileFieldMoveSpeed  EntityProfileField = 11
	EntityProfileField_EntityProfileFieldPushDmg    EntityProfileField = 12
	EntityProfileField_EntityProfileFieldPushDist   EntityProfileField = 13
	EntityProfileField_EntityProfileFieldHpCurrent  EntityProfileField = 14
	EntityProfileField_EntityProfileFieldRename     EntityProfileField = 15
	EntityProfileField_EntityProfileFieldHpRecovery EntityProfileField = 16
)

// Enum value maps for EntityProfileField.
var (
	EntityProfileField_name = map[int32]string{
		0:  "EntityProfileFieldUnKnown",
		3:  "EntityProfileFieldAtt",
		4:  "EntityProfileFieldAttSpeed",
		5:  "EntityProfileFieldDef",
		6:  "EntityProfileFieldHpLimit",
		7:  "EntityProfileFieldCritRate",
		8:  "EntityProfileFieldCritDamage",
		9:  "EntityProfileFieldHitRate",
		10: "EntityProfileFieldMissRate",
		11: "EntityProfileFieldMoveSpeed",
		12: "EntityProfileFieldPushDmg",
		13: "EntityProfileFieldPushDist",
		14: "EntityProfileFieldHpCurrent",
		15: "EntityProfileFieldRename",
		16: "EntityProfileFieldHpRecovery",
	}
	EntityProfileField_value = map[string]int32{
		"EntityProfileFieldUnKnown":    0,
		"EntityProfileFieldAtt":        3,
		"EntityProfileFieldAttSpeed":   4,
		"EntityProfileFieldDef":        5,
		"EntityProfileFieldHpLimit":    6,
		"EntityProfileFieldCritRate":   7,
		"EntityProfileFieldCritDamage": 8,
		"EntityProfileFieldHitRate":    9,
		"EntityProfileFieldMissRate":   10,
		"EntityProfileFieldMoveSpeed":  11,
		"EntityProfileFieldPushDmg":    12,
		"EntityProfileFieldPushDist":   13,
		"EntityProfileFieldHpCurrent":  14,
		"EntityProfileFieldRename":     15,
		"EntityProfileFieldHpRecovery": 16,
	}
)

func (x EntityProfileField) Enum() *EntityProfileField {
	p := new(EntityProfileField)
	*p = x
	return p
}

func (x EntityProfileField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityProfileField) Descriptor() protoreflect.EnumDescriptor {
	return file_entityProfile_proto_enumTypes[0].Descriptor()
}

func (EntityProfileField) Type() protoreflect.EnumType {
	return &file_entityProfile_proto_enumTypes[0]
}

func (x EntityProfileField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityProfileField.Descriptor instead.
func (EntityProfileField) EnumDescriptor() ([]byte, []int) {
	return file_entityProfile_proto_rawDescGZIP(), []int{0}
}

type EntityProfileUpdateTriggers int32

const (
	EntityProfileUpdateTriggers_EntityProfileUpdateTriggersUnKnown       EntityProfileUpdateTriggers = 0
	EntityProfileUpdateTriggers_EntityProfileUpdateTriggersBeAttack      EntityProfileUpdateTriggers = 1
	EntityProfileUpdateTriggers_EntityProfileUpdateTriggersUsConsumable  EntityProfileUpdateTriggers = 2
	EntityProfileUpdateTriggers_EntityProfileUpdateTriggersAutoRecoverHp EntityProfileUpdateTriggers = 3
	EntityProfileUpdateTriggers_EntityProfileUpdateTriggersBuff          EntityProfileUpdateTriggers = 4
)

// Enum value maps for EntityProfileUpdateTriggers.
var (
	EntityProfileUpdateTriggers_name = map[int32]string{
		0: "EntityProfileUpdateTriggersUnKnown",
		1: "EntityProfileUpdateTriggersBeAttack",
		2: "EntityProfileUpdateTriggersUsConsumable",
		3: "EntityProfileUpdateTriggersAutoRecoverHp",
		4: "EntityProfileUpdateTriggersBuff",
	}
	EntityProfileUpdateTriggers_value = map[string]int32{
		"EntityProfileUpdateTriggersUnKnown":       0,
		"EntityProfileUpdateTriggersBeAttack":      1,
		"EntityProfileUpdateTriggersUsConsumable":  2,
		"EntityProfileUpdateTriggersAutoRecoverHp": 3,
		"EntityProfileUpdateTriggersBuff":          4,
	}
)

func (x EntityProfileUpdateTriggers) Enum() *EntityProfileUpdateTriggers {
	p := new(EntityProfileUpdateTriggers)
	*p = x
	return p
}

func (x EntityProfileUpdateTriggers) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityProfileUpdateTriggers) Descriptor() protoreflect.EnumDescriptor {
	return file_entityProfile_proto_enumTypes[1].Descriptor()
}

func (EntityProfileUpdateTriggers) Type() protoreflect.EnumType {
	return &file_entityProfile_proto_enumTypes[1]
}

func (x EntityProfileUpdateTriggers) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityProfileUpdateTriggers.Descriptor instead.
func (EntityProfileUpdateTriggers) EnumDescriptor() ([]byte, []int) {
	return file_entityProfile_proto_rawDescGZIP(), []int{1}
}

type EntityProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 等级
	UnKnown int32 `protobuf:"varint,1,opt,name=UnKnown,proto3" json:"UnKnown,omitempty"`
	// 攻击力
	Att int32 `protobuf:"varint,3,opt,name=att,proto3" json:"att,omitempty"`
	// 攻击速率
	AttSpeed int32 `protobuf:"varint,4,opt,name=att_speed,json=attSpeed,proto3" json:"att_speed,omitempty"`
	// 防御力
	Def int32 `protobuf:"varint,5,opt,name=def,proto3" json:"def,omitempty"`
	// 当前血量
	HpCurrent int32 `protobuf:"varint,6,opt,name=hp_current,json=hpCurrent,proto3" json:"hp_current,omitempty"`
	// 血量上限
	HpLimit int32 `protobuf:"varint,7,opt,name=hp_limit,json=hpLimit,proto3" json:"hp_limit,omitempty"`
	// 暴击(Critical Strikes)率
	CritRate int32 `protobuf:"varint,8,opt,name=crit_rate,json=critRate,proto3" json:"crit_rate,omitempty"`
	// 暴击伤害
	CritDmg int32 `protobuf:"varint,9,opt,name=crit_dmg,json=critDmg,proto3" json:"crit_dmg,omitempty"`
	// 名中率
	HitRate int32 `protobuf:"varint,10,opt,name=hit_rate,json=hitRate,proto3" json:"hit_rate,omitempty"`
	// 闪避率
	MissRate int32 `protobuf:"varint,11,opt,name=miss_rate,json=missRate,proto3" json:"miss_rate,omitempty"`
	// 移动速度
	MoveSpeed float32 `protobuf:"fixed32,12,opt,name=move_speed,json=moveSpeed,proto3" json:"move_speed,omitempty"`
	// 击退伤害
	PushDmg int32 `protobuf:"varint,13,opt,name=push_dmg,json=pushDmg,proto3" json:"push_dmg,omitempty"`
	// 击退距离
	PushDist int32 `protobuf:"varint,14,opt,name=push_dist,json=pushDist,proto3" json:"push_dist,omitempty"`
	// HP恢复
	HpRecovery int32 `protobuf:"varint,21,opt,name=hp_recovery,json=hpRecovery,proto3" json:"hp_recovery,omitempty"`
}

func (x *EntityProfile) Reset() {
	*x = EntityProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entityProfile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityProfile) ProtoMessage() {}

func (x *EntityProfile) ProtoReflect() protoreflect.Message {
	mi := &file_entityProfile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityProfile.ProtoReflect.Descriptor instead.
func (*EntityProfile) Descriptor() ([]byte, []int) {
	return file_entityProfile_proto_rawDescGZIP(), []int{0}
}

func (x *EntityProfile) GetUnKnown() int32 {
	if x != nil {
		return x.UnKnown
	}
	return 0
}

func (x *EntityProfile) GetAtt() int32 {
	if x != nil {
		return x.Att
	}
	return 0
}

func (x *EntityProfile) GetAttSpeed() int32 {
	if x != nil {
		return x.AttSpeed
	}
	return 0
}

func (x *EntityProfile) GetDef() int32 {
	if x != nil {
		return x.Def
	}
	return 0
}

func (x *EntityProfile) GetHpCurrent() int32 {
	if x != nil {
		return x.HpCurrent
	}
	return 0
}

func (x *EntityProfile) GetHpLimit() int32 {
	if x != nil {
		return x.HpLimit
	}
	return 0
}

func (x *EntityProfile) GetCritRate() int32 {
	if x != nil {
		return x.CritRate
	}
	return 0
}

func (x *EntityProfile) GetCritDmg() int32 {
	if x != nil {
		return x.CritDmg
	}
	return 0
}

func (x *EntityProfile) GetHitRate() int32 {
	if x != nil {
		return x.HitRate
	}
	return 0
}

func (x *EntityProfile) GetMissRate() int32 {
	if x != nil {
		return x.MissRate
	}
	return 0
}

func (x *EntityProfile) GetMoveSpeed() float32 {
	if x != nil {
		return x.MoveSpeed
	}
	return 0
}

func (x *EntityProfile) GetPushDmg() int32 {
	if x != nil {
		return x.PushDmg
	}
	return 0
}

func (x *EntityProfile) GetPushDist() int32 {
	if x != nil {
		return x.PushDist
	}
	return 0
}

func (x *EntityProfile) GetHpRecovery() int32 {
	if x != nil {
		return x.HpRecovery
	}
	return 0
}

type EntityProfileUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field          EntityProfileField `protobuf:"varint,1,opt,name=field,proto3,enum=gameMessageCore.EntityProfileField" json:"field,omitempty"`
	CurValue       int32              `protobuf:"varint,2,opt,name=cur_value,json=curValue,proto3" json:"cur_value,omitempty"`
	CurValueStr    string             `protobuf:"bytes,3,opt,name=cur_value_str,json=curValueStr,proto3" json:"cur_value_str,omitempty"`
	UseStringValue bool               `protobuf:"varint,4,opt,name=use_string_value,json=useStringValue,proto3" json:"use_string_value,omitempty"`
}

func (x *EntityProfileUpdate) Reset() {
	*x = EntityProfileUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entityProfile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityProfileUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityProfileUpdate) ProtoMessage() {}

func (x *EntityProfileUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_entityProfile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityProfileUpdate.ProtoReflect.Descriptor instead.
func (*EntityProfileUpdate) Descriptor() ([]byte, []int) {
	return file_entityProfile_proto_rawDescGZIP(), []int{1}
}

func (x *EntityProfileUpdate) GetField() EntityProfileField {
	if x != nil {
		return x.Field
	}
	return EntityProfileField_EntityProfileFieldUnKnown
}

func (x *EntityProfileUpdate) GetCurValue() int32 {
	if x != nil {
		return x.CurValue
	}
	return 0
}

func (x *EntityProfileUpdate) GetCurValueStr() string {
	if x != nil {
		return x.CurValueStr
	}
	return ""
}

func (x *EntityProfileUpdate) GetUseStringValue() bool {
	if x != nil {
		return x.UseStringValue
	}
	return false
}

var File_entityProfile_proto protoreflect.FileDescriptor

var file_entityProfile_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x22, 0x8c, 0x03, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x6e, 0x4b, 0x6e,
	0x6f, 0x77, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x55, 0x6e, 0x4b, 0x6e, 0x6f,
	0x77, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x74, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x61, 0x74, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x74, 0x74, 0x53, 0x70, 0x65, 0x65,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x64, 0x65, 0x66, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x70, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x68, 0x70, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x70, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x68, 0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x72, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x63, 0x72, 0x69, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x72,
	0x69, 0x74, 0x5f, 0x64, 0x6d, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x72,
	0x69, 0x74, 0x44, 0x6d, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x68, 0x69, 0x74, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x69, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x75, 0x73, 0x68, 0x5f, 0x64, 0x6d, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x70, 0x75, 0x73, 0x68, 0x44, 0x6d, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x73, 0x68, 0x5f,
	0x64, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68,
	0x44, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x70, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x70, 0x52, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x22, 0xbb, 0x01, 0x0a, 0x13, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x75, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75,
	0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x75, 0x73, 0x65,
	0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x2a, 0xea, 0x03, 0x0a, 0x12, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x55, 0x6e, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x41,
	0x74, 0x74, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x41, 0x74, 0x74, 0x53, 0x70, 0x65,
	0x65, 0x64, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x66, 0x10, 0x05, 0x12,
	0x1d, 0x0a, 0x19, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x10, 0x06, 0x12, 0x1e,
	0x0a, 0x1a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x43, 0x72, 0x69, 0x74, 0x52, 0x61, 0x74, 0x65, 0x10, 0x07, 0x12, 0x20,
	0x0a, 0x1c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x43, 0x72, 0x69, 0x74, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x10, 0x08,
	0x12, 0x1d, 0x0a, 0x19, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x69, 0x74, 0x52, 0x61, 0x74, 0x65, 0x10, 0x09, 0x12,
	0x1e, 0x0a, 0x1a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x69, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x10, 0x0a, 0x12,
	0x1f, 0x0a, 0x1b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x6f, 0x76, 0x65, 0x53, 0x70, 0x65, 0x65, 0x64, 0x10, 0x0b,
	0x12, 0x1d, 0x0a, 0x19, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x75, 0x73, 0x68, 0x44, 0x6d, 0x67, 0x10, 0x0c, 0x12,
	0x1e, 0x0a, 0x1a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x75, 0x73, 0x68, 0x44, 0x69, 0x73, 0x74, 0x10, 0x0d, 0x12,
	0x1f, 0x0a, 0x1b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x70, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x10, 0x0e,
	0x12, 0x1c, 0x0a, 0x18, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x10, 0x0f, 0x12, 0x20,
	0x0a, 0x1c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x48, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x10, 0x10,
	0x2a, 0xee, 0x01, 0x0a, 0x1b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73,
	0x12, 0x26, 0x0a, 0x22, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x55,
	0x6e, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x27, 0x0a, 0x23, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x42, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x10,
	0x01, 0x12, 0x2b, 0x0a, 0x27, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73,
	0x55, 0x73, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x02, 0x12, 0x2c,
	0x0a, 0x28, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x41, 0x75, 0x74,
	0x6f, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x48, 0x70, 0x10, 0x03, 0x12, 0x23, 0x0a, 0x1f,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x42, 0x75, 0x66, 0x66, 0x10,
	0x04, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entityProfile_proto_rawDescOnce sync.Once
	file_entityProfile_proto_rawDescData = file_entityProfile_proto_rawDesc
)

func file_entityProfile_proto_rawDescGZIP() []byte {
	file_entityProfile_proto_rawDescOnce.Do(func() {
		file_entityProfile_proto_rawDescData = protoimpl.X.CompressGZIP(file_entityProfile_proto_rawDescData)
	})
	return file_entityProfile_proto_rawDescData
}

var file_entityProfile_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_entityProfile_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_entityProfile_proto_goTypes = []interface{}{
	(EntityProfileField)(0),          // 0: gameMessageCore.EntityProfileField
	(EntityProfileUpdateTriggers)(0), // 1: gameMessageCore.EntityProfileUpdateTriggers
	(*EntityProfile)(nil),            // 2: gameMessageCore.EntityProfile
	(*EntityProfileUpdate)(nil),      // 3: gameMessageCore.EntityProfileUpdate
}
var file_entityProfile_proto_depIdxs = []int32{
	0, // 0: gameMessageCore.EntityProfileUpdate.field:type_name -> gameMessageCore.EntityProfileField
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_entityProfile_proto_init() }
func file_entityProfile_proto_init() {
	if File_entityProfile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entityProfile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityProfile); i {
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
		file_entityProfile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityProfileUpdate); i {
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
			RawDescriptor: file_entityProfile_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entityProfile_proto_goTypes,
		DependencyIndexes: file_entityProfile_proto_depIdxs,
		EnumInfos:         file_entityProfile_proto_enumTypes,
		MessageInfos:      file_entityProfile_proto_msgTypes,
	}.Build()
	File_entityProfile_proto = out.File
	file_entityProfile_proto_rawDesc = nil
	file_entityProfile_proto_goTypes = nil
	file_entityProfile_proto_depIdxs = nil
}
