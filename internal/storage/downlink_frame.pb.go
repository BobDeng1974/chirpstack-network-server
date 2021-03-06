// Code generated by protoc-gen-go. DO NOT EDIT.
// source: downlink_frames.proto

package storage

import (
	fmt "fmt"
	gw "github.com/brocaar/chirpstack-api/go/v3/gw"
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

type DownlinkFrame struct {
	// Token.
	Token uint32 `protobuf:"varint,1,opt,name=token,proto3" json:"token,omitempty"`
	// DevEUI.
	DevEui []byte `protobuf:"bytes,2,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	// Multicast Group ID.
	MulticastGroupId []byte `protobuf:"bytes,3,opt,name=multicast_group_id,json=multicastGroupId,proto3" json:"multicast_group_id,omitempty"`
	// Downlink frames.
	DownlinkFrame *gw.DownlinkFrame `protobuf:"bytes,4,opt,name=downlink_frame,json=downlinkFrame,proto3" json:"downlink_frame,omitempty"`
	// Routing Profile ID.
	RoutingProfileId []byte `protobuf:"bytes,5,opt,name=routing_profile_id,json=routingProfileId,proto3" json:"routing_profile_id,omitempty"`
	// Downlink frame-counter (full).
	FCnt uint32 `protobuf:"varint,6,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
	// Encrypted FOpts (LoRaWAN 1.1).
	EncryptedFopts bool `protobuf:"varint,7,opt,name=encrypted_fopts,json=encryptedFopts,proto3" json:"encrypted_fopts,omitempty"`
	// Network session encryption key (for FOpts).
	NwkSEncKey           []byte   `protobuf:"bytes,8,opt,name=nwk_s_enc_key,json=nwkSEncKey,proto3" json:"nwk_s_enc_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownlinkFrame) Reset()         { *m = DownlinkFrame{} }
func (m *DownlinkFrame) String() string { return proto.CompactTextString(m) }
func (*DownlinkFrame) ProtoMessage()    {}
func (*DownlinkFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_a06c04f39a283b6b, []int{0}
}

func (m *DownlinkFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownlinkFrame.Unmarshal(m, b)
}
func (m *DownlinkFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownlinkFrame.Marshal(b, m, deterministic)
}
func (m *DownlinkFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownlinkFrame.Merge(m, src)
}
func (m *DownlinkFrame) XXX_Size() int {
	return xxx_messageInfo_DownlinkFrame.Size(m)
}
func (m *DownlinkFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_DownlinkFrame.DiscardUnknown(m)
}

var xxx_messageInfo_DownlinkFrame proto.InternalMessageInfo

func (m *DownlinkFrame) GetToken() uint32 {
	if m != nil {
		return m.Token
	}
	return 0
}

func (m *DownlinkFrame) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *DownlinkFrame) GetMulticastGroupId() []byte {
	if m != nil {
		return m.MulticastGroupId
	}
	return nil
}

func (m *DownlinkFrame) GetDownlinkFrame() *gw.DownlinkFrame {
	if m != nil {
		return m.DownlinkFrame
	}
	return nil
}

func (m *DownlinkFrame) GetRoutingProfileId() []byte {
	if m != nil {
		return m.RoutingProfileId
	}
	return nil
}

func (m *DownlinkFrame) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

func (m *DownlinkFrame) GetEncryptedFopts() bool {
	if m != nil {
		return m.EncryptedFopts
	}
	return false
}

func (m *DownlinkFrame) GetNwkSEncKey() []byte {
	if m != nil {
		return m.NwkSEncKey
	}
	return nil
}

func init() {
	proto.RegisterType((*DownlinkFrame)(nil), "storage.DownlinkFrame")
}

func init() { proto.RegisterFile("downlink_frames.proto", fileDescriptor_a06c04f39a283b6b) }

var fileDescriptor_a06c04f39a283b6b = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xe9, 0xdc, 0xd6, 0x91, 0xd9, 0xa9, 0x51, 0x31, 0x78, 0xaa, 0x5e, 0xec, 0x61, 0x54,
	0xd0, 0x8b, 0x77, 0xdd, 0x64, 0x78, 0x91, 0xfa, 0x03, 0x42, 0x6d, 0xbe, 0x86, 0xd0, 0xee, 0x4b,
	0x49, 0xd3, 0x85, 0xfe, 0x30, 0xff, 0x9f, 0xb4, 0x1d, 0x83, 0x1e, 0xf3, 0x3e, 0x2f, 0xbc, 0x4f,
	0x3e, 0x72, 0x2b, 0xb4, 0xc3, 0x52, 0x61, 0xc1, 0x73, 0x93, 0xee, 0xa1, 0x8e, 0x2b, 0xa3, 0xad,
	0xa6, 0x7e, 0x6d, 0xb5, 0x49, 0x25, 0xdc, 0x2f, 0xa5, 0x7b, 0x96, 0x6e, 0x48, 0x1f, 0xff, 0x26,
	0x24, 0xf8, 0x38, 0xf6, 0xb7, 0x5d, 0x9d, 0xde, 0x90, 0x99, 0xd5, 0x05, 0x20, 0xf3, 0x42, 0x2f,
	0x0a, 0x92, 0xe1, 0x41, 0xef, 0x88, 0x2f, 0xe0, 0xc0, 0xa1, 0x51, 0x6c, 0x12, 0x7a, 0xd1, 0x79,
	0x32, 0x17, 0x70, 0xd8, 0x34, 0x8a, 0xae, 0x09, 0xdd, 0x37, 0xa5, 0x55, 0x59, 0x5a, 0x5b, 0x2e,
	0x8d, 0x6e, 0x2a, 0xae, 0x04, 0x3b, 0xeb, 0x3b, 0x97, 0x27, 0xf2, 0xd9, 0x81, 0x9d, 0xa0, 0x6f,
	0x64, 0x35, 0xb6, 0x63, 0xd3, 0xd0, 0x8b, 0x96, 0x2f, 0x57, 0xb1, 0x74, 0xf1, 0xc8, 0x23, 0x09,
	0xc4, 0x48, 0x6b, 0x4d, 0xa8, 0xd1, 0x8d, 0x55, 0x28, 0x79, 0x65, 0x74, 0xae, 0x4a, 0xe8, 0x76,
	0x66, 0xc3, 0xce, 0x91, 0x7c, 0x0f, 0x60, 0x27, 0xe8, 0x35, 0x99, 0xe5, 0x3c, 0x43, 0xcb, 0xe6,
	0xfd, 0x27, 0xa6, 0xf9, 0x3b, 0x5a, 0xfa, 0x44, 0x2e, 0x00, 0x33, 0xd3, 0x56, 0x16, 0x04, 0xcf,
	0x75, 0x65, 0x6b, 0xe6, 0x87, 0x5e, 0xb4, 0x48, 0x56, 0xa7, 0x78, 0xdb, 0xa5, 0xf4, 0x81, 0x04,
	0xe8, 0x0a, 0x5e, 0x73, 0xc0, 0x8c, 0x17, 0xd0, 0xb2, 0x45, 0x3f, 0x43, 0xd0, 0x15, 0x3f, 0x1b,
	0xcc, 0xbe, 0xa0, 0xfd, 0x9d, 0xf7, 0xe7, 0x7b, 0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x9c,
	0x67, 0xb9, 0x6d, 0x01, 0x00, 0x00,
}
