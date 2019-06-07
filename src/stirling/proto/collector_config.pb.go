// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/stirling/proto/collector_config.proto

package stirlingpb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	io "io"
	math "math"
	proto1 "pixielabs.ai/pixielabs/src/shared/types/proto"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Element struct {
	Name  string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type  proto1.DataType    `protobuf:"varint,2,opt,name=type,proto3,enum=pl.types.DataType" json:"type,omitempty"`
	Ptype proto1.PatternType `protobuf:"varint,3,opt,name=ptype,proto3,enum=pl.types.PatternType" json:"ptype,omitempty"`
}

func (m *Element) Reset()      { *m = Element{} }
func (*Element) ProtoMessage() {}
func (*Element) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0c6d532c5976665, []int{0}
}
func (m *Element) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Element) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Element.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Element) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Element.Merge(m, src)
}
func (m *Element) XXX_Size() int {
	return m.Size()
}
func (m *Element) XXX_DiscardUnknown() {
	xxx_messageInfo_Element.DiscardUnknown(m)
}

var xxx_messageInfo_Element proto.InternalMessageInfo

func (m *Element) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Element) GetType() proto1.DataType {
	if m != nil {
		return m.Type
	}
	return proto1.DATA_TYPE_UNKNOWN
}

func (m *Element) GetPtype() proto1.PatternType {
	if m != nil {
		return m.Ptype
	}
	return proto1.UNSPECIFIED
}

type InfoClass struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   uint64            `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Elements             []*Element        `protobuf:"bytes,3,rep,name=elements,proto3" json:"elements,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Subscribed           bool              `protobuf:"varint,5,opt,name=subscribed,proto3" json:"subscribed,omitempty"`
	SamplingPeriodMillis uint32            `protobuf:"varint,6,opt,name=sampling_period_millis,json=samplingPeriodMillis,proto3" json:"sampling_period_millis,omitempty"`
	PushPeriodMillis     uint32            `protobuf:"varint,7,opt,name=push_period_millis,json=pushPeriodMillis,proto3" json:"push_period_millis,omitempty"`
}

func (m *InfoClass) Reset()      { *m = InfoClass{} }
func (*InfoClass) ProtoMessage() {}
func (*InfoClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0c6d532c5976665, []int{1}
}
func (m *InfoClass) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InfoClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InfoClass.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InfoClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoClass.Merge(m, src)
}
func (m *InfoClass) XXX_Size() int {
	return m.Size()
}
func (m *InfoClass) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoClass.DiscardUnknown(m)
}

var xxx_messageInfo_InfoClass proto.InternalMessageInfo

func (m *InfoClass) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InfoClass) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InfoClass) GetElements() []*Element {
	if m != nil {
		return m.Elements
	}
	return nil
}

func (m *InfoClass) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *InfoClass) GetSubscribed() bool {
	if m != nil {
		return m.Subscribed
	}
	return false
}

func (m *InfoClass) GetSamplingPeriodMillis() uint32 {
	if m != nil {
		return m.SamplingPeriodMillis
	}
	return 0
}

func (m *InfoClass) GetPushPeriodMillis() uint32 {
	if m != nil {
		return m.PushPeriodMillis
	}
	return 0
}

type Publish struct {
	PublishedInfoClasses []*InfoClass `protobuf:"bytes,1,rep,name=published_info_classes,json=publishedInfoClasses,proto3" json:"published_info_classes,omitempty"`
}

func (m *Publish) Reset()      { *m = Publish{} }
func (*Publish) ProtoMessage() {}
func (*Publish) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0c6d532c5976665, []int{2}
}
func (m *Publish) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Publish) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Publish.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Publish) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Publish.Merge(m, src)
}
func (m *Publish) XXX_Size() int {
	return m.Size()
}
func (m *Publish) XXX_DiscardUnknown() {
	xxx_messageInfo_Publish.DiscardUnknown(m)
}

var xxx_messageInfo_Publish proto.InternalMessageInfo

func (m *Publish) GetPublishedInfoClasses() []*InfoClass {
	if m != nil {
		return m.PublishedInfoClasses
	}
	return nil
}

type Subscribe struct {
	SubscribedInfoClasses []*InfoClass `protobuf:"bytes,1,rep,name=subscribed_info_classes,json=subscribedInfoClasses,proto3" json:"subscribed_info_classes,omitempty"`
}

func (m *Subscribe) Reset()      { *m = Subscribe{} }
func (*Subscribe) ProtoMessage() {}
func (*Subscribe) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0c6d532c5976665, []int{3}
}
func (m *Subscribe) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Subscribe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Subscribe.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Subscribe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscribe.Merge(m, src)
}
func (m *Subscribe) XXX_Size() int {
	return m.Size()
}
func (m *Subscribe) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscribe.DiscardUnknown(m)
}

var xxx_messageInfo_Subscribe proto.InternalMessageInfo

func (m *Subscribe) GetSubscribedInfoClasses() []*InfoClass {
	if m != nil {
		return m.SubscribedInfoClasses
	}
	return nil
}

func init() {
	proto.RegisterType((*Element)(nil), "pl.stirling.stirlingpb.Element")
	proto.RegisterType((*InfoClass)(nil), "pl.stirling.stirlingpb.InfoClass")
	proto.RegisterMapType((map[string]string)(nil), "pl.stirling.stirlingpb.InfoClass.MetadataEntry")
	proto.RegisterType((*Publish)(nil), "pl.stirling.stirlingpb.Publish")
	proto.RegisterType((*Subscribe)(nil), "pl.stirling.stirlingpb.Subscribe")
}

func init() {
	proto.RegisterFile("src/stirling/proto/collector_config.proto", fileDescriptor_f0c6d532c5976665)
}

var fileDescriptor_f0c6d532c5976665 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4f, 0x8b, 0xd3, 0x40,
	0x14, 0xcf, 0xb4, 0xdd, 0x6d, 0xfb, 0x64, 0x97, 0x65, 0xe8, 0xd6, 0xb0, 0x87, 0xb1, 0xf6, 0x20,
	0x15, 0x25, 0x81, 0xd5, 0x83, 0xb8, 0x17, 0x51, 0xf7, 0x20, 0xb2, 0x50, 0xa2, 0x20, 0x7a, 0x09,
	0x93, 0x64, 0xba, 0x1d, 0x9c, 0x24, 0xc3, 0xcc, 0x54, 0xe8, 0xcd, 0x8f, 0xe0, 0xc7, 0xf0, 0xa3,
	0x78, 0xec, 0x71, 0x8f, 0x36, 0xbd, 0x78, 0xdc, 0xb3, 0x27, 0xc9, 0xa4, 0x4d, 0xab, 0xac, 0x08,
	0xde, 0x5e, 0xde, 0xfb, 0xfd, 0x79, 0xef, 0x17, 0x06, 0xee, 0x6b, 0x15, 0xfb, 0xda, 0x70, 0x25,
	0x78, 0x76, 0xe9, 0x4b, 0x95, 0x9b, 0xdc, 0x8f, 0x73, 0x21, 0x58, 0x6c, 0x72, 0x15, 0xc6, 0x79,
	0x36, 0xe1, 0x97, 0x9e, 0x6d, 0xe3, 0xbe, 0x14, 0xde, 0x06, 0x59, 0x17, 0x32, 0x3a, 0x19, 0x5a,
	0x89, 0x29, 0x55, 0x2c, 0xf1, 0xcd, 0x5c, 0x32, 0xbd, 0x96, 0xb1, 0x75, 0xc5, 0x1d, 0x2a, 0x68,
	0x9f, 0x0b, 0x96, 0xb2, 0xcc, 0x60, 0x0c, 0xad, 0x8c, 0xa6, 0xcc, 0x45, 0x03, 0x34, 0xea, 0x06,
	0xb6, 0xc6, 0xf7, 0xa0, 0x55, 0xa2, 0xdd, 0xc6, 0x00, 0x8d, 0x0e, 0x4f, 0xb1, 0x27, 0x85, 0x57,
	0xb1, 0x5f, 0x52, 0x43, 0xdf, 0xce, 0x25, 0x0b, 0xec, 0x1c, 0x3f, 0x80, 0x3d, 0x69, 0x81, 0x4d,
	0x0b, 0x3c, 0xde, 0x02, 0xc7, 0xd4, 0x18, 0xa6, 0x32, 0x8b, 0xad, 0x30, 0xc3, 0x9f, 0x0d, 0xe8,
	0xbe, 0xca, 0x26, 0xf9, 0x0b, 0x41, 0xb5, 0xbe, 0xd1, 0xf6, 0x10, 0x1a, 0x3c, 0xb1, 0xa6, 0xad,
	0xa0, 0xc1, 0x13, 0x7c, 0x06, 0x1d, 0x56, 0x6d, 0xa9, 0xdd, 0xe6, 0xa0, 0x39, 0xba, 0x75, 0x7a,
	0xc7, 0xbb, 0xf9, 0x68, 0x6f, 0x7d, 0x4d, 0x50, 0x13, 0xf0, 0x6b, 0xe8, 0xa4, 0xcc, 0xd0, 0x84,
	0x1a, 0xea, 0xb6, 0x2c, 0xd9, 0xff, 0x1b, 0xb9, 0xde, 0xca, 0xbb, 0x58, 0x33, 0xce, 0x33, 0xa3,
	0xe6, 0x41, 0x2d, 0x80, 0x09, 0x80, 0x9e, 0x45, 0x3a, 0x56, 0x3c, 0x62, 0x89, 0xbb, 0x37, 0x40,
	0xa3, 0x4e, 0xb0, 0xd3, 0xc1, 0x8f, 0xa1, 0xaf, 0x69, 0x2a, 0x4b, 0xbd, 0x50, 0x32, 0xc5, 0xf3,
	0x24, 0x4c, 0xb9, 0x10, 0x5c, 0xbb, 0xfb, 0x03, 0x34, 0x3a, 0x08, 0x7a, 0x9b, 0xe9, 0xd8, 0x0e,
	0x2f, 0xec, 0x0c, 0x3f, 0x04, 0x2c, 0x67, 0x7a, 0xfa, 0x07, 0xa3, 0x6d, 0x19, 0x47, 0xe5, 0x64,
	0x17, 0x7d, 0x72, 0x06, 0x07, 0xbf, 0xad, 0x87, 0x8f, 0xa0, 0xf9, 0x91, 0xcd, 0xd7, 0x09, 0x96,
	0x25, 0xee, 0xc1, 0xde, 0x27, 0x2a, 0x66, 0xd5, 0x8f, 0xeb, 0x06, 0xd5, 0xc7, 0xd3, 0xc6, 0x13,
	0x34, 0x8c, 0xa0, 0x3d, 0x9e, 0x45, 0x82, 0xeb, 0x29, 0x7e, 0x07, 0x7d, 0x59, 0x95, 0x2c, 0x09,
	0x79, 0x36, 0xc9, 0xc3, 0xb8, 0xbc, 0x9d, 0x69, 0x17, 0xd9, 0x98, 0xee, 0xfe, 0x33, 0xa6, 0xa0,
	0x57, 0x0b, 0xd4, 0x3d, 0xa6, 0x87, 0x13, 0xe8, 0xbe, 0xd9, 0x44, 0x82, 0xdf, 0xc3, 0xed, 0x6d,
	0x3e, 0xff, 0x69, 0x73, 0xbc, 0x55, 0xd8, 0xf1, 0x79, 0xfe, 0x6c, 0xb1, 0x24, 0xce, 0xd5, 0x92,
	0x38, 0xd7, 0x4b, 0x82, 0x3e, 0x17, 0x04, 0x7d, 0x2d, 0x08, 0xfa, 0x56, 0x10, 0xb4, 0x28, 0x08,
	0xfa, 0x5e, 0x10, 0xf4, 0xa3, 0x20, 0xce, 0x75, 0x41, 0xd0, 0x97, 0x15, 0x71, 0x16, 0x2b, 0xe2,
	0x5c, 0xad, 0x88, 0xf3, 0x01, 0xb6, 0x16, 0xd1, 0xbe, 0x7d, 0x05, 0x8f, 0x7e, 0x05, 0x00, 0x00,
	0xff, 0xff, 0xea, 0xb2, 0x26, 0xf1, 0x6e, 0x03, 0x00, 0x00,
}

func (this *Element) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Element)
	if !ok {
		that2, ok := that.(Element)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Ptype != that1.Ptype {
		return false
	}
	return true
}
func (this *InfoClass) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*InfoClass)
	if !ok {
		that2, ok := that.(InfoClass)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if len(this.Elements) != len(that1.Elements) {
		return false
	}
	for i := range this.Elements {
		if !this.Elements[i].Equal(that1.Elements[i]) {
			return false
		}
	}
	if len(this.Metadata) != len(that1.Metadata) {
		return false
	}
	for i := range this.Metadata {
		if this.Metadata[i] != that1.Metadata[i] {
			return false
		}
	}
	if this.Subscribed != that1.Subscribed {
		return false
	}
	if this.SamplingPeriodMillis != that1.SamplingPeriodMillis {
		return false
	}
	if this.PushPeriodMillis != that1.PushPeriodMillis {
		return false
	}
	return true
}
func (this *Publish) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Publish)
	if !ok {
		that2, ok := that.(Publish)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.PublishedInfoClasses) != len(that1.PublishedInfoClasses) {
		return false
	}
	for i := range this.PublishedInfoClasses {
		if !this.PublishedInfoClasses[i].Equal(that1.PublishedInfoClasses[i]) {
			return false
		}
	}
	return true
}
func (this *Subscribe) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Subscribe)
	if !ok {
		that2, ok := that.(Subscribe)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.SubscribedInfoClasses) != len(that1.SubscribedInfoClasses) {
		return false
	}
	for i := range this.SubscribedInfoClasses {
		if !this.SubscribedInfoClasses[i].Equal(that1.SubscribedInfoClasses[i]) {
			return false
		}
	}
	return true
}
func (this *Element) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&stirlingpb.Element{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Ptype: "+fmt.Sprintf("%#v", this.Ptype)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InfoClass) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&stirlingpb.InfoClass{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	if this.Elements != nil {
		s = append(s, "Elements: "+fmt.Sprintf("%#v", this.Elements)+",\n")
	}
	keysForMetadata := make([]string, 0, len(this.Metadata))
	for k, _ := range this.Metadata {
		keysForMetadata = append(keysForMetadata, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForMetadata)
	mapStringForMetadata := "map[string]string{"
	for _, k := range keysForMetadata {
		mapStringForMetadata += fmt.Sprintf("%#v: %#v,", k, this.Metadata[k])
	}
	mapStringForMetadata += "}"
	if this.Metadata != nil {
		s = append(s, "Metadata: "+mapStringForMetadata+",\n")
	}
	s = append(s, "Subscribed: "+fmt.Sprintf("%#v", this.Subscribed)+",\n")
	s = append(s, "SamplingPeriodMillis: "+fmt.Sprintf("%#v", this.SamplingPeriodMillis)+",\n")
	s = append(s, "PushPeriodMillis: "+fmt.Sprintf("%#v", this.PushPeriodMillis)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Publish) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&stirlingpb.Publish{")
	if this.PublishedInfoClasses != nil {
		s = append(s, "PublishedInfoClasses: "+fmt.Sprintf("%#v", this.PublishedInfoClasses)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Subscribe) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&stirlingpb.Subscribe{")
	if this.SubscribedInfoClasses != nil {
		s = append(s, "SubscribedInfoClasses: "+fmt.Sprintf("%#v", this.SubscribedInfoClasses)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCollectorConfig(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Element) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Element) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCollectorConfig(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Type != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCollectorConfig(dAtA, i, uint64(m.Type))
	}
	if m.Ptype != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintCollectorConfig(dAtA, i, uint64(m.Ptype))
	}
	return i, nil
}

func (m *InfoClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InfoClass) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCollectorConfig(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Id != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCollectorConfig(dAtA, i, uint64(m.Id))
	}
	if len(m.Elements) > 0 {
		for _, msg := range m.Elements {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintCollectorConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Metadata) > 0 {
		for k, _ := range m.Metadata {
			dAtA[i] = 0x22
			i++
			v := m.Metadata[k]
			mapSize := 1 + len(k) + sovCollectorConfig(uint64(len(k))) + 1 + len(v) + sovCollectorConfig(uint64(len(v)))
			i = encodeVarintCollectorConfig(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintCollectorConfig(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintCollectorConfig(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.Subscribed {
		dAtA[i] = 0x28
		i++
		if m.Subscribed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.SamplingPeriodMillis != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintCollectorConfig(dAtA, i, uint64(m.SamplingPeriodMillis))
	}
	if m.PushPeriodMillis != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintCollectorConfig(dAtA, i, uint64(m.PushPeriodMillis))
	}
	return i, nil
}

func (m *Publish) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Publish) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PublishedInfoClasses) > 0 {
		for _, msg := range m.PublishedInfoClasses {
			dAtA[i] = 0xa
			i++
			i = encodeVarintCollectorConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Subscribe) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Subscribe) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.SubscribedInfoClasses) > 0 {
		for _, msg := range m.SubscribedInfoClasses {
			dAtA[i] = 0xa
			i++
			i = encodeVarintCollectorConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintCollectorConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Element) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCollectorConfig(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovCollectorConfig(uint64(m.Type))
	}
	if m.Ptype != 0 {
		n += 1 + sovCollectorConfig(uint64(m.Ptype))
	}
	return n
}

func (m *InfoClass) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCollectorConfig(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovCollectorConfig(uint64(m.Id))
	}
	if len(m.Elements) > 0 {
		for _, e := range m.Elements {
			l = e.Size()
			n += 1 + l + sovCollectorConfig(uint64(l))
		}
	}
	if len(m.Metadata) > 0 {
		for k, v := range m.Metadata {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovCollectorConfig(uint64(len(k))) + 1 + len(v) + sovCollectorConfig(uint64(len(v)))
			n += mapEntrySize + 1 + sovCollectorConfig(uint64(mapEntrySize))
		}
	}
	if m.Subscribed {
		n += 2
	}
	if m.SamplingPeriodMillis != 0 {
		n += 1 + sovCollectorConfig(uint64(m.SamplingPeriodMillis))
	}
	if m.PushPeriodMillis != 0 {
		n += 1 + sovCollectorConfig(uint64(m.PushPeriodMillis))
	}
	return n
}

func (m *Publish) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PublishedInfoClasses) > 0 {
		for _, e := range m.PublishedInfoClasses {
			l = e.Size()
			n += 1 + l + sovCollectorConfig(uint64(l))
		}
	}
	return n
}

func (m *Subscribe) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SubscribedInfoClasses) > 0 {
		for _, e := range m.SubscribedInfoClasses {
			l = e.Size()
			n += 1 + l + sovCollectorConfig(uint64(l))
		}
	}
	return n
}

func sovCollectorConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCollectorConfig(x uint64) (n int) {
	return sovCollectorConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Element) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Element{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Ptype:` + fmt.Sprintf("%v", this.Ptype) + `,`,
		`}`,
	}, "")
	return s
}
func (this *InfoClass) String() string {
	if this == nil {
		return "nil"
	}
	keysForMetadata := make([]string, 0, len(this.Metadata))
	for k, _ := range this.Metadata {
		keysForMetadata = append(keysForMetadata, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForMetadata)
	mapStringForMetadata := "map[string]string{"
	for _, k := range keysForMetadata {
		mapStringForMetadata += fmt.Sprintf("%v: %v,", k, this.Metadata[k])
	}
	mapStringForMetadata += "}"
	s := strings.Join([]string{`&InfoClass{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`Elements:` + strings.Replace(fmt.Sprintf("%v", this.Elements), "Element", "Element", 1) + `,`,
		`Metadata:` + mapStringForMetadata + `,`,
		`Subscribed:` + fmt.Sprintf("%v", this.Subscribed) + `,`,
		`SamplingPeriodMillis:` + fmt.Sprintf("%v", this.SamplingPeriodMillis) + `,`,
		`PushPeriodMillis:` + fmt.Sprintf("%v", this.PushPeriodMillis) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Publish) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Publish{`,
		`PublishedInfoClasses:` + strings.Replace(fmt.Sprintf("%v", this.PublishedInfoClasses), "InfoClass", "InfoClass", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Subscribe) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Subscribe{`,
		`SubscribedInfoClasses:` + strings.Replace(fmt.Sprintf("%v", this.SubscribedInfoClasses), "InfoClass", "InfoClass", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCollectorConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Element) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollectorConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Element: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Element: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectorConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectorConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= proto1.DataType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ptype", wireType)
			}
			m.Ptype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectorConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ptype |= proto1.PatternType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCollectorConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InfoClass) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollectorConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InfoClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InfoClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectorConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectorConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Elements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectorConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Elements = append(m.Elements, &Element{})
			if err := m.Elements[len(m.Elements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectorConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCollectorConfig
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCollectorConfig
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthCollectorConfig
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthCollectorConfig
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCollectorConfig
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthCollectorConfig
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthCollectorConfig
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipCollectorConfig(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthCollectorConfig
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Metadata[mapkey] = mapvalue
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscribed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectorConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Subscribed = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SamplingPeriodMillis", wireType)
			}
			m.SamplingPeriodMillis = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectorConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SamplingPeriodMillis |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PushPeriodMillis", wireType)
			}
			m.PushPeriodMillis = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectorConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PushPeriodMillis |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCollectorConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Publish) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollectorConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Publish: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Publish: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublishedInfoClasses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectorConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublishedInfoClasses = append(m.PublishedInfoClasses, &InfoClass{})
			if err := m.PublishedInfoClasses[len(m.PublishedInfoClasses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollectorConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Subscribe) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollectorConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Subscribe: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Subscribe: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscribedInfoClasses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectorConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubscribedInfoClasses = append(m.SubscribedInfoClasses, &InfoClass{})
			if err := m.SubscribedInfoClasses[len(m.SubscribedInfoClasses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollectorConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCollectorConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCollectorConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCollectorConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCollectorConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCollectorConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCollectorConfig
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCollectorConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCollectorConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCollectorConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCollectorConfig
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCollectorConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCollectorConfig   = fmt.Errorf("proto: integer overflow")
)
