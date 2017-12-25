// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: coprocessor.proto

/*
	Package coprocessor is a generated protocol buffer package.

	It is generated from these files:
		coprocessor.proto

	It has these top-level messages:
		KeyRange
		Request
		Response
*/
package coprocessor

import (
	"fmt"
	"io"
	"math"

	proto "github.com/golang/protobuf/proto"

	errorpb "github.com/pingcap/kvproto/pkg/errorpb"

	kvrpcpb "github.com/pingcap/kvproto/pkg/kvrpcpb"

	github_com_pingcap_tipb_sharedbytes "github.com/pingcap/tipb/sharedbytes"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// [start, end)
type KeyRange struct {
	Start []byte `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   []byte `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (m *KeyRange) Reset()                    { *m = KeyRange{} }
func (m *KeyRange) String() string            { return proto.CompactTextString(m) }
func (*KeyRange) ProtoMessage()               {}
func (*KeyRange) Descriptor() ([]byte, []int) { return fileDescriptorCoprocessor, []int{0} }

func (m *KeyRange) GetStart() []byte {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *KeyRange) GetEnd() []byte {
	if m != nil {
		return m.End
	}
	return nil
}

type Request struct {
	Context *kvrpcpb.Context `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	Tp      int64            `protobuf:"varint,2,opt,name=tp,proto3" json:"tp,omitempty"`
	Data    []byte           `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Ranges  []*KeyRange      `protobuf:"bytes,4,rep,name=ranges" json:"ranges,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptorCoprocessor, []int{1} }

func (m *Request) GetContext() *kvrpcpb.Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Request) GetTp() int64 {
	if m != nil {
		return m.Tp
	}
	return 0
}

func (m *Request) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Request) GetRanges() []*KeyRange {
	if m != nil {
		return m.Ranges
	}
	return nil
}

type Response struct {
	Data        github_com_pingcap_tipb_sharedbytes.SharedBytes `protobuf:"bytes,1,opt,name=data,proto3,customtype=github.com/pingcap/tipb/sharedbytes.SharedBytes" json:"data"`
	RegionError *errorpb.Error                                  `protobuf:"bytes,2,opt,name=region_error,json=regionError" json:"region_error,omitempty"`
	Locked      *kvrpcpb.LockInfo                               `protobuf:"bytes,3,opt,name=locked" json:"locked,omitempty"`
	OtherError  string                                          `protobuf:"bytes,4,opt,name=other_error,json=otherError,proto3" json:"other_error,omitempty"`
	Range       *KeyRange                                       `protobuf:"bytes,5,opt,name=range" json:"range,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptorCoprocessor, []int{2} }

func (m *Response) GetRegionError() *errorpb.Error {
	if m != nil {
		return m.RegionError
	}
	return nil
}

func (m *Response) GetLocked() *kvrpcpb.LockInfo {
	if m != nil {
		return m.Locked
	}
	return nil
}

func (m *Response) GetOtherError() string {
	if m != nil {
		return m.OtherError
	}
	return ""
}

func (m *Response) GetRange() *KeyRange {
	if m != nil {
		return m.Range
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyRange)(nil), "coprocessor.KeyRange")
	proto.RegisterType((*Request)(nil), "coprocessor.Request")
	proto.RegisterType((*Response)(nil), "coprocessor.Response")
}
func (m *KeyRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyRange) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Start) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(len(m.Start)))
		i += copy(dAtA[i:], m.Start)
	}
	if len(m.End) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(len(m.End)))
		i += copy(dAtA[i:], m.End)
	}
	return i, nil
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Context != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(m.Context.Size()))
		n1, err := m.Context.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Tp != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(m.Tp))
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if len(m.Ranges) > 0 {
		for _, msg := range m.Ranges {
			dAtA[i] = 0x22
			i++
			i = encodeVarintCoprocessor(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintCoprocessor(dAtA, i, uint64(m.Data.Size()))
	n2, err := m.Data.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.RegionError != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(m.RegionError.Size()))
		n3, err := m.RegionError.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Locked != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(m.Locked.Size()))
		n4, err := m.Locked.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if len(m.OtherError) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(len(m.OtherError)))
		i += copy(dAtA[i:], m.OtherError)
	}
	if m.Range != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintCoprocessor(dAtA, i, uint64(m.Range.Size()))
		n5, err := m.Range.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}

func encodeVarintCoprocessor(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *KeyRange) Size() (n int) {
	var l int
	_ = l
	l = len(m.Start)
	if l > 0 {
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	l = len(m.End)
	if l > 0 {
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	return n
}

func (m *Request) Size() (n int) {
	var l int
	_ = l
	if m.Context != nil {
		l = m.Context.Size()
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	if m.Tp != 0 {
		n += 1 + sovCoprocessor(uint64(m.Tp))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	if len(m.Ranges) > 0 {
		for _, e := range m.Ranges {
			l = e.Size()
			n += 1 + l + sovCoprocessor(uint64(l))
		}
	}
	return n
}

func (m *Response) Size() (n int) {
	var l int
	_ = l
	l = m.Data.Size()
	n += 1 + l + sovCoprocessor(uint64(l))
	if m.RegionError != nil {
		l = m.RegionError.Size()
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	if m.Locked != nil {
		l = m.Locked.Size()
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	l = len(m.OtherError)
	if l > 0 {
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	if m.Range != nil {
		l = m.Range.Size()
		n += 1 + l + sovCoprocessor(uint64(l))
	}
	return n
}

func sovCoprocessor(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCoprocessor(x uint64) (n int) {
	return sovCoprocessor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeyRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoprocessor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: KeyRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Start = append(m.Start[:0], dAtA[iNdEx:postIndex]...)
			if m.Start == nil {
				m.Start = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.End = append(m.End[:0], dAtA[iNdEx:postIndex]...)
			if m.End == nil {
				m.End = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoprocessor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCoprocessor
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
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoprocessor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Context == nil {
				m.Context = &kvrpcpb.Context{}
			}
			if err := m.Context.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tp", wireType)
			}
			m.Tp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ranges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ranges = append(m.Ranges, &KeyRange{})
			if err := m.Ranges[len(m.Ranges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoprocessor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCoprocessor
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
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoprocessor
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionError", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RegionError == nil {
				m.RegionError = &errorpb.Error{}
			}
			if err := m.RegionError.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locked", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Locked == nil {
				m.Locked = &kvrpcpb.LockInfo{}
			}
			if err := m.Locked.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OtherError", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OtherError = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Range", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoprocessor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Range == nil {
				m.Range = &KeyRange{}
			}
			if err := m.Range.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoprocessor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCoprocessor
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
func skipCoprocessor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCoprocessor
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
					return 0, ErrIntOverflowCoprocessor
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCoprocessor
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCoprocessor
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCoprocessor
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCoprocessor(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCoprocessor = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCoprocessor   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("coprocessor.proto", fileDescriptorCoprocessor) }

var fileDescriptorCoprocessor = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xb1, 0x6e, 0xd4, 0x40,
	0x10, 0x8d, 0xed, 0xbb, 0x4b, 0x18, 0x27, 0x51, 0xb2, 0x0a, 0x92, 0x95, 0xc2, 0x77, 0xba, 0xea,
	0x08, 0x62, 0x2d, 0x4c, 0x41, 0x7f, 0x88, 0x02, 0x85, 0x6a, 0xf9, 0x00, 0x64, 0xaf, 0x17, 0x9f,
	0x65, 0xe2, 0x59, 0x76, 0x37, 0x27, 0x52, 0xd3, 0xf0, 0x09, 0x7c, 0x52, 0x4a, 0x6a, 0x8a, 0x08,
	0x1d, 0x3f, 0x82, 0x3c, 0x6b, 0x47, 0xd7, 0xa4, 0x9a, 0x37, 0xcf, 0xcf, 0xcf, 0x6f, 0x9e, 0xe1,
	0x5c, 0xa2, 0x36, 0x28, 0x95, 0xb5, 0x68, 0xb8, 0x36, 0xe8, 0x90, 0xc5, 0x7b, 0xd4, 0xe5, 0x89,
	0x32, 0x06, 0x8d, 0x2e, 0xfd, 0xb3, 0xcb, 0x93, 0x76, 0x6b, 0xb4, 0x7c, 0x5c, 0x2f, 0x6a, 0xac,
	0x91, 0x60, 0xd6, 0x23, 0xcf, 0x2e, 0x73, 0x38, 0xba, 0x56, 0x77, 0xa2, 0xe8, 0x6a, 0xc5, 0x2e,
	0x60, 0x6a, 0x5d, 0x61, 0x5c, 0x12, 0x2c, 0x82, 0xd5, 0xb1, 0xf0, 0x0b, 0x3b, 0x83, 0x48, 0x75,
	0x55, 0x12, 0x12, 0xd7, 0xc3, 0xe5, 0xcf, 0x00, 0x0e, 0x85, 0xfa, 0x76, 0xab, 0xac, 0x63, 0x57,
	0x70, 0x28, 0xb1, 0x73, 0xea, 0xbb, 0x7f, 0x2b, 0xce, 0xcf, 0xf8, 0xf8, 0xd9, 0x77, 0x9e, 0x17,
	0xa3, 0x80, 0x9d, 0x42, 0xe8, 0x34, 0x19, 0x45, 0x22, 0x74, 0x9a, 0x31, 0x98, 0x54, 0x85, 0x2b,
	0x92, 0x88, 0xac, 0x09, 0xb3, 0x57, 0x30, 0x33, 0x7d, 0x18, 0x9b, 0x4c, 0x16, 0xd1, 0x2a, 0xce,
	0x9f, 0xf3, 0xfd, 0xa3, 0xc7, 0xa8, 0x62, 0x10, 0x2d, 0x7f, 0x84, 0x70, 0x24, 0x94, 0xd5, 0xd8,
	0x59, 0xc5, 0xae, 0x07, 0x3f, 0x8a, 0xbf, 0x7e, 0x7b, 0xff, 0x30, 0x3f, 0xf8, 0xf3, 0x30, 0xcf,
	0xea, 0xc6, 0x6d, 0x6e, 0x4b, 0x2e, 0xf1, 0x26, 0xd3, 0x4d, 0x57, 0xcb, 0x42, 0x67, 0xae, 0xd1,
	0x65, 0x66, 0x37, 0x85, 0x51, 0x55, 0x79, 0xe7, 0x94, 0xe5, 0x9f, 0x08, 0xaf, 0x7b, 0x3c, 0x04,
	0x79, 0x0d, 0xc7, 0x46, 0xd5, 0x0d, 0x76, 0x9f, 0xa9, 0x55, 0x8a, 0x1d, 0xe7, 0xa7, 0x7c, 0xec,
	0xf8, 0x7d, 0x3f, 0x45, 0xec, 0x35, 0xb4, 0xb0, 0x17, 0x30, 0xfb, 0x8a, 0xb2, 0x55, 0x15, 0x5d,
	0x14, 0xe7, 0xe7, 0x8f, 0x55, 0x7c, 0x44, 0xd9, 0x7e, 0xe8, 0xbe, 0xa0, 0x18, 0x04, 0x6c, 0x0e,
	0x31, 0xba, 0x8d, 0x32, 0x83, 0xf9, 0x64, 0x11, 0xac, 0x9e, 0x09, 0x20, 0xca, 0x7b, 0xbd, 0x84,
	0x29, 0x9d, 0x98, 0x4c, 0xc9, 0xea, 0x89, 0x1a, 0xbc, 0x66, 0x7d, 0x75, 0xbf, 0x4b, 0x83, 0xdf,
	0xbb, 0x34, 0xf8, 0xbb, 0x4b, 0x83, 0x5f, 0xff, 0xd2, 0x03, 0x48, 0x24, 0xde, 0xf0, 0xe1, 0x5a,
	0xee, 0x9a, 0x76, 0xcb, 0xdb, 0x2d, 0xfd, 0xf0, 0x72, 0x46, 0xe3, 0xcd, 0xff, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xa7, 0x5f, 0x92, 0x38, 0x4d, 0x02, 0x00, 0x00,
}
