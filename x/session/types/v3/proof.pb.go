// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qubetics/session/v3/proof.proto

package v3

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

type Proof struct {
	ID            uint64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DownloadBytes cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=download_bytes,json=downloadBytes,proto3,customtype=cosmossdk.io/math.Int" json:"download_bytes"`
	UploadBytes   cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=upload_bytes,json=uploadBytes,proto3,customtype=cosmossdk.io/math.Int" json:"upload_bytes"`
	Duration      time.Duration         `protobuf:"bytes,4,opt,name=duration,proto3,stdduration" json:"duration"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c2ff0f3c42d4139, []int{0}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return m.Size()
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Proof)(nil), "qubetics.session.v3.Proof")
}

func init() { proto.RegisterFile("qubetics/session/v3/proof.proto", fileDescriptor_4c2ff0f3c42d4139) }

var fileDescriptor_4c2ff0f3c42d4139 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0xe3, 0x50, 0xaa, 0x92, 0x02, 0x43, 0xf8, 0xa3, 0x52, 0x09, 0xa7, 0x62, 0xea, 0x82,
	0x2d, 0x35, 0x33, 0x02, 0x45, 0x5d, 0xba, 0xa1, 0x8e, 0x2c, 0x55, 0x1c, 0xa7, 0xa9, 0xd5, 0x36,
	0x17, 0x6a, 0x27, 0xd0, 0xb7, 0x60, 0xe4, 0x11, 0x78, 0x94, 0x8e, 0x1d, 0x11, 0x43, 0x81, 0x74,
	0xe1, 0x31, 0x50, 0x92, 0x26, 0x62, 0x64, 0x3b, 0xfb, 0x7e, 0xdf, 0xa7, 0x3b, 0x9d, 0x61, 0x3d,
	0xc6, 0xcc, 0x57, 0xc2, 0x93, 0x54, 0xfa, 0x52, 0x0a, 0x08, 0x69, 0x62, 0xd3, 0x68, 0x01, 0x30,
	0x26, 0xd1, 0x02, 0x14, 0x98, 0x27, 0x65, 0x80, 0xec, 0x02, 0x24, 0xb1, 0xdb, 0xa7, 0x01, 0x04,
	0x90, 0xf7, 0x69, 0x56, 0x15, 0xd1, 0x36, 0x0e, 0x00, 0x82, 0x99, 0x4f, 0xf3, 0x17, 0x8b, 0xc7,
	0x94, 0xc7, 0x0b, 0x57, 0x65, 0x48, 0xfe, 0x73, 0xf5, 0x83, 0x8c, 0xfd, 0xfb, 0x4c, 0x6d, 0x9e,
	0x1b, 0xba, 0xe0, 0x2d, 0xd4, 0x41, 0xdd, 0x9a, 0x53, 0x4f, 0x37, 0x96, 0x3e, 0xe8, 0x0f, 0x75,
	0xc1, 0xcd, 0xbe, 0x71, 0xcc, 0xe1, 0x29, 0x9c, 0x81, 0xcb, 0x47, 0x6c, 0xa9, 0x7c, 0xd9, 0xd2,
	0x3b, 0xa8, 0x7b, 0xe0, 0x5c, 0xae, 0x36, 0x96, 0xf6, 0xb1, 0xb1, 0xce, 0x3c, 0x90, 0x73, 0x90,
	0x92, 0x4f, 0x89, 0x00, 0x3a, 0x77, 0xd5, 0x84, 0x0c, 0x42, 0x35, 0x3c, 0x2a, 0x21, 0x27, 0x63,
	0xcc, 0x3b, 0xe3, 0x30, 0x8e, 0xfe, 0x38, 0xf6, 0xfe, 0xe3, 0x68, 0x16, 0x48, 0x61, 0xb8, 0x35,
	0x1a, 0xe5, 0xec, 0xad, 0x5a, 0x07, 0x75, 0x9b, 0xbd, 0x0b, 0x52, 0x2c, 0x47, 0xca, 0xe5, 0x48,
	0x7f, 0x17, 0x70, 0x1a, 0x99, 0xf8, 0xf5, 0xd3, 0x42, 0xc3, 0x0a, 0x72, 0x46, 0xab, 0x6f, 0xac,
	0xbd, 0xa5, 0x58, 0x5b, 0xa5, 0x18, 0xad, 0x53, 0x8c, 0xbe, 0x52, 0x8c, 0x5e, 0xb6, 0x58, 0x5b,
	0x6f, 0xb1, 0xf6, 0xbe, 0xc5, 0xda, 0xc3, 0x4d, 0x20, 0xd4, 0x24, 0x66, 0xc4, 0x83, 0x39, 0xad,
	0x6e, 0x50, 0x16, 0xd7, 0x6c, 0x06, 0xde, 0xd4, 0x9b, 0xb8, 0x22, 0xa4, 0x49, 0x8f, 0x3e, 0x57,
	0xc7, 0x51, 0xcb, 0xc8, 0x97, 0x34, 0xb1, 0x59, 0x3d, 0x9f, 0xc3, 0xfe, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0x19, 0x90, 0xb7, 0xc7, 0xc0, 0x01, 0x00, 0x00,
}

func (m *Proof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Duration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintProof(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	{
		size := m.UploadBytes.Size()
		i -= size
		if _, err := m.UploadBytes.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProof(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.DownloadBytes.Size()
		i -= size
		if _, err := m.DownloadBytes.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProof(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.ID != 0 {
		i = encodeVarintProof(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProof(dAtA []byte, offset int, v uint64) int {
	offset -= sovProof(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Proof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovProof(uint64(m.ID))
	}
	l = m.DownloadBytes.Size()
	n += 1 + l + sovProof(uint64(l))
	l = m.UploadBytes.Size()
	n += 1 + l + sovProof(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovProof(uint64(l))
	return n
}

func sovProof(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProof(x uint64) (n int) {
	return sovProof(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Proof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProof
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
			return fmt.Errorf("proto: Proof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return fmt.Errorf("proto: wrong wireType = %d for field DownloadBytes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DownloadBytes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UploadBytes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UploadBytes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProof
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
func skipProof(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
				return 0, ErrInvalidLengthProof
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProof
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProof
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProof        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProof          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProof = fmt.Errorf("proto: unexpected end of group")
)
