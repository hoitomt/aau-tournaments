// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: ee_request.proto

package v1

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

type ExposureEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate   string `protobuf:"bytes,1,opt,name=startDate,json=StartDateString,proto3" json:"startDate,omitempty"`
	EndDate     string `protobuf:"bytes,2,opt,name=endDate,json=EndDateString,proto3" json:"endDate,omitempty"`
	Gender      string `protobuf:"bytes,3,opt,name=gender,json=Gender,proto3" json:"gender,omitempty"`
	Page        int32  `protobuf:"varint,4,opt,name=page,json=Page,proto3" json:"page,omitempty"`
	InviteType  string `protobuf:"bytes,5,opt,name=inviteType,json=InviteType,proto3" json:"inviteType,omitempty"`
	SearchToken string `protobuf:"bytes,6,opt,name=searchToken,json=SearchToken,proto3" json:"searchToken,omitempty"`
	SportType   string `protobuf:"bytes,7,opt,name=sportType,proto3" json:"sportType,omitempty"`
}

func (x *ExposureEventRequest) Reset() {
	*x = ExposureEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ee_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExposureEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExposureEventRequest) ProtoMessage() {}

func (x *ExposureEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ee_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExposureEventRequest.ProtoReflect.Descriptor instead.
func (*ExposureEventRequest) Descriptor() ([]byte, []int) {
	return file_ee_request_proto_rawDescGZIP(), []int{0}
}

func (x *ExposureEventRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *ExposureEventRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *ExposureEventRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *ExposureEventRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ExposureEventRequest) GetInviteType() string {
	if x != nil {
		return x.InviteType
	}
	return ""
}

func (x *ExposureEventRequest) GetSearchToken() string {
	if x != nil {
		return x.SearchToken
	}
	return ""
}

func (x *ExposureEventRequest) GetSportType() string {
	if x != nil {
		return x.SportType
	}
	return ""
}

var File_ee_request_proto protoreflect.FileDescriptor

var file_ee_request_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0xe6, 0x01, 0x0a, 0x14, 0x45,
	0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x2d,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ee_request_proto_rawDescOnce sync.Once
	file_ee_request_proto_rawDescData = file_ee_request_proto_rawDesc
)

func file_ee_request_proto_rawDescGZIP() []byte {
	file_ee_request_proto_rawDescOnce.Do(func() {
		file_ee_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_ee_request_proto_rawDescData)
	})
	return file_ee_request_proto_rawDescData
}

var file_ee_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ee_request_proto_goTypes = []interface{}{
	(*ExposureEventRequest)(nil), // 0: protos.ExposureEventRequest
}
var file_ee_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ee_request_proto_init() }
func file_ee_request_proto_init() {
	if File_ee_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ee_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExposureEventRequest); i {
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
			RawDescriptor: file_ee_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ee_request_proto_goTypes,
		DependencyIndexes: file_ee_request_proto_depIdxs,
		MessageInfos:      file_ee_request_proto_msgTypes,
	}.Build()
	File_ee_request_proto = out.File
	file_ee_request_proto_rawDesc = nil
	file_ee_request_proto_goTypes = nil
	file_ee_request_proto_depIdxs = nil
}
