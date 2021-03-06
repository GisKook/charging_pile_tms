// Code generated by protoc-gen-go.
// source: manage.proto
// DO NOT EDIT!

/*
Package Report is a generated protocol buffer package.

It is generated from these files:
	manage.proto
	param.proto

It has these top-level messages:
	Command
	Param
*/
package Report

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Command_CommandType int32

const (
	Command_CMT_INVALID Command_CommandType = 0
	// das<->tms
	Command_CMT_REQ_LOGIN Command_CommandType = 1
	Command_CMT_REP_LOGIN Command_CommandType = 32769
	// das<->tms
	Command_CMT_REQ_SETTING Command_CommandType = 16
	Command_CMT_REP_SETTING Command_CommandType = 32784
	// das->tss
	Command_CMT_REQ_HEART Command_CommandType = 5
	// das<->tms
	Command_CMT_REQ_PRICE Command_CommandType = 3
	Command_CMT_REP_PRICE Command_CommandType = 32771
	// wechat <-> das
	Command_CMT_REQ_GET_GUN_STATUS Command_CommandType = 32774
	Command_CMT_REP_GET_GUN_STATUS Command_CommandType = 6
	// wechat <-> das
	Command_CMT_REQ_CHARGING Command_CommandType = 32775
	Command_CMT_REP_CHARGING Command_CommandType = 7
	// wechat <-> das
	Command_CMT_REQ_STOP_CHARGING Command_CommandType = 32782
	Command_CMT_REP_STOP_CHARGING Command_CommandType = 14
	// web<->das
	Command_CMT_REQ_NOTIFY_SET_PRICE Command_CommandType = 32783
	Command_CMT_REP_NOTIFY_SET_PRICE Command_CommandType = 15
	// das->wechat
	Command_CMT_REP_CHARGING_STARTED Command_CommandType = 8
	// das->wechat
	Command_CMT_REP_CHARGING_DATA_UPLOAD Command_CommandType = 9
	// das->wechat
	Command_CMT_REP_CHARGING_STOPPED Command_CommandType = 11
	// das <-> web
	Command_CMT_REQ_PIN Command_CommandType = 32780
	Command_CMT_REP_PIN Command_CommandType = 12
	// das -> ?
	Command_CMT_REP_OFFLINE_DATA Command_CommandType = 13
)

var Command_CommandType_name = map[int32]string{
	0:     "CMT_INVALID",
	1:     "CMT_REQ_LOGIN",
	32769: "CMT_REP_LOGIN",
	16:    "CMT_REQ_SETTING",
	32784: "CMT_REP_SETTING",
	5:     "CMT_REQ_HEART",
	3:     "CMT_REQ_PRICE",
	32771: "CMT_REP_PRICE",
	32774: "CMT_REQ_GET_GUN_STATUS",
	6:     "CMT_REP_GET_GUN_STATUS",
	32775: "CMT_REQ_CHARGING",
	7:     "CMT_REP_CHARGING",
	32782: "CMT_REQ_STOP_CHARGING",
	14:    "CMT_REP_STOP_CHARGING",
	32783: "CMT_REQ_NOTIFY_SET_PRICE",
	15:    "CMT_REP_NOTIFY_SET_PRICE",
	8:     "CMT_REP_CHARGING_STARTED",
	9:     "CMT_REP_CHARGING_DATA_UPLOAD",
	11:    "CMT_REP_CHARGING_STOPPED",
	32780: "CMT_REQ_PIN",
	12:    "CMT_REP_PIN",
	13:    "CMT_REP_OFFLINE_DATA",
}
var Command_CommandType_value = map[string]int32{
	"CMT_INVALID":                  0,
	"CMT_REQ_LOGIN":                1,
	"CMT_REP_LOGIN":                32769,
	"CMT_REQ_SETTING":              16,
	"CMT_REP_SETTING":              32784,
	"CMT_REQ_HEART":                5,
	"CMT_REQ_PRICE":                3,
	"CMT_REP_PRICE":                32771,
	"CMT_REQ_GET_GUN_STATUS":       32774,
	"CMT_REP_GET_GUN_STATUS":       6,
	"CMT_REQ_CHARGING":             32775,
	"CMT_REP_CHARGING":             7,
	"CMT_REQ_STOP_CHARGING":        32782,
	"CMT_REP_STOP_CHARGING":        14,
	"CMT_REQ_NOTIFY_SET_PRICE":     32783,
	"CMT_REP_NOTIFY_SET_PRICE":     15,
	"CMT_REP_CHARGING_STARTED":     8,
	"CMT_REP_CHARGING_DATA_UPLOAD": 9,
	"CMT_REP_CHARGING_STOPPED":     11,
	"CMT_REQ_PIN":                  32780,
	"CMT_REP_PIN":                  12,
	"CMT_REP_OFFLINE_DATA":         13,
}

func (x Command_CommandType) String() string {
	return proto.EnumName(Command_CommandType_name, int32(x))
}
func (Command_CommandType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Command struct {
	Type         Command_CommandType `protobuf:"varint,1,opt,name=type,enum=Report.Command_CommandType" json:"type,omitempty"`
	Uuid         string              `protobuf:"bytes,2,opt,name=uuid" json:"uuid,omitempty"`
	Tid          uint64              `protobuf:"varint,3,opt,name=tid" json:"tid,omitempty"`
	SerialNumber uint32              `protobuf:"varint,4,opt,name=serial_number" json:"serial_number,omitempty"`
	Paras        []*Param            `protobuf:"bytes,5,rep,name=paras" json:"paras,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Command) GetType() Command_CommandType {
	if m != nil {
		return m.Type
	}
	return Command_CMT_INVALID
}

func (m *Command) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Command) GetTid() uint64 {
	if m != nil {
		return m.Tid
	}
	return 0
}

func (m *Command) GetSerialNumber() uint32 {
	if m != nil {
		return m.SerialNumber
	}
	return 0
}

func (m *Command) GetParas() []*Param {
	if m != nil {
		return m.Paras
	}
	return nil
}

func init() {
	proto.RegisterType((*Command)(nil), "Report.Command")
	proto.RegisterEnum("Report.Command_CommandType", Command_CommandType_name, Command_CommandType_value)
}

func init() { proto.RegisterFile("manage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xcd, 0x6e, 0x9b, 0x40,
	0x10, 0x80, 0x0b, 0xc6, 0x4e, 0x33, 0x98, 0x78, 0x33, 0x89, 0x23, 0x9a, 0x58, 0x15, 0xca, 0x89,
	0x5e, 0x7c, 0x48, 0x9f, 0x00, 0x99, 0x35, 0x59, 0xc9, 0x5d, 0xb6, 0x78, 0x5d, 0xa9, 0x27, 0x44,
	0x64, 0x54, 0x59, 0x2a, 0x36, 0x22, 0xf6, 0x21, 0x37, 0xaa, 0x4a, 0xed, 0xa5, 0x6a, 0xfb, 0x16,
	0xbd, 0xf5, 0x19, 0x2b, 0xc0, 0x9b, 0xa2, 0x58, 0x39, 0x21, 0x7d, 0xdf, 0xfc, 0xee, 0x00, 0xfd,
	0x2c, 0x59, 0x27, 0x9f, 0xd2, 0x71, 0x5e, 0x6c, 0xb6, 0x1b, 0xec, 0x45, 0x69, 0xbe, 0x29, 0xb6,
	0x97, 0x66, 0x9e, 0x14, 0x49, 0xd6, 0xc0, 0xeb, 0xbf, 0x5d, 0x38, 0x9a, 0x6c, 0xb2, 0x2c, 0x59,
	0x2f, 0xf1, 0x0d, 0x18, 0xdb, 0x87, 0x3c, 0xb5, 0x35, 0x47, 0x73, 0x4f, 0x6e, 0xae, 0xc6, 0x4d,
	0xfc, 0x78, 0xaf, 0xd5, 0x57, 0x3e, 0xe4, 0x29, 0xf6, 0xc1, 0xd8, 0xed, 0x56, 0x4b, 0x5b, 0x77,
	0x34, 0xf7, 0x18, 0x4d, 0xe8, 0x6c, 0x57, 0x4b, 0xbb, 0xe3, 0x68, 0xae, 0x81, 0x43, 0xb0, 0xee,
	0xd3, 0x62, 0x95, 0x7c, 0x8e, 0xd7, 0xbb, 0xec, 0x2e, 0x2d, 0x6c, 0xc3, 0xd1, 0x5c, 0x0b, 0x47,
	0xd0, 0xad, 0xfa, 0xde, 0xdb, 0x5d, 0xa7, 0xe3, 0x9a, 0x37, 0x96, 0xaa, 0x2e, 0xaa, 0x61, 0xae,
	0xff, 0x18, 0x60, 0xb6, 0xeb, 0x0f, 0xc0, 0x9c, 0xbc, 0x93, 0x31, 0xe3, 0x1f, 0xbc, 0x19, 0xf3,
	0xc9, 0x0b, 0x3c, 0x05, 0xab, 0x02, 0x11, 0x7d, 0x1f, 0xcf, 0xc2, 0x80, 0x71, 0xa2, 0xe1, 0x99,
	0x42, 0x62, 0x8f, 0xbe, 0x94, 0x3a, 0x9e, 0xc1, 0x40, 0xc5, 0xcd, 0xa9, 0x94, 0x8c, 0x07, 0x84,
	0xe0, 0x50, 0x41, 0xf1, 0x08, 0x7f, 0x97, 0x7a, 0xbb, 0xe6, 0x2d, 0xf5, 0x22, 0x49, 0xba, 0x6d,
	0x24, 0x22, 0x36, 0xa1, 0xa4, 0xd3, 0x6e, 0xd3, 0xa0, 0xaf, 0xa5, 0x8e, 0x23, 0xb8, 0x50, 0x71,
	0x01, 0x95, 0x71, 0xb0, 0xe0, 0xf1, 0x5c, 0x7a, 0x72, 0x31, 0x27, 0xdf, 0x4a, 0x1d, 0x2f, 0x95,
	0x15, 0x4f, 0x6d, 0x0f, 0x2f, 0x80, 0xa8, 0xcc, 0xc9, 0xad, 0x17, 0x05, 0xd5, 0x30, 0xdf, 0x4b,
	0x1d, 0xcf, 0x15, 0x17, 0xff, 0xf9, 0x11, 0x5e, 0xc1, 0xf0, 0x71, 0x1d, 0x19, 0xb6, 0xd4, 0xcf,
	0x52, 0xc7, 0x57, 0x4a, 0x8a, 0x27, 0xf2, 0x04, 0x5f, 0x83, 0xad, 0xf2, 0x78, 0x28, 0xd9, 0xf4,
	0x63, 0xb5, 0xf8, 0x7e, 0xfe, 0x5f, 0xf5, 0xfc, 0xb6, 0x4a, 0x3d, 0xf0, 0x83, 0xb6, 0x55, 0x35,
	0xab, 0x05, 0x22, 0x49, 0x7d, 0xf2, 0x12, 0x1d, 0x18, 0x1d, 0x58, 0xdf, 0x93, 0x5e, 0xbc, 0x10,
	0xb3, 0xd0, 0xf3, 0xc9, 0xf1, 0x33, 0xf9, 0xa1, 0x10, 0xd4, 0x27, 0x26, 0x9e, 0x36, 0xb7, 0xad,
	0xdf, 0x98, 0x71, 0xf2, 0xa3, 0xd4, 0xd5, 0xb9, 0xeb, 0x37, 0x66, 0x9c, 0xf4, 0xd1, 0x86, 0x73,
	0x05, 0xc2, 0xe9, 0x74, 0xc6, 0x38, 0xad, 0x5b, 0x10, 0xeb, 0xae, 0x57, 0xff, 0xb7, 0x6f, 0xff,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x15, 0x02, 0x7f, 0x26, 0xdc, 0x02, 0x00, 0x00,
}
