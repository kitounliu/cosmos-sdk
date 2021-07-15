// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/airdrop/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgAirDrop represents a message to create an airdrop fund for distribution
type MsgAirDrop struct {
	FromAddress string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	Fund        *Fund  `protobuf:"bytes,2,opt,name=fund,proto3" json:"fund,omitempty" yaml:"fund"`
}

func (m *MsgAirDrop) Reset()         { *m = MsgAirDrop{} }
func (m *MsgAirDrop) String() string { return proto.CompactTextString(m) }
func (*MsgAirDrop) ProtoMessage()    {}
func (*MsgAirDrop) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbea7cc5769c257f, []int{0}
}
func (m *MsgAirDrop) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAirDrop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAirDrop.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAirDrop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAirDrop.Merge(m, src)
}
func (m *MsgAirDrop) XXX_Size() int {
	return m.Size()
}
func (m *MsgAirDrop) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAirDrop.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAirDrop proto.InternalMessageInfo

// MsgAirDropResponse represents a message for the response
type MsgAirDropResponse struct {
}

func (m *MsgAirDropResponse) Reset()         { *m = MsgAirDropResponse{} }
func (m *MsgAirDropResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAirDropResponse) ProtoMessage()    {}
func (*MsgAirDropResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbea7cc5769c257f, []int{1}
}
func (m *MsgAirDropResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAirDropResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAirDropResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAirDropResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAirDropResponse.Merge(m, src)
}
func (m *MsgAirDropResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAirDropResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAirDropResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAirDropResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAirDrop)(nil), "cosmos.airdrop.v1beta1.MsgAirDrop")
	proto.RegisterType((*MsgAirDropResponse)(nil), "cosmos.airdrop.v1beta1.MsgAirDropResponse")
}

func init() { proto.RegisterFile("cosmos/airdrop/v1beta1/tx.proto", fileDescriptor_bbea7cc5769c257f) }

var fileDescriptor_bbea7cc5769c257f = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0xcc, 0x2c, 0x4a, 0x29, 0xca, 0x2f, 0xd0, 0x2f, 0x33, 0x4c, 0x4a, 0x2d,
	0x49, 0x34, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x83, 0x28, 0xd0,
	0x83, 0x2a, 0xd0, 0x83, 0x2a, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd1, 0x07, 0xb1,
	0x20, 0xaa, 0xa5, 0x54, 0x70, 0x18, 0x07, 0xd3, 0x0d, 0x56, 0xa5, 0x34, 0x95, 0x91, 0x8b, 0xcb,
	0xb7, 0x38, 0xdd, 0x31, 0xb3, 0xc8, 0xa5, 0x28, 0xbf, 0x40, 0xc8, 0x8a, 0x8b, 0x27, 0xad, 0x28,
	0x3f, 0x37, 0x3e, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x58, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3,
	0x49, 0xfc, 0xd3, 0x3d, 0x79, 0xe1, 0xca, 0xc4, 0xdc, 0x1c, 0x2b, 0x25, 0x64, 0x59, 0xa5, 0x20,
	0x6e, 0x10, 0xd7, 0x11, 0xc2, 0x13, 0x72, 0xe4, 0x62, 0x49, 0x2b, 0xcd, 0x4b, 0x91, 0x60, 0x52,
	0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd1, 0xc3, 0xee, 0x5a, 0x3d, 0xb7, 0xd2, 0xbc, 0x14, 0x27, 0xfe,
	0x4f, 0xf7, 0xe4, 0xb9, 0xa1, 0x26, 0x96, 0xe6, 0xa5, 0x28, 0x05, 0x81, 0xb5, 0x5a, 0x71, 0x74,
	0x2c, 0x90, 0x67, 0x78, 0xb1, 0x40, 0x9e, 0x41, 0x49, 0x84, 0x4b, 0x08, 0xe1, 0xac, 0xa0, 0xd4,
	0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xa3, 0x04, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0xa1, 0x48, 0x2e,
	0x76, 0x98, 0x83, 0x95, 0x70, 0x59, 0x83, 0xd0, 0x2d, 0xa5, 0x45, 0x58, 0x0d, 0xcc, 0x06, 0x27,
	0xb7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63,
	0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0x06, 0x2d, 0x84, 0xd2, 0x2d, 0x4e, 0xc9, 0xd6,
	0xaf, 0x80, 0x87, 0x73, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0x78, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x3d, 0x89, 0x5e, 0x78, 0xd5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// AirDrop defines a method for sending coins to the airdrop module for distribution
	AirDrop(ctx context.Context, in *MsgAirDrop, opts ...grpc.CallOption) (*MsgAirDropResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AirDrop(ctx context.Context, in *MsgAirDrop, opts ...grpc.CallOption) (*MsgAirDropResponse, error) {
	out := new(MsgAirDropResponse)
	err := c.cc.Invoke(ctx, "/cosmos.airdrop.v1beta1.Msg/AirDrop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// AirDrop defines a method for sending coins to the airdrop module for distribution
	AirDrop(context.Context, *MsgAirDrop) (*MsgAirDropResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AirDrop(ctx context.Context, req *MsgAirDrop) (*MsgAirDropResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AirDrop not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AirDrop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAirDrop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AirDrop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.airdrop.v1beta1.Msg/AirDrop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AirDrop(ctx, req.(*MsgAirDrop))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.airdrop.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AirDrop",
			Handler:    _Msg_AirDrop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/airdrop/v1beta1/tx.proto",
}

func (m *MsgAirDrop) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAirDrop) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAirDrop) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fund != nil {
		{
			size, err := m.Fund.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAirDropResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAirDropResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAirDropResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAirDrop) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Fund != nil {
		l = m.Fund.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAirDropResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAirDrop) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAirDrop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAirDrop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fund", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fund == nil {
				m.Fund = &Fund{}
			}
			if err := m.Fund.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAirDropResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAirDropResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAirDropResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
