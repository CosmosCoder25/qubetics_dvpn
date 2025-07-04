// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qubetics/provider/v1/provider.proto

package v1

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Provider struct {
	Address     string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Identity    string `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	Website     string `protobuf:"bytes,4,opt,name=website,proto3" json:"website,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *Provider) Reset()         { *m = Provider{} }
func (m *Provider) String() string { return proto.CompactTextString(m) }
func (*Provider) ProtoMessage()    {}
func (*Provider) Descriptor() ([]byte, []int) {
	return fileDescriptor_91ca446284c6dc76, []int{0}
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
	proto.RegisterType((*Provider)(nil), "qubetics.provider.v1.Provider")
}

func init() {
	proto.RegisterFile("qubetics/provider/v1/provider.proto", fileDescriptor_91ca446284c6dc76)
}

var fileDescriptor_91ca446284c6dc76 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3f, 0x4e, 0xc3, 0x30,
	0x14, 0xc6, 0x6d, 0x28, 0x50, 0xcc, 0x66, 0x75, 0xb0, 0x3a, 0x58, 0x15, 0x2c, 0x2c, 0xc4, 0x0a,
	0xec, 0x0c, 0x9c, 0x00, 0x31, 0x32, 0x11, 0xff, 0x51, 0xfa, 0x04, 0x8d, 0x83, 0xed, 0x1a, 0x7a,
	0x8b, 0x1e, 0x83, 0xa3, 0x74, 0xec, 0xc8, 0x08, 0xc9, 0x45, 0x50, 0xdd, 0x24, 0x74, 0xfb, 0x7d,
	0xdf, 0xfb, 0xe9, 0x49, 0xef, 0x91, 0xab, 0xf7, 0xa5, 0x34, 0x01, 0x94, 0x17, 0xb5, 0xb3, 0x11,
	0xb4, 0x71, 0x22, 0xe6, 0x03, 0x67, 0xb5, 0xb3, 0xc1, 0xd2, 0x49, 0x2f, 0x65, 0xc3, 0x20, 0xe6,
	0xd3, 0x49, 0x69, 0x4b, 0x9b, 0x04, 0xb1, 0xa3, 0xbd, 0x7b, 0xb9, 0xc6, 0x64, 0xfc, 0xd8, 0x59,
	0x94, 0x91, 0xb3, 0x42, 0x6b, 0x67, 0xbc, 0x67, 0x78, 0x86, 0xaf, 0xcf, 0x9f, 0xfa, 0x48, 0x29,
	0x19, 0x55, 0xc5, 0xc2, 0xb0, 0xa3, 0x54, 0x27, 0xa6, 0x53, 0x32, 0x06, 0x6d, 0xaa, 0x00, 0x61,
	0xc5, 0x8e, 0x53, 0x3f, 0xe4, 0xdd, 0xa6, 0x0f, 0x23, 0x3d, 0x04, 0xc3, 0x46, 0xfb, 0x4d, 0x5d,
	0xa4, 0x33, 0x72, 0xa1, 0x8d, 0x57, 0x0e, 0xea, 0x00, 0xb6, 0x62, 0x27, 0x69, 0x7a, 0x58, 0x3d,
	0xbc, 0x6c, 0x7e, 0x39, 0xfa, 0x6a, 0x38, 0xda, 0x34, 0x1c, 0x6f, 0x1b, 0x8e, 0x7f, 0x1a, 0x8e,
	0xd7, 0x2d, 0x47, 0xdb, 0x96, 0xa3, 0xef, 0x96, 0xa3, 0xe7, 0xfb, 0x12, 0xc2, 0x7c, 0x29, 0x33,
	0x65, 0x17, 0x62, 0x78, 0x48, 0x0f, 0x37, 0xf2, 0xcd, 0xaa, 0x57, 0x35, 0x2f, 0xa0, 0x12, 0xf1,
	0x56, 0x7c, 0xfe, 0x7f, 0x2a, 0xac, 0x6a, 0xe3, 0x45, 0xcc, 0xe5, 0x69, 0xba, 0xfd, 0xee, 0x2f,
	0x00, 0x00, 0xff, 0xff, 0x1c, 0x9a, 0x3a, 0x06, 0x4e, 0x01, 0x00, 0x00,
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
