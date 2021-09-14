// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: botanydist/botanydist.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Params struct {
	Active  bool     `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty" yaml:"active"`
	Periods []Period `protobuf:"bytes,2,rep,name=periods,proto3" json:"periods" yaml:"periods"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_210c77e58829b608, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Params) GetPeriods() []Period {
	if m != nil {
		return m.Periods
	}
	return nil
}

type Period struct {
	Start     time.Time                              `protobuf:"bytes,1,opt,name=start,proto3,stdtime" json:"start" yaml:"start"`
	End       time.Time                              `protobuf:"bytes,2,opt,name=end,proto3,stdtime" json:"end" yaml:"end"`
	Inflation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=inflation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation" yaml:"inflation"`
}

func (m *Period) Reset()         { *m = Period{} }
func (m *Period) String() string { return proto.CompactTextString(m) }
func (*Period) ProtoMessage()    {}
func (*Period) Descriptor() ([]byte, []int) {
	return fileDescriptor_210c77e58829b608, []int{1}
}
func (m *Period) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Period) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Period.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Period) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Period.Merge(m, src)
}
func (m *Period) XXX_Size() int {
	return m.Size()
}
func (m *Period) XXX_DiscardUnknown() {
	xxx_messageInfo_Period.DiscardUnknown(m)
}

var xxx_messageInfo_Period proto.InternalMessageInfo

func (m *Period) GetStart() time.Time {
	if m != nil {
		return m.Start
	}
	return time.Time{}
}

func (m *Period) GetEnd() time.Time {
	if m != nil {
		return m.End
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Params)(nil), "botany.botanydist.Params")
	proto.RegisterType((*Period)(nil), "botany.botanydist.Period")
}

func init() { proto.RegisterFile("botanydist/botanydist.proto", fileDescriptor_210c77e58829b608) }

var fileDescriptor_210c77e58829b608 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x31, 0x6f, 0x9b, 0x40,
	0x14, 0xc7, 0x39, 0x5b, 0x75, 0xeb, 0x73, 0x5b, 0xd5, 0xa8, 0xaa, 0xa8, 0x2b, 0x81, 0xc5, 0x50,
	0xd1, 0xa1, 0x77, 0xb2, 0xbb, 0x75, 0x44, 0x9e, 0xda, 0xc5, 0x42, 0x9d, 0x3a, 0xe5, 0x80, 0x33,
	0xb9, 0x04, 0x38, 0xc4, 0x9d, 0x2d, 0xb3, 0xe5, 0x23, 0xf8, 0x63, 0x79, 0xf4, 0x18, 0x65, 0x20,
	0x91, 0xfd, 0x0d, 0x3c, 0x67, 0x88, 0xe0, 0x20, 0xb6, 0x94, 0x21, 0x13, 0x8f, 0xf7, 0xfe, 0xef,
	0xf7, 0xfe, 0x7f, 0x04, 0xfc, 0xe6, 0x73, 0x49, 0xd2, 0x22, 0x64, 0x42, 0xe2, 0x53, 0x89, 0xb2,
	0x9c, 0x4b, 0xae, 0x0f, 0x55, 0x07, 0x9d, 0x06, 0xa3, 0xcf, 0x11, 0x8f, 0x78, 0x3d, 0xc5, 0x55,
	0xa5, 0x84, 0x23, 0x2b, 0xe2, 0x3c, 0x8a, 0x29, 0xae, 0xdf, 0xfc, 0xe5, 0x02, 0x4b, 0x96, 0x50,
	0x21, 0x49, 0x92, 0x35, 0x02, 0x33, 0xe0, 0x22, 0xe1, 0x02, 0xfb, 0x44, 0x50, 0xbc, 0x9a, 0xf8,
	0x54, 0x92, 0x09, 0x0e, 0x38, 0x4b, 0xd5, 0xdc, 0xbe, 0x01, 0xb0, 0x37, 0x27, 0x39, 0x49, 0x84,
	0xfe, 0x03, 0xf6, 0x48, 0x20, 0xd9, 0x8a, 0x1a, 0x60, 0x0c, 0x9c, 0x77, 0xee, 0xf0, 0x58, 0x5a,
	0x1f, 0x0a, 0x92, 0xc4, 0xbf, 0x6d, 0xd5, 0xb7, 0xbd, 0x46, 0xa0, 0xff, 0x85, 0x6f, 0x33, 0x9a,
	0x33, 0x1e, 0x0a, 0xa3, 0x33, 0xee, 0x3a, 0x83, 0xe9, 0x57, 0xf4, 0xc2, 0x31, 0x9a, 0xd7, 0x0a,
	0xf7, 0xcb, 0xb6, 0xb4, 0xb4, 0x63, 0x69, 0x7d, 0x54, 0xa8, 0x66, 0xcf, 0xf6, 0x5a, 0x82, 0xfd,
	0x58, 0x59, 0xa8, 0x6b, 0xfd, 0x0f, 0x7c, 0x23, 0x24, 0xc9, 0x65, 0xed, 0x60, 0x30, 0x1d, 0x21,
	0x15, 0x0f, 0xb5, 0xf1, 0xd0, 0xbf, 0x36, 0x9e, 0x6b, 0x34, 0xd8, 0xf7, 0x0a, 0x5b, 0xaf, 0xd9,
	0x9b, 0x7b, 0x0b, 0x78, 0x0a, 0xa1, 0xcf, 0x60, 0x97, 0xa6, 0xa1, 0xd1, 0x79, 0x95, 0xd4, 0x1a,
	0x84, 0x8a, 0x44, 0xd3, 0x50, 0x71, 0xaa, 0x75, 0xfd, 0x02, 0xf6, 0x59, 0xba, 0x88, 0x89, 0x64,
	0x3c, 0x35, 0x06, 0x63, 0xe0, 0xf4, 0x5d, 0xb7, 0xd2, 0xdf, 0x95, 0xd6, 0xf7, 0x88, 0xc9, 0xcb,
	0xa5, 0x8f, 0x02, 0x9e, 0xe0, 0xe6, 0x2b, 0xab, 0xc7, 0x4f, 0x11, 0x5e, 0x63, 0x59, 0x64, 0x54,
	0xa0, 0x19, 0x0d, 0x8e, 0xa5, 0xf5, 0x49, 0x91, 0x9f, 0x41, 0xb6, 0x77, 0x82, 0xba, 0xee, 0x76,
	0x6f, 0x82, 0xdd, 0xde, 0x04, 0x0f, 0x7b, 0x13, 0x6c, 0x0e, 0xa6, 0xb6, 0x3b, 0x98, 0xda, 0xed,
	0xc1, 0xd4, 0xfe, 0x3b, 0x67, 0x07, 0xe2, 0x20, 0xa5, 0x09, 0xbe, 0xca, 0x8a, 0x35, 0x5e, 0x9f,
	0xfd, 0x2f, 0xea, 0x8c, 0xdf, 0xab, 0x63, 0xfd, 0x7a, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xe4,
	0xa7, 0xf0, 0x55, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Periods) > 0 {
		for iNdEx := len(m.Periods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Periods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBotanydist(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Period) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Period) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Period) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBotanydist(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.End, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.End):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintBotanydist(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Start, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Start):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintBotanydist(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintBotanydist(dAtA []byte, offset int, v uint64) int {
	offset -= sovBotanydist(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Active {
		n += 2
	}
	if len(m.Periods) > 0 {
		for _, e := range m.Periods {
			l = e.Size()
			n += 1 + l + sovBotanydist(uint64(l))
		}
	}
	return n
}

func (m *Period) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Start)
	n += 1 + l + sovBotanydist(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.End)
	n += 1 + l + sovBotanydist(uint64(l))
	l = m.Inflation.Size()
	n += 1 + l + sovBotanydist(uint64(l))
	return n
}

func sovBotanydist(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBotanydist(x uint64) (n int) {
	return sovBotanydist(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBotanydist
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotanydist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Active = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Periods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotanydist
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
				return ErrInvalidLengthBotanydist
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBotanydist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Periods = append(m.Periods, Period{})
			if err := m.Periods[len(m.Periods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBotanydist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBotanydist
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
func (m *Period) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBotanydist
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
			return fmt.Errorf("proto: Period: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Period: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotanydist
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
				return ErrInvalidLengthBotanydist
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBotanydist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Start, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotanydist
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
				return ErrInvalidLengthBotanydist
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBotanydist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.End, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotanydist
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
				return ErrInvalidLengthBotanydist
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBotanydist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBotanydist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBotanydist
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
func skipBotanydist(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBotanydist
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
					return 0, ErrIntOverflowBotanydist
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
					return 0, ErrIntOverflowBotanydist
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
				return 0, ErrInvalidLengthBotanydist
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBotanydist
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBotanydist
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBotanydist        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBotanydist          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBotanydist = fmt.Errorf("proto: unexpected end of group")
)