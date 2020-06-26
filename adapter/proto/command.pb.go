// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.0
// source: adapter/proto/command.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type OrderCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*OrderCommand_Create
	Payload isOrderCommand_Payload `protobuf_oneof:"payload"`
}

func (x *OrderCommand) Reset() {
	*x = OrderCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adapter_proto_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCommand) ProtoMessage() {}

func (x *OrderCommand) ProtoReflect() protoreflect.Message {
	mi := &file_adapter_proto_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCommand.ProtoReflect.Descriptor instead.
func (*OrderCommand) Descriptor() ([]byte, []int) {
	return file_adapter_proto_command_proto_rawDescGZIP(), []int{0}
}

func (m *OrderCommand) GetPayload() isOrderCommand_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *OrderCommand) GetCreate() *CreateOrderCommand {
	if x, ok := x.GetPayload().(*OrderCommand_Create); ok {
		return x.Create
	}
	return nil
}

type isOrderCommand_Payload interface {
	isOrderCommand_Payload()
}

type OrderCommand_Create struct {
	Create *CreateOrderCommand `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

func (*OrderCommand_Create) isOrderCommand_Payload() {}

type CreateOrderCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID              string   `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	CustomerID           string   `protobuf:"bytes,2,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	ProductID            string   `protobuf:"bytes,3,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	Quantity             uint32   `protobuf:"varint,4,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	ExpectedDeliveryDate string   `protobuf:"bytes,5,opt,name=ExpectedDeliveryDate,proto3" json:"ExpectedDeliveryDate,omitempty"`
	PickupDate           string   `protobuf:"bytes,6,opt,name=PickupDate,proto3" json:"PickupDate,omitempty"`
	PickupAddress        *Address `protobuf:"bytes,7,opt,name=PickupAddress,proto3" json:"PickupAddress,omitempty"`
	DestinationAddress   *Address `protobuf:"bytes,8,opt,name=DestinationAddress,proto3" json:"DestinationAddress,omitempty"`
}

func (x *CreateOrderCommand) Reset() {
	*x = CreateOrderCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adapter_proto_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderCommand) ProtoMessage() {}

func (x *CreateOrderCommand) ProtoReflect() protoreflect.Message {
	mi := &file_adapter_proto_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderCommand.ProtoReflect.Descriptor instead.
func (*CreateOrderCommand) Descriptor() ([]byte, []int) {
	return file_adapter_proto_command_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderCommand) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *CreateOrderCommand) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

func (x *CreateOrderCommand) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *CreateOrderCommand) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CreateOrderCommand) GetExpectedDeliveryDate() string {
	if x != nil {
		return x.ExpectedDeliveryDate
	}
	return ""
}

func (x *CreateOrderCommand) GetPickupDate() string {
	if x != nil {
		return x.PickupDate
	}
	return ""
}

func (x *CreateOrderCommand) GetPickupAddress() *Address {
	if x != nil {
		return x.PickupAddress
	}
	return nil
}

func (x *CreateOrderCommand) GetDestinationAddress() *Address {
	if x != nil {
		return x.DestinationAddress
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street  string `protobuf:"bytes,1,opt,name=Street,proto3" json:"Street,omitempty"`
	City    string `protobuf:"bytes,2,opt,name=City,proto3" json:"City,omitempty"`
	Country string `protobuf:"bytes,3,opt,name=Country,proto3" json:"Country,omitempty"`
	State   string `protobuf:"bytes,4,opt,name=State,proto3" json:"State,omitempty"`
	ZipCode string `protobuf:"bytes,5,opt,name=ZipCode,proto3" json:"ZipCode,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adapter_proto_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_adapter_proto_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_adapter_proto_command_proto_rawDescGZIP(), []int{2}
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

var File_adapter_proto_command_proto protoreflect.FileDescriptor

var file_adapter_proto_command_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x56, 0x0a, 0x0c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x3b, 0x0a, 0x06,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48,
	0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0xe2, 0x02, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x32, 0x0a, 0x14, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x45,
	0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x0d, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x46, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x12, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x7f, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x5a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x5a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x61, 0x64,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_adapter_proto_command_proto_rawDescOnce sync.Once
	file_adapter_proto_command_proto_rawDescData = file_adapter_proto_command_proto_rawDesc
)

func file_adapter_proto_command_proto_rawDescGZIP() []byte {
	file_adapter_proto_command_proto_rawDescOnce.Do(func() {
		file_adapter_proto_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_adapter_proto_command_proto_rawDescData)
	})
	return file_adapter_proto_command_proto_rawDescData
}

var file_adapter_proto_command_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_adapter_proto_command_proto_goTypes = []interface{}{
	(*OrderCommand)(nil),       // 0: order.command.OrderCommand
	(*CreateOrderCommand)(nil), // 1: order.command.CreateOrderCommand
	(*Address)(nil),            // 2: order.command.Address
}
var file_adapter_proto_command_proto_depIdxs = []int32{
	1, // 0: order.command.OrderCommand.create:type_name -> order.command.CreateOrderCommand
	2, // 1: order.command.CreateOrderCommand.PickupAddress:type_name -> order.command.Address
	2, // 2: order.command.CreateOrderCommand.DestinationAddress:type_name -> order.command.Address
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_adapter_proto_command_proto_init() }
func file_adapter_proto_command_proto_init() {
	if File_adapter_proto_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_adapter_proto_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCommand); i {
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
		file_adapter_proto_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderCommand); i {
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
		file_adapter_proto_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
	file_adapter_proto_command_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OrderCommand_Create)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_adapter_proto_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_adapter_proto_command_proto_goTypes,
		DependencyIndexes: file_adapter_proto_command_proto_depIdxs,
		MessageInfos:      file_adapter_proto_command_proto_msgTypes,
	}.Build()
	File_adapter_proto_command_proto = out.File
	file_adapter_proto_command_proto_rawDesc = nil
	file_adapter_proto_command_proto_goTypes = nil
	file_adapter_proto_command_proto_depIdxs = nil
}
