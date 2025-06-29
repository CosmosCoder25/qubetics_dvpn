// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qubetics/lease/v1/lease.proto

package v1

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	v1 "github.com/qubetics/qubetics-blockchain/v2/types/v1"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Lease struct {
	ID                 uint64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProvAddress        string                `protobuf:"bytes,2,opt,name=prov_address,json=provAddress,proto3" json:"prov_address,omitempty"`
	NodeAddress        string                `protobuf:"bytes,3,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
	Price              v1.Price              `protobuf:"bytes,4,opt,name=price,proto3" json:"price"`
	Hours              int64                 `protobuf:"varint,5,opt,name=hours,proto3" json:"hours,omitempty"`
	MaxHours           int64                 `protobuf:"varint,6,opt,name=max_hours,json=maxHours,proto3" json:"max_hours,omitempty"`
	RenewalPricePolicy v1.RenewalPricePolicy `protobuf:"varint,7,opt,name=renewal_price_policy,json=renewalPricePolicy,proto3,enum=qubetics.types.v1.RenewalPricePolicy" json:"renewal_price_policy,omitempty"`
	StartAt            time.Time             `protobuf:"bytes,8,opt,name=start_at,json=startAt,proto3,stdtime" json:"start_at"`
}

func (m *Lease) Reset()         { *m = Lease{} }
func (m *Lease) String() string { return proto.CompactTextString(m) }
func (*Lease) ProtoMessage()    {}
func (*Lease) Descriptor() ([]byte, []int) {
	return fileDescriptor_a230a2afd1bdf1dd, []int{0}
}
func (m *Lease) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Lease.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Lease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lease.Merge(m, src)
}
func (m *Lease) XXX_Size() int {
	return m.Size()
}
func (m *Lease) XXX_DiscardUnknown() {
	xxx_messageInfo_Lease.DiscardUnknown(m)
}

var xxx_messageInfo_Lease proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Lease)(nil), "qubetics.lease.v1.Lease")
}

func init() { proto.RegisterFile("qubetics/lease/v1/lease.proto", fileDescriptor_a230a2afd1bdf1dd) }

var fileDescriptor_a230a2afd1bdf1dd = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x8a, 0xd3, 0x40,
	0x18, 0xc7, 0x33, 0xd9, 0xb6, 0xdb, 0x9d, 0x15, 0xc1, 0x50, 0x24, 0x54, 0x4c, 0xa3, 0x20, 0xe4,
	0xe2, 0x0c, 0xad, 0xde, 0x3c, 0xc8, 0x16, 0x0f, 0x0a, 0x1e, 0x96, 0x20, 0x08, 0x82, 0x84, 0x49,
	0x32, 0xa6, 0x83, 0x49, 0x27, 0xce, 0x4c, 0x62, 0xf7, 0x2d, 0xf6, 0x31, 0x7c, 0x94, 0x1e, 0xf7,
	0xe8, 0x69, 0xd5, 0xf4, 0xe2, 0x63, 0xc8, 0xcc, 0x34, 0xf5, 0xd0, 0xbd, 0x7d, 0xf3, 0xff, 0xff,
	0x26, 0xdf, 0xf7, 0xe5, 0x3f, 0xf0, 0xf1, 0xb7, 0x26, 0xa5, 0x8a, 0x65, 0x12, 0x97, 0x94, 0x48,
	0x8a, 0xdb, 0xb9, 0x2d, 0x50, 0x2d, 0xb8, 0xe2, 0xde, 0x83, 0xde, 0x46, 0x56, 0x6d, 0xe7, 0xd3,
	0x49, 0xc1, 0x0b, 0x6e, 0x5c, 0xac, 0x2b, 0x0b, 0x4e, 0x67, 0x05, 0xe7, 0x45, 0x49, 0xb1, 0x39,
	0xa5, 0xcd, 0x17, 0xac, 0x58, 0x45, 0xa5, 0x22, 0x55, 0xbd, 0x07, 0xfe, 0x37, 0x52, 0x57, 0x35,
	0x95, 0xba, 0x51, 0x2d, 0x58, 0x46, 0xfb, 0xfb, 0xc7, 0xb6, 0xa0, 0x6b, 0xfa, 0x9d, 0x94, 0x16,
	0x78, 0xfa, 0xd7, 0x85, 0xc3, 0xf7, 0x7a, 0x06, 0xef, 0x21, 0x74, 0x59, 0xee, 0x83, 0x10, 0x44,
	0x83, 0xe5, 0xa8, 0xbb, 0x9d, 0xb9, 0xef, 0xde, 0xc4, 0x2e, 0xcb, 0xbd, 0x27, 0xf0, 0x5e, 0x2d,
	0x78, 0x9b, 0x90, 0x3c, 0x17, 0x54, 0x4a, 0xdf, 0x0d, 0x41, 0x74, 0x16, 0x9f, 0x6b, 0xed, 0xc2,
	0x4a, 0x1a, 0x59, 0xf3, 0x9c, 0x1e, 0x90, 0x13, 0x8b, 0x68, 0xad, 0x47, 0x5e, 0xc2, 0xa1, 0x99,
	0xcb, 0x1f, 0x84, 0x20, 0x3a, 0x5f, 0xf8, 0xe8, 0xf0, 0x07, 0xcc, 0x60, 0xa8, 0x9d, 0xa3, 0x4b,
	0xed, 0x2f, 0x07, 0xdb, 0xdb, 0x99, 0x13, 0x5b, 0xd8, 0x9b, 0xc0, 0xe1, 0x8a, 0x37, 0x42, 0xfa,
	0xc3, 0x10, 0x44, 0x27, 0xb1, 0x3d, 0x78, 0x8f, 0xe0, 0x59, 0x45, 0x36, 0x89, 0x75, 0x46, 0xc6,
	0x19, 0x57, 0x64, 0xf3, 0xd6, 0x98, 0x1f, 0xe1, 0x64, 0xbf, 0x61, 0x62, 0xbe, 0x91, 0xd4, 0xbc,
	0x64, 0xd9, 0x95, 0x7f, 0x1a, 0x82, 0xe8, 0xfe, 0xe2, 0xd9, 0x1d, 0x7d, 0x63, 0x8b, 0x9b, 0xf6,
	0x97, 0x06, 0x8e, 0x3d, 0x71, 0xa4, 0x79, 0xaf, 0xe1, 0x58, 0x2a, 0x22, 0x54, 0x42, 0x94, 0x3f,
	0x36, 0x4b, 0x4c, 0x91, 0x4d, 0x07, 0xf5, 0xe9, 0xa0, 0x0f, 0x7d, 0x3a, 0xcb, 0xb1, 0x5e, 0xe3,
	0xfa, 0xd7, 0x0c, 0xc4, 0xa7, 0xe6, 0xd6, 0x85, 0x5a, 0x7e, 0xde, 0xfe, 0x09, 0x9c, 0x1f, 0x5d,
	0xe0, 0x6c, 0xbb, 0x00, 0xdc, 0x74, 0x01, 0xf8, 0xdd, 0x05, 0xe0, 0x7a, 0x17, 0x38, 0x37, 0xbb,
	0xc0, 0xf9, 0xb9, 0x0b, 0x9c, 0x4f, 0xaf, 0x0a, 0xa6, 0x56, 0x4d, 0x8a, 0x32, 0x5e, 0xe1, 0x43,
	0x70, 0x7d, 0xf1, 0x3c, 0x2d, 0x79, 0xf6, 0x35, 0x5b, 0x11, 0xb6, 0xc6, 0xed, 0x02, 0x6f, 0xf6,
	0x2f, 0xab, 0xcf, 0x35, 0x1d, 0x99, 0x29, 0x5e, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x83, 0xff,
	0x09, 0xd3, 0x7b, 0x02, 0x00, 0x00,
}

func (m *Lease) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lease) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Lease) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.StartAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintLease(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x42
	if m.RenewalPricePolicy != 0 {
		i = encodeVarintLease(dAtA, i, uint64(m.RenewalPricePolicy))
		i--
		dAtA[i] = 0x38
	}
	if m.MaxHours != 0 {
		i = encodeVarintLease(dAtA, i, uint64(m.MaxHours))
		i--
		dAtA[i] = 0x30
	}
	if m.Hours != 0 {
		i = encodeVarintLease(dAtA, i, uint64(m.Hours))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLease(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintLease(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ProvAddress) > 0 {
		i -= len(m.ProvAddress)
		copy(dAtA[i:], m.ProvAddress)
		i = encodeVarintLease(dAtA, i, uint64(len(m.ProvAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintLease(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLease(dAtA []byte, offset int, v uint64) int {
	offset -= sovLease(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Lease) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovLease(uint64(m.ID))
	}
	l = len(m.ProvAddress)
	if l > 0 {
		n += 1 + l + sovLease(uint64(l))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovLease(uint64(l))
	}
	l = m.Price.Size()
	n += 1 + l + sovLease(uint64(l))
	if m.Hours != 0 {
		n += 1 + sovLease(uint64(m.Hours))
	}
	if m.MaxHours != 0 {
		n += 1 + sovLease(uint64(m.MaxHours))
	}
	if m.RenewalPricePolicy != 0 {
		n += 1 + sovLease(uint64(m.RenewalPricePolicy))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartAt)
	n += 1 + l + sovLease(uint64(l))
	return n
}

func sovLease(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLease(x uint64) (n int) {
	return sovLease(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Lease) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLease
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Lease: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lease: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProvAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProvAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hours", wireType)
			}
			m.Hours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hours |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxHours", wireType)
			}
			m.MaxHours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxHours |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RenewalPricePolicy", wireType)
			}
			m.RenewalPricePolicy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RenewalPricePolicy |= v1.RenewalPricePolicy(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLease
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLease
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLease
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.StartAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLease(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLease
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLease(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLease
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLease
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLease
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLease
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLease
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLease
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLease        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLease          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLease = fmt.Errorf("proto: unexpected end of group")
)
