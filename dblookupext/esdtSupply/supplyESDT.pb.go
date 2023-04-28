// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: supplyESDT.proto

package esdtSupply

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_multiversx_mx_chain_core_go_data "github.com/multiversx/mx-chain-core-go/data"
	io "io"
	math "math"
	math_big "math/big"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// SupplyESDT is used to store information a shard esdt token supply
type SupplyESDT struct {
	Supply           *math_big.Int `protobuf:"bytes,1,opt,name=Supply,proto3,casttypewith=math/big.Int;github.com/multiversx/mx-chain-core-go/data.BigIntCaster" json:"value"`
	Burned           *math_big.Int `protobuf:"bytes,2,opt,name=Burned,proto3,casttypewith=math/big.Int;github.com/multiversx/mx-chain-core-go/data.BigIntCaster" json:"burned"`
	Minted           *math_big.Int `protobuf:"bytes,3,opt,name=Minted,proto3,casttypewith=math/big.Int;github.com/multiversx/mx-chain-core-go/data.BigIntCaster" json:"minted"`
	RecomputedSupply bool          `protobuf:"varint,4,opt,name=RecomputedSupply,proto3" json:"RecomputedSupply"`
}

func (m *SupplyESDT) Reset()      { *m = SupplyESDT{} }
func (*SupplyESDT) ProtoMessage() {}
func (*SupplyESDT) Descriptor() ([]byte, []int) {
	return fileDescriptor_173c6d56cc05b222, []int{0}
}
func (m *SupplyESDT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SupplyESDT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SupplyESDT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupplyESDT.Merge(m, src)
}
func (m *SupplyESDT) XXX_Size() int {
	return m.Size()
}
func (m *SupplyESDT) XXX_DiscardUnknown() {
	xxx_messageInfo_SupplyESDT.DiscardUnknown(m)
}

var xxx_messageInfo_SupplyESDT proto.InternalMessageInfo

func (m *SupplyESDT) GetSupply() *math_big.Int {
	if m != nil {
		return m.Supply
	}
	return nil
}

func (m *SupplyESDT) GetBurned() *math_big.Int {
	if m != nil {
		return m.Burned
	}
	return nil
}

func (m *SupplyESDT) GetMinted() *math_big.Int {
	if m != nil {
		return m.Minted
	}
	return nil
}

func (m *SupplyESDT) GetRecomputedSupply() bool {
	if m != nil {
		return m.RecomputedSupply
	}
	return false
}

func init() {
	proto.RegisterType((*SupplyESDT)(nil), "proto.SupplyESDT")
}

func init() { proto.RegisterFile("supplyESDT.proto", fileDescriptor_173c6d56cc05b222) }

var fileDescriptor_173c6d56cc05b222 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0xd1, 0x3f, 0x4e, 0xc3, 0x30,
	0x14, 0x06, 0xf0, 0x98, 0xd2, 0x08, 0x59, 0x0c, 0x55, 0xc4, 0x50, 0x31, 0xbc, 0x54, 0x4c, 0x5d,
	0x92, 0x0c, 0x8c, 0x2c, 0x28, 0xb4, 0x43, 0x07, 0x96, 0x94, 0x89, 0x2d, 0x7f, 0x8c, 0x6b, 0xa8,
	0xe3, 0x28, 0xb1, 0xab, 0xb2, 0x71, 0x04, 0x06, 0x0e, 0x81, 0x38, 0x09, 0x63, 0xc7, 0x4e, 0x85,
	0xba, 0x0b, 0xea, 0xd4, 0x23, 0x20, 0x9c, 0x0a, 0x90, 0xba, 0x76, 0xb2, 0xbf, 0xcf, 0xf2, 0xfb,
	0x49, 0x36, 0x6e, 0x55, 0xaa, 0x28, 0xc6, 0x8f, 0xfd, 0x61, 0xef, 0xc6, 0x2f, 0x4a, 0x21, 0x85,
	0xd3, 0x34, 0xcb, 0xa9, 0x47, 0x99, 0x1c, 0xa9, 0xc4, 0x4f, 0x05, 0x0f, 0xa8, 0xa0, 0x22, 0x30,
	0x75, 0xa2, 0xee, 0x4c, 0x32, 0xc1, 0xec, 0xea, 0x5b, 0x67, 0x2f, 0x0d, 0x8c, 0x87, 0xbf, 0xa3,
	0x9c, 0x7b, 0x6c, 0xd7, 0xa9, 0x8d, 0x3a, 0xa8, 0x7b, 0x1c, 0x46, 0xeb, 0x85, 0xdb, 0x9c, 0xc4,
	0x63, 0x45, 0xde, 0x3e, 0xdc, 0x3e, 0x8f, 0xe5, 0x28, 0x48, 0x18, 0xf5, 0x07, 0xb9, 0xbc, 0xf8,
	0xe7, 0x70, 0x35, 0x96, 0x6c, 0x42, 0xca, 0x6a, 0x1a, 0xf0, 0xa9, 0x97, 0x8e, 0x62, 0x96, 0x7b,
	0xa9, 0x28, 0x89, 0x47, 0x45, 0x90, 0xc5, 0x32, 0xf6, 0x43, 0x46, 0x07, 0xb9, 0xbc, 0x8a, 0x2b,
	0x49, 0xca, 0x68, 0x2b, 0x38, 0x0f, 0xd8, 0x0e, 0x55, 0x99, 0x93, 0xac, 0x7d, 0x60, 0xac, 0xe1,
	0x7a, 0xe1, 0xda, 0x89, 0x69, 0xf6, 0x88, 0xd5, 0xc4, 0x0f, 0x76, 0xcd, 0x72, 0x49, 0xb2, 0x76,
	0xe3, 0x0f, 0xe3, 0xa6, 0xd9, 0x23, 0x56, 0x13, 0xce, 0x25, 0x6e, 0x45, 0x24, 0x15, 0xbc, 0x50,
	0x92, 0x64, 0xdb, 0xf7, 0x3c, 0xec, 0xa0, 0xee, 0x51, 0x78, 0xb2, 0x5e, 0xb8, 0x3b, 0x67, 0xd1,
	0x4e, 0x13, 0xf6, 0x66, 0x4b, 0xb0, 0xe6, 0x4b, 0xb0, 0x36, 0x4b, 0x40, 0x4f, 0x1a, 0xd0, 0xab,
	0x06, 0xf4, 0xae, 0x01, 0xcd, 0x34, 0xa0, 0xb9, 0x06, 0xf4, 0xa9, 0x01, 0x7d, 0x69, 0xb0, 0x36,
	0x1a, 0xd0, 0xf3, 0x0a, 0xac, 0xd9, 0x0a, 0xac, 0xf9, 0x0a, 0xac, 0x5b, 0x4c, 0xaa, 0x4c, 0xd6,
	0x53, 0x12, 0xdb, 0xfc, 0xf1, 0xf9, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x06, 0x34, 0x97,
	0x2d, 0x02, 0x00, 0x00,
}

func (this *SupplyESDT) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SupplyESDT)
	if !ok {
		that2, ok := that.(SupplyESDT)
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
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		if !__caster.Equal(this.Supply, that1.Supply) {
			return false
		}
	}
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		if !__caster.Equal(this.Burned, that1.Burned) {
			return false
		}
	}
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		if !__caster.Equal(this.Minted, that1.Minted) {
			return false
		}
	}
	if this.RecomputedSupply != that1.RecomputedSupply {
		return false
	}
	return true
}
func (this *SupplyESDT) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&esdtSupply.SupplyESDT{")
	s = append(s, "Supply: "+fmt.Sprintf("%#v", this.Supply)+",\n")
	s = append(s, "Burned: "+fmt.Sprintf("%#v", this.Burned)+",\n")
	s = append(s, "Minted: "+fmt.Sprintf("%#v", this.Minted)+",\n")
	s = append(s, "RecomputedSupply: "+fmt.Sprintf("%#v", this.RecomputedSupply)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSupplyESDT(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *SupplyESDT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SupplyESDT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SupplyESDT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RecomputedSupply {
		i--
		if m.RecomputedSupply {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		size := __caster.Size(m.Minted)
		i -= size
		if _, err := __caster.MarshalTo(m.Minted, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSupplyESDT(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		size := __caster.Size(m.Burned)
		i -= size
		if _, err := __caster.MarshalTo(m.Burned, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSupplyESDT(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		size := __caster.Size(m.Supply)
		i -= size
		if _, err := __caster.MarshalTo(m.Supply, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSupplyESDT(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintSupplyESDT(dAtA []byte, offset int, v uint64) int {
	offset -= sovSupplyESDT(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SupplyESDT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		l = __caster.Size(m.Supply)
		n += 1 + l + sovSupplyESDT(uint64(l))
	}
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		l = __caster.Size(m.Burned)
		n += 1 + l + sovSupplyESDT(uint64(l))
	}
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		l = __caster.Size(m.Minted)
		n += 1 + l + sovSupplyESDT(uint64(l))
	}
	if m.RecomputedSupply {
		n += 2
	}
	return n
}

func sovSupplyESDT(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSupplyESDT(x uint64) (n int) {
	return sovSupplyESDT(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SupplyESDT) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SupplyESDT{`,
		`Supply:` + fmt.Sprintf("%v", this.Supply) + `,`,
		`Burned:` + fmt.Sprintf("%v", this.Burned) + `,`,
		`Minted:` + fmt.Sprintf("%v", this.Minted) + `,`,
		`RecomputedSupply:` + fmt.Sprintf("%v", this.RecomputedSupply) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSupplyESDT(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SupplyESDT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSupplyESDT
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
			return fmt.Errorf("proto: SupplyESDT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SupplyESDT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplyESDT
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSupplyESDT
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSupplyESDT
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Supply = tmp
				}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Burned", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplyESDT
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSupplyESDT
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSupplyESDT
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Burned = tmp
				}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minted", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplyESDT
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSupplyESDT
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSupplyESDT
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Minted = tmp
				}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecomputedSupply", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplyESDT
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
			m.RecomputedSupply = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSupplyESDT(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSupplyESDT
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSupplyESDT
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
func skipSupplyESDT(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSupplyESDT
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
					return 0, ErrIntOverflowSupplyESDT
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
					return 0, ErrIntOverflowSupplyESDT
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
				return 0, ErrInvalidLengthSupplyESDT
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSupplyESDT
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSupplyESDT
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSupplyESDT        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSupplyESDT          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSupplyESDT = fmt.Errorf("proto: unexpected end of group")
)
