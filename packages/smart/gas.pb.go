// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gas.proto

package smart

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PaymentType int32

const (
	PaymentType_INVALID          PaymentType = 0
	PaymentType_ContractCaller   PaymentType = 1
	PaymentType_ContractBinder   PaymentType = 2
	PaymentType_EcosystemAddress PaymentType = 3
)

var PaymentType_name = map[int32]string{
	0: "INVALID",
	1: "ContractCaller",
	2: "ContractBinder",
	3: "EcosystemAddress",
}

var PaymentType_value = map[string]int32{
	"INVALID":          0,
	"ContractCaller":   1,
	"ContractBinder":   2,
	"EcosystemAddress": 3,
}

func (x PaymentType) String() string {
	return proto.EnumName(PaymentType_name, int32(x))
}

func (PaymentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df176b4a803aa869, []int{0}
}

type GasScenesType int32

const (
	GasScenesType_Unknown    GasScenesType = 0
	GasScenesType_Reward     GasScenesType = 1
	GasScenesType_Taxes      GasScenesType = 2
	GasScenesType_Direct     GasScenesType = 15
	GasScenesType_Combustion GasScenesType = 16
)

var GasScenesType_name = map[int32]string{
	0:  "Unknown",
	1:  "Reward",
	2:  "Taxes",
	15: "Direct",
	16: "Combustion",
}

var GasScenesType_value = map[string]int32{
	"Unknown":    0,
	"Reward":     1,
	"Taxes":      2,
	"Direct":     15,
	"Combustion": 16,
}

func (x GasScenesType) String() string {
	return proto.EnumName(GasScenesType_name, int32(x))
}

func (GasScenesType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df176b4a803aa869, []int{1}
}

type GasPayAbleType int32

const (
	GasPayAbleType_Invalid GasPayAbleType = 0
	GasPayAbleType_Unable  GasPayAbleType = 1
	GasPayAbleType_Capable GasPayAbleType = 2
)

var GasPayAbleType_name = map[int32]string{
	0: "Invalid",
	1: "Unable",
	2: "Capable",
}

var GasPayAbleType_value = map[string]int32{
	"Invalid": 0,
	"Unable":  1,
	"Capable": 2,
}

func (x GasPayAbleType) String() string {
	return proto.EnumName(GasPayAbleType_name, int32(x))
}

func (GasPayAbleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df176b4a803aa869, []int{2}
}

type FuelType int32

const (
	FuelType_UNKNOWN      FuelType = 0
	FuelType_vmCost_fee   FuelType = 1
	FuelType_storage_fee  FuelType = 2
	FuelType_element_fee  FuelType = 3
	FuelType_expedite_fee FuelType = 4
)

var FuelType_name = map[int32]string{
	0: "UNKNOWN",
	1: "vmCost_fee",
	2: "storage_fee",
	3: "element_fee",
	4: "expedite_fee",
}

var FuelType_value = map[string]int32{
	"UNKNOWN":      0,
	"vmCost_fee":   1,
	"storage_fee":  2,
	"element_fee":  3,
	"expedite_fee": 4,
}

func (x FuelType) String() string {
	return proto.EnumName(FuelType_name, int32(x))
}

func (FuelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df176b4a803aa869, []int{3}
}

func init() {
	proto.RegisterEnum("smart.PaymentType", PaymentType_name, PaymentType_value)
	proto.RegisterEnum("smart.GasScenesType", GasScenesType_name, GasScenesType_value)
	proto.RegisterEnum("smart.GasPayAbleType", GasPayAbleType_name, GasPayAbleType_value)
	proto.RegisterEnum("smart.FuelType", FuelType_name, FuelType_value)
}

func init() { proto.RegisterFile("gas.proto", fileDescriptor_df176b4a803aa869) }

var fileDescriptor_df176b4a803aa869 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0x6d, 0x53, 0x68, 0x59, 0x5a, 0x58, 0xad, 0x7a, 0xf6, 0xbd, 0x96, 0xc0, 0x87, 0x4a,
	0xbd, 0x1b, 0xd3, 0x22, 0xd4, 0xca, 0xa0, 0x16, 0x5a, 0xd4, 0x1c, 0xa2, 0xb1, 0x3d, 0x71, 0x56,
	0xd8, 0xbb, 0xd6, 0xee, 0xf2, 0xc7, 0x6f, 0x91, 0xc7, 0xca, 0x91, 0x63, 0x8e, 0x11, 0xbc, 0x48,
	0x64, 0xa3, 0xa0, 0x1c, 0xe7, 0x37, 0x33, 0xdf, 0x7c, 0x9a, 0x8f, 0x74, 0x33, 0xd0, 0xa3, 0x52,
	0x49, 0x23, 0x59, 0x5b, 0x17, 0xa0, 0x8c, 0xb7, 0x26, 0xbd, 0x05, 0x54, 0x05, 0x0a, 0xb3, 0xac,
	0x4a, 0x64, 0x3d, 0xf2, 0x7e, 0x16, 0xfd, 0x0d, 0x7e, 0xcd, 0x26, 0xd4, 0x62, 0x8c, 0xf4, 0x43,
	0x29, 0x8c, 0x82, 0xc4, 0x84, 0x90, 0xe7, 0xa8, 0xa8, 0xfd, 0x96, 0x8d, 0xb9, 0x48, 0x51, 0x51,
	0x87, 0x7d, 0x26, 0xf4, 0x7b, 0x22, 0x75, 0xa5, 0x0d, 0x16, 0x41, 0x9a, 0x2a, 0xd4, 0x9a, 0xb6,
	0xbc, 0x39, 0xf9, 0x34, 0x05, 0xfd, 0x27, 0x41, 0x81, 0xfa, 0x55, 0x7b, 0x25, 0x36, 0x42, 0xee,
	0x05, 0xb5, 0x18, 0x21, 0x9d, 0xdf, 0xb8, 0x07, 0x95, 0x52, 0x9b, 0x75, 0x49, 0x7b, 0x09, 0x07,
	0xd4, 0xd4, 0xa9, 0xf1, 0x84, 0x2b, 0x4c, 0x0c, 0x1d, 0xb0, 0x3e, 0x21, 0xa1, 0x2c, 0xe2, 0xad,
	0x36, 0x5c, 0x0a, 0x4a, 0xbd, 0x6f, 0xa4, 0x3f, 0x05, 0xbd, 0x80, 0x2a, 0x88, 0x73, 0xbc, 0xba,
	0x15, 0x3b, 0xc8, 0x79, 0x7a, 0x51, 0x5c, 0x09, 0x88, 0x73, 0xa4, 0x76, 0xdd, 0x08, 0xa1, 0x6c,
	0x0a, 0xc7, 0xbb, 0x21, 0x1f, 0x7e, 0x6c, 0x31, 0xbf, 0x7a, 0x88, 0x7e, 0x46, 0xf3, 0x7f, 0x11,
	0xb5, 0xea, 0x03, 0xbb, 0x22, 0x94, 0xda, 0xdc, 0xde, 0x61, 0xbd, 0x35, 0x20, 0x3d, 0x6d, 0xa4,
	0x82, 0x0c, 0x1b, 0xe0, 0xd4, 0x00, 0x73, 0xac, 0x9f, 0xd3, 0x80, 0x16, 0xa3, 0xe4, 0x23, 0x1e,
	0x4a, 0x4c, 0xb9, 0xb9, 0x8c, 0xbc, 0x1b, 0x87, 0x8f, 0x27, 0xd7, 0x3e, 0x9e, 0x5c, 0xfb, 0xf9,
	0xe4, 0xda, 0x0f, 0x67, 0xd7, 0x3a, 0x9e, 0x5d, 0xeb, 0xe9, 0xec, 0x5a, 0xff, 0xbf, 0x64, 0xdc,
	0xdc, 0x6f, 0xe3, 0x51, 0x22, 0x0b, 0x7f, 0x36, 0x0e, 0xd6, 0x43, 0x2e, 0xfd, 0x4c, 0x0e, 0x79,
	0x0c, 0x07, 0xbf, 0x84, 0x64, 0x03, 0x19, 0x6a, 0xbf, 0x09, 0x21, 0xee, 0x34, 0x91, 0x7c, 0x7d,
	0x09, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x5f, 0x34, 0x50, 0x9f, 0x01, 0x00, 0x00,
}