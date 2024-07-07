// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: protos/equipments.proto

package genprotos

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

type RawCountry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryName string  `protobuf:"bytes,1,opt,name=country_name,json=countryName,proto3" json:"country_name,omitempty"`
	Latitude    float32 `protobuf:"fixed32,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude   float32 `protobuf:"fixed32,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *RawCountry) Reset() {
	*x = RawCountry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_equipments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawCountry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawCountry) ProtoMessage() {}

func (x *RawCountry) ProtoReflect() protoreflect.Message {
	mi := &file_protos_equipments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawCountry.ProtoReflect.Descriptor instead.
func (*RawCountry) Descriptor() ([]byte, []int) {
	return file_protos_equipments_proto_rawDescGZIP(), []int{0}
}

func (x *RawCountry) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

func (x *RawCountry) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *RawCountry) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryId   int32   `protobuf:"varint,1,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	CountryName string  `protobuf:"bytes,2,opt,name=country_name,json=countryName,proto3" json:"country_name,omitempty"`
	Latitude    float32 `protobuf:"fixed32,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude   float32 `protobuf:"fixed32,4,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Country) Reset() {
	*x = Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_equipments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_protos_equipments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_protos_equipments_proto_rawDescGZIP(), []int{1}
}

func (x *Country) GetCountryId() int32 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *Country) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

func (x *Country) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Country) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type GetCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryId int32 `protobuf:"varint,1,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
}

func (x *GetCountryRequest) Reset() {
	*x = GetCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_equipments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCountryRequest) ProtoMessage() {}

func (x *GetCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_equipments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCountryRequest.ProtoReflect.Descriptor instead.
func (*GetCountryRequest) Descriptor() ([]byte, []int) {
	return file_protos_equipments_proto_rawDescGZIP(), []int{2}
}

func (x *GetCountryRequest) GetCountryId() int32 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

type GetClosestCountryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Countries []*Country `protobuf:"bytes,1,rep,name=countries,proto3" json:"countries,omitempty"`
}

func (x *GetClosestCountryResponse) Reset() {
	*x = GetClosestCountryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_equipments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClosestCountryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClosestCountryResponse) ProtoMessage() {}

func (x *GetClosestCountryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_equipments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClosestCountryResponse.ProtoReflect.Descriptor instead.
func (*GetClosestCountryResponse) Descriptor() ([]byte, []int) {
	return file_protos_equipments_proto_rawDescGZIP(), []int{3}
}

func (x *GetClosestCountryResponse) GetCountries() []*Country {
	if x != nil {
		return x.Countries
	}
	return nil
}

var File_protos_equipments_proto protoreflect.FileDescriptor

var file_protos_equipments_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0a, 0x52, 0x61, 0x77,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x32, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64,
	0x22, 0x43, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x32, 0xad, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0b, 0x2e, 0x52, 0x61, 0x77, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x08, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x43, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_equipments_proto_rawDescOnce sync.Once
	file_protos_equipments_proto_rawDescData = file_protos_equipments_proto_rawDesc
)

func file_protos_equipments_proto_rawDescGZIP() []byte {
	file_protos_equipments_proto_rawDescOnce.Do(func() {
		file_protos_equipments_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_equipments_proto_rawDescData)
	})
	return file_protos_equipments_proto_rawDescData
}

var file_protos_equipments_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_equipments_proto_goTypes = []any{
	(*RawCountry)(nil),                // 0: RawCountry
	(*Country)(nil),                   // 1: Country
	(*GetCountryRequest)(nil),         // 2: GetCountryRequest
	(*GetClosestCountryResponse)(nil), // 3: GetClosestCountryResponse
}
var file_protos_equipments_proto_depIdxs = []int32{
	1, // 0: GetClosestCountryResponse.countries:type_name -> Country
	0, // 1: CountryService.CreateCountry:input_type -> RawCountry
	2, // 2: CountryService.GetClosestCountry:input_type -> GetCountryRequest
	2, // 3: CountryService.GetCountryById:input_type -> GetCountryRequest
	1, // 4: CountryService.CreateCountry:output_type -> Country
	3, // 5: CountryService.GetClosestCountry:output_type -> GetClosestCountryResponse
	1, // 6: CountryService.GetCountryById:output_type -> Country
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_equipments_proto_init() }
func file_protos_equipments_proto_init() {
	if File_protos_equipments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_equipments_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RawCountry); i {
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
		file_protos_equipments_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Country); i {
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
		file_protos_equipments_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetCountryRequest); i {
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
		file_protos_equipments_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetClosestCountryResponse); i {
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
			RawDescriptor: file_protos_equipments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_equipments_proto_goTypes,
		DependencyIndexes: file_protos_equipments_proto_depIdxs,
		MessageInfos:      file_protos_equipments_proto_msgTypes,
	}.Build()
	File_protos_equipments_proto = out.File
	file_protos_equipments_proto_rawDesc = nil
	file_protos_equipments_proto_goTypes = nil
	file_protos_equipments_proto_depIdxs = nil
}
