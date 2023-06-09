// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: proto/suicide.proto

package grpc

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_suicide_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_suicide_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_suicide_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

type Suicide struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetStatsData *GetStatsData `protobuf:"bytes,1,opt,name=getStatsData,proto3" json:"GET_STATS_DATA"` //@gotags: json:"GET_STATS_DATA"
}

func (x *Suicide) Reset() {
	*x = Suicide{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_suicide_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Suicide) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Suicide) ProtoMessage() {}

func (x *Suicide) ProtoReflect() protoreflect.Message {
	mi := &file_proto_suicide_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Suicide.ProtoReflect.Descriptor instead.
func (*Suicide) Descriptor() ([]byte, []int) {
	return file_proto_suicide_proto_rawDescGZIP(), []int{1}
}

func (x *Suicide) GetGetStatsData() *GetStatsData {
	if x != nil {
		return x.GetStatsData
	}
	return nil
}

type GetStatsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatisticalData *StatisticalData `protobuf:"bytes,1,opt,name=statisticalData,proto3" json:"STATISTICAL_DATA"` //@gotags: json:"STATISTICAL_DATA"
}

func (x *GetStatsData) Reset() {
	*x = GetStatsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_suicide_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsData) ProtoMessage() {}

func (x *GetStatsData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_suicide_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsData.ProtoReflect.Descriptor instead.
func (*GetStatsData) Descriptor() ([]byte, []int) {
	return file_proto_suicide_proto_rawDescGZIP(), []int{2}
}

func (x *GetStatsData) GetStatisticalData() *StatisticalData {
	if x != nil {
		return x.StatisticalData
	}
	return nil
}

type StatisticalData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataInf *DataInf `protobuf:"bytes,1,opt,name=dataInf,proto3" json:"DATA_INF"` //@gotags: json:"DATA_INF"
}

func (x *StatisticalData) Reset() {
	*x = StatisticalData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_suicide_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticalData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticalData) ProtoMessage() {}

func (x *StatisticalData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_suicide_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticalData.ProtoReflect.Descriptor instead.
func (*StatisticalData) Descriptor() ([]byte, []int) {
	return file_proto_suicide_proto_rawDescGZIP(), []int{3}
}

func (x *StatisticalData) GetDataInf() *DataInf {
	if x != nil {
		return x.DataInf
	}
	return nil
}

type DataInf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []*Value `protobuf:"bytes,1,rep,name=value,proto3" json:"VALUE"` //@gotags: json:"VALUE"
}

func (x *DataInf) Reset() {
	*x = DataInf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_suicide_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataInf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataInf) ProtoMessage() {}

func (x *DataInf) ProtoReflect() protoreflect.Message {
	mi := &file_proto_suicide_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataInf.ProtoReflect.Descriptor instead.
func (*DataInf) Descriptor() ([]byte, []int) {
	return file_proto_suicide_proto_rawDescGZIP(), []int{4}
}

func (x *DataInf) GetValue() []*Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time string `protobuf:"bytes,1,opt,name=time,proto3" json:"@time"` //@gotags: json:"@time"
	Unit string `protobuf:"bytes,2,opt,name=unit,proto3" json:"@unit"` //@gotags: json:"@unit"
	V    string `protobuf:"bytes,3,opt,name=v,proto3" json:"$"`       //@gotags: json:"$"
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_suicide_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_proto_suicide_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_proto_suicide_proto_rawDescGZIP(), []int{5}
}

func (x *Value) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Value) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Value) GetV() string {
	if x != nil {
		return x.V
	}
	return ""
}

var File_proto_suicide_proto protoreflect.FileDescriptor

var file_proto_suicide_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75, 0x69, 0x63, 0x69, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x23, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x41, 0x0a, 0x07, 0x53, 0x75, 0x69, 0x63, 0x69, 0x64, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x67,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x4f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x49,
	0x6e, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66,
	0x22, 0x2c, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x12, 0x21, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3d,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12,
	0x0c, 0x0a, 0x01, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x76, 0x32, 0x40, 0x0a,
	0x0e, 0x53, 0x75, 0x69, 0x63, 0x69, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2e, 0x0a, 0x0e, 0x53, 0x75, 0x69, 0x63, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x75, 0x69, 0x63, 0x69, 0x64, 0x65, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_suicide_proto_rawDescOnce sync.Once
	file_proto_suicide_proto_rawDescData = file_proto_suicide_proto_rawDesc
)

func file_proto_suicide_proto_rawDescGZIP() []byte {
	file_proto_suicide_proto_rawDescOnce.Do(func() {
		file_proto_suicide_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_suicide_proto_rawDescData)
	})
	return file_proto_suicide_proto_rawDescData
}

var file_proto_suicide_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_suicide_proto_goTypes = []interface{}{
	(*Request)(nil),         // 0: main.Request
	(*Suicide)(nil),         // 1: main.Suicide
	(*GetStatsData)(nil),    // 2: main.GetStatsData
	(*StatisticalData)(nil), // 3: main.StatisticalData
	(*DataInf)(nil),         // 4: main.DataInf
	(*Value)(nil),           // 5: main.Value
}
var file_proto_suicide_proto_depIdxs = []int32{
	2, // 0: main.Suicide.getStatsData:type_name -> main.GetStatsData
	3, // 1: main.GetStatsData.statisticalData:type_name -> main.StatisticalData
	4, // 2: main.StatisticalData.dataInf:type_name -> main.DataInf
	5, // 3: main.DataInf.value:type_name -> main.Value
	0, // 4: main.SuicideService.SuicideRequest:input_type -> main.Request
	1, // 5: main.SuicideService.SuicideRequest:output_type -> main.Suicide
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_suicide_proto_init() }
func file_proto_suicide_proto_init() {
	if File_proto_suicide_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_suicide_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_suicide_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Suicide); i {
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
		file_proto_suicide_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatsData); i {
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
		file_proto_suicide_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatisticalData); i {
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
		file_proto_suicide_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataInf); i {
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
		file_proto_suicide_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
			RawDescriptor: file_proto_suicide_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_suicide_proto_goTypes,
		DependencyIndexes: file_proto_suicide_proto_depIdxs,
		MessageInfos:      file_proto_suicide_proto_msgTypes,
	}.Build()
	File_proto_suicide_proto = out.File
	file_proto_suicide_proto_rawDesc = nil
	file_proto_suicide_proto_goTypes = nil
	file_proto_suicide_proto_depIdxs = nil
}
