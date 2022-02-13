// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ssi/did.proto

package types

import (
	fmt "fmt"
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

type Did struct {
	Context              []string           `protobuf:"bytes,1,rep,name=context,proto3" json:"context,omitempty"`
	Type                 string             `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Id                   string             `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string             `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	PublicKey            []*PublicKeyStruct `protobuf:"bytes,5,rep,name=publicKey,proto3" json:"publicKey,omitempty"`
	Authentication       []string           `protobuf:"bytes,6,rep,name=authentication,proto3" json:"authentication,omitempty"`
	AssertionMethod      []string           `protobuf:"bytes,7,rep,name=assertionMethod,proto3" json:"assertionMethod,omitempty"`
	KeyAgreement         []string           `protobuf:"bytes,8,rep,name=keyAgreement,proto3" json:"keyAgreement,omitempty"`
	CapabilityInvocation []string           `protobuf:"bytes,9,rep,name=capabilityInvocation,proto3" json:"capabilityInvocation,omitempty"`
	Created              string             `protobuf:"bytes,10,opt,name=created,proto3" json:"created,omitempty"`
	Updated              string             `protobuf:"bytes,11,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (m *Did) Reset()         { *m = Did{} }
func (m *Did) String() string { return proto.CompactTextString(m) }
func (*Did) ProtoMessage()    {}
func (*Did) Descriptor() ([]byte, []int) {
	return fileDescriptor_1200da98fc4194d9, []int{0}
}
func (m *Did) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Did) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Did.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Did) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Did.Merge(m, src)
}
func (m *Did) XXX_Size() int {
	return m.Size()
}
func (m *Did) XXX_DiscardUnknown() {
	xxx_messageInfo_Did.DiscardUnknown(m)
}

var xxx_messageInfo_Did proto.InternalMessageInfo

func (m *Did) GetContext() []string {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Did) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Did) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Did) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Did) GetPublicKey() []*PublicKeyStruct {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Did) GetAuthentication() []string {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *Did) GetAssertionMethod() []string {
	if m != nil {
		return m.AssertionMethod
	}
	return nil
}

func (m *Did) GetKeyAgreement() []string {
	if m != nil {
		return m.KeyAgreement
	}
	return nil
}

func (m *Did) GetCapabilityInvocation() []string {
	if m != nil {
		return m.CapabilityInvocation
	}
	return nil
}

func (m *Did) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Did) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

func init() {
	proto.RegisterType((*Did)(nil), "hypersignprotocol.hidnode.ssi.Did")
}

func init() { proto.RegisterFile("ssi/did.proto", fileDescriptor_1200da98fc4194d9) }

var fileDescriptor_1200da98fc4194d9 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x9b, 0xa6, 0xb7, 0xbd, 0x99, 0xf6, 0xf6, 0xc2, 0x70, 0x17, 0xc3, 0x05, 0x43, 0xe9,
	0x42, 0xb2, 0x69, 0x02, 0xf5, 0x09, 0x14, 0x37, 0x62, 0x05, 0xa9, 0x3b, 0x77, 0x49, 0xe6, 0xd0,
	0x1c, 0x6c, 0x67, 0x42, 0xe6, 0x44, 0x9a, 0xb5, 0x2f, 0xe0, 0x63, 0xb9, 0xec, 0xd2, 0xa5, 0xb4,
	0x2f, 0x22, 0x99, 0x36, 0x8a, 0x45, 0xdc, 0x9d, 0xff, 0xff, 0x3f, 0x66, 0x0e, 0xff, 0x61, 0x7f,
	0x8c, 0xc1, 0x48, 0xa2, 0x0c, 0xf3, 0x42, 0x93, 0xe6, 0x27, 0x59, 0x95, 0x43, 0x61, 0x70, 0xa1,
	0xac, 0x4e, 0xf5, 0x32, 0xcc, 0x50, 0x2a, 0x2d, 0x21, 0x34, 0x06, 0xff, 0x0f, 0x6a, 0x9a, 0xd6,
	0x7b, 0x78, 0xfc, 0xe4, 0x32, 0xf7, 0x12, 0x25, 0x17, 0xac, 0x97, 0x6a, 0x45, 0xb0, 0x26, 0xe1,
	0x8c, 0xdc, 0xc0, 0x9b, 0x37, 0x92, 0x73, 0xd6, 0xa1, 0x2a, 0x07, 0xd1, 0x1e, 0x39, 0x81, 0x37,
	0xb7, 0x33, 0x1f, 0xb2, 0x36, 0x4a, 0xe1, 0x5a, 0xa7, 0x8d, 0xb2, 0x66, 0x54, 0xbc, 0x02, 0xd1,
	0xd9, 0x33, 0xf5, 0xcc, 0x67, 0xcc, 0xcb, 0xcb, 0x64, 0x89, 0xe9, 0x35, 0x54, 0xe2, 0xd7, 0xc8,
	0x0d, 0xfa, 0xd3, 0x30, 0xfc, 0x71, 0xb5, 0xf0, 0xb6, 0xe1, 0xef, 0xa8, 0x28, 0x53, 0x9a, 0x7f,
	0x3e, 0xc0, 0x4f, 0xd9, 0x30, 0x2e, 0x29, 0x03, 0x45, 0x98, 0xc6, 0x84, 0x5a, 0x89, 0xae, 0x5d,
	0xf3, 0xc8, 0xe5, 0x01, 0xfb, 0x1b, 0x1b, 0x03, 0x45, 0x2d, 0x6e, 0x80, 0x32, 0x2d, 0x45, 0xcf,
	0x82, 0xc7, 0x36, 0x1f, 0xb3, 0xc1, 0x03, 0x54, 0xe7, 0x8b, 0x02, 0x60, 0x05, 0x8a, 0xc4, 0x6f,
	0x8b, 0x7d, 0xf1, 0xf8, 0x94, 0xfd, 0x4b, 0xe3, 0x3c, 0x4e, 0x70, 0x89, 0x54, 0x5d, 0xa9, 0x47,
	0x7d, 0xf8, 0xdb, 0xb3, 0xec, 0xb7, 0x99, 0x6d, 0xb2, 0x80, 0x98, 0x40, 0x0a, 0x66, 0xeb, 0x68,
	0x64, 0x9d, 0x94, 0xb9, 0xb4, 0x49, 0x7f, 0x9f, 0x1c, 0xe4, 0xc5, 0xec, 0x65, 0xeb, 0x3b, 0x9b,
	0xad, 0xef, 0xbc, 0x6d, 0x7d, 0xe7, 0x79, 0xe7, 0xb7, 0x36, 0x3b, 0xbf, 0xf5, 0xba, 0xf3, 0x5b,
	0xf7, 0xd3, 0x05, 0x52, 0x56, 0x26, 0x61, 0xaa, 0x57, 0xd1, 0x47, 0x79, 0x93, 0xa6, 0xbd, 0x28,
	0x43, 0x39, 0xa9, 0xeb, 0x8b, 0xd6, 0x91, 0x3d, 0x6b, 0x95, 0x83, 0x49, 0xba, 0x36, 0x3e, 0x7b,
	0x0f, 0x00, 0x00, 0xff, 0xff, 0x40, 0x5f, 0xbe, 0xf7, 0x18, 0x02, 0x00, 0x00,
}

func (m *Did) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Did) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Did) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Updated) > 0 {
		i -= len(m.Updated)
		copy(dAtA[i:], m.Updated)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Updated)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Created) > 0 {
		i -= len(m.Created)
		copy(dAtA[i:], m.Created)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Created)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.CapabilityInvocation) > 0 {
		for iNdEx := len(m.CapabilityInvocation) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CapabilityInvocation[iNdEx])
			copy(dAtA[i:], m.CapabilityInvocation[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.CapabilityInvocation[iNdEx])))
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.KeyAgreement) > 0 {
		for iNdEx := len(m.KeyAgreement) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.KeyAgreement[iNdEx])
			copy(dAtA[i:], m.KeyAgreement[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.KeyAgreement[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.AssertionMethod) > 0 {
		for iNdEx := len(m.AssertionMethod) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AssertionMethod[iNdEx])
			copy(dAtA[i:], m.AssertionMethod[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.AssertionMethod[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Authentication) > 0 {
		for iNdEx := len(m.Authentication) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Authentication[iNdEx])
			copy(dAtA[i:], m.Authentication[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.Authentication[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.PublicKey) > 0 {
		for iNdEx := len(m.PublicKey) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PublicKey[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDid(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Context) > 0 {
		for iNdEx := len(m.Context) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Context[iNdEx])
			copy(dAtA[i:], m.Context[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.Context[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintDid(dAtA []byte, offset int, v uint64) int {
	offset -= sovDid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Did) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Context) > 0 {
		for _, s := range m.Context {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	if len(m.PublicKey) > 0 {
		for _, e := range m.PublicKey {
			l = e.Size()
			n += 1 + l + sovDid(uint64(l))
		}
	}
	if len(m.Authentication) > 0 {
		for _, s := range m.Authentication {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	if len(m.AssertionMethod) > 0 {
		for _, s := range m.AssertionMethod {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	if len(m.KeyAgreement) > 0 {
		for _, s := range m.KeyAgreement {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	if len(m.CapabilityInvocation) > 0 {
		for _, s := range m.CapabilityInvocation {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	l = len(m.Created)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.Updated)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	return n
}

func sovDid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDid(x uint64) (n int) {
	return sovDid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Did) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDid
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
			return fmt.Errorf("proto: Did: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Did: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Context = append(m.Context, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = append(m.PublicKey, &PublicKeyStruct{})
			if err := m.PublicKey[len(m.PublicKey)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authentication", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authentication = append(m.Authentication, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssertionMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssertionMethod = append(m.AssertionMethod, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyAgreement", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyAgreement = append(m.KeyAgreement, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapabilityInvocation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CapabilityInvocation = append(m.CapabilityInvocation, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Created = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Updated", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Updated = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDid
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
func skipDid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDid
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
					return 0, ErrIntOverflowDid
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
					return 0, ErrIntOverflowDid
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
				return 0, ErrInvalidLengthDid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDid = fmt.Errorf("proto: unexpected end of group")
)
