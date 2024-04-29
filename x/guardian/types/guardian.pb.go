// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: guardian/v1/guardian.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type STATUS int32

const (
	STATUS_CLAIMABLE STATUS = 0
	STATUS_CLAIMED   STATUS = 1
)

var STATUS_name = map[int32]string{
	0: "CLAIMABLE",
	1: "CLAIMED",
}

var STATUS_value = map[string]int32{
	"CLAIMABLE": 0,
	"CLAIMED":   1,
}

func (x STATUS) String() string {
	return proto.EnumName(STATUS_name, int32(x))
}

func (STATUS) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7f8d20ee917ce1bf, []int{0}
}

type Payee struct {
	Payee  string `protobuf:"bytes,1,opt,name=payee,proto3" json:"payee,omitempty"`
	Status STATUS `protobuf:"varint,2,opt,name=status,proto3,enum=guardian.v1.STATUS" json:"status,omitempty"`
}

func (m *Payee) Reset()         { *m = Payee{} }
func (m *Payee) String() string { return proto.CompactTextString(m) }
func (*Payee) ProtoMessage()    {}
func (*Payee) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f8d20ee917ce1bf, []int{0}
}
func (m *Payee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Payee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Payee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Payee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payee.Merge(m, src)
}
func (m *Payee) XXX_Size() int {
	return m.Size()
}
func (m *Payee) XXX_DiscardUnknown() {
	xxx_messageInfo_Payee.DiscardUnknown(m)
}

var xxx_messageInfo_Payee proto.InternalMessageInfo

func (m *Payee) GetPayee() string {
	if m != nil {
		return m.Payee
	}
	return ""
}

func (m *Payee) GetStatus() STATUS {
	if m != nil {
		return m.Status
	}
	return STATUS_CLAIMABLE
}

type GuardedFee struct {
	Payer  string                                   `protobuf:"bytes,1,opt,name=payer,proto3" json:"payer,omitempty"`
	Payees []*Payee                                 `protobuf:"bytes,2,rep,name=payees,proto3" json:"payees,omitempty"`
	Fee    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=fee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee"`
}

func (m *GuardedFee) Reset()         { *m = GuardedFee{} }
func (m *GuardedFee) String() string { return proto.CompactTextString(m) }
func (*GuardedFee) ProtoMessage()    {}
func (*GuardedFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f8d20ee917ce1bf, []int{1}
}
func (m *GuardedFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GuardedFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GuardedFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GuardedFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuardedFee.Merge(m, src)
}
func (m *GuardedFee) XXX_Size() int {
	return m.Size()
}
func (m *GuardedFee) XXX_DiscardUnknown() {
	xxx_messageInfo_GuardedFee.DiscardUnknown(m)
}

var xxx_messageInfo_GuardedFee proto.InternalMessageInfo

func (m *GuardedFee) GetPayer() string {
	if m != nil {
		return m.Payer
	}
	return ""
}

func (m *GuardedFee) GetPayees() []*Payee {
	if m != nil {
		return m.Payees
	}
	return nil
}

func (m *GuardedFee) GetFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Fee
	}
	return nil
}

func init() {
	proto.RegisterEnum("guardian.v1.STATUS", STATUS_name, STATUS_value)
	proto.RegisterType((*Payee)(nil), "guardian.v1.Payee")
	proto.RegisterType((*GuardedFee)(nil), "guardian.v1.GuardedFee")
}

func init() { proto.RegisterFile("guardian/v1/guardian.proto", fileDescriptor_7f8d20ee917ce1bf) }

var fileDescriptor_7f8d20ee917ce1bf = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x2f, 0x4d, 0x2c,
	0x4a, 0xc9, 0x4c, 0xcc, 0xd3, 0x2f, 0x33, 0xd4, 0x87, 0xb1, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0xb8, 0xe1, 0xfc, 0x32, 0x43, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0xb8, 0x3e, 0x88,
	0x05, 0x51, 0x22, 0x25, 0x97, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0xac, 0x9f, 0x94, 0x58, 0x9c, 0xaa,
	0x5f, 0x66, 0x98, 0x94, 0x5a, 0x92, 0x68, 0xa8, 0x9f, 0x9c, 0x9f, 0x09, 0x35, 0x42, 0x29, 0x80,
	0x8b, 0x35, 0x20, 0xb1, 0x32, 0x35, 0x55, 0x48, 0x84, 0x8b, 0xb5, 0x00, 0xc4, 0x90, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0xb4, 0xb9, 0xd8, 0x8a, 0x4b, 0x12, 0x4b, 0x4a, 0x8b,
	0x25, 0x98, 0x14, 0x18, 0x35, 0xf8, 0x8c, 0x84, 0xf5, 0x90, 0xac, 0xd4, 0x0b, 0x0e, 0x71, 0x0c,
	0x09, 0x0d, 0x0e, 0x82, 0x2a, 0xb1, 0x62, 0x79, 0xb1, 0x40, 0x9e, 0x51, 0x69, 0x33, 0x23, 0x17,
	0x97, 0x3b, 0x48, 0x51, 0x6a, 0x8a, 0x1b, 0xc2, 0xdc, 0x22, 0x64, 0x73, 0x8b, 0x84, 0xb4, 0xb8,
	0xd8, 0xc0, 0x16, 0x80, 0xcc, 0x65, 0xd6, 0xe0, 0x36, 0x12, 0x42, 0x31, 0x17, 0xec, 0xa2, 0x20,
	0xa8, 0x0a, 0xa1, 0x58, 0x2e, 0xe6, 0xb4, 0xd4, 0x54, 0x09, 0x66, 0xb0, 0x42, 0x49, 0x3d, 0x88,
	0x87, 0xf4, 0x40, 0x1e, 0xd2, 0x83, 0x7a, 0x48, 0xcf, 0x39, 0x3f, 0x33, 0xcf, 0xc9, 0xe0, 0xc4,
	0x3d, 0x79, 0x86, 0x55, 0xf7, 0xe5, 0x35, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3,
	0x73, 0xf5, 0xa1, 0xbe, 0x87, 0x50, 0xba, 0xc5, 0x29, 0xd9, 0xfa, 0x25, 0x95, 0x05, 0xa9, 0xc5,
	0x60, 0x0d, 0xc5, 0x41, 0x20, 0x73, 0x21, 0xae, 0xd6, 0x52, 0xe1, 0x62, 0x83, 0xf8, 0x46, 0x88,
	0x97, 0x8b, 0xd3, 0xd9, 0xc7, 0xd1, 0xd3, 0xd7, 0xd1, 0xc9, 0xc7, 0x55, 0x80, 0x41, 0x88, 0x9b,
	0x8b, 0x1d, 0xcc, 0x75, 0x75, 0x11, 0x60, 0x74, 0xf2, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23,
	0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6,
	0x63, 0x39, 0x86, 0x28, 0x3d, 0x24, 0x4b, 0x93, 0x12, 0xf3, 0x52, 0xc0, 0xa1, 0x9b, 0x9c, 0x9f,
	0xa3, 0x9f, 0x9c, 0x91, 0x98, 0x99, 0xa7, 0x5f, 0x01, 0x8f, 0x3a, 0x88, 0x03, 0x92, 0xd8, 0xc0,
	0x0a, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x78, 0xeb, 0x62, 0xdf, 0x01, 0x00, 0x00,
}

func (this *Payee) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Payee)
	if !ok {
		that2, ok := that.(Payee)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Payee != that1.Payee {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	return true
}
func (this *GuardedFee) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GuardedFee)
	if !ok {
		that2, ok := that.(GuardedFee)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Payer != that1.Payer {
		return false
	}
	if len(this.Payees) != len(that1.Payees) {
		return false
	}
	for i := range this.Payees {
		if !this.Payees[i].Equal(that1.Payees[i]) {
			return false
		}
	}
	if len(this.Fee) != len(that1.Fee) {
		return false
	}
	for i := range this.Fee {
		if !this.Fee[i].Equal(&that1.Fee[i]) {
			return false
		}
	}
	return true
}
func (m *Payee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Payee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Payee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintGuardian(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Payee) > 0 {
		i -= len(m.Payee)
		copy(dAtA[i:], m.Payee)
		i = encodeVarintGuardian(dAtA, i, uint64(len(m.Payee)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GuardedFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GuardedFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GuardedFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fee) > 0 {
		for iNdEx := len(m.Fee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGuardian(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Payees) > 0 {
		for iNdEx := len(m.Payees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Payees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGuardian(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Payer) > 0 {
		i -= len(m.Payer)
		copy(dAtA[i:], m.Payer)
		i = encodeVarintGuardian(dAtA, i, uint64(len(m.Payer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGuardian(dAtA []byte, offset int, v uint64) int {
	offset -= sovGuardian(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Payee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Payee)
	if l > 0 {
		n += 1 + l + sovGuardian(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovGuardian(uint64(m.Status))
	}
	return n
}

func (m *GuardedFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Payer)
	if l > 0 {
		n += 1 + l + sovGuardian(uint64(l))
	}
	if len(m.Payees) > 0 {
		for _, e := range m.Payees {
			l = e.Size()
			n += 1 + l + sovGuardian(uint64(l))
		}
	}
	if len(m.Fee) > 0 {
		for _, e := range m.Fee {
			l = e.Size()
			n += 1 + l + sovGuardian(uint64(l))
		}
	}
	return n
}

func sovGuardian(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGuardian(x uint64) (n int) {
	return sovGuardian(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Payee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuardian
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
			return fmt.Errorf("proto: Payee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Payee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardian
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
				return ErrInvalidLengthGuardian
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGuardian
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardian
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= STATUS(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGuardian(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGuardian
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
func (m *GuardedFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuardian
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
			return fmt.Errorf("proto: GuardedFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GuardedFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardian
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
				return ErrInvalidLengthGuardian
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGuardian
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardian
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
				return ErrInvalidLengthGuardian
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGuardian
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payees = append(m.Payees, &Payee{})
			if err := m.Payees[len(m.Payees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardian
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
				return ErrInvalidLengthGuardian
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGuardian
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = append(m.Fee, types.Coin{})
			if err := m.Fee[len(m.Fee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGuardian(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGuardian
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
func skipGuardian(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGuardian
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
					return 0, ErrIntOverflowGuardian
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
					return 0, ErrIntOverflowGuardian
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
				return 0, ErrInvalidLengthGuardian
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGuardian
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGuardian
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGuardian        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGuardian          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGuardian = fmt.Errorf("proto: unexpected end of group")
)
