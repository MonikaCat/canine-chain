// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: canine_chain/oracle/feed.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

type Feed struct {
	Owner      string    `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Data       string    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	LastUpdate time.Time `protobuf:"bytes,3,opt,name=last_update,json=lastUpdate,proto3,stdtime" json:"last_update"`
	Name       string    `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Feed) Reset()         { *m = Feed{} }
func (m *Feed) String() string { return proto.CompactTextString(m) }
func (*Feed) ProtoMessage()    {}
func (*Feed) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b6c62baf4208166, []int{0}
}
func (m *Feed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Feed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Feed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Feed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feed.Merge(m, src)
}
func (m *Feed) XXX_Size() int {
	return m.Size()
}
func (m *Feed) XXX_DiscardUnknown() {
	xxx_messageInfo_Feed.DiscardUnknown(m)
}

var xxx_messageInfo_Feed proto.InternalMessageInfo

func (m *Feed) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Feed) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Feed) GetLastUpdate() time.Time {
	if m != nil {
		return m.LastUpdate
	}
	return time.Time{}
}

func (m *Feed) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Feed)(nil), "canine_chain.oracle.Feed")
}

func init() { proto.RegisterFile("canine_chain/oracle/feed.proto", fileDescriptor_3b6c62baf4208166) }

var fileDescriptor_3b6c62baf4208166 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x50, 0x3d, 0x4e, 0xc3, 0x30,
	0x14, 0x8e, 0x21, 0x20, 0x70, 0xb7, 0xd0, 0x21, 0xca, 0xe0, 0x54, 0x4c, 0x5d, 0xb0, 0x05, 0xdc,
	0xa0, 0x12, 0x2c, 0x30, 0x55, 0xb0, 0xb0, 0x54, 0x2f, 0xc9, 0xab, 0x1b, 0x48, 0xec, 0x28, 0x71,
	0x04, 0x9c, 0x81, 0xa5, 0xc7, 0xea, 0xd8, 0x91, 0x09, 0x50, 0x72, 0x11, 0x14, 0x5b, 0x91, 0xba,
	0x7d, 0xcf, 0xdf, 0x9f, 0xf5, 0x51, 0x96, 0x82, 0xca, 0x15, 0xae, 0xd2, 0x0d, 0xe4, 0x4a, 0xe8,
	0x1a, 0xd2, 0x02, 0xc5, 0x1a, 0x31, 0xe3, 0x55, 0xad, 0x8d, 0x0e, 0x2e, 0x0e, 0x79, 0xee, 0xf8,
	0x68, 0x2a, 0xb5, 0xd4, 0x96, 0x17, 0x03, 0x72, 0xd2, 0x28, 0x96, 0x5a, 0xcb, 0x02, 0x85, 0xbd,
	0x92, 0x76, 0x2d, 0x4c, 0x5e, 0x62, 0x63, 0xa0, 0xac, 0x9c, 0xe0, 0xf2, 0x8b, 0x50, 0xff, 0x1e,
	0x31, 0x0b, 0xa6, 0xf4, 0x44, 0xbf, 0x2b, 0xac, 0x43, 0x32, 0x23, 0xf3, 0xf3, 0xa5, 0x3b, 0x82,
	0x80, 0xfa, 0x19, 0x18, 0x08, 0x8f, 0xec, 0xa3, 0xc5, 0xc1, 0x1d, 0x9d, 0x14, 0xd0, 0x98, 0x55,
	0x5b, 0x65, 0x60, 0x30, 0x3c, 0x9e, 0x91, 0xf9, 0xe4, 0x26, 0xe2, 0xae, 0x89, 0x8f, 0x4d, 0xfc,
	0x69, 0x6c, 0x5a, 0x9c, 0xed, 0x7e, 0x62, 0x6f, 0xfb, 0x1b, 0x93, 0x25, 0x1d, 0x8c, 0xcf, 0xd6,
	0x37, 0x44, 0x2b, 0x28, 0x31, 0xf4, 0x5d, 0xf4, 0x80, 0x17, 0x0f, 0xbb, 0x8e, 0x91, 0x7d, 0xc7,
	0xc8, 0x5f, 0xc7, 0xc8, 0xb6, 0x67, 0xde, 0xbe, 0x67, 0xde, 0x77, 0xcf, 0xbc, 0x97, 0x6b, 0x99,
	0x9b, 0x4d, 0x9b, 0xf0, 0x54, 0x97, 0xe2, 0x15, 0xd2, 0x37, 0x28, 0x1e, 0x21, 0x69, 0x84, 0x5b,
	0xe2, 0xca, 0x2d, 0xf5, 0x31, 0x6e, 0x65, 0x3e, 0x2b, 0x6c, 0x92, 0x53, 0xfb, 0x95, 0xdb, 0xff,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x2e, 0x9f, 0xfc, 0x4f, 0x01, 0x00, 0x00,
}

func (m *Feed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Feed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Feed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintFeed(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastUpdate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastUpdate):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintFeed(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintFeed(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintFeed(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFeed(dAtA []byte, offset int, v uint64) int {
	offset -= sovFeed(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Feed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovFeed(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovFeed(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastUpdate)
	n += 1 + l + sovFeed(uint64(l))
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovFeed(uint64(l))
	}
	return n
}

func sovFeed(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFeed(x uint64) (n int) {
	return sovFeed(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Feed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeed
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
			return fmt.Errorf("proto: Feed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Feed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeed
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
				return ErrInvalidLengthFeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeed
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
				return ErrInvalidLengthFeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeed
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
				return ErrInvalidLengthFeed
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastUpdate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeed
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
				return ErrInvalidLengthFeed
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeed
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeed(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeed
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
func skipFeed(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFeed
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
					return 0, ErrIntOverflowFeed
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
					return 0, ErrIntOverflowFeed
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
				return 0, ErrInvalidLengthFeed
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFeed
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFeed
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFeed        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFeed          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFeed = fmt.Errorf("proto: unexpected end of group")
)
