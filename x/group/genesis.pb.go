// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/group/v1alpha1/genesis.proto

package group

import (
	fmt "fmt"
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

// GenesisState defines the group module's genesis state.
type GenesisState struct {
	// group_seq is the group table orm.Sequence,
	// it is used to get the next group ID.
	GroupSeq uint64 `protobuf:"varint,1,opt,name=group_seq,json=groupSeq,proto3" json:"group_seq,omitempty"`
	// groups is the list of groups info.
	Groups []*GroupInfo `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
	// group_members is the list of groups members.
	GroupMembers []*GroupMember `protobuf:"bytes,3,rep,name=group_members,json=groupMembers,proto3" json:"group_members,omitempty"`
	// group_account_seq is the group account table orm.Sequence,
	// it is used to generate the next group account address.
	GroupAccountSeq uint64 `protobuf:"varint,4,opt,name=group_account_seq,json=groupAccountSeq,proto3" json:"group_account_seq,omitempty"`
	// group_accounts is the list of group accounts info.
	GroupAccounts []*GroupAccountInfo `protobuf:"bytes,5,rep,name=group_accounts,json=groupAccounts,proto3" json:"group_accounts,omitempty"`
	// proposal_seq is the proposal table orm.Sequence,
	// it is used to get the next proposal ID.
	ProposalSeq uint64 `protobuf:"varint,6,opt,name=proposal_seq,json=proposalSeq,proto3" json:"proposal_seq,omitempty"`
	// proposals is the list of proposals.
	Proposals []*Proposal `protobuf:"bytes,7,rep,name=proposals,proto3" json:"proposals,omitempty"`
	// votes is the list of votes.
	Votes []*Vote `protobuf:"bytes,8,rep,name=votes,proto3" json:"votes,omitempty"`
	// polls is the list of polls.
	Polls []*Poll `protobuf:"bytes,9,rep,name=polls,proto3" json:"polls,omitempty"`
	// poll_seq is the poll table orm.Sequence,
	// it is used to get the next poll ID.
	PollSeq uint64 `protobuf:"varint,10,opt,name=poll_seq,json=pollSeq,proto3" json:"poll_seq,omitempty"`
	// votes is the list of votes for poll.
	VotesForPoll []*VotePoll `protobuf:"bytes,11,rep,name=votes_for_poll,json=votesForPoll,proto3" json:"votes_for_poll,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ccc5d002e96a4ab, []int{0}
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

func (m *GenesisState) GetGroupSeq() uint64 {
	if m != nil {
		return m.GroupSeq
	}
	return 0
}

func (m *GenesisState) GetGroups() []*GroupInfo {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *GenesisState) GetGroupMembers() []*GroupMember {
	if m != nil {
		return m.GroupMembers
	}
	return nil
}

func (m *GenesisState) GetGroupAccountSeq() uint64 {
	if m != nil {
		return m.GroupAccountSeq
	}
	return 0
}

func (m *GenesisState) GetGroupAccounts() []*GroupAccountInfo {
	if m != nil {
		return m.GroupAccounts
	}
	return nil
}

func (m *GenesisState) GetProposalSeq() uint64 {
	if m != nil {
		return m.ProposalSeq
	}
	return 0
}

func (m *GenesisState) GetProposals() []*Proposal {
	if m != nil {
		return m.Proposals
	}
	return nil
}

func (m *GenesisState) GetVotes() []*Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *GenesisState) GetPolls() []*Poll {
	if m != nil {
		return m.Polls
	}
	return nil
}

func (m *GenesisState) GetPollSeq() uint64 {
	if m != nil {
		return m.PollSeq
	}
	return 0
}

func (m *GenesisState) GetVotesForPoll() []*VotePoll {
	if m != nil {
		return m.VotesForPoll
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "regen.group.v1alpha1.GenesisState")
}

func init() {
	proto.RegisterFile("regen/group/v1alpha1/genesis.proto", fileDescriptor_6ccc5d002e96a4ab)
}

var fileDescriptor_6ccc5d002e96a4ab = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0xaa, 0xda, 0x40,
	0x14, 0xc7, 0x4d, 0xfd, 0x1e, 0xa3, 0xa5, 0x83, 0x8b, 0xa9, 0x85, 0x54, 0xa5, 0x14, 0x29, 0x34,
	0xa9, 0xed, 0xa2, 0x9b, 0x52, 0x68, 0x29, 0x4a, 0x17, 0x82, 0x44, 0xe8, 0xa2, 0x1b, 0x89, 0xe9,
	0x18, 0xa5, 0x49, 0x4e, 0xcc, 0x8c, 0xd2, 0xbe, 0x45, 0x1f, 0xab, 0x9b, 0x0b, 0x2e, 0xef, 0xf2,
	0xa2, 0x2f, 0x72, 0x99, 0x33, 0x11, 0xef, 0x85, 0x90, 0x55, 0xe6, 0x9c, 0xfc, 0xfe, 0x1f, 0x03,
	0x43, 0x86, 0x29, 0x0f, 0x78, 0xec, 0x04, 0x29, 0xec, 0x13, 0xe7, 0x30, 0xf6, 0xc2, 0x64, 0xe3,
	0x8d, 0x9d, 0x80, 0xc7, 0x5c, 0x6c, 0x85, 0x9d, 0xa4, 0x20, 0x81, 0x76, 0x91, 0xb1, 0x91, 0xb1,
	0x2f, 0x4c, 0xaf, 0x1b, 0x40, 0x00, 0x08, 0x38, 0xea, 0xa4, 0xd9, 0x5e, 0x3f, 0xd7, 0x4f, 0xfe,
	0x4d, 0x78, 0xe6, 0x36, 0xbc, 0xa9, 0x10, 0x73, 0xaa, 0xfd, 0x17, 0xd2, 0x93, 0x9c, 0xbe, 0x20,
	0x4d, 0xc4, 0x97, 0x82, 0xef, 0x98, 0xd1, 0x37, 0x46, 0x15, 0xb7, 0x81, 0x8b, 0x05, 0xdf, 0xd1,
	0x8f, 0xa4, 0x86, 0x67, 0xc1, 0x9e, 0xf4, 0xcb, 0xa3, 0xd6, 0xfb, 0x97, 0x76, 0x5e, 0x19, 0x7b,
	0xaa, 0xc6, 0xef, 0xf1, 0x1a, 0xdc, 0x0c, 0xa7, 0x13, 0xd2, 0xd6, 0xae, 0x11, 0x8f, 0x56, 0x3c,
	0x15, 0xac, 0x8c, 0xfa, 0x41, 0x81, 0x7e, 0x86, 0xa4, 0x6b, 0x06, 0xd7, 0x41, 0xd0, 0x37, 0xe4,
	0x99, 0xf6, 0xf1, 0x7c, 0x1f, 0xf6, 0xb1, 0xc4, 0x96, 0x15, 0x6c, 0xf9, 0x14, 0x7f, 0x7c, 0xd1,
	0x7b, 0x55, 0x76, 0x46, 0x3a, 0x8f, 0x58, 0xc1, 0xaa, 0x18, 0xfa, 0xba, 0x20, 0x34, 0x93, 0x63,
	0xf7, 0xf6, 0x43, 0x43, 0x41, 0x07, 0xc4, 0x4c, 0x52, 0x48, 0x40, 0x78, 0x21, 0xa6, 0xd6, 0x30,
	0xb5, 0x75, 0xd9, 0xa9, 0xc4, 0x4f, 0xa4, 0x79, 0x19, 0x05, 0xab, 0x63, 0x98, 0x95, 0x1f, 0x36,
	0xcf, 0x30, 0xf7, 0x2a, 0xa0, 0xef, 0x48, 0xf5, 0x00, 0x92, 0x0b, 0xd6, 0x40, 0x65, 0x2f, 0x5f,
	0xf9, 0x03, 0x24, 0x77, 0x35, 0xa8, 0x14, 0x09, 0x84, 0xa1, 0x60, 0xcd, 0x22, 0xc5, 0x1c, 0xc2,
	0xd0, 0xd5, 0x20, 0x7d, 0x4e, 0x1a, 0xea, 0x80, 0x17, 0x20, 0x78, 0x81, 0xba, 0x9a, 0x55, 0xf9,
	0x6f, 0xa4, 0x83, 0xae, 0xcb, 0x35, 0xa4, 0x4b, 0xb5, 0x64, 0xad, 0xa2, 0x1b, 0xa8, 0x1e, 0xe8,
	0x6c, 0xa2, 0x6a, 0x02, 0xa9, 0x9a, 0xbe, 0x7e, 0xfe, 0x7f, 0xb2, 0x8c, 0xe3, 0xc9, 0x32, 0xee,
	0x4e, 0x96, 0xf1, 0xef, 0x6c, 0x95, 0x8e, 0x67, 0xab, 0x74, 0x7b, 0xb6, 0x4a, 0x3f, 0x5f, 0x05,
	0x5b, 0xb9, 0xd9, 0xaf, 0x6c, 0x1f, 0x22, 0xc7, 0x07, 0x11, 0x81, 0xc8, 0x3e, 0x6f, 0xc5, 0xaf,
	0xdf, 0xce, 0x1f, 0xfd, 0x48, 0x57, 0x35, 0x7c, 0x96, 0x1f, 0xee, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x2c, 0xd4, 0x57, 0x01, 0x0a, 0x03, 0x00, 0x00,
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
	if len(m.VotesForPoll) > 0 {
		for iNdEx := len(m.VotesForPoll) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VotesForPoll[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if m.PollSeq != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PollSeq))
		i--
		dAtA[i] = 0x50
	}
	if len(m.Polls) > 0 {
		for iNdEx := len(m.Polls) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Polls[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Proposals) > 0 {
		for iNdEx := len(m.Proposals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proposals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.ProposalSeq != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ProposalSeq))
		i--
		dAtA[i] = 0x30
	}
	if len(m.GroupAccounts) > 0 {
		for iNdEx := len(m.GroupAccounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GroupAccounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.GroupAccountSeq != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.GroupAccountSeq))
		i--
		dAtA[i] = 0x20
	}
	if len(m.GroupMembers) > 0 {
		for iNdEx := len(m.GroupMembers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GroupMembers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Groups) > 0 {
		for iNdEx := len(m.Groups) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Groups[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.GroupSeq != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.GroupSeq))
		i--
		dAtA[i] = 0x8
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
	if m.GroupSeq != 0 {
		n += 1 + sovGenesis(uint64(m.GroupSeq))
	}
	if len(m.Groups) > 0 {
		for _, e := range m.Groups {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GroupMembers) > 0 {
		for _, e := range m.GroupMembers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.GroupAccountSeq != 0 {
		n += 1 + sovGenesis(uint64(m.GroupAccountSeq))
	}
	if len(m.GroupAccounts) > 0 {
		for _, e := range m.GroupAccounts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ProposalSeq != 0 {
		n += 1 + sovGenesis(uint64(m.ProposalSeq))
	}
	if len(m.Proposals) > 0 {
		for _, e := range m.Proposals {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Polls) > 0 {
		for _, e := range m.Polls {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.PollSeq != 0 {
		n += 1 + sovGenesis(uint64(m.PollSeq))
	}
	if len(m.VotesForPoll) > 0 {
		for _, e := range m.VotesForPoll {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupSeq", wireType)
			}
			m.GroupSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Groups", wireType)
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
			m.Groups = append(m.Groups, &GroupInfo{})
			if err := m.Groups[len(m.Groups)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupMembers", wireType)
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
			m.GroupMembers = append(m.GroupMembers, &GroupMember{})
			if err := m.GroupMembers[len(m.GroupMembers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupAccountSeq", wireType)
			}
			m.GroupAccountSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupAccountSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupAccounts", wireType)
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
			m.GroupAccounts = append(m.GroupAccounts, &GroupAccountInfo{})
			if err := m.GroupAccounts[len(m.GroupAccounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalSeq", wireType)
			}
			m.ProposalSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposals", wireType)
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
			m.Proposals = append(m.Proposals, &Proposal{})
			if err := m.Proposals[len(m.Proposals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
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
			m.Votes = append(m.Votes, &Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Polls", wireType)
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
			m.Polls = append(m.Polls, &Poll{})
			if err := m.Polls[len(m.Polls)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PollSeq", wireType)
			}
			m.PollSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PollSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotesForPoll", wireType)
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
			m.VotesForPoll = append(m.VotesForPoll, &VotePoll{})
			if err := m.VotesForPoll[len(m.VotesForPoll)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
