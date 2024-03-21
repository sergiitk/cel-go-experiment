// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.5
// source: xds/type/v4/cel.proto

package v4

import (
	expr "cel.wtf/expr"
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

type CelExpression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CelExprChecked *expr.CheckedExpr `protobuf:"bytes,1,opt,name=cel_expr_checked,json=celExprChecked,proto3" json:"cel_expr_checked,omitempty"`
}

func (x *CelExpression) Reset() {
	*x = CelExpression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_type_v4_cel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CelExpression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CelExpression) ProtoMessage() {}

func (x *CelExpression) ProtoReflect() protoreflect.Message {
	mi := &file_xds_type_v4_cel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CelExpression.ProtoReflect.Descriptor instead.
func (*CelExpression) Descriptor() ([]byte, []int) {
	return file_xds_type_v4_cel_proto_rawDescGZIP(), []int{0}
}

func (x *CelExpression) GetCelExprChecked() *expr.CheckedExpr {
	if x != nil {
		return x.CelExprChecked
	}
	return nil
}

var File_xds_type_v4_cel_proto protoreflect.FileDescriptor

var file_xds_type_v4_cel_proto_rawDesc = []byte{
	0x0a, 0x15, 0x78, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x34, 0x2f, 0x63, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x76, 0x33, 0x1a, 0x16, 0x63, 0x65, 0x6c, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x0d,
	0x43, 0x65, 0x6c, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a,
	0x10, 0x63, 0x65, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x65, 0x6c, 0x2e, 0x65, 0x78,
	0x70, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x45, 0x78, 0x70, 0x72, 0x52, 0x0e,
	0x63, 0x65, 0x6c, 0x45, 0x78, 0x70, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x48,
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x78, 0x64, 0x73,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x34, 0x42, 0x08, 0x43, 0x65, 0x6c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6e, 0x63, 0x66, 0x2f, 0x78, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x78, 0x64, 0x73,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xds_type_v4_cel_proto_rawDescOnce sync.Once
	file_xds_type_v4_cel_proto_rawDescData = file_xds_type_v4_cel_proto_rawDesc
)

func file_xds_type_v4_cel_proto_rawDescGZIP() []byte {
	file_xds_type_v4_cel_proto_rawDescOnce.Do(func() {
		file_xds_type_v4_cel_proto_rawDescData = protoimpl.X.CompressGZIP(file_xds_type_v4_cel_proto_rawDescData)
	})
	return file_xds_type_v4_cel_proto_rawDescData
}

var file_xds_type_v4_cel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_xds_type_v4_cel_proto_goTypes = []interface{}{
	(*CelExpression)(nil),    // 0: xds.type.v3.CelExpression
	(*expr.CheckedExpr)(nil), // 1: cel.expr.CheckedExpr
}
var file_xds_type_v4_cel_proto_depIdxs = []int32{
	1, // 0: xds.type.v3.CelExpression.cel_expr_checked:type_name -> cel.expr.CheckedExpr
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_xds_type_v4_cel_proto_init() }
func file_xds_type_v4_cel_proto_init() {
	if File_xds_type_v4_cel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xds_type_v4_cel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CelExpression); i {
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
			RawDescriptor: file_xds_type_v4_cel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_xds_type_v4_cel_proto_goTypes,
		DependencyIndexes: file_xds_type_v4_cel_proto_depIdxs,
		MessageInfos:      file_xds_type_v4_cel_proto_msgTypes,
	}.Build()
	File_xds_type_v4_cel_proto = out.File
	file_xds_type_v4_cel_proto_rawDesc = nil
	file_xds_type_v4_cel_proto_goTypes = nil
	file_xds_type_v4_cel_proto_depIdxs = nil
}
