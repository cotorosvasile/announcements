// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: announce.proto

package announcements

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

type Announcement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Categories *Announcement_Categories `protobuf:"bytes,2,opt,name=categories,proto3" json:"categories,omitempty"`
	Title      *Announcement_Title      `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Type       string                   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Posted     float64                  `protobuf:"fixed64,5,opt,name=posted,proto3" json:"posted,omitempty"`
}

func (x *Announcement) Reset() {
	*x = Announcement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_announce_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Announcement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Announcement) ProtoMessage() {}

func (x *Announcement) ProtoReflect() protoreflect.Message {
	mi := &file_announce_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Announcement.ProtoReflect.Descriptor instead.
func (*Announcement) Descriptor() ([]byte, []int) {
	return file_announce_proto_rawDescGZIP(), []int{0}
}

func (x *Announcement) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Announcement) GetCategories() *Announcement_Categories {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *Announcement) GetTitle() *Announcement_Title {
	if x != nil {
		return x.Title
	}
	return nil
}

func (x *Announcement) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Announcement) GetPosted() float64 {
	if x != nil {
		return x.Posted
	}
	return 0
}

type AnnouncementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response float64 `protobuf:"fixed64,5,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *AnnouncementResponse) Reset() {
	*x = AnnouncementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_announce_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnouncementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnouncementResponse) ProtoMessage() {}

func (x *AnnouncementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_announce_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnouncementResponse.ProtoReflect.Descriptor instead.
func (*AnnouncementResponse) Descriptor() ([]byte, []int) {
	return file_announce_proto_rawDescGZIP(), []int{1}
}

func (x *AnnouncementResponse) GetResponse() float64 {
	if x != nil {
		return x.Response
	}
	return 0
}

type GetAllAnnouncementsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllAnnouncementsRequest) Reset() {
	*x = GetAllAnnouncementsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_announce_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAnnouncementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAnnouncementsRequest) ProtoMessage() {}

func (x *GetAllAnnouncementsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_announce_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAnnouncementsRequest.ProtoReflect.Descriptor instead.
func (*GetAllAnnouncementsRequest) Descriptor() ([]byte, []int) {
	return file_announce_proto_rawDescGZIP(), []int{2}
}

type GetAllAnnouncementsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Announcements []*Announcement `protobuf:"bytes,1,rep,name=announcements,proto3" json:"announcements,omitempty"`
}

func (x *GetAllAnnouncementsResponse) Reset() {
	*x = GetAllAnnouncementsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_announce_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAnnouncementsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAnnouncementsResponse) ProtoMessage() {}

func (x *GetAllAnnouncementsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_announce_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAnnouncementsResponse.ProtoReflect.Descriptor instead.
func (*GetAllAnnouncementsResponse) Descriptor() ([]byte, []int) {
	return file_announce_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllAnnouncementsResponse) GetAnnouncements() []*Announcement {
	if x != nil {
		return x.Announcements
	}
	return nil
}

type Announcement_Categories struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subcategory string `protobuf:"bytes,1,opt,name=subcategory,proto3" json:"subcategory,omitempty"`
}

func (x *Announcement_Categories) Reset() {
	*x = Announcement_Categories{}
	if protoimpl.UnsafeEnabled {
		mi := &file_announce_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Announcement_Categories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Announcement_Categories) ProtoMessage() {}

func (x *Announcement_Categories) ProtoReflect() protoreflect.Message {
	mi := &file_announce_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Announcement_Categories.ProtoReflect.Descriptor instead.
func (*Announcement_Categories) Descriptor() ([]byte, []int) {
	return file_announce_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Announcement_Categories) GetSubcategory() string {
	if x != nil {
		return x.Subcategory
	}
	return ""
}

type Announcement_Title struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ro string `protobuf:"bytes,1,opt,name=ro,proto3" json:"ro,omitempty"`
	Ru string `protobuf:"bytes,2,opt,name=ru,proto3" json:"ru,omitempty"`
}

func (x *Announcement_Title) Reset() {
	*x = Announcement_Title{}
	if protoimpl.UnsafeEnabled {
		mi := &file_announce_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Announcement_Title) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Announcement_Title) ProtoMessage() {}

func (x *Announcement_Title) ProtoReflect() protoreflect.Message {
	mi := &file_announce_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Announcement_Title.ProtoReflect.Descriptor instead.
func (*Announcement_Title) Descriptor() ([]byte, []int) {
	return file_announce_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Announcement_Title) GetRo() string {
	if x != nil {
		return x.Ro
	}
	return ""
}

func (x *Announcement_Title) GetRu() string {
	if x != nil {
		return x.Ru
	}
	return ""
}

var File_announce_proto protoreflect.FileDescriptor

var file_announce_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0xa4, 0x02, 0x0a, 0x0c, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x46, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x1a, 0x2e, 0x0a,
	0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x73,
	0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x75, 0x62, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x27, 0x0a,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x72, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x75, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x72, 0x75, 0x22, 0x32, 0x0a, 0x14, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x60, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x61, 0x6e, 0x6e, 0x6f, 0x75,
	0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41,
	0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x61, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xe0, 0x01, 0x0a, 0x14, 0x41,
	0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x29, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x6e, 0x6e, 0x6f, 0x75,
	0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x30, 0x5a,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x74, 0x6f, 0x72, 0x6f, 0x73, 0x76, 0x61, 0x73, 0x69, 0x6c,
	0x65, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_announce_proto_rawDescOnce sync.Once
	file_announce_proto_rawDescData = file_announce_proto_rawDesc
)

func file_announce_proto_rawDescGZIP() []byte {
	file_announce_proto_rawDescOnce.Do(func() {
		file_announce_proto_rawDescData = protoimpl.X.CompressGZIP(file_announce_proto_rawDescData)
	})
	return file_announce_proto_rawDescData
}

var file_announce_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_announce_proto_goTypes = []interface{}{
	(*Announcement)(nil),                // 0: announcements.Announcement
	(*AnnouncementResponse)(nil),        // 1: announcements.AnnouncementResponse
	(*GetAllAnnouncementsRequest)(nil),  // 2: announcements.GetAllAnnouncementsRequest
	(*GetAllAnnouncementsResponse)(nil), // 3: announcements.GetAllAnnouncementsResponse
	(*Announcement_Categories)(nil),     // 4: announcements.Announcement.Categories
	(*Announcement_Title)(nil),          // 5: announcements.Announcement.Title
}
var file_announce_proto_depIdxs = []int32{
	4, // 0: announcements.Announcement.categories:type_name -> announcements.Announcement.Categories
	5, // 1: announcements.Announcement.title:type_name -> announcements.Announcement.Title
	0, // 2: announcements.GetAllAnnouncementsResponse.announcements:type_name -> announcements.Announcement
	0, // 3: announcements.AnnouncementsService.CreateAnnouncement:input_type -> announcements.Announcement
	2, // 4: announcements.AnnouncementsService.GetAllAnnouncements:input_type -> announcements.GetAllAnnouncementsRequest
	1, // 5: announcements.AnnouncementsService.CreateAnnouncement:output_type -> announcements.AnnouncementResponse
	3, // 6: announcements.AnnouncementsService.GetAllAnnouncements:output_type -> announcements.GetAllAnnouncementsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_announce_proto_init() }
func file_announce_proto_init() {
	if File_announce_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_announce_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Announcement); i {
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
		file_announce_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnouncementResponse); i {
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
		file_announce_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAnnouncementsRequest); i {
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
		file_announce_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAnnouncementsResponse); i {
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
		file_announce_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Announcement_Categories); i {
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
		file_announce_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Announcement_Title); i {
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
			RawDescriptor: file_announce_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_announce_proto_goTypes,
		DependencyIndexes: file_announce_proto_depIdxs,
		MessageInfos:      file_announce_proto_msgTypes,
	}.Build()
	File_announce_proto = out.File
	file_announce_proto_rawDesc = nil
	file_announce_proto_goTypes = nil
	file_announce_proto_depIdxs = nil
}