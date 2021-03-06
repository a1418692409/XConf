// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/config.proto

package config

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

type AppRequest struct {
	AppName              string   `protobuf:"bytes,4,opt,name=appName,proto3" json:"appName"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppRequest) Reset()         { *m = AppRequest{} }
func (m *AppRequest) String() string { return proto.CompactTextString(m) }
func (*AppRequest) ProtoMessage()    {}
func (*AppRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{0}
}

func (m *AppRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppRequest.Unmarshal(m, b)
}
func (m *AppRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppRequest.Marshal(b, m, deterministic)
}
func (m *AppRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppRequest.Merge(m, src)
}
func (m *AppRequest) XXX_Size() int {
	return xxx_messageInfo_AppRequest.Size(m)
}
func (m *AppRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppRequest proto.InternalMessageInfo

func (m *AppRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *AppRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type AppResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt"`
	UpdatedAt            int64    `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt"`
	AppName              string   `protobuf:"bytes,4,opt,name=appName,proto3" json:"appName"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppResponse) Reset()         { *m = AppResponse{} }
func (m *AppResponse) String() string { return proto.CompactTextString(m) }
func (*AppResponse) ProtoMessage()    {}
func (*AppResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{1}
}

func (m *AppResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppResponse.Unmarshal(m, b)
}
func (m *AppResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppResponse.Marshal(b, m, deterministic)
}
func (m *AppResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppResponse.Merge(m, src)
}
func (m *AppResponse) XXX_Size() int {
	return xxx_messageInfo_AppResponse.Size(m)
}
func (m *AppResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AppResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AppResponse proto.InternalMessageInfo

func (m *AppResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AppResponse) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *AppResponse) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *AppResponse) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *AppResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type AppsResponse struct {
	Apps                 []*AppResponse `protobuf:"bytes,1,rep,name=apps,proto3" json:"apps"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AppsResponse) Reset()         { *m = AppsResponse{} }
func (m *AppsResponse) String() string { return proto.CompactTextString(m) }
func (*AppsResponse) ProtoMessage()    {}
func (*AppsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{2}
}

func (m *AppsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppsResponse.Unmarshal(m, b)
}
func (m *AppsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppsResponse.Marshal(b, m, deterministic)
}
func (m *AppsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppsResponse.Merge(m, src)
}
func (m *AppsResponse) XXX_Size() int {
	return xxx_messageInfo_AppsResponse.Size(m)
}
func (m *AppsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AppsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AppsResponse proto.InternalMessageInfo

func (m *AppsResponse) GetApps() []*AppResponse {
	if m != nil {
		return m.Apps
	}
	return nil
}

type ClusterRequest struct {
	AppName              string   `protobuf:"bytes,4,opt,name=appName,proto3" json:"appName"`
	ClusterName          string   `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName"`
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterRequest) Reset()         { *m = ClusterRequest{} }
func (m *ClusterRequest) String() string { return proto.CompactTextString(m) }
func (*ClusterRequest) ProtoMessage()    {}
func (*ClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{3}
}

func (m *ClusterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterRequest.Unmarshal(m, b)
}
func (m *ClusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterRequest.Marshal(b, m, deterministic)
}
func (m *ClusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterRequest.Merge(m, src)
}
func (m *ClusterRequest) XXX_Size() int {
	return xxx_messageInfo_ClusterRequest.Size(m)
}
func (m *ClusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterRequest proto.InternalMessageInfo

func (m *ClusterRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *ClusterRequest) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ClusterResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt"`
	UpdatedAt            int64    `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt"`
	AppName              string   `protobuf:"bytes,4,opt,name=appName,proto3" json:"appName"`
	ClusterName          string   `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName"`
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterResponse) Reset()         { *m = ClusterResponse{} }
func (m *ClusterResponse) String() string { return proto.CompactTextString(m) }
func (*ClusterResponse) ProtoMessage()    {}
func (*ClusterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{4}
}

func (m *ClusterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterResponse.Unmarshal(m, b)
}
func (m *ClusterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterResponse.Marshal(b, m, deterministic)
}
func (m *ClusterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterResponse.Merge(m, src)
}
func (m *ClusterResponse) XXX_Size() int {
	return xxx_messageInfo_ClusterResponse.Size(m)
}
func (m *ClusterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterResponse proto.InternalMessageInfo

func (m *ClusterResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ClusterResponse) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *ClusterResponse) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *ClusterResponse) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *ClusterResponse) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ClustersResponse struct {
	Clusters             []*ClusterResponse `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ClustersResponse) Reset()         { *m = ClustersResponse{} }
func (m *ClustersResponse) String() string { return proto.CompactTextString(m) }
func (*ClustersResponse) ProtoMessage()    {}
func (*ClustersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{5}
}

func (m *ClustersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClustersResponse.Unmarshal(m, b)
}
func (m *ClustersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClustersResponse.Marshal(b, m, deterministic)
}
func (m *ClustersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClustersResponse.Merge(m, src)
}
func (m *ClustersResponse) XXX_Size() int {
	return xxx_messageInfo_ClustersResponse.Size(m)
}
func (m *ClustersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClustersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClustersResponse proto.InternalMessageInfo

func (m *ClustersResponse) GetClusters() []*ClusterResponse {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type NamespaceRequest struct {
	AppName              string   `protobuf:"bytes,4,opt,name=appName,proto3" json:"appName"`
	ClusterName          string   `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName"`
	NamespaceName        string   `protobuf:"bytes,6,opt,name=namespaceName,proto3" json:"namespaceName"`
	Format               string   `protobuf:"bytes,7,opt,name=format,proto3" json:"format"`
	Description          string   `protobuf:"bytes,10,opt,name=description,proto3" json:"description"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamespaceRequest) Reset()         { *m = NamespaceRequest{} }
func (m *NamespaceRequest) String() string { return proto.CompactTextString(m) }
func (*NamespaceRequest) ProtoMessage()    {}
func (*NamespaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{6}
}

func (m *NamespaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamespaceRequest.Unmarshal(m, b)
}
func (m *NamespaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamespaceRequest.Marshal(b, m, deterministic)
}
func (m *NamespaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespaceRequest.Merge(m, src)
}
func (m *NamespaceRequest) XXX_Size() int {
	return xxx_messageInfo_NamespaceRequest.Size(m)
}
func (m *NamespaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NamespaceRequest proto.InternalMessageInfo

func (m *NamespaceRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *NamespaceRequest) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *NamespaceRequest) GetNamespaceName() string {
	if m != nil {
		return m.NamespaceName
	}
	return ""
}

func (m *NamespaceRequest) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *NamespaceRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type NamespaceResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt"`
	UpdatedAt            int64    `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt"`
	AppName              string   `protobuf:"bytes,4,opt,name=appName,proto3" json:"appName"`
	ClusterName          string   `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName"`
	NamespaceName        string   `protobuf:"bytes,6,opt,name=namespaceName,proto3" json:"namespaceName"`
	Format               string   `protobuf:"bytes,7,opt,name=format,proto3" json:"format"`
	Value                string   `protobuf:"bytes,8,opt,name=value,proto3" json:"value"`
	Released             bool     `protobuf:"varint,9,opt,name=released,proto3" json:"released"`
	EditValue            string   `protobuf:"bytes,10,opt,name=editValue,proto3" json:"editValue"`
	Description          string   `protobuf:"bytes,11,opt,name=description,proto3" json:"description"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamespaceResponse) Reset()         { *m = NamespaceResponse{} }
func (m *NamespaceResponse) String() string { return proto.CompactTextString(m) }
func (*NamespaceResponse) ProtoMessage()    {}
func (*NamespaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{7}
}

func (m *NamespaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamespaceResponse.Unmarshal(m, b)
}
func (m *NamespaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamespaceResponse.Marshal(b, m, deterministic)
}
func (m *NamespaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespaceResponse.Merge(m, src)
}
func (m *NamespaceResponse) XXX_Size() int {
	return xxx_messageInfo_NamespaceResponse.Size(m)
}
func (m *NamespaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NamespaceResponse proto.InternalMessageInfo

func (m *NamespaceResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NamespaceResponse) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *NamespaceResponse) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *NamespaceResponse) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *NamespaceResponse) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *NamespaceResponse) GetNamespaceName() string {
	if m != nil {
		return m.NamespaceName
	}
	return ""
}

func (m *NamespaceResponse) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *NamespaceResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *NamespaceResponse) GetReleased() bool {
	if m != nil {
		return m.Released
	}
	return false
}

func (m *NamespaceResponse) GetEditValue() string {
	if m != nil {
		return m.EditValue
	}
	return ""
}

func (m *NamespaceResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type NamespacesResponse struct {
	Namespaces           []*NamespaceResponse `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NamespacesResponse) Reset()         { *m = NamespacesResponse{} }
func (m *NamespacesResponse) String() string { return proto.CompactTextString(m) }
func (*NamespacesResponse) ProtoMessage()    {}
func (*NamespacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{8}
}

func (m *NamespacesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamespacesResponse.Unmarshal(m, b)
}
func (m *NamespacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamespacesResponse.Marshal(b, m, deterministic)
}
func (m *NamespacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespacesResponse.Merge(m, src)
}
func (m *NamespacesResponse) XXX_Size() int {
	return xxx_messageInfo_NamespacesResponse.Size(m)
}
func (m *NamespacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NamespacesResponse proto.InternalMessageInfo

func (m *NamespacesResponse) GetNamespaces() []*NamespaceResponse {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

type UpdateConfigRequest struct {
	AppName              string   `protobuf:"bytes,4,opt,name=appName,proto3" json:"appName"`
	ClusterName          string   `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName"`
	NamespaceName        string   `protobuf:"bytes,6,opt,name=namespaceName,proto3" json:"namespaceName"`
	Value                string   `protobuf:"bytes,8,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateConfigRequest) Reset()         { *m = UpdateConfigRequest{} }
func (m *UpdateConfigRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateConfigRequest) ProtoMessage()    {}
func (*UpdateConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{9}
}

func (m *UpdateConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateConfigRequest.Unmarshal(m, b)
}
func (m *UpdateConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateConfigRequest.Marshal(b, m, deterministic)
}
func (m *UpdateConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateConfigRequest.Merge(m, src)
}
func (m *UpdateConfigRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateConfigRequest.Size(m)
}
func (m *UpdateConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateConfigRequest proto.InternalMessageInfo

func (m *UpdateConfigRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *UpdateConfigRequest) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *UpdateConfigRequest) GetNamespaceName() string {
	if m != nil {
		return m.NamespaceName
	}
	return ""
}

func (m *UpdateConfigRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type QueryConfigRequest struct {
	AppName              string   `protobuf:"bytes,4,opt,name=appName,proto3" json:"appName"`
	ClusterName          string   `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName"`
	NamespaceName        string   `protobuf:"bytes,6,opt,name=namespaceName,proto3" json:"namespaceName"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryConfigRequest) Reset()         { *m = QueryConfigRequest{} }
func (m *QueryConfigRequest) String() string { return proto.CompactTextString(m) }
func (*QueryConfigRequest) ProtoMessage()    {}
func (*QueryConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{10}
}

func (m *QueryConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryConfigRequest.Unmarshal(m, b)
}
func (m *QueryConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryConfigRequest.Marshal(b, m, deterministic)
}
func (m *QueryConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryConfigRequest.Merge(m, src)
}
func (m *QueryConfigRequest) XXX_Size() int {
	return xxx_messageInfo_QueryConfigRequest.Size(m)
}
func (m *QueryConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryConfigRequest proto.InternalMessageInfo

func (m *QueryConfigRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *QueryConfigRequest) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *QueryConfigRequest) GetNamespaceName() string {
	if m != nil {
		return m.NamespaceName
	}
	return ""
}

type ConfigResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt"`
	UpdatedAt            int64    `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt"`
	AppName              string   `protobuf:"bytes,4,opt,name=appName,proto3" json:"appName"`
	ClusterName          string   `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName"`
	NamespaceName        string   `protobuf:"bytes,6,opt,name=namespaceName,proto3" json:"namespaceName"`
	Format               string   `protobuf:"bytes,7,opt,name=format,proto3" json:"format"`
	Value                string   `protobuf:"bytes,8,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigResponse) Reset()         { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()    {}
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{11}
}

func (m *ConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigResponse.Unmarshal(m, b)
}
func (m *ConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigResponse.Marshal(b, m, deterministic)
}
func (m *ConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigResponse.Merge(m, src)
}
func (m *ConfigResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigResponse.Size(m)
}
func (m *ConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigResponse proto.InternalMessageInfo

func (m *ConfigResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ConfigResponse) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *ConfigResponse) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *ConfigResponse) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *ConfigResponse) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ConfigResponse) GetNamespaceName() string {
	if m != nil {
		return m.NamespaceName
	}
	return ""
}

func (m *ConfigResponse) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *ConfigResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ReleaseRequest struct {
	AppName              string   `protobuf:"bytes,4,opt,name=appName,proto3" json:"appName"`
	ClusterName          string   `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName"`
	NamespaceName        string   `protobuf:"bytes,6,opt,name=namespaceName,proto3" json:"namespaceName"`
	Tag                  string   `protobuf:"bytes,7,opt,name=tag,proto3" json:"tag"`
	Comment              string   `protobuf:"bytes,9,opt,name=comment,proto3" json:"comment"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseRequest) Reset()         { *m = ReleaseRequest{} }
func (m *ReleaseRequest) String() string { return proto.CompactTextString(m) }
func (*ReleaseRequest) ProtoMessage()    {}
func (*ReleaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{12}
}

func (m *ReleaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseRequest.Unmarshal(m, b)
}
func (m *ReleaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseRequest.Marshal(b, m, deterministic)
}
func (m *ReleaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseRequest.Merge(m, src)
}
func (m *ReleaseRequest) XXX_Size() int {
	return xxx_messageInfo_ReleaseRequest.Size(m)
}
func (m *ReleaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseRequest proto.InternalMessageInfo

func (m *ReleaseRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *ReleaseRequest) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ReleaseRequest) GetNamespaceName() string {
	if m != nil {
		return m.NamespaceName
	}
	return ""
}

func (m *ReleaseRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *ReleaseRequest) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type ReleaseResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt"`
	UpdatedAt            int64    `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt"`
	AppName              string   `protobuf:"bytes,4,opt,name=appName,proto3" json:"appName"`
	ClusterName          string   `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName"`
	NamespaceName        string   `protobuf:"bytes,6,opt,name=namespaceName,proto3" json:"namespaceName"`
	Tag                  string   `protobuf:"bytes,7,opt,name=tag,proto3" json:"tag"`
	Value                string   `protobuf:"bytes,8,opt,name=value,proto3" json:"value"`
	Comment              string   `protobuf:"bytes,9,opt,name=comment,proto3" json:"comment"`
	Type                 string   `protobuf:"bytes,10,opt,name=type,proto3" json:"type"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseResponse) Reset()         { *m = ReleaseResponse{} }
func (m *ReleaseResponse) String() string { return proto.CompactTextString(m) }
func (*ReleaseResponse) ProtoMessage()    {}
func (*ReleaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{13}
}

func (m *ReleaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseResponse.Unmarshal(m, b)
}
func (m *ReleaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseResponse.Marshal(b, m, deterministic)
}
func (m *ReleaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseResponse.Merge(m, src)
}
func (m *ReleaseResponse) XXX_Size() int {
	return xxx_messageInfo_ReleaseResponse.Size(m)
}
func (m *ReleaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseResponse proto.InternalMessageInfo

func (m *ReleaseResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReleaseResponse) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *ReleaseResponse) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *ReleaseResponse) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *ReleaseResponse) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ReleaseResponse) GetNamespaceName() string {
	if m != nil {
		return m.NamespaceName
	}
	return ""
}

func (m *ReleaseResponse) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *ReleaseResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ReleaseResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *ReleaseResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type ReleaseHistoryResponse struct {
	ReleaseHistory       []*ReleaseResponse `protobuf:"bytes,1,rep,name=releaseHistory,proto3" json:"releaseHistory"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ReleaseHistoryResponse) Reset()         { *m = ReleaseHistoryResponse{} }
func (m *ReleaseHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*ReleaseHistoryResponse) ProtoMessage()    {}
func (*ReleaseHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{14}
}

func (m *ReleaseHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseHistoryResponse.Unmarshal(m, b)
}
func (m *ReleaseHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseHistoryResponse.Marshal(b, m, deterministic)
}
func (m *ReleaseHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseHistoryResponse.Merge(m, src)
}
func (m *ReleaseHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_ReleaseHistoryResponse.Size(m)
}
func (m *ReleaseHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseHistoryResponse proto.InternalMessageInfo

func (m *ReleaseHistoryResponse) GetReleaseHistory() []*ReleaseResponse {
	if m != nil {
		return m.ReleaseHistory
	}
	return nil
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{15}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{16}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AppRequest)(nil), "AppRequest")
	proto.RegisterType((*AppResponse)(nil), "AppResponse")
	proto.RegisterType((*AppsResponse)(nil), "AppsResponse")
	proto.RegisterType((*ClusterRequest)(nil), "ClusterRequest")
	proto.RegisterType((*ClusterResponse)(nil), "ClusterResponse")
	proto.RegisterType((*ClustersResponse)(nil), "ClustersResponse")
	proto.RegisterType((*NamespaceRequest)(nil), "NamespaceRequest")
	proto.RegisterType((*NamespaceResponse)(nil), "NamespaceResponse")
	proto.RegisterType((*NamespacesResponse)(nil), "NamespacesResponse")
	proto.RegisterType((*UpdateConfigRequest)(nil), "UpdateConfigRequest")
	proto.RegisterType((*QueryConfigRequest)(nil), "QueryConfigRequest")
	proto.RegisterType((*ConfigResponse)(nil), "ConfigResponse")
	proto.RegisterType((*ReleaseRequest)(nil), "ReleaseRequest")
	proto.RegisterType((*ReleaseResponse)(nil), "ReleaseResponse")
	proto.RegisterType((*ReleaseHistoryResponse)(nil), "ReleaseHistoryResponse")
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("config/config.proto", fileDescriptor_cc332a44e926b360) }

var fileDescriptor_cc332a44e926b360 = []byte{
	// 757 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x57, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xb6, 0x93, 0x34, 0xb1, 0x27, 0x89, 0x9d, 0x6e, 0xaa, 0x62, 0x45, 0x1c, 0xa2, 0x15, 0x3f,
	0x39, 0xd0, 0xa5, 0xb4, 0x1c, 0x0a, 0x27, 0x42, 0x39, 0xf4, 0x80, 0x90, 0xb0, 0x04, 0x9c, 0x5d,
	0x7b, 0x5b, 0x2c, 0x92, 0x78, 0xb1, 0x37, 0x95, 0x7a, 0xe4, 0x01, 0x38, 0xf6, 0xca, 0x23, 0xf0,
	0x02, 0x3c, 0x02, 0xaf, 0xc2, 0x43, 0x20, 0xaf, 0x7f, 0xe2, 0xb5, 0x5d, 0x0a, 0x14, 0x45, 0xc0,
	0x29, 0xde, 0x99, 0xd9, 0x9d, 0x6f, 0xbe, 0x9d, 0x9d, 0x99, 0xc0, 0xd0, 0x0d, 0x16, 0x27, 0xfe,
	0xe9, 0xfd, 0xe4, 0x87, 0xb0, 0x30, 0xe0, 0x01, 0x3e, 0x02, 0x98, 0x32, 0x66, 0xd3, 0xf7, 0x4b,
	0x1a, 0x71, 0x64, 0x41, 0xc7, 0x61, 0xec, 0x85, 0x33, 0xa7, 0x56, 0x6b, 0xac, 0x4e, 0x74, 0x3b,
	0x5b, 0xa2, 0x31, 0x74, 0x3d, 0x1a, 0xb9, 0xa1, 0xcf, 0xb8, 0x1f, 0x2c, 0xac, 0x0d, 0xa1, 0x2d,
	0x8a, 0xf0, 0x85, 0x0a, 0x5d, 0x71, 0x54, 0xc4, 0x82, 0x45, 0x44, 0x91, 0x01, 0x0d, 0xdf, 0xb3,
	0xd4, 0xb1, 0x3a, 0x69, 0xda, 0x0d, 0xdf, 0x43, 0x37, 0x41, 0x77, 0x43, 0xea, 0x70, 0xea, 0x4d,
	0xb9, 0xd5, 0x10, 0xe2, 0x95, 0x20, 0xd6, 0x2e, 0x99, 0x97, 0x6a, 0x9b, 0x89, 0x36, 0x17, 0x5c,
	0x0b, 0xd7, 0x2e, 0xf4, 0xa6, 0x8c, 0x45, 0x39, 0xae, 0x31, 0xb4, 0x1c, 0xc6, 0x22, 0x4b, 0x1d,
	0x37, 0x27, 0xdd, 0xbd, 0x1e, 0x29, 0x60, 0xb6, 0x85, 0x06, 0x2f, 0xc0, 0x38, 0x9c, 0x2d, 0x23,
	0x4e, 0xc3, 0x9f, 0xe2, 0xc5, 0x4d, 0x6c, 0x85, 0x36, 0xf5, 0x5f, 0x10, 0x95, 0x11, 0xb6, 0xab,
	0x08, 0xbf, 0xa8, 0x60, 0xe6, 0x0e, 0xd7, 0xcd, 0xde, 0xb5, 0xd1, 0x3f, 0x81, 0x41, 0x0a, 0x7e,
	0xc5, 0xf1, 0x3d, 0xd0, 0xd2, 0x43, 0x32, 0x9e, 0x07, 0xa4, 0x14, 0xa1, 0x9d, 0x5b, 0xe0, 0xcf,
	0x2a, 0x0c, 0x62, 0x67, 0x11, 0x73, 0x5c, 0xfa, 0x27, 0x28, 0xbf, 0x05, 0xfd, 0x45, 0x76, 0x9e,
	0xb0, 0x49, 0x60, 0xcb, 0x42, 0xb4, 0x0d, 0xed, 0x93, 0x20, 0x9c, 0x3b, 0xdc, 0xea, 0x08, 0x75,
	0xba, 0x2a, 0x87, 0x0c, 0xd5, 0x90, 0xbf, 0x36, 0x60, 0xb3, 0x00, 0xf8, 0x2f, 0xbb, 0xb2, 0xeb,
	0x45, 0xbf, 0x05, 0x1b, 0x67, 0xce, 0x6c, 0x49, 0x2d, 0x4d, 0x88, 0x93, 0x05, 0x1a, 0x81, 0x16,
	0xd2, 0x19, 0x75, 0x22, 0xea, 0x59, 0xfa, 0x58, 0x9d, 0x68, 0x76, 0xbe, 0x8e, 0x23, 0xa1, 0x9e,
	0xcf, 0x5f, 0x8b, 0x5d, 0x09, 0x5b, 0x2b, 0x41, 0x99, 0xcd, 0x6e, 0x95, 0xcd, 0x23, 0x40, 0x39,
	0x99, 0xab, 0x14, 0xda, 0x03, 0xc8, 0x01, 0x67, 0x49, 0x84, 0x48, 0x85, 0x75, 0xbb, 0x60, 0x85,
	0x3f, 0xaa, 0x30, 0x7c, 0x25, 0x38, 0x3c, 0x14, 0x35, 0x6e, 0x7d, 0xb9, 0x54, 0xcb, 0x1a, 0x3e,
	0x03, 0xf4, 0x72, 0x49, 0xc3, 0xf3, 0x35, 0xa3, 0xc1, 0xdf, 0x54, 0x30, 0x32, 0x9f, 0xff, 0x7f,
	0x72, 0xe2, 0x4f, 0x2a, 0x18, 0x76, 0x92, 0x8d, 0xeb, 0xbb, 0xf1, 0x01, 0x34, 0xb9, 0x73, 0x9a,
	0xe2, 0x8b, 0x3f, 0x63, 0x9f, 0x6e, 0x30, 0x9f, 0xd3, 0x05, 0x17, 0x4f, 0x44, 0xb7, 0xb3, 0x25,
	0xbe, 0x68, 0x80, 0x99, 0x03, 0xfc, 0x27, 0x2f, 0xa4, 0x1a, 0x6d, 0x7d, 0x9d, 0xb8, 0x94, 0x03,
	0x84, 0xa0, 0xc5, 0xcf, 0x59, 0x56, 0x20, 0xc4, 0x37, 0xb6, 0x61, 0x3b, 0xa5, 0xe5, 0xc8, 0x8f,
	0x78, 0x10, 0x9e, 0xe7, 0xec, 0x1c, 0x80, 0x11, 0x4a, 0x9a, 0xbc, 0x8d, 0x94, 0x78, 0xb4, 0x4b,
	0x76, 0x58, 0x87, 0x4e, 0x9a, 0x04, 0x18, 0x40, 0xcb, 0xcc, 0xf6, 0x3e, 0x74, 0xa0, 0x9d, 0x3c,
	0x09, 0x34, 0x01, 0xfd, 0x50, 0x10, 0x3b, 0x65, 0x0c, 0x75, 0xc9, 0x6a, 0xfc, 0x19, 0x49, 0xc3,
	0x00, 0x56, 0xd0, 0x5d, 0xd0, 0xc4, 0xfb, 0xbd, 0xd2, 0xf0, 0x36, 0xe8, 0xcf, 0xe8, 0x8c, 0xd6,
	0x1c, 0xa9, 0x13, 0xc9, 0x4c, 0x7b, 0xee, 0x47, 0x3c, 0x1e, 0x47, 0x90, 0x46, 0x32, 0x93, 0x3e,
	0x29, 0xce, 0x27, 0x58, 0x41, 0x0f, 0xa1, 0x9f, 0x00, 0x4c, 0x5b, 0x26, 0x32, 0x89, 0x3c, 0x8f,
	0x8c, 0x2a, 0xdd, 0x14, 0x2b, 0x68, 0x1f, 0x7a, 0x49, 0xb1, 0xf9, 0x95, 0x4d, 0x3b, 0xd0, 0x4f,
	0x80, 0x5f, 0xba, 0x4b, 0x0a, 0x60, 0x17, 0x7a, 0x71, 0x00, 0x59, 0xbf, 0x97, 0x43, 0xdd, 0x24,
	0xe5, 0x39, 0x00, 0x2b, 0xe8, 0x31, 0x98, 0x49, 0x2c, 0x79, 0xe5, 0x46, 0x9b, 0xa4, 0xdc, 0xec,
	0x47, 0x35, 0x85, 0x1d, 0x2b, 0xe8, 0x01, 0x98, 0x09, 0xb8, 0x1f, 0xee, 0x95, 0x00, 0x3e, 0x02,
	0x43, 0x90, 0xf0, 0x1b, 0xde, 0x0e, 0xc0, 0x88, 0x63, 0x5b, 0xb5, 0xa2, 0x2a, 0x17, 0x43, 0x52,
	0x6d, 0x54, 0x02, 0x67, 0xaf, 0xd8, 0x75, 0xd0, 0x16, 0xa9, 0x69, 0x42, 0x32, 0xce, 0x1d, 0xe8,
	0xa7, 0x89, 0x9c, 0xee, 0x31, 0x89, 0x5c, 0xc1, 0x64, 0xf3, 0xa7, 0x80, 0x62, 0x6c, 0xf2, 0x63,
	0xa9, 0x0b, 0xed, 0x06, 0xa9, 0x7f, 0x50, 0x58, 0x41, 0x13, 0xd0, 0xec, 0x60, 0x36, 0x3b, 0x76,
	0xdc, 0x77, 0x57, 0x78, 0x23, 0xd0, 0xb2, 0xa9, 0xe3, 0xa1, 0x21, 0xa9, 0x76, 0xaf, 0x91, 0x49,
	0xe4, 0xce, 0x82, 0x15, 0x74, 0x07, 0x36, 0xde, 0x38, 0xdc, 0x7d, 0x5b, 0xc8, 0xe9, 0xaa, 0xd5,
	0xae, 0x7a, 0xdc, 0x16, 0x7f, 0x39, 0xf6, 0xbf, 0x07, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x1e, 0xd4,
	0xe3, 0x89, 0x0c, 0x00, 0x00,
}
