// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	RoleList
	TargetName
	Target
	TargetWithRole
	TargetWithRoleList
	TargetByNameAction
	Signature
	PublicKey
	DelegationRole
	TargetSigned
	TargetSignedList
	BasicResponse
*/
package api

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// RoleList message holds a list of TUF role names
type RoleList struct {
	Roles []string `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
}

func (m *RoleList) Reset()                    { *m = RoleList{} }
func (m *RoleList) String() string            { return proto.CompactTextString(m) }
func (*RoleList) ProtoMessage()               {}
func (*RoleList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RoleList) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

type TargetName struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *TargetName) Reset()                    { *m = TargetName{} }
func (m *TargetName) String() string            { return proto.CompactTextString(m) }
func (*TargetName) ProtoMessage()               {}
func (*TargetName) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TargetName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Target message describes a TUF target
type Target struct {
	Gun    string            `protobuf:"bytes,1,opt,name=gun" json:"gun,omitempty"`
	Name   string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Length int64             `protobuf:"varint,3,opt,name=length" json:"length,omitempty"`
	Hashes map[string][]byte `protobuf:"bytes,4,rep,name=hashes" json:"hashes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Custom []byte            `protobuf:"bytes,5,opt,name=custom,proto3" json:"custom,omitempty"`
}

func (m *Target) Reset()                    { *m = Target{} }
func (m *Target) String() string            { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()               {}
func (*Target) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Target) GetGun() string {
	if m != nil {
		return m.Gun
	}
	return ""
}

func (m *Target) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Target) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Target) GetHashes() map[string][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

func (m *Target) GetCustom() []byte {
	if m != nil {
		return m.Custom
	}
	return nil
}

// TargetWithRole represents a Target that exists in a particular role
type TargetWithRole struct {
	Target *Target `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
	Role   string  `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
}

func (m *TargetWithRole) Reset()                    { *m = TargetWithRole{} }
func (m *TargetWithRole) String() string            { return proto.CompactTextString(m) }
func (*TargetWithRole) ProtoMessage()               {}
func (*TargetWithRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TargetWithRole) GetTarget() *Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *TargetWithRole) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type TargetWithRoleList struct {
	Targets []*TargetWithRole `protobuf:"bytes,1,rep,name=targets" json:"targets,omitempty"`
}

func (m *TargetWithRoleList) Reset()                    { *m = TargetWithRoleList{} }
func (m *TargetWithRoleList) String() string            { return proto.CompactTextString(m) }
func (*TargetWithRoleList) ProtoMessage()               {}
func (*TargetWithRoleList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TargetWithRoleList) GetTargets() []*TargetWithRole {
	if m != nil {
		return m.Targets
	}
	return nil
}

type TargetByNameAction struct {
	Name  string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Roles *RoleList `protobuf:"bytes,2,opt,name=roles" json:"roles,omitempty"`
}

func (m *TargetByNameAction) Reset()                    { *m = TargetByNameAction{} }
func (m *TargetByNameAction) String() string            { return proto.CompactTextString(m) }
func (*TargetByNameAction) ProtoMessage()               {}
func (*TargetByNameAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TargetByNameAction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TargetByNameAction) GetRoles() *RoleList {
	if m != nil {
		return m.Roles
	}
	return nil
}

type Signature struct {
	KeyID     string `protobuf:"bytes,1,opt,name=keyID" json:"keyID,omitempty"`
	Method    string `protobuf:"bytes,2,opt,name=method" json:"method,omitempty"`
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	IsValid   bool   `protobuf:"varint,4,opt,name=isValid" json:"isValid,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Signature) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *Signature) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Signature) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

type PublicKey struct {
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type DelegationRole struct {
	Keys      map[string]*PublicKey `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Name      string                `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Threshold int32                 `protobuf:"varint,3,opt,name=threshold" json:"threshold,omitempty"`
	Paths     []string              `protobuf:"bytes,4,rep,name=paths" json:"paths,omitempty"`
}

func (m *DelegationRole) Reset()                    { *m = DelegationRole{} }
func (m *DelegationRole) String() string            { return proto.CompactTextString(m) }
func (*DelegationRole) ProtoMessage()               {}
func (*DelegationRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DelegationRole) GetKeys() map[string]*PublicKey {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *DelegationRole) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DelegationRole) GetThreshold() int32 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *DelegationRole) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

type TargetSigned struct {
	Role       *DelegationRole `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	Target     *Target         `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	Signatures []*Signature    `protobuf:"bytes,3,rep,name=signatures" json:"signatures,omitempty"`
}

func (m *TargetSigned) Reset()                    { *m = TargetSigned{} }
func (m *TargetSigned) String() string            { return proto.CompactTextString(m) }
func (*TargetSigned) ProtoMessage()               {}
func (*TargetSigned) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *TargetSigned) GetRole() *DelegationRole {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *TargetSigned) GetTarget() *Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *TargetSigned) GetSignatures() []*Signature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type TargetSignedList struct {
	Targets []*TargetSigned `protobuf:"bytes,1,rep,name=targets" json:"targets,omitempty"`
}

func (m *TargetSignedList) Reset()                    { *m = TargetSignedList{} }
func (m *TargetSignedList) String() string            { return proto.CompactTextString(m) }
func (*TargetSignedList) ProtoMessage()               {}
func (*TargetSignedList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *TargetSignedList) GetTargets() []*TargetSigned {
	if m != nil {
		return m.Targets
	}
	return nil
}

// BasicResponse describes a response with a true/false success indicator,
// and if false, an error type and message. See the errors.go file in this
// package for the possible errors and a translation function between the
// BasicResponse object and a concrete error type.
type BasicResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *BasicResponse) Reset()                    { *m = BasicResponse{} }
func (m *BasicResponse) String() string            { return proto.CompactTextString(m) }
func (*BasicResponse) ProtoMessage()               {}
func (*BasicResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *BasicResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *BasicResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RoleList)(nil), "api.RoleList")
	proto.RegisterType((*TargetName)(nil), "api.TargetName")
	proto.RegisterType((*Target)(nil), "api.Target")
	proto.RegisterType((*TargetWithRole)(nil), "api.TargetWithRole")
	proto.RegisterType((*TargetWithRoleList)(nil), "api.TargetWithRoleList")
	proto.RegisterType((*TargetByNameAction)(nil), "api.TargetByNameAction")
	proto.RegisterType((*Signature)(nil), "api.Signature")
	proto.RegisterType((*PublicKey)(nil), "api.PublicKey")
	proto.RegisterType((*DelegationRole)(nil), "api.DelegationRole")
	proto.RegisterType((*TargetSigned)(nil), "api.TargetSigned")
	proto.RegisterType((*TargetSignedList)(nil), "api.TargetSignedList")
	proto.RegisterType((*BasicResponse)(nil), "api.BasicResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Notary service

type NotaryClient interface {
	// AddTarget adds a target to the TUF repository and re-signs.
	AddTarget(ctx context.Context, in *Target, opts ...grpc.CallOption) (*BasicResponse, error)
	// RemoveTarget deletes a target from the TUF repository and re-signs. It only
	// uses the `name` field from the Target object, ignoring all other fields
	RemoveTarget(ctx context.Context, in *Target, opts ...grpc.CallOption) (*BasicResponse, error)
	// ListTargets list the targets for the specified roles in the TUF repository
	ListTargets(ctx context.Context, in *RoleList, opts ...grpc.CallOption) (*TargetWithRoleList, error)
	// GetTargetByName returns a target by the given name.
	GetTargetByName(ctx context.Context, in *TargetByNameAction, opts ...grpc.CallOption) (*TargetWithRole, error)
	// GetAllTargetMetadataByName
	GetAllTargetMetadataByName(ctx context.Context, in *TargetName, opts ...grpc.CallOption) (*TargetSignedList, error)
}

type notaryClient struct {
	cc *grpc.ClientConn
}

func NewNotaryClient(cc *grpc.ClientConn) NotaryClient {
	return &notaryClient{cc}
}

func (c *notaryClient) AddTarget(ctx context.Context, in *Target, opts ...grpc.CallOption) (*BasicResponse, error) {
	out := new(BasicResponse)
	err := grpc.Invoke(ctx, "/api.Notary/AddTarget", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notaryClient) RemoveTarget(ctx context.Context, in *Target, opts ...grpc.CallOption) (*BasicResponse, error) {
	out := new(BasicResponse)
	err := grpc.Invoke(ctx, "/api.Notary/RemoveTarget", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notaryClient) ListTargets(ctx context.Context, in *RoleList, opts ...grpc.CallOption) (*TargetWithRoleList, error) {
	out := new(TargetWithRoleList)
	err := grpc.Invoke(ctx, "/api.Notary/ListTargets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notaryClient) GetTargetByName(ctx context.Context, in *TargetByNameAction, opts ...grpc.CallOption) (*TargetWithRole, error) {
	out := new(TargetWithRole)
	err := grpc.Invoke(ctx, "/api.Notary/GetTargetByName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notaryClient) GetAllTargetMetadataByName(ctx context.Context, in *TargetName, opts ...grpc.CallOption) (*TargetSignedList, error) {
	out := new(TargetSignedList)
	err := grpc.Invoke(ctx, "/api.Notary/GetAllTargetMetadataByName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Notary service

type NotaryServer interface {
	// AddTarget adds a target to the TUF repository and re-signs.
	AddTarget(context.Context, *Target) (*BasicResponse, error)
	// RemoveTarget deletes a target from the TUF repository and re-signs. It only
	// uses the `name` field from the Target object, ignoring all other fields
	RemoveTarget(context.Context, *Target) (*BasicResponse, error)
	// ListTargets list the targets for the specified roles in the TUF repository
	ListTargets(context.Context, *RoleList) (*TargetWithRoleList, error)
	// GetTargetByName returns a target by the given name.
	GetTargetByName(context.Context, *TargetByNameAction) (*TargetWithRole, error)
	// GetAllTargetMetadataByName
	GetAllTargetMetadataByName(context.Context, *TargetName) (*TargetSignedList, error)
}

func RegisterNotaryServer(s *grpc.Server, srv NotaryServer) {
	s.RegisterService(&_Notary_serviceDesc, srv)
}

func _Notary_AddTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Target)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotaryServer).AddTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Notary/AddTarget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotaryServer).AddTarget(ctx, req.(*Target))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notary_RemoveTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Target)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotaryServer).RemoveTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Notary/RemoveTarget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotaryServer).RemoveTarget(ctx, req.(*Target))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notary_ListTargets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotaryServer).ListTargets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Notary/ListTargets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotaryServer).ListTargets(ctx, req.(*RoleList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notary_GetTargetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TargetByNameAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotaryServer).GetTargetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Notary/GetTargetByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotaryServer).GetTargetByName(ctx, req.(*TargetByNameAction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notary_GetAllTargetMetadataByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TargetName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotaryServer).GetAllTargetMetadataByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Notary/GetAllTargetMetadataByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotaryServer).GetAllTargetMetadataByName(ctx, req.(*TargetName))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notary_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Notary",
	HandlerType: (*NotaryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTarget",
			Handler:    _Notary_AddTarget_Handler,
		},
		{
			MethodName: "RemoveTarget",
			Handler:    _Notary_RemoveTarget_Handler,
		},
		{
			MethodName: "ListTargets",
			Handler:    _Notary_ListTargets_Handler,
		},
		{
			MethodName: "GetTargetByName",
			Handler:    _Notary_GetTargetByName_Handler,
		},
		{
			MethodName: "GetAllTargetMetadataByName",
			Handler:    _Notary_GetAllTargetMetadataByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 662 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6f, 0xd3, 0x3c,
	0x14, 0x5e, 0x9a, 0xae, 0x5b, 0x4e, 0xba, 0x8f, 0xd7, 0x7b, 0x61, 0x51, 0x34, 0xa4, 0x28, 0x43,
	0xa2, 0x12, 0xa2, 0x68, 0xe5, 0x82, 0x8f, 0x1b, 0xd4, 0x6d, 0x30, 0xa6, 0xb1, 0x09, 0x19, 0x04,
	0xd7, 0x5e, 0x73, 0xd4, 0x44, 0x4d, 0x93, 0x12, 0xbb, 0x93, 0xf2, 0x1f, 0xf8, 0x2f, 0xfc, 0x0f,
	0xae, 0xf9, 0x41, 0xc8, 0x76, 0xdc, 0xba, 0x50, 0x24, 0xee, 0x72, 0x7c, 0x3e, 0xfc, 0x3c, 0xcf,
	0x79, 0x1c, 0xf0, 0xd8, 0x2c, 0xeb, 0xcf, 0xaa, 0x52, 0x94, 0xc4, 0x65, 0xb3, 0x2c, 0x8e, 0x60,
	0x9b, 0x96, 0x39, 0xbe, 0xcf, 0xb8, 0x20, 0xff, 0xc3, 0x66, 0x55, 0xe6, 0xc8, 0x03, 0x27, 0x72,
	0x7b, 0x1e, 0xd5, 0x41, 0x1c, 0x01, 0x7c, 0x62, 0xd5, 0x18, 0xc5, 0x0d, 0x9b, 0x22, 0x21, 0xd0,
	0x2e, 0xd8, 0x14, 0x03, 0x27, 0x72, 0x7a, 0x1e, 0x55, 0xdf, 0xf1, 0x0f, 0x07, 0x3a, 0xba, 0x84,
	0xec, 0x83, 0x3b, 0x9e, 0x17, 0x4d, 0x56, 0x7e, 0x2e, 0x1a, 0x5a, 0xcb, 0x06, 0x72, 0x1f, 0x3a,
	0x39, 0x16, 0x63, 0x91, 0x06, 0x6e, 0xe4, 0xf4, 0x5c, 0xda, 0x44, 0xe4, 0x29, 0x74, 0x52, 0xc6,
	0x53, 0xe4, 0x41, 0x3b, 0x72, 0x7b, 0xfe, 0xe0, 0xb0, 0x2f, 0xd1, 0xea, 0xd1, 0xfd, 0x77, 0x2a,
	0xf3, 0xa6, 0x10, 0x55, 0x4d, 0x9b, 0x32, 0x39, 0x68, 0x34, 0xe7, 0xa2, 0x9c, 0x06, 0x9b, 0x91,
	0xd3, 0xeb, 0xd2, 0x26, 0x0a, 0x5f, 0x82, 0x6f, 0x95, 0x4b, 0x54, 0x13, 0xac, 0x0d, 0xaa, 0x09,
	0xd6, 0x92, 0xea, 0x1d, 0xcb, 0xe7, 0x1a, 0x56, 0x97, 0xea, 0xe0, 0x55, 0xeb, 0x85, 0x13, 0x5f,
	0xc2, 0xae, 0xbe, 0xf0, 0x4b, 0x26, 0x52, 0x29, 0x0d, 0x39, 0x86, 0x8e, 0x50, 0x27, 0x6a, 0x80,
	0x3f, 0xf0, 0x2d, 0x54, 0xb4, 0x49, 0x49, 0x9a, 0x52, 0x2e, 0x43, 0x53, 0x7e, 0xc7, 0x67, 0x40,
	0x56, 0x47, 0x29, 0x95, 0x9f, 0xc0, 0x96, 0xee, 0xd1, 0x3a, 0xfb, 0x83, 0x03, 0x6b, 0x9e, 0xa9,
	0xa4, 0xa6, 0x26, 0xbe, 0x36, 0x43, 0x4e, 0x6b, 0xb9, 0x80, 0xe1, 0x48, 0x64, 0x65, 0xb1, 0x6e,
	0x0d, 0xe4, 0xd8, 0xac, 0xaf, 0xa5, 0x60, 0xee, 0xa8, 0xb1, 0xe6, 0x5a, 0xb3, 0xcd, 0xaf, 0xe0,
	0x7d, 0xcc, 0xc6, 0x05, 0x13, 0xf3, 0x0a, 0xa5, 0x0a, 0x13, 0xac, 0x2f, 0xcf, 0x9b, 0x31, 0x3a,
	0x90, 0xa2, 0x4e, 0x51, 0xa4, 0x65, 0xd2, 0x90, 0x69, 0x22, 0x72, 0x04, 0x1e, 0x37, 0xad, 0x6a,
	0x71, 0x5d, 0xba, 0x3c, 0x20, 0x01, 0x6c, 0x65, 0xfc, 0x33, 0xcb, 0xb3, 0x24, 0x68, 0x47, 0x4e,
	0x6f, 0x9b, 0x9a, 0x30, 0xf6, 0xc1, 0xfb, 0x30, 0xbf, 0xcd, 0xb3, 0xd1, 0x15, 0xd6, 0xf1, 0x4f,
	0x07, 0x76, 0xcf, 0x31, 0xc7, 0x31, 0x93, 0x3c, 0x94, 0xbe, 0x27, 0xd0, 0x9e, 0x60, 0x6d, 0xd4,
	0x78, 0xa0, 0x60, 0xaf, 0x96, 0xf4, 0xaf, 0xb0, 0x6e, 0x36, 0xaf, 0x4a, 0xd7, 0x9a, 0xea, 0x08,
	0x3c, 0x91, 0x56, 0xc8, 0xd3, 0x32, 0x4f, 0x14, 0xbc, 0x4d, 0xba, 0x3c, 0x90, 0x54, 0x67, 0x4c,
	0xa4, 0xda, 0x59, 0x1e, 0xd5, 0x41, 0x78, 0x01, 0xde, 0x62, 0xf4, 0x1a, 0x97, 0x3c, 0xb4, 0x5d,
	0xe2, 0x0f, 0x76, 0x15, 0xb4, 0x05, 0x17, 0xdb, 0x35, 0xdf, 0x1c, 0xe8, 0xea, 0x35, 0x49, 0x75,
	0x31, 0x21, 0x8f, 0x1a, 0x3f, 0x68, 0xcb, 0x1c, 0xac, 0x21, 0xa5, 0x4d, 0x62, 0xb9, 0xab, 0xf5,
	0x77, 0x77, 0xf5, 0x01, 0x16, 0x4a, 0xf3, 0xc0, 0x55, 0x42, 0x69, 0x34, 0x8b, 0x65, 0x52, 0xab,
	0x22, 0x7e, 0x0d, 0xfb, 0x36, 0x1a, 0xe5, 0xbb, 0xc7, 0xbf, 0xfb, 0xee, 0x3f, 0xeb, 0x26, 0x5d,
	0xb7, 0x74, 0xdd, 0x19, 0xec, 0x9c, 0x32, 0x9e, 0x8d, 0x28, 0xf2, 0x59, 0x59, 0x70, 0xb5, 0x5e,
	0x3e, 0x1f, 0x8d, 0x90, 0x73, 0x45, 0x69, 0x9b, 0x9a, 0x50, 0x66, 0xa6, 0xc8, 0x39, 0x1b, 0x9b,
	0x75, 0x98, 0x70, 0xf0, 0xbd, 0x05, 0x9d, 0x9b, 0x52, 0xb0, 0xaa, 0x26, 0x7d, 0xf0, 0x86, 0x49,
	0xd2, 0xfc, 0x24, 0x6c, 0x8a, 0x21, 0x51, 0xc1, 0xca, 0x65, 0xf1, 0x06, 0x39, 0x81, 0x2e, 0xc5,
	0x69, 0x79, 0x87, 0xff, 0xde, 0xf2, 0x1c, 0x7c, 0xc9, 0x53, 0xd7, 0x70, 0xb2, 0x6a, 0xff, 0xf0,
	0x70, 0xcd, 0x23, 0x93, 0x89, 0x78, 0x83, 0x0c, 0x61, 0xef, 0x02, 0x85, 0xfd, 0xc8, 0x88, 0x5d,
	0x6d, 0xbf, 0xbb, 0x70, 0xdd, 0x5b, 0x8d, 0x37, 0xc8, 0x5b, 0x08, 0x2f, 0x50, 0x0c, 0xf3, 0x5c,
	0x67, 0xae, 0x51, 0xb0, 0x84, 0x09, 0xd6, 0x4c, 0xdb, 0xb3, 0x9a, 0xe4, 0x41, 0x78, 0xef, 0x0f,
	0xe5, 0x35, 0x94, 0xdb, 0x8e, 0xfa, 0x33, 0x3f, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x81, 0xf5,
	0x43, 0xd6, 0xa6, 0x05, 0x00, 0x00,
}
