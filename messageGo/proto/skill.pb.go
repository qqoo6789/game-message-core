// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: skill.proto

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

type EntityCombatState int32

const (
	EntityCombatState_EntityCombatStateUnknown EntityCombatState = 0
	//技能开始
	EntityCombatState_EntityCombatStateBegin EntityCombatState = 1
	//前摇结束
	EntityCombatState_EntityCombatStateEnd EntityCombatState = 2
	//技能打断
	EntityCombatState_EntityCombatStateInterrupted EntityCombatState = 3
	//飞行物打中目标时
	EntityCombatState_EntityCombatStateFlyHit EntityCombatState = 4
)

// Enum value maps for EntityCombatState.
var (
	EntityCombatState_name = map[int32]string{
		0: "EntityCombatStateUnknown",
		1: "EntityCombatStateBegin",
		2: "EntityCombatStateEnd",
		3: "EntityCombatStateInterrupted",
		4: "EntityCombatStateFlyHit",
	}
	EntityCombatState_value = map[string]int32{
		"EntityCombatStateUnknown":     0,
		"EntityCombatStateBegin":       1,
		"EntityCombatStateEnd":         2,
		"EntityCombatStateInterrupted": 3,
		"EntityCombatStateFlyHit":      4,
	}
)

func (x EntityCombatState) Enum() *EntityCombatState {
	p := new(EntityCombatState)
	*p = x
	return p
}

func (x EntityCombatState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityCombatState) Descriptor() protoreflect.EnumDescriptor {
	return file_skill_proto_enumTypes[0].Descriptor()
}

func (EntityCombatState) Type() protoreflect.EnumType {
	return &file_skill_proto_enumTypes[0]
}

func (x EntityCombatState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityCombatState.Descriptor instead.
func (EntityCombatState) EnumDescriptor() ([]byte, []int) {
	return file_skill_proto_rawDescGZIP(), []int{0}
}

type DamageEffectId int32

const (
	DamageEffectId_DamageEffectIdUnknown DamageEffectId = 0
	// 普通攻击效果
	DamageEffectId_DamageEffectId1001 DamageEffectId = 1001
	// 普通攻击效果（采集木头类）
	DamageEffectId_DamageEffectId1002 DamageEffectId = 1002
	// 普通攻击效果（采集石头类）
	DamageEffectId_DamageEffectId1003 DamageEffectId = 1003
	// 普通攻击效果（用手采集类）
	DamageEffectId_DamageEffectId1004 DamageEffectId = 1004
	// 持续伤害效果
	DamageEffectId_DamageEffectId1005 DamageEffectId = 1005
	// 持续加血效果
	DamageEffectId_DamageEffectId1006 DamageEffectId = 1006
	// 一次性加血效果
	DamageEffectId_DamageEffectId1007 DamageEffectId = 1007
	// 减速效果
	DamageEffectId_DamageEffectId1008 DamageEffectId = 1008
	// 击退效果
	DamageEffectId_DamageEffectId1009 DamageEffectId = 1009
)

// Enum value maps for DamageEffectId.
var (
	DamageEffectId_name = map[int32]string{
		0:    "DamageEffectIdUnknown",
		1001: "DamageEffectId1001",
		1002: "DamageEffectId1002",
		1003: "DamageEffectId1003",
		1004: "DamageEffectId1004",
		1005: "DamageEffectId1005",
		1006: "DamageEffectId1006",
		1007: "DamageEffectId1007",
		1008: "DamageEffectId1008",
		1009: "DamageEffectId1009",
	}
	DamageEffectId_value = map[string]int32{
		"DamageEffectIdUnknown": 0,
		"DamageEffectId1001":    1001,
		"DamageEffectId1002":    1002,
		"DamageEffectId1003":    1003,
		"DamageEffectId1004":    1004,
		"DamageEffectId1005":    1005,
		"DamageEffectId1006":    1006,
		"DamageEffectId1007":    1007,
		"DamageEffectId1008":    1008,
		"DamageEffectId1009":    1009,
	}
)

func (x DamageEffectId) Enum() *DamageEffectId {
	p := new(DamageEffectId)
	*p = x
	return p
}

func (x DamageEffectId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DamageEffectId) Descriptor() protoreflect.EnumDescriptor {
	return file_skill_proto_enumTypes[1].Descriptor()
}

func (DamageEffectId) Type() protoreflect.EnumType {
	return &file_skill_proto_enumTypes[1]
}

func (x DamageEffectId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DamageEffectId.Descriptor instead.
func (DamageEffectId) EnumDescriptor() ([]byte, []int) {
	return file_skill_proto_rawDescGZIP(), []int{1}
}

type DamageState int32

const (
	// 普通伤害
	DamageState_DamageStateNormal DamageState = 0
	// 暴击伤害
	DamageState_DamageStateCrit DamageState = 1
	// 未命中
	DamageState_DamageStateMiss DamageState = 3
	// 掉落伤害
	DamageState_DamageStateFall DamageState = 4
)

// Enum value maps for DamageState.
var (
	DamageState_name = map[int32]string{
		0: "DamageStateNormal",
		1: "DamageStateCrit",
		3: "DamageStateMiss",
		4: "DamageStateFall",
	}
	DamageState_value = map[string]int32{
		"DamageStateNormal": 0,
		"DamageStateCrit":   1,
		"DamageStateMiss":   3,
		"DamageStateFall":   4,
	}
)

func (x DamageState) Enum() *DamageState {
	p := new(DamageState)
	*p = x
	return p
}

func (x DamageState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DamageState) Descriptor() protoreflect.EnumDescriptor {
	return file_skill_proto_enumTypes[2].Descriptor()
}

func (DamageState) Type() protoreflect.EnumType {
	return &file_skill_proto_enumTypes[2]
}

func (x DamageState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DamageState.Descriptor instead.
func (DamageState) EnumDescriptor() ([]byte, []int) {
	return file_skill_proto_rawDescGZIP(), []int{2}
}

type DamageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DmgState DamageState `protobuf:"varint,1,opt,name=dmg_state,json=dmgState,proto3,enum=gameMessageCore.DamageState" json:"dmg_state,omitempty"`
	// 当前数值
	CurrentInt int32 `protobuf:"varint,2,opt,name=current_int,json=currentInt,proto3" json:"current_int,omitempty"`
	// 变更数值 delta = current - 原来 (<0为伤害, >0回血)
	DeltaInt int32 `protobuf:"varint,3,opt,name=delta_int,json=deltaInt,proto3" json:"delta_int,omitempty"`
}

func (x *DamageData) Reset() {
	*x = DamageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skill_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DamageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DamageData) ProtoMessage() {}

func (x *DamageData) ProtoReflect() protoreflect.Message {
	mi := &file_skill_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DamageData.ProtoReflect.Descriptor instead.
func (*DamageData) Descriptor() ([]byte, []int) {
	return file_skill_proto_rawDescGZIP(), []int{0}
}

func (x *DamageData) GetDmgState() DamageState {
	if x != nil {
		return x.DmgState
	}
	return DamageState_DamageStateNormal
}

func (x *DamageData) GetCurrentInt() int32 {
	if x != nil {
		return x.CurrentInt
	}
	return 0
}

func (x *DamageData) GetDeltaInt() int32 {
	if x != nil {
		return x.DeltaInt
	}
	return 0
}

type DamageEffectBeatBack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurLoc    *EntityLocation `protobuf:"bytes,1,opt,name=cur_loc,json=curLoc,proto3" json:"cur_loc,omitempty"`
	BackToPos *EntityLocation `protobuf:"bytes,2,opt,name=back_to_pos,json=backToPos,proto3" json:"back_to_pos,omitempty"`
}

func (x *DamageEffectBeatBack) Reset() {
	*x = DamageEffectBeatBack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skill_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DamageEffectBeatBack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DamageEffectBeatBack) ProtoMessage() {}

func (x *DamageEffectBeatBack) ProtoReflect() protoreflect.Message {
	mi := &file_skill_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DamageEffectBeatBack.ProtoReflect.Descriptor instead.
func (*DamageEffectBeatBack) Descriptor() ([]byte, []int) {
	return file_skill_proto_rawDescGZIP(), []int{1}
}

func (x *DamageEffectBeatBack) GetCurLoc() *EntityLocation {
	if x != nil {
		return x.CurLoc
	}
	return nil
}

func (x *DamageEffectBeatBack) GetBackToPos() *EntityLocation {
	if x != nil {
		return x.BackToPos
	}
	return nil
}

type DamageEffect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EffectType DamageEffectId `protobuf:"varint,1,opt,name=effectType,proto3,enum=gameMessageCore.DamageEffectId" json:"effectType,omitempty"`
	// 效果过期时间
	ExpiredAt int32 `protobuf:"varint,2,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	// Types that are assignable to Data:
	//	*DamageEffect_IntValue
	//	*DamageEffect_StrValue
	//	*DamageEffect_BeatBackValue
	//	*DamageEffect_DamageValue
	Data isDamageEffect_Data `protobuf_oneof:"data"`
}

func (x *DamageEffect) Reset() {
	*x = DamageEffect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skill_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DamageEffect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DamageEffect) ProtoMessage() {}

func (x *DamageEffect) ProtoReflect() protoreflect.Message {
	mi := &file_skill_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DamageEffect.ProtoReflect.Descriptor instead.
func (*DamageEffect) Descriptor() ([]byte, []int) {
	return file_skill_proto_rawDescGZIP(), []int{2}
}

func (x *DamageEffect) GetEffectType() DamageEffectId {
	if x != nil {
		return x.EffectType
	}
	return DamageEffectId_DamageEffectIdUnknown
}

func (x *DamageEffect) GetExpiredAt() int32 {
	if x != nil {
		return x.ExpiredAt
	}
	return 0
}

func (m *DamageEffect) GetData() isDamageEffect_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *DamageEffect) GetIntValue() int32 {
	if x, ok := x.GetData().(*DamageEffect_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *DamageEffect) GetStrValue() string {
	if x, ok := x.GetData().(*DamageEffect_StrValue); ok {
		return x.StrValue
	}
	return ""
}

func (x *DamageEffect) GetBeatBackValue() *DamageEffectBeatBack {
	if x, ok := x.GetData().(*DamageEffect_BeatBackValue); ok {
		return x.BeatBackValue
	}
	return nil
}

func (x *DamageEffect) GetDamageValue() *DamageData {
	if x, ok := x.GetData().(*DamageEffect_DamageValue); ok {
		return x.DamageValue
	}
	return nil
}

type isDamageEffect_Data interface {
	isDamageEffect_Data()
}

type DamageEffect_IntValue struct {
	IntValue int32 `protobuf:"varint,10,opt,name=int_value,json=intValue,proto3,oneof"`
}

type DamageEffect_StrValue struct {
	StrValue string `protobuf:"bytes,11,opt,name=str_value,json=strValue,proto3,oneof"`
}

type DamageEffect_BeatBackValue struct {
	BeatBackValue *DamageEffectBeatBack `protobuf:"bytes,12,opt,name=beat_back_value,json=beatBackValue,proto3,oneof"`
}

type DamageEffect_DamageValue struct {
	DamageValue *DamageData `protobuf:"bytes,13,opt,name=damage_value,json=damageValue,proto3,oneof"`
}

func (*DamageEffect_IntValue) isDamageEffect_Data() {}

func (*DamageEffect_StrValue) isDamageEffect_Data() {}

func (*DamageEffect_BeatBackValue) isDamageEffect_Data() {}

func (*DamageEffect_DamageValue) isDamageEffect_Data() {}

type EntityDamage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity  *EntityId       `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	SkillId int32           `protobuf:"varint,2,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
	Effect  []*DamageEffect `protobuf:"bytes,3,rep,name=effect,proto3" json:"effect,omitempty"`
}

func (x *EntityDamage) Reset() {
	*x = EntityDamage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skill_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityDamage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityDamage) ProtoMessage() {}

func (x *EntityDamage) ProtoReflect() protoreflect.Message {
	mi := &file_skill_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityDamage.ProtoReflect.Descriptor instead.
func (*EntityDamage) Descriptor() ([]byte, []int) {
	return file_skill_proto_rawDescGZIP(), []int{3}
}

func (x *EntityDamage) GetEntity() *EntityId {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *EntityDamage) GetSkillId() int32 {
	if x != nil {
		return x.SkillId
	}
	return 0
}

func (x *EntityDamage) GetEffect() []*DamageEffect {
	if x != nil {
		return x.Effect
	}
	return nil
}

//技能飞行物数据
type SkillFlyerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target    *EntityId `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	DestStamp int64     `protobuf:"varint,2,opt,name=dest_stamp,json=destStamp,proto3" json:"dest_stamp,omitempty"`
	Dir       *Vector3  `protobuf:"bytes,3,opt,name=dir,proto3" json:"dir,omitempty"`
}

func (x *SkillFlyerData) Reset() {
	*x = SkillFlyerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skill_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkillFlyerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkillFlyerData) ProtoMessage() {}

func (x *SkillFlyerData) ProtoReflect() protoreflect.Message {
	mi := &file_skill_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkillFlyerData.ProtoReflect.Descriptor instead.
func (*SkillFlyerData) Descriptor() ([]byte, []int) {
	return file_skill_proto_rawDescGZIP(), []int{4}
}

func (x *SkillFlyerData) GetTarget() *EntityId {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *SkillFlyerData) GetDestStamp() int64 {
	if x != nil {
		return x.DestStamp
	}
	return 0
}

func (x *SkillFlyerData) GetDir() *Vector3 {
	if x != nil {
		return x.Dir
	}
	return nil
}

var File_skill_proto protoreflect.FileDescriptor

var file_skill_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67,
	0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x1a, 0x0c,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x76, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x0a, 0x44,
	0x61, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x09, 0x64, 0x6d, 0x67,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x44,
	0x61, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x08, 0x64, 0x6d, 0x67, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x5f, 0x69,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x49,
	0x6e, 0x74, 0x22, 0x91, 0x01, 0x0a, 0x14, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x42, 0x65, 0x61, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x12, 0x38, 0x0a, 0x07, 0x63,
	0x75, 0x72, 0x5f, 0x6c, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x63,
	0x75, 0x72, 0x4c, 0x6f, 0x63, 0x12, 0x3f, 0x0a, 0x0b, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x6f,
	0x5f, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x62, 0x61, 0x63,
	0x6b, 0x54, 0x6f, 0x50, 0x6f, 0x73, 0x22, 0xc7, 0x02, 0x0a, 0x0c, 0x44, 0x61, 0x6d, 0x61, 0x67,
	0x65, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x3f, 0x0a, 0x0a, 0x65, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61,
	0x6d, 0x61, 0x67, 0x65, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x49, 0x64, 0x52, 0x0a, 0x65, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x73, 0x74, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4f, 0x0a, 0x0f, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65,
	0x2e, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x42, 0x65, 0x61,
	0x74, 0x42, 0x61, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x0d, 0x62, 0x65, 0x61, 0x74, 0x42, 0x61, 0x63,
	0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x44,
	0x61, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x61, 0x6d,
	0x61, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x93, 0x01, 0x0a, 0x0c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x61, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x31, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x52, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x12,
	0x35, 0x0a, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72,
	0x65, 0x2e, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x06,
	0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x0e, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x46, 0x6c, 0x79, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x64, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x64, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2a, 0x0a, 0x03, 0x64,
	0x69, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x33, 0x52, 0x03, 0x64, 0x69, 0x72, 0x2a, 0xa6, 0x01, 0x0a, 0x11, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a,
	0x18, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x42, 0x65, 0x67, 0x69, 0x6e, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x10,
	0x02, 0x12, 0x20, 0x0a, 0x1c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6d, 0x62, 0x61,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74, 0x65,
	0x64, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6d,
	0x62, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x79, 0x48, 0x69, 0x74, 0x10, 0x04,
	0x2a, 0x8c, 0x02, 0x0a, 0x0e, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x17,
	0x0a, 0x12, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x31, 0x30, 0x30, 0x31, 0x10, 0xe9, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x44, 0x61, 0x6d, 0x61, 0x67,
	0x65, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x49, 0x64, 0x31, 0x30, 0x30, 0x32, 0x10, 0xea, 0x07,
	0x12, 0x17, 0x0a, 0x12, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x31, 0x30, 0x30, 0x33, 0x10, 0xeb, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x44, 0x61, 0x6d,
	0x61, 0x67, 0x65, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x49, 0x64, 0x31, 0x30, 0x30, 0x34, 0x10,
	0xec, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x31, 0x30, 0x30, 0x35, 0x10, 0xed, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x44,
	0x61, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x49, 0x64, 0x31, 0x30, 0x30,
	0x36, 0x10, 0xee, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x49, 0x64, 0x31, 0x30, 0x30, 0x37, 0x10, 0xef, 0x07, 0x12, 0x17, 0x0a,
	0x12, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x49, 0x64, 0x31,
	0x30, 0x30, 0x38, 0x10, 0xf0, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65,
	0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x49, 0x64, 0x31, 0x30, 0x30, 0x39, 0x10, 0xf1, 0x07, 0x2a,
	0x63, 0x0a, 0x0b, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15,
	0x0a, 0x11, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x72,
	0x6d, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x43, 0x72, 0x69, 0x74, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x61,
	0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x10, 0x03, 0x12,
	0x13, 0x0a, 0x0f, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x46, 0x61,
	0x6c, 0x6c, 0x10, 0x04, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x47, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_skill_proto_rawDescOnce sync.Once
	file_skill_proto_rawDescData = file_skill_proto_rawDesc
)

func file_skill_proto_rawDescGZIP() []byte {
	file_skill_proto_rawDescOnce.Do(func() {
		file_skill_proto_rawDescData = protoimpl.X.CompressGZIP(file_skill_proto_rawDescData)
	})
	return file_skill_proto_rawDescData
}

var file_skill_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_skill_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_skill_proto_goTypes = []interface{}{
	(EntityCombatState)(0),       // 0: gameMessageCore.EntityCombatState
	(DamageEffectId)(0),          // 1: gameMessageCore.DamageEffectId
	(DamageState)(0),             // 2: gameMessageCore.DamageState
	(*DamageData)(nil),           // 3: gameMessageCore.DamageData
	(*DamageEffectBeatBack)(nil), // 4: gameMessageCore.DamageEffectBeatBack
	(*DamageEffect)(nil),         // 5: gameMessageCore.DamageEffect
	(*EntityDamage)(nil),         // 6: gameMessageCore.EntityDamage
	(*SkillFlyerData)(nil),       // 7: gameMessageCore.SkillFlyerData
	(*EntityLocation)(nil),       // 8: gameMessageCore.EntityLocation
	(*EntityId)(nil),             // 9: gameMessageCore.EntityId
	(*Vector3)(nil),              // 10: gameMessageCore.Vector3
}
var file_skill_proto_depIdxs = []int32{
	2,  // 0: gameMessageCore.DamageData.dmg_state:type_name -> gameMessageCore.DamageState
	8,  // 1: gameMessageCore.DamageEffectBeatBack.cur_loc:type_name -> gameMessageCore.EntityLocation
	8,  // 2: gameMessageCore.DamageEffectBeatBack.back_to_pos:type_name -> gameMessageCore.EntityLocation
	1,  // 3: gameMessageCore.DamageEffect.effectType:type_name -> gameMessageCore.DamageEffectId
	4,  // 4: gameMessageCore.DamageEffect.beat_back_value:type_name -> gameMessageCore.DamageEffectBeatBack
	3,  // 5: gameMessageCore.DamageEffect.damage_value:type_name -> gameMessageCore.DamageData
	9,  // 6: gameMessageCore.EntityDamage.entity:type_name -> gameMessageCore.EntityId
	5,  // 7: gameMessageCore.EntityDamage.effect:type_name -> gameMessageCore.DamageEffect
	9,  // 8: gameMessageCore.SkillFlyerData.target:type_name -> gameMessageCore.EntityId
	10, // 9: gameMessageCore.SkillFlyerData.dir:type_name -> gameMessageCore.Vector3
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_skill_proto_init() }
func file_skill_proto_init() {
	if File_skill_proto != nil {
		return
	}
	file_entity_proto_init()
	file_vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_skill_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DamageData); i {
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
		file_skill_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DamageEffectBeatBack); i {
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
		file_skill_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DamageEffect); i {
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
		file_skill_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityDamage); i {
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
		file_skill_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkillFlyerData); i {
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
	file_skill_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*DamageEffect_IntValue)(nil),
		(*DamageEffect_StrValue)(nil),
		(*DamageEffect_BeatBackValue)(nil),
		(*DamageEffect_DamageValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_skill_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_skill_proto_goTypes,
		DependencyIndexes: file_skill_proto_depIdxs,
		EnumInfos:         file_skill_proto_enumTypes,
		MessageInfos:      file_skill_proto_msgTypes,
	}.Build()
	File_skill_proto = out.File
	file_skill_proto_rawDesc = nil
	file_skill_proto_goTypes = nil
	file_skill_proto_depIdxs = nil
}
