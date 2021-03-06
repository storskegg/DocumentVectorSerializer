// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vectr.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	vectr.proto

It has these top-level messages:
	DocumentVector
*/
package main

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

type DocumentVector struct {
	Vector []float32 `protobuf:"fixed32,1,rep,packed,name=vector" json:"vector,omitempty"`
}

func (m *DocumentVector) Reset()                    { *m = DocumentVector{} }
func (m *DocumentVector) String() string            { return proto.CompactTextString(m) }
func (*DocumentVector) ProtoMessage()               {}
func (*DocumentVector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DocumentVector) GetVector() []float32 {
	if m != nil {
		return m.Vector
	}
	return nil
}

func init() {
	proto.RegisterType((*DocumentVector)(nil), "main.DocumentVector")
}

func init() { proto.RegisterFile("vectr.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 82 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x4b, 0x4d, 0x2e,
	0x29, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0xd2, 0xe0,
	0xe2, 0x73, 0xc9, 0x4f, 0x2e, 0xcd, 0x4d, 0xcd, 0x2b, 0x09, 0x4b, 0x4d, 0x2e, 0xc9, 0x2f, 0x12,
	0x12, 0xe3, 0x62, 0x2b, 0x03, 0xb3, 0x24, 0x18, 0x15, 0x98, 0x35, 0x98, 0x82, 0xa0, 0xbc, 0x24,
	0x36, 0xb0, 0x36, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x02, 0x03, 0xf9, 0x45, 0x00,
	0x00, 0x00,
}
