// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: home.proto

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

type CollectResourceType int32

const (
	CollectResourceType_CollectResourceTypeUnknown CollectResourceType = 0
	// 土壤作物
	CollectResourceType_CollectResourceTypeSoil CollectResourceType = 1
	// 采集资源
	CollectResourceType_HomeResource CollectResourceType = 2
)

// Enum value maps for CollectResourceType.
var (
	CollectResourceType_name = map[int32]string{
		0: "CollectResourceTypeUnknown",
		1: "CollectResourceTypeSoil",
		2: "HomeResource",
	}
	CollectResourceType_value = map[string]int32{
		"CollectResourceTypeUnknown": 0,
		"CollectResourceTypeSoil":    1,
		"HomeResource":               2,
	}
)

func (x CollectResourceType) Enum() *CollectResourceType {
	p := new(CollectResourceType)
	*p = x
	return p
}

func (x CollectResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CollectResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_home_proto_enumTypes[0].Descriptor()
}

func (CollectResourceType) Type() protoreflect.EnumType {
	return &file_home_proto_enumTypes[0]
}

func (x CollectResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CollectResourceType.Descriptor instead.
func (CollectResourceType) EnumDescriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{0}
}

type ProxySoilData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SoilStatus       int32  `protobuf:"varint,2,opt,name=soilStatus,proto3" json:"soilStatus,omitempty"`
	StatusStartStamp int64  `protobuf:"varint,3,opt,name=statusStartStamp,proto3" json:"statusStartStamp,omitempty"`
	GrowingStage     int32  `protobuf:"varint,4,opt,name=growingStage,proto3" json:"growingStage,omitempty"`
	SeedCid          int32  `protobuf:"varint,5,opt,name=seedCid,proto3" json:"seedCid,omitempty"`
	SowingValid      bool   `protobuf:"varint,6,opt,name=sowingValid,proto3" json:"sowingValid,omitempty"`
	ManureCid        int32  `protobuf:"varint,7,opt,name=manureCid,proto3" json:"manureCid,omitempty"`
	ManureValid      bool   `protobuf:"varint,8,opt,name=manureValid,proto3" json:"manureValid,omitempty"`
	ExtraWateringNum int32  `protobuf:"varint,9,opt,name=extraWateringNum,proto3" json:"extraWateringNum,omitempty"`
}

func (x *ProxySoilData) Reset() {
	*x = ProxySoilData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxySoilData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxySoilData) ProtoMessage() {}

func (x *ProxySoilData) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxySoilData.ProtoReflect.Descriptor instead.
func (*ProxySoilData) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{0}
}

func (x *ProxySoilData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProxySoilData) GetSoilStatus() int32 {
	if x != nil {
		return x.SoilStatus
	}
	return 0
}

func (x *ProxySoilData) GetStatusStartStamp() int64 {
	if x != nil {
		return x.StatusStartStamp
	}
	return 0
}

func (x *ProxySoilData) GetGrowingStage() int32 {
	if x != nil {
		return x.GrowingStage
	}
	return 0
}

func (x *ProxySoilData) GetSeedCid() int32 {
	if x != nil {
		return x.SeedCid
	}
	return 0
}

func (x *ProxySoilData) GetSowingValid() bool {
	if x != nil {
		return x.SowingValid
	}
	return false
}

func (x *ProxySoilData) GetManureCid() int32 {
	if x != nil {
		return x.ManureCid
	}
	return 0
}

func (x *ProxySoilData) GetManureValid() bool {
	if x != nil {
		return x.ManureValid
	}
	return false
}

func (x *ProxySoilData) GetExtraWateringNum() int32 {
	if x != nil {
		return x.ExtraWateringNum
	}
	return 0
}

type CollectResourceBaseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type CollectResourceType `protobuf:"varint,2,opt,name=type,proto3,enum=gameMessageCore.CollectResourceType" json:"type,omitempty"`
}

func (x *CollectResourceBaseInfo) Reset() {
	*x = CollectResourceBaseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectResourceBaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectResourceBaseInfo) ProtoMessage() {}

func (x *CollectResourceBaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectResourceBaseInfo.ProtoReflect.Descriptor instead.
func (*CollectResourceBaseInfo) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{1}
}

func (x *CollectResourceBaseInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CollectResourceBaseInfo) GetType() CollectResourceType {
	if x != nil {
		return x.Type
	}
	return CollectResourceType_CollectResourceTypeUnknown
}

//使用采集技能的信息
type UseCollectResourceSkillInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Targets []*CollectResourceBaseInfo `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`
	//消耗道具的cid 比如种子 肥料
	CostItemCid int32 `protobuf:"varint,2,opt,name=costItemCid,proto3" json:"costItemCid,omitempty"`
}

func (x *UseCollectResourceSkillInfo) Reset() {
	*x = UseCollectResourceSkillInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UseCollectResourceSkillInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UseCollectResourceSkillInfo) ProtoMessage() {}

func (x *UseCollectResourceSkillInfo) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UseCollectResourceSkillInfo.ProtoReflect.Descriptor instead.
func (*UseCollectResourceSkillInfo) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{2}
}

func (x *UseCollectResourceSkillInfo) GetTargets() []*CollectResourceBaseInfo {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *UseCollectResourceSkillInfo) GetCostItemCid() int32 {
	if x != nil {
		return x.CostItemCid
	}
	return 0
}

//土地进度信息
type HomeSoilProgressInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64                         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProgressInfo *CollectResourceProgressResult `protobuf:"bytes,2,opt,name=progressInfo,proto3" json:"progressInfo,omitempty"`
}

func (x *HomeSoilProgressInfo) Reset() {
	*x = HomeSoilProgressInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeSoilProgressInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeSoilProgressInfo) ProtoMessage() {}

func (x *HomeSoilProgressInfo) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeSoilProgressInfo.ProtoReflect.Descriptor instead.
func (*HomeSoilProgressInfo) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{3}
}

func (x *HomeSoilProgressInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HomeSoilProgressInfo) GetProgressInfo() *CollectResourceProgressResult {
	if x != nil {
		return x.ProgressInfo
	}
	return nil
}

//家园某个采集资源操作结果
type CollectResourceOperateResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetInfo *CollectResourceBaseInfo `protobuf:"bytes,1,opt,name=targetInfo,proto3" json:"targetInfo,omitempty"`
	//修改进度结果 不影响操作执行 仅仅表示进度的变化
	ProgressResult *CollectResourceProgressResult `protobuf:"bytes,2,opt,name=progressResult,proto3" json:"progressResult,omitempty"`
	//动作执行结果
	ExecuteResult *CollectResourceExecuteResult `protobuf:"bytes,3,opt,name=executeResult,proto3" json:"executeResult,omitempty"`
}

func (x *CollectResourceOperateResult) Reset() {
	*x = CollectResourceOperateResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectResourceOperateResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectResourceOperateResult) ProtoMessage() {}

func (x *CollectResourceOperateResult) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectResourceOperateResult.ProtoReflect.Descriptor instead.
func (*CollectResourceOperateResult) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{4}
}

func (x *CollectResourceOperateResult) GetTargetInfo() *CollectResourceBaseInfo {
	if x != nil {
		return x.TargetInfo
	}
	return nil
}

func (x *CollectResourceOperateResult) GetProgressResult() *CollectResourceProgressResult {
	if x != nil {
		return x.ProgressResult
	}
	return nil
}

func (x *CollectResourceOperateResult) GetExecuteResult() *CollectResourceExecuteResult {
	if x != nil {
		return x.ExecuteResult
	}
	return nil
}

type CollectResourceProgressResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//最新总进度值 适用于分段式的进度
	TotalProgress int32 `protobuf:"varint,1,opt,name=totalProgress,proto3" json:"totalProgress,omitempty"`
	//进度充满时的时间戳 适用于浇水连续式的进度
	ProgressFullStamp int64 `protobuf:"varint,2,opt,name=progressFullStamp,proto3" json:"progressFullStamp,omitempty"`
}

func (x *CollectResourceProgressResult) Reset() {
	*x = CollectResourceProgressResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectResourceProgressResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectResourceProgressResult) ProtoMessage() {}

func (x *CollectResourceProgressResult) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectResourceProgressResult.ProtoReflect.Descriptor instead.
func (*CollectResourceProgressResult) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{5}
}

func (x *CollectResourceProgressResult) GetTotalProgress() int32 {
	if x != nil {
		return x.TotalProgress
	}
	return 0
}

func (x *CollectResourceProgressResult) GetProgressFullStamp() int64 {
	if x != nil {
		return x.ProgressFullStamp
	}
	return 0
}

type CollectResourceExecuteResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//道具有效  在播种和施肥时候专精不够时候 会返回false
	ItemValid bool `protobuf:"varint,1,opt,name=itemValid,proto3" json:"itemValid,omitempty"`
	//额外浇水次数
	ExtraWateringNum int32           `protobuf:"varint,2,opt,name=ExtraWateringNum,proto3" json:"ExtraWateringNum,omitempty"`
	DropList         []*ItemBaseInfo `protobuf:"bytes,3,rep,name=dropList,proto3" json:"dropList,omitempty"`
}

func (x *CollectResourceExecuteResult) Reset() {
	*x = CollectResourceExecuteResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectResourceExecuteResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectResourceExecuteResult) ProtoMessage() {}

func (x *CollectResourceExecuteResult) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectResourceExecuteResult.ProtoReflect.Descriptor instead.
func (*CollectResourceExecuteResult) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{6}
}

func (x *CollectResourceExecuteResult) GetItemValid() bool {
	if x != nil {
		return x.ItemValid
	}
	return false
}

func (x *CollectResourceExecuteResult) GetExtraWateringNum() int32 {
	if x != nil {
		return x.ExtraWateringNum
	}
	return 0
}

func (x *CollectResourceExecuteResult) GetDropList() []*ItemBaseInfo {
	if x != nil {
		return x.DropList
	}
	return nil
}

type ProxyAnimalData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimId          uint64            `protobuf:"varint,1,opt,name=animId,proto3" json:"animId,omitempty"`
	HungerProgress  int32             `protobuf:"varint,2,opt,name=hungerProgress,proto3" json:"hungerProgress,omitempty"`
	HarvestProgress int32             `protobuf:"varint,3,opt,name=harvestProgress,proto3" json:"harvestProgress,omitempty"`
	IsComforted     bool              `protobuf:"varint,4,opt,name=isComforted,proto3" json:"isComforted,omitempty"`
	IsDead          bool              `protobuf:"varint,5,opt,name=isDead,proto3" json:"isDead,omitempty"`
	ProductData     *ProxyProductData `protobuf:"bytes,6,opt,name=productData,proto3" json:"productData,omitempty"`
}

func (x *ProxyAnimalData) Reset() {
	*x = ProxyAnimalData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyAnimalData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyAnimalData) ProtoMessage() {}

func (x *ProxyAnimalData) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyAnimalData.ProtoReflect.Descriptor instead.
func (*ProxyAnimalData) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{7}
}

func (x *ProxyAnimalData) GetAnimId() uint64 {
	if x != nil {
		return x.AnimId
	}
	return 0
}

func (x *ProxyAnimalData) GetHungerProgress() int32 {
	if x != nil {
		return x.HungerProgress
	}
	return 0
}

func (x *ProxyAnimalData) GetHarvestProgress() int32 {
	if x != nil {
		return x.HarvestProgress
	}
	return 0
}

func (x *ProxyAnimalData) GetIsComforted() bool {
	if x != nil {
		return x.IsComforted
	}
	return false
}

func (x *ProxyAnimalData) GetIsDead() bool {
	if x != nil {
		return x.IsDead
	}
	return false
}

func (x *ProxyAnimalData) GetProductData() *ProxyProductData {
	if x != nil {
		return x.ProductData
	}
	return nil
}

type ProxyAnimalBaseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimId       uint64 `protobuf:"varint,1,opt,name=animId,proto3" json:"animId,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cid          int32  `protobuf:"varint,3,opt,name=cid,proto3" json:"cid,omitempty"`
	Favorability int32  `protobuf:"varint,4,opt,name=favorability,proto3" json:"favorability,omitempty"`
}

func (x *ProxyAnimalBaseData) Reset() {
	*x = ProxyAnimalBaseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyAnimalBaseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyAnimalBaseData) ProtoMessage() {}

func (x *ProxyAnimalBaseData) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyAnimalBaseData.ProtoReflect.Descriptor instead.
func (*ProxyAnimalBaseData) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{8}
}

func (x *ProxyAnimalBaseData) GetAnimId() uint64 {
	if x != nil {
		return x.AnimId
	}
	return 0
}

func (x *ProxyAnimalBaseData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProxyAnimalBaseData) GetCid() int32 {
	if x != nil {
		return x.Cid
	}
	return 0
}

func (x *ProxyAnimalBaseData) GetFavorability() int32 {
	if x != nil {
		return x.Favorability
	}
	return 0
}

type ProxyProductData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int32    `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	ItemCid   int32    `protobuf:"varint,2,opt,name=itemCid,proto3" json:"itemCid,omitempty"`
	Position  *Vector3 `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *ProxyProductData) Reset() {
	*x = ProxyProductData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_home_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyProductData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyProductData) ProtoMessage() {}

func (x *ProxyProductData) ProtoReflect() protoreflect.Message {
	mi := &file_home_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyProductData.ProtoReflect.Descriptor instead.
func (*ProxyProductData) Descriptor() ([]byte, []int) {
	return file_home_proto_rawDescGZIP(), []int{9}
}

func (x *ProxyProductData) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *ProxyProductData) GetItemCid() int32 {
	if x != nil {
		return x.ItemCid
	}
	return 0
}

func (x *ProxyProductData) GetPosition() *Vector3 {
	if x != nil {
		return x.Position
	}
	return nil
}

var File_home_proto protoreflect.FileDescriptor

var file_home_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x61,
	0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x1a, 0x0a, 0x69,
	0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x02, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x53, 0x6f, 0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x69,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x6f, 0x69, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x67, 0x72, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x65,
	0x64, 0x43, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x65, 0x65, 0x64,
	0x43, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x75, 0x72, 0x65, 0x43,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x75, 0x72, 0x65,
	0x43, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x6e, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6d, 0x61, 0x6e, 0x75, 0x72, 0x65,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x72, 0x61, 0x57, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x10, 0x65, 0x78, 0x74, 0x72, 0x61, 0x57, 0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x4e, 0x75,
	0x6d, 0x22, 0x63, 0x0a, 0x17, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x1b, 0x55, 0x73, 0x65, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x42, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x63, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x69, 0x64, 0x22, 0x7a, 0x0a, 0x14,
	0x48, 0x6f, 0x6d, 0x65, 0x53, 0x6f, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x52, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x95, 0x02, 0x0a, 0x1c, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x48, 0x0a, 0x0a, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42,
	0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x56, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0e, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x53, 0x0a, 0x0d, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x73, 0x0a, 0x1d, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x46, 0x75, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x46, 0x75, 0x6c, 0x6c,
	0x53, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xa3, 0x01, 0x0a, 0x1c, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x45, 0x78, 0x74, 0x72, 0x61, 0x57, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x57, 0x61, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d,
	0x12, 0x39, 0x0a, 0x08, 0x64, 0x72, 0x6f, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x08, 0x64, 0x72, 0x6f, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xfa, 0x01, 0x0a, 0x0f,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6e, 0x69, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x61, 0x6e, 0x69, 0x6d, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x68, 0x75, 0x6e, 0x67, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x68, 0x75, 0x6e, 0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x28, 0x0a, 0x0f, 0x68, 0x61, 0x72, 0x76, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x68, 0x61, 0x72, 0x76, 0x65, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x43,
	0x6f, 0x6d, 0x66, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x69, 0x73, 0x43, 0x6f, 0x6d, 0x66, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x73, 0x44, 0x65, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x44,
	0x65, 0x61, 0x64, 0x12, 0x43, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x77, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x42, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6e, 0x69, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x61, 0x6e, 0x69, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x22, 0x80, 0x01, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x43, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x43, 0x69, 0x64, 0x12, 0x34,
	0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x72, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x64, 0x0a, 0x13, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x53, 0x6f, 0x69, 0x6c, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x6f, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x10, 0x02, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_home_proto_rawDescOnce sync.Once
	file_home_proto_rawDescData = file_home_proto_rawDesc
)

func file_home_proto_rawDescGZIP() []byte {
	file_home_proto_rawDescOnce.Do(func() {
		file_home_proto_rawDescData = protoimpl.X.CompressGZIP(file_home_proto_rawDescData)
	})
	return file_home_proto_rawDescData
}

var file_home_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_home_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_home_proto_goTypes = []interface{}{
	(CollectResourceType)(0),              // 0: gameMessageCore.CollectResourceType
	(*ProxySoilData)(nil),                 // 1: gameMessageCore.ProxySoilData
	(*CollectResourceBaseInfo)(nil),       // 2: gameMessageCore.CollectResourceBaseInfo
	(*UseCollectResourceSkillInfo)(nil),   // 3: gameMessageCore.UseCollectResourceSkillInfo
	(*HomeSoilProgressInfo)(nil),          // 4: gameMessageCore.HomeSoilProgressInfo
	(*CollectResourceOperateResult)(nil),  // 5: gameMessageCore.CollectResourceOperateResult
	(*CollectResourceProgressResult)(nil), // 6: gameMessageCore.CollectResourceProgressResult
	(*CollectResourceExecuteResult)(nil),  // 7: gameMessageCore.CollectResourceExecuteResult
	(*ProxyAnimalData)(nil),               // 8: gameMessageCore.ProxyAnimalData
	(*ProxyAnimalBaseData)(nil),           // 9: gameMessageCore.ProxyAnimalBaseData
	(*ProxyProductData)(nil),              // 10: gameMessageCore.ProxyProductData
	(*ItemBaseInfo)(nil),                  // 11: gameMessageCore.ItemBaseInfo
	(*Vector3)(nil),                       // 12: gameMessageCore.Vector3
}
var file_home_proto_depIdxs = []int32{
	0,  // 0: gameMessageCore.CollectResourceBaseInfo.type:type_name -> gameMessageCore.CollectResourceType
	2,  // 1: gameMessageCore.UseCollectResourceSkillInfo.targets:type_name -> gameMessageCore.CollectResourceBaseInfo
	6,  // 2: gameMessageCore.HomeSoilProgressInfo.progressInfo:type_name -> gameMessageCore.CollectResourceProgressResult
	2,  // 3: gameMessageCore.CollectResourceOperateResult.targetInfo:type_name -> gameMessageCore.CollectResourceBaseInfo
	6,  // 4: gameMessageCore.CollectResourceOperateResult.progressResult:type_name -> gameMessageCore.CollectResourceProgressResult
	7,  // 5: gameMessageCore.CollectResourceOperateResult.executeResult:type_name -> gameMessageCore.CollectResourceExecuteResult
	11, // 6: gameMessageCore.CollectResourceExecuteResult.dropList:type_name -> gameMessageCore.ItemBaseInfo
	10, // 7: gameMessageCore.ProxyAnimalData.productData:type_name -> gameMessageCore.ProxyProductData
	12, // 8: gameMessageCore.ProxyProductData.position:type_name -> gameMessageCore.Vector3
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_home_proto_init() }
func file_home_proto_init() {
	if File_home_proto != nil {
		return
	}
	file_item_proto_init()
	file_vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_home_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxySoilData); i {
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
		file_home_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectResourceBaseInfo); i {
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
		file_home_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UseCollectResourceSkillInfo); i {
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
		file_home_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeSoilProgressInfo); i {
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
		file_home_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectResourceOperateResult); i {
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
		file_home_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectResourceProgressResult); i {
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
		file_home_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectResourceExecuteResult); i {
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
		file_home_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyAnimalData); i {
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
		file_home_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyAnimalBaseData); i {
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
		file_home_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyProductData); i {
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
			RawDescriptor: file_home_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_home_proto_goTypes,
		DependencyIndexes: file_home_proto_depIdxs,
		EnumInfos:         file_home_proto_enumTypes,
		MessageInfos:      file_home_proto_msgTypes,
	}.Build()
	File_home_proto = out.File
	file_home_proto_rawDesc = nil
	file_home_proto_goTypes = nil
	file_home_proto_depIdxs = nil
}
