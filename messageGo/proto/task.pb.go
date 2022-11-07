// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: task.proto

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

type TaskOptionType int32

const (
	TaskOptionType_UnknownTaskOptionType TaskOptionType = 0
	// 上交若干数量的指定道具(寻物任务)
	TaskOptionType_HandInItem TaskOptionType = 1
	// 使用若干数量的指定道具
	TaskOptionType_UseItem TaskOptionType = 2
	// 获得若干数量的指定道具
	TaskOptionType_PickUpItem TaskOptionType = 3
	// 杀死若干数量的指定怪物
	TaskOptionType_KillMonster TaskOptionType = 4
	//角色达到某等级
	TaskOptionType_UserLevel TaskOptionType = 5
	//指定插槽达到某等级
	TaskOptionType_TargetSlotLevel TaskOptionType = 6
	//指定数量插槽都达到某等级
	TaskOptionType_SlotLevelCount TaskOptionType = 7
	//指定合成技能达到某等级
	TaskOptionType_CraftSkillLevel TaskOptionType = 8
	//使用指定的配方合成
	TaskOptionType_RecipeUse TaskOptionType = 9
	//完成若干数量的指定类型任务
	TaskOptionType_TaskListTypeCount TaskOptionType = 10
	//到达指定坐标点指定半径范围内的区域
	TaskOptionType_TargetPosition TaskOptionType = 11
)

// Enum value maps for TaskOptionType.
var (
	TaskOptionType_name = map[int32]string{
		0:  "UnknownTaskOptionType",
		1:  "HandInItem",
		2:  "UseItem",
		3:  "PickUpItem",
		4:  "KillMonster",
		5:  "UserLevel",
		6:  "TargetSlotLevel",
		7:  "SlotLevelCount",
		8:  "CraftSkillLevel",
		9:  "RecipeUse",
		10: "TaskListTypeCount",
		11: "TargetPosition",
	}
	TaskOptionType_value = map[string]int32{
		"UnknownTaskOptionType": 0,
		"HandInItem":            1,
		"UseItem":               2,
		"PickUpItem":            3,
		"KillMonster":           4,
		"UserLevel":             5,
		"TargetSlotLevel":       6,
		"SlotLevelCount":        7,
		"CraftSkillLevel":       8,
		"RecipeUse":             9,
		"TaskListTypeCount":     10,
		"TargetPosition":        11,
	}
)

func (x TaskOptionType) Enum() *TaskOptionType {
	p := new(TaskOptionType)
	*p = x
	return p
}

func (x TaskOptionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskOptionType) Descriptor() protoreflect.EnumDescriptor {
	return file_task_proto_enumTypes[0].Descriptor()
}

func (TaskOptionType) Type() protoreflect.EnumType {
	return &file_task_proto_enumTypes[0]
}

func (x TaskOptionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskOptionType.Descriptor instead.
func (TaskOptionType) EnumDescriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{0}
}

type TaskListType int32

const (
	TaskListType_TaskListTypeUnknown TaskListType = 0
	// 每日任务
	TaskListType_TaskListTypeDaily TaskListType = 1
	// 悬赏任务
	TaskListType_TaskListTypeRewarded TaskListType = 2
	// 主线任务(引导任务)
	TaskListType_TaskListTypeGuide TaskListType = 3
)

// Enum value maps for TaskListType.
var (
	TaskListType_name = map[int32]string{
		0: "TaskListTypeUnknown",
		1: "TaskListTypeDaily",
		2: "TaskListTypeRewarded",
		3: "TaskListTypeGuide",
	}
	TaskListType_value = map[string]int32{
		"TaskListTypeUnknown":  0,
		"TaskListTypeDaily":    1,
		"TaskListTypeRewarded": 2,
		"TaskListTypeGuide":    3,
	}
)

func (x TaskListType) Enum() *TaskListType {
	p := new(TaskListType)
	*p = x
	return p
}

func (x TaskListType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskListType) Descriptor() protoreflect.EnumDescriptor {
	return file_task_proto_enumTypes[1].Descriptor()
}

func (TaskListType) Type() protoreflect.EnumType {
	return &file_task_proto_enumTypes[1]
}

func (x TaskListType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskListType.Descriptor instead.
func (TaskListType) EnumDescriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{1}
}

type TaskOptionItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemCid int32  `protobuf:"varint,1,opt,name=item_cid,json=itemCid,proto3" json:"item_cid,omitempty"`
	Num     int32  `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"` // max 10
	NftId   string `protobuf:"bytes,3,opt,name=nft_id,json=nftId,proto3" json:"nft_id,omitempty"`
}

func (x *TaskOptionItem) Reset() {
	*x = TaskOptionItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskOptionItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskOptionItem) ProtoMessage() {}

func (x *TaskOptionItem) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskOptionItem.ProtoReflect.Descriptor instead.
func (*TaskOptionItem) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{0}
}

func (x *TaskOptionItem) GetItemCid() int32 {
	if x != nil {
		return x.ItemCid
	}
	return 0
}

func (x *TaskOptionItem) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *TaskOptionItem) GetNftId() string {
	if x != nil {
		return x.NftId
	}
	return ""
}

type TaskOptionMonster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MonCid int32 `protobuf:"varint,1,opt,name=mon_cid,json=monCid,proto3" json:"mon_cid,omitempty"`
	Num    int32 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *TaskOptionMonster) Reset() {
	*x = TaskOptionMonster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskOptionMonster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskOptionMonster) ProtoMessage() {}

func (x *TaskOptionMonster) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskOptionMonster.ProtoReflect.Descriptor instead.
func (*TaskOptionMonster) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskOptionMonster) GetMonCid() int32 {
	if x != nil {
		return x.MonCid
	}
	return 0
}

func (x *TaskOptionMonster) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type TaskOptionMoveTo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	R int32 `protobuf:"varint,1,opt,name=r,proto3" json:"r,omitempty"`
	C int32 `protobuf:"varint,2,opt,name=c,proto3" json:"c,omitempty"`
}

func (x *TaskOptionMoveTo) Reset() {
	*x = TaskOptionMoveTo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskOptionMoveTo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskOptionMoveTo) ProtoMessage() {}

func (x *TaskOptionMoveTo) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskOptionMoveTo.ProtoReflect.Descriptor instead.
func (*TaskOptionMoveTo) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{2}
}

func (x *TaskOptionMoveTo) GetR() int32 {
	if x != nil {
		return x.R
	}
	return 0
}

func (x *TaskOptionMoveTo) GetC() int32 {
	if x != nil {
		return x.C
	}
	return 0
}

type TaskOptionQuiz struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuizType int32 `protobuf:"varint,1,opt,name=quiz_type,json=quizType,proto3" json:"quiz_type,omitempty"`
	QuizNum  int32 `protobuf:"varint,2,opt,name=quiz_num,json=quizNum,proto3" json:"quiz_num,omitempty"`
}

func (x *TaskOptionQuiz) Reset() {
	*x = TaskOptionQuiz{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskOptionQuiz) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskOptionQuiz) ProtoMessage() {}

func (x *TaskOptionQuiz) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskOptionQuiz.ProtoReflect.Descriptor instead.
func (*TaskOptionQuiz) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{3}
}

func (x *TaskOptionQuiz) GetQuizType() int32 {
	if x != nil {
		return x.QuizType
	}
	return 0
}

func (x *TaskOptionQuiz) GetQuizNum() int32 {
	if x != nil {
		return x.QuizNum
	}
	return 0
}

type TaskOptionCnf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind TaskOptionType `protobuf:"varint,1,opt,name=kind,proto3,enum=gameMessageCore.TaskOptionType" json:"kind,omitempty"`
	// Types that are assignable to Data:
	//	*TaskOptionCnf_Num
	//	*TaskOptionCnf_Item
	//	*TaskOptionCnf_MonInfo
	//	*TaskOptionCnf_TarPos
	//	*TaskOptionCnf_QuizInfo
	Data isTaskOptionCnf_Data `protobuf_oneof:"data"`
}

func (x *TaskOptionCnf) Reset() {
	*x = TaskOptionCnf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskOptionCnf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskOptionCnf) ProtoMessage() {}

func (x *TaskOptionCnf) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskOptionCnf.ProtoReflect.Descriptor instead.
func (*TaskOptionCnf) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{4}
}

func (x *TaskOptionCnf) GetKind() TaskOptionType {
	if x != nil {
		return x.Kind
	}
	return TaskOptionType_UnknownTaskOptionType
}

func (m *TaskOptionCnf) GetData() isTaskOptionCnf_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *TaskOptionCnf) GetNum() int32 {
	if x, ok := x.GetData().(*TaskOptionCnf_Num); ok {
		return x.Num
	}
	return 0
}

func (x *TaskOptionCnf) GetItem() *TaskOptionItem {
	if x, ok := x.GetData().(*TaskOptionCnf_Item); ok {
		return x.Item
	}
	return nil
}

func (x *TaskOptionCnf) GetMonInfo() *TaskOptionMonster {
	if x, ok := x.GetData().(*TaskOptionCnf_MonInfo); ok {
		return x.MonInfo
	}
	return nil
}

func (x *TaskOptionCnf) GetTarPos() *TaskOptionMoveTo {
	if x, ok := x.GetData().(*TaskOptionCnf_TarPos); ok {
		return x.TarPos
	}
	return nil
}

func (x *TaskOptionCnf) GetQuizInfo() *TaskOptionQuiz {
	if x, ok := x.GetData().(*TaskOptionCnf_QuizInfo); ok {
		return x.QuizInfo
	}
	return nil
}

type isTaskOptionCnf_Data interface {
	isTaskOptionCnf_Data()
}

type TaskOptionCnf_Num struct {
	Num int32 `protobuf:"varint,11,opt,name=num,proto3,oneof"`
}

type TaskOptionCnf_Item struct {
	Item *TaskOptionItem `protobuf:"bytes,12,opt,name=item,proto3,oneof"`
}

type TaskOptionCnf_MonInfo struct {
	MonInfo *TaskOptionMonster `protobuf:"bytes,13,opt,name=mon_info,json=monInfo,proto3,oneof"`
}

type TaskOptionCnf_TarPos struct {
	TarPos *TaskOptionMoveTo `protobuf:"bytes,14,opt,name=tar_pos,json=tarPos,proto3,oneof"`
}

type TaskOptionCnf_QuizInfo struct {
	QuizInfo *TaskOptionQuiz `protobuf:"bytes,15,opt,name=quiz_info,json=quizInfo,proto3,oneof"`
}

func (*TaskOptionCnf_Num) isTaskOptionCnf_Data() {}

func (*TaskOptionCnf_Item) isTaskOptionCnf_Data() {}

func (*TaskOptionCnf_MonInfo) isTaskOptionCnf_Data() {}

func (*TaskOptionCnf_TarPos) isTaskOptionCnf_Data() {}

func (*TaskOptionCnf_QuizInfo) isTaskOptionCnf_Data() {}

type TaskOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 配置
	OptionCnf *TaskOptionCnf `protobuf:"bytes,1,opt,name=option_cnf,json=optionCnf,proto3" json:"option_cnf,omitempty"`
	// 进度
	Rate int32 `protobuf:"varint,2,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (x *TaskOption) Reset() {
	*x = TaskOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskOption) ProtoMessage() {}

func (x *TaskOption) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskOption.ProtoReflect.Descriptor instead.
func (*TaskOption) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{5}
}

func (x *TaskOption) GetOptionCnf() *TaskOptionCnf {
	if x != nil {
		return x.OptionCnf
	}
	return nil
}

func (x *TaskOption) GetRate() int32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

// 任务
type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务 ID
	TaskId int32 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// 任务类型
	TaskKind TaskOptionType `protobuf:"varint,2,opt,name=task_kind,json=taskKind,proto3,enum=gameMessageCore.TaskOptionType" json:"task_kind,omitempty"`
	// 任务的接取时间
	CreatedAtSec int64 `protobuf:"varint,3,opt,name=created_at_sec,json=createdAtSec,proto3" json:"created_at_sec,omitempty"`
	// 子项信息
	Options []*TaskOption `protobuf:"bytes,4,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{6}
}

func (x *Task) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *Task) GetTaskKind() TaskOptionType {
	if x != nil {
		return x.TaskKind
	}
	return TaskOptionType_UnknownTaskOptionType
}

func (x *Task) GetCreatedAtSec() int64 {
	if x != nil {
		return x.CreatedAtSec
	}
	return 0
}

func (x *Task) GetOptions() []*TaskOption {
	if x != nil {
		return x.Options
	}
	return nil
}

type TaskList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// task list id
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// task list type
	Kind TaskListType `protobuf:"varint,2,opt,name=kind,proto3,enum=gameMessageCore.TaskListType" json:"kind,omitempty"`
	// 任务链是否可以在接取
	CanReceive bool `protobuf:"varint,3,opt,name=can_receive,json=canReceive,proto3" json:"can_receive,omitempty"`
	// 任务链是否正在进行
	Doing bool `protobuf:"varint,4,opt,name=doing,proto3" json:"doing,omitempty"`
	// 进度
	Rate int32 `protobuf:"varint,5,opt,name=rate,proto3" json:"rate,omitempty"`
	// current task
	CurTask *Task `protobuf:"bytes,6,opt,name=cur_task,json=curTask,proto3" json:"cur_task,omitempty"`
	// 任务奖励的领取进度(0, 1, 2)
	ReceiveReward int32 `protobuf:"varint,7,opt,name=receive_reward,json=receiveReward,proto3" json:"receive_reward,omitempty"`
}

func (x *TaskList) Reset() {
	*x = TaskList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskList) ProtoMessage() {}

func (x *TaskList) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskList.ProtoReflect.Descriptor instead.
func (*TaskList) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{7}
}

func (x *TaskList) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskList) GetKind() TaskListType {
	if x != nil {
		return x.Kind
	}
	return TaskListType_TaskListTypeUnknown
}

func (x *TaskList) GetCanReceive() bool {
	if x != nil {
		return x.CanReceive
	}
	return false
}

func (x *TaskList) GetDoing() bool {
	if x != nil {
		return x.Doing
	}
	return false
}

func (x *TaskList) GetRate() int32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *TaskList) GetCurTask() *Task {
	if x != nil {
		return x.CurTask
	}
	return nil
}

func (x *TaskList) GetReceiveReward() int32 {
	if x != nil {
		return x.ReceiveReward
	}
	return 0
}

// 用户已接任务列表
type PlayerTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskLists []*TaskList `protobuf:"bytes,1,rep,name=task_lists,json=taskLists,proto3" json:"task_lists,omitempty"`
}

func (x *PlayerTask) Reset() {
	*x = PlayerTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerTask) ProtoMessage() {}

func (x *PlayerTask) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerTask.ProtoReflect.Descriptor instead.
func (*PlayerTask) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{8}
}

func (x *PlayerTask) GetTaskLists() []*TaskList {
	if x != nil {
		return x.TaskLists
	}
	return nil
}

var File_task_proto protoreflect.FileDescriptor

var file_task_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x61,
	0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x22, 0x54, 0x0a,
	0x0e, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x43, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x15, 0x0a, 0x06,
	0x6e, 0x66, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x66,
	0x74, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x5f,
	0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x6f, 0x6e, 0x43, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6e, 0x75, 0x6d, 0x22, 0x2e, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x63, 0x22, 0x48, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x51, 0x75, 0x69, 0x7a, 0x12, 0x1b, 0x0a, 0x09, 0x71, 0x75, 0x69, 0x7a, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x69, 0x7a, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x69, 0x7a, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x71, 0x75, 0x69, 0x7a, 0x4e, 0x75, 0x6d, 0x22, 0xd6, 0x02,
	0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6e, 0x66, 0x12,
	0x33, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x35, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x48, 0x00, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12,
	0x3f, 0x0a, 0x08, 0x6d, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f,
	0x6e, 0x73, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x3c, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f,
	0x76, 0x65, 0x54, 0x6f, 0x48, 0x00, 0x52, 0x06, 0x74, 0x61, 0x72, 0x50, 0x6f, 0x73, 0x12, 0x3e,
	0x0a, 0x09, 0x71, 0x75, 0x69, 0x7a, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75,
	0x69, 0x7a, 0x48, 0x00, 0x52, 0x08, 0x71, 0x75, 0x69, 0x7a, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x06,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5f, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0a, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x6e, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6e, 0x66, 0x52, 0x09, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6e, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x22, 0xba, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x09, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x74,
	0x61, 0x73, 0x6b, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x53, 0x65, 0x63, 0x12, 0x35, 0x0a,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0xf1, 0x01, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x31, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72,
	0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x61, 0x6e, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x6f, 0x69, 0x6e, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x6f, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12,
	0x30, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x07, 0x63, 0x75, 0x72, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x22, 0x46, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x38, 0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x73,
	0x2a, 0xf0, 0x01, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x61,
	0x73, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x48, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x50,
	0x69, 0x63, 0x6b, 0x55, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x4b,
	0x69, 0x6c, 0x6c, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x10, 0x06,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x6c, 0x6f, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x72, 0x61, 0x66, 0x74, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x55, 0x73, 0x65, 0x10, 0x09, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x0a, 0x12,
	0x12, 0x0a, 0x0e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x10, 0x0b, 0x2a, 0x6f, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11,
	0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x44, 0x61, 0x69, 0x6c,
	0x79, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64, 0x10, 0x02, 0x12, 0x15, 0x0a,
	0x11, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x47, 0x75, 0x69,
	0x64, 0x65, 0x10, 0x03, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x47, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_task_proto_rawDescOnce sync.Once
	file_task_proto_rawDescData = file_task_proto_rawDesc
)

func file_task_proto_rawDescGZIP() []byte {
	file_task_proto_rawDescOnce.Do(func() {
		file_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_proto_rawDescData)
	})
	return file_task_proto_rawDescData
}

var file_task_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_task_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_task_proto_goTypes = []interface{}{
	(TaskOptionType)(0),       // 0: gameMessageCore.TaskOptionType
	(TaskListType)(0),         // 1: gameMessageCore.TaskListType
	(*TaskOptionItem)(nil),    // 2: gameMessageCore.TaskOptionItem
	(*TaskOptionMonster)(nil), // 3: gameMessageCore.TaskOptionMonster
	(*TaskOptionMoveTo)(nil),  // 4: gameMessageCore.TaskOptionMoveTo
	(*TaskOptionQuiz)(nil),    // 5: gameMessageCore.TaskOptionQuiz
	(*TaskOptionCnf)(nil),     // 6: gameMessageCore.TaskOptionCnf
	(*TaskOption)(nil),        // 7: gameMessageCore.TaskOption
	(*Task)(nil),              // 8: gameMessageCore.Task
	(*TaskList)(nil),          // 9: gameMessageCore.TaskList
	(*PlayerTask)(nil),        // 10: gameMessageCore.PlayerTask
}
var file_task_proto_depIdxs = []int32{
	0,  // 0: gameMessageCore.TaskOptionCnf.kind:type_name -> gameMessageCore.TaskOptionType
	2,  // 1: gameMessageCore.TaskOptionCnf.item:type_name -> gameMessageCore.TaskOptionItem
	3,  // 2: gameMessageCore.TaskOptionCnf.mon_info:type_name -> gameMessageCore.TaskOptionMonster
	4,  // 3: gameMessageCore.TaskOptionCnf.tar_pos:type_name -> gameMessageCore.TaskOptionMoveTo
	5,  // 4: gameMessageCore.TaskOptionCnf.quiz_info:type_name -> gameMessageCore.TaskOptionQuiz
	6,  // 5: gameMessageCore.TaskOption.option_cnf:type_name -> gameMessageCore.TaskOptionCnf
	0,  // 6: gameMessageCore.Task.task_kind:type_name -> gameMessageCore.TaskOptionType
	7,  // 7: gameMessageCore.Task.options:type_name -> gameMessageCore.TaskOption
	1,  // 8: gameMessageCore.TaskList.kind:type_name -> gameMessageCore.TaskListType
	8,  // 9: gameMessageCore.TaskList.cur_task:type_name -> gameMessageCore.Task
	9,  // 10: gameMessageCore.PlayerTask.task_lists:type_name -> gameMessageCore.TaskList
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_task_proto_init() }
func file_task_proto_init() {
	if File_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskOptionItem); i {
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
		file_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskOptionMonster); i {
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
		file_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskOptionMoveTo); i {
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
		file_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskOptionQuiz); i {
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
		file_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskOptionCnf); i {
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
		file_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskOption); i {
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
		file_task_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_task_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskList); i {
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
		file_task_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerTask); i {
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
	file_task_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*TaskOptionCnf_Num)(nil),
		(*TaskOptionCnf_Item)(nil),
		(*TaskOptionCnf_MonInfo)(nil),
		(*TaskOptionCnf_TarPos)(nil),
		(*TaskOptionCnf_QuizInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_task_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_task_proto_goTypes,
		DependencyIndexes: file_task_proto_depIdxs,
		EnumInfos:         file_task_proto_enumTypes,
		MessageInfos:      file_task_proto_msgTypes,
	}.Build()
	File_task_proto = out.File
	file_task_proto_rawDesc = nil
	file_task_proto_goTypes = nil
	file_task_proto_depIdxs = nil
}
