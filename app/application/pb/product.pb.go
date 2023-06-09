// バージョン3を利用する

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: proto/product.proto

// パッケージの宣言

package pb

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

// *****
// パラメータメッセージの定義
// *****
type ProductParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      *string `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`           // 商品Id
	Keyword *string `protobuf:"bytes,2,opt,name=keyword,proto3,oneof" json:"keyword,omitempty"` // キーワード
}

func (x *ProductParam) Reset() {
	*x = ProductParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductParam) ProtoMessage() {}

func (x *ProductParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductParam.ProtoReflect.Descriptor instead.
func (*ProductParam) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductParam) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *ProductParam) GetKeyword() string {
	if x != nil && x.Keyword != nil {
		return *x.Keyword
	}
	return ""
}

// *****
// 商品検索結果メッセージの定義
// *****
type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`        // 商品Id
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`    // 商品名
	Price int32  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"` // 単価
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

// *****
// 検索結果メッセージの定義
// *****
type ProductResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// エラーか検索結果のいずれかを返す
	//
	// Types that are assignable to Result:
	//
	//	*ProductResult_Product
	//	*ProductResult_Error
	Result isProductResult_Result `protobuf_oneof:"result"`
}

func (x *ProductResult) Reset() {
	*x = ProductResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResult) ProtoMessage() {}

func (x *ProductResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResult.ProtoReflect.Descriptor instead.
func (*ProductResult) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{2}
}

func (m *ProductResult) GetResult() isProductResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *ProductResult) GetProduct() *Product {
	if x, ok := x.GetResult().(*ProductResult_Product); ok {
		return x.Product
	}
	return nil
}

func (x *ProductResult) GetError() *Error {
	if x, ok := x.GetResult().(*ProductResult_Error); ok {
		return x.Error
	}
	return nil
}

type isProductResult_Result interface {
	isProductResult_Result()
}

type ProductResult_Product struct {
	Product *Product `protobuf:"bytes,1,opt,name=product,proto3,oneof"` // 検索結果
}

type ProductResult_Error struct {
	Error *Error `protobuf:"bytes,2,opt,name=error,proto3,oneof"` // 検索エラー
}

func (*ProductResult_Product) isProductResult_Result() {}

func (*ProductResult_Error) isProductResult_Result() {}

// *****
// 商品一覧メッセージの定義
// *****
type ProductsResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"` // 一覧
	Error    *Error     `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"` // エラー
}

func (x *ProductsResult) Reset() {
	*x = ProductsResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsResult) ProtoMessage() {}

func (x *ProductsResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsResult.ProtoReflect.Descriptor instead.
func (*ProductsResult) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{3}
}

func (x *ProductsResult) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *ProductsResult) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_proto_product_proto protoreflect.FileDescriptor

var file_proto_product_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x55, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12,
	0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x43, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x6b, 0x0a, 0x0d, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2a, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x6f, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xb0, 0x01, 0x0a, 0x0e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x31, 0x0a, 0x04, 0x42, 0x79, 0x49, 0x64, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x11, 0x5a, 0x0f,
	0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_product_proto_rawDescOnce sync.Once
	file_proto_product_proto_rawDescData = file_proto_product_proto_rawDesc
)

func file_proto_product_proto_rawDescGZIP() []byte {
	file_proto_product_proto_rawDescOnce.Do(func() {
		file_proto_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_product_proto_rawDescData)
	})
	return file_proto_product_proto_rawDescData
}

var file_proto_product_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_product_proto_goTypes = []interface{}{
	(*ProductParam)(nil),   // 0: proto.ProductParam
	(*Product)(nil),        // 1: proto.Product
	(*ProductResult)(nil),  // 2: proto.ProductResult
	(*ProductsResult)(nil), // 3: proto.ProductsResult
	(*Error)(nil),          // 4: proto.Error
}
var file_proto_product_proto_depIdxs = []int32{
	1, // 0: proto.ProductResult.product:type_name -> proto.Product
	4, // 1: proto.ProductResult.error:type_name -> proto.Error
	1, // 2: proto.ProductsResult.products:type_name -> proto.Product
	4, // 3: proto.ProductsResult.error:type_name -> proto.Error
	0, // 4: proto.ProductService.List:input_type -> proto.ProductParam
	0, // 5: proto.ProductService.ById:input_type -> proto.ProductParam
	0, // 6: proto.ProductService.ByKeyword:input_type -> proto.ProductParam
	3, // 7: proto.ProductService.List:output_type -> proto.ProductsResult
	2, // 8: proto.ProductService.ById:output_type -> proto.ProductResult
	3, // 9: proto.ProductService.ByKeyword:output_type -> proto.ProductsResult
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_product_proto_init() }
func file_proto_product_proto_init() {
	if File_proto_product_proto != nil {
		return
	}
	file_proto_error_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductParam); i {
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
		file_proto_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_proto_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductResult); i {
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
		file_proto_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsResult); i {
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
	file_proto_product_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_product_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ProductResult_Product)(nil),
		(*ProductResult_Error)(nil),
	}
	file_proto_product_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_product_proto_goTypes,
		DependencyIndexes: file_proto_product_proto_depIdxs,
		MessageInfos:      file_proto_product_proto_msgTypes,
	}.Build()
	File_proto_product_proto = out.File
	file_proto_product_proto_rawDesc = nil
	file_proto_product_proto_goTypes = nil
	file_proto_product_proto_depIdxs = nil
}
