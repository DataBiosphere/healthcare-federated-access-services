// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consents/v1/consents.proto

// Package v1 provides protocol buffer versions of Remembered Consents API for
// listing and revoking Remembered consents.
package v1

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// RequestMatchType defines what request is valid to use this consent.
type Consent_RequestMatchType int32

const (
	// NONE : do not remember.
	Consent_NONE Consent_RequestMatchType = 0
	// SUBSET : request resource and scopes are subset of resource and scopes in
	// this item.
	Consent_SUBSET Consent_RequestMatchType = 1
	// ANYTHING : request anything.
	Consent_ANYTHING Consent_RequestMatchType = 2
)

var Consent_RequestMatchType_name = map[int32]string{
	0: "NONE",
	1: "SUBSET",
	2: "ANYTHING",
}

var Consent_RequestMatchType_value = map[string]int32{
	"NONE":     0,
	"SUBSET":   1,
	"ANYTHING": 2,
}

func (x Consent_RequestMatchType) String() string {
	return proto.EnumName(Consent_RequestMatchType_name, int32(x))
}

func (Consent_RequestMatchType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0bb2e176a5b7c192, []int{5, 0}
}

// ReleaseType defines what to release.
type Consent_ReleaseType int32

const (
	Consent_UNSPECIFIED Consent_ReleaseType = 0
	// SELECTED : release selected visas of this item.
	Consent_SELECTED Consent_ReleaseType = 1
	// ANYTHING_NEEDED: release anything request needed.
	Consent_ANYTHING_NEEDED Consent_ReleaseType = 2
)

var Consent_ReleaseType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "SELECTED",
	2: "ANYTHING_NEEDED",
}

var Consent_ReleaseType_value = map[string]int32{
	"UNSPECIFIED":     0,
	"SELECTED":        1,
	"ANYTHING_NEEDED": 2,
}

func (x Consent_ReleaseType) String() string {
	return proto.EnumName(Consent_ReleaseType_name, int32(x))
}

func (Consent_ReleaseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0bb2e176a5b7c192, []int{5, 1}
}

// OAuth Consent.
//
// Deprecated: Do not use.
type MockConsent struct {
	// Name of the OAuth Consent.
	// Format: `consents/{consent_id}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Identifies the user who gave the OAuth consent.
	// E.g. subject or account number
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// Identifies the client for which the OAuth consent was given.
	Client string `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
	// Identifies the items for which the OAuth consent was given.
	// E.g. JTI of a Visa JWT.
	Items []string `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	// Identifies the scopes for which the OAuth consent was given.
	Scopes []string `protobuf:"bytes,5,rep,name=scopes,proto3" json:"scopes,omitempty"`
	// Identifies the resources for which the OAuth consent was given.
	Resouces []string `protobuf:"bytes,6,rep,name=resouces,proto3" json:"resouces,omitempty"`
	// Time at which OAuth consent was first given.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Time at which consent was last updated.
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MockConsent) Reset()         { *m = MockConsent{} }
func (m *MockConsent) String() string { return proto.CompactTextString(m) }
func (*MockConsent) ProtoMessage()    {}
func (*MockConsent) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb2e176a5b7c192, []int{0}
}

func (m *MockConsent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MockConsent.Unmarshal(m, b)
}
func (m *MockConsent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MockConsent.Marshal(b, m, deterministic)
}
func (m *MockConsent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MockConsent.Merge(m, src)
}
func (m *MockConsent) XXX_Size() int {
	return xxx_messageInfo_MockConsent.Size(m)
}
func (m *MockConsent) XXX_DiscardUnknown() {
	xxx_messageInfo_MockConsent.DiscardUnknown(m)
}

var xxx_messageInfo_MockConsent proto.InternalMessageInfo

func (m *MockConsent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MockConsent) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *MockConsent) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

func (m *MockConsent) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *MockConsent) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *MockConsent) GetResouces() []string {
	if m != nil {
		return m.Resouces
	}
	return nil
}

func (m *MockConsent) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *MockConsent) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

type DeleteConsentRequest struct {
	Realm                string   `protobuf:"bytes,1,opt,name=realm,proto3" json:"realm,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ConsentId            string   `protobuf:"bytes,3,opt,name=consent_id,json=consentId,proto3" json:"consent_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteConsentRequest) Reset()         { *m = DeleteConsentRequest{} }
func (m *DeleteConsentRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteConsentRequest) ProtoMessage()    {}
func (*DeleteConsentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb2e176a5b7c192, []int{1}
}

func (m *DeleteConsentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteConsentRequest.Unmarshal(m, b)
}
func (m *DeleteConsentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteConsentRequest.Marshal(b, m, deterministic)
}
func (m *DeleteConsentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteConsentRequest.Merge(m, src)
}
func (m *DeleteConsentRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteConsentRequest.Size(m)
}
func (m *DeleteConsentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteConsentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteConsentRequest proto.InternalMessageInfo

func (m *DeleteConsentRequest) GetRealm() string {
	if m != nil {
		return m.Realm
	}
	return ""
}

func (m *DeleteConsentRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *DeleteConsentRequest) GetConsentId() string {
	if m != nil {
		return m.ConsentId
	}
	return ""
}

type ListConsentsRequest struct {
	Realm                string   `protobuf:"bytes,1,opt,name=realm,proto3" json:"realm,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Filter               string   `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	PageSize             int32    `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            string   `protobuf:"bytes,5,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListConsentsRequest) Reset()         { *m = ListConsentsRequest{} }
func (m *ListConsentsRequest) String() string { return proto.CompactTextString(m) }
func (*ListConsentsRequest) ProtoMessage()    {}
func (*ListConsentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb2e176a5b7c192, []int{2}
}

func (m *ListConsentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListConsentsRequest.Unmarshal(m, b)
}
func (m *ListConsentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListConsentsRequest.Marshal(b, m, deterministic)
}
func (m *ListConsentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListConsentsRequest.Merge(m, src)
}
func (m *ListConsentsRequest) XXX_Size() int {
	return xxx_messageInfo_ListConsentsRequest.Size(m)
}
func (m *ListConsentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListConsentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListConsentsRequest proto.InternalMessageInfo

func (m *ListConsentsRequest) GetRealm() string {
	if m != nil {
		return m.Realm
	}
	return ""
}

func (m *ListConsentsRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ListConsentsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListConsentsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListConsentsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type MockListConsentsResponse struct {
	// Deprecated: keep for +1 backward compatibility
	Consents             []*MockConsent `protobuf:"bytes,1,rep,name=consents,proto3" json:"consents,omitempty"`
	NextPageToken        string         `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MockListConsentsResponse) Reset()         { *m = MockListConsentsResponse{} }
func (m *MockListConsentsResponse) String() string { return proto.CompactTextString(m) }
func (*MockListConsentsResponse) ProtoMessage()    {}
func (*MockListConsentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb2e176a5b7c192, []int{3}
}

func (m *MockListConsentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MockListConsentsResponse.Unmarshal(m, b)
}
func (m *MockListConsentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MockListConsentsResponse.Marshal(b, m, deterministic)
}
func (m *MockListConsentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MockListConsentsResponse.Merge(m, src)
}
func (m *MockListConsentsResponse) XXX_Size() int {
	return xxx_messageInfo_MockListConsentsResponse.Size(m)
}
func (m *MockListConsentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MockListConsentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MockListConsentsResponse proto.InternalMessageInfo

func (m *MockListConsentsResponse) GetConsents() []*MockConsent {
	if m != nil {
		return m.Consents
	}
	return nil
}

func (m *MockListConsentsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type ListConsentsResponse struct {
	Consents             []*Consent `protobuf:"bytes,1,rep,name=consents,proto3" json:"consents,omitempty"`
	NextPageToken        string     `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListConsentsResponse) Reset()         { *m = ListConsentsResponse{} }
func (m *ListConsentsResponse) String() string { return proto.CompactTextString(m) }
func (*ListConsentsResponse) ProtoMessage()    {}
func (*ListConsentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb2e176a5b7c192, []int{4}
}

func (m *ListConsentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListConsentsResponse.Unmarshal(m, b)
}
func (m *ListConsentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListConsentsResponse.Marshal(b, m, deterministic)
}
func (m *ListConsentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListConsentsResponse.Merge(m, src)
}
func (m *ListConsentsResponse) XXX_Size() int {
	return xxx_messageInfo_ListConsentsResponse.Size(m)
}
func (m *ListConsentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListConsentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListConsentsResponse proto.InternalMessageInfo

func (m *ListConsentsResponse) GetConsents() []*Consent {
	if m != nil {
		return m.Consents
	}
	return nil
}

func (m *ListConsentsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Consent contains the consent a user has given for release of visas to a
// specific OAuth client.
type Consent struct {
	// name iin format: /users/{user_id}/consents/{consent_id}
	Name                 string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Client               *Consent_Client          `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	CreateTime           *timestamp.Timestamp     `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	ExpireTime           *timestamp.Timestamp     `protobuf:"bytes,4,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	RequestMatchType     Consent_RequestMatchType `protobuf:"varint,5,opt,name=request_match_type,json=requestMatchType,proto3,enum=consents.v1.Consent_RequestMatchType" json:"request_match_type,omitempty"`
	RequestedResources   []string                 `protobuf:"bytes,6,rep,name=requested_resources,json=requestedResources,proto3" json:"requested_resources,omitempty"`
	RequestedScopes      []string                 `protobuf:"bytes,7,rep,name=requested_scopes,json=requestedScopes,proto3" json:"requested_scopes,omitempty"`
	ReleaseType          Consent_ReleaseType      `protobuf:"varint,8,opt,name=release_type,json=releaseType,proto3,enum=consents.v1.Consent_ReleaseType" json:"release_type,omitempty"`
	SelectedVisas        []*Consent_Visa          `protobuf:"bytes,9,rep,name=selected_visas,json=selectedVisas,proto3" json:"selected_visas,omitempty"`
	ReleaseProfileName   bool                     `protobuf:"varint,10,opt,name=release_profile_name,json=releaseProfileName,proto3" json:"release_profile_name,omitempty"`
	ReleaseProfileEmail  bool                     `protobuf:"varint,11,opt,name=release_profile_email,json=releaseProfileEmail,proto3" json:"release_profile_email,omitempty"`
	ReleaseProfileOther  bool                     `protobuf:"varint,12,opt,name=release_profile_other,json=releaseProfileOther,proto3" json:"release_profile_other,omitempty"`
	ReleaseAccountAdmin  bool                     `protobuf:"varint,13,opt,name=release_account_admin,json=releaseAccountAdmin,proto3" json:"release_account_admin,omitempty"`
	ReleaseLink          bool                     `protobuf:"varint,14,opt,name=release_link,json=releaseLink,proto3" json:"release_link,omitempty"`
	ReleaseIdentities    bool                     `protobuf:"varint,15,opt,name=release_identities,json=releaseIdentities,proto3" json:"release_identities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Consent) Reset()         { *m = Consent{} }
func (m *Consent) String() string { return proto.CompactTextString(m) }
func (*Consent) ProtoMessage()    {}
func (*Consent) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb2e176a5b7c192, []int{5}
}

func (m *Consent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consent.Unmarshal(m, b)
}
func (m *Consent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consent.Marshal(b, m, deterministic)
}
func (m *Consent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consent.Merge(m, src)
}
func (m *Consent) XXX_Size() int {
	return xxx_messageInfo_Consent.Size(m)
}
func (m *Consent) XXX_DiscardUnknown() {
	xxx_messageInfo_Consent.DiscardUnknown(m)
}

var xxx_messageInfo_Consent proto.InternalMessageInfo

func (m *Consent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Consent) GetClient() *Consent_Client {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *Consent) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Consent) GetExpireTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpireTime
	}
	return nil
}

func (m *Consent) GetRequestMatchType() Consent_RequestMatchType {
	if m != nil {
		return m.RequestMatchType
	}
	return Consent_NONE
}

func (m *Consent) GetRequestedResources() []string {
	if m != nil {
		return m.RequestedResources
	}
	return nil
}

func (m *Consent) GetRequestedScopes() []string {
	if m != nil {
		return m.RequestedScopes
	}
	return nil
}

func (m *Consent) GetReleaseType() Consent_ReleaseType {
	if m != nil {
		return m.ReleaseType
	}
	return Consent_UNSPECIFIED
}

func (m *Consent) GetSelectedVisas() []*Consent_Visa {
	if m != nil {
		return m.SelectedVisas
	}
	return nil
}

func (m *Consent) GetReleaseProfileName() bool {
	if m != nil {
		return m.ReleaseProfileName
	}
	return false
}

func (m *Consent) GetReleaseProfileEmail() bool {
	if m != nil {
		return m.ReleaseProfileEmail
	}
	return false
}

func (m *Consent) GetReleaseProfileOther() bool {
	if m != nil {
		return m.ReleaseProfileOther
	}
	return false
}

func (m *Consent) GetReleaseAccountAdmin() bool {
	if m != nil {
		return m.ReleaseAccountAdmin
	}
	return false
}

func (m *Consent) GetReleaseLink() bool {
	if m != nil {
		return m.ReleaseLink
	}
	return false
}

func (m *Consent) GetReleaseIdentities() bool {
	if m != nil {
		return m.ReleaseIdentities
	}
	return false
}

// Visa contains fields to match released visas user have.
type Consent_Visa struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Source               string   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	By                   string   `protobuf:"bytes,3,opt,name=by,proto3" json:"by,omitempty"`
	Iss                  string   `protobuf:"bytes,4,opt,name=iss,proto3" json:"iss,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Consent_Visa) Reset()         { *m = Consent_Visa{} }
func (m *Consent_Visa) String() string { return proto.CompactTextString(m) }
func (*Consent_Visa) ProtoMessage()    {}
func (*Consent_Visa) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb2e176a5b7c192, []int{5, 0}
}

func (m *Consent_Visa) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consent_Visa.Unmarshal(m, b)
}
func (m *Consent_Visa) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consent_Visa.Marshal(b, m, deterministic)
}
func (m *Consent_Visa) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consent_Visa.Merge(m, src)
}
func (m *Consent_Visa) XXX_Size() int {
	return xxx_messageInfo_Consent_Visa.Size(m)
}
func (m *Consent_Visa) XXX_DiscardUnknown() {
	xxx_messageInfo_Consent_Visa.DiscardUnknown(m)
}

var xxx_messageInfo_Consent_Visa proto.InternalMessageInfo

func (m *Consent_Visa) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Consent_Visa) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Consent_Visa) GetBy() string {
	if m != nil {
		return m.By
	}
	return ""
}

func (m *Consent_Visa) GetIss() string {
	if m != nil {
		return m.Iss
	}
	return ""
}

// Client : remember the consent for this client.
type Consent_Client struct {
	ClientId             string            `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Ui                   map[string]string `protobuf:"bytes,3,rep,name=ui,proto3" json:"ui,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Consent_Client) Reset()         { *m = Consent_Client{} }
func (m *Consent_Client) String() string { return proto.CompactTextString(m) }
func (*Consent_Client) ProtoMessage()    {}
func (*Consent_Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb2e176a5b7c192, []int{5, 1}
}

func (m *Consent_Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consent_Client.Unmarshal(m, b)
}
func (m *Consent_Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consent_Client.Marshal(b, m, deterministic)
}
func (m *Consent_Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consent_Client.Merge(m, src)
}
func (m *Consent_Client) XXX_Size() int {
	return xxx_messageInfo_Consent_Client.Size(m)
}
func (m *Consent_Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Consent_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Consent_Client proto.InternalMessageInfo

func (m *Consent_Client) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Consent_Client) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Consent_Client) GetUi() map[string]string {
	if m != nil {
		return m.Ui
	}
	return nil
}

func init() {
	proto.RegisterEnum("consents.v1.Consent_RequestMatchType", Consent_RequestMatchType_name, Consent_RequestMatchType_value)
	proto.RegisterEnum("consents.v1.Consent_ReleaseType", Consent_ReleaseType_name, Consent_ReleaseType_value)
	proto.RegisterType((*MockConsent)(nil), "consents.v1.MockConsent")
	proto.RegisterType((*DeleteConsentRequest)(nil), "consents.v1.DeleteConsentRequest")
	proto.RegisterType((*ListConsentsRequest)(nil), "consents.v1.ListConsentsRequest")
	proto.RegisterType((*MockListConsentsResponse)(nil), "consents.v1.MockListConsentsResponse")
	proto.RegisterType((*ListConsentsResponse)(nil), "consents.v1.ListConsentsResponse")
	proto.RegisterType((*Consent)(nil), "consents.v1.Consent")
	proto.RegisterType((*Consent_Visa)(nil), "consents.v1.Consent.Visa")
	proto.RegisterType((*Consent_Client)(nil), "consents.v1.Consent.Client")
	proto.RegisterMapType((map[string]string)(nil), "consents.v1.Consent.Client.UiEntry")
}

func init() {
	proto.RegisterFile("proto/consents/v1/consents.proto", fileDescriptor_0bb2e176a5b7c192)
}

var fileDescriptor_0bb2e176a5b7c192 = []byte{
	// 1118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0x1b, 0xd5,
	0x13, 0xef, 0xae, 0x1d, 0xc7, 0x1e, 0xe7, 0xc3, 0x3d, 0xc9, 0xbf, 0xff, 0xc5, 0x29, 0xe0, 0x98,
	0xb6, 0x72, 0x0b, 0xb1, 0x9b, 0x04, 0x10, 0x94, 0x1b, 0x52, 0x67, 0x29, 0x96, 0x52, 0x37, 0xb2,
	0x1d, 0x10, 0xdc, 0x58, 0x9b, 0xdd, 0x49, 0x72, 0x94, 0xf5, 0xee, 0xb2, 0xe7, 0x38, 0x8a, 0x1b,
	0xe5, 0x02, 0x24, 0x9e, 0x00, 0x21, 0x71, 0x03, 0x77, 0x3c, 0x11, 0xaf, 0xc0, 0x83, 0xa0, 0xf3,
	0xb1, 0xeb, 0xd8, 0x31, 0xe9, 0x87, 0xb8, 0x3b, 0x33, 0xbf, 0x99, 0x9d, 0xdf, 0x99, 0x33, 0x1f,
	0x0b, 0x95, 0x28, 0x0e, 0x79, 0xd8, 0x70, 0xc3, 0x80, 0x61, 0xc0, 0x59, 0xe3, 0x6c, 0x33, 0x3d,
	0xd7, 0x25, 0x44, 0x8a, 0xa9, 0x7c, 0xb6, 0x59, 0xbe, 0x7b, 0x1c, 0x86, 0xc7, 0x3e, 0x36, 0x9c,
	0x88, 0x36, 0x9c, 0x20, 0x08, 0xb9, 0xc3, 0x69, 0x18, 0x68, 0xd3, 0xf2, 0x9a, 0x46, 0xa5, 0x74,
	0x38, 0x3c, 0x6a, 0xe0, 0x20, 0xe2, 0x23, 0x0d, 0xbe, 0x3f, 0x0d, 0x72, 0x3a, 0x40, 0xc6, 0x9d,
	0x41, 0xa4, 0x0c, 0xaa, 0xbf, 0x9a, 0x50, 0x7c, 0x1e, 0xba, 0xa7, 0x4d, 0x15, 0x8f, 0x10, 0xc8,
	0x06, 0xce, 0x00, 0x2d, 0xa3, 0x62, 0xd4, 0x0a, 0x1d, 0x79, 0x16, 0xba, 0x21, 0xc3, 0xd8, 0x32,
	0x95, 0x4e, 0x9c, 0xc9, 0x1d, 0xc8, 0xb9, 0x3e, 0xc5, 0x80, 0x5b, 0x19, 0xa9, 0xd5, 0x12, 0x59,
	0x85, 0x39, 0xca, 0x71, 0xc0, 0xac, 0x6c, 0x25, 0x53, 0x2b, 0x74, 0x94, 0x20, 0xac, 0x99, 0x1b,
	0x46, 0xc8, 0xac, 0x39, 0xa9, 0xd6, 0x12, 0x29, 0x43, 0x3e, 0x46, 0x16, 0x0e, 0x5d, 0x64, 0x56,
	0x4e, 0x22, 0xa9, 0x4c, 0xbe, 0x80, 0xa2, 0x1b, 0xa3, 0xc3, 0xb1, 0x2f, 0x38, 0x5b, 0xf3, 0x15,
	0xa3, 0x56, 0xdc, 0x2a, 0xd7, 0xd5, 0x85, 0xea, 0xc9, 0x85, 0xea, 0xbd, 0xe4, 0x42, 0x1d, 0x50,
	0xe6, 0x42, 0x21, 0x9c, 0x87, 0x91, 0x97, 0x3a, 0xe7, 0x5f, 0xed, 0xac, 0xcc, 0x85, 0xe2, 0x89,
	0x69, 0x19, 0x55, 0x0f, 0x56, 0x77, 0xd1, 0x47, 0x8e, 0x3a, 0x31, 0x1d, 0xfc, 0x61, 0x88, 0x4c,
	0xde, 0x2f, 0x46, 0xc7, 0x1f, 0xe8, 0x04, 0x29, 0x81, 0xfc, 0x1f, 0xe6, 0x45, 0x56, 0xfa, 0xd4,
	0xd3, 0x49, 0xca, 0x09, 0xb1, 0xe5, 0x91, 0x77, 0x01, 0xf4, 0x4b, 0x0a, 0x4c, 0xa5, 0xaa, 0xa0,
	0x35, 0x2d, 0xaf, 0xfa, 0x9b, 0x01, 0x2b, 0x7b, 0x94, 0x71, 0x1d, 0x84, 0xbd, 0x65, 0x94, 0x3b,
	0x90, 0x3b, 0xa2, 0x3e, 0xc7, 0x38, 0x79, 0x0c, 0x25, 0x91, 0x35, 0x28, 0x44, 0xce, 0x31, 0xf6,
	0x19, 0x7d, 0x89, 0x56, 0xb6, 0x62, 0xd4, 0xe6, 0x3a, 0x79, 0xa1, 0xe8, 0xd2, 0x97, 0x28, 0xa8,
	0x49, 0x90, 0x87, 0xa7, 0x18, 0x58, 0x73, 0x8a, 0x9a, 0xd0, 0xf4, 0x84, 0xa2, 0x7a, 0x0e, 0x96,
	0xa8, 0x8b, 0x49, 0x76, 0x2c, 0x12, 0x47, 0xf2, 0x31, 0xe4, 0x93, 0xfa, 0xb4, 0x8c, 0x4a, 0xa6,
	0x56, 0xdc, 0xb2, 0xea, 0x57, 0x0a, 0xb6, 0x7e, 0xa5, 0xa0, 0x3a, 0xa9, 0x25, 0x79, 0x00, 0xcb,
	0x01, 0x9e, 0xf3, 0xfe, 0x95, 0xa8, 0xea, 0x1a, 0x8b, 0x42, 0xbd, 0x9f, 0x46, 0x8e, 0x60, 0x75,
	0x66, 0xd4, 0xc7, 0xd7, 0xa2, 0xae, 0x4e, 0x44, 0x7d, 0xfb, 0x88, 0xbf, 0x17, 0x60, 0xfe, 0xa6,
	0x06, 0xd8, 0x4e, 0x8b, 0xdd, 0x94, 0x85, 0xb4, 0x36, 0x2b, 0x6e, 0xbd, 0x29, 0x4d, 0xd2, 0x4e,
	0x98, 0xaa, 0xdf, 0xcc, 0x9b, 0xd6, 0x2f, 0x9e, 0x47, 0x34, 0xd6, 0xce, 0xd9, 0x57, 0x3b, 0x2b,
	0x73, 0xe9, 0xdc, 0x05, 0x12, 0xab, 0x42, 0xea, 0x0f, 0x1c, 0xee, 0x9e, 0xf4, 0xf9, 0x28, 0x42,
	0xf9, 0xc2, 0x4b, 0x5b, 0xf7, 0x67, 0x52, 0xd7, 0x75, 0xf7, 0x5c, 0x58, 0xf7, 0x46, 0x11, 0x76,
	0x4a, 0xf1, 0x94, 0x86, 0x34, 0x60, 0x45, 0xeb, 0xd0, 0xeb, 0xcb, 0x26, 0x8d, 0xc7, 0x5d, 0x4b,
	0x52, 0xa8, 0x93, 0x20, 0xe4, 0x21, 0x94, 0xc6, 0x0e, 0xba, 0xfb, 0xe7, 0xa5, 0xf5, 0x72, 0xaa,
	0xef, 0xaa, 0x31, 0xd0, 0x84, 0x85, 0x18, 0x7d, 0x74, 0x18, 0x2a, 0xaa, 0x79, 0x49, 0xb5, 0xf2,
	0x2f, 0x54, 0xa5, 0xa1, 0x64, 0x59, 0x8c, 0xc7, 0x02, 0xf9, 0x12, 0x96, 0x18, 0xfa, 0xe8, 0x8a,
	0x70, 0x67, 0x94, 0x39, 0xcc, 0x2a, 0xc8, 0x22, 0x79, 0x67, 0xe6, 0x67, 0xbe, 0xa1, 0xcc, 0xe9,
	0x2c, 0x26, 0x0e, 0x42, 0x62, 0xe4, 0x31, 0xac, 0x26, 0x34, 0xa2, 0x38, 0x3c, 0xa2, 0x3e, 0xf6,
	0x65, 0x29, 0x40, 0xc5, 0xa8, 0xe5, 0xc5, 0x1d, 0x25, 0xb6, 0xaf, 0xa0, 0xb6, 0x28, 0x8c, 0x2d,
	0xf8, 0xdf, 0xb4, 0x07, 0x0e, 0x1c, 0xea, 0x5b, 0x45, 0xe9, 0xb2, 0x32, 0xe9, 0x62, 0x0b, 0x68,
	0x96, 0x4f, 0xc8, 0x4f, 0x30, 0xb6, 0x16, 0x66, 0xf9, 0xbc, 0x10, 0xd0, 0x55, 0x1f, 0xc7, 0x75,
	0xc3, 0x61, 0xc0, 0xfb, 0x8e, 0x37, 0xa0, 0x81, 0xb5, 0x38, 0xe1, 0xb3, 0xa3, 0xb0, 0x1d, 0x01,
	0x91, 0xf5, 0x71, 0x52, 0x7d, 0x1a, 0x9c, 0x5a, 0x4b, 0xd2, 0x34, 0x49, 0xd9, 0x1e, 0x0d, 0x4e,
	0xc9, 0x06, 0x24, 0x97, 0xea, 0x53, 0x0f, 0x03, 0x4e, 0x39, 0x45, 0x66, 0x2d, 0x4b, 0xc3, 0xdb,
	0x1a, 0x69, 0xa5, 0x40, 0xb9, 0x07, 0x59, 0x91, 0x28, 0xd1, 0x22, 0xf2, 0x99, 0x74, 0x8b, 0x88,
	0xb3, 0x9c, 0xf0, 0xf2, 0xe1, 0x93, 0xd1, 0xa4, 0x24, 0xb2, 0x04, 0xe6, 0xe1, 0x48, 0x8f, 0x25,
	0xf3, 0x70, 0x44, 0x4a, 0x90, 0xa1, 0x8c, 0xc9, 0x82, 0x2e, 0x74, 0xc4, 0xb1, 0xfc, 0xa7, 0x01,
	0x39, 0xd5, 0x3a, 0x62, 0x5e, 0xa9, 0xe6, 0x11, 0x23, 0x4e, 0x7d, 0x3d, 0xaf, 0x14, 0x2d, 0x2f,
	0x6d, 0x4c, 0x73, 0xa2, 0x31, 0xcd, 0x21, 0xb5, 0x32, 0xf2, 0x9d, 0x3f, 0xb8, 0xa1, 0x29, 0xeb,
	0x07, 0xd4, 0x0e, 0x78, 0x3c, 0xea, 0x98, 0x43, 0x5a, 0xfe, 0x04, 0xe6, 0xb5, 0x28, 0xd8, 0x9c,
	0xe2, 0x48, 0x87, 0x12, 0x47, 0x31, 0x79, 0xcf, 0x1c, 0x7f, 0x98, 0x84, 0x51, 0xc2, 0x13, 0xf3,
	0x33, 0xa3, 0xfa, 0x29, 0x94, 0xa6, 0xdb, 0x84, 0xe4, 0x21, 0xdb, 0x7e, 0xd1, 0xb6, 0x4b, 0xb7,
	0x08, 0x40, 0xae, 0x7b, 0xf0, 0xb4, 0x6b, 0xf7, 0x4a, 0x06, 0x59, 0x80, 0xfc, 0x4e, 0xfb, 0xbb,
	0xde, 0xd7, 0xad, 0xf6, 0xb3, 0x92, 0x59, 0xdd, 0x81, 0xe2, 0x95, 0x9a, 0x25, 0xcb, 0x50, 0x3c,
	0x68, 0x77, 0xf7, 0xed, 0x66, 0xeb, 0xab, 0x96, 0xbd, 0x5b, 0xba, 0x25, 0xac, 0xbb, 0xf6, 0x9e,
	0xdd, 0xec, 0xd9, 0xbb, 0x25, 0x83, 0xac, 0xc0, 0x72, 0xe2, 0xdb, 0x6f, 0xdb, 0xf6, 0xae, 0xbd,
	0x5b, 0x32, 0xb7, 0xfe, 0xc8, 0x42, 0x3e, 0x19, 0x87, 0xe4, 0x47, 0x03, 0x6e, 0x8b, 0x01, 0x3b,
	0xb1, 0x9e, 0xc8, 0xfa, 0xc4, 0xed, 0x67, 0xad, 0xae, 0xf2, 0x9d, 0x6b, 0xe3, 0xc3, 0x16, 0x7f,
	0x0a, 0xd5, 0xcd, 0x9f, 0xfe, 0xfa, 0xfb, 0x17, 0xf3, 0xc3, 0x47, 0x0f, 0xc5, 0x7f, 0x88, 0xd8,
	0x28, 0xac, 0x71, 0xa1, 0xf7, 0xcc, 0xe5, 0xf8, 0x1f, 0xe5, 0x62, 0xbc, 0xc7, 0x2e, 0x05, 0x87,
	0xd2, 0xf4, 0x76, 0x20, 0x93, 0xfd, 0x3a, 0x63, 0xad, 0x95, 0xef, 0x5f, 0xdb, 0x12, 0xb3, 0x06,
	0x7d, 0xf5, 0x9e, 0x24, 0xf4, 0x1e, 0xb9, 0x7b, 0x13, 0x21, 0xf2, 0xb3, 0x01, 0x8b, 0xff, 0x59,
	0x0e, 0x3e, 0x97, 0x21, 0xb7, 0x1f, 0x6d, 0x8a, 0x90, 0x17, 0x72, 0xdd, 0x5e, 0xbe, 0x7e, 0x2e,
	0x16, 0xde, 0x30, 0x0f, 0xeb, 0x37, 0x58, 0xe8, 0x1c, 0x7c, 0x24, 0x09, 0x3d, 0x20, 0xf7, 0x5e,
	0x87, 0xd0, 0xd3, 0x6f, 0xbf, 0x3f, 0x38, 0xa6, 0xfc, 0x64, 0x78, 0x58, 0x77, 0xc3, 0x41, 0xe3,
	0x99, 0xbc, 0x62, 0xd3, 0x0f, 0x87, 0xde, 0xbe, 0xef, 0xf0, 0xa3, 0x30, 0x1e, 0x34, 0x4e, 0xd0,
	0xf1, 0xf9, 0x89, 0xeb, 0xc4, 0xb8, 0x71, 0x84, 0x1e, 0xc6, 0x0e, 0x47, 0x6f, 0xc3, 0x71, 0x5d,
	0x64, 0x6c, 0x83, 0x61, 0x7c, 0x46, 0x5d, 0x64, 0x8d, 0x6b, 0xbf, 0xa5, 0x87, 0x39, 0xa9, 0xda,
	0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x0e, 0xfc, 0x51, 0xb2, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConsentsClient is the client API for Consents service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsentsClient interface {
	// Deletes the specified Remembered Consent.
	// Deprecated: keep for +1 backward compatibility
	MockDeleteConsent(ctx context.Context, in *DeleteConsentRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Lists the Remembered Consents.
	// Deprecated: keep for +1 backward compatibility
	MockListConsents(ctx context.Context, in *ListConsentsRequest, opts ...grpc.CallOption) (*MockListConsentsResponse, error)
	// Deletes the specified Remembered Consent.
	DeleteConsent(ctx context.Context, in *DeleteConsentRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Lists the Remembered Consents.
	ListConsents(ctx context.Context, in *ListConsentsRequest, opts ...grpc.CallOption) (*ListConsentsResponse, error)
}

type consentsClient struct {
	cc grpc.ClientConnInterface
}

func NewConsentsClient(cc grpc.ClientConnInterface) ConsentsClient {
	return &consentsClient{cc}
}

func (c *consentsClient) MockDeleteConsent(ctx context.Context, in *DeleteConsentRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/consents.v1.Consents/MockDeleteConsent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consentsClient) MockListConsents(ctx context.Context, in *ListConsentsRequest, opts ...grpc.CallOption) (*MockListConsentsResponse, error) {
	out := new(MockListConsentsResponse)
	err := c.cc.Invoke(ctx, "/consents.v1.Consents/MockListConsents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consentsClient) DeleteConsent(ctx context.Context, in *DeleteConsentRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/consents.v1.Consents/DeleteConsent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consentsClient) ListConsents(ctx context.Context, in *ListConsentsRequest, opts ...grpc.CallOption) (*ListConsentsResponse, error) {
	out := new(ListConsentsResponse)
	err := c.cc.Invoke(ctx, "/consents.v1.Consents/ListConsents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsentsServer is the server API for Consents service.
type ConsentsServer interface {
	// Deletes the specified Remembered Consent.
	// Deprecated: keep for +1 backward compatibility
	MockDeleteConsent(context.Context, *DeleteConsentRequest) (*empty.Empty, error)
	// Lists the Remembered Consents.
	// Deprecated: keep for +1 backward compatibility
	MockListConsents(context.Context, *ListConsentsRequest) (*MockListConsentsResponse, error)
	// Deletes the specified Remembered Consent.
	DeleteConsent(context.Context, *DeleteConsentRequest) (*empty.Empty, error)
	// Lists the Remembered Consents.
	ListConsents(context.Context, *ListConsentsRequest) (*ListConsentsResponse, error)
}

// UnimplementedConsentsServer can be embedded to have forward compatible implementations.
type UnimplementedConsentsServer struct {
}

func (*UnimplementedConsentsServer) MockDeleteConsent(ctx context.Context, req *DeleteConsentRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MockDeleteConsent not implemented")
}
func (*UnimplementedConsentsServer) MockListConsents(ctx context.Context, req *ListConsentsRequest) (*MockListConsentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MockListConsents not implemented")
}
func (*UnimplementedConsentsServer) DeleteConsent(ctx context.Context, req *DeleteConsentRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConsent not implemented")
}
func (*UnimplementedConsentsServer) ListConsents(ctx context.Context, req *ListConsentsRequest) (*ListConsentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConsents not implemented")
}

func RegisterConsentsServer(s *grpc.Server, srv ConsentsServer) {
	s.RegisterService(&_Consents_serviceDesc, srv)
}

func _Consents_MockDeleteConsent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConsentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsentsServer).MockDeleteConsent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consents.v1.Consents/MockDeleteConsent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsentsServer).MockDeleteConsent(ctx, req.(*DeleteConsentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consents_MockListConsents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConsentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsentsServer).MockListConsents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consents.v1.Consents/MockListConsents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsentsServer).MockListConsents(ctx, req.(*ListConsentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consents_DeleteConsent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConsentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsentsServer).DeleteConsent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consents.v1.Consents/DeleteConsent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsentsServer).DeleteConsent(ctx, req.(*DeleteConsentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consents_ListConsents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConsentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsentsServer).ListConsents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consents.v1.Consents/ListConsents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsentsServer).ListConsents(ctx, req.(*ListConsentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Consents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consents.v1.Consents",
	HandlerType: (*ConsentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MockDeleteConsent",
			Handler:    _Consents_MockDeleteConsent_Handler,
		},
		{
			MethodName: "MockListConsents",
			Handler:    _Consents_MockListConsents_Handler,
		},
		{
			MethodName: "DeleteConsent",
			Handler:    _Consents_DeleteConsent_Handler,
		},
		{
			MethodName: "ListConsents",
			Handler:    _Consents_ListConsents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/consents/v1/consents.proto",
}
