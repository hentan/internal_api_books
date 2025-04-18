// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: authors/authors.proto

package authors

import (
	books "github.com/hentan/internal_api_books/gen/go/books"
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

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NameAuthor    string `protobuf:"bytes,2,opt,name=name_author,json=nameAuthor,proto3" json:"name_author,omitempty"`
	SurnameAuthor string `protobuf:"bytes,3,opt,name=surname_author,json=surnameAuthor,proto3" json:"surname_author,omitempty"`
	Biography     string `protobuf:"bytes,4,opt,name=biography,proto3" json:"biography,omitempty"`
	Birthday      string `protobuf:"bytes,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_authors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_authors_authors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_authors_authors_proto_rawDescGZIP(), []int{0}
}

func (x *Author) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Author) GetNameAuthor() string {
	if x != nil {
		return x.NameAuthor
	}
	return ""
}

func (x *Author) GetSurnameAuthor() string {
	if x != nil {
		return x.SurnameAuthor
	}
	return ""
}

func (x *Author) GetBiography() string {
	if x != nil {
		return x.Biography
	}
	return ""
}

func (x *Author) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

type AuthorList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authors []*Author `protobuf:"bytes,1,rep,name=Authors,proto3" json:"Authors,omitempty"`
}

func (x *AuthorList) Reset() {
	*x = AuthorList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_authors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorList) ProtoMessage() {}

func (x *AuthorList) ProtoReflect() protoreflect.Message {
	mi := &file_authors_authors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorList.ProtoReflect.Descriptor instead.
func (*AuthorList) Descriptor() ([]byte, []int) {
	return file_authors_authors_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorList) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

type AuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AuthorRequest) Reset() {
	*x = AuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_authors_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorRequest) ProtoMessage() {}

func (x *AuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authors_authors_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorRequest.ProtoReflect.Descriptor instead.
func (*AuthorRequest) Descriptor() ([]byte, []int) {
	return file_authors_authors_proto_rawDescGZIP(), []int{2}
}

func (x *AuthorRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NameAuthor    string `protobuf:"bytes,1,opt,name=name_author,json=nameAuthor,proto3" json:"name_author,omitempty"`
	SurnameAuthor string `protobuf:"bytes,2,opt,name=surname_author,json=surnameAuthor,proto3" json:"surname_author,omitempty"`
	Biography     string `protobuf:"bytes,3,opt,name=biography,proto3" json:"biography,omitempty"`
	Birthday      string `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
}

func (x *CreateAuthorRequest) Reset() {
	*x = CreateAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_authors_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthorRequest) ProtoMessage() {}

func (x *CreateAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authors_authors_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthorRequest.ProtoReflect.Descriptor instead.
func (*CreateAuthorRequest) Descriptor() ([]byte, []int) {
	return file_authors_authors_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAuthorRequest) GetNameAuthor() string {
	if x != nil {
		return x.NameAuthor
	}
	return ""
}

func (x *CreateAuthorRequest) GetSurnameAuthor() string {
	if x != nil {
		return x.SurnameAuthor
	}
	return ""
}

func (x *CreateAuthorRequest) GetBiography() string {
	if x != nil {
		return x.Biography
	}
	return ""
}

func (x *CreateAuthorRequest) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

type UpdateAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NameAuthor    string `protobuf:"bytes,2,opt,name=name_author,json=nameAuthor,proto3" json:"name_author,omitempty"`
	SurnameAuthor string `protobuf:"bytes,3,opt,name=surname_author,json=surnameAuthor,proto3" json:"surname_author,omitempty"`
	Biography     string `protobuf:"bytes,4,opt,name=biography,proto3" json:"biography,omitempty"`
	Birthday      string `protobuf:"bytes,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
}

func (x *UpdateAuthorRequest) Reset() {
	*x = UpdateAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_authors_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuthorRequest) ProtoMessage() {}

func (x *UpdateAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authors_authors_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuthorRequest.ProtoReflect.Descriptor instead.
func (*UpdateAuthorRequest) Descriptor() ([]byte, []int) {
	return file_authors_authors_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAuthorRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateAuthorRequest) GetNameAuthor() string {
	if x != nil {
		return x.NameAuthor
	}
	return ""
}

func (x *UpdateAuthorRequest) GetSurnameAuthor() string {
	if x != nil {
		return x.SurnameAuthor
	}
	return ""
}

func (x *UpdateAuthorRequest) GetBiography() string {
	if x != nil {
		return x.Biography
	}
	return ""
}

func (x *UpdateAuthorRequest) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_authors_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authors_authors_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_authors_authors_proto_rawDescGZIP(), []int{5}
}

func (x *StatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateAuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author *Author `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *CreateAuthorResponse) Reset() {
	*x = CreateAuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_authors_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthorResponse) ProtoMessage() {}

func (x *CreateAuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authors_authors_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthorResponse.ProtoReflect.Descriptor instead.
func (*CreateAuthorResponse) Descriptor() ([]byte, []int) {
	return file_authors_authors_proto_rawDescGZIP(), []int{6}
}

func (x *CreateAuthorResponse) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

var File_authors_authors_proto protoreflect.FileDescriptor

var file_authors_authors_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73,
	0x1a, 0x11, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x25, 0x0a, 0x0e, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69, 0x6f, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x6f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79,
	0x22, 0x37, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29,
	0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x52, 0x07, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x22, 0x1f, 0x0a, 0x0d, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69,
	0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x22, 0xa7, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x22, 0x2a,
	0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3f, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x32, 0xd2, 0x02, 0x0a, 0x0d,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x0c,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x38, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x4b, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x65, 0x6e, 0x74, 0x61, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x61,
	0x70, 0x69, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authors_authors_proto_rawDescOnce sync.Once
	file_authors_authors_proto_rawDescData = file_authors_authors_proto_rawDesc
)

func file_authors_authors_proto_rawDescGZIP() []byte {
	file_authors_authors_proto_rawDescOnce.Do(func() {
		file_authors_authors_proto_rawDescData = protoimpl.X.CompressGZIP(file_authors_authors_proto_rawDescData)
	})
	return file_authors_authors_proto_rawDescData
}

var file_authors_authors_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_authors_authors_proto_goTypes = []interface{}{
	(*Author)(nil),               // 0: authors.Author
	(*AuthorList)(nil),           // 1: authors.AuthorList
	(*AuthorRequest)(nil),        // 2: authors.AuthorRequest
	(*CreateAuthorRequest)(nil),  // 3: authors.CreateAuthorRequest
	(*UpdateAuthorRequest)(nil),  // 4: authors.UpdateAuthorRequest
	(*StatusResponse)(nil),       // 5: authors.StatusResponse
	(*CreateAuthorResponse)(nil), // 6: authors.CreateAuthorResponse
	(*books.Empty)(nil),          // 7: books.Empty
}
var file_authors_authors_proto_depIdxs = []int32{
	0, // 0: authors.AuthorList.Authors:type_name -> authors.Author
	0, // 1: authors.CreateAuthorResponse.author:type_name -> authors.Author
	7, // 2: authors.AuthorService.GetAllAuthors:input_type -> books.Empty
	2, // 3: authors.AuthorService.GetAuthorById:input_type -> authors.AuthorRequest
	3, // 4: authors.AuthorService.CreateAuthor:input_type -> authors.CreateAuthorRequest
	4, // 5: authors.AuthorService.UpdateAuthor:input_type -> authors.UpdateAuthorRequest
	2, // 6: authors.AuthorService.DeleteAuthor:input_type -> authors.AuthorRequest
	1, // 7: authors.AuthorService.GetAllAuthors:output_type -> authors.AuthorList
	0, // 8: authors.AuthorService.GetAuthorById:output_type -> authors.Author
	6, // 9: authors.AuthorService.CreateAuthor:output_type -> authors.CreateAuthorResponse
	5, // 10: authors.AuthorService.UpdateAuthor:output_type -> authors.StatusResponse
	5, // 11: authors.AuthorService.DeleteAuthor:output_type -> authors.StatusResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_authors_authors_proto_init() }
func file_authors_authors_proto_init() {
	if File_authors_authors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authors_authors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_authors_authors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorList); i {
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
		file_authors_authors_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorRequest); i {
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
		file_authors_authors_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthorRequest); i {
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
		file_authors_authors_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAuthorRequest); i {
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
		file_authors_authors_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_authors_authors_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthorResponse); i {
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
			RawDescriptor: file_authors_authors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authors_authors_proto_goTypes,
		DependencyIndexes: file_authors_authors_proto_depIdxs,
		MessageInfos:      file_authors_authors_proto_msgTypes,
	}.Build()
	File_authors_authors_proto = out.File
	file_authors_authors_proto_rawDesc = nil
	file_authors_authors_proto_goTypes = nil
	file_authors_authors_proto_depIdxs = nil
}
