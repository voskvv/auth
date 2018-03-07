// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth_types.proto

package authProto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type StoredToken struct {
	TokenId       *UUID  `protobuf:"bytes,1,opt,name=token_id,json=tokenId" json:"token_id,omitempty"`
	UserAgent     string `protobuf:"bytes,2,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
	Platform      string `protobuf:"bytes,3,opt,name=platform" json:"platform,omitempty"`
	Fingerprint   string `protobuf:"bytes,4,opt,name=fingerprint" json:"fingerprint,omitempty"`
	UserId        *UUID  `protobuf:"bytes,5,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserRole      string `protobuf:"bytes,6,opt,name=user_role,json=userRole" json:"user_role,omitempty"`
	UserNamespace string `protobuf:"bytes,7,opt,name=user_namespace,json=userNamespace" json:"user_namespace,omitempty"`
	UserVolume    string `protobuf:"bytes,8,opt,name=user_volume,json=userVolume" json:"user_volume,omitempty"`
	RwAccess      bool   `protobuf:"varint,9,opt,name=rw_access,json=rwAccess" json:"rw_access,omitempty"`
	// @inject_tag: binding:"ip"
	UserIp      string                     `protobuf:"bytes,10,opt,name=user_ip,json=userIp" json:"user_ip,omitempty" binding:"ip"`
	PartTokenId *UUID                      `protobuf:"bytes,11,opt,name=part_token_id,json=partTokenId" json:"part_token_id,omitempty"`
	CreatedAt   *google_protobuf.Timestamp `protobuf:"bytes,12,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	LifeTime    *google_protobuf1.Duration `protobuf:"bytes,13,opt,name=life_time,json=lifeTime" json:"life_time,omitempty"`
}

func (m *StoredToken) Reset()                    { *m = StoredToken{} }
func (m *StoredToken) String() string            { return proto.CompactTextString(m) }
func (*StoredToken) ProtoMessage()               {}
func (*StoredToken) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *StoredToken) GetTokenId() *UUID {
	if m != nil {
		return m.TokenId
	}
	return nil
}

func (m *StoredToken) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *StoredToken) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *StoredToken) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *StoredToken) GetUserId() *UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *StoredToken) GetUserRole() string {
	if m != nil {
		return m.UserRole
	}
	return ""
}

func (m *StoredToken) GetUserNamespace() string {
	if m != nil {
		return m.UserNamespace
	}
	return ""
}

func (m *StoredToken) GetUserVolume() string {
	if m != nil {
		return m.UserVolume
	}
	return ""
}

func (m *StoredToken) GetRwAccess() bool {
	if m != nil {
		return m.RwAccess
	}
	return false
}

func (m *StoredToken) GetUserIp() string {
	if m != nil {
		return m.UserIp
	}
	return ""
}

func (m *StoredToken) GetPartTokenId() *UUID {
	if m != nil {
		return m.PartTokenId
	}
	return nil
}

func (m *StoredToken) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *StoredToken) GetLifeTime() *google_protobuf1.Duration {
	if m != nil {
		return m.LifeTime
	}
	return nil
}

type AccessObject struct {
	Label string `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
	// @inject_tag: binding:"uuid"
	Id     string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty" binding:"uuid"`
	Access string `protobuf:"bytes,3,opt,name=access" json:"access,omitempty"`
}

func (m *AccessObject) Reset()                    { *m = AccessObject{} }
func (m *AccessObject) String() string            { return proto.CompactTextString(m) }
func (*AccessObject) ProtoMessage()               {}
func (*AccessObject) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *AccessObject) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *AccessObject) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AccessObject) GetAccess() string {
	if m != nil {
		return m.Access
	}
	return ""
}

type ResourcesAccess struct {
	Namespace []*AccessObject `protobuf:"bytes,1,rep,name=namespace" json:"namespace,omitempty"`
	Volume    []*AccessObject `protobuf:"bytes,2,rep,name=volume" json:"volume,omitempty"`
}

func (m *ResourcesAccess) Reset()                    { *m = ResourcesAccess{} }
func (m *ResourcesAccess) String() string            { return proto.CompactTextString(m) }
func (*ResourcesAccess) ProtoMessage()               {}
func (*ResourcesAccess) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ResourcesAccess) GetNamespace() []*AccessObject {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ResourcesAccess) GetVolume() []*AccessObject {
	if m != nil {
		return m.Volume
	}
	return nil
}

type StoredTokenForUser struct {
	TokenId   *UUID  `protobuf:"bytes,1,opt,name=token_id,json=tokenId" json:"token_id,omitempty"`
	UserAgent string `protobuf:"bytes,2,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
	// @inject_tag: binding:"ip"
	Ip        string `protobuf:"bytes,3,opt,name=ip" json:"ip,omitempty" binding:"ip"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *StoredTokenForUser) Reset()                    { *m = StoredTokenForUser{} }
func (m *StoredTokenForUser) String() string            { return proto.CompactTextString(m) }
func (*StoredTokenForUser) ProtoMessage()               {}
func (*StoredTokenForUser) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *StoredTokenForUser) GetTokenId() *UUID {
	if m != nil {
		return m.TokenId
	}
	return nil
}

func (m *StoredTokenForUser) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *StoredTokenForUser) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *StoredTokenForUser) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*StoredToken)(nil), "StoredToken")
	proto.RegisterType((*AccessObject)(nil), "AccessObject")
	proto.RegisterType((*ResourcesAccess)(nil), "ResourcesAccess")
	proto.RegisterType((*StoredTokenForUser)(nil), "StoredTokenForUser")
}

func init() { proto.RegisterFile("auth_types.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x25, 0xad, 0x6d, 0x93, 0x9b, 0xed, 0x2a, 0x83, 0xe8, 0x58, 0x71, 0x37, 0x14, 0x16, 0x2a,
	0x42, 0x16, 0x56, 0x10, 0x7c, 0xac, 0x2c, 0x42, 0x41, 0x54, 0x62, 0xeb, 0x83, 0x2f, 0x61, 0x9a,
	0xb9, 0xad, 0xa3, 0x49, 0x66, 0x98, 0x4c, 0x5c, 0x7c, 0xf7, 0x57, 0xfd, 0x0f, 0x99, 0xc9, 0x74,
	0x1b, 0xd4, 0xb7, 0x7d, 0xcb, 0x3d, 0xe7, 0xcc, 0xdc, 0xb9, 0xe7, 0xdc, 0xc0, 0x03, 0xd6, 0x9a,
	0xaf, 0xb9, 0xf9, 0xa9, 0xb0, 0x49, 0x95, 0x96, 0x46, 0xce, 0xa0, 0x6d, 0x05, 0xf7, 0xdf, 0xe7,
	0x7b, 0x29, 0xf7, 0x25, 0x5e, 0xba, 0x6a, 0xdb, 0xee, 0x2e, 0x8d, 0xa8, 0xb0, 0x31, 0xac, 0x52,
	0x5e, 0x70, 0xf6, 0xb7, 0x80, 0xb7, 0x9a, 0x19, 0x21, 0xeb, 0x8e, 0x9f, 0xff, 0x1e, 0x42, 0xfc,
	0xc9, 0x48, 0x8d, 0x7c, 0x2d, 0xbf, 0x63, 0x4d, 0x12, 0x08, 0x8d, 0xfd, 0xc8, 0x05, 0xa7, 0x41,
	0x12, 0x2c, 0xe2, 0xab, 0x51, 0xba, 0xd9, 0xac, 0xae, 0xb3, 0x89, 0x83, 0x57, 0x9c, 0x3c, 0x03,
	0x68, 0x1b, 0xd4, 0x39, 0xdb, 0x63, 0x6d, 0xe8, 0x20, 0x09, 0x16, 0x51, 0x16, 0x59, 0x64, 0x69,
	0x01, 0x32, 0x83, 0x50, 0x95, 0xcc, 0xec, 0xa4, 0xae, 0xe8, 0xd0, 0x91, 0xb7, 0x35, 0x49, 0x20,
	0xde, 0x89, 0x7a, 0x8f, 0x5a, 0x69, 0x51, 0x1b, 0x7a, 0xcf, 0xd1, 0x7d, 0x88, 0x9c, 0xc1, 0xc4,
	0x5d, 0x2e, 0x38, 0x1d, 0xf5, 0xbb, 0x8f, 0x2d, 0xba, 0xe2, 0xe4, 0x29, 0xb8, 0x56, 0xb9, 0x96,
	0x25, 0xd2, 0x71, 0x77, 0xbd, 0x05, 0x32, 0x59, 0x22, 0xb9, 0x80, 0x53, 0x47, 0xd6, 0xac, 0xc2,
	0x46, 0xb1, 0x02, 0xe9, 0xc4, 0x29, 0xa6, 0x16, 0x7d, 0x7f, 0x00, 0xc9, 0x39, 0xc4, 0x4e, 0xf6,
	0x43, 0x96, 0x6d, 0x85, 0x34, 0x74, 0x1a, 0x37, 0xd3, 0x67, 0x87, 0xd8, 0x26, 0xfa, 0x26, 0x67,
	0x45, 0x81, 0x4d, 0x43, 0xa3, 0x24, 0x58, 0x84, 0x59, 0xa8, 0x6f, 0x96, 0xae, 0x26, 0x8f, 0x0f,
	0x2f, 0x54, 0x14, 0xdc, 0xc9, 0xee, 0x69, 0x8a, 0x3c, 0x87, 0xa9, 0x62, 0xda, 0xe4, 0xb7, 0xf6,
	0xc5, 0xfd, 0x01, 0x62, 0xcb, 0xad, 0xbd, 0x85, 0xaf, 0x01, 0x0a, 0x8d, 0xcc, 0x20, 0xcf, 0x99,
	0xa1, 0x27, 0x4e, 0x37, 0x4b, 0xbb, 0xa4, 0xd2, 0x43, 0x52, 0xe9, 0xfa, 0x10, 0x65, 0x16, 0x79,
	0xf5, 0xd2, 0x90, 0x57, 0x10, 0x95, 0x62, 0x87, 0xb9, 0xcd, 0x99, 0x4e, 0xdd, 0xc9, 0x27, 0xff,
	0x9c, 0xbc, 0xf6, 0x19, 0x67, 0xa1, 0xd5, 0xda, 0x7b, 0xe6, 0xef, 0xe0, 0xa4, 0x1b, 0xe0, 0xc3,
	0xf6, 0x1b, 0x16, 0x86, 0x3c, 0x84, 0x51, 0xc9, 0xb6, 0x58, 0xba, 0x90, 0xa3, 0xac, 0x2b, 0xc8,
	0x29, 0x0c, 0x04, 0xf7, 0x99, 0x0e, 0x04, 0x27, 0x8f, 0x60, 0xec, 0x6d, 0xe8, 0xa2, 0xf4, 0xd5,
	0x1c, 0xe1, 0x7e, 0x86, 0x8d, 0x6c, 0x75, 0x81, 0x8d, 0xf7, 0xe5, 0x05, 0x44, 0x47, 0xdf, 0x83,
	0x64, 0xb8, 0x88, 0xaf, 0xa6, 0x69, 0xbf, 0x65, 0x76, 0xe4, 0xc9, 0x05, 0x8c, 0xbd, 0xfb, 0x83,
	0xff, 0x29, 0x3d, 0x39, 0xff, 0x15, 0x00, 0xe9, 0x2d, 0xe7, 0x5b, 0xa9, 0x37, 0x0d, 0xea, 0xbb,
	0xef, 0xa8, 0x1d, 0x53, 0xf9, 0x91, 0x06, 0x42, 0x59, 0x79, 0x2f, 0x8f, 0x6e, 0x2d, 0x8f, 0x9e,
	0xbf, 0x89, 0xbf, 0x44, 0xf6, 0x27, 0xfc, 0xe8, 0xec, 0x1d, 0x3b, 0x97, 0x5f, 0xfe, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x59, 0x54, 0x68, 0x88, 0x98, 0x03, 0x00, 0x00,
}
