// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qubetics/provider/v2/provider.proto

package v2

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

type Provider struct {
	Address     string    `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Name        string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Identity    string    `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	Website     string    `protobuf:"bytes,4,opt,name=website,proto3" json:"website,omitempty"`
	Description string    `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Status      v1.Status `protobuf:"varint,6,opt,name=status,proto3,enum=qubetics.types.v1.Status" json:"status,omitempty"`
	StatusAt    time.Time `protobuf:"bytes,7,opt,name=status_at,json=statusAt,proto3,stdtime" json:"status_at"`
}

func (m *Provider) Reset()         { *m = Provider{} }
func (m *Provider) String() string { return proto.CompactTextString(m) }
func (*Provider) ProtoMessage()    {}
func (*Provider) Descriptor() ([]byte, []int) {
	return fileDescriptor_1df231b4f455f513, []int{0}
}
func (m *Provider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Provider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Provider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Provider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Provider.Merge(m, src)
}
func (m *Provider) XXX_Size() int {
	return m.Size()
}
func (m *Provider) XXX_DiscardUnknown() {
	xxx_messageInfo_Provider.DiscardUnknown(m)
}

var xxx_messageInfo_Provider proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Provider)(nil), "qubetics.provider.v2.Provider")
}

func init() {
	proto.RegisterFile("qubetics/provider/v2/provider.proto", fileDescriptor_1df231b4f455f513)
}

var fileDescriptor_1df231b4f455f513 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x51, 0xbb, 0x72, 0xe2, 0x30,
	0x14, 0xb5, 0x58, 0x16, 0x8c, 0x98, 0xd9, 0x42, 0x43, 0xa1, 0x75, 0x21, 0x3c, 0xbb, 0x0d, 0xcd,
	0x4a, 0x83, 0xb7, 0xcf, 0x0c, 0x7c, 0x41, 0x86, 0xa4, 0x4a, 0x93, 0xf8, 0xa1, 0x18, 0x4d, 0xb0,
	0xe5, 0x58, 0xb2, 0x13, 0x7e, 0x20, 0x35, 0x9f, 0x91, 0x4f, 0xa1, 0xa4, 0x4c, 0x95, 0x87, 0xf9,
	0x91, 0x0c, 0xf2, 0x23, 0xe9, 0xce, 0xb9, 0xe7, 0xe8, 0xde, 0xab, 0x73, 0xe1, 0xdf, 0xfb, 0x22,
	0xe0, 0x5a, 0x84, 0x8a, 0x65, 0xb9, 0x2c, 0x45, 0xc4, 0x73, 0x56, 0x7a, 0x1d, 0xa6, 0x59, 0x2e,
	0xb5, 0x44, 0x93, 0xd6, 0x44, 0x3b, 0xa1, 0xf4, 0x9c, 0x49, 0x2c, 0x63, 0x69, 0x0c, 0xec, 0x84,
	0x6a, 0xaf, 0x33, 0x8d, 0xa5, 0x8c, 0x37, 0x9c, 0x19, 0x16, 0x14, 0xb7, 0x4c, 0x8b, 0x84, 0x2b,
	0xed, 0x27, 0x59, 0x63, 0x20, 0xdd, 0x44, 0xbd, 0xcd, 0xb8, 0x62, 0xe5, 0x9c, 0x29, 0xed, 0xeb,
	0x42, 0xd5, 0xfa, 0x9f, 0xa7, 0x1e, 0xb4, 0xcf, 0x9b, 0x31, 0x08, 0xc3, 0xa1, 0x1f, 0x45, 0x39,
	0x57, 0x0a, 0x03, 0x17, 0xcc, 0x46, 0xab, 0x96, 0x22, 0x04, 0xfb, 0xa9, 0x9f, 0x70, 0xdc, 0x33,
	0x65, 0x83, 0x91, 0x03, 0x6d, 0x11, 0xf1, 0x54, 0x0b, 0xbd, 0xc5, 0x3f, 0x4c, 0xbd, 0xe3, 0xa7,
	0x4e, 0x0f, 0x3c, 0x50, 0x42, 0x73, 0xdc, 0xaf, 0x3b, 0x35, 0x14, 0xb9, 0x70, 0x1c, 0x71, 0x15,
	0xe6, 0x22, 0xd3, 0x42, 0xa6, 0xf8, 0xa7, 0x51, 0xbf, 0x97, 0xd0, 0x1c, 0x0e, 0xea, 0x15, 0xf1,
	0xc0, 0x05, 0xb3, 0x5f, 0xde, 0x6f, 0xda, 0x05, 0x62, 0xfe, 0x40, 0xcb, 0x39, 0xbd, 0x30, 0x86,
	0x55, 0x63, 0x44, 0x0b, 0x38, 0xaa, 0xd1, 0xb5, 0xaf, 0xf1, 0xd0, 0x05, 0xb3, 0xb1, 0xe7, 0xd0,
	0x3a, 0x1a, 0xda, 0x46, 0x43, 0x2f, 0xdb, 0x68, 0x96, 0xf6, 0xfe, 0x75, 0x6a, 0xed, 0xde, 0xa6,
	0x60, 0x65, 0xd7, 0xcf, 0x16, 0x7a, 0x79, 0xb3, 0xff, 0x20, 0xd6, 0x73, 0x45, 0xac, 0x7d, 0x45,
	0xc0, 0xa1, 0x22, 0xe0, 0xbd, 0x22, 0x60, 0x77, 0x24, 0xd6, 0xe1, 0x48, 0xac, 0x97, 0x23, 0xb1,
	0xae, 0xce, 0x62, 0xa1, 0xd7, 0x45, 0x40, 0x43, 0x99, 0xb0, 0x2e, 0xd5, 0x16, 0xfc, 0x0b, 0x36,
	0x32, 0xbc, 0x0b, 0xd7, 0xbe, 0x48, 0x4f, 0x27, 0x7d, 0xfc, 0x3a, 0x70, 0x93, 0xbb, 0x17, 0x0c,
	0xcc, 0x26, 0xff, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x76, 0x2c, 0x22, 0x30, 0x05, 0x02, 0x00,
	0x00,
}

func (m *Provider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Provider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Provider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.StatusAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StatusAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintProvider(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x3a
	if m.Status != 0 {
		i = encodeVarintProvider(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintProvider(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProvider(dAtA []byte, offset int, v uint64) int {
	offset -= sovProvider(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Provider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProvider(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovProvider(uint64(m.Status))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StatusAt)
	n += 1 + l + sovProvider(uint64(l))
	return n
}

func sovProvider(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProvider(x uint64) (n int) {
	return sovProvider(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Provider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProvider
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
			return fmt.Errorf("proto: Provider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Provider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= v1.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvider
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
				return ErrInvalidLengthProvider
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProvider
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.StatusAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProvider(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProvider
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
func skipProvider(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProvider
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
					return 0, ErrIntOverflowProvider
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
					return 0, ErrIntOverflowProvider
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
				return 0, ErrInvalidLengthProvider
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProvider
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProvider
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProvider        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProvider          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProvider = fmt.Errorf("proto: unexpected end of group")
)
