// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Users struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Passwd               string   `protobuf:"bytes,3,opt,name=passwd,proto3" json:"passwd,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{1}
}

func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (m *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(m, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Users) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Users) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

func (m *Users) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type Id struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{2}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UpdateRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	User                 *Users   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{3}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateRequest) GetUser() *Users {
	if m != nil {
		return m.User
	}
	return nil
}

type Status struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{4}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type DeleteResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{5}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*Users)(nil), "Users")
	proto.RegisterType((*Id)(nil), "Id")
	proto.RegisterType((*UpdateRequest)(nil), "UpdateRequest")
	proto.RegisterType((*Status)(nil), "Status")
	proto.RegisterType((*DeleteResponse)(nil), "DeleteResponse")
}

func init() {
	proto.RegisterFile("proto/user.proto", fileDescriptor_d570e3e37e5899c5)
}

var fileDescriptor_d570e3e37e5899c5 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4f, 0x6b, 0xfa, 0x40,
	0x10, 0x25, 0xd1, 0xc4, 0x9f, 0x23, 0x3f, 0x5b, 0x16, 0x29, 0x41, 0x4a, 0x2b, 0x4b, 0x0f, 0xd2,
	0xc3, 0x5a, 0xec, 0xb1, 0xa7, 0x4a, 0x4b, 0xc9, 0x35, 0xe2, 0xa5, 0xb7, 0x95, 0x1d, 0x8a, 0x90,
	0x68, 0xba, 0xb3, 0xb1, 0xf8, 0xb1, 0xfa, 0x0d, 0xcb, 0xce, 0x6a, 0xc0, 0xd2, 0x5b, 0xde, 0x9f,
	0x99, 0x37, 0x2f, 0x0b, 0x97, 0xb5, 0xdd, 0xb9, 0xdd, 0xac, 0x21, 0xb4, 0x8a, 0x3f, 0x65, 0x0f,
	0x92, 0xd7, 0xaa, 0x76, 0x07, 0xa9, 0x21, 0x59, 0x11, 0x5a, 0x12, 0x43, 0x88, 0x37, 0x26, 0x8b,
	0x26, 0xd1, 0xb4, 0x53, 0xc4, 0x1b, 0x23, 0xc6, 0xf0, 0xcf, 0xfb, 0xb7, 0xba, 0xc2, 0x2c, 0x9e,
	0x44, 0xd3, 0x7e, 0xd1, 0x62, 0x71, 0x05, 0x69, 0xad, 0x89, 0xbe, 0x4c, 0xd6, 0x61, 0xe5, 0x88,
	0xc4, 0x08, 0x12, 0xac, 0xf4, 0xa6, 0xcc, 0xba, 0x4c, 0x07, 0x20, 0x47, 0x10, 0xe7, 0xe6, 0xf7,
	0x7e, 0xf9, 0x04, 0xff, 0x57, 0xb5, 0xd1, 0x0e, 0x0b, 0xfc, 0x6c, 0x90, 0xdc, 0x1f, 0x07, 0x74,
	0x7d, 0x20, 0x87, 0x0f, 0xe6, 0xa9, 0xe2, 0x33, 0x0b, 0xe6, 0xe4, 0x0d, 0xa4, 0x4b, 0xa7, 0x5d,
	0x43, 0x3e, 0x72, 0xaf, 0xcb, 0x06, 0x79, 0x30, 0x29, 0x02, 0x90, 0xf7, 0x30, 0x7c, 0xc1, 0x12,
	0xfd, 0x72, 0xaa, 0x77, 0x5b, 0x42, 0x91, 0x41, 0xaf, 0x42, 0x22, 0xfd, 0x11, 0x9c, 0xfd, 0xe2,
	0x04, 0xe7, 0xdf, 0x11, 0x80, 0xdf, 0xbd, 0x44, 0xbb, 0x47, 0x2b, 0xc6, 0x30, 0x78, 0x43, 0xe7,
	0x89, 0xc5, 0x21, 0x37, 0xa2, 0xa3, 0x72, 0x33, 0x3e, 0x86, 0x8b, 0x5b, 0xd6, 0x9e, 0xcb, 0x32,
	0xc0, 0x54, 0xf1, 0x3f, 0x3c, 0xc9, 0x0f, 0x91, 0xb8, 0x06, 0xc8, 0xb7, 0x84, 0x96, 0xe7, 0xc5,
	0x91, 0x6f, 0xc7, 0xef, 0x00, 0x42, 0x65, 0x56, 0x87, 0xea, 0xac, 0x7f, 0xeb, 0x92, 0x00, 0xe1,
	0x76, 0x76, 0x71, 0xfe, 0x85, 0x3a, 0x6f, 0xb3, 0x48, 0xdf, 0xbb, 0x6a, 0x56, 0xaf, 0xd7, 0x29,
	0xbf, 0xe6, 0xe3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x6a, 0xe1, 0xc0, 0xe1, 0x01, 0x00,
	0x00,
}
