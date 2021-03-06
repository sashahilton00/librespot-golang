// Code generated by protoc-gen-go. DO NOT EDIT.
// source: playlist4content.proto

package Spotify

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Item struct {
	Uri              *string         `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	Attributes       *ItemAttributes `protobuf:"bytes,2,opt,name=attributes" json:"attributes,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *Item) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *Item) GetAttributes() *ItemAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type ListItems struct {
	Pos              *int32  `protobuf:"varint,1,opt,name=pos" json:"pos,omitempty"`
	Truncated        *bool   `protobuf:"varint,2,opt,name=truncated" json:"truncated,omitempty"`
	Items            []*Item `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ListItems) Reset()                    { *m = ListItems{} }
func (m *ListItems) String() string            { return proto.CompactTextString(m) }
func (*ListItems) ProtoMessage()               {}
func (*ListItems) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *ListItems) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func (m *ListItems) GetTruncated() bool {
	if m != nil && m.Truncated != nil {
		return *m.Truncated
	}
	return false
}

func (m *ListItems) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type ContentRange struct {
	Pos              *int32 `protobuf:"varint,1,opt,name=pos" json:"pos,omitempty"`
	Length           *int32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ContentRange) Reset()                    { *m = ContentRange{} }
func (m *ContentRange) String() string            { return proto.CompactTextString(m) }
func (*ContentRange) ProtoMessage()               {}
func (*ContentRange) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *ContentRange) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func (m *ContentRange) GetLength() int32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

type ListContentSelection struct {
	WantRevision          *bool                  `protobuf:"varint,1,opt,name=wantRevision" json:"wantRevision,omitempty"`
	WantLength            *bool                  `protobuf:"varint,2,opt,name=wantLength" json:"wantLength,omitempty"`
	WantAttributes        *bool                  `protobuf:"varint,3,opt,name=wantAttributes" json:"wantAttributes,omitempty"`
	WantChecksum          *bool                  `protobuf:"varint,4,opt,name=wantChecksum" json:"wantChecksum,omitempty"`
	WantContent           *bool                  `protobuf:"varint,5,opt,name=wantContent" json:"wantContent,omitempty"`
	ContentRange          *ContentRange          `protobuf:"bytes,6,opt,name=contentRange" json:"contentRange,omitempty"`
	WantDiff              *bool                  `protobuf:"varint,7,opt,name=wantDiff" json:"wantDiff,omitempty"`
	BaseRevision          []byte                 `protobuf:"bytes,8,opt,name=baseRevision" json:"baseRevision,omitempty"`
	HintRevision          []byte                 `protobuf:"bytes,9,opt,name=hintRevision" json:"hintRevision,omitempty"`
	WantNothingIfUpToDate *bool                  `protobuf:"varint,10,opt,name=wantNothingIfUpToDate" json:"wantNothingIfUpToDate,omitempty"`
	WantResolveAction     *bool                  `protobuf:"varint,12,opt,name=wantResolveAction" json:"wantResolveAction,omitempty"`
	Issues                []*ClientIssue         `protobuf:"bytes,13,rep,name=issues" json:"issues,omitempty"`
	ResolveAction         []*ClientResolveAction `protobuf:"bytes,14,rep,name=resolveAction" json:"resolveAction,omitempty"`
	XXX_unrecognized      []byte                 `json:"-"`
}

func (m *ListContentSelection) Reset()                    { *m = ListContentSelection{} }
func (m *ListContentSelection) String() string            { return proto.CompactTextString(m) }
func (*ListContentSelection) ProtoMessage()               {}
func (*ListContentSelection) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *ListContentSelection) GetWantRevision() bool {
	if m != nil && m.WantRevision != nil {
		return *m.WantRevision
	}
	return false
}

func (m *ListContentSelection) GetWantLength() bool {
	if m != nil && m.WantLength != nil {
		return *m.WantLength
	}
	return false
}

func (m *ListContentSelection) GetWantAttributes() bool {
	if m != nil && m.WantAttributes != nil {
		return *m.WantAttributes
	}
	return false
}

func (m *ListContentSelection) GetWantChecksum() bool {
	if m != nil && m.WantChecksum != nil {
		return *m.WantChecksum
	}
	return false
}

func (m *ListContentSelection) GetWantContent() bool {
	if m != nil && m.WantContent != nil {
		return *m.WantContent
	}
	return false
}

func (m *ListContentSelection) GetContentRange() *ContentRange {
	if m != nil {
		return m.ContentRange
	}
	return nil
}

func (m *ListContentSelection) GetWantDiff() bool {
	if m != nil && m.WantDiff != nil {
		return *m.WantDiff
	}
	return false
}

func (m *ListContentSelection) GetBaseRevision() []byte {
	if m != nil {
		return m.BaseRevision
	}
	return nil
}

func (m *ListContentSelection) GetHintRevision() []byte {
	if m != nil {
		return m.HintRevision
	}
	return nil
}

func (m *ListContentSelection) GetWantNothingIfUpToDate() bool {
	if m != nil && m.WantNothingIfUpToDate != nil {
		return *m.WantNothingIfUpToDate
	}
	return false
}

func (m *ListContentSelection) GetWantResolveAction() bool {
	if m != nil && m.WantResolveAction != nil {
		return *m.WantResolveAction
	}
	return false
}

func (m *ListContentSelection) GetIssues() []*ClientIssue {
	if m != nil {
		return m.Issues
	}
	return nil
}

func (m *ListContentSelection) GetResolveAction() []*ClientResolveAction {
	if m != nil {
		return m.ResolveAction
	}
	return nil
}

func init() {
	proto.RegisterType((*Item)(nil), "Spotify.Item")
	proto.RegisterType((*ListItems)(nil), "Spotify.ListItems")
	proto.RegisterType((*ContentRange)(nil), "Spotify.ContentRange")
	proto.RegisterType((*ListContentSelection)(nil), "Spotify.ListContentSelection")
}

func init() { proto.RegisterFile("playlist4content.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x5b, 0x6b, 0xdb, 0x30,
	0x14, 0x26, 0xcb, 0xa5, 0xce, 0x89, 0x53, 0x36, 0xad, 0xe9, 0x44, 0x28, 0xc3, 0x78, 0x30, 0xf2,
	0x50, 0xf2, 0x50, 0x0a, 0xdb, 0x1e, 0xbb, 0xf6, 0x25, 0x50, 0x06, 0x53, 0xb7, 0xf7, 0xa9, 0xde,
	0x49, 0x22, 0xe6, 0x48, 0xc6, 0x3a, 0xee, 0xe8, 0x8f, 0xdb, 0x7f, 0x1b, 0x92, 0x52, 0x5f, 0xda,
	0xbc, 0xf9, 0x7c, 0xe7, 0xbb, 0xc8, 0x9f, 0x04, 0xa7, 0x45, 0x2e, 0x1f, 0x73, 0x65, 0xe9, 0x32,
	0x33, 0x9a, 0x50, 0xd3, 0xb2, 0x28, 0x0d, 0x19, 0x76, 0x74, 0x57, 0x18, 0x52, 0xeb, 0xc7, 0xf9,
	0xdb, 0x9a, 0xb0, 0x43, 0x92, 0x61, 0x3b, 0x9f, 0xd5, 0xa0, 0xb2, 0xb6, 0x42, 0x1b, 0xe0, 0xf4,
	0x3b, 0x0c, 0x56, 0x84, 0x3b, 0xf6, 0x1a, 0xfa, 0x55, 0xa9, 0x78, 0x2f, 0xe9, 0x2d, 0xc6, 0xc2,
	0x7d, 0xb2, 0x4f, 0x00, 0x92, 0xa8, 0x54, 0xf7, 0x15, 0xa1, 0xe5, 0xaf, 0x92, 0xde, 0x62, 0x72,
	0xf1, 0x6e, 0xb9, 0xcf, 0x58, 0x3a, 0xd1, 0x55, 0xbd, 0x16, 0x2d, 0x6a, 0xfa, 0x0b, 0xc6, 0xb7,
	0xca, 0x92, 0x63, 0x58, 0xe7, 0x5b, 0x18, 0xeb, 0x7d, 0x87, 0xc2, 0x7d, 0xb2, 0x33, 0x18, 0x53,
	0x59, 0xe9, 0x4c, 0x12, 0xfe, 0xf6, 0xb6, 0x91, 0x68, 0x00, 0xf6, 0x01, 0x86, 0xca, 0x09, 0x79,
	0x3f, 0xe9, 0x2f, 0x26, 0x17, 0xd3, 0x4e, 0xa0, 0x08, 0xbb, 0xf4, 0x33, 0xc4, 0xd7, 0xe1, 0xd7,
	0x85, 0xd4, 0x1b, 0x3c, 0x10, 0x72, 0x0a, 0xa3, 0x1c, 0xf5, 0x86, 0xb6, 0x3e, 0x61, 0x28, 0xf6,
	0x53, 0xfa, 0x6f, 0x00, 0x27, 0xee, 0x70, 0x7b, 0xf9, 0x1d, 0xe6, 0x98, 0x91, 0x32, 0x9a, 0xa5,
	0x10, 0xff, 0x95, 0x9a, 0x04, 0x3e, 0x28, 0xab, 0x8c, 0xf6, 0x5e, 0x91, 0xe8, 0x60, 0xec, 0x3d,
	0x80, 0x9b, 0x6f, 0x1b, 0xe3, 0x48, 0xb4, 0x10, 0xf6, 0x11, 0x8e, 0xdd, 0xd4, 0xd4, 0xc2, 0xfb,
	0x9e, 0xf3, 0x0c, 0x7d, 0xca, 0xba, 0xde, 0x62, 0xf6, 0xc7, 0x56, 0x3b, 0x3e, 0x68, 0xb2, 0x9e,
	0x30, 0x96, 0xc0, 0xc4, 0xcf, 0xe1, 0x9c, 0x7c, 0xe8, 0x29, 0x6d, 0x88, 0x7d, 0x81, 0x38, 0x6b,
	0x95, 0xc0, 0x47, 0xfe, 0x86, 0x66, 0x75, 0x61, 0xed, 0x86, 0x44, 0x87, 0xca, 0xe6, 0x10, 0x39,
	0xa7, 0x1b, 0xb5, 0x5e, 0xf3, 0x23, 0xef, 0x5c, 0xcf, 0xee, 0x70, 0xf7, 0xd2, 0x62, 0x5d, 0x44,
	0x94, 0xf4, 0x16, 0xb1, 0xe8, 0x60, 0x8e, 0xb3, 0x55, 0xad, 0xb2, 0xc6, 0x81, 0xd3, 0xc6, 0xd8,
	0x25, 0xcc, 0x9c, 0xe7, 0x37, 0x43, 0x5b, 0xa5, 0x37, 0xab, 0xf5, 0xcf, 0xe2, 0x87, 0xb9, 0x91,
	0x84, 0x1c, 0x7c, 0xe0, 0xe1, 0x25, 0x3b, 0x87, 0x37, 0xa1, 0x72, 0x6b, 0xf2, 0x07, 0xbc, 0xf2,
	0x77, 0xc3, 0x63, 0xaf, 0x78, 0xb9, 0x60, 0xe7, 0x30, 0x0a, 0x8f, 0x99, 0x4f, 0xfd, 0x6b, 0x39,
	0x69, 0x7e, 0x3e, 0x57, 0xa8, 0x69, 0xe5, 0x96, 0x62, 0xcf, 0x61, 0x5f, 0x61, 0x5a, 0x76, 0x7c,
	0x8f, 0xbd, 0xe8, 0xec, 0x99, 0xa8, 0x13, 0x21, 0xba, 0x92, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xea, 0x4e, 0xf4, 0x8f, 0x7c, 0x03, 0x00, 0x00,
}
