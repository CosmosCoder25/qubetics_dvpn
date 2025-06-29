// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qubetics/types/v1/status.proto

package v1

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type Status int32

const (
	StatusUnspecified     Status = 0
	StatusActive          Status = 1
	StatusInactivePending Status = 2
	StatusInactive        Status = 3
)

var Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "STATUS_ACTIVE",
	2: "STATUS_INACTIVE_PENDING",
	3: "STATUS_INACTIVE",
}

var Status_value = map[string]int32{
	"STATUS_UNSPECIFIED":      0,
	"STATUS_ACTIVE":           1,
	"STATUS_INACTIVE_PENDING": 2,
	"STATUS_INACTIVE":         3,
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b19e02092cd15eeb, []int{0}
}

func init() {
	proto.RegisterEnum("qubetics.types.v1.Status", Status_name, Status_value)
}

func init() { proto.RegisterFile("qubetics/types/v1/status.proto", fileDescriptor_b19e02092cd15eeb) }

var fileDescriptor_b19e02092cd15eeb = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x2c, 0x4d, 0x4a,
	0x2d, 0xc9, 0x4c, 0x2e, 0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0x33, 0xd4, 0x2f, 0x2e,
	0x49, 0x2c, 0x29, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0xc9, 0xeb, 0x81,
	0xe5, 0xf5, 0xca, 0x0c, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xb2, 0xfa, 0x20, 0x16, 0x44,
	0xa1, 0xd6, 0x5e, 0x46, 0x2e, 0xb6, 0x60, 0xb0, 0x4e, 0x21, 0x5d, 0x2e, 0xa1, 0xe0, 0x10, 0xc7,
	0x90, 0xd0, 0xe0, 0xf8, 0x50, 0xbf, 0xe0, 0x00, 0x57, 0x67, 0x4f, 0x37, 0x4f, 0x57, 0x17, 0x01,
	0x06, 0x29, 0xd1, 0xae, 0xb9, 0x0a, 0x82, 0x10, 0x35, 0xa1, 0x79, 0xc5, 0x05, 0xa9, 0xc9, 0x99,
	0x69, 0x99, 0xa9, 0x29, 0x42, 0xca, 0x5c, 0xbc, 0x50, 0xe5, 0x8e, 0xce, 0x21, 0x9e, 0x61, 0xae,
	0x02, 0x8c, 0x52, 0x02, 0x5d, 0x73, 0x15, 0x78, 0x20, 0x2a, 0x1d, 0x93, 0x4b, 0x32, 0xcb, 0x52,
	0x85, 0xcc, 0xb8, 0xc4, 0xa1, 0x8a, 0x3c, 0xfd, 0x20, 0xca, 0xe2, 0x03, 0x5c, 0xfd, 0x5c, 0x3c,
	0xfd, 0xdc, 0x05, 0x98, 0xa4, 0x24, 0xbb, 0xe6, 0x2a, 0x88, 0x42, 0x94, 0x7b, 0xe6, 0x25, 0x82,
	0x35, 0x04, 0xa4, 0xe6, 0xa5, 0x64, 0xe6, 0xa5, 0x0b, 0xa9, 0x73, 0xf1, 0xa3, 0xe9, 0x13, 0x60,
	0x96, 0x12, 0xea, 0x9a, 0xab, 0xc0, 0x87, 0xaa, 0xde, 0x29, 0xf2, 0xc4, 0x43, 0x39, 0x86, 0x0b,
	0x0f, 0xe5, 0x18, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6,
	0x17, 0x8f, 0xe4, 0x18, 0x26, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x86, 0x28, 0xe3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x78, 0xc8,
	0xc1, 0x18, 0xba, 0x49, 0x39, 0xf9, 0xc9, 0xd9, 0xc9, 0x19, 0x89, 0x99, 0x79, 0xfa, 0x65, 0x46,
	0xf0, 0x00, 0x4d, 0x62, 0x03, 0x87, 0x90, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x13, 0x7b, 0x46,
	0x8b, 0x6c, 0x01, 0x00, 0x00,
}
