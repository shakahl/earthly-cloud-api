// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: runner.proto

package runner

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

type GitInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// https://docs.github.com/en/rest/checks/runs?apiVersion=2022-11-28
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Repo  string `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
	Hash  string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *GitInfo) Reset() {
	*x = GitInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitInfo) ProtoMessage() {}

func (x *GitInfo) ProtoReflect() protoreflect.Message {
	mi := &file_runner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitInfo.ProtoReflect.Descriptor instead.
func (*GitInfo) Descriptor() ([]byte, []int) {
	return file_runner_proto_rawDescGZIP(), []int{0}
}

func (x *GitInfo) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GitInfo) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *GitInfo) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type BuildRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url          string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Path         string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Tag          string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	Target       string   `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty"`
	RepoToken    string   `protobuf:"bytes,5,opt,name=repo_token,json=repoToken,proto3" json:"repo_token,omitempty"`
	BuildId      string   `protobuf:"bytes,6,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	PipelineId   string   `protobuf:"bytes,7,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id,omitempty"`
	UserId       string   `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrgId        string   `protobuf:"bytes,9,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	ProjectId    string   `protobuf:"bytes,10,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	PipelineName string   `protobuf:"bytes,11,opt,name=pipeline_name,json=pipelineName,proto3" json:"pipeline_name,omitempty"`
	GitInfo      *GitInfo `protobuf:"bytes,12,opt,name=git_info,json=gitInfo,proto3" json:"git_info,omitempty"`
}

func (x *BuildRequest) Reset() {
	*x = BuildRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildRequest) ProtoMessage() {}

func (x *BuildRequest) ProtoReflect() protoreflect.Message {
	mi := &file_runner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildRequest.ProtoReflect.Descriptor instead.
func (*BuildRequest) Descriptor() ([]byte, []int) {
	return file_runner_proto_rawDescGZIP(), []int{1}
}

func (x *BuildRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *BuildRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *BuildRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *BuildRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *BuildRequest) GetRepoToken() string {
	if x != nil {
		return x.RepoToken
	}
	return ""
}

func (x *BuildRequest) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

func (x *BuildRequest) GetPipelineId() string {
	if x != nil {
		return x.PipelineId
	}
	return ""
}

func (x *BuildRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BuildRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *BuildRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *BuildRequest) GetPipelineName() string {
	if x != nil {
		return x.PipelineName
	}
	return ""
}

func (x *BuildRequest) GetGitInfo() *GitInfo {
	if x != nil {
		return x.GitInfo
	}
	return nil
}

type BuildResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BuildResponse) Reset() {
	*x = BuildResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildResponse) ProtoMessage() {}

func (x *BuildResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildResponse.ProtoReflect.Descriptor instead.
func (*BuildResponse) Descriptor() ([]byte, []int) {
	return file_runner_proto_rawDescGZIP(), []int{2}
}

var File_runner_proto protoreflect.FileDescriptor

var file_runner_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65,
	0x72, 0x22, 0x47, 0x0a, 0x07, 0x47, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0xe4, 0x02, 0x0a, 0x0c, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x74, 0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x70, 0x6f, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x70, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x67, 0x69,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72,
	0x2e, 0x47, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x67, 0x69, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x56, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x05,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_runner_proto_rawDescOnce sync.Once
	file_runner_proto_rawDescData = file_runner_proto_rawDesc
)

func file_runner_proto_rawDescGZIP() []byte {
	file_runner_proto_rawDescOnce.Do(func() {
		file_runner_proto_rawDescData = protoimpl.X.CompressGZIP(file_runner_proto_rawDescData)
	})
	return file_runner_proto_rawDescData
}

var file_runner_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_runner_proto_goTypes = []interface{}{
	(*GitInfo)(nil),       // 0: api.public.runner.GitInfo
	(*BuildRequest)(nil),  // 1: api.public.runner.BuildRequest
	(*BuildResponse)(nil), // 2: api.public.runner.BuildResponse
}
var file_runner_proto_depIdxs = []int32{
	0, // 0: api.public.runner.BuildRequest.git_info:type_name -> api.public.runner.GitInfo
	1, // 1: api.public.runner.Runner.Build:input_type -> api.public.runner.BuildRequest
	2, // 2: api.public.runner.Runner.Build:output_type -> api.public.runner.BuildResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_runner_proto_init() }
func file_runner_proto_init() {
	if File_runner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_runner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitInfo); i {
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
		file_runner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildRequest); i {
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
		file_runner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildResponse); i {
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
			RawDescriptor: file_runner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_runner_proto_goTypes,
		DependencyIndexes: file_runner_proto_depIdxs,
		MessageInfos:      file_runner_proto_msgTypes,
	}.Build()
	File_runner_proto = out.File
	file_runner_proto_rawDesc = nil
	file_runner_proto_goTypes = nil
	file_runner_proto_depIdxs = nil
}
