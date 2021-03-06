// Code generated by protoc-gen-go.
// source: common/msp_principal.proto
// DO NOT EDIT!

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MSPPrincipal_Classification int32

const (
	MSPPrincipal_ByMSPRole MSPPrincipal_Classification = 0
	// one of a member of MSP network, and the one of an
	// administrator of an MSP network
	MSPPrincipal_ByOrganizationUnit MSPPrincipal_Classification = 1
	// groupping of entities, per MSP affiliation
	// E.g., this can well be represented by an MSP's
	// Organization unit
	MSPPrincipal_ByIdentity MSPPrincipal_Classification = 2
)

var MSPPrincipal_Classification_name = map[int32]string{
	0: "ByMSPRole",
	1: "ByOrganizationUnit",
	2: "ByIdentity",
}
var MSPPrincipal_Classification_value = map[string]int32{
	"ByMSPRole":          0,
	"ByOrganizationUnit": 1,
	"ByIdentity":         2,
}

func (x MSPPrincipal_Classification) String() string {
	return proto.EnumName(MSPPrincipal_Classification_name, int32(x))
}
func (MSPPrincipal_Classification) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 0}
}

type MSPRole_MSPRoleType int32

const (
	MSPRole_Member MSPRole_MSPRoleType = 0
	MSPRole_Admin  MSPRole_MSPRoleType = 1
)

var MSPRole_MSPRoleType_name = map[int32]string{
	0: "Member",
	1: "Admin",
}
var MSPRole_MSPRoleType_value = map[string]int32{
	"Member": 0,
	"Admin":  1,
}

func (x MSPRole_MSPRoleType) String() string {
	return proto.EnumName(MSPRole_MSPRoleType_name, int32(x))
}
func (MSPRole_MSPRoleType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{2, 0} }

// MSPPrincipal aims to represent an MSP-centric set of identities.
// In particular, this structure allows for definition of
//  - a group of identities that are member of the same MSP
//  - a group of identities that are member of the same organization unit
//    in the same MSP
//  - a group of identities that are administering a specific MSP
//  - a specific identity
// Expressing these groups is done given two fields of the fields below
//  - Classification, that defines the type of classification of identities
//    in an MSP this principal would be defined on; Classification can take
//    three values:
//     (i)  ByMSPRole: that represents a classification of identities within
//          MSP based on one of the two pre-defined MSP rules, "member" and "admin"
//     (ii) ByOrganizationUnit: that represents a classification of identities
//          within MSP based on the organization unit an identity belongs to
//     (iii)ByIdentity that denotes that MSPPrincipal is mapped to a single
//          identity/certificate; this would mean that the Principal bytes
//          message
type MSPPrincipal struct {
	// Classification describes the way that one should process
	// Principal. An Classification value of "ByOrganizationUnit" reflects
	// that "Principal" contains the name of an organization this MSP
	// handles. A Classification value "ByIdentity" means that
	// "Principal" contains a specific identity. Default value
	// denotes that Principal contains one of the groups by
	// default supported by all MSPs ("admin" or "member").
	PrincipalClassification MSPPrincipal_Classification `protobuf:"varint,1,opt,name=PrincipalClassification,enum=common.MSPPrincipal_Classification" json:"PrincipalClassification,omitempty"`
	// Principal completes the policy principal definition. For the default
	// principal types, Principal can be either "Admin" or "Member".
	// For the ByOrganizationUnit/ByIdentity values of Classification,
	// PolicyPrincipal acquires its value from an organization unit or
	// identity, respectively.
	Principal []byte `protobuf:"bytes,2,opt,name=Principal,proto3" json:"Principal,omitempty"`
}

func (m *MSPPrincipal) Reset()                    { *m = MSPPrincipal{} }
func (m *MSPPrincipal) String() string            { return proto.CompactTextString(m) }
func (*MSPPrincipal) ProtoMessage()               {}
func (*MSPPrincipal) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// OrganizationUnit governs the organization of the Principal
// field of a policy principal when a specific organization unity members
// are to be defined within a policy principal.
type OrganizationUnit struct {
	// MSPIdentifier represents the identifier of the MSP this organization unit
	// refers to
	MSPIdentifier string `protobuf:"bytes,1,opt,name=MSPIdentifier" json:"MSPIdentifier,omitempty"`
	// OrganizationUnitIdentifier defines the organization unit under the
	// MSP identified with MSPIdentifier
	OrganizationUnitIdentifier string `protobuf:"bytes,2,opt,name=OrganizationUnitIdentifier" json:"OrganizationUnitIdentifier,omitempty"`
}

func (m *OrganizationUnit) Reset()                    { *m = OrganizationUnit{} }
func (m *OrganizationUnit) String() string            { return proto.CompactTextString(m) }
func (*OrganizationUnit) ProtoMessage()               {}
func (*OrganizationUnit) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

// MSPRole governs the organization of the Principal
// field of an MSPPrincipal when it aims to define one of the
// two dedicated roles within an MSP: Admin and Members.
type MSPRole struct {
	// MSPIdentifier represents the identifier of the MSP this principal
	// refers to
	MSPIdentifier string `protobuf:"bytes,1,opt,name=MSPIdentifier" json:"MSPIdentifier,omitempty"`
	// MSPRoleType defines which of the available, pre-defined MSP-roles
	// an identiy should posess inside the MSP with identifier MSPidentifier
	Role MSPRole_MSPRoleType `protobuf:"varint,2,opt,name=Role,enum=common.MSPRole_MSPRoleType" json:"Role,omitempty"`
}

func (m *MSPRole) Reset()                    { *m = MSPRole{} }
func (m *MSPRole) String() string            { return proto.CompactTextString(m) }
func (*MSPRole) ProtoMessage()               {}
func (*MSPRole) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func init() {
	proto.RegisterType((*MSPPrincipal)(nil), "common.MSPPrincipal")
	proto.RegisterType((*OrganizationUnit)(nil), "common.OrganizationUnit")
	proto.RegisterType((*MSPRole)(nil), "common.MSPRole")
	proto.RegisterEnum("common.MSPPrincipal_Classification", MSPPrincipal_Classification_name, MSPPrincipal_Classification_value)
	proto.RegisterEnum("common.MSPRole_MSPRoleType", MSPRole_MSPRoleType_name, MSPRole_MSPRoleType_value)
}

func init() { proto.RegisterFile("common/msp_principal.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x97, 0xa1, 0x93, 0x3e, 0xb7, 0x52, 0x72, 0xd0, 0x31, 0x3d, 0x8c, 0xba, 0xc3, 0x40,
	0x6c, 0x61, 0xde, 0x05, 0xeb, 0x41, 0x3c, 0x14, 0x4b, 0xa7, 0x17, 0x41, 0xa4, 0xed, 0xb2, 0xed,
	0x41, 0x9b, 0x86, 0x34, 0x82, 0xf1, 0x0f, 0xf0, 0x2f, 0xf4, 0x0f, 0x92, 0xa5, 0x6e, 0x76, 0x03,
	0xc5, 0x53, 0x79, 0xdf, 0xf7, 0xfb, 0xbe, 0xd7, 0x84, 0xc0, 0x20, 0x2b, 0x8b, 0xa2, 0xe4, 0x7e,
	0x51, 0x89, 0x17, 0x21, 0x91, 0x67, 0x28, 0x92, 0xdc, 0x13, 0xb2, 0x54, 0x25, 0xed, 0xd4, 0x9e,
	0xfb, 0x49, 0xa0, 0x1b, 0x4e, 0xa3, 0x68, 0x6d, 0xd3, 0x67, 0x38, 0xde, 0x0c, 0x37, 0x79, 0x52,
	0x55, 0x38, 0xc7, 0x2c, 0x51, 0x58, 0xf2, 0x3e, 0x19, 0x92, 0xb1, 0x3d, 0x39, 0xf3, 0xea, 0xa8,
	0xd7, 0x8c, 0x79, 0xdb, 0x68, 0xfc, 0x5b, 0x07, 0x3d, 0x05, 0x6b, 0x63, 0xf5, 0xdb, 0x43, 0x32,
	0xee, 0xc6, 0x3f, 0x82, 0x7b, 0x0b, 0xf6, 0x0e, 0xdf, 0x03, 0x2b, 0xd0, 0xe1, 0x34, 0x8a, 0xcb,
	0x9c, 0x39, 0x2d, 0x7a, 0x04, 0x34, 0xd0, 0xf7, 0x72, 0x91, 0x70, 0x7c, 0x37, 0xc0, 0x23, 0x47,
	0xe5, 0x10, 0x6a, 0x03, 0x04, 0xfa, 0x6e, 0xc6, 0xb8, 0x42, 0xa5, 0x9d, 0xb6, 0xfb, 0x06, 0xce,
	0x2e, 0x45, 0x47, 0xd0, 0x0b, 0xa7, 0x51, 0x0d, 0xcd, 0x91, 0x49, 0x73, 0x1e, 0x2b, 0xde, 0x16,
	0xe9, 0x15, 0x0c, 0x76, 0x93, 0x8d, 0x48, 0xdb, 0x44, 0xfe, 0x20, 0xdc, 0x0f, 0x02, 0x07, 0xdf,
	0xff, 0xfb, 0xcf, 0x8d, 0x3e, 0xec, 0xad, 0x68, 0xd3, 0x6d, 0x4f, 0x4e, 0x1a, 0xd7, 0xbb, 0x92,
	0xd7, 0xdf, 0x07, 0x2d, 0x58, 0x6c, 0x40, 0x77, 0x04, 0x87, 0x0d, 0x91, 0x02, 0x74, 0x42, 0x56,
	0xa4, 0x4c, 0x3a, 0x2d, 0x6a, 0xc1, 0xfe, 0xf5, 0xac, 0x40, 0xee, 0x90, 0xe0, 0xe2, 0xe9, 0x7c,
	0x81, 0x6a, 0xf9, 0x9a, 0xae, 0x0a, 0xfd, 0xa5, 0x16, 0x4c, 0xe6, 0x6c, 0xb6, 0x60, 0xd2, 0x9f,
	0x27, 0xa9, 0xc4, 0xcc, 0x37, 0x0f, 0xa1, 0xf2, 0xeb, 0x75, 0x69, 0xc7, 0x8c, 0x97, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x81, 0x8f, 0xfe, 0x6f, 0x35, 0x02, 0x00, 0x00,
}
