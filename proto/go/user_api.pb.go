// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: user_api.proto

package protoapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// request object for the `/users` endpoint
type UserCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// required
	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	RoleId   *UUID  `protobuf:"bytes,4,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	RoleName string `protobuf:"bytes,5,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
}

func (x *UserCreateRequest) Reset() {
	*x = UserCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateRequest) ProtoMessage() {}

func (x *UserCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateRequest.ProtoReflect.Descriptor instead.
func (*UserCreateRequest) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{0}
}

func (x *UserCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserCreateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserCreateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserCreateRequest) GetRoleId() *UUID {
	if x != nil {
		return x.RoleId
	}
	return nil
}

func (x *UserCreateRequest) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

type UserCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserCreateResponse) Reset() {
	*x = UserCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateResponse) ProtoMessage() {}

func (x *UserCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateResponse.ProtoReflect.Descriptor instead.
func (*UserCreateResponse) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{1}
}

func (x *UserCreateResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   *UUID  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	RoleId   *UUID  `protobuf:"bytes,5,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	RoleName string `protobuf:"bytes,6,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
}

func (x *UserUpdateRequest) Reset() {
	*x = UserUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdateRequest) ProtoMessage() {}

func (x *UserUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdateRequest.ProtoReflect.Descriptor instead.
func (*UserUpdateRequest) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{2}
}

func (x *UserUpdateRequest) GetUserId() *UUID {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *UserUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserUpdateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserUpdateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserUpdateRequest) GetRoleId() *UUID {
	if x != nil {
		return x.RoleId
	}
	return nil
}

func (x *UserUpdateRequest) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

type UserUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserUpdateResponse) Reset() {
	*x = UserUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdateResponse) ProtoMessage() {}

func (x *UserUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdateResponse.ProtoReflect.Descriptor instead.
func (*UserUpdateResponse) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{3}
}

func (x *UserUpdateResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId *UUID `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserDeleteRequest) Reset() {
	*x = UserDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeleteRequest) ProtoMessage() {}

func (x *UserDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeleteRequest.ProtoReflect.Descriptor instead.
func (*UserDeleteRequest) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{4}
}

func (x *UserDeleteRequest) GetUserId() *UUID {
	if x != nil {
		return x.UserId
	}
	return nil
}

type UserDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserDeleteResponse) Reset() {
	*x = UserDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeleteResponse) ProtoMessage() {}

func (x *UserDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeleteResponse.ProtoReflect.Descriptor instead.
func (*UserDeleteResponse) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{5}
}

type UserFindOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId *UUID `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserFindOneRequest) Reset() {
	*x = UserFindOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFindOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFindOneRequest) ProtoMessage() {}

func (x *UserFindOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFindOneRequest.ProtoReflect.Descriptor instead.
func (*UserFindOneRequest) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{6}
}

func (x *UserFindOneRequest) GetUserId() *UUID {
	if x != nil {
		return x.UserId
	}
	return nil
}

type UserFindOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserFindOneResponse) Reset() {
	*x = UserFindOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFindOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFindOneResponse) ProtoMessage() {}

func (x *UserFindOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFindOneResponse.ProtoReflect.Descriptor instead.
func (*UserFindOneResponse) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{7}
}

func (x *UserFindOneResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserFindManyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UserFindManyResponse) Reset() {
	*x = UserFindManyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFindManyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFindManyResponse) ProtoMessage() {}

func (x *UserFindManyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFindManyResponse.ProtoReflect.Descriptor instead.
func (*UserFindManyResponse) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{8}
}

func (x *UserFindManyResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type UserChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      *UUID  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OldPassword string `protobuf:"bytes,2,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"`
	NewPassword string `protobuf:"bytes,3,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
}

func (x *UserChangePasswordRequest) Reset() {
	*x = UserChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserChangePasswordRequest) ProtoMessage() {}

func (x *UserChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*UserChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{9}
}

func (x *UserChangePasswordRequest) GetUserId() *UUID {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *UserChangePasswordRequest) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *UserChangePasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type UserChangePasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserChangePasswordResponse) Reset() {
	*x = UserChangePasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserChangePasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserChangePasswordResponse) ProtoMessage() {}

func (x *UserChangePasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserChangePasswordResponse.ProtoReflect.Descriptor instead.
func (*UserChangePasswordResponse) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{10}
}

type UserChangeUsernameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      *UUID  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OldUsername string `protobuf:"bytes,2,opt,name=old_username,json=oldUsername,proto3" json:"old_username,omitempty"`
	NewUsername string `protobuf:"bytes,3,opt,name=new_username,json=newUsername,proto3" json:"new_username,omitempty"`
}

func (x *UserChangeUsernameRequest) Reset() {
	*x = UserChangeUsernameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserChangeUsernameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserChangeUsernameRequest) ProtoMessage() {}

func (x *UserChangeUsernameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserChangeUsernameRequest.ProtoReflect.Descriptor instead.
func (*UserChangeUsernameRequest) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{11}
}

func (x *UserChangeUsernameRequest) GetUserId() *UUID {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *UserChangeUsernameRequest) GetOldUsername() string {
	if x != nil {
		return x.OldUsername
	}
	return ""
}

func (x *UserChangeUsernameRequest) GetNewUsername() string {
	if x != nil {
		return x.NewUsername
	}
	return ""
}

type UserChangeUsernameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserChangeUsernameResponse) Reset() {
	*x = UserChangeUsernameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserChangeUsernameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserChangeUsernameResponse) ProtoMessage() {}

func (x *UserChangeUsernameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserChangeUsernameResponse.ProtoReflect.Descriptor instead.
func (*UserChangeUsernameResponse) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{12}
}

type UserChangeEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   *UUID  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OldEmail string `protobuf:"bytes,2,opt,name=old_email,json=oldEmail,proto3" json:"old_email,omitempty"`
	NewEmail string `protobuf:"bytes,3,opt,name=new_email,json=newEmail,proto3" json:"new_email,omitempty"`
}

func (x *UserChangeEmailRequest) Reset() {
	*x = UserChangeEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserChangeEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserChangeEmailRequest) ProtoMessage() {}

func (x *UserChangeEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserChangeEmailRequest.ProtoReflect.Descriptor instead.
func (*UserChangeEmailRequest) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{13}
}

func (x *UserChangeEmailRequest) GetUserId() *UUID {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *UserChangeEmailRequest) GetOldEmail() string {
	if x != nil {
		return x.OldEmail
	}
	return ""
}

func (x *UserChangeEmailRequest) GetNewEmail() string {
	if x != nil {
		return x.NewEmail
	}
	return ""
}

type UserChangeEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserChangeEmailResponse) Reset() {
	*x = UserChangeEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserChangeEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserChangeEmailResponse) ProtoMessage() {}

func (x *UserChangeEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserChangeEmailResponse.ProtoReflect.Descriptor instead.
func (*UserChangeEmailResponse) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{14}
}

var File_user_api_proto protoreflect.FileDescriptor

var file_user_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01,
	0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x07, 0x72, 0x6f, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x55, 0x49,
	0x44, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xb6, 0x01, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x06, 0x72, 0x6f, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x2f, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x33, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x0a, 0x12,
	0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x30, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x22, 0x33, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64,
	0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x19, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x6c, 0x64, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f,
	0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65,
	0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x1c, 0x0a,
	0x1a, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x19,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x55, 0x49,
	0x44, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x6c, 0x64,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6f, 0x6c, 0x64, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x6e, 0x65, 0x77, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x1c, 0x0a, 0x1a, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x72, 0x0a,
	0x16, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6c, 0x64, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d,
	0x2e, 0x2f, 0x67, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_api_proto_rawDescOnce sync.Once
	file_user_api_proto_rawDescData = file_user_api_proto_rawDesc
)

func file_user_api_proto_rawDescGZIP() []byte {
	file_user_api_proto_rawDescOnce.Do(func() {
		file_user_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_api_proto_rawDescData)
	})
	return file_user_api_proto_rawDescData
}

var file_user_api_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_user_api_proto_goTypes = []interface{}{
	(*UserCreateRequest)(nil),          // 0: UserCreateRequest
	(*UserCreateResponse)(nil),         // 1: UserCreateResponse
	(*UserUpdateRequest)(nil),          // 2: UserUpdateRequest
	(*UserUpdateResponse)(nil),         // 3: UserUpdateResponse
	(*UserDeleteRequest)(nil),          // 4: UserDeleteRequest
	(*UserDeleteResponse)(nil),         // 5: UserDeleteResponse
	(*UserFindOneRequest)(nil),         // 6: UserFindOneRequest
	(*UserFindOneResponse)(nil),        // 7: UserFindOneResponse
	(*UserFindManyResponse)(nil),       // 8: UserFindManyResponse
	(*UserChangePasswordRequest)(nil),  // 9: UserChangePasswordRequest
	(*UserChangePasswordResponse)(nil), // 10: UserChangePasswordResponse
	(*UserChangeUsernameRequest)(nil),  // 11: UserChangeUsernameRequest
	(*UserChangeUsernameResponse)(nil), // 12: UserChangeUsernameResponse
	(*UserChangeEmailRequest)(nil),     // 13: UserChangeEmailRequest
	(*UserChangeEmailResponse)(nil),    // 14: UserChangeEmailResponse
	(*UUID)(nil),                       // 15: UUID
	(*User)(nil),                       // 16: User
}
var file_user_api_proto_depIdxs = []int32{
	15, // 0: UserCreateRequest.role_id:type_name -> UUID
	16, // 1: UserCreateResponse.user:type_name -> User
	15, // 2: UserUpdateRequest.user_id:type_name -> UUID
	15, // 3: UserUpdateRequest.role_id:type_name -> UUID
	16, // 4: UserUpdateResponse.user:type_name -> User
	15, // 5: UserDeleteRequest.user_id:type_name -> UUID
	15, // 6: UserFindOneRequest.user_id:type_name -> UUID
	16, // 7: UserFindOneResponse.user:type_name -> User
	16, // 8: UserFindManyResponse.users:type_name -> User
	15, // 9: UserChangePasswordRequest.user_id:type_name -> UUID
	15, // 10: UserChangeUsernameRequest.user_id:type_name -> UUID
	15, // 11: UserChangeEmailRequest.user_id:type_name -> UUID
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_user_api_proto_init() }
func file_user_api_proto_init() {
	if File_user_api_proto != nil {
		return
	}
	file_generic_proto_init()
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreateRequest); i {
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
		file_user_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreateResponse); i {
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
		file_user_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdateRequest); i {
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
		file_user_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdateResponse); i {
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
		file_user_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDeleteRequest); i {
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
		file_user_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDeleteResponse); i {
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
		file_user_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFindOneRequest); i {
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
		file_user_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFindOneResponse); i {
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
		file_user_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFindManyResponse); i {
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
		file_user_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChangePasswordRequest); i {
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
		file_user_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChangePasswordResponse); i {
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
		file_user_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChangeUsernameRequest); i {
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
		file_user_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChangeUsernameResponse); i {
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
		file_user_api_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChangeEmailRequest); i {
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
		file_user_api_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChangeEmailResponse); i {
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
			RawDescriptor: file_user_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_api_proto_goTypes,
		DependencyIndexes: file_user_api_proto_depIdxs,
		MessageInfos:      file_user_api_proto_msgTypes,
	}.Build()
	File_user_api_proto = out.File
	file_user_api_proto_rawDesc = nil
	file_user_api_proto_goTypes = nil
	file_user_api_proto_depIdxs = nil
}
