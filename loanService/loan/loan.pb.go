// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: loan.proto

package loan

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

// Request message for ApplyLoan
type ApplyLoanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Amount   float32 `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Duration int32   `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *ApplyLoanRequest) Reset() {
	*x = ApplyLoanRequest{}
	mi := &file_loan_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApplyLoanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyLoanRequest) ProtoMessage() {}

func (x *ApplyLoanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loan_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyLoanRequest.ProtoReflect.Descriptor instead.
func (*ApplyLoanRequest) Descriptor() ([]byte, []int) {
	return file_loan_proto_rawDescGZIP(), []int{0}
}

func (x *ApplyLoanRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ApplyLoanRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ApplyLoanRequest) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

// Response message for ApplyLoan
type ApplyLoanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status     bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	StatusCode int32  `protobuf:"varint,3,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	LoandId    string `protobuf:"bytes,4,opt,name=loandId,proto3" json:"loandId,omitempty"`
}

func (x *ApplyLoanResponse) Reset() {
	*x = ApplyLoanResponse{}
	mi := &file_loan_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApplyLoanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyLoanResponse) ProtoMessage() {}

func (x *ApplyLoanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loan_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyLoanResponse.ProtoReflect.Descriptor instead.
func (*ApplyLoanResponse) Descriptor() ([]byte, []int) {
	return file_loan_proto_rawDescGZIP(), []int{1}
}

func (x *ApplyLoanResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ApplyLoanResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ApplyLoanResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ApplyLoanResponse) GetLoandId() string {
	if x != nil {
		return x.LoandId
	}
	return ""
}

type ApproveLoanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoanId         string  `protobuf:"bytes,1,opt,name=loanId,proto3" json:"loanId,omitempty"`
	ApprovedAmount float32 `protobuf:"fixed32,2,opt,name=approvedAmount,proto3" json:"approvedAmount,omitempty"`
	UserId         string  `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *ApproveLoanRequest) Reset() {
	*x = ApproveLoanRequest{}
	mi := &file_loan_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApproveLoanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveLoanRequest) ProtoMessage() {}

func (x *ApproveLoanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loan_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveLoanRequest.ProtoReflect.Descriptor instead.
func (*ApproveLoanRequest) Descriptor() ([]byte, []int) {
	return file_loan_proto_rawDescGZIP(), []int{2}
}

func (x *ApproveLoanRequest) GetLoanId() string {
	if x != nil {
		return x.LoanId
	}
	return ""
}

func (x *ApproveLoanRequest) GetApprovedAmount() float32 {
	if x != nil {
		return x.ApprovedAmount
	}
	return 0
}

func (x *ApproveLoanRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ApproveLoanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status     bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	StatusCode int32  `protobuf:"varint,3,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
}

func (x *ApproveLoanResponse) Reset() {
	*x = ApproveLoanResponse{}
	mi := &file_loan_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApproveLoanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveLoanResponse) ProtoMessage() {}

func (x *ApproveLoanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loan_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveLoanResponse.ProtoReflect.Descriptor instead.
func (*ApproveLoanResponse) Descriptor() ([]byte, []int) {
	return file_loan_proto_rawDescGZIP(), []int{3}
}

func (x *ApproveLoanResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ApproveLoanResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ApproveLoanResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

type RejectLoanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoanId string `protobuf:"bytes,1,opt,name=loanId,proto3" json:"loanId,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *RejectLoanRequest) Reset() {
	*x = RejectLoanRequest{}
	mi := &file_loan_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RejectLoanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectLoanRequest) ProtoMessage() {}

func (x *RejectLoanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loan_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectLoanRequest.ProtoReflect.Descriptor instead.
func (*RejectLoanRequest) Descriptor() ([]byte, []int) {
	return file_loan_proto_rawDescGZIP(), []int{4}
}

func (x *RejectLoanRequest) GetLoanId() string {
	if x != nil {
		return x.LoanId
	}
	return ""
}

func (x *RejectLoanRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type RejectLoanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status     bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	StatusCode int32  `protobuf:"varint,3,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
}

func (x *RejectLoanResponse) Reset() {
	*x = RejectLoanResponse{}
	mi := &file_loan_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RejectLoanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectLoanResponse) ProtoMessage() {}

func (x *RejectLoanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loan_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectLoanResponse.ProtoReflect.Descriptor instead.
func (*RejectLoanResponse) Descriptor() ([]byte, []int) {
	return file_loan_proto_rawDescGZIP(), []int{5}
}

func (x *RejectLoanResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RejectLoanResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *RejectLoanResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

var File_loan_proto protoreflect.FileDescriptor

var file_loan_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x10,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x6f, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7f, 0x0a, 0x11,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x6f, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x6c, 0x0a,
	0x12, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x61,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x13, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x43, 0x0a, 0x11, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f,
	0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x61,
	0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x61, 0x6e, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x66, 0x0a, 0x12, 0x52, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0xb2, 0x01, 0x0a, 0x0b, 0x4c, 0x6f, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x32, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x6f, 0x61, 0x6e, 0x12, 0x11,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x6f, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x6f, 0x61, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x4c, 0x6f, 0x61, 0x6e, 0x12, 0x13, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x4c, 0x6f,
	0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x0a, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x12, 0x12, 0x2e,
	0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x6e, 0x6c, 0x69, 0x6b, 0x65, 0x68, 0x65, 0x6e, 0x72,
	0x79, 0x79, 0x2f, 0x6c, 0x6f, 0x61, 0x6e, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6c,
	0x6f, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x61, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loan_proto_rawDescOnce sync.Once
	file_loan_proto_rawDescData = file_loan_proto_rawDesc
)

func file_loan_proto_rawDescGZIP() []byte {
	file_loan_proto_rawDescOnce.Do(func() {
		file_loan_proto_rawDescData = protoimpl.X.CompressGZIP(file_loan_proto_rawDescData)
	})
	return file_loan_proto_rawDescData
}

var file_loan_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_loan_proto_goTypes = []any{
	(*ApplyLoanRequest)(nil),    // 0: ApplyLoanRequest
	(*ApplyLoanResponse)(nil),   // 1: ApplyLoanResponse
	(*ApproveLoanRequest)(nil),  // 2: ApproveLoanRequest
	(*ApproveLoanResponse)(nil), // 3: ApproveLoanResponse
	(*RejectLoanRequest)(nil),   // 4: RejectLoanRequest
	(*RejectLoanResponse)(nil),  // 5: RejectLoanResponse
}
var file_loan_proto_depIdxs = []int32{
	0, // 0: LoanService.ApplyLoan:input_type -> ApplyLoanRequest
	2, // 1: LoanService.ApproveLoan:input_type -> ApproveLoanRequest
	4, // 2: LoanService.RejectLoan:input_type -> RejectLoanRequest
	1, // 3: LoanService.ApplyLoan:output_type -> ApplyLoanResponse
	3, // 4: LoanService.ApproveLoan:output_type -> ApproveLoanResponse
	5, // 5: LoanService.RejectLoan:output_type -> RejectLoanResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_loan_proto_init() }
func file_loan_proto_init() {
	if File_loan_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_loan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_loan_proto_goTypes,
		DependencyIndexes: file_loan_proto_depIdxs,
		MessageInfos:      file_loan_proto_msgTypes,
	}.Build()
	File_loan_proto = out.File
	file_loan_proto_rawDesc = nil
	file_loan_proto_goTypes = nil
	file_loan_proto_depIdxs = nil
}