// Code generated by protoc-gen-gogo.
// source: proto3_proto/proto3.proto
// DO NOT EDIT!

/*
Package proto3_proto is a generated protocol buffer package.

It is generated from these files:
	proto3_proto/proto3.proto

It has these top-level messages:
	Message
	Nested
	MessageWithMap
*/
package proto3_proto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import testdata "github.com/gogo/protobuf/proto/testdata"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Message_Humour int32

const (
	Message_UNKNOWN     Message_Humour = 0
	Message_PUNS        Message_Humour = 1
	Message_SLAPSTICK   Message_Humour = 2
	Message_BILL_BAILEY Message_Humour = 3
)

var Message_Humour_name = map[int32]string{
	0: "UNKNOWN",
	1: "PUNS",
	2: "SLAPSTICK",
	3: "BILL_BAILEY",
}
var Message_Humour_value = map[string]int32{
	"UNKNOWN":     0,
	"PUNS":        1,
	"SLAPSTICK":   2,
	"BILL_BAILEY": 3,
}

func (x Message_Humour) String() string {
	return proto.EnumName(Message_Humour_name, int32(x))
}
func (Message_Humour) EnumDescriptor() ([]byte, []int) { return fileDescriptorProto3, []int{0, 0} }

type Message struct {
	Name         string                           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hilarity     Message_Humour                   `protobuf:"varint,2,opt,name=hilarity,proto3,enum=proto3_proto.Message_Humour" json:"hilarity,omitempty"`
	HeightInCm   uint32                           `protobuf:"varint,3,opt,name=height_in_cm,proto3" json:"height_in_cm,omitempty"`
	Data         []byte                           `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	ResultCount  int64                            `protobuf:"varint,7,opt,name=result_count,proto3" json:"result_count,omitempty"`
	TrueScotsman bool                             `protobuf:"varint,8,opt,name=true_scotsman,proto3" json:"true_scotsman,omitempty"`
	Score        float32                          `protobuf:"fixed32,9,opt,name=score,proto3" json:"score,omitempty"`
	Key          []uint64                         `protobuf:"varint,5,rep,name=key" json:"key,omitempty"`
	Nested       *Nested                          `protobuf:"bytes,6,opt,name=nested" json:"nested,omitempty"`
	Terrain      map[string]*Nested               `protobuf:"bytes,10,rep,name=terrain" json:"terrain,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	Proto2Field  *testdata.SubDefaults            `protobuf:"bytes,11,opt,name=proto2_field" json:"proto2_field,omitempty"`
	Proto2Value  map[string]*testdata.SubDefaults `protobuf:"bytes,13,rep,name=proto2_value" json:"proto2_value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptorProto3, []int{0} }

func (m *Message) GetNested() *Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *Message) GetTerrain() map[string]*Nested {
	if m != nil {
		return m.Terrain
	}
	return nil
}

func (m *Message) GetProto2Field() *testdata.SubDefaults {
	if m != nil {
		return m.Proto2Field
	}
	return nil
}

func (m *Message) GetProto2Value() map[string]*testdata.SubDefaults {
	if m != nil {
		return m.Proto2Value
	}
	return nil
}

type Nested struct {
	Bunny string `protobuf:"bytes,1,opt,name=bunny,proto3" json:"bunny,omitempty"`
}

func (m *Nested) Reset()                    { *m = Nested{} }
func (m *Nested) String() string            { return proto.CompactTextString(m) }
func (*Nested) ProtoMessage()               {}
func (*Nested) Descriptor() ([]byte, []int) { return fileDescriptorProto3, []int{1} }

type MessageWithMap struct {
	ByteMapping map[bool][]byte `protobuf:"bytes,1,rep,name=byte_mapping" json:"byte_mapping,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *MessageWithMap) Reset()                    { *m = MessageWithMap{} }
func (m *MessageWithMap) String() string            { return proto.CompactTextString(m) }
func (*MessageWithMap) ProtoMessage()               {}
func (*MessageWithMap) Descriptor() ([]byte, []int) { return fileDescriptorProto3, []int{2} }

func (m *MessageWithMap) GetByteMapping() map[bool][]byte {
	if m != nil {
		return m.ByteMapping
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "proto3_proto.Message")
	proto.RegisterType((*Nested)(nil), "proto3_proto.Nested")
	proto.RegisterType((*MessageWithMap)(nil), "proto3_proto.MessageWithMap")
	proto.RegisterEnum("proto3_proto.Message_Humour", Message_Humour_name, Message_Humour_value)
}

var fileDescriptorProto3 = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x53, 0x6d, 0x8f, 0xd2, 0x40,
	0x10, 0x96, 0x77, 0x98, 0x96, 0xb3, 0xd9, 0x60, 0xb2, 0x12, 0x63, 0x10, 0x13, 0x73, 0xf1, 0xa5,
	0x24, 0xf8, 0xe5, 0x62, 0x8c, 0xe6, 0xc0, 0x33, 0x92, 0x03, 0x24, 0xcb, 0xe1, 0xc5, 0x4f, 0xcd,
	0x16, 0x96, 0xd2, 0x48, 0xb7, 0xa4, 0xdd, 0x9a, 0xf0, 0x77, 0xfc, 0x55, 0xfe, 0x1c, 0xb7, 0xbb,
	0xe5, 0xae, 0x77, 0xc1, 0xfb, 0xd4, 0xd9, 0x67, 0x9e, 0x99, 0x67, 0xf6, 0x99, 0x2d, 0x3c, 0xdd,
	0x45, 0xa1, 0x08, 0xdf, 0x3b, 0xea, 0xd3, 0xd3, 0x07, 0x5b, 0x7d, 0x90, 0x99, 0x4f, 0xb5, 0xfb,
	0x9e, 0x2f, 0x36, 0x89, 0x6b, 0x2f, 0xc3, 0xa0, 0xe7, 0x85, 0x5e, 0xc6, 0x75, 0x93, 0xb5, 0x0e,
	0x7a, 0x82, 0xc5, 0x62, 0x45, 0x05, 0x55, 0x81, 0xee, 0xd0, 0xfd, 0x5b, 0x81, 0xda, 0x84, 0xc5,
	0x31, 0xf5, 0x18, 0x42, 0x50, 0xe6, 0x34, 0x60, 0xb8, 0xd0, 0x29, 0x9c, 0x36, 0x88, 0x8a, 0xd1,
	0x19, 0xd4, 0x37, 0xfe, 0x96, 0x46, 0xbe, 0xd8, 0xe3, 0xa2, 0xc4, 0x4f, 0xfa, 0xcf, 0xec, 0xbc,
	0xa8, 0x9d, 0x15, 0xdb, 0xdf, 0x92, 0x20, 0x4c, 0x22, 0x72, 0xc3, 0x46, 0x1d, 0x30, 0x37, 0xcc,
	0xf7, 0x36, 0xc2, 0xf1, 0xb9, 0xb3, 0x0c, 0x70, 0x49, 0x56, 0x37, 0x09, 0x68, 0x6c, 0xc4, 0x87,
	0x41, 0xaa, 0x97, 0x8e, 0x83, 0xcb, 0x32, 0x63, 0x12, 0x15, 0xa3, 0x17, 0x60, 0x46, 0x2c, 0x4e,
	0xb6, 0xc2, 0x59, 0x86, 0x09, 0x17, 0xb8, 0x26, 0x73, 0x25, 0x62, 0x68, 0x6c, 0x98, 0x42, 0xe8,
	0x25, 0x34, 0x45, 0x94, 0x30, 0x27, 0x5e, 0x86, 0x22, 0x0e, 0x28, 0xc7, 0x75, 0xc9, 0xa9, 0x13,
	0x33, 0x05, 0xe7, 0x19, 0x86, 0x5a, 0x50, 0x91, 0xf9, 0x88, 0xe1, 0x86, 0x4c, 0x16, 0x89, 0x3e,
	0x20, 0x0b, 0x4a, 0xbf, 0xd8, 0x1e, 0x57, 0x3a, 0xa5, 0xd3, 0x32, 0x49, 0x43, 0xf4, 0x16, 0xaa,
	0x5c, 0xba, 0xc1, 0x56, 0xb8, 0x2a, 0x89, 0x46, 0xbf, 0x75, 0xf7, 0x76, 0x53, 0x95, 0x23, 0x19,
	0x07, 0x7d, 0x84, 0x9a, 0x60, 0x51, 0x44, 0x7d, 0x8e, 0x41, 0xf6, 0x30, 0xfa, 0xdd, 0xe3, 0x66,
	0x5c, 0x69, 0xd2, 0x05, 0x17, 0xd1, 0x9e, 0x1c, 0x4a, 0xa4, 0x97, 0x7a, 0x5f, 0x7d, 0x67, 0xed,
	0xb3, 0xed, 0x0a, 0x1b, 0x4a, 0xf1, 0x89, 0x7d, 0xd8, 0x8b, 0x3d, 0x4f, 0xdc, 0x2f, 0x6c, 0x4d,
	0xe5, 0x4d, 0x63, 0x62, 0x68, 0xea, 0xd7, 0x94, 0x89, 0x46, 0x37, 0x95, 0xbf, 0xe9, 0x36, 0x61,
	0xb8, 0xa9, 0xc4, 0x5f, 0x1d, 0x17, 0x9f, 0x29, 0xe6, 0x8f, 0x94, 0xa8, 0x07, 0xc8, 0x5a, 0x29,
	0xa4, 0x3d, 0x03, 0x33, 0x3f, 0xdd, 0xc1, 0x12, 0xbd, 0x73, 0x65, 0xc9, 0x6b, 0xa8, 0x68, 0x95,
	0xe2, 0x03, 0x8e, 0x68, 0xca, 0x87, 0xe2, 0x59, 0xa1, 0xbd, 0x00, 0xeb, 0xbe, 0xe4, 0x91, 0xae,
	0x6f, 0xee, 0x76, 0xfd, 0xcf, 0xad, 0x6f, 0xdb, 0x76, 0x3f, 0x43, 0x55, 0xbf, 0x29, 0x64, 0x40,
	0x6d, 0x31, 0xbd, 0x9c, 0x7e, 0xbf, 0x9e, 0x5a, 0x8f, 0x50, 0x1d, 0xca, 0xb3, 0xc5, 0x74, 0x6e,
	0x15, 0x50, 0x13, 0x1a, 0xf3, 0xf1, 0xf9, 0x6c, 0x7e, 0x35, 0x1a, 0x5e, 0x5a, 0x45, 0xf4, 0x18,
	0x8c, 0xc1, 0x68, 0x3c, 0x76, 0x06, 0xe7, 0xa3, 0xf1, 0xc5, 0x4f, 0xab, 0xd4, 0x7d, 0x0e, 0x55,
	0x3d, 0x6c, 0xfa, 0x18, 0xdc, 0x84, 0xf3, 0xc3, 0x3c, 0xfa, 0xd0, 0xfd, 0x53, 0x80, 0x93, 0xcc,
	0xb3, 0x6b, 0xf9, 0xe3, 0x4c, 0xe8, 0x0e, 0x49, 0x73, 0xdc, 0xbd, 0x60, 0x4e, 0x40, 0x77, 0x3b,
	0x9f, 0x7b, 0x92, 0x9f, 0xfa, 0xfc, 0xee, 0xa8, 0xcf, 0x59, 0x8d, 0x3d, 0x90, 0x05, 0x13, 0xcd,
	0xcf, 0xec, 0x76, 0x6f, 0x91, 0xf6, 0x27, 0xb0, 0xee, 0x13, 0xf2, 0xe6, 0xd4, 0xb5, 0x39, 0xad,
	0xbc, 0x39, 0x66, 0xce, 0x05, 0xb7, 0xaa, 0xa5, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xf0,
	0x9a, 0xf4, 0x05, 0x04, 0x00, 0x00,
}
