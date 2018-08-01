// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chainscript.proto

package chainscript

/*
ChainScript is an open standard for representing Proof of Process data.
Proof of Process is a protocol that allows partners to follow the execution
of a shared process.
Proof of Process provides immutability and auditability of every step in the
process.
*/

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

// A segment describes an atomic step in your process.
type Segment struct {
	// The link is the immutable part of a segment.
	// It contains the details of the step.
	Link *Link `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	// The link can be enriched with potentially mutable metadata.
	Meta                 *SegmentMeta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Segment) Reset()         { *m = Segment{} }
func (m *Segment) String() string { return proto.CompactTextString(m) }
func (*Segment) ProtoMessage()    {}
func (*Segment) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainscript_621c374498d1a002, []int{0}
}
func (m *Segment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Segment.Unmarshal(m, b)
}
func (m *Segment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Segment.Marshal(b, m, deterministic)
}
func (dst *Segment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Segment.Merge(dst, src)
}
func (m *Segment) XXX_Size() int {
	return xxx_messageInfo_Segment.Size(m)
}
func (m *Segment) XXX_DiscardUnknown() {
	xxx_messageInfo_Segment.DiscardUnknown(m)
}

var xxx_messageInfo_Segment proto.InternalMessageInfo

func (m *Segment) GetLink() *Link {
	if m != nil {
		return m.Link
	}
	return nil
}

func (m *Segment) GetMeta() *SegmentMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

// Segment metadata. This is the potentially mutable part of a segment.
// It contains some invariants (hash of the immutable link) and evidences
// for the link that can be produced after the link is created.
type SegmentMeta struct {
	// Hash of the segment's link.
	LinkHash []byte `protobuf:"bytes,1,opt,name=link_hash,json=linkHash,proto3" json:"link_hash,omitempty"`
	// Evidences produced for the segment's link.
	Evidences            []*Evidence `protobuf:"bytes,10,rep,name=evidences,proto3" json:"evidences,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SegmentMeta) Reset()         { *m = SegmentMeta{} }
func (m *SegmentMeta) String() string { return proto.CompactTextString(m) }
func (*SegmentMeta) ProtoMessage()    {}
func (*SegmentMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainscript_621c374498d1a002, []int{1}
}
func (m *SegmentMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentMeta.Unmarshal(m, b)
}
func (m *SegmentMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentMeta.Marshal(b, m, deterministic)
}
func (dst *SegmentMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentMeta.Merge(dst, src)
}
func (m *SegmentMeta) XXX_Size() int {
	return xxx_messageInfo_SegmentMeta.Size(m)
}
func (m *SegmentMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentMeta.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentMeta proto.InternalMessageInfo

func (m *SegmentMeta) GetLinkHash() []byte {
	if m != nil {
		return m.LinkHash
	}
	return nil
}

func (m *SegmentMeta) GetEvidences() []*Evidence {
	if m != nil {
		return m.Evidences
	}
	return nil
}

// Evidences can be used to externally verify a link's existence at a given
// moment in time.
// An evidence can be a proof of inclusion in a public blockchain, a timestamp
// signed by a trusted authority or anything that you trust to provide an
// immutable ordering of your process' steps.
type Evidence struct {
	// Version of the evidence format.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Identifier of the evidence type.
	// For example, in the case of a timestamp on the Bitcoin blockchain,
	// this would be "bitcoin".
	Backend string `protobuf:"bytes,10,opt,name=backend,proto3" json:"backend,omitempty"`
	// Instance of the backend used.
	// For example, in the case of a timestamp on the Bitcoin blockchain,
	// this would be the chain ID (to identify testnet from mainnet).
	Provider string `protobuf:"bytes,11,opt,name=provider,proto3" json:"provider,omitempty"`
	// Data that should be usable offline by any client wishing to validate
	// the evidence.
	// For backwards compatibility, you should update the evidence version
	// when the structure of this proof changes.
	Proof                []byte   `protobuf:"bytes,20,opt,name=proof,proto3" json:"proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Evidence) Reset()         { *m = Evidence{} }
func (m *Evidence) String() string { return proto.CompactTextString(m) }
func (*Evidence) ProtoMessage()    {}
func (*Evidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainscript_621c374498d1a002, []int{2}
}
func (m *Evidence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Evidence.Unmarshal(m, b)
}
func (m *Evidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Evidence.Marshal(b, m, deterministic)
}
func (dst *Evidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Evidence.Merge(dst, src)
}
func (m *Evidence) XXX_Size() int {
	return xxx_messageInfo_Evidence.Size(m)
}
func (m *Evidence) XXX_DiscardUnknown() {
	xxx_messageInfo_Evidence.DiscardUnknown(m)
}

var xxx_messageInfo_Evidence proto.InternalMessageInfo

func (m *Evidence) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Evidence) GetBackend() string {
	if m != nil {
		return m.Backend
	}
	return ""
}

func (m *Evidence) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Evidence) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

// A link is the immutable part of a segment.
// A link contains all the data that represents a process' step.
type Link struct {
	// Version of the link format.
	// You can for example use the git tag of the code used to create the link.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Data representing the process' step details.
	// For backwards compatibility, you should update the link version
	// in meta when the structure/encoding of this field changes.
	Data []byte `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty"`
	// Metadata associated to the process' step.
	// Some of this metadata is used to provide filtering options when
	// fetching links.
	Meta *LinkMeta `protobuf:"bytes,11,opt,name=meta,proto3" json:"meta,omitempty"`
	// (Optional) Signatures of configurable parts of the link.
	Signatures           []*Signature `protobuf:"bytes,20,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainscript_621c374498d1a002, []int{3}
}
func (m *Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Link.Unmarshal(m, b)
}
func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Link.Marshal(b, m, deterministic)
}
func (dst *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(dst, src)
}
func (m *Link) XXX_Size() int {
	return xxx_messageInfo_Link.Size(m)
}
func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

func (m *Link) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Link) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Link) GetMeta() *LinkMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Link) GetSignatures() []*Signature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// A process represents a real-world process that is shared between multiple
// independent actors.
type Process struct {
	// The name of the process.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The current state of the process.
	State                string   `protobuf:"bytes,10,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Process) Reset()         { *m = Process{} }
func (m *Process) String() string { return proto.CompactTextString(m) }
func (*Process) ProtoMessage()    {}
func (*Process) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainscript_621c374498d1a002, []int{4}
}
func (m *Process) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process.Unmarshal(m, b)
}
func (m *Process) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process.Marshal(b, m, deterministic)
}
func (dst *Process) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process.Merge(dst, src)
}
func (m *Process) XXX_Size() int {
	return xxx_messageInfo_Process.Size(m)
}
func (m *Process) XXX_DiscardUnknown() {
	xxx_messageInfo_Process.DiscardUnknown(m)
}

var xxx_messageInfo_Process proto.InternalMessageInfo

func (m *Process) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Process) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

// Metadata associated to a process' step.
// Once included in a segment, this is immutable.
type LinkMeta struct {
	// The Client ID should be set by the client code creating the link.
	// Use a unique ID that easily identifies your library, for example the
	// github url of your repository.
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Hash of the previous link (in the same process).
	PrevLinkHash []byte `protobuf:"bytes,10,opt,name=prev_link_hash,json=prevLinkHash,proto3" json:"prev_link_hash,omitempty"`
	// Priority of the link.
	// Can be used to order and filter search results.
	Priority float64 `protobuf:"fixed64,11,opt,name=priority,proto3" json:"priority,omitempty"`
	// References to related links (potentially in other processes).
	Refs []*LinkReference `protobuf:"bytes,12,rep,name=refs,proto3" json:"refs,omitempty"`
	// A link is a step in a given process.
	Process *Process `protobuf:"bytes,20,opt,name=process,proto3" json:"process,omitempty"`
	// A link always belongs to a specific map in that process.
	// A map is an instance of a process.
	MapId string `protobuf:"bytes,21,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	// (Optional) Action in the process that resulted in the link's creation.
	// Can be used to filter link search results.
	Action string `protobuf:"bytes,30,opt,name=action,proto3" json:"action,omitempty"`
	// (Optional) Step of the process that results from the action.
	// Can be used to help deserialize link data or filter link search results.
	Step string `protobuf:"bytes,31,opt,name=step,proto3" json:"step,omitempty"`
	// (Optional) Tags that can be used to filter link search results.
	Tags []string `protobuf:"bytes,32,rep,name=tags,proto3" json:"tags,omitempty"`
	// (Optional) Additional metadata needed by your business logic.
	// For backwards compatibility, you should update the link version when the
	// structure of this field changes.
	Data                 []byte   `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkMeta) Reset()         { *m = LinkMeta{} }
func (m *LinkMeta) String() string { return proto.CompactTextString(m) }
func (*LinkMeta) ProtoMessage()    {}
func (*LinkMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainscript_621c374498d1a002, []int{5}
}
func (m *LinkMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkMeta.Unmarshal(m, b)
}
func (m *LinkMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkMeta.Marshal(b, m, deterministic)
}
func (dst *LinkMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkMeta.Merge(dst, src)
}
func (m *LinkMeta) XXX_Size() int {
	return xxx_messageInfo_LinkMeta.Size(m)
}
func (m *LinkMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkMeta.DiscardUnknown(m)
}

var xxx_messageInfo_LinkMeta proto.InternalMessageInfo

func (m *LinkMeta) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *LinkMeta) GetPrevLinkHash() []byte {
	if m != nil {
		return m.PrevLinkHash
	}
	return nil
}

func (m *LinkMeta) GetPriority() float64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *LinkMeta) GetRefs() []*LinkReference {
	if m != nil {
		return m.Refs
	}
	return nil
}

func (m *LinkMeta) GetProcess() *Process {
	if m != nil {
		return m.Process
	}
	return nil
}

func (m *LinkMeta) GetMapId() string {
	if m != nil {
		return m.MapId
	}
	return ""
}

func (m *LinkMeta) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *LinkMeta) GetStep() string {
	if m != nil {
		return m.Step
	}
	return ""
}

func (m *LinkMeta) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *LinkMeta) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// A reference to a link that can be in another process.
type LinkReference struct {
	// Hash of the referenced link.
	LinkHash []byte `protobuf:"bytes,1,opt,name=link_hash,json=linkHash,proto3" json:"link_hash,omitempty"`
	// Process containing the referenced link.
	Process              string   `protobuf:"bytes,10,opt,name=process,proto3" json:"process,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkReference) Reset()         { *m = LinkReference{} }
func (m *LinkReference) String() string { return proto.CompactTextString(m) }
func (*LinkReference) ProtoMessage()    {}
func (*LinkReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainscript_621c374498d1a002, []int{6}
}
func (m *LinkReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkReference.Unmarshal(m, b)
}
func (m *LinkReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkReference.Marshal(b, m, deterministic)
}
func (dst *LinkReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkReference.Merge(dst, src)
}
func (m *LinkReference) XXX_Size() int {
	return xxx_messageInfo_LinkReference.Size(m)
}
func (m *LinkReference) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkReference.DiscardUnknown(m)
}

var xxx_messageInfo_LinkReference proto.InternalMessageInfo

func (m *LinkReference) GetLinkHash() []byte {
	if m != nil {
		return m.LinkHash
	}
	return nil
}

func (m *LinkReference) GetProcess() string {
	if m != nil {
		return m.Process
	}
	return ""
}

// A signature of configurable parts of a link.
// Different signature types and versions are allowed to sign different
// encodings of the data, but we recommend signing a hash of the
// protobuf-encoded bytes.
type Signature struct {
	// Version of the signature format.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Signature algorithm used (for example, "EdDSA").
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// A description of the parts of the links that are signed.
	// This should unambiguously let the verifier recompute the signed payload
	// bytes from the link's content.
	PayloadPath string `protobuf:"bytes,10,opt,name=payload_path,json=payloadPath,proto3" json:"payload_path,omitempty"`
	// Encoded signer's public key.
	// For backwards compatibility, you should update the signature version
	// or the signature type when changing the encoding used.
	PublicKey []byte `protobuf:"bytes,20,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Signature bytes.
	Signature            []byte   `protobuf:"bytes,21,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainscript_621c374498d1a002, []int{7}
}
func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (dst *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(dst, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Signature) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Signature) GetPayloadPath() string {
	if m != nil {
		return m.PayloadPath
	}
	return ""
}

func (m *Signature) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*Segment)(nil), "stratumn.chainscript.Segment")
	proto.RegisterType((*SegmentMeta)(nil), "stratumn.chainscript.SegmentMeta")
	proto.RegisterType((*Evidence)(nil), "stratumn.chainscript.Evidence")
	proto.RegisterType((*Link)(nil), "stratumn.chainscript.Link")
	proto.RegisterType((*Process)(nil), "stratumn.chainscript.Process")
	proto.RegisterType((*LinkMeta)(nil), "stratumn.chainscript.LinkMeta")
	proto.RegisterType((*LinkReference)(nil), "stratumn.chainscript.LinkReference")
	proto.RegisterType((*Signature)(nil), "stratumn.chainscript.Signature")
}

func init() { proto.RegisterFile("chainscript.proto", fileDescriptor_chainscript_621c374498d1a002) }

var fileDescriptor_chainscript_621c374498d1a002 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0xc5, 0x89, 0x1b, 0x5b, 0x23, 0xa5, 0xd0, 0xc5, 0x29, 0x4b, 0xda, 0x24, 0x8e, 0xda, 0x43,
	0x4e, 0x3e, 0x38, 0x94, 0x5c, 0x0a, 0x85, 0x42, 0x4b, 0x43, 0x53, 0x08, 0xea, 0xad, 0x17, 0xb3,
	0x91, 0xc6, 0xd6, 0x62, 0x4b, 0x5a, 0x76, 0xd7, 0x06, 0x7d, 0x4a, 0x3f, 0xa1, 0x3f, 0xd2, 0xef,
	0x2a, 0x3b, 0x5a, 0x59, 0x2e, 0x38, 0xbe, 0xed, 0xbc, 0x9d, 0xd9, 0x99, 0xf7, 0xde, 0x48, 0xf0,
	0x2a, 0xcd, 0x85, 0x2c, 0x4d, 0xaa, 0xa5, 0xb2, 0x13, 0xa5, 0x2b, 0x5b, 0xb1, 0x91, 0xb1, 0x5a,
	0xd8, 0x75, 0x51, 0x4e, 0x76, 0xee, 0x62, 0x05, 0x83, 0x9f, 0xb8, 0x28, 0xb0, 0xb4, 0x6c, 0x02,
	0xfd, 0x95, 0x2c, 0x97, 0xbc, 0x37, 0xee, 0xdd, 0x84, 0xd3, 0xf3, 0xc9, 0xbe, 0xfc, 0xc9, 0x83,
	0x2c, 0x97, 0x09, 0xe5, 0xb1, 0x0f, 0xd0, 0x2f, 0xd0, 0x0a, 0x7e, 0x44, 0xf9, 0xd7, 0xfb, 0xf3,
	0xfd, 0xe3, 0x3f, 0xd0, 0x8a, 0x84, 0xd2, 0xe3, 0x1c, 0xc2, 0x1d, 0x90, 0xbd, 0x81, 0xc0, 0xbd,
	0x36, 0xcb, 0x85, 0xc9, 0xa9, 0x75, 0x94, 0x0c, 0x1d, 0xf0, 0x4d, 0x98, 0x9c, 0x7d, 0x84, 0x00,
	0x37, 0x32, 0xc3, 0x32, 0x45, 0xc3, 0x61, 0x7c, 0x7c, 0x13, 0x4e, 0x2f, 0xf7, 0xf7, 0xf9, 0xe2,
	0xd3, 0x92, 0xae, 0x20, 0x56, 0x30, 0x6c, 0x61, 0xc6, 0x61, 0xb0, 0x41, 0x6d, 0x64, 0x55, 0x52,
	0x93, 0x20, 0x69, 0x43, 0x77, 0xf3, 0x24, 0xd2, 0x25, 0x96, 0x19, 0x87, 0xe6, 0xc6, 0x87, 0xec,
	0x1c, 0x86, 0x4a, 0x57, 0xee, 0x05, 0xcd, 0x43, 0xba, 0xda, 0xc6, 0x6c, 0x04, 0x2f, 0x94, 0xae,
	0xaa, 0x39, 0x1f, 0xd1, 0xc8, 0x4d, 0x10, 0xff, 0xe9, 0x41, 0xdf, 0x29, 0x74, 0xa0, 0x1d, 0x83,
	0x7e, 0x26, 0xac, 0xa0, 0x5e, 0x51, 0x42, 0x67, 0x36, 0xf5, 0x4a, 0x86, 0xa4, 0xe4, 0xe5, 0xf3,
	0xca, 0x77, 0x32, 0xb2, 0x4f, 0x00, 0x46, 0x2e, 0x4a, 0x61, 0xd7, 0x1a, 0x0d, 0x1f, 0x91, 0x36,
	0x57, 0xcf, 0x78, 0xd0, 0xe6, 0x25, 0x3b, 0x25, 0xf1, 0x2d, 0x0c, 0x1e, 0x75, 0x95, 0xa2, 0x31,
	0x6e, 0xa6, 0x52, 0x14, 0xe8, 0x47, 0xa5, 0xb3, 0x23, 0x68, 0xac, 0xb0, 0xe8, 0x45, 0x69, 0x82,
	0xf8, 0xef, 0x11, 0x0c, 0xdb, 0x41, 0x9c, 0x75, 0xe9, 0x4a, 0x62, 0x69, 0x67, 0x32, 0xf3, 0xb5,
	0xc3, 0x06, 0xb8, 0xcf, 0xd8, 0x7b, 0x78, 0xa9, 0x34, 0x6e, 0x66, 0x9d, 0xb9, 0x0d, 0xe3, 0xc8,
	0xa1, 0x0f, 0xad, 0xc1, 0x24, 0xb1, 0xac, 0xb4, 0xb4, 0x35, 0xb1, 0xef, 0x25, 0xdb, 0x98, 0xdd,
	0x41, 0x5f, 0xe3, 0xdc, 0xf0, 0x88, 0xb8, 0xbd, 0x3b, 0xb0, 0x8f, 0x38, 0x47, 0x4d, 0xe6, 0x53,
	0x01, 0xbb, 0x83, 0x81, 0x6a, 0x98, 0x91, 0x3b, 0xe1, 0xf4, 0x62, 0x7f, 0xad, 0xa7, 0x9f, 0xb4,
	0xd9, 0xec, 0x0c, 0x4e, 0x0a, 0xa1, 0x1c, 0x9b, 0xb3, 0x86, 0x74, 0x21, 0xd4, 0x7d, 0xc6, 0x5e,
	0xc3, 0x89, 0x48, 0xad, 0xf3, 0xf2, 0x92, 0x60, 0x1f, 0x39, 0xd9, 0x8c, 0x45, 0xc5, 0xaf, 0x1a,
	0xd9, 0xdc, 0xd9, 0x61, 0x56, 0x2c, 0x0c, 0x1f, 0x8f, 0x8f, 0x1d, 0xe6, 0xce, 0x5b, 0xcb, 0xb3,
	0xce, 0xf2, 0xf8, 0x2b, 0x9c, 0xfe, 0x37, 0xfa, 0xe1, 0xef, 0x80, 0x77, 0x8c, 0xfc, 0x8e, 0xfa,
	0x30, 0xfe, 0xdd, 0x83, 0x60, 0xeb, 0xef, 0xe1, 0xb5, 0xb3, 0xb5, 0x42, 0xfa, 0x58, 0xdd, 0x5c,
	0xb5, 0x42, 0x76, 0x0d, 0x91, 0x12, 0xf5, 0xaa, 0x12, 0xd9, 0x4c, 0x09, 0x9b, 0xfb, 0xa7, 0x43,
	0x8f, 0x3d, 0x0a, 0x9b, 0xb3, 0x0b, 0x00, 0xb5, 0x7e, 0x5a, 0xc9, 0x74, 0xb6, 0xc4, 0xda, 0xef,
	0x7a, 0xd0, 0x20, 0xdf, 0xb1, 0x66, 0x6f, 0x21, 0xd8, 0x6e, 0x14, 0x69, 0x16, 0x25, 0x1d, 0xf0,
	0xf9, 0xf4, 0x57, 0xb8, 0x23, 0xf7, 0xd3, 0x09, 0xfd, 0x87, 0x6e, 0xff, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x90, 0x7a, 0x05, 0x92, 0x9c, 0x04, 0x00, 0x00,
}
