// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: store/branch.proto

package store

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

type BranchSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema         string                  `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	Metadata       *DatabaseSchemaMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	DatabaseConfig *DatabaseConfig         `protobuf:"bytes,3,opt,name=database_config,json=databaseConfig,proto3" json:"database_config,omitempty"`
}

func (x *BranchSnapshot) Reset() {
	*x = BranchSnapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_branch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchSnapshot) ProtoMessage() {}

func (x *BranchSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_store_branch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchSnapshot.ProtoReflect.Descriptor instead.
func (*BranchSnapshot) Descriptor() ([]byte, []int) {
	return file_store_branch_proto_rawDescGZIP(), []int{0}
}

func (x *BranchSnapshot) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *BranchSnapshot) GetMetadata() *DatabaseSchemaMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *BranchSnapshot) GetDatabaseConfig() *DatabaseConfig {
	if x != nil {
		return x.DatabaseConfig
	}
	return nil
}

type BranchConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of source database.
	// Optional.
	// Example: instances/instance-id/databases/database-name.
	SourceDatabase int64 `protobuf:"varint,1,opt,name=source_database,json=sourceDatabase,proto3" json:"source_database,omitempty"`
	// The name of the source branch.
	// Optional.
	// Example: projects/project-id/branches/branch-id.
	SourceBranch int64 `protobuf:"varint,2,opt,name=source_branch,json=sourceBranch,proto3" json:"source_branch,omitempty"`
}

func (x *BranchConfig) Reset() {
	*x = BranchConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_branch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchConfig) ProtoMessage() {}

func (x *BranchConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_branch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchConfig.ProtoReflect.Descriptor instead.
func (*BranchConfig) Descriptor() ([]byte, []int) {
	return file_store_branch_proto_rawDescGZIP(), []int{1}
}

func (x *BranchConfig) GetSourceDatabase() int64 {
	if x != nil {
		return x.SourceDatabase
	}
	return 0
}

func (x *BranchConfig) GetSourceBranch() int64 {
	if x != nil {
		return x.SourceBranch
	}
	return 0
}

var File_store_branch_proto protoreflect.FileDescriptor

var file_store_branch_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x1a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x0e, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x42, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x0f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x5c, 0x0a, 0x0c, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_branch_proto_rawDescOnce sync.Once
	file_store_branch_proto_rawDescData = file_store_branch_proto_rawDesc
)

func file_store_branch_proto_rawDescGZIP() []byte {
	file_store_branch_proto_rawDescOnce.Do(func() {
		file_store_branch_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_branch_proto_rawDescData)
	})
	return file_store_branch_proto_rawDescData
}

var file_store_branch_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_store_branch_proto_goTypes = []interface{}{
	(*BranchSnapshot)(nil),         // 0: bytebase.store.BranchSnapshot
	(*BranchConfig)(nil),           // 1: bytebase.store.BranchConfig
	(*DatabaseSchemaMetadata)(nil), // 2: bytebase.store.DatabaseSchemaMetadata
	(*DatabaseConfig)(nil),         // 3: bytebase.store.DatabaseConfig
}
var file_store_branch_proto_depIdxs = []int32{
	2, // 0: bytebase.store.BranchSnapshot.metadata:type_name -> bytebase.store.DatabaseSchemaMetadata
	3, // 1: bytebase.store.BranchSnapshot.database_config:type_name -> bytebase.store.DatabaseConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_store_branch_proto_init() }
func file_store_branch_proto_init() {
	if File_store_branch_proto != nil {
		return
	}
	file_store_database_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_store_branch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BranchSnapshot); i {
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
		file_store_branch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BranchConfig); i {
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
			RawDescriptor: file_store_branch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_branch_proto_goTypes,
		DependencyIndexes: file_store_branch_proto_depIdxs,
		MessageInfos:      file_store_branch_proto_msgTypes,
	}.Build()
	File_store_branch_proto = out.File
	file_store_branch_proto_rawDesc = nil
	file_store_branch_proto_goTypes = nil
	file_store_branch_proto_depIdxs = nil
}
