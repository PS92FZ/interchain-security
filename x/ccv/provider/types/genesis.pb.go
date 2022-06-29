// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchain_security/ccv/provider/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/interchain-security/x/ccv/types"
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

// GenesisState defines the CCV provider chain genesis state
type GenesisState struct {
	ConsumerStates []ConsumerState `protobuf:"bytes,1,rep,name=consumer_states,json=consumerStates,proto3" json:"consumer_states" yaml:"consumer_states"`
	Params         Params          `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_48411d9c7900d48e, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetConsumerStates() []ConsumerState {
	if m != nil {
		return m.ConsumerStates
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// ConsumerState defines the state that the provider chain stores for each consumer chain
type ConsumerState struct {
	ChainId   string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (m *ConsumerState) Reset()         { *m = ConsumerState{} }
func (m *ConsumerState) String() string { return proto.CompactTextString(m) }
func (*ConsumerState) ProtoMessage()    {}
func (*ConsumerState) Descriptor() ([]byte, []int) {
	return fileDescriptor_48411d9c7900d48e, []int{1}
}
func (m *ConsumerState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConsumerState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConsumerState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConsumerState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumerState.Merge(m, src)
}
func (m *ConsumerState) XXX_Size() int {
	return m.Size()
}
func (m *ConsumerState) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumerState.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumerState proto.InternalMessageInfo

func (m *ConsumerState) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *ConsumerState) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "interchain_security.ccv.provider.v1.GenesisState")
	proto.RegisterType((*ConsumerState)(nil), "interchain_security.ccv.provider.v1.ConsumerState")
}

func init() {
	proto.RegisterFile("interchain_security/ccv/provider/v1/genesis.proto", fileDescriptor_48411d9c7900d48e)
}

var fileDescriptor_48411d9c7900d48e = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x33, 0xfd, 0x3e, 0xaa, 0x9d, 0xfa, 0x07, 0x82, 0x48, 0x2d, 0x38, 0x2d, 0xd1, 0x45,
	0x41, 0x9c, 0x21, 0x71, 0xd7, 0x65, 0x5d, 0x48, 0x76, 0x52, 0x5d, 0xb9, 0x29, 0xe9, 0x64, 0x48,
	0x07, 0x9a, 0x4c, 0x98, 0x99, 0x06, 0x8b, 0x2f, 0xe1, 0x63, 0x75, 0xd9, 0x95, 0xb8, 0x2a, 0xd2,
	0xbe, 0x81, 0x4f, 0x20, 0x99, 0xc6, 0x6a, 0x45, 0x21, 0xbb, 0x99, 0x7b, 0xef, 0xef, 0x9c, 0x03,
	0x07, 0xba, 0x3c, 0xd1, 0x4c, 0xd2, 0x51, 0xc0, 0x93, 0x81, 0x62, 0x74, 0x22, 0xb9, 0x9e, 0x12,
	0x4a, 0x33, 0x92, 0x4a, 0x91, 0xf1, 0x90, 0x49, 0x92, 0xb9, 0x24, 0x62, 0x09, 0x53, 0x5c, 0xe1,
	0x54, 0x0a, 0x2d, 0xec, 0xb3, 0x5f, 0x10, 0x4c, 0x69, 0x86, 0x3f, 0x11, 0x9c, 0xb9, 0xcd, 0xa3,
	0x48, 0x44, 0xc2, 0xdc, 0x93, 0xfc, 0xb5, 0x46, 0x9b, 0xe7, 0x7f, 0xb9, 0x65, 0x2e, 0x29, 0x14,
	0xb4, 0x68, 0x7a, 0x65, 0x32, 0x6d, 0xcc, 0x0c, 0xe3, 0xbc, 0x00, 0xb8, 0x77, 0xb3, 0x8e, 0x79,
	0xa7, 0x03, 0xcd, 0xec, 0x27, 0x78, 0x48, 0x45, 0xa2, 0x26, 0x31, 0x93, 0x03, 0x95, 0x4f, 0x54,
	0x03, 0xb4, 0xff, 0x75, 0xea, 0x9e, 0x87, 0x4b, 0xe4, 0xc7, 0xd7, 0x05, 0x6b, 0xc4, 0x7a, 0x68,
	0xb6, 0x68, 0x59, 0xef, 0x8b, 0xd6, 0xf1, 0x34, 0x88, 0xc7, 0x5d, 0xe7, 0x87, 0xb0, 0xd3, 0x3f,
	0xa0, 0xdf, 0xcf, 0x95, 0xed, 0xc3, 0x6a, 0x1a, 0xc8, 0x20, 0x56, 0x8d, 0x4a, 0x1b, 0x74, 0xea,
	0xde, 0x45, 0x29, 0xcf, 0x5b, 0x83, 0xf4, 0xfe, 0xe7, 0x66, 0xfd, 0x42, 0xc0, 0xf1, 0xe1, 0xfe,
	0x56, 0x16, 0xfb, 0x04, 0xee, 0xae, 0x75, 0x78, 0xd8, 0x00, 0x6d, 0xd0, 0xa9, 0xf5, 0x77, 0xcc,
	0xdf, 0x0f, 0xed, 0x53, 0x08, 0xe9, 0x28, 0x48, 0x12, 0x36, 0xce, 0x97, 0x15, 0xb3, 0xac, 0x15,
	0x13, 0x3f, 0xec, 0xdd, 0xcf, 0x96, 0x08, 0xcc, 0x97, 0x08, 0xbc, 0x2d, 0x11, 0x78, 0x5e, 0x21,
	0x6b, 0xbe, 0x42, 0xd6, 0xeb, 0x0a, 0x59, 0x0f, 0xdd, 0x88, 0xeb, 0xd1, 0x64, 0x88, 0xa9, 0x88,
	0x09, 0x15, 0x2a, 0x16, 0x8a, 0x7c, 0x05, 0xbe, 0xdc, 0x74, 0xf0, 0xb8, 0xdd, 0x82, 0x9e, 0xa6,
	0x4c, 0x0d, 0xab, 0xa6, 0x80, 0xab, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0x40, 0x06, 0x36,
	0x4a, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ConsumerStates) > 0 {
		for iNdEx := len(m.ConsumerStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ConsumerStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ConsumerState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsumerState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConsumerState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ConsumerStates) > 0 {
		for _, e := range m.ConsumerStates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *ConsumerState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsumerStates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsumerStates = append(m.ConsumerStates, ConsumerState{})
			if err := m.ConsumerStates[len(m.ConsumerStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ConsumerState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ConsumerState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsumerState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)