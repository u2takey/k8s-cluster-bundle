// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/apis/bundle/v1alpha1/cluster_bundle.proto

package v1alpha1

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// All the components necessary for deploying (or re-deploying) the cluster. A
// Cluster Bundle can be thought of as a cookie cutter for creating Kubernetes
// clusters.
type ClusterBundle struct {
	// API Version for the Bundle. Should have the form <namespace>/<phase>.
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// The Kubernetes `kind` for this object. Should be 'ClusterBundle'.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// Kubernetes ObjectMeta proto.
	Metadata *ObjectMeta `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Spec for the ClusterBundle, which specifies the intended Kubernetes cluster
	// configuration.
	Spec                 *ClusterBundleSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ClusterBundle) Reset()         { *m = ClusterBundle{} }
func (m *ClusterBundle) String() string { return proto.CompactTextString(m) }
func (*ClusterBundle) ProtoMessage()    {}
func (*ClusterBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b30f666a9d42204, []int{0}
}

func (m *ClusterBundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterBundle.Unmarshal(m, b)
}
func (m *ClusterBundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterBundle.Marshal(b, m, deterministic)
}
func (m *ClusterBundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterBundle.Merge(m, src)
}
func (m *ClusterBundle) XXX_Size() int {
	return xxx_messageInfo_ClusterBundle.Size(m)
}
func (m *ClusterBundle) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterBundle.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterBundle proto.InternalMessageInfo

func (m *ClusterBundle) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *ClusterBundle) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *ClusterBundle) GetMetadata() *ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ClusterBundle) GetSpec() *ClusterBundleSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// A spec object that wraps the Cluster Bundle.
type ClusterBundleSpec struct {
	// Configuration for the nodes.
	NodeConfigs []*NodeConfig `protobuf:"bytes,1,rep,name=node_configs,json=nodeConfigs,proto3" json:"node_configs,omitempty"`
	// Kubernetes objects grouped into cluster components and versioned together.
	// These could be applications or they could be some sort of supporting
	// object collection.
	Components []*ClusterComponent `protobuf:"bytes,2,rep,name=components,proto3" json:"components,omitempty"`
	// Options examples: These are examples of options that can be applied to the
	// cluster bundle in order to specialize it based on user input.
	OptionsExamples      []*ClusterObjectKey `protobuf:"bytes,3,rep,name=options_examples,json=optionsExamples,proto3" json:"options_examples,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ClusterBundleSpec) Reset()         { *m = ClusterBundleSpec{} }
func (m *ClusterBundleSpec) String() string { return proto.CompactTextString(m) }
func (*ClusterBundleSpec) ProtoMessage()    {}
func (*ClusterBundleSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b30f666a9d42204, []int{1}
}

func (m *ClusterBundleSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterBundleSpec.Unmarshal(m, b)
}
func (m *ClusterBundleSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterBundleSpec.Marshal(b, m, deterministic)
}
func (m *ClusterBundleSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterBundleSpec.Merge(m, src)
}
func (m *ClusterBundleSpec) XXX_Size() int {
	return xxx_messageInfo_ClusterBundleSpec.Size(m)
}
func (m *ClusterBundleSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterBundleSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterBundleSpec proto.InternalMessageInfo

func (m *ClusterBundleSpec) GetNodeConfigs() []*NodeConfig {
	if m != nil {
		return m.NodeConfigs
	}
	return nil
}

func (m *ClusterBundleSpec) GetComponents() []*ClusterComponent {
	if m != nil {
		return m.Components
	}
	return nil
}

func (m *ClusterBundleSpec) GetOptionsExamples() []*ClusterObjectKey {
	if m != nil {
		return m.OptionsExamples
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterBundle)(nil), "pkg.apis.bundle.v1alpha1.ClusterBundle")
	proto.RegisterType((*ClusterBundleSpec)(nil), "pkg.apis.bundle.v1alpha1.ClusterBundleSpec")
}

func init() {
	proto.RegisterFile("pkg/apis/bundle/v1alpha1/cluster_bundle.proto", fileDescriptor_6b30f666a9d42204)
}

var fileDescriptor_6b30f666a9d42204 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4e, 0xe3, 0x30,
	0x10, 0x87, 0x95, 0xb6, 0x5a, 0xed, 0xba, 0xbb, 0x5a, 0xf0, 0xc9, 0xea, 0x85, 0xaa, 0xe2, 0x50,
	0x15, 0x9a, 0x50, 0xb8, 0x70, 0x03, 0x35, 0x42, 0x95, 0xca, 0x5f, 0x05, 0xc1, 0x81, 0x4b, 0xe4,
	0x24, 0x6e, 0x6a, 0x92, 0x78, 0xac, 0xd8, 0xa9, 0xe0, 0x09, 0x79, 0x2c, 0x50, 0x9d, 0xb4, 0x2a,
	0x42, 0xa1, 0xdc, 0x9c, 0xf1, 0xef, 0xfb, 0x3c, 0x9a, 0x0c, 0x1a, 0xca, 0x24, 0x76, 0xa8, 0xe4,
	0xca, 0x09, 0x0a, 0x11, 0xa5, 0xcc, 0x59, 0x8c, 0x68, 0x2a, 0xe7, 0x74, 0xe4, 0x84, 0x69, 0xa1,
	0x34, 0xcb, 0xfd, 0xb2, 0x6e, 0xcb, 0x1c, 0x34, 0x60, 0x22, 0x93, 0xd8, 0x5e, 0xc6, 0xed, 0xaa,
	0xbc, 0x8a, 0x77, 0x0e, 0x6b, 0x45, 0xe5, 0xb7, 0x1f, 0x42, 0x96, 0x81, 0x28, 0x3d, 0x9d, 0xa3,
	0xad, 0xcf, 0x86, 0x90, 0x49, 0x10, 0x4c, 0xe8, 0x8a, 0x18, 0xd4, 0x12, 0x02, 0xa2, 0xa5, 0x5d,
	0xcc, 0x78, 0xbc, 0x35, 0x0b, 0xc1, 0x33, 0x0b, 0xb5, 0x9f, 0x31, 0x4d, 0xcb, 0x6c, 0xef, 0xcd,
	0x42, 0xff, 0xdc, 0xf2, 0xcd, 0xb1, 0x09, 0xe3, 0x3d, 0xd4, 0xa6, 0x92, 0xfb, 0x0b, 0x96, 0x2b,
	0x0e, 0x82, 0x58, 0x5d, 0xab, 0xff, 0xc7, 0x43, 0x54, 0xf2, 0xc7, 0xb2, 0x82, 0x31, 0x6a, 0x25,
	0x5c, 0x44, 0xa4, 0x61, 0x6e, 0xcc, 0x19, 0x9f, 0xa3, 0xdf, 0x4b, 0x69, 0x44, 0x35, 0x25, 0xcd,
	0xae, 0xd5, 0x6f, 0x1f, 0xef, 0xdb, 0x75, 0xb3, 0xb2, 0x6f, 0x4d, 0x17, 0xd7, 0x4c, 0x53, 0x6f,
	0x4d, 0xe1, 0x33, 0xd4, 0x52, 0x92, 0x85, 0xa4, 0x65, 0xe8, 0x83, 0x7a, 0xfa, 0x53, 0xb7, 0xf7,
	0x92, 0x85, 0x9e, 0x01, 0x7b, 0xef, 0x16, 0xda, 0xfd, 0x72, 0x87, 0x27, 0xe8, 0xef, 0xc6, 0x80,
	0x14, 0xb1, 0xba, 0xcd, 0xef, 0x9b, 0xbb, 0x81, 0x88, 0xb9, 0x26, 0xec, 0xb5, 0xc5, 0xfa, 0xac,
	0xf0, 0x14, 0xa1, 0xf5, 0x3f, 0x51, 0xa4, 0x61, 0x34, 0x83, 0xad, 0x5d, 0xba, 0x2b, 0xc4, 0xdb,
	0xa0, 0xf1, 0x03, 0xda, 0x01, 0xa9, 0x39, 0x08, 0xe5, 0xb3, 0x17, 0x9a, 0xc9, 0x94, 0x29, 0xd2,
	0xfc, 0xa1, 0xb1, 0x1c, 0xde, 0x25, 0x7b, 0xf5, 0xfe, 0x57, 0x8e, 0x8b, 0x4a, 0x31, 0xbe, 0x7a,
	0x9a, 0xc6, 0x5c, 0xcf, 0x8b, 0xc0, 0x0e, 0x21, 0x73, 0x26, 0x00, 0x71, 0xca, 0xdc, 0x14, 0x8a,
	0xe8, 0x2e, 0xa5, 0x7a, 0x06, 0x79, 0xe6, 0x24, 0xa7, 0x6a, 0x58, 0x6d, 0xd8, 0xb0, 0xda, 0x8d,
	0xba, 0x5d, 0x09, 0x7e, 0x99, 0x05, 0x39, 0xf9, 0x08, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x32, 0x64,
	0xff, 0x23, 0x03, 0x00, 0x00,
}