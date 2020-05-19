// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/trace/v4alpha/xray.proto

package envoy_config_trace_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v4alpha "github.com/datawire/ambassador/pkg/api/envoy/config/core/v4alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type XRayConfig struct {
	// The UDP endpoint of the X-Ray Daemon where the spans will be sent.
	// If this value is not set, the default value of 127.0.0.1:2000 will be used.
	DaemonEndpoint *v4alpha.SocketAddress `protobuf:"bytes,1,opt,name=daemon_endpoint,json=daemonEndpoint,proto3" json:"daemon_endpoint,omitempty"`
	// The name of the X-Ray segment.
	SegmentName string `protobuf:"bytes,2,opt,name=segment_name,json=segmentName,proto3" json:"segment_name,omitempty"`
	// The location of a local custom sampling rules JSON file.
	// For an example of the sampling rules see:
	// `X-Ray SDK documentation
	// <https://docs.aws.amazon.com/xray/latest/devguide/xray-sdk-go-configuration.html#xray-sdk-go-configuration-sampling>`_
	SamplingRuleManifest *v4alpha.DataSource `protobuf:"bytes,3,opt,name=sampling_rule_manifest,json=samplingRuleManifest,proto3" json:"sampling_rule_manifest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *XRayConfig) Reset()         { *m = XRayConfig{} }
func (m *XRayConfig) String() string { return proto.CompactTextString(m) }
func (*XRayConfig) ProtoMessage()    {}
func (*XRayConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_edb5acf508aeba3e, []int{0}
}
func (m *XRayConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *XRayConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_XRayConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *XRayConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XRayConfig.Merge(m, src)
}
func (m *XRayConfig) XXX_Size() int {
	return m.Size()
}
func (m *XRayConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_XRayConfig.DiscardUnknown(m)
}

var xxx_messageInfo_XRayConfig proto.InternalMessageInfo

func (m *XRayConfig) GetDaemonEndpoint() *v4alpha.SocketAddress {
	if m != nil {
		return m.DaemonEndpoint
	}
	return nil
}

func (m *XRayConfig) GetSegmentName() string {
	if m != nil {
		return m.SegmentName
	}
	return ""
}

func (m *XRayConfig) GetSamplingRuleManifest() *v4alpha.DataSource {
	if m != nil {
		return m.SamplingRuleManifest
	}
	return nil
}

func init() {
	proto.RegisterType((*XRayConfig)(nil), "envoy.config.trace.v4alpha.XRayConfig")
}

func init() {
	proto.RegisterFile("envoy/config/trace/v4alpha/xray.proto", fileDescriptor_edb5acf508aeba3e)
}

var fileDescriptor_edb5acf508aeba3e = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x99, 0x28, 0x95, 0x4e, 0x45, 0x4b, 0x10, 0x2d, 0x05, 0x43, 0x2d, 0x96, 0x06, 0x17,
	0x13, 0xb0, 0xae, 0xdc, 0x19, 0x75, 0x25, 0x4a, 0x4d, 0x37, 0x05, 0x17, 0xe1, 0x34, 0x39, 0x8d,
	0x83, 0xc9, 0x4c, 0x98, 0x99, 0x84, 0x66, 0xe7, 0xd2, 0xad, 0x5b, 0x1f, 0xc5, 0xbd, 0xd0, 0xa5,
	0xbe, 0x81, 0xf4, 0x31, 0x5c, 0x49, 0xf3, 0x87, 0x52, 0xee, 0xed, 0xdd, 0x05, 0xce, 0x2f, 0xbf,
	0xf9, 0xce, 0x77, 0xe8, 0x0c, 0x45, 0x29, 0x2b, 0x2f, 0x92, 0x62, 0xcb, 0x13, 0xcf, 0x28, 0x88,
	0xd0, 0x2b, 0x5f, 0x40, 0x9a, 0x7f, 0x06, 0x6f, 0xa7, 0xa0, 0x62, 0xb9, 0x92, 0x46, 0xda, 0xe3,
	0x1a, 0x63, 0x0d, 0xc6, 0x6a, 0x8c, 0xb5, 0xd8, 0x78, 0x7e, 0xa6, 0x88, 0xa4, 0x3a, 0x19, 0x20,
	0x8e, 0x15, 0x6a, 0xdd, 0x48, 0xc6, 0x4f, 0x2f, 0x83, 0x1b, 0xd0, 0xd8, 0x52, 0x8f, 0x8b, 0x38,
	0x07, 0x0f, 0x84, 0x90, 0x06, 0x0c, 0x97, 0x42, 0x7b, 0xda, 0x80, 0x29, 0x3a, 0xc9, 0x93, 0x2b,
	0xe3, 0x12, 0x95, 0xe6, 0x52, 0x70, 0x91, 0xb4, 0xc8, 0xa3, 0x12, 0x52, 0x1e, 0x83, 0x41, 0xaf,
	0xfb, 0x68, 0x06, 0xd3, 0xef, 0x16, 0xa5, 0xeb, 0x00, 0xaa, 0xd7, 0x75, 0x02, 0xfb, 0x23, 0xbd,
	0x1f, 0x03, 0x66, 0x52, 0x84, 0x28, 0xe2, 0x5c, 0x72, 0x61, 0x46, 0x64, 0x42, 0xdc, 0xc1, 0x73,
	0x97, 0x9d, 0xad, 0x7b, 0x4c, 0xda, 0x6d, 0xcb, 0x56, 0x32, 0xfa, 0x82, 0xe6, 0x55, 0xb3, 0x58,
	0x70, 0xaf, 0x11, 0xbc, 0x6d, 0xff, 0xb7, 0x9f, 0xd1, 0xbb, 0x1a, 0x93, 0x0c, 0x85, 0x09, 0x05,
	0x64, 0x38, 0xb2, 0x26, 0xc4, 0xed, 0xfb, 0x77, 0xfe, 0xf9, 0xb7, 0x95, 0x35, 0x24, 0xc1, 0xa0,
	0x1d, 0x7e, 0x80, 0x0c, 0xed, 0x4f, 0xf4, 0xa1, 0x86, 0x2c, 0x4f, 0xb9, 0x48, 0x42, 0x55, 0xa4,
	0x18, 0x66, 0x20, 0xf8, 0x16, 0xb5, 0x19, 0xdd, 0xaa, 0x53, 0xcc, 0x6e, 0x48, 0xf1, 0x06, 0x0c,
	0xac, 0x64, 0xa1, 0x22, 0x0c, 0x1e, 0x74, 0x92, 0xa0, 0x48, 0xf1, 0x7d, 0xab, 0x78, 0x39, 0xff,
	0xf1, 0xeb, 0x9b, 0x33, 0xa5, 0x93, 0xeb, 0xee, 0xb6, 0x60, 0xa7, 0x12, 0xfc, 0x77, 0xfb, 0x83,
	0x43, 0x7e, 0x1f, 0x1c, 0xf2, 0xf7, 0xe0, 0x90, 0x9f, 0x5f, 0xf7, 0x7f, 0x7a, 0xd6, 0xd0, 0xa2,
	0x2e, 0x97, 0xcd, 0xeb, 0xb9, 0x92, 0xbb, 0x8a, 0x5d, 0xbe, 0xbe, 0xdf, 0x5f, 0x2b, 0xa8, 0x96,
	0xc7, 0x7a, 0x97, 0x64, 0xd3, 0xab, 0x7b, 0x5e, 0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x66, 0xcc,
	0x43, 0x32, 0x56, 0x02, 0x00, 0x00,
}

func (m *XRayConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *XRayConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *XRayConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.SamplingRuleManifest != nil {
		{
			size, err := m.SamplingRuleManifest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintXray(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SegmentName) > 0 {
		i -= len(m.SegmentName)
		copy(dAtA[i:], m.SegmentName)
		i = encodeVarintXray(dAtA, i, uint64(len(m.SegmentName)))
		i--
		dAtA[i] = 0x12
	}
	if m.DaemonEndpoint != nil {
		{
			size, err := m.DaemonEndpoint.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintXray(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintXray(dAtA []byte, offset int, v uint64) int {
	offset -= sovXray(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *XRayConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DaemonEndpoint != nil {
		l = m.DaemonEndpoint.Size()
		n += 1 + l + sovXray(uint64(l))
	}
	l = len(m.SegmentName)
	if l > 0 {
		n += 1 + l + sovXray(uint64(l))
	}
	if m.SamplingRuleManifest != nil {
		l = m.SamplingRuleManifest.Size()
		n += 1 + l + sovXray(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovXray(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozXray(x uint64) (n int) {
	return sovXray(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *XRayConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowXray
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
			return fmt.Errorf("proto: XRayConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: XRayConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaemonEndpoint", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowXray
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
				return ErrInvalidLengthXray
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthXray
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DaemonEndpoint == nil {
				m.DaemonEndpoint = &v4alpha.SocketAddress{}
			}
			if err := m.DaemonEndpoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SegmentName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowXray
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
				return ErrInvalidLengthXray
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthXray
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SegmentName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SamplingRuleManifest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowXray
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
				return ErrInvalidLengthXray
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthXray
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SamplingRuleManifest == nil {
				m.SamplingRuleManifest = &v4alpha.DataSource{}
			}
			if err := m.SamplingRuleManifest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipXray(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthXray
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthXray
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipXray(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowXray
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
					return 0, ErrIntOverflowXray
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
					return 0, ErrIntOverflowXray
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
				return 0, ErrInvalidLengthXray
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupXray
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthXray
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthXray        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowXray          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupXray = fmt.Errorf("proto: unexpected end of group")
)