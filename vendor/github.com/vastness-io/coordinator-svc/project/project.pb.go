// Code generated by protoc-gen-go. DO NOT EDIT.
// source: project.proto

package project

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GetProjectRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *GetProjectRequest) Reset()                    { *m = GetProjectRequest{} }
func (m *GetProjectRequest) String() string            { return proto.CompactTextString(m) }
func (*GetProjectRequest) ProtoMessage()               {}
func (*GetProjectRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *GetProjectRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetProjectRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Project struct {
	Name         string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type         string        `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Repositories []*Repository `protobuf:"bytes,3,rep,name=repositories" json:"repositories,omitempty"`
}

func (m *Project) Reset()                    { *m = Project{} }
func (m *Project) String() string            { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()               {}
func (*Project) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Project) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Project) GetRepositories() []*Repository {
	if m != nil {
		return m.Repositories
	}
	return nil
}

type GetProjectsResponse struct {
	Meta     *GetProjectsResponse_Meta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Projects []*Project                `protobuf:"bytes,2,rep,name=projects" json:"projects,omitempty"`
}

func (m *GetProjectsResponse) Reset()                    { *m = GetProjectsResponse{} }
func (m *GetProjectsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetProjectsResponse) ProtoMessage()               {}
func (*GetProjectsResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *GetProjectsResponse) GetMeta() *GetProjectsResponse_Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetProjectsResponse) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

type GetProjectsResponse_Meta struct {
	CurrentPage int32 `protobuf:"varint,1,opt,name=current_page,json=currentPage" json:"current_page,omitempty"`
	LastPage    int32 `protobuf:"varint,3,opt,name=last_page,json=lastPage" json:"last_page,omitempty"`
	PerPage     int32 `protobuf:"varint,4,opt,name=per_page,json=perPage" json:"per_page,omitempty"`
	TotalCount  int32 `protobuf:"varint,5,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
}

func (m *GetProjectsResponse_Meta) Reset()                    { *m = GetProjectsResponse_Meta{} }
func (m *GetProjectsResponse_Meta) String() string            { return proto.CompactTextString(m) }
func (*GetProjectsResponse_Meta) ProtoMessage()               {}
func (*GetProjectsResponse_Meta) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2, 0} }

func (m *GetProjectsResponse_Meta) GetCurrentPage() int32 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *GetProjectsResponse_Meta) GetLastPage() int32 {
	if m != nil {
		return m.LastPage
	}
	return 0
}

func (m *GetProjectsResponse_Meta) GetPerPage() int32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *GetProjectsResponse_Meta) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type GetProjectsRequest struct {
	StartPage int32 `protobuf:"varint,1,opt,name=start_page,json=startPage" json:"start_page,omitempty"`
	Limit     int32 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *GetProjectsRequest) Reset()                    { *m = GetProjectsRequest{} }
func (m *GetProjectsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetProjectsRequest) ProtoMessage()               {}
func (*GetProjectsRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *GetProjectsRequest) GetStartPage() int32 {
	if m != nil {
		return m.StartPage
	}
	return 0
}

func (m *GetProjectsRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func init() {
	proto.RegisterType((*GetProjectRequest)(nil), "project.GetProjectRequest")
	proto.RegisterType((*Project)(nil), "project.Project")
	proto.RegisterType((*GetProjectsResponse)(nil), "project.GetProjectsResponse")
	proto.RegisterType((*GetProjectsResponse_Meta)(nil), "project.GetProjectsResponse.Meta")
	proto.RegisterType((*GetProjectsRequest)(nil), "project.GetProjectsRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Projects service

type ProjectsClient interface {
	GetProjects(ctx context.Context, in *GetProjectsRequest, opts ...grpc.CallOption) (*GetProjectsResponse, error)
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error)
}

type projectsClient struct {
	cc *grpc.ClientConn
}

func NewProjectsClient(cc *grpc.ClientConn) ProjectsClient {
	return &projectsClient{cc}
}

func (c *projectsClient) GetProjects(ctx context.Context, in *GetProjectsRequest, opts ...grpc.CallOption) (*GetProjectsResponse, error) {
	out := new(GetProjectsResponse)
	err := grpc.Invoke(ctx, "/project.Projects/GetProjects", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := grpc.Invoke(ctx, "/project.Projects/GetProject", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Projects service

type ProjectsServer interface {
	GetProjects(context.Context, *GetProjectsRequest) (*GetProjectsResponse, error)
	GetProject(context.Context, *GetProjectRequest) (*Project, error)
}

func RegisterProjectsServer(s *grpc.Server, srv ProjectsServer) {
	s.RegisterService(&_Projects_serviceDesc, srv)
}

func _Projects_GetProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).GetProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.Projects/GetProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).GetProjects(ctx, req.(*GetProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.Projects/GetProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Projects_serviceDesc = grpc.ServiceDesc{
	ServiceName: "project.Projects",
	HandlerType: (*ProjectsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProjects",
			Handler:    _Projects_GetProjects_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _Projects_GetProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "project.proto",
}

func init() { proto.RegisterFile("project.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6e, 0xe2, 0x30,
	0x10, 0xc6, 0x15, 0x48, 0x96, 0x30, 0x61, 0x25, 0xd6, 0xec, 0x21, 0x1b, 0xb6, 0x2a, 0xe4, 0xc4,
	0xa1, 0xe2, 0x40, 0x55, 0x55, 0x6a, 0x8f, 0x3d, 0xb4, 0x3d, 0x54, 0x42, 0x7e, 0x01, 0xe4, 0xa2,
	0x11, 0x0a, 0x82, 0xd8, 0xb5, 0x87, 0x03, 0xd7, 0x3e, 0x41, 0xdf, 0xb8, 0x55, 0x26, 0x7f, 0x28,
	0x2a, 0xaa, 0x7a, 0xb3, 0xbf, 0xdf, 0x37, 0xe3, 0xcf, 0x63, 0xc3, 0x6f, 0x63, 0xf5, 0x1a, 0x97,
	0x34, 0x35, 0x56, 0x93, 0x16, 0x9d, 0x6a, 0x9b, 0xf4, 0x2d, 0x1a, 0xed, 0x32, 0xd2, 0x76, 0x5f,
	0xa2, 0xf4, 0x16, 0xfe, 0xdc, 0x23, 0xcd, 0x4b, 0x2e, 0xf1, 0x65, 0x87, 0x8e, 0x84, 0x00, 0x3f,
	0x57, 0x5b, 0x8c, 0xbd, 0x91, 0x37, 0xe9, 0x4a, 0x5e, 0x17, 0x1a, 0xed, 0x0d, 0xc6, 0xad, 0x52,
	0x2b, 0xd6, 0xe9, 0x1a, 0x3a, 0x55, 0xe5, 0x4f, 0x4b, 0xc4, 0x35, 0xf4, 0x9a, 0x0c, 0x19, 0xba,
	0xb8, 0x3d, 0x6a, 0x4f, 0xa2, 0xd9, 0x60, 0x5a, 0x07, 0x96, 0x4d, 0x40, 0x79, 0x64, 0x4c, 0xdf,
	0x3d, 0x18, 0x1c, 0x92, 0x3a, 0x89, 0xce, 0xe8, 0xdc, 0xa1, 0xb8, 0x02, 0x7f, 0x8b, 0xa4, 0xf8,
	0xe0, 0x68, 0x36, 0x6e, 0x1a, 0x9d, 0xf0, 0x4e, 0x9f, 0x90, 0x94, 0x64, 0xbb, 0xb8, 0x80, 0xb0,
	0x72, 0xba, 0xb8, 0xc5, 0x19, 0xfa, 0x4d, 0x69, 0x3d, 0x8d, 0xc6, 0x91, 0xbc, 0x7a, 0xe0, 0x17,
	0xc5, 0x62, 0x0c, 0xbd, 0xe5, 0xce, 0x5a, 0xcc, 0x69, 0x61, 0xd4, 0xaa, 0xbc, 0x6e, 0x20, 0xa3,
	0x4a, 0x9b, 0xab, 0x15, 0x8a, 0x21, 0x74, 0x37, 0xca, 0x55, 0xbc, 0xcd, 0x3c, 0x2c, 0x04, 0x86,
	0xff, 0x20, 0x34, 0x68, 0x4b, 0xe6, 0x33, 0xeb, 0x18, 0xb4, 0x8c, 0xce, 0x21, 0x22, 0x4d, 0x6a,
	0xb3, 0x58, 0xea, 0x5d, 0x4e, 0x71, 0xc0, 0x14, 0x58, 0xba, 0x2b, 0x94, 0xf4, 0x11, 0xc4, 0xd1,
	0xa5, 0xca, 0xb7, 0x3a, 0x03, 0x70, 0xa4, 0xec, 0x51, 0x9e, 0x2e, 0x2b, 0xdc, 0xf5, 0x2f, 0x04,
	0x9b, 0x6c, 0x9b, 0x11, 0x3f, 0x42, 0x20, 0xcb, 0xcd, 0xec, 0xcd, 0x83, 0xb0, 0x6e, 0x24, 0x1e,
	0x20, 0xfa, 0xd4, 0x57, 0x0c, 0x4f, 0x8f, 0x90, 0x4f, 0x4b, 0xfe, 0x7f, 0x37, 0x5f, 0x71, 0x03,
	0x70, 0x90, 0x45, 0x72, 0xc2, 0x5b, 0xf7, 0xf9, 0x32, 0xec, 0xe7, 0x5f, 0xfc, 0x1f, 0x2f, 0x3f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x87, 0xb2, 0x4c, 0xbb, 0x02, 0x00, 0x00,
}
