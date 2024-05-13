// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/sanction/v1beta1/sanction.proto

package sanction

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

// TempStatus is whether a temporary entry is a sanction or unsanction.
type TempStatus int32

const (
	// TEMP_STATUS_UNSPECIFIED represents and unspecified status value.
	TEMP_STATUS_UNSPECIFIED TempStatus = 0
	// TEMP_STATUS_SANCTIONED indicates a sanction is in place.
	TEMP_STATUS_SANCTIONED TempStatus = 1
	// TEMP_STATUS_UNSANCTIONED indicates an unsanctioned is in place.
	TEMP_STATUS_UNSANCTIONED TempStatus = 2
)

var TempStatus_name = map[int32]string{
	0: "TEMP_STATUS_UNSPECIFIED",
	1: "TEMP_STATUS_SANCTIONED",
	2: "TEMP_STATUS_UNSANCTIONED",
}

var TempStatus_value = map[string]int32{
	"TEMP_STATUS_UNSPECIFIED":  0,
	"TEMP_STATUS_SANCTIONED":   1,
	"TEMP_STATUS_UNSANCTIONED": 2,
}

func (x TempStatus) String() string {
	return proto.EnumName(TempStatus_name, int32(x))
}

func (TempStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9e632afabc7910f0, []int{0}
}

// Params defines the configurable parameters of the sanction module.
type Params struct {
	// immediate_sanction_min_deposit is the minimum deposit for a sanction to happen immediately.
	// If this is zero, immediate sanctioning is not available.
	// Otherwise, if a sanction governance proposal is issued with a deposit at least this large, a temporary sanction
	// will be immediately issued that will expire when voting ends on the governance proposal.
	ImmediateSanctionMinDeposit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=immediate_sanction_min_deposit,json=immediateSanctionMinDeposit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"immediate_sanction_min_deposit"`
	// immediate_unsanction_min_deposit is the minimum deposit for an unsanction to happen immediately.
	// If this is zero, immediate unsanctioning is not available.
	// Otherwise, if an unsanction governance proposal is issued with a deposit at least this large, a temporary
	// unsanction will be immediately issued that will expire when voting ends on the governance proposal.
	ImmediateUnsanctionMinDeposit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=immediate_unsanction_min_deposit,json=immediateUnsanctionMinDeposit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"immediate_unsanction_min_deposit"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e632afabc7910f0, []int{0}
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

func (m *Params) GetImmediateSanctionMinDeposit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ImmediateSanctionMinDeposit
	}
	return nil
}

func (m *Params) GetImmediateUnsanctionMinDeposit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ImmediateUnsanctionMinDeposit
	}
	return nil
}

// TemporaryEntry defines the information involved in a temporary sanction or unsanction.
type TemporaryEntry struct {
	// address is the address of this temporary entry.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// proposal_id is the governance proposal id associated with this temporary entry.
	ProposalId uint64 `protobuf:"varint,2,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	// status is whether the entry is a sanction or unsanction.
	Status TempStatus `protobuf:"varint,3,opt,name=status,proto3,enum=cosmos.sanction.v1beta1.TempStatus" json:"status,omitempty"`
}

func (m *TemporaryEntry) Reset()         { *m = TemporaryEntry{} }
func (m *TemporaryEntry) String() string { return proto.CompactTextString(m) }
func (*TemporaryEntry) ProtoMessage()    {}
func (*TemporaryEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e632afabc7910f0, []int{1}
}
func (m *TemporaryEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TemporaryEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TemporaryEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TemporaryEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemporaryEntry.Merge(m, src)
}
func (m *TemporaryEntry) XXX_Size() int {
	return m.Size()
}
func (m *TemporaryEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_TemporaryEntry.DiscardUnknown(m)
}

var xxx_messageInfo_TemporaryEntry proto.InternalMessageInfo

func (m *TemporaryEntry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *TemporaryEntry) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

func (m *TemporaryEntry) GetStatus() TempStatus {
	if m != nil {
		return m.Status
	}
	return TEMP_STATUS_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("cosmos.sanction.v1beta1.TempStatus", TempStatus_name, TempStatus_value)
	proto.RegisterType((*Params)(nil), "cosmos.sanction.v1beta1.Params")
	proto.RegisterType((*TemporaryEntry)(nil), "cosmos.sanction.v1beta1.TemporaryEntry")
}

func init() {
	proto.RegisterFile("cosmos/sanction/v1beta1/sanction.proto", fileDescriptor_9e632afabc7910f0)
}

var fileDescriptor_9e632afabc7910f0 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x7d, 0x69, 0x15, 0xc4, 0x15, 0x55, 0xc5, 0xaa, 0xa8, 0x9b, 0x82, 0x63, 0x15, 0x09,
	0x59, 0x91, 0x62, 0xab, 0x61, 0x64, 0xca, 0x1f, 0x57, 0x64, 0x68, 0x88, 0x62, 0x67, 0x61, 0xb1,
	0x2e, 0xf6, 0xc9, 0x9c, 0xa8, 0xef, 0xac, 0xbb, 0x4b, 0x45, 0xbe, 0x01, 0x23, 0x33, 0x23, 0x48,
	0x08, 0x75, 0xea, 0xc0, 0x87, 0xe8, 0x58, 0x31, 0x31, 0x15, 0x94, 0x0c, 0xfd, 0x1a, 0x28, 0xf6,
	0xc5, 0x89, 0x10, 0xac, 0x2c, 0xb6, 0xef, 0x79, 0x9e, 0xf7, 0xd1, 0xcf, 0xf2, 0x6b, 0xf8, 0x2c,
	0x62, 0x22, 0x65, 0xc2, 0x15, 0x88, 0x46, 0x92, 0x30, 0xea, 0x5e, 0x9c, 0x4c, 0xb0, 0x44, 0x27,
	0xa5, 0xe0, 0x64, 0x9c, 0x49, 0xa6, 0x1f, 0x14, 0x39, 0xa7, 0x94, 0x55, 0xae, 0xf6, 0x10, 0xa5,
	0x84, 0x32, 0x37, 0xbf, 0x16, 0xd9, 0x9a, 0xa9, 0x3a, 0x27, 0x48, 0xe0, 0xb2, 0x2f, 0x62, 0x44,
	0x75, 0xd5, 0x0e, 0x0b, 0x3f, 0xcc, 0x4f, 0xae, 0x2a, 0x2e, 0xac, 0xfd, 0x84, 0x25, 0xac, 0xd0,
	0x97, 0x4f, 0x85, 0x7a, 0x7c, 0x5b, 0x81, 0xd5, 0x21, 0xe2, 0x28, 0x15, 0xfa, 0x17, 0x00, 0x4d,
	0x92, 0xa6, 0x38, 0x26, 0x48, 0xe2, 0x70, 0x45, 0x13, 0xa6, 0x84, 0x86, 0x31, 0xce, 0x98, 0x20,
	0xd2, 0x00, 0xd6, 0x96, 0xbd, 0xd3, 0x3a, 0x74, 0x54, 0xf1, 0x92, 0x62, 0x45, 0xeb, 0x74, 0x19,
	0xa1, 0x9d, 0xd3, 0xeb, 0xdb, 0xba, 0x76, 0xf9, 0xb3, 0x6e, 0x27, 0x44, 0xbe, 0x99, 0x4e, 0x9c,
	0x88, 0xa5, 0x8a, 0x42, 0xdd, 0x9a, 0x22, 0x7e, 0xeb, 0xca, 0x59, 0x86, 0x45, 0x3e, 0x20, 0x3e,
	0xde, 0x5d, 0x35, 0x1e, 0x9c, 0xe3, 0x04, 0x45, 0xb3, 0x70, 0xf9, 0x1e, 0xe2, 0xeb, 0xdd, 0x55,
	0x03, 0x8c, 0x8e, 0x4a, 0x10, 0x5f, 0x71, 0x9c, 0x11, 0xda, 0x2b, 0x28, 0xf4, 0x4b, 0x00, 0xad,
	0x35, 0xe8, 0x94, 0xfe, 0x15, 0xb5, 0xf2, 0xbf, 0x50, 0x9f, 0x94, 0x28, 0xe3, 0x92, 0x64, 0x0d,
	0x7b, 0xfc, 0x09, 0xc0, 0xdd, 0x00, 0xa7, 0x19, 0xe3, 0x88, 0xcf, 0x3c, 0x2a, 0xf9, 0x4c, 0x6f,
	0xc1, 0x7b, 0x28, 0x8e, 0x39, 0x16, 0xc2, 0x00, 0x16, 0xb0, 0xef, 0x77, 0x8c, 0xef, 0xdf, 0x9a,
	0xfb, 0x0a, 0xb4, 0x5d, 0x38, 0xbe, 0xe4, 0x84, 0x26, 0xa3, 0x55, 0x50, 0xaf, 0xc3, 0x9d, 0x8c,
	0xb3, 0x8c, 0x09, 0x74, 0x1e, 0x92, 0xd8, 0xa8, 0x58, 0xc0, 0xde, 0x1e, 0xc1, 0x95, 0xd4, 0x8f,
	0xf5, 0x17, 0xb0, 0x2a, 0x24, 0x92, 0x53, 0x61, 0x6c, 0x59, 0xc0, 0xde, 0x6d, 0x3d, 0x75, 0xfe,
	0xb1, 0x56, 0xce, 0x92, 0xc6, 0xcf, 0xa3, 0x23, 0x35, 0xd2, 0x20, 0x10, 0xae, 0x55, 0xfd, 0x08,
	0x1e, 0x04, 0xde, 0xd9, 0x30, 0xf4, 0x83, 0x76, 0x30, 0xf6, 0xc3, 0xf1, 0xc0, 0x1f, 0x7a, 0xdd,
	0xfe, 0x69, 0xdf, 0xeb, 0xed, 0x69, 0x7a, 0x0d, 0x3e, 0xda, 0x34, 0xfd, 0xf6, 0xa0, 0x1b, 0xf4,
	0x5f, 0x0d, 0xbc, 0xde, 0x1e, 0xd0, 0x1f, 0x43, 0xe3, 0x8f, 0xc1, 0xb5, 0x5b, 0xa9, 0x6d, 0xbf,
	0xff, 0x6c, 0x6a, 0x9d, 0x97, 0xd7, 0x73, 0x13, 0xdc, 0xcc, 0x4d, 0xf0, 0x6b, 0x6e, 0x82, 0x0f,
	0x0b, 0x53, 0xbb, 0x59, 0x98, 0xda, 0x8f, 0x85, 0xa9, 0xbd, 0x76, 0x36, 0x3e, 0x44, 0xc6, 0xd9,
	0x05, 0xa6, 0x88, 0x46, 0xb8, 0x49, 0xd8, 0xc6, 0xc9, 0x7d, 0x57, 0xfe, 0x3d, 0x93, 0x6a, 0xbe,
	0xc1, 0xcf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x31, 0x26, 0x2e, 0x68, 0x03, 0x00, 0x00,
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
	if len(m.ImmediateUnsanctionMinDeposit) > 0 {
		for iNdEx := len(m.ImmediateUnsanctionMinDeposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ImmediateUnsanctionMinDeposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSanction(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ImmediateSanctionMinDeposit) > 0 {
		for iNdEx := len(m.ImmediateSanctionMinDeposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ImmediateSanctionMinDeposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSanction(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TemporaryEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TemporaryEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TemporaryEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintSanction(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if m.ProposalId != 0 {
		i = encodeVarintSanction(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSanction(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSanction(dAtA []byte, offset int, v uint64) int {
	offset -= sovSanction(v)
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
	if len(m.ImmediateSanctionMinDeposit) > 0 {
		for _, e := range m.ImmediateSanctionMinDeposit {
			l = e.Size()
			n += 1 + l + sovSanction(uint64(l))
		}
	}
	if len(m.ImmediateUnsanctionMinDeposit) > 0 {
		for _, e := range m.ImmediateUnsanctionMinDeposit {
			l = e.Size()
			n += 1 + l + sovSanction(uint64(l))
		}
	}
	return n
}

func (m *TemporaryEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSanction(uint64(l))
	}
	if m.ProposalId != 0 {
		n += 1 + sovSanction(uint64(m.ProposalId))
	}
	if m.Status != 0 {
		n += 1 + sovSanction(uint64(m.Status))
	}
	return n
}

func sovSanction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSanction(x uint64) (n int) {
	return sovSanction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSanction
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmediateSanctionMinDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSanction
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
				return ErrInvalidLengthSanction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSanction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImmediateSanctionMinDeposit = append(m.ImmediateSanctionMinDeposit, types.Coin{})
			if err := m.ImmediateSanctionMinDeposit[len(m.ImmediateSanctionMinDeposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmediateUnsanctionMinDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSanction
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
				return ErrInvalidLengthSanction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSanction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImmediateUnsanctionMinDeposit = append(m.ImmediateUnsanctionMinDeposit, types.Coin{})
			if err := m.ImmediateUnsanctionMinDeposit[len(m.ImmediateUnsanctionMinDeposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSanction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSanction
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
func (m *TemporaryEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSanction
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
			return fmt.Errorf("proto: TemporaryEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TemporaryEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSanction
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
				return ErrInvalidLengthSanction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSanction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSanction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSanction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= TempStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSanction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSanction
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
func skipSanction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSanction
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
					return 0, ErrIntOverflowSanction
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
					return 0, ErrIntOverflowSanction
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
				return 0, ErrInvalidLengthSanction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSanction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSanction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSanction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSanction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSanction = fmt.Errorf("proto: unexpected end of group")
)
