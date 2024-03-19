// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: adhocprofiles/v1/adhocprofiles.proto

package adhocprofilesv1

import (
	_ "github.com/grafana/pyroscope/api/gen/proto/go/types/v1"
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

type AdHocProfilesUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is typically the file name and it serves as a human readable name for the profile.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// This is the profile encoded in base64. The supported formats are pprof, json, collapsed and perf-script.
	Profile string `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	// Max nodes can be used to truncate the response.
	MaxNodes *int64 `protobuf:"varint,3,opt,name=max_nodes,json=maxNodes,proto3,oneof" json:"max_nodes,omitempty"`
}

func (x *AdHocProfilesUploadRequest) Reset() {
	*x = AdHocProfilesUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdHocProfilesUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdHocProfilesUploadRequest) ProtoMessage() {}

func (x *AdHocProfilesUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdHocProfilesUploadRequest.ProtoReflect.Descriptor instead.
func (*AdHocProfilesUploadRequest) Descriptor() ([]byte, []int) {
	return file_adhocprofiles_v1_adhocprofiles_proto_rawDescGZIP(), []int{0}
}

func (x *AdHocProfilesUploadRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdHocProfilesUploadRequest) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *AdHocProfilesUploadRequest) GetMaxNodes() int64 {
	if x != nil && x.MaxNodes != nil {
		return *x.MaxNodes
	}
	return 0
}

type AdHocProfilesGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the profile.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The desired profile type (e.g., cpu, samples) for the returned flame graph. If omitted the first profile is returned.
	ProfileType *string `protobuf:"bytes,2,opt,name=profile_type,json=profileType,proto3,oneof" json:"profile_type,omitempty"`
	// Max nodes can be used to truncate the response.
	MaxNodes *int64 `protobuf:"varint,3,opt,name=max_nodes,json=maxNodes,proto3,oneof" json:"max_nodes,omitempty"`
}

func (x *AdHocProfilesGetRequest) Reset() {
	*x = AdHocProfilesGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdHocProfilesGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdHocProfilesGetRequest) ProtoMessage() {}

func (x *AdHocProfilesGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdHocProfilesGetRequest.ProtoReflect.Descriptor instead.
func (*AdHocProfilesGetRequest) Descriptor() ([]byte, []int) {
	return file_adhocprofiles_v1_adhocprofiles_proto_rawDescGZIP(), []int{1}
}

func (x *AdHocProfilesGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AdHocProfilesGetRequest) GetProfileType() string {
	if x != nil && x.ProfileType != nil {
		return *x.ProfileType
	}
	return ""
}

func (x *AdHocProfilesGetRequest) GetMaxNodes() int64 {
	if x != nil && x.MaxNodes != nil {
		return *x.MaxNodes
	}
	return 0
}

type AdHocProfilesGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// timestamp in milliseconds
	UploadedAt  int64  `protobuf:"varint,3,opt,name=uploaded_at,json=uploadedAt,proto3" json:"uploaded_at,omitempty"`
	ProfileType string `protobuf:"bytes,4,opt,name=profile_type,json=profileType,proto3" json:"profile_type,omitempty"`
	// Some profiles formats (like pprof) can contain multiple profile (sample) types inside. One of these can be passed
	// in the Get request using the profile_type field.
	ProfileTypes       []string `protobuf:"bytes,5,rep,name=profile_types,json=profileTypes,proto3" json:"profile_types,omitempty"`
	FlamebearerProfile string   `protobuf:"bytes,6,opt,name=flamebearer_profile,json=flamebearerProfile,proto3" json:"flamebearer_profile,omitempty"`
}

func (x *AdHocProfilesGetResponse) Reset() {
	*x = AdHocProfilesGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdHocProfilesGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdHocProfilesGetResponse) ProtoMessage() {}

func (x *AdHocProfilesGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdHocProfilesGetResponse.ProtoReflect.Descriptor instead.
func (*AdHocProfilesGetResponse) Descriptor() ([]byte, []int) {
	return file_adhocprofiles_v1_adhocprofiles_proto_rawDescGZIP(), []int{2}
}

func (x *AdHocProfilesGetResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AdHocProfilesGetResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdHocProfilesGetResponse) GetUploadedAt() int64 {
	if x != nil {
		return x.UploadedAt
	}
	return 0
}

func (x *AdHocProfilesGetResponse) GetProfileType() string {
	if x != nil {
		return x.ProfileType
	}
	return ""
}

func (x *AdHocProfilesGetResponse) GetProfileTypes() []string {
	if x != nil {
		return x.ProfileTypes
	}
	return nil
}

func (x *AdHocProfilesGetResponse) GetFlamebearerProfile() string {
	if x != nil {
		return x.FlamebearerProfile
	}
	return ""
}

type AdHocProfilesListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdHocProfilesListRequest) Reset() {
	*x = AdHocProfilesListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdHocProfilesListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdHocProfilesListRequest) ProtoMessage() {}

func (x *AdHocProfilesListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdHocProfilesListRequest.ProtoReflect.Descriptor instead.
func (*AdHocProfilesListRequest) Descriptor() ([]byte, []int) {
	return file_adhocprofiles_v1_adhocprofiles_proto_rawDescGZIP(), []int{3}
}

type AdHocProfilesListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profiles []*AdHocProfilesProfileMetadata `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles,omitempty"`
}

func (x *AdHocProfilesListResponse) Reset() {
	*x = AdHocProfilesListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdHocProfilesListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdHocProfilesListResponse) ProtoMessage() {}

func (x *AdHocProfilesListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdHocProfilesListResponse.ProtoReflect.Descriptor instead.
func (*AdHocProfilesListResponse) Descriptor() ([]byte, []int) {
	return file_adhocprofiles_v1_adhocprofiles_proto_rawDescGZIP(), []int{4}
}

func (x *AdHocProfilesListResponse) GetProfiles() []*AdHocProfilesProfileMetadata {
	if x != nil {
		return x.Profiles
	}
	return nil
}

type AdHocProfilesProfileMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// timestamp in milliseconds
	UploadedAt int64 `protobuf:"varint,3,opt,name=uploaded_at,json=uploadedAt,proto3" json:"uploaded_at,omitempty"`
}

func (x *AdHocProfilesProfileMetadata) Reset() {
	*x = AdHocProfilesProfileMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdHocProfilesProfileMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdHocProfilesProfileMetadata) ProtoMessage() {}

func (x *AdHocProfilesProfileMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdHocProfilesProfileMetadata.ProtoReflect.Descriptor instead.
func (*AdHocProfilesProfileMetadata) Descriptor() ([]byte, []int) {
	return file_adhocprofiles_v1_adhocprofiles_proto_rawDescGZIP(), []int{5}
}

func (x *AdHocProfilesProfileMetadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AdHocProfilesProfileMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdHocProfilesProfileMetadata) GetUploadedAt() int64 {
	if x != nil {
		return x.UploadedAt
	}
	return 0
}

var File_adhocprofiles_v1_adhocprofiles_proto protoreflect.FileDescriptor

var file_adhocprofiles_v1_adhocprofiles_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x64, 0x68, 0x6f, 0x63, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x64, 0x68, 0x6f, 0x63, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x64, 0x68, 0x6f, 0x63, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a,
	0x0a, 0x1a, 0x41, 0x64, 0x48, 0x6f, 0x63, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x09, 0x6d, 0x61,
	0x78, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x08, 0x6d, 0x61, 0x78, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x17, 0x41,
	0x64, 0x48, 0x6f, 0x63, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20,
	0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x01, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x88, 0x01, 0x01,
	0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22,
	0xd8, 0x01, 0x0a, 0x18, 0x41, 0x64, 0x48, 0x6f, 0x63, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x66, 0x6c, 0x61,
	0x6d, 0x65, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x66, 0x6c, 0x61, 0x6d, 0x65, 0x62, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x1a, 0x0a, 0x18, 0x41, 0x64,
	0x48, 0x6f, 0x63, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x67, 0x0a, 0x19, 0x41, 0x64, 0x48, 0x6f, 0x63, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x64, 0x68, 0x6f, 0x63, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x48, 0x6f, 0x63, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22,
	0x63, 0x0a, 0x1c, 0x41, 0x64, 0x48, 0x6f, 0x63, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x64, 0x41, 0x74, 0x32, 0xbe, 0x02, 0x0a, 0x13, 0x41, 0x64, 0x48, 0x6f, 0x63, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x06,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2c, 0x2e, 0x61, 0x64, 0x68, 0x6f, 0x63, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x48, 0x6f, 0x63, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x64, 0x68, 0x6f, 0x63, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x48, 0x6f, 0x63, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x29, 0x2e, 0x61, 0x64, 0x68, 0x6f,
	0x63, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x48,
	0x6f, 0x63, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x64, 0x68, 0x6f, 0x63, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x48, 0x6f, 0x63, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x61, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x2e, 0x61, 0x64, 0x68,
	0x6f, 0x63, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x48, 0x6f, 0x63, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x64, 0x68, 0x6f, 0x63, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x48, 0x6f, 0x63, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xdb, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x64,
	0x68, 0x6f, 0x63, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x12,
	0x41, 0x64, 0x68, 0x6f, 0x63, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x79, 0x72, 0x6f, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x64, 0x68, 0x6f, 0x63, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x64, 0x68, 0x6f, 0x63, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x10, 0x41, 0x64, 0x68,
	0x6f, 0x63, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10,
	0x41, 0x64, 0x68, 0x6f, 0x63, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x1c, 0x41, 0x64, 0x68, 0x6f, 0x63, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x11, 0x41, 0x64, 0x68, 0x6f, 0x63, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_adhocprofiles_v1_adhocprofiles_proto_rawDescOnce sync.Once
	file_adhocprofiles_v1_adhocprofiles_proto_rawDescData = file_adhocprofiles_v1_adhocprofiles_proto_rawDesc
)

func file_adhocprofiles_v1_adhocprofiles_proto_rawDescGZIP() []byte {
	file_adhocprofiles_v1_adhocprofiles_proto_rawDescOnce.Do(func() {
		file_adhocprofiles_v1_adhocprofiles_proto_rawDescData = protoimpl.X.CompressGZIP(file_adhocprofiles_v1_adhocprofiles_proto_rawDescData)
	})
	return file_adhocprofiles_v1_adhocprofiles_proto_rawDescData
}

var file_adhocprofiles_v1_adhocprofiles_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_adhocprofiles_v1_adhocprofiles_proto_goTypes = []interface{}{
	(*AdHocProfilesUploadRequest)(nil),   // 0: adhocprofiles.v1.AdHocProfilesUploadRequest
	(*AdHocProfilesGetRequest)(nil),      // 1: adhocprofiles.v1.AdHocProfilesGetRequest
	(*AdHocProfilesGetResponse)(nil),     // 2: adhocprofiles.v1.AdHocProfilesGetResponse
	(*AdHocProfilesListRequest)(nil),     // 3: adhocprofiles.v1.AdHocProfilesListRequest
	(*AdHocProfilesListResponse)(nil),    // 4: adhocprofiles.v1.AdHocProfilesListResponse
	(*AdHocProfilesProfileMetadata)(nil), // 5: adhocprofiles.v1.AdHocProfilesProfileMetadata
}
var file_adhocprofiles_v1_adhocprofiles_proto_depIdxs = []int32{
	5, // 0: adhocprofiles.v1.AdHocProfilesListResponse.profiles:type_name -> adhocprofiles.v1.AdHocProfilesProfileMetadata
	0, // 1: adhocprofiles.v1.AdHocProfileService.Upload:input_type -> adhocprofiles.v1.AdHocProfilesUploadRequest
	1, // 2: adhocprofiles.v1.AdHocProfileService.Get:input_type -> adhocprofiles.v1.AdHocProfilesGetRequest
	3, // 3: adhocprofiles.v1.AdHocProfileService.List:input_type -> adhocprofiles.v1.AdHocProfilesListRequest
	2, // 4: adhocprofiles.v1.AdHocProfileService.Upload:output_type -> adhocprofiles.v1.AdHocProfilesGetResponse
	2, // 5: adhocprofiles.v1.AdHocProfileService.Get:output_type -> adhocprofiles.v1.AdHocProfilesGetResponse
	4, // 6: adhocprofiles.v1.AdHocProfileService.List:output_type -> adhocprofiles.v1.AdHocProfilesListResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_adhocprofiles_v1_adhocprofiles_proto_init() }
func file_adhocprofiles_v1_adhocprofiles_proto_init() {
	if File_adhocprofiles_v1_adhocprofiles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdHocProfilesUploadRequest); i {
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
		file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdHocProfilesGetRequest); i {
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
		file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdHocProfilesGetResponse); i {
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
		file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdHocProfilesListRequest); i {
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
		file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdHocProfilesListResponse); i {
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
		file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdHocProfilesProfileMetadata); i {
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
	file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_adhocprofiles_v1_adhocprofiles_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_adhocprofiles_v1_adhocprofiles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_adhocprofiles_v1_adhocprofiles_proto_goTypes,
		DependencyIndexes: file_adhocprofiles_v1_adhocprofiles_proto_depIdxs,
		MessageInfos:      file_adhocprofiles_v1_adhocprofiles_proto_msgTypes,
	}.Build()
	File_adhocprofiles_v1_adhocprofiles_proto = out.File
	file_adhocprofiles_v1_adhocprofiles_proto_rawDesc = nil
	file_adhocprofiles_v1_adhocprofiles_proto_goTypes = nil
	file_adhocprofiles_v1_adhocprofiles_proto_depIdxs = nil
}