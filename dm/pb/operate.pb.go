// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: operate.proto

package pb

import (
	"fmt"

	proto "github.com/gogo/protobuf/proto"

	math "math"

	_ "google.golang.org/genproto/googleapis/api/annotations"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OperateType int32

const (
	OperateType_Unknown                 OperateType = 0
	OperateType_MigrateWorkerRelay      OperateType = 1
	OperateType_UpdateWorkerRelayConfig OperateType = 2
	OperateType_StartTask               OperateType = 3
	OperateType_UpdateMasterConfig      OperateType = 4
	OperateType_OperateTask             OperateType = 5
	OperateType_UpdateTask              OperateType = 6
	OperateType_QueryStatus             OperateType = 7
	OperateType_QueryError              OperateType = 8
	OperateType_ShowDDLLocks            OperateType = 9
	OperateType_UnlockDDLLock           OperateType = 10
	OperateType_BreakWorkerDDLLock      OperateType = 11
	OperateType_SwitchWorkerRelayMaster OperateType = 12
	OperateType_OperateWorkerRelay      OperateType = 13
	OperateType_RefreshWorkerTasks      OperateType = 14
	OperateType_HandleSQLs              OperateType = 15
	OperateType_PurgeWorkerRelay        OperateType = 16
	OperateType_CheckTask               OperateType = 17
)

var OperateType_name = map[int32]string{
	0:  "Unknown",
	1:  "MigrateWorkerRelay",
	2:  "UpdateWorkerRelayConfig",
	3:  "StartTask",
	4:  "UpdateMasterConfig",
	5:  "OperateTask",
	6:  "UpdateTask",
	7:  "QueryStatus",
	8:  "QueryError",
	9:  "ShowDDLLocks",
	10: "UnlockDDLLock",
	11: "BreakWorkerDDLLock",
	12: "SwitchWorkerRelayMaster",
	13: "OperateWorkerRelay",
	14: "RefreshWorkerTasks",
	15: "HandleSQLs",
	16: "PurgeWorkerRelay",
	17: "CheckTask",
}
var OperateType_value = map[string]int32{
	"Unknown":                 0,
	"MigrateWorkerRelay":      1,
	"UpdateWorkerRelayConfig": 2,
	"StartTask":               3,
	"UpdateMasterConfig":      4,
	"OperateTask":             5,
	"UpdateTask":              6,
	"QueryStatus":             7,
	"QueryError":              8,
	"ShowDDLLocks":            9,
	"UnlockDDLLock":           10,
	"BreakWorkerDDLLock":      11,
	"SwitchWorkerRelayMaster": 12,
	"OperateWorkerRelay":      13,
	"RefreshWorkerTasks":      14,
	"HandleSQLs":              15,
	"PurgeWorkerRelay":        16,
	"CheckTask":               17,
}

func (x OperateType) String() string {
	return proto.EnumName(OperateType_name, int32(x))
}
func (OperateType) EnumDescriptor() ([]byte, []int) { return fileDescriptorOperate, []int{0} }

type Operate struct {
	Tp       OperateType `protobuf:"varint,1,opt,name=tp,proto3,enum=pb.OperateType" json:"tp,omitempty"`
	Request  []byte      `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Response []byte      `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	Err      string      `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
}

func (m *Operate) Reset()                    { *m = Operate{} }
func (m *Operate) String() string            { return proto.CompactTextString(m) }
func (*Operate) ProtoMessage()               {}
func (*Operate) Descriptor() ([]byte, []int) { return fileDescriptorOperate, []int{0} }

func (m *Operate) GetTp() OperateType {
	if m != nil {
		return m.Tp
	}
	return OperateType_Unknown
}

func (m *Operate) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Operate) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Operate) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*Operate)(nil), "pb.Operate")
	proto.RegisterEnum("pb.OperateType", OperateType_name, OperateType_value)
}
func (m *Operate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Operate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Tp != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintOperate(dAtA, i, uint64(m.Tp))
	}
	if len(m.Request) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOperate(dAtA, i, uint64(len(m.Request)))
		i += copy(dAtA[i:], m.Request)
	}
	if len(m.Response) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintOperate(dAtA, i, uint64(len(m.Response)))
		i += copy(dAtA[i:], m.Response)
	}
	if len(m.Err) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintOperate(dAtA, i, uint64(len(m.Err)))
		i += copy(dAtA[i:], m.Err)
	}
	return i, nil
}

func encodeVarintOperate(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Operate) Size() (n int) {
	var l int
	_ = l
	if m.Tp != 0 {
		n += 1 + sovOperate(uint64(m.Tp))
	}
	l = len(m.Request)
	if l > 0 {
		n += 1 + l + sovOperate(uint64(l))
	}
	l = len(m.Response)
	if l > 0 {
		n += 1 + l + sovOperate(uint64(l))
	}
	l = len(m.Err)
	if l > 0 {
		n += 1 + l + sovOperate(uint64(l))
	}
	return n
}

func sovOperate(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozOperate(x uint64) (n int) {
	return sovOperate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Operate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperate
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
			return fmt.Errorf("proto: Operate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Operate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tp", wireType)
			}
			m.Tp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tp |= (OperateType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperate
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
				return ErrInvalidLengthOperate
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Request = append(m.Request[:0], dAtA[iNdEx:postIndex]...)
			if m.Request == nil {
				m.Request = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperate
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
				return ErrInvalidLengthOperate
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Response = append(m.Response[:0], dAtA[iNdEx:postIndex]...)
			if m.Response == nil {
				m.Response = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Err", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperate
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
				return ErrInvalidLengthOperate
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Err = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOperate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOperate
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
func skipOperate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOperate
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
					return 0, ErrIntOverflowOperate
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
					return 0, ErrIntOverflowOperate
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
				return 0, ErrInvalidLengthOperate
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOperate
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
				next, err := skipOperate(dAtA[start:])
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
	ErrInvalidLengthOperate = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOperate   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("operate.proto", fileDescriptorOperate) }

var fileDescriptorOperate = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xbb, 0x4e, 0x69, 0x9a, 0xc9, 0xbf, 0xe9, 0x0a, 0x81, 0x55, 0x50, 0x88, 0x38, 0x45,
	0x1c, 0x52, 0x09, 0xde, 0xa0, 0x2d, 0x12, 0x87, 0x54, 0x50, 0xbb, 0x11, 0xe7, 0x4d, 0x3a, 0x75,
	0x2c, 0x5b, 0xbb, 0xcb, 0xec, 0x5a, 0x51, 0xde, 0x82, 0xc7, 0xe2, 0xc8, 0x23, 0xa0, 0xf0, 0x10,
	0x5c, 0xd1, 0xda, 0x2e, 0xa4, 0x37, 0xcf, 0x37, 0xbf, 0x6f, 0xf7, 0xfb, 0xac, 0x85, 0xa1, 0xb1,
	0xc4, 0xca, 0xd3, 0xdc, 0xb2, 0xf1, 0x46, 0x46, 0x76, 0x75, 0xfe, 0x3a, 0x33, 0x26, 0x2b, 0xe9,
	0x42, 0xd9, 0xfc, 0x42, 0x69, 0x6d, 0xbc, 0xf2, 0xb9, 0xd1, 0xae, 0x21, 0xde, 0x32, 0x74, 0x3f,
	0x37, 0x16, 0xf9, 0x06, 0x22, 0x6f, 0x63, 0x31, 0x15, 0xb3, 0xd1, 0xfb, 0xf1, 0xdc, 0xae, 0xe6,
	0xed, 0xe2, 0x6e, 0x67, 0x29, 0x89, 0xbc, 0x95, 0x31, 0x74, 0x99, 0xbe, 0x55, 0xe4, 0x7c, 0x1c,
	0x4d, 0xc5, 0x6c, 0x90, 0x3c, 0x8e, 0xf2, 0x1c, 0x4e, 0x99, 0x9c, 0x35, 0xda, 0x51, 0xdc, 0xa9,
	0x57, 0xff, 0x66, 0x89, 0xd0, 0x21, 0xe6, 0xf8, 0x78, 0x2a, 0x66, 0xbd, 0x24, 0x7c, 0xbe, 0xfb,
	0x13, 0x41, 0xff, 0xe0, 0x6c, 0xd9, 0x87, 0xee, 0x52, 0x17, 0xda, 0x6c, 0x35, 0x1e, 0xc9, 0x17,
	0x20, 0x6f, 0xf2, 0x2c, 0xec, 0xbe, 0x1a, 0x2e, 0x88, 0x13, 0x2a, 0xd5, 0x0e, 0x85, 0x7c, 0x05,
	0x2f, 0x97, 0xf6, 0xfe, 0xa9, 0x7c, 0x65, 0xf4, 0x43, 0x9e, 0x61, 0x24, 0x87, 0xd0, 0x4b, 0xbd,
	0x62, 0x7f, 0xa7, 0x5c, 0x81, 0x9d, 0x70, 0x46, 0xc3, 0xde, 0x28, 0xe7, 0x89, 0x5b, 0xec, 0x58,
	0x8e, 0xff, 0xdf, 0x1b, 0xc0, 0x67, 0x72, 0x04, 0xd0, 0x80, 0xf5, 0x7c, 0x12, 0x80, 0xdb, 0x8a,
	0x78, 0x97, 0x7a, 0xe5, 0x2b, 0x87, 0xdd, 0x00, 0xd4, 0xc2, 0x47, 0x66, 0xc3, 0x78, 0x2a, 0x11,
	0x06, 0xe9, 0xc6, 0x6c, 0xaf, 0xaf, 0x17, 0x0b, 0xb3, 0x2e, 0x1c, 0xf6, 0xe4, 0x19, 0x0c, 0x97,
	0xba, 0x34, 0xeb, 0xa2, 0xd5, 0x10, 0xc2, 0xf5, 0x97, 0x4c, 0xaa, 0x68, 0x92, 0x3e, 0xea, 0xfd,
	0x50, 0x21, 0xdd, 0xe6, 0x7e, 0xbd, 0x39, 0xa8, 0xd0, 0x24, 0xc4, 0x41, 0x30, 0xb5, 0xd9, 0x0e,
	0x7b, 0x0f, 0x83, 0x9e, 0xd0, 0x03, 0x93, 0x6b, 0x5d, 0x21, 0xa9, 0xc3, 0x51, 0x48, 0xf6, 0x49,
	0xe9, 0xfb, 0x92, 0xd2, 0xdb, 0x85, 0xc3, 0xb1, 0x7c, 0x0e, 0xf8, 0xa5, 0xe2, 0xec, 0x89, 0x1b,
	0xc3, 0x8f, 0xb9, 0xda, 0xd0, 0xba, 0xa8, 0xfb, 0x9d, 0x5d, 0xe2, 0x8f, 0xfd, 0x44, 0xfc, 0xdc,
	0x4f, 0xc4, 0xaf, 0xfd, 0x44, 0x7c, 0xff, 0x3d, 0x39, 0x5a, 0x9d, 0xd4, 0xcf, 0xe0, 0xc3, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xe1, 0x3b, 0x13, 0x39, 0x02, 0x00, 0x00,
}
