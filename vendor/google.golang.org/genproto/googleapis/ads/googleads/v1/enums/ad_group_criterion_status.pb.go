// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/ad_group_criterion_status.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The possible statuses of an AdGroupCriterion.
type AdGroupCriterionStatusEnum_AdGroupCriterionStatus int32

const (
	// No value has been specified.
	AdGroupCriterionStatusEnum_UNSPECIFIED AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdGroupCriterionStatusEnum_UNKNOWN AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 1
	// The ad group criterion is enabled.
	AdGroupCriterionStatusEnum_ENABLED AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 2
	// The ad group criterion is paused.
	AdGroupCriterionStatusEnum_PAUSED AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 3
	// The ad group criterion is removed.
	AdGroupCriterionStatusEnum_REMOVED AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 4
)

var AdGroupCriterionStatusEnum_AdGroupCriterionStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "PAUSED",
	4: "REMOVED",
}

var AdGroupCriterionStatusEnum_AdGroupCriterionStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"PAUSED":      3,
	"REMOVED":     4,
}

func (x AdGroupCriterionStatusEnum_AdGroupCriterionStatus) String() string {
	return proto.EnumName(AdGroupCriterionStatusEnum_AdGroupCriterionStatus_name, int32(x))
}

func (AdGroupCriterionStatusEnum_AdGroupCriterionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_53e0a4674d62bdc5, []int{0, 0}
}

// Message describing AdGroupCriterion statuses.
type AdGroupCriterionStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupCriterionStatusEnum) Reset()         { *m = AdGroupCriterionStatusEnum{} }
func (m *AdGroupCriterionStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterionStatusEnum) ProtoMessage()    {}
func (*AdGroupCriterionStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_53e0a4674d62bdc5, []int{0}
}

func (m *AdGroupCriterionStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterionStatusEnum.Unmarshal(m, b)
}
func (m *AdGroupCriterionStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterionStatusEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupCriterionStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterionStatusEnum.Merge(m, src)
}
func (m *AdGroupCriterionStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterionStatusEnum.Size(m)
}
func (m *AdGroupCriterionStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterionStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterionStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.AdGroupCriterionStatusEnum_AdGroupCriterionStatus", AdGroupCriterionStatusEnum_AdGroupCriterionStatus_name, AdGroupCriterionStatusEnum_AdGroupCriterionStatus_value)
	proto.RegisterType((*AdGroupCriterionStatusEnum)(nil), "google.ads.googleads.v1.enums.AdGroupCriterionStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/ad_group_criterion_status.proto", fileDescriptor_53e0a4674d62bdc5)
}

var fileDescriptor_53e0a4674d62bdc5 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0xf3, 0x30,
	0x1c, 0xc6, 0xdf, 0x75, 0x2f, 0x13, 0xb2, 0x83, 0xa5, 0x07, 0x0f, 0xd3, 0x1d, 0xb6, 0x0f, 0x90,
	0x52, 0xbc, 0x45, 0x3c, 0xa4, 0x6b, 0x1c, 0x43, 0xed, 0x8a, 0x63, 0x15, 0xa4, 0x30, 0xe2, 0x52,
	0x42, 0x61, 0x4b, 0x4a, 0x93, 0xee, 0xe0, 0xc7, 0xf1, 0xe8, 0x47, 0xf1, 0xa3, 0x78, 0xf2, 0x23,
	0x48, 0xd2, 0xb5, 0xa7, 0xe9, 0x25, 0x3c, 0xc9, 0xf3, 0x7f, 0x7e, 0x3c, 0xf9, 0x83, 0x5b, 0x2e,
	0x25, 0xdf, 0xe5, 0x3e, 0x65, 0xca, 0x6f, 0xa4, 0x51, 0x87, 0xc0, 0xcf, 0x45, 0xbd, 0x57, 0x3e,
	0x65, 0x1b, 0x5e, 0xc9, 0xba, 0xdc, 0x6c, 0xab, 0x42, 0xe7, 0x55, 0x21, 0xc5, 0x46, 0x69, 0xaa,
	0x6b, 0x05, 0xcb, 0x4a, 0x6a, 0xe9, 0x8d, 0x9b, 0x0c, 0xa4, 0x4c, 0xc1, 0x2e, 0x0e, 0x0f, 0x01,
	0xb4, 0xf1, 0xd1, 0x55, 0x4b, 0x2f, 0x0b, 0x9f, 0x0a, 0x21, 0x35, 0xd5, 0x85, 0x14, 0xc7, 0xf0,
	0xf4, 0x0d, 0x8c, 0x30, 0x9b, 0x1b, 0xfc, 0xac, 0xa5, 0xaf, 0x2c, 0x9c, 0x88, 0x7a, 0x3f, 0xcd,
	0xc0, 0xc5, 0x69, 0xd7, 0x3b, 0x07, 0xc3, 0x75, 0xbc, 0x4a, 0xc8, 0x6c, 0x71, 0xb7, 0x20, 0x91,
	0xfb, 0xcf, 0x1b, 0x82, 0xb3, 0x75, 0x7c, 0x1f, 0x2f, 0x9f, 0x63, 0xb7, 0x67, 0x2e, 0x24, 0xc6,
	0xe1, 0x03, 0x89, 0x5c, 0xc7, 0x03, 0x60, 0x90, 0xe0, 0xf5, 0x8a, 0x44, 0x6e, 0xdf, 0x18, 0x4f,
	0xe4, 0x71, 0x99, 0x92, 0xc8, 0xfd, 0x1f, 0x7e, 0xf7, 0xc0, 0x64, 0x2b, 0xf7, 0xf0, 0xcf, 0xfe,
	0xe1, 0xe5, 0xe9, 0x06, 0x89, 0xa9, 0x9f, 0xf4, 0x5e, 0xc2, 0x63, 0x9a, 0xcb, 0x1d, 0x15, 0x1c,
	0xca, 0x8a, 0xfb, 0x3c, 0x17, 0xf6, 0x73, 0xed, 0x32, 0xcb, 0x42, 0xfd, 0xb2, 0xdb, 0x1b, 0x7b,
	0xbe, 0x3b, 0xfd, 0x39, 0xc6, 0x1f, 0xce, 0x78, 0xde, 0xa0, 0x30, 0x53, 0xb0, 0x91, 0x46, 0xa5,
	0x01, 0x34, 0xbb, 0x50, 0x9f, 0xad, 0x9f, 0x61, 0xa6, 0xb2, 0xce, 0xcf, 0xd2, 0x20, 0xb3, 0xfe,
	0x97, 0x33, 0x69, 0x1e, 0x11, 0xc2, 0x4c, 0x21, 0xd4, 0x4d, 0x20, 0x94, 0x06, 0x08, 0xd9, 0x99,
	0xd7, 0x81, 0x2d, 0x76, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xfe, 0x68, 0x64, 0xf3, 0x01,
	0x00, 0x00,
}
