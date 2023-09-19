// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/community/v1beta1/proposal.proto

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

// CommunityPoolLendDepositProposal deposits from the community pool into lend
type CommunityPoolLendDepositProposal struct {
	Title       string                                   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                                   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Amount      github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *CommunityPoolLendDepositProposal) Reset()      { *m = CommunityPoolLendDepositProposal{} }
func (*CommunityPoolLendDepositProposal) ProtoMessage() {}
func (*CommunityPoolLendDepositProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_64aa83b2ed448ec1, []int{0}
}
func (m *CommunityPoolLendDepositProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommunityPoolLendDepositProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommunityPoolLendDepositProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommunityPoolLendDepositProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunityPoolLendDepositProposal.Merge(m, src)
}
func (m *CommunityPoolLendDepositProposal) XXX_Size() int {
	return m.Size()
}
func (m *CommunityPoolLendDepositProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunityPoolLendDepositProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CommunityPoolLendDepositProposal proto.InternalMessageInfo

// CommunityPoolLendWithdrawProposal withdraws a lend position back to the community pool
type CommunityPoolLendWithdrawProposal struct {
	Title       string                                   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                                   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Amount      github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *CommunityPoolLendWithdrawProposal) Reset()      { *m = CommunityPoolLendWithdrawProposal{} }
func (*CommunityPoolLendWithdrawProposal) ProtoMessage() {}
func (*CommunityPoolLendWithdrawProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_64aa83b2ed448ec1, []int{1}
}
func (m *CommunityPoolLendWithdrawProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommunityPoolLendWithdrawProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommunityPoolLendWithdrawProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommunityPoolLendWithdrawProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunityPoolLendWithdrawProposal.Merge(m, src)
}
func (m *CommunityPoolLendWithdrawProposal) XXX_Size() int {
	return m.Size()
}
func (m *CommunityPoolLendWithdrawProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunityPoolLendWithdrawProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CommunityPoolLendWithdrawProposal proto.InternalMessageInfo

// CommunityCDPRepayDebtProposal repays a cdp debt position owned by the community module
// This proposal exists primarily to allow committees to repay community module cdp debts.
type CommunityCDPRepayDebtProposal struct {
	Title          string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description    string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CollateralType string     `protobuf:"bytes,3,opt,name=collateral_type,json=collateralType,proto3" json:"collateral_type,omitempty"`
	Payment        types.Coin `protobuf:"bytes,4,opt,name=payment,proto3" json:"payment"`
}

func (m *CommunityCDPRepayDebtProposal) Reset()      { *m = CommunityCDPRepayDebtProposal{} }
func (*CommunityCDPRepayDebtProposal) ProtoMessage() {}
func (*CommunityCDPRepayDebtProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_64aa83b2ed448ec1, []int{2}
}
func (m *CommunityCDPRepayDebtProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommunityCDPRepayDebtProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommunityCDPRepayDebtProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommunityCDPRepayDebtProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunityCDPRepayDebtProposal.Merge(m, src)
}
func (m *CommunityCDPRepayDebtProposal) XXX_Size() int {
	return m.Size()
}
func (m *CommunityCDPRepayDebtProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunityCDPRepayDebtProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CommunityCDPRepayDebtProposal proto.InternalMessageInfo

// CommunityCDPWithdrawCollateralProposal withdraws cdp collateral owned by the community module
// This proposal exists primarily to allow committees to withdraw community module cdp collateral.
type CommunityCDPWithdrawCollateralProposal struct {
	Title          string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description    string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CollateralType string     `protobuf:"bytes,3,opt,name=collateral_type,json=collateralType,proto3" json:"collateral_type,omitempty"`
	Collateral     types.Coin `protobuf:"bytes,4,opt,name=collateral,proto3" json:"collateral"`
}

func (m *CommunityCDPWithdrawCollateralProposal) Reset() {
	*m = CommunityCDPWithdrawCollateralProposal{}
}
func (*CommunityCDPWithdrawCollateralProposal) ProtoMessage() {}
func (*CommunityCDPWithdrawCollateralProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_64aa83b2ed448ec1, []int{3}
}
func (m *CommunityCDPWithdrawCollateralProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommunityCDPWithdrawCollateralProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommunityCDPWithdrawCollateralProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommunityCDPWithdrawCollateralProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunityCDPWithdrawCollateralProposal.Merge(m, src)
}
func (m *CommunityCDPWithdrawCollateralProposal) XXX_Size() int {
	return m.Size()
}
func (m *CommunityCDPWithdrawCollateralProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunityCDPWithdrawCollateralProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CommunityCDPWithdrawCollateralProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CommunityPoolLendDepositProposal)(nil), "kava.community.v1beta1.CommunityPoolLendDepositProposal")
	proto.RegisterType((*CommunityPoolLendWithdrawProposal)(nil), "kava.community.v1beta1.CommunityPoolLendWithdrawProposal")
	proto.RegisterType((*CommunityCDPRepayDebtProposal)(nil), "kava.community.v1beta1.CommunityCDPRepayDebtProposal")
	proto.RegisterType((*CommunityCDPWithdrawCollateralProposal)(nil), "kava.community.v1beta1.CommunityCDPWithdrawCollateralProposal")
}

func init() {
	proto.RegisterFile("kava/community/v1beta1/proposal.proto", fileDescriptor_64aa83b2ed448ec1)
}

var fileDescriptor_64aa83b2ed448ec1 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x93, 0x3f, 0x0b, 0xd3, 0x40,
	0x18, 0xc6, 0x73, 0xb6, 0x56, 0xbd, 0x82, 0x42, 0x28, 0x12, 0x0b, 0x26, 0xb1, 0xa0, 0x16, 0xa4,
	0x39, 0xab, 0x93, 0x2e, 0x42, 0x53, 0x37, 0x87, 0x12, 0x04, 0xc1, 0x45, 0x2e, 0xc9, 0xd1, 0x1e,
	0x4d, 0xf2, 0x1e, 0xb9, 0x6b, 0x35, 0xdf, 0xc0, 0xd1, 0xd1, 0xb1, 0xb3, 0xdf, 0x43, 0xa8, 0x4e,
	0x1d, 0x1c, 0x9c, 0x54, 0xda, 0x2f, 0x22, 0xf9, 0xdb, 0x80, 0x20, 0x82, 0x20, 0x38, 0xe5, 0xcd,
	0x7b, 0xcf, 0x7b, 0xf7, 0xfc, 0x78, 0x78, 0xf1, 0xed, 0x35, 0xdd, 0x52, 0x12, 0x40, 0x1c, 0x6f,
	0x12, 0xae, 0x32, 0xb2, 0x9d, 0xfa, 0x4c, 0xd1, 0x29, 0x11, 0x29, 0x08, 0x90, 0x34, 0x72, 0x44,
	0x0a, 0x0a, 0xf4, 0xeb, 0xb9, 0xcc, 0x69, 0x64, 0x4e, 0x25, 0x1b, 0x9a, 0x01, 0xc8, 0x18, 0x24,
	0xf1, 0xa9, 0x64, 0xcd, 0x6c, 0x00, 0x3c, 0x29, 0xe7, 0x86, 0x83, 0x25, 0x2c, 0xa1, 0x28, 0x49,
	0x5e, 0x95, 0xdd, 0xd1, 0x27, 0x84, 0x6d, 0xb7, 0xbe, 0x6b, 0x01, 0x10, 0x3d, 0x63, 0x49, 0x38,
	0x67, 0x02, 0x24, 0x57, 0x8b, 0xea, 0x61, 0x7d, 0x80, 0x2f, 0x2a, 0xae, 0x22, 0x66, 0x20, 0x1b,
	0x8d, 0xaf, 0x78, 0xe5, 0x8f, 0x6e, 0xe3, 0x7e, 0xc8, 0x64, 0x90, 0x72, 0xa1, 0x38, 0x24, 0xc6,
	0x85, 0xe2, 0xac, 0xdd, 0xd2, 0x03, 0xdc, 0xa3, 0x31, 0x6c, 0x12, 0x65, 0x74, 0xec, 0xce, 0xb8,
	0xff, 0xe0, 0x86, 0x53, 0x7a, 0x74, 0x72, 0x8f, 0xb5, 0x71, 0xc7, 0x05, 0x9e, 0xcc, 0xee, 0xef,
	0xbf, 0x59, 0xda, 0x87, 0xef, 0xd6, 0x78, 0xc9, 0xd5, 0x6a, 0xe3, 0xe7, 0x7c, 0xa4, 0x02, 0x2a,
	0x3f, 0x13, 0x19, 0xae, 0x89, 0xca, 0x04, 0x93, 0xc5, 0x80, 0xf4, 0xaa, 0xab, 0x1f, 0x5f, 0x7e,
	0xbb, 0xb3, 0xb4, 0xf7, 0x3b, 0x4b, 0x1b, 0x7d, 0x46, 0xf8, 0xd6, 0x2f, 0x2c, 0x2f, 0xb8, 0x5a,
	0x85, 0x29, 0x7d, 0xfd, 0xbf, 0xc1, 0x7c, 0x44, 0xf8, 0x66, 0x03, 0xe3, 0xce, 0x17, 0x1e, 0x13,
	0x34, 0x9b, 0x33, 0xff, 0xef, 0x53, 0xb9, 0x8b, 0xaf, 0x05, 0x10, 0x45, 0x54, 0xb1, 0x94, 0x46,
	0xaf, 0x72, 0x17, 0x46, 0xa7, 0x50, 0x5d, 0x3d, 0xb7, 0x9f, 0x67, 0x82, 0xe9, 0x8f, 0xf0, 0x25,
	0x41, 0xb3, 0x98, 0x25, 0xca, 0xe8, 0xda, 0xe8, 0xf7, 0xc8, 0xdd, 0x1c, 0xd9, 0xab, 0xf5, 0x2d,
	0x8e, 0x2f, 0x08, 0xdf, 0x69, 0x73, 0xd4, 0x79, 0xb8, 0xcd, 0x5b, 0xff, 0x0e, 0xe8, 0x09, 0xc6,
	0xe7, 0xce, 0x9f, 0x32, 0xb5, 0x46, 0xce, 0x58, 0xb3, 0xa7, 0xfb, 0xa3, 0x89, 0x0e, 0x47, 0x13,
	0xfd, 0x38, 0x9a, 0xe8, 0xdd, 0xc9, 0xd4, 0x0e, 0x27, 0x53, 0xfb, 0x7a, 0x32, 0xb5, 0x97, 0xf7,
	0x5a, 0xa1, 0xe7, 0xab, 0x3a, 0x89, 0xa8, 0x2f, 0x8b, 0x8a, 0xbc, 0x69, 0x6d, 0x77, 0x91, 0xbe,
	0xdf, 0x2b, 0xb6, 0xf0, 0xe1, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0x58, 0xb9, 0x21, 0xfc,
	0x03, 0x00, 0x00,
}

func (m *CommunityPoolLendDepositProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommunityPoolLendDepositProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommunityPoolLendDepositProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommunityPoolLendWithdrawProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommunityPoolLendWithdrawProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommunityPoolLendWithdrawProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommunityCDPRepayDebtProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommunityCDPRepayDebtProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommunityCDPRepayDebtProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Payment.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.CollateralType) > 0 {
		i -= len(m.CollateralType)
		copy(dAtA[i:], m.CollateralType)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.CollateralType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommunityCDPWithdrawCollateralProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommunityCDPWithdrawCollateralProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommunityCDPWithdrawCollateralProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Collateral.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.CollateralType) > 0 {
		i -= len(m.CollateralType)
		copy(dAtA[i:], m.CollateralType)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.CollateralType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CommunityPoolLendDepositProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func (m *CommunityPoolLendWithdrawProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func (m *CommunityCDPRepayDebtProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.CollateralType)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = m.Payment.Size()
	n += 1 + l + sovProposal(uint64(l))
	return n
}

func (m *CommunityCDPWithdrawCollateralProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.CollateralType)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = m.Collateral.Size()
	n += 1 + l + sovProposal(uint64(l))
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommunityPoolLendDepositProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: CommunityPoolLendDepositProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommunityPoolLendDepositProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *CommunityPoolLendWithdrawProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: CommunityPoolLendWithdrawProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommunityPoolLendWithdrawProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *CommunityCDPRepayDebtProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: CommunityCDPRepayDebtProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommunityCDPRepayDebtProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Payment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *CommunityCDPWithdrawCollateralProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: CommunityCDPWithdrawCollateralProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommunityCDPWithdrawCollateralProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Collateral.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
