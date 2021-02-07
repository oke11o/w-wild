// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: trainer.proto

package trainer

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type IsHourAvailableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *IsHourAvailableRequest) Reset() {
	*x = IsHourAvailableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsHourAvailableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsHourAvailableRequest) ProtoMessage() {}

func (x *IsHourAvailableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsHourAvailableRequest.ProtoReflect.Descriptor instead.
func (*IsHourAvailableRequest) Descriptor() ([]byte, []int) {
	return file_trainer_proto_rawDescGZIP(), []int{0}
}

func (x *IsHourAvailableRequest) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type IsHourAvailableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAvailable bool `protobuf:"varint,1,opt,name=is_available,json=isAvailable,proto3" json:"is_available,omitempty"`
}

func (x *IsHourAvailableResponse) Reset() {
	*x = IsHourAvailableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsHourAvailableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsHourAvailableResponse) ProtoMessage() {}

func (x *IsHourAvailableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsHourAvailableResponse.ProtoReflect.Descriptor instead.
func (*IsHourAvailableResponse) Descriptor() ([]byte, []int) {
	return file_trainer_proto_rawDescGZIP(), []int{1}
}

func (x *IsHourAvailableResponse) GetIsAvailable() bool {
	if x != nil {
		return x.IsAvailable
	}
	return false
}

type UpdateHourRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time                 *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	HasTrainingScheduled bool                   `protobuf:"varint,2,opt,name=has_training_scheduled,json=hasTrainingScheduled,proto3" json:"has_training_scheduled,omitempty"`
	Available            bool                   `protobuf:"varint,3,opt,name=available,proto3" json:"available,omitempty"`
}

func (x *UpdateHourRequest) Reset() {
	*x = UpdateHourRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHourRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHourRequest) ProtoMessage() {}

func (x *UpdateHourRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHourRequest.ProtoReflect.Descriptor instead.
func (*UpdateHourRequest) Descriptor() ([]byte, []int) {
	return file_trainer_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateHourRequest) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *UpdateHourRequest) GetHasTrainingScheduled() bool {
	if x != nil {
		return x.HasTrainingScheduled
	}
	return false
}

func (x *UpdateHourRequest) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trainer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trainer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_trainer_proto_rawDescGZIP(), []int{3}
}

var File_trainer_proto protoreflect.FileDescriptor

var file_trainer_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x16, 0x49, 0x73, 0x48,
	0x6f, 0x75, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x17, 0x49, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x22, 0x97, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x68, 0x61, 0x73, 0x5f, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x68, 0x61, 0x73, 0x54, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xac, 0x01, 0x0a,
	0x0e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x56, 0x0a, 0x0f, 0x49, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x49, 0x73, 0x48,
	0x6f, 0x75, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x49, 0x73,
	0x48, 0x6f, 0x75, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x48, 0x6f, 0x75, 0x72, 0x12, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e,
	0x3b, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trainer_proto_rawDescOnce sync.Once
	file_trainer_proto_rawDescData = file_trainer_proto_rawDesc
)

func file_trainer_proto_rawDescGZIP() []byte {
	file_trainer_proto_rawDescOnce.Do(func() {
		file_trainer_proto_rawDescData = protoimpl.X.CompressGZIP(file_trainer_proto_rawDescData)
	})
	return file_trainer_proto_rawDescData
}

var file_trainer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_trainer_proto_goTypes = []interface{}{
	(*IsHourAvailableRequest)(nil),  // 0: trainer.IsHourAvailableRequest
	(*IsHourAvailableResponse)(nil), // 1: trainer.IsHourAvailableResponse
	(*UpdateHourRequest)(nil),       // 2: trainer.UpdateHourRequest
	(*EmptyResponse)(nil),           // 3: trainer.EmptyResponse
	(*timestamppb.Timestamp)(nil),   // 4: google.protobuf.Timestamp
}
var file_trainer_proto_depIdxs = []int32{
	4, // 0: trainer.IsHourAvailableRequest.time:type_name -> google.protobuf.Timestamp
	4, // 1: trainer.UpdateHourRequest.time:type_name -> google.protobuf.Timestamp
	0, // 2: trainer.TrainerService.IsHourAvailable:input_type -> trainer.IsHourAvailableRequest
	2, // 3: trainer.TrainerService.UpdateHour:input_type -> trainer.UpdateHourRequest
	1, // 4: trainer.TrainerService.IsHourAvailable:output_type -> trainer.IsHourAvailableResponse
	3, // 5: trainer.TrainerService.UpdateHour:output_type -> trainer.EmptyResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_trainer_proto_init() }
func file_trainer_proto_init() {
	if File_trainer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trainer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsHourAvailableRequest); i {
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
		file_trainer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsHourAvailableResponse); i {
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
		file_trainer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHourRequest); i {
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
		file_trainer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
			RawDescriptor: file_trainer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trainer_proto_goTypes,
		DependencyIndexes: file_trainer_proto_depIdxs,
		MessageInfos:      file_trainer_proto_msgTypes,
	}.Build()
	File_trainer_proto = out.File
	file_trainer_proto_rawDesc = nil
	file_trainer_proto_goTypes = nil
	file_trainer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TrainerServiceClient is the client API for TrainerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TrainerServiceClient interface {
	IsHourAvailable(ctx context.Context, in *IsHourAvailableRequest, opts ...grpc.CallOption) (*IsHourAvailableResponse, error)
	UpdateHour(ctx context.Context, in *UpdateHourRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type trainerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainerServiceClient(cc grpc.ClientConnInterface) TrainerServiceClient {
	return &trainerServiceClient{cc}
}

func (c *trainerServiceClient) IsHourAvailable(ctx context.Context, in *IsHourAvailableRequest, opts ...grpc.CallOption) (*IsHourAvailableResponse, error) {
	out := new(IsHourAvailableResponse)
	err := c.cc.Invoke(ctx, "/trainer.TrainerService/IsHourAvailable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainerServiceClient) UpdateHour(ctx context.Context, in *UpdateHourRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/trainer.TrainerService/UpdateHour", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainerServiceServer is the server API for TrainerService service.
type TrainerServiceServer interface {
	IsHourAvailable(context.Context, *IsHourAvailableRequest) (*IsHourAvailableResponse, error)
	UpdateHour(context.Context, *UpdateHourRequest) (*EmptyResponse, error)
}

// UnimplementedTrainerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTrainerServiceServer struct {
}

func (*UnimplementedTrainerServiceServer) IsHourAvailable(context.Context, *IsHourAvailableRequest) (*IsHourAvailableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsHourAvailable not implemented")
}
func (*UnimplementedTrainerServiceServer) UpdateHour(context.Context, *UpdateHourRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHour not implemented")
}

func RegisterTrainerServiceServer(s *grpc.Server, srv TrainerServiceServer) {
	s.RegisterService(&_TrainerService_serviceDesc, srv)
}

func _TrainerService_IsHourAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsHourAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainerServiceServer).IsHourAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trainer.TrainerService/IsHourAvailable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainerServiceServer).IsHourAvailable(ctx, req.(*IsHourAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainerService_UpdateHour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainerServiceServer).UpdateHour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trainer.TrainerService/UpdateHour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainerServiceServer).UpdateHour(ctx, req.(*UpdateHourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TrainerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "trainer.TrainerService",
	HandlerType: (*TrainerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsHourAvailable",
			Handler:    _TrainerService_IsHourAvailable_Handler,
		},
		{
			MethodName: "UpdateHour",
			Handler:    _TrainerService_UpdateHour_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trainer.proto",
}
