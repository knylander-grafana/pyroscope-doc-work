// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: metastore/v1/types.proto

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

// BlockMeta is a metadata entry that describes the block's contents. A block
// is a collection of datasets that share certain properties, such as shard ID,
// compaction level, tenant ID, time range, creation time, and more.
//
// The block content's format denotes the binary format of the datasets and the
// metadata entry (to address logical dependencies). Each dataset has its own
// table of contents that lists the sections within the dataset. Each dataset
// has its own set of attributes (labels) that describe its specific contents.
type BlockMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FormatVersion uint32 `protobuf:"varint,1,opt,name=format_version,json=formatVersion,proto3" json:"format_version,omitempty"`
	// Block ID is a unique identifier for the block.
	// This is the only field that is not included into
	// the string table.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// If empty, datasets belong to distinct tenants.
	Tenant          int32      `protobuf:"varint,3,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Shard           uint32     `protobuf:"varint,4,opt,name=shard,proto3" json:"shard,omitempty"`
	CompactionLevel uint32     `protobuf:"varint,5,opt,name=compaction_level,json=compactionLevel,proto3" json:"compaction_level,omitempty"`
	MinTime         int64      `protobuf:"varint,6,opt,name=min_time,json=minTime,proto3" json:"min_time,omitempty"`
	MaxTime         int64      `protobuf:"varint,7,opt,name=max_time,json=maxTime,proto3" json:"max_time,omitempty"`
	CreatedBy       int32      `protobuf:"varint,8,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Size            uint64     `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	Datasets        []*Dataset `protobuf:"bytes,10,rep,name=datasets,proto3" json:"datasets,omitempty"`
	// String table contains strings of the block.
	// By convention, the first string is always an empty string.
	StringTable []string `protobuf:"bytes,11,rep,name=string_table,json=stringTable,proto3" json:"string_table,omitempty"`
}

func (x *BlockMeta) Reset() {
	*x = BlockMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockMeta) ProtoMessage() {}

func (x *BlockMeta) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockMeta.ProtoReflect.Descriptor instead.
func (*BlockMeta) Descriptor() ([]byte, []int) {
	return file_metastore_v1_types_proto_rawDescGZIP(), []int{0}
}

func (x *BlockMeta) GetFormatVersion() uint32 {
	if x != nil {
		return x.FormatVersion
	}
	return 0
}

func (x *BlockMeta) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BlockMeta) GetTenant() int32 {
	if x != nil {
		return x.Tenant
	}
	return 0
}

func (x *BlockMeta) GetShard() uint32 {
	if x != nil {
		return x.Shard
	}
	return 0
}

func (x *BlockMeta) GetCompactionLevel() uint32 {
	if x != nil {
		return x.CompactionLevel
	}
	return 0
}

func (x *BlockMeta) GetMinTime() int64 {
	if x != nil {
		return x.MinTime
	}
	return 0
}

func (x *BlockMeta) GetMaxTime() int64 {
	if x != nil {
		return x.MaxTime
	}
	return 0
}

func (x *BlockMeta) GetCreatedBy() int32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *BlockMeta) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *BlockMeta) GetDatasets() []*Dataset {
	if x != nil {
		return x.Datasets
	}
	return nil
}

func (x *BlockMeta) GetStringTable() []string {
	if x != nil {
		return x.StringTable
	}
	return nil
}

type Dataset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant  int32 `protobuf:"varint,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Name    int32 `protobuf:"varint,2,opt,name=name,proto3" json:"name,omitempty"`
	MinTime int64 `protobuf:"varint,3,opt,name=min_time,json=minTime,proto3" json:"min_time,omitempty"`
	MaxTime int64 `protobuf:"varint,4,opt,name=max_time,json=maxTime,proto3" json:"max_time,omitempty"`
	// Table of contents lists data sections within the tenant
	// service region. The offsets are absolute.
	//
	// The interpretation of the table of contents is specific
	// to the metadata format version. By default, the sections are:
	//   - 0: profiles.parquet
	//   - 1: index.tsdb
	//   - 2: symbols.symdb
	TableOfContents []uint64 `protobuf:"varint,5,rep,packed,name=table_of_contents,json=tableOfContents,proto3" json:"table_of_contents,omitempty"`
	// Size of the dataset in bytes.
	Size uint64 `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	// Length prefixed label key-value pairs.
	//
	// Multiple label sets can be associated with a dataset to denote relationships
	// across multiple dimensions. For example, each dataset currently stores data
	// for multiple profile types:
	//   - service_name=A, profile_type=cpu
	//   - service_name=A, profile_type=memory
	//
	// Labels are primarily used to filter datasets based on their attributes.
	// For instance, labels can be used to select datasets containing a specific
	// service.
	//
	// The set of attributes is extensible and can grow over time. For example, a
	// namespace attribute could be added to datasets:
	//   - service_name=A, profile_type=cpu
	//   - service_name=A, profile_type=memory
	//   - service_name=B, namespace=N, profile_type=cpu
	//   - service_name=B, namespace=N, profile_type=memory
	//   - service_name=C, namespace=N, profile_type=cpu
	//   - service_name=C, namespace=N, profile_type=memory
	//
	// This organization enables querying datasets by namespace without accessing
	// the block contents, which significantly improves performance.
	//
	// Metadata labels are not required to be included in the block's TSDB index
	// and may be orthogonal to the data dimensions. Generally, attributes serve
	// two primary purposes:
	//   - To create data scopes that span multiple service, reducing the need to
	//     scan the entire set of block satisfying the query expression, i.e.,
	//     the time range and tenant ID.
	//   - To provide additional information about datasets without altering the
	//     storage schema or access methods.
	//
	// For example, this approach can support cost attribution or similar breakdown
	// analyses. It can also handle data dependencies (e.g., links to external data)
	// using labels.
	//
	// The cardinality of the labels is expected to remain relatively low (fewer
	// than a million unique combinations globally). However, this depends on the
	// metadata storage system.
	//
	// Metadata labels are represented as a slice of `int32` values that refer to
	// strings in the metadata entry's string table. The slice is a sequence of
	// length-prefixed key-value (KV) pairs:
	//
	// len(2) | k1 | v1 | k2 | v2 | len(3) | k1 | v3 | k2 | v4 | k3 | v5
	//
	// The order of KV pairs is not defined. The format is optimized for indexing
	// rather than querying, and it is not intended to be the most space-efficient
	// representation. Since entries are supposed to be indexed, the redundancy of
	// denormalized relationships is not a concern.
	Labels []int32 `protobuf:"varint,8,rep,packed,name=labels,proto3" json:"labels,omitempty"`
}

func (x *Dataset) Reset() {
	*x = Dataset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dataset) ProtoMessage() {}

func (x *Dataset) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dataset.ProtoReflect.Descriptor instead.
func (*Dataset) Descriptor() ([]byte, []int) {
	return file_metastore_v1_types_proto_rawDescGZIP(), []int{1}
}

func (x *Dataset) GetTenant() int32 {
	if x != nil {
		return x.Tenant
	}
	return 0
}

func (x *Dataset) GetName() int32 {
	if x != nil {
		return x.Name
	}
	return 0
}

func (x *Dataset) GetMinTime() int64 {
	if x != nil {
		return x.MinTime
	}
	return 0
}

func (x *Dataset) GetMaxTime() int64 {
	if x != nil {
		return x.MaxTime
	}
	return 0
}

func (x *Dataset) GetTableOfContents() []uint64 {
	if x != nil {
		return x.TableOfContents
	}
	return nil
}

func (x *Dataset) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Dataset) GetLabels() []int32 {
	if x != nil {
		return x.Labels
	}
	return nil
}

type BlockList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant string   `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Shard  uint32   `protobuf:"varint,2,opt,name=shard,proto3" json:"shard,omitempty"`
	Blocks []string `protobuf:"bytes,3,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *BlockList) Reset() {
	*x = BlockList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockList) ProtoMessage() {}

func (x *BlockList) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockList.ProtoReflect.Descriptor instead.
func (*BlockList) Descriptor() ([]byte, []int) {
	return file_metastore_v1_types_proto_rawDescGZIP(), []int{2}
}

func (x *BlockList) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *BlockList) GetShard() uint32 {
	if x != nil {
		return x.Shard
	}
	return 0
}

func (x *BlockList) GetBlocks() []string {
	if x != nil {
		return x.Blocks
	}
	return nil
}

var File_metastore_v1_types_proto protoreflect.FileDescriptor

var file_metastore_v1_types_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x74, 0x61,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xda, 0x02, 0x0a, 0x09, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x31, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xc9, 0x01, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x6d, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6f, 0x66, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0f,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x66, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x4a, 0x04, 0x08, 0x07, 0x10,
	0x08, 0x22, 0x51, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x42, 0xb7, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x79, 0x72, 0x6f,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0d, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metastore_v1_types_proto_rawDescOnce sync.Once
	file_metastore_v1_types_proto_rawDescData = file_metastore_v1_types_proto_rawDesc
)

func file_metastore_v1_types_proto_rawDescGZIP() []byte {
	file_metastore_v1_types_proto_rawDescOnce.Do(func() {
		file_metastore_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_metastore_v1_types_proto_rawDescData)
	})
	return file_metastore_v1_types_proto_rawDescData
}

var file_metastore_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_metastore_v1_types_proto_goTypes = []any{
	(*BlockMeta)(nil), // 0: metastore.v1.BlockMeta
	(*Dataset)(nil),   // 1: metastore.v1.Dataset
	(*BlockList)(nil), // 2: metastore.v1.BlockList
}
var file_metastore_v1_types_proto_depIdxs = []int32{
	1, // 0: metastore.v1.BlockMeta.datasets:type_name -> metastore.v1.Dataset
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_metastore_v1_types_proto_init() }
func file_metastore_v1_types_proto_init() {
	if File_metastore_v1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metastore_v1_types_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BlockMeta); i {
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
		file_metastore_v1_types_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Dataset); i {
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
		file_metastore_v1_types_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BlockList); i {
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
			RawDescriptor: file_metastore_v1_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_metastore_v1_types_proto_goTypes,
		DependencyIndexes: file_metastore_v1_types_proto_depIdxs,
		MessageInfos:      file_metastore_v1_types_proto_msgTypes,
	}.Build()
	File_metastore_v1_types_proto = out.File
	file_metastore_v1_types_proto_rawDesc = nil
	file_metastore_v1_types_proto_goTypes = nil
	file_metastore_v1_types_proto_depIdxs = nil
}
