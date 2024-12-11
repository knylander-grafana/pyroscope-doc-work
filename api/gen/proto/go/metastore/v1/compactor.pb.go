// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: metastore/v1/compactor.proto

package metastorev1

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

type CompactionJobStatus int32

const (
	CompactionJobStatus_COMPACTION_STATUS_UNSPECIFIED CompactionJobStatus = 0
	CompactionJobStatus_COMPACTION_STATUS_IN_PROGRESS CompactionJobStatus = 1
	CompactionJobStatus_COMPACTION_STATUS_SUCCESS     CompactionJobStatus = 2
)

// Enum value maps for CompactionJobStatus.
var (
	CompactionJobStatus_name = map[int32]string{
		0: "COMPACTION_STATUS_UNSPECIFIED",
		1: "COMPACTION_STATUS_IN_PROGRESS",
		2: "COMPACTION_STATUS_SUCCESS",
	}
	CompactionJobStatus_value = map[string]int32{
		"COMPACTION_STATUS_UNSPECIFIED": 0,
		"COMPACTION_STATUS_IN_PROGRESS": 1,
		"COMPACTION_STATUS_SUCCESS":     2,
	}
)

func (x CompactionJobStatus) Enum() *CompactionJobStatus {
	p := new(CompactionJobStatus)
	*p = x
	return p
}

func (x CompactionJobStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompactionJobStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_metastore_v1_compactor_proto_enumTypes[0].Descriptor()
}

func (CompactionJobStatus) Type() protoreflect.EnumType {
	return &file_metastore_v1_compactor_proto_enumTypes[0]
}

func (x CompactionJobStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompactionJobStatus.Descriptor instead.
func (CompactionJobStatus) EnumDescriptor() ([]byte, []int) {
	return file_metastore_v1_compactor_proto_rawDescGZIP(), []int{0}
}

type PollCompactionJobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusUpdates []*CompactionJobStatusUpdate `protobuf:"bytes,1,rep,name=status_updates,json=statusUpdates,proto3" json:"status_updates,omitempty"`
	// How many new jobs a worker can be assigned to.
	JobCapacity uint32 `protobuf:"varint,2,opt,name=job_capacity,json=jobCapacity,proto3" json:"job_capacity,omitempty"`
}

func (x *PollCompactionJobsRequest) Reset() {
	*x = PollCompactionJobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_compactor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollCompactionJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollCompactionJobsRequest) ProtoMessage() {}

func (x *PollCompactionJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_compactor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollCompactionJobsRequest.ProtoReflect.Descriptor instead.
func (*PollCompactionJobsRequest) Descriptor() ([]byte, []int) {
	return file_metastore_v1_compactor_proto_rawDescGZIP(), []int{0}
}

func (x *PollCompactionJobsRequest) GetStatusUpdates() []*CompactionJobStatusUpdate {
	if x != nil {
		return x.StatusUpdates
	}
	return nil
}

func (x *PollCompactionJobsRequest) GetJobCapacity() uint32 {
	if x != nil {
		return x.JobCapacity
	}
	return 0
}

type PollCompactionJobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompactionJobs []*CompactionJob           `protobuf:"bytes,1,rep,name=compaction_jobs,json=compactionJobs,proto3" json:"compaction_jobs,omitempty"`
	Assignments    []*CompactionJobAssignment `protobuf:"bytes,2,rep,name=assignments,proto3" json:"assignments,omitempty"`
}

func (x *PollCompactionJobsResponse) Reset() {
	*x = PollCompactionJobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_compactor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollCompactionJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollCompactionJobsResponse) ProtoMessage() {}

func (x *PollCompactionJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_compactor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollCompactionJobsResponse.ProtoReflect.Descriptor instead.
func (*PollCompactionJobsResponse) Descriptor() ([]byte, []int) {
	return file_metastore_v1_compactor_proto_rawDescGZIP(), []int{1}
}

func (x *PollCompactionJobsResponse) GetCompactionJobs() []*CompactionJob {
	if x != nil {
		return x.CompactionJobs
	}
	return nil
}

func (x *PollCompactionJobsResponse) GetAssignments() []*CompactionJobAssignment {
	if x != nil {
		return x.Assignments
	}
	return nil
}

type CompactionJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Shard           uint32        `protobuf:"varint,2,opt,name=shard,proto3" json:"shard,omitempty"`
	Tenant          string        `protobuf:"bytes,3,opt,name=tenant,proto3" json:"tenant,omitempty"`
	CompactionLevel uint32        `protobuf:"varint,4,opt,name=compaction_level,json=compactionLevel,proto3" json:"compaction_level,omitempty"`
	SourceBlocks    []string      `protobuf:"bytes,5,rep,name=source_blocks,json=sourceBlocks,proto3" json:"source_blocks,omitempty"`
	Tombstones      []*Tombstones `protobuf:"bytes,6,rep,name=tombstones,proto3" json:"tombstones,omitempty"`
}

func (x *CompactionJob) Reset() {
	*x = CompactionJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_compactor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactionJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactionJob) ProtoMessage() {}

func (x *CompactionJob) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_compactor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactionJob.ProtoReflect.Descriptor instead.
func (*CompactionJob) Descriptor() ([]byte, []int) {
	return file_metastore_v1_compactor_proto_rawDescGZIP(), []int{2}
}

func (x *CompactionJob) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CompactionJob) GetShard() uint32 {
	if x != nil {
		return x.Shard
	}
	return 0
}

func (x *CompactionJob) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *CompactionJob) GetCompactionLevel() uint32 {
	if x != nil {
		return x.CompactionLevel
	}
	return 0
}

func (x *CompactionJob) GetSourceBlocks() []string {
	if x != nil {
		return x.SourceBlocks
	}
	return nil
}

func (x *CompactionJob) GetTombstones() []*Tombstones {
	if x != nil {
		return x.Tombstones
	}
	return nil
}

// Tombstones represent objects removed from the index but still stored.
type Tombstones struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks *BlockTombstones `protobuf:"bytes,1,opt,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *Tombstones) Reset() {
	*x = Tombstones{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_compactor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tombstones) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tombstones) ProtoMessage() {}

func (x *Tombstones) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_compactor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tombstones.ProtoReflect.Descriptor instead.
func (*Tombstones) Descriptor() ([]byte, []int) {
	return file_metastore_v1_compactor_proto_rawDescGZIP(), []int{3}
}

func (x *Tombstones) GetBlocks() *BlockTombstones {
	if x != nil {
		return x.Blocks
	}
	return nil
}

type BlockTombstones struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Shard           uint32   `protobuf:"varint,2,opt,name=shard,proto3" json:"shard,omitempty"`
	Tenant          string   `protobuf:"bytes,3,opt,name=tenant,proto3" json:"tenant,omitempty"`
	CompactionLevel uint32   `protobuf:"varint,4,opt,name=compaction_level,json=compactionLevel,proto3" json:"compaction_level,omitempty"`
	Blocks          []string `protobuf:"bytes,5,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *BlockTombstones) Reset() {
	*x = BlockTombstones{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_compactor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockTombstones) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockTombstones) ProtoMessage() {}

func (x *BlockTombstones) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_compactor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockTombstones.ProtoReflect.Descriptor instead.
func (*BlockTombstones) Descriptor() ([]byte, []int) {
	return file_metastore_v1_compactor_proto_rawDescGZIP(), []int{4}
}

func (x *BlockTombstones) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BlockTombstones) GetShard() uint32 {
	if x != nil {
		return x.Shard
	}
	return 0
}

func (x *BlockTombstones) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *BlockTombstones) GetCompactionLevel() uint32 {
	if x != nil {
		return x.CompactionLevel
	}
	return 0
}

func (x *BlockTombstones) GetBlocks() []string {
	if x != nil {
		return x.Blocks
	}
	return nil
}

type CompactionJobAssignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Token          uint64 `protobuf:"varint,2,opt,name=token,proto3" json:"token,omitempty"`
	LeaseExpiresAt int64  `protobuf:"varint,3,opt,name=lease_expires_at,json=leaseExpiresAt,proto3" json:"lease_expires_at,omitempty"`
}

func (x *CompactionJobAssignment) Reset() {
	*x = CompactionJobAssignment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_compactor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactionJobAssignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactionJobAssignment) ProtoMessage() {}

func (x *CompactionJobAssignment) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_compactor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactionJobAssignment.ProtoReflect.Descriptor instead.
func (*CompactionJobAssignment) Descriptor() ([]byte, []int) {
	return file_metastore_v1_compactor_proto_rawDescGZIP(), []int{5}
}

func (x *CompactionJobAssignment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CompactionJobAssignment) GetToken() uint64 {
	if x != nil {
		return x.Token
	}
	return 0
}

func (x *CompactionJobAssignment) GetLeaseExpiresAt() int64 {
	if x != nil {
		return x.LeaseExpiresAt
	}
	return 0
}

type CompactionJobStatusUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Token  uint64              `protobuf:"varint,2,opt,name=token,proto3" json:"token,omitempty"`
	Status CompactionJobStatus `protobuf:"varint,3,opt,name=status,proto3,enum=metastore.v1.CompactionJobStatus" json:"status,omitempty"`
	// Only present if the job completed successfully.
	CompactedBlocks *CompactedBlocks `protobuf:"bytes,4,opt,name=compacted_blocks,json=compactedBlocks,proto3" json:"compacted_blocks,omitempty"`
}

func (x *CompactionJobStatusUpdate) Reset() {
	*x = CompactionJobStatusUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_compactor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactionJobStatusUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactionJobStatusUpdate) ProtoMessage() {}

func (x *CompactionJobStatusUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_compactor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactionJobStatusUpdate.ProtoReflect.Descriptor instead.
func (*CompactionJobStatusUpdate) Descriptor() ([]byte, []int) {
	return file_metastore_v1_compactor_proto_rawDescGZIP(), []int{6}
}

func (x *CompactionJobStatusUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CompactionJobStatusUpdate) GetToken() uint64 {
	if x != nil {
		return x.Token
	}
	return 0
}

func (x *CompactionJobStatusUpdate) GetStatus() CompactionJobStatus {
	if x != nil {
		return x.Status
	}
	return CompactionJobStatus_COMPACTION_STATUS_UNSPECIFIED
}

func (x *CompactionJobStatusUpdate) GetCompactedBlocks() *CompactedBlocks {
	if x != nil {
		return x.CompactedBlocks
	}
	return nil
}

type CompactedBlocks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceBlocks *BlockList   `protobuf:"bytes,1,opt,name=source_blocks,json=sourceBlocks,proto3" json:"source_blocks,omitempty"`
	NewBlocks    []*BlockMeta `protobuf:"bytes,2,rep,name=new_blocks,json=newBlocks,proto3" json:"new_blocks,omitempty"`
}

func (x *CompactedBlocks) Reset() {
	*x = CompactedBlocks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_compactor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactedBlocks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactedBlocks) ProtoMessage() {}

func (x *CompactedBlocks) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_compactor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactedBlocks.ProtoReflect.Descriptor instead.
func (*CompactedBlocks) Descriptor() ([]byte, []int) {
	return file_metastore_v1_compactor_proto_rawDescGZIP(), []int{7}
}

func (x *CompactedBlocks) GetSourceBlocks() *BlockList {
	if x != nil {
		return x.SourceBlocks
	}
	return nil
}

func (x *CompactedBlocks) GetNewBlocks() []*BlockMeta {
	if x != nil {
		return x.NewBlocks
	}
	return nil
}

var File_metastore_v1_compactor_proto protoreflect.FileDescriptor

var file_metastore_v1_compactor_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x6d, 0x65,
	0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x19, 0x50, 0x6f, 0x6c, 0x6c, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x61, 0x70, 0x61,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6a, 0x6f, 0x62, 0x43,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x22, 0xab, 0x01, 0x0a, 0x1a, 0x50, 0x6f, 0x6c, 0x6c,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x0e, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x47, 0x0a, 0x0b,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xdb, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x68, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x38, 0x0a, 0x0a, 0x74, 0x6f, 0x6d,
	0x62, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6d,
	0x62, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x0a, 0x74, 0x6f, 0x6d, 0x62, 0x73, 0x74, 0x6f,
	0x6e, 0x65, 0x73, 0x22, 0x43, 0x0a, 0x0a, 0x54, 0x6f, 0x6d, 0x62, 0x73, 0x74, 0x6f, 0x6e, 0x65,
	0x73, 0x12, 0x35, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x6f, 0x6d, 0x62, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73,
	0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x96, 0x01, 0x0a, 0x0f, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x54, 0x6f, 0x6d, 0x62, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x29,
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x22, 0x6d, 0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a,
	0x6f, 0x62, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74,
	0x22, 0xca, 0x01, 0x0a, 0x19, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a,
	0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x48, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x65, 0x64,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x63, 0x74, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x0f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x87, 0x01,
	0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x12, 0x3c, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12,
	0x36, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x09, 0x6e, 0x65,
	0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x2a, 0x7a, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21,
	0x0a, 0x1d, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45,
	0x53, 0x53, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x02, 0x32, 0x7e, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x12, 0x50, 0x6f, 0x6c, 0x6c,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x27,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0xbb, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x2f, 0x70,
	0x79, 0x72, 0x6f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x4d, 0x65, 0x74, 0x61, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metastore_v1_compactor_proto_rawDescOnce sync.Once
	file_metastore_v1_compactor_proto_rawDescData = file_metastore_v1_compactor_proto_rawDesc
)

func file_metastore_v1_compactor_proto_rawDescGZIP() []byte {
	file_metastore_v1_compactor_proto_rawDescOnce.Do(func() {
		file_metastore_v1_compactor_proto_rawDescData = protoimpl.X.CompressGZIP(file_metastore_v1_compactor_proto_rawDescData)
	})
	return file_metastore_v1_compactor_proto_rawDescData
}

var file_metastore_v1_compactor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_metastore_v1_compactor_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_metastore_v1_compactor_proto_goTypes = []any{
	(CompactionJobStatus)(0),           // 0: metastore.v1.CompactionJobStatus
	(*PollCompactionJobsRequest)(nil),  // 1: metastore.v1.PollCompactionJobsRequest
	(*PollCompactionJobsResponse)(nil), // 2: metastore.v1.PollCompactionJobsResponse
	(*CompactionJob)(nil),              // 3: metastore.v1.CompactionJob
	(*Tombstones)(nil),                 // 4: metastore.v1.Tombstones
	(*BlockTombstones)(nil),            // 5: metastore.v1.BlockTombstones
	(*CompactionJobAssignment)(nil),    // 6: metastore.v1.CompactionJobAssignment
	(*CompactionJobStatusUpdate)(nil),  // 7: metastore.v1.CompactionJobStatusUpdate
	(*CompactedBlocks)(nil),            // 8: metastore.v1.CompactedBlocks
	(*BlockList)(nil),                  // 9: metastore.v1.BlockList
	(*BlockMeta)(nil),                  // 10: metastore.v1.BlockMeta
}
var file_metastore_v1_compactor_proto_depIdxs = []int32{
	7,  // 0: metastore.v1.PollCompactionJobsRequest.status_updates:type_name -> metastore.v1.CompactionJobStatusUpdate
	3,  // 1: metastore.v1.PollCompactionJobsResponse.compaction_jobs:type_name -> metastore.v1.CompactionJob
	6,  // 2: metastore.v1.PollCompactionJobsResponse.assignments:type_name -> metastore.v1.CompactionJobAssignment
	4,  // 3: metastore.v1.CompactionJob.tombstones:type_name -> metastore.v1.Tombstones
	5,  // 4: metastore.v1.Tombstones.blocks:type_name -> metastore.v1.BlockTombstones
	0,  // 5: metastore.v1.CompactionJobStatusUpdate.status:type_name -> metastore.v1.CompactionJobStatus
	8,  // 6: metastore.v1.CompactionJobStatusUpdate.compacted_blocks:type_name -> metastore.v1.CompactedBlocks
	9,  // 7: metastore.v1.CompactedBlocks.source_blocks:type_name -> metastore.v1.BlockList
	10, // 8: metastore.v1.CompactedBlocks.new_blocks:type_name -> metastore.v1.BlockMeta
	1,  // 9: metastore.v1.CompactionService.PollCompactionJobs:input_type -> metastore.v1.PollCompactionJobsRequest
	2,  // 10: metastore.v1.CompactionService.PollCompactionJobs:output_type -> metastore.v1.PollCompactionJobsResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_metastore_v1_compactor_proto_init() }
func file_metastore_v1_compactor_proto_init() {
	if File_metastore_v1_compactor_proto != nil {
		return
	}
	file_metastore_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_metastore_v1_compactor_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PollCompactionJobsRequest); i {
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
		file_metastore_v1_compactor_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PollCompactionJobsResponse); i {
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
		file_metastore_v1_compactor_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CompactionJob); i {
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
		file_metastore_v1_compactor_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Tombstones); i {
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
		file_metastore_v1_compactor_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*BlockTombstones); i {
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
		file_metastore_v1_compactor_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CompactionJobAssignment); i {
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
		file_metastore_v1_compactor_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CompactionJobStatusUpdate); i {
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
		file_metastore_v1_compactor_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CompactedBlocks); i {
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
			RawDescriptor: file_metastore_v1_compactor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metastore_v1_compactor_proto_goTypes,
		DependencyIndexes: file_metastore_v1_compactor_proto_depIdxs,
		EnumInfos:         file_metastore_v1_compactor_proto_enumTypes,
		MessageInfos:      file_metastore_v1_compactor_proto_msgTypes,
	}.Build()
	File_metastore_v1_compactor_proto = out.File
	file_metastore_v1_compactor_proto_rawDesc = nil
	file_metastore_v1_compactor_proto_goTypes = nil
	file_metastore_v1_compactor_proto_depIdxs = nil
}