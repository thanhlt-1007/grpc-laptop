// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: messages/laptop_message/message.proto

package laptop_message

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	cpu_message "grpc-laptop/go_protos/messages/cpu_message"
	gpu_message "grpc-laptop/go_protos/messages/gpu_message"
	keyboard_message "grpc-laptop/go_protos/messages/keyboard_message"
	memory_message "grpc-laptop/go_protos/messages/memory_message"
	screen_message "grpc-laptop/go_protos/messages/screen_message"
	storage_message "grpc-laptop/go_protos/messages/storage_message"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Laptop struct {
	state    protoimpl.MessageState     `protogen:"open.v1"`
	Id       string                     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand    string                     `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Name     string                     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Cpu      *cpu_message.CPU           `protobuf:"bytes,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory   *memory_message.Memory     `protobuf:"bytes,5,opt,name=memory,proto3" json:"memory,omitempty"`
	Gpu      []*gpu_message.GPU         `protobuf:"bytes,6,rep,name=gpu,proto3" json:"gpu,omitempty"`
	Storage  []*storage_message.Storage `protobuf:"bytes,7,rep,name=storage,proto3" json:"storage,omitempty"`
	Screen   *screen_message.Screen     `protobuf:"bytes,8,opt,name=screen,proto3" json:"screen,omitempty"`
	Keyboard *keyboard_message.Keyboard `protobuf:"bytes,9,opt,name=keyboard,proto3" json:"keyboard,omitempty"`
	// Types that are valid to be assigned to Weight:
	//
	//	*Laptop_WeightKg
	//	*Laptop_WeightLb
	Weight        isLaptop_Weight        `protobuf_oneof:"weight"`
	PriceUsd      float64                `protobuf:"fixed64,12,opt,name=price_usd,json=priceUsd,proto3" json:"price_usd,omitempty"`
	ReleaseYear   uint32                 `protobuf:"varint,13,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Laptop) Reset() {
	*x = Laptop{}
	mi := &file_messages_laptop_message_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Laptop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Laptop) ProtoMessage() {}

func (x *Laptop) ProtoReflect() protoreflect.Message {
	mi := &file_messages_laptop_message_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Laptop.ProtoReflect.Descriptor instead.
func (*Laptop) Descriptor() ([]byte, []int) {
	return file_messages_laptop_message_message_proto_rawDescGZIP(), []int{0}
}

func (x *Laptop) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Laptop) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Laptop) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Laptop) GetCpu() *cpu_message.CPU {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *Laptop) GetMemory() *memory_message.Memory {
	if x != nil {
		return x.Memory
	}
	return nil
}

func (x *Laptop) GetGpu() []*gpu_message.GPU {
	if x != nil {
		return x.Gpu
	}
	return nil
}

func (x *Laptop) GetStorage() []*storage_message.Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *Laptop) GetScreen() *screen_message.Screen {
	if x != nil {
		return x.Screen
	}
	return nil
}

func (x *Laptop) GetKeyboard() *keyboard_message.Keyboard {
	if x != nil {
		return x.Keyboard
	}
	return nil
}

func (x *Laptop) GetWeight() isLaptop_Weight {
	if x != nil {
		return x.Weight
	}
	return nil
}

func (x *Laptop) GetWeightKg() float64 {
	if x != nil {
		if x, ok := x.Weight.(*Laptop_WeightKg); ok {
			return x.WeightKg
		}
	}
	return 0
}

func (x *Laptop) GetWeightLb() float64 {
	if x != nil {
		if x, ok := x.Weight.(*Laptop_WeightLb); ok {
			return x.WeightLb
		}
	}
	return 0
}

func (x *Laptop) GetPriceUsd() float64 {
	if x != nil {
		return x.PriceUsd
	}
	return 0
}

func (x *Laptop) GetReleaseYear() uint32 {
	if x != nil {
		return x.ReleaseYear
	}
	return 0
}

func (x *Laptop) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type isLaptop_Weight interface {
	isLaptop_Weight()
}

type Laptop_WeightKg struct {
	WeightKg float64 `protobuf:"fixed64,10,opt,name=weight_kg,json=weightKg,proto3,oneof"`
}

type Laptop_WeightLb struct {
	WeightLb float64 `protobuf:"fixed64,11,opt,name=weight_lb,json=weightLb,proto3,oneof"`
}

func (*Laptop_WeightKg) isLaptop_Weight() {}

func (*Laptop_WeightLb) isLaptop_Weight() {}

var File_messages_laptop_message_message_proto protoreflect.FileDescriptor

var file_messages_laptop_message_message_proto_rawDesc = string([]byte{
	0x0a, 0x25, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x6c, 0x61, 0x70, 0x74, 0x6f,
	0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2f, 0x63, 0x70, 0x75, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x67, 0x70, 0x75,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f,
	0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf9, 0x04, 0x0a, 0x06, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x63, 0x70, 0x75, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x43, 0x50, 0x55, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x3e, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x03, 0x67, 0x70, 0x75, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x67, 0x70, 0x75, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x47, 0x50, 0x55, 0x52, 0x03, 0x67, 0x70, 0x75, 0x12, 0x42, 0x0a, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x12, 0x3e, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x06, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x12, 0x46, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x08,
	0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x5f, 0x6b, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x08, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x4b, 0x67, 0x12, 0x1d, 0x0a, 0x09, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x5f, 0x6c, 0x62, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x08, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x4c, 0x62, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f,
	0x75, 0x73, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x55, 0x73, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x79,
	0x65, 0x61, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x59, 0x65, 0x61, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x42, 0x08, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x2f, 0x5a, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x2f, 0x67, 0x6f, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x6c, 0x61,
	0x70, 0x74, 0x6f, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_messages_laptop_message_message_proto_rawDescOnce sync.Once
	file_messages_laptop_message_message_proto_rawDescData []byte
)

func file_messages_laptop_message_message_proto_rawDescGZIP() []byte {
	file_messages_laptop_message_message_proto_rawDescOnce.Do(func() {
		file_messages_laptop_message_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_messages_laptop_message_message_proto_rawDesc), len(file_messages_laptop_message_message_proto_rawDesc)))
	})
	return file_messages_laptop_message_message_proto_rawDescData
}

var file_messages_laptop_message_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_messages_laptop_message_message_proto_goTypes = []any{
	(*Laptop)(nil),                    // 0: protos.messages.laptop_message.Laptop
	(*cpu_message.CPU)(nil),           // 1: protos.messages.cpu_message.CPU
	(*memory_message.Memory)(nil),     // 2: protos.messages.memory_message.Memory
	(*gpu_message.GPU)(nil),           // 3: protos.messages.gpu_message.GPU
	(*storage_message.Storage)(nil),   // 4: protos.messages.storage_message.Storage
	(*screen_message.Screen)(nil),     // 5: protos.messages.screen_message.Screen
	(*keyboard_message.Keyboard)(nil), // 6: protos.messages.keyboard_message.Keyboard
	(*timestamppb.Timestamp)(nil),     // 7: google.protobuf.Timestamp
}
var file_messages_laptop_message_message_proto_depIdxs = []int32{
	1, // 0: protos.messages.laptop_message.Laptop.cpu:type_name -> protos.messages.cpu_message.CPU
	2, // 1: protos.messages.laptop_message.Laptop.memory:type_name -> protos.messages.memory_message.Memory
	3, // 2: protos.messages.laptop_message.Laptop.gpu:type_name -> protos.messages.gpu_message.GPU
	4, // 3: protos.messages.laptop_message.Laptop.storage:type_name -> protos.messages.storage_message.Storage
	5, // 4: protos.messages.laptop_message.Laptop.screen:type_name -> protos.messages.screen_message.Screen
	6, // 5: protos.messages.laptop_message.Laptop.keyboard:type_name -> protos.messages.keyboard_message.Keyboard
	7, // 6: protos.messages.laptop_message.Laptop.updated_at:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_messages_laptop_message_message_proto_init() }
func file_messages_laptop_message_message_proto_init() {
	if File_messages_laptop_message_message_proto != nil {
		return
	}
	file_messages_laptop_message_message_proto_msgTypes[0].OneofWrappers = []any{
		(*Laptop_WeightKg)(nil),
		(*Laptop_WeightLb)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_messages_laptop_message_message_proto_rawDesc), len(file_messages_laptop_message_message_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_laptop_message_message_proto_goTypes,
		DependencyIndexes: file_messages_laptop_message_message_proto_depIdxs,
		MessageInfos:      file_messages_laptop_message_message_proto_msgTypes,
	}.Build()
	File_messages_laptop_message_message_proto = out.File
	file_messages_laptop_message_message_proto_goTypes = nil
	file_messages_laptop_message_message_proto_depIdxs = nil
}
