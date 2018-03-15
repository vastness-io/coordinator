// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commit_author.proto

/*
Package vcs is a generated protocol buffer package.

It is generated from these files:
	commit_author.proto
	push_event.proto
	repository.proto
	user.proto
	vcs_event.proto

It has these top-level messages:
	CommitAuthor
	PushCommit
	VcsPushEvent
	Repository
	User
*/
package vcs

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

type CommitAuthor struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
	Date     string `protobuf:"bytes,4,opt,name=date" json:"date,omitempty"`
}

func (m *CommitAuthor) Reset()                    { *m = CommitAuthor{} }
func (m *CommitAuthor) String() string            { return proto.CompactTextString(m) }
func (*CommitAuthor) ProtoMessage()               {}
func (*CommitAuthor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CommitAuthor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CommitAuthor) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CommitAuthor) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CommitAuthor) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func init() {
	proto.RegisterType((*CommitAuthor)(nil), "vcs.CommitAuthor")
}

func init() { proto.RegisterFile("commit_author.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0xcf, 0xcd,
	0xcd, 0x2c, 0x89, 0x4f, 0x2c, 0x2d, 0xc9, 0xc8, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2e, 0x4b, 0x2e, 0x56, 0xca, 0xe0, 0xe2, 0x71, 0x06, 0xcb, 0x39, 0x82, 0xa5, 0x84, 0x84,
	0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21,
	0x11, 0x2e, 0xd6, 0xd4, 0xdc, 0xc4, 0xcc, 0x1c, 0x09, 0x26, 0xb0, 0x20, 0x84, 0x23, 0x24, 0xc5,
	0xc5, 0x51, 0x5a, 0x9c, 0x5a, 0x04, 0x56, 0xcd, 0x0c, 0x96, 0x80, 0xf3, 0x41, 0xa6, 0xa4, 0x24,
	0x96, 0xa4, 0x4a, 0xb0, 0x40, 0x4c, 0x01, 0xb1, 0x93, 0xd8, 0xc0, 0xb6, 0x1a, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x17, 0x57, 0x3c, 0xd9, 0x8c, 0x00, 0x00, 0x00,
}
