// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: redpanda/api/console/v1alpha1/authentication.proto

package consolev1alpha1

import (
	reflect "reflect"
	sync "sync"

	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SASLMechanism int32

const (
	// The SASL mechanism is unspecified.
	SASLMechanism_SASL_MECHANISM_UNSPECIFIED SASLMechanism = 0
	// The SASL mechanism using SCRAM-SHA-256.
	SASLMechanism_SASL_MECHANISM_SCRAM_SHA_256 SASLMechanism = 1
	// The SASL mechanism using SCRAM-SHA-512.
	SASLMechanism_SASL_MECHANISM_SCRAM_SHA_512 SASLMechanism = 2
)

// Enum value maps for SASLMechanism.
var (
	SASLMechanism_name = map[int32]string{
		0: "SASL_MECHANISM_UNSPECIFIED",
		1: "SASL_MECHANISM_SCRAM_SHA_256",
		2: "SASL_MECHANISM_SCRAM_SHA_512",
	}
	SASLMechanism_value = map[string]int32{
		"SASL_MECHANISM_UNSPECIFIED":   0,
		"SASL_MECHANISM_SCRAM_SHA_256": 1,
		"SASL_MECHANISM_SCRAM_SHA_512": 2,
	}
)

func (x SASLMechanism) Enum() *SASLMechanism {
	p := new(SASLMechanism)
	*p = x
	return p
}

func (x SASLMechanism) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SASLMechanism) Descriptor() protoreflect.EnumDescriptor {
	return file_redpanda_api_console_v1alpha1_authentication_proto_enumTypes[0].Descriptor()
}

func (SASLMechanism) Type() protoreflect.EnumType {
	return &file_redpanda_api_console_v1alpha1_authentication_proto_enumTypes[0]
}

func (x SASLMechanism) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SASLMechanism.Descriptor instead.
func (SASLMechanism) EnumDescriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_authentication_proto_rawDescGZIP(), []int{0}
}

type AuthenticationMethod int32

const (
	// The authentication method is unspecified.
	AuthenticationMethod_AUTHENTICATION_METHOD_UNSPECIFIED AuthenticationMethod = 0
	// The authentication method using OpenID Connect.
	AuthenticationMethod_AUTHENTICATION_METHOD_OIDC AuthenticationMethod = 1
	// The authentication method using SASL-SCRAM.
	AuthenticationMethod_AUTHENTICATION_METHOD_SASL_SCRAM AuthenticationMethod = 2
	// The authentication method for Redpanda Cloud.
	AuthenticationMethod_AUTHENTICATION_METHOD_REDPANDA_CLOUD AuthenticationMethod = 3
)

// Enum value maps for AuthenticationMethod.
var (
	AuthenticationMethod_name = map[int32]string{
		0: "AUTHENTICATION_METHOD_UNSPECIFIED",
		1: "AUTHENTICATION_METHOD_OIDC",
		2: "AUTHENTICATION_METHOD_SASL_SCRAM",
		3: "AUTHENTICATION_METHOD_REDPANDA_CLOUD",
	}
	AuthenticationMethod_value = map[string]int32{
		"AUTHENTICATION_METHOD_UNSPECIFIED":    0,
		"AUTHENTICATION_METHOD_OIDC":           1,
		"AUTHENTICATION_METHOD_SASL_SCRAM":     2,
		"AUTHENTICATION_METHOD_REDPANDA_CLOUD": 3,
	}
)

func (x AuthenticationMethod) Enum() *AuthenticationMethod {
	p := new(AuthenticationMethod)
	*p = x
	return p
}

func (x AuthenticationMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthenticationMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_redpanda_api_console_v1alpha1_authentication_proto_enumTypes[1].Descriptor()
}

func (AuthenticationMethod) Type() protoreflect.EnumType {
	return &file_redpanda_api_console_v1alpha1_authentication_proto_enumTypes[1]
}

func (x AuthenticationMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthenticationMethod.Descriptor instead.
func (AuthenticationMethod) EnumDescriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_authentication_proto_rawDescGZIP(), []int{1}
}

type LoginSaslScramRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The username for the login request.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// The password for the login request.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// The SASL mechanism to be used for authentication.
	Mechanism SASLMechanism `protobuf:"varint,3,opt,name=mechanism,proto3,enum=redpanda.api.console.v1alpha1.SASLMechanism" json:"mechanism,omitempty"`
	// Whether or not the session token should be returned in the body.
	ReturnToken bool `protobuf:"varint,4,opt,name=return_token,json=returnToken,proto3" json:"return_token,omitempty"`
}

func (x *LoginSaslScramRequest) Reset() {
	*x = LoginSaslScramRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginSaslScramRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginSaslScramRequest) ProtoMessage() {}

func (x *LoginSaslScramRequest) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginSaslScramRequest.ProtoReflect.Descriptor instead.
func (*LoginSaslScramRequest) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *LoginSaslScramRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginSaslScramRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginSaslScramRequest) GetMechanism() SASLMechanism {
	if x != nil {
		return x.Mechanism
	}
	return SASLMechanism_SASL_MECHANISM_UNSPECIFIED
}

func (x *LoginSaslScramRequest) GetReturnToken() bool {
	if x != nil {
		return x.ReturnToken
	}
	return false
}

type LoginSaslScramResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SessionToken is only set if return_token is set to true in the request.
	SessionToken string `protobuf:"bytes,1,opt,name=session_token,json=sessionToken,proto3" json:"session_token,omitempty"`
}

func (x *LoginSaslScramResponse) Reset() {
	*x = LoginSaslScramResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_authentication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginSaslScramResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginSaslScramResponse) ProtoMessage() {}

func (x *LoginSaslScramResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_authentication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginSaslScramResponse.ProtoReflect.Descriptor instead.
func (*LoginSaslScramResponse) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *LoginSaslScramResponse) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

type ListAuthenticationMethodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAuthenticationMethodsRequest) Reset() {
	*x = ListAuthenticationMethodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_authentication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuthenticationMethodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthenticationMethodsRequest) ProtoMessage() {}

func (x *ListAuthenticationMethodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_authentication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthenticationMethodsRequest.ProtoReflect.Descriptor instead.
func (*ListAuthenticationMethodsRequest) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_authentication_proto_rawDescGZIP(), []int{2}
}

type ListAuthenticationMethodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of available authentication methods.
	Methods []AuthenticationMethod `protobuf:"varint,1,rep,packed,name=methods,proto3,enum=redpanda.api.console.v1alpha1.AuthenticationMethod" json:"methods,omitempty"`
}

func (x *ListAuthenticationMethodsResponse) Reset() {
	*x = ListAuthenticationMethodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_authentication_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuthenticationMethodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthenticationMethodsResponse) ProtoMessage() {}

func (x *ListAuthenticationMethodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_authentication_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthenticationMethodsResponse.ProtoReflect.Descriptor instead.
func (*ListAuthenticationMethodsResponse) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_authentication_proto_rawDescGZIP(), []int{3}
}

func (x *ListAuthenticationMethodsResponse) GetMethods() []AuthenticationMethod {
	if x != nil {
		return x.Methods
	}
	return nil
}

var File_redpanda_api_console_v1alpha1_authentication_proto protoreflect.FileDescriptor

var file_redpanda_api_console_v1alpha1_authentication_proto_rawDesc = []byte{
	0x0a, 0x32, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x15, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x61, 0x73, 0x6c, 0x53,
	0x63, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xe0,
	0x41, 0x02, 0xba, 0x48, 0x0a, 0xc8, 0x01, 0x01, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xe0, 0x41, 0x02,
	0xba, 0x48, 0x0a, 0xc8, 0x01, 0x01, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x5d, 0x0a, 0x09, 0x6d, 0x65, 0x63, 0x68, 0x61,
	0x6e, 0x69, 0x73, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x72, 0x65, 0x64,
	0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x41, 0x53, 0x4c, 0x4d,
	0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x42, 0x11, 0xe0, 0x41, 0x02, 0xba, 0x48, 0x0b,
	0xc8, 0x01, 0x01, 0x82, 0x01, 0x05, 0x10, 0x01, 0x22, 0x01, 0x00, 0x52, 0x09, 0x6d, 0x65, 0x63,
	0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3d, 0x0a, 0x16, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x53, 0x61, 0x73, 0x6c, 0x53, 0x63, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x22, 0x0a, 0x20, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x77, 0x0a, 0x21,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x52, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x33, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x73, 0x2a, 0x73, 0x0a, 0x0d, 0x53, 0x41, 0x53, 0x4c, 0x4d, 0x65, 0x63,
	0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x41, 0x53, 0x4c, 0x5f, 0x4d,
	0x45, 0x43, 0x48, 0x41, 0x4e, 0x49, 0x53, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x41, 0x53, 0x4c, 0x5f, 0x4d,
	0x45, 0x43, 0x48, 0x41, 0x4e, 0x49, 0x53, 0x4d, 0x5f, 0x53, 0x43, 0x52, 0x41, 0x4d, 0x5f, 0x53,
	0x48, 0x41, 0x5f, 0x32, 0x35, 0x36, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x41, 0x53, 0x4c,
	0x5f, 0x4d, 0x45, 0x43, 0x48, 0x41, 0x4e, 0x49, 0x53, 0x4d, 0x5f, 0x53, 0x43, 0x52, 0x41, 0x4d,
	0x5f, 0x53, 0x48, 0x41, 0x5f, 0x35, 0x31, 0x32, 0x10, 0x02, 0x2a, 0xad, 0x01, 0x0a, 0x14, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x25, 0x0a, 0x21, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x55,
	0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x54,
	0x48, 0x4f, 0x44, 0x5f, 0x4f, 0x49, 0x44, 0x43, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20, 0x41, 0x55,
	0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x54,
	0x48, 0x4f, 0x44, 0x5f, 0x53, 0x41, 0x53, 0x4c, 0x5f, 0x53, 0x43, 0x52, 0x41, 0x4d, 0x10, 0x02,
	0x12, 0x28, 0x0a, 0x24, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x52, 0x45, 0x44, 0x50, 0x41, 0x4e,
	0x44, 0x41, 0x5f, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x10, 0x03, 0x32, 0xbb, 0x02, 0x0a, 0x15, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x7f, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x61, 0x73,
	0x6c, 0x53, 0x63, 0x72, 0x61, 0x6d, 0x12, 0x34, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x61, 0x73, 0x6c,
	0x53, 0x63, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x72,
	0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x53, 0x61, 0x73, 0x6c, 0x53, 0x63, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa0, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x73, 0x12, 0x3f, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xb4, 0x02, 0x0a, 0x21, 0x63, 0x6f, 0x6d,
	0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x13,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x64,
	0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x73, 0x6f,
	0x6c, 0x65, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x41, 0x43,
	0xaa, 0x02, 0x1d, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x41, 0x70, 0x69, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xca, 0x02, 0x1d, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x5c, 0x41, 0x70, 0x69, 0x5c,
	0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xe2, 0x02, 0x29, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x5c, 0x41, 0x70, 0x69, 0x5c,
	0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x52,
	0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f,
	0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_redpanda_api_console_v1alpha1_authentication_proto_rawDescOnce sync.Once
	file_redpanda_api_console_v1alpha1_authentication_proto_rawDescData = file_redpanda_api_console_v1alpha1_authentication_proto_rawDesc
)

func file_redpanda_api_console_v1alpha1_authentication_proto_rawDescGZIP() []byte {
	file_redpanda_api_console_v1alpha1_authentication_proto_rawDescOnce.Do(func() {
		file_redpanda_api_console_v1alpha1_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_redpanda_api_console_v1alpha1_authentication_proto_rawDescData)
	})
	return file_redpanda_api_console_v1alpha1_authentication_proto_rawDescData
}

var file_redpanda_api_console_v1alpha1_authentication_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_redpanda_api_console_v1alpha1_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_redpanda_api_console_v1alpha1_authentication_proto_goTypes = []interface{}{
	(SASLMechanism)(0),                        // 0: redpanda.api.console.v1alpha1.SASLMechanism
	(AuthenticationMethod)(0),                 // 1: redpanda.api.console.v1alpha1.AuthenticationMethod
	(*LoginSaslScramRequest)(nil),             // 2: redpanda.api.console.v1alpha1.LoginSaslScramRequest
	(*LoginSaslScramResponse)(nil),            // 3: redpanda.api.console.v1alpha1.LoginSaslScramResponse
	(*ListAuthenticationMethodsRequest)(nil),  // 4: redpanda.api.console.v1alpha1.ListAuthenticationMethodsRequest
	(*ListAuthenticationMethodsResponse)(nil), // 5: redpanda.api.console.v1alpha1.ListAuthenticationMethodsResponse
}
var file_redpanda_api_console_v1alpha1_authentication_proto_depIdxs = []int32{
	0, // 0: redpanda.api.console.v1alpha1.LoginSaslScramRequest.mechanism:type_name -> redpanda.api.console.v1alpha1.SASLMechanism
	1, // 1: redpanda.api.console.v1alpha1.ListAuthenticationMethodsResponse.methods:type_name -> redpanda.api.console.v1alpha1.AuthenticationMethod
	2, // 2: redpanda.api.console.v1alpha1.AuthenticationService.LoginSaslScram:input_type -> redpanda.api.console.v1alpha1.LoginSaslScramRequest
	4, // 3: redpanda.api.console.v1alpha1.AuthenticationService.ListAuthenticationMethods:input_type -> redpanda.api.console.v1alpha1.ListAuthenticationMethodsRequest
	3, // 4: redpanda.api.console.v1alpha1.AuthenticationService.LoginSaslScram:output_type -> redpanda.api.console.v1alpha1.LoginSaslScramResponse
	5, // 5: redpanda.api.console.v1alpha1.AuthenticationService.ListAuthenticationMethods:output_type -> redpanda.api.console.v1alpha1.ListAuthenticationMethodsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_redpanda_api_console_v1alpha1_authentication_proto_init() }
func file_redpanda_api_console_v1alpha1_authentication_proto_init() {
	if File_redpanda_api_console_v1alpha1_authentication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_redpanda_api_console_v1alpha1_authentication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginSaslScramRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_redpanda_api_console_v1alpha1_authentication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginSaslScramResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_redpanda_api_console_v1alpha1_authentication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuthenticationMethodsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_redpanda_api_console_v1alpha1_authentication_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuthenticationMethodsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_redpanda_api_console_v1alpha1_authentication_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_redpanda_api_console_v1alpha1_authentication_proto_goTypes,
		DependencyIndexes: file_redpanda_api_console_v1alpha1_authentication_proto_depIdxs,
		EnumInfos:         file_redpanda_api_console_v1alpha1_authentication_proto_enumTypes,
		MessageInfos:      file_redpanda_api_console_v1alpha1_authentication_proto_msgTypes,
	}.Build()
	File_redpanda_api_console_v1alpha1_authentication_proto = out.File
	file_redpanda_api_console_v1alpha1_authentication_proto_rawDesc = nil
	file_redpanda_api_console_v1alpha1_authentication_proto_goTypes = nil
	file_redpanda_api_console_v1alpha1_authentication_proto_depIdxs = nil
}
