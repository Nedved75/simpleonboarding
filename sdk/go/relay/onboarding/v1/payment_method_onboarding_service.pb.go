// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: relay/onboarding/v1/payment_method_onboarding_service.proto

package v1

import (
	relay "github.com/Nedved75/simpleonboarding/sdk/go/relay"
	_ "github.com/Nedved75/simpleonboarding/sdk/go/relay/onboarding/v1/paymentmethods"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PaymentMethodOnboardingStatus int32

const (
	PaymentMethodOnboardingStatus_PENDING       PaymentMethodOnboardingStatus = 0
	PaymentMethodOnboardingStatus_ACTIVATED     PaymentMethodOnboardingStatus = 1
	PaymentMethodOnboardingStatus_NON_COMPLIANT PaymentMethodOnboardingStatus = 2
	PaymentMethodOnboardingStatus_DEACTIVATED   PaymentMethodOnboardingStatus = 3
	PaymentMethodOnboardingStatus_TERMINATED    PaymentMethodOnboardingStatus = 4
)

// Enum value maps for PaymentMethodOnboardingStatus.
var (
	PaymentMethodOnboardingStatus_name = map[int32]string{
		0: "PENDING",
		1: "ACTIVATED",
		2: "NON_COMPLIANT",
		3: "DEACTIVATED",
		4: "TERMINATED",
	}
	PaymentMethodOnboardingStatus_value = map[string]int32{
		"PENDING":       0,
		"ACTIVATED":     1,
		"NON_COMPLIANT": 2,
		"DEACTIVATED":   3,
		"TERMINATED":    4,
	}
)

func (x PaymentMethodOnboardingStatus) Enum() *PaymentMethodOnboardingStatus {
	p := new(PaymentMethodOnboardingStatus)
	*p = x
	return p
}

func (x PaymentMethodOnboardingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentMethodOnboardingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_relay_onboarding_v1_payment_method_onboarding_service_proto_enumTypes[0].Descriptor()
}

func (PaymentMethodOnboardingStatus) Type() protoreflect.EnumType {
	return &file_relay_onboarding_v1_payment_method_onboarding_service_proto_enumTypes[0]
}

func (x PaymentMethodOnboardingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentMethodOnboardingStatus.Descriptor instead.
func (PaymentMethodOnboardingStatus) EnumDescriptor() ([]byte, []int) {
	return file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDescGZIP(), []int{0}
}

type CompletelyNew struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// Types that are assignable to Paymentselection:
	//
	//	*CompletelyNew_InitReq
	//	*CompletelyNew_InitRes
	Paymentselection isCompletelyNew_Paymentselection `protobuf_oneof:"paymentselection"`
}

func (x *CompletelyNew) Reset() {
	*x = CompletelyNew{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletelyNew) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletelyNew) ProtoMessage() {}

func (x *CompletelyNew) ProtoReflect() protoreflect.Message {
	mi := &file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletelyNew.ProtoReflect.Descriptor instead.
func (*CompletelyNew) Descriptor() ([]byte, []int) {
	return file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDescGZIP(), []int{0}
}

func (x *CompletelyNew) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (m *CompletelyNew) GetPaymentselection() isCompletelyNew_Paymentselection {
	if m != nil {
		return m.Paymentselection
	}
	return nil
}

func (x *CompletelyNew) GetInitReq() string {
	if x, ok := x.GetPaymentselection().(*CompletelyNew_InitReq); ok {
		return x.InitReq
	}
	return ""
}

func (x *CompletelyNew) GetInitRes() string {
	if x, ok := x.GetPaymentselection().(*CompletelyNew_InitRes); ok {
		return x.InitRes
	}
	return ""
}

type isCompletelyNew_Paymentselection interface {
	isCompletelyNew_Paymentselection()
}

type CompletelyNew_InitReq struct {
	InitReq string `protobuf:"bytes,2,opt,name=InitReq,proto3,oneof"`
}

type CompletelyNew_InitRes struct {
	InitRes string `protobuf:"bytes,3,opt,name=InitRes,proto3,oneof"` //message InitializeRequestData = 4;
}

func (*CompletelyNew_InitReq) isCompletelyNew_Paymentselection() {}

func (*CompletelyNew_InitRes) isCompletelyNew_Paymentselection() {}

type NewIndeed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *NewIndeed) Reset() {
	*x = NewIndeed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewIndeed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewIndeed) ProtoMessage() {}

func (x *NewIndeed) ProtoReflect() protoreflect.Message {
	mi := &file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewIndeed.ProtoReflect.Descriptor instead.
func (*NewIndeed) Descriptor() ([]byte, []int) {
	return file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDescGZIP(), []int{1}
}

func (x *NewIndeed) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type InitializeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Regex: ^.{1,36}$
	Reference     string              `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`
	PaymentMethod relay.PaymentMethod `protobuf:"varint,2,opt,name=payment_method,json=paymentMethod,proto3,enum=relay.PaymentMethod" json:"payment_method,omitempty"`
	// Payment method specific information.
	Data *anypb.Any `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *InitializeRequest) Reset() {
	*x = InitializeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeRequest) ProtoMessage() {}

func (x *InitializeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeRequest.ProtoReflect.Descriptor instead.
func (*InitializeRequest) Descriptor() ([]byte, []int) {
	return file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDescGZIP(), []int{2}
}

func (x *InitializeRequest) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *InitializeRequest) GetPaymentMethod() relay.PaymentMethod {
	if x != nil {
		return x.PaymentMethod
	}
	return relay.PaymentMethod(0)
}

func (x *InitializeRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type InitializeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentMethodConfigurationId string `protobuf:"bytes,1,opt,name=payment_method_configuration_id,json=paymentMethodConfigurationId,proto3" json:"payment_method_configuration_id,omitempty"`
	// Payment method specific information.
	Data   *anypb.Any    `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
	Status *relay.Status `protobuf:"bytes,101,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *InitializeResponse) Reset() {
	*x = InitializeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeResponse) ProtoMessage() {}

func (x *InitializeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeResponse.ProtoReflect.Descriptor instead.
func (*InitializeResponse) Descriptor() ([]byte, []int) {
	return file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDescGZIP(), []int{3}
}

func (x *InitializeResponse) GetPaymentMethodConfigurationId() string {
	if x != nil {
		return x.PaymentMethodConfigurationId
	}
	return ""
}

func (x *InitializeResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *InitializeResponse) GetStatus() *relay.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentMethodConfigurationId string              `protobuf:"bytes,1,opt,name=payment_method_configuration_id,json=paymentMethodConfigurationId,proto3" json:"payment_method_configuration_id,omitempty"`
	PaymentMethod                relay.PaymentMethod `protobuf:"varint,2,opt,name=payment_method,json=paymentMethod,proto3,enum=relay.PaymentMethod" json:"payment_method,omitempty"`
	// Payment method specific information.
	Data *anypb.Any `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRequest) GetPaymentMethodConfigurationId() string {
	if x != nil {
		return x.PaymentMethodConfigurationId
	}
	return ""
}

func (x *UpdateRequest) GetPaymentMethod() relay.PaymentMethod {
	if x != nil {
		return x.PaymentMethod
	}
	return relay.PaymentMethod(0)
}

func (x *UpdateRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Payment method specific information.
	Data   *anypb.Any    `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
	Status *relay.Status `protobuf:"bytes,101,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UpdateResponse) GetStatus() *relay.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentMethodConfigurationId string              `protobuf:"bytes,1,opt,name=payment_method_configuration_id,json=paymentMethodConfigurationId,proto3" json:"payment_method_configuration_id,omitempty"`
	PaymentMethod                relay.PaymentMethod `protobuf:"varint,2,opt,name=payment_method,json=paymentMethod,proto3,enum=relay.PaymentMethod" json:"payment_method,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetRequest) GetPaymentMethodConfigurationId() string {
	if x != nil {
		return x.PaymentMethodConfigurationId
	}
	return ""
}

func (x *GetRequest) GetPaymentMethod() relay.PaymentMethod {
	if x != nil {
		return x.PaymentMethod
	}
	return relay.PaymentMethod(0)
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentMethodOnboardingStatus PaymentMethodOnboardingStatus `protobuf:"varint,1,opt,name=payment_method_onboarding_status,json=paymentMethodOnboardingStatus,proto3,enum=relay.onboarding.v1.PaymentMethodOnboardingStatus" json:"payment_method_onboarding_status,omitempty"`
	// Payment method specific information.
	Data   *anypb.Any    `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
	Status *relay.Status `protobuf:"bytes,101,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetResponse) GetPaymentMethodOnboardingStatus() PaymentMethodOnboardingStatus {
	if x != nil {
		return x.PaymentMethodOnboardingStatus
	}
	return PaymentMethodOnboardingStatus_PENDING
}

func (x *GetResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetResponse) GetStatus() *relay.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_relay_onboarding_v1_payment_method_onboarding_service_proto protoreflect.FileDescriptor

var file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2d, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x73, 0x2f, 0x70, 0x70, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7d, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6c, 0x79, 0x4e, 0x65,
	0x77, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x07, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x1a, 0x0a, 0x07, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x07, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x42, 0x12, 0x0a, 0x10, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x2d, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x98,
	0x01, 0x0a, 0x11, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xac, 0x01, 0x0a, 0x12, 0x49, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x45, 0x0a, 0x1f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1c, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xbd, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x1f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x1c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x3b, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52,
	0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x28,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x61, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x65,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x1f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x1c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x3b, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52,
	0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0xdb,
	0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7b,
	0x0a, 0x20, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x5f, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x6e, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x1d, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x6e, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x6f, 0x0a, 0x1d,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x6e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x4e,
	0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b,
	0x44, 0x45, 0x41, 0x43, 0x54, 0x49, 0x56, 0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0e, 0x0a,
	0x0a, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x44, 0x10, 0x04, 0x32, 0xe8, 0x02,
	0x0a, 0x17, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
	0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x5f, 0x0a, 0x0a, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x12, 0x26, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e,
	0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x6f, 0x6e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x6f,
	0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e,
	0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x03, 0x4e,
	0x65, 0x77, 0x12, 0x22, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x6c, 0x79, 0x4e, 0x65, 0x77, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x6f,
	0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77,
	0x49, 0x6e, 0x64, 0x65, 0x65, 0x64, 0x22, 0x00, 0x42, 0x57, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x65, 0x64, 0x76, 0x65, 0x64, 0x37, 0x35, 0x2f,
	0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x6f, 0x6e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0xaa, 0x02, 0x13, 0x52, 0x65,
	0x6c, 0x61, 0x79, 0x2e, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDescOnce sync.Once
	file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDescData = file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDesc
)

func file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDescGZIP() []byte {
	file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDescOnce.Do(func() {
		file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDescData)
	})
	return file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDescData
}

var file_relay_onboarding_v1_payment_method_onboarding_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_relay_onboarding_v1_payment_method_onboarding_service_proto_goTypes = []interface{}{
	(PaymentMethodOnboardingStatus)(0), // 0: relay.onboarding.v1.PaymentMethodOnboardingStatus
	(*CompletelyNew)(nil),              // 1: relay.onboarding.v1.CompletelyNew
	(*NewIndeed)(nil),                  // 2: relay.onboarding.v1.NewIndeed
	(*InitializeRequest)(nil),          // 3: relay.onboarding.v1.InitializeRequest
	(*InitializeResponse)(nil),         // 4: relay.onboarding.v1.InitializeResponse
	(*UpdateRequest)(nil),              // 5: relay.onboarding.v1.UpdateRequest
	(*UpdateResponse)(nil),             // 6: relay.onboarding.v1.UpdateResponse
	(*GetRequest)(nil),                 // 7: relay.onboarding.v1.GetRequest
	(*GetResponse)(nil),                // 8: relay.onboarding.v1.GetResponse
	(relay.PaymentMethod)(0),           // 9: relay.PaymentMethod
	(*anypb.Any)(nil),                  // 10: google.protobuf.Any
	(*relay.Status)(nil),               // 11: relay.Status
}
var file_relay_onboarding_v1_payment_method_onboarding_service_proto_depIdxs = []int32{
	9,  // 0: relay.onboarding.v1.InitializeRequest.payment_method:type_name -> relay.PaymentMethod
	10, // 1: relay.onboarding.v1.InitializeRequest.data:type_name -> google.protobuf.Any
	10, // 2: relay.onboarding.v1.InitializeResponse.data:type_name -> google.protobuf.Any
	11, // 3: relay.onboarding.v1.InitializeResponse.status:type_name -> relay.Status
	9,  // 4: relay.onboarding.v1.UpdateRequest.payment_method:type_name -> relay.PaymentMethod
	10, // 5: relay.onboarding.v1.UpdateRequest.data:type_name -> google.protobuf.Any
	10, // 6: relay.onboarding.v1.UpdateResponse.data:type_name -> google.protobuf.Any
	11, // 7: relay.onboarding.v1.UpdateResponse.status:type_name -> relay.Status
	9,  // 8: relay.onboarding.v1.GetRequest.payment_method:type_name -> relay.PaymentMethod
	0,  // 9: relay.onboarding.v1.GetResponse.payment_method_onboarding_status:type_name -> relay.onboarding.v1.PaymentMethodOnboardingStatus
	10, // 10: relay.onboarding.v1.GetResponse.data:type_name -> google.protobuf.Any
	11, // 11: relay.onboarding.v1.GetResponse.status:type_name -> relay.Status
	3,  // 12: relay.onboarding.v1.PaymentMethodOnboarding.Initialize:input_type -> relay.onboarding.v1.InitializeRequest
	5,  // 13: relay.onboarding.v1.PaymentMethodOnboarding.Update:input_type -> relay.onboarding.v1.UpdateRequest
	7,  // 14: relay.onboarding.v1.PaymentMethodOnboarding.Get:input_type -> relay.onboarding.v1.GetRequest
	1,  // 15: relay.onboarding.v1.PaymentMethodOnboarding.New:input_type -> relay.onboarding.v1.CompletelyNew
	4,  // 16: relay.onboarding.v1.PaymentMethodOnboarding.Initialize:output_type -> relay.onboarding.v1.InitializeResponse
	6,  // 17: relay.onboarding.v1.PaymentMethodOnboarding.Update:output_type -> relay.onboarding.v1.UpdateResponse
	8,  // 18: relay.onboarding.v1.PaymentMethodOnboarding.Get:output_type -> relay.onboarding.v1.GetResponse
	2,  // 19: relay.onboarding.v1.PaymentMethodOnboarding.New:output_type -> relay.onboarding.v1.NewIndeed
	16, // [16:20] is the sub-list for method output_type
	12, // [12:16] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_relay_onboarding_v1_payment_method_onboarding_service_proto_init() }
func file_relay_onboarding_v1_payment_method_onboarding_service_proto_init() {
	if File_relay_onboarding_v1_payment_method_onboarding_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletelyNew); i {
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
		file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewIndeed); i {
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
		file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeRequest); i {
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
		file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeResponse); i {
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
		file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
	file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CompletelyNew_InitReq)(nil),
		(*CompletelyNew_InitRes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relay_onboarding_v1_payment_method_onboarding_service_proto_goTypes,
		DependencyIndexes: file_relay_onboarding_v1_payment_method_onboarding_service_proto_depIdxs,
		EnumInfos:         file_relay_onboarding_v1_payment_method_onboarding_service_proto_enumTypes,
		MessageInfos:      file_relay_onboarding_v1_payment_method_onboarding_service_proto_msgTypes,
	}.Build()
	File_relay_onboarding_v1_payment_method_onboarding_service_proto = out.File
	file_relay_onboarding_v1_payment_method_onboarding_service_proto_rawDesc = nil
	file_relay_onboarding_v1_payment_method_onboarding_service_proto_goTypes = nil
	file_relay_onboarding_v1_payment_method_onboarding_service_proto_depIdxs = nil
}
