// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: logistica.proto

package logistica

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type OrdenPyme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Producto    string `protobuf:"bytes,2,opt,name=Producto,proto3" json:"Producto,omitempty"`
	Valor       int32  `protobuf:"varint,3,opt,name=Valor,proto3" json:"Valor,omitempty"`
	Origen      string `protobuf:"bytes,4,opt,name=Origen,proto3" json:"Origen,omitempty"`
	Destino     string `protobuf:"bytes,5,opt,name=Destino,proto3" json:"Destino,omitempty"`
	Prioritario int32  `protobuf:"varint,6,opt,name=Prioritario,proto3" json:"Prioritario,omitempty"`
}

func (x *OrdenPyme) Reset() {
	*x = OrdenPyme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistica_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdenPyme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdenPyme) ProtoMessage() {}

func (x *OrdenPyme) ProtoReflect() protoreflect.Message {
	mi := &file_logistica_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdenPyme.ProtoReflect.Descriptor instead.
func (*OrdenPyme) Descriptor() ([]byte, []int) {
	return file_logistica_proto_rawDescGZIP(), []int{0}
}

func (x *OrdenPyme) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrdenPyme) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *OrdenPyme) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *OrdenPyme) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *OrdenPyme) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *OrdenPyme) GetPrioritario() int32 {
	if x != nil {
		return x.Prioritario
	}
	return 0
}

type SeguimientoPaqueteSolicitado struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDPaquete string `protobuf:"bytes,1,opt,name=IDPaquete,proto3" json:"IDPaquete,omitempty"`
	Estado    string `protobuf:"bytes,2,opt,name=Estado,proto3" json:"Estado,omitempty"`
}

func (x *SeguimientoPaqueteSolicitado) Reset() {
	*x = SeguimientoPaqueteSolicitado{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistica_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeguimientoPaqueteSolicitado) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeguimientoPaqueteSolicitado) ProtoMessage() {}

func (x *SeguimientoPaqueteSolicitado) ProtoReflect() protoreflect.Message {
	mi := &file_logistica_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeguimientoPaqueteSolicitado.ProtoReflect.Descriptor instead.
func (*SeguimientoPaqueteSolicitado) Descriptor() ([]byte, []int) {
	return file_logistica_proto_rawDescGZIP(), []int{1}
}

func (x *SeguimientoPaqueteSolicitado) GetIDPaquete() string {
	if x != nil {
		return x.IDPaquete
	}
	return ""
}

func (x *SeguimientoPaqueteSolicitado) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type SeguimientoPyme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *SeguimientoPyme) Reset() {
	*x = SeguimientoPyme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistica_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeguimientoPyme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeguimientoPyme) ProtoMessage() {}

func (x *SeguimientoPyme) ProtoReflect() protoreflect.Message {
	mi := &file_logistica_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeguimientoPyme.ProtoReflect.Descriptor instead.
func (*SeguimientoPyme) Descriptor() ([]byte, []int) {
	return file_logistica_proto_rawDescGZIP(), []int{2}
}

func (x *SeguimientoPyme) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OrdenRetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Producto string `protobuf:"bytes,2,opt,name=Producto,proto3" json:"Producto,omitempty"`
	Valor    int32  `protobuf:"varint,3,opt,name=Valor,proto3" json:"Valor,omitempty"`
	Origen   string `protobuf:"bytes,4,opt,name=Origen,proto3" json:"Origen,omitempty"`
	Destino  string `protobuf:"bytes,5,opt,name=Destino,proto3" json:"Destino,omitempty"`
}

func (x *OrdenRetail) Reset() {
	*x = OrdenRetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistica_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdenRetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdenRetail) ProtoMessage() {}

func (x *OrdenRetail) ProtoReflect() protoreflect.Message {
	mi := &file_logistica_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdenRetail.ProtoReflect.Descriptor instead.
func (*OrdenRetail) Descriptor() ([]byte, []int) {
	return file_logistica_proto_rawDescGZIP(), []int{3}
}

func (x *OrdenRetail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrdenRetail) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *OrdenRetail) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *OrdenRetail) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *OrdenRetail) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

type SeguimientoRetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *SeguimientoRetail) Reset() {
	*x = SeguimientoRetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistica_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeguimientoRetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeguimientoRetail) ProtoMessage() {}

func (x *SeguimientoRetail) ProtoReflect() protoreflect.Message {
	mi := &file_logistica_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeguimientoRetail.ProtoReflect.Descriptor instead.
func (*SeguimientoRetail) Descriptor() ([]byte, []int) {
	return file_logistica_proto_rawDescGZIP(), []int{4}
}

func (x *SeguimientoRetail) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Camion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Tipo          string `protobuf:"bytes,2,opt,name=Tipo,proto3" json:"Tipo,omitempty"`
	EntregaRetail bool   `protobuf:"varint,3,opt,name=EntregaRetail,proto3" json:"EntregaRetail,omitempty"`
}

func (x *Camion) Reset() {
	*x = Camion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistica_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Camion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Camion) ProtoMessage() {}

func (x *Camion) ProtoReflect() protoreflect.Message {
	mi := &file_logistica_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Camion.ProtoReflect.Descriptor instead.
func (*Camion) Descriptor() ([]byte, []int) {
	return file_logistica_proto_rawDescGZIP(), []int{5}
}

func (x *Camion) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Camion) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *Camion) GetEntregaRetail() bool {
	if x != nil {
		return x.EntregaRetail
	}
	return false
}

type PaqueteCamion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Tipo    string `protobuf:"bytes,2,opt,name=Tipo,proto3" json:"Tipo,omitempty"`
	Valor   int32  `protobuf:"varint,3,opt,name=Valor,proto3" json:"Valor,omitempty"`
	Origen  string `protobuf:"bytes,4,opt,name=Origen,proto3" json:"Origen,omitempty"`
	Destino string `protobuf:"bytes,5,opt,name=Destino,proto3" json:"Destino,omitempty"`
}

func (x *PaqueteCamion) Reset() {
	*x = PaqueteCamion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistica_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaqueteCamion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaqueteCamion) ProtoMessage() {}

func (x *PaqueteCamion) ProtoReflect() protoreflect.Message {
	mi := &file_logistica_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaqueteCamion.ProtoReflect.Descriptor instead.
func (*PaqueteCamion) Descriptor() ([]byte, []int) {
	return file_logistica_proto_rawDescGZIP(), []int{6}
}

func (x *PaqueteCamion) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PaqueteCamion) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *PaqueteCamion) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *PaqueteCamion) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *PaqueteCamion) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistica_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logistica_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_logistica_proto_rawDescGZIP(), []int{7}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistica_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_logistica_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_logistica_proto_rawDescGZIP(), []int{8}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_logistica_proto protoreflect.FileDescriptor

var file_logistica_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x22, 0xa1, 0x01, 0x0a,
	0x09, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x50, 0x79, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x4f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x72,
	0x69, 0x67, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x12, 0x20,
	0x0a, 0x0b, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x61, 0x72, 0x69, 0x6f, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x61, 0x72, 0x69, 0x6f,
	0x22, 0x54, 0x0a, 0x1c, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x50,
	0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x64, 0x6f,
	0x12, 0x1c, 0x0a, 0x09, 0x49, 0x44, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x44, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x21, 0x0a, 0x0f, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d,
	0x69, 0x65, 0x6e, 0x74, 0x6f, 0x50, 0x79, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x0b, 0x4f, 0x72,
	0x64, 0x65, 0x6e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x72, 0x69, 0x67, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x72, 0x69,
	0x67, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x22, 0x23, 0x0a,
	0x11, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x52, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x49, 0x64, 0x22, 0x52, 0x0a, 0x06, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x69, 0x70, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x69, 0x70, 0x6f,
	0x12, 0x24, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x72, 0x65, 0x67, 0x61, 0x52, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x45, 0x6e, 0x74, 0x72, 0x65, 0x67, 0x61,
	0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x7b, 0x0a, 0x0d, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x70, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x69, 0x70, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x6f,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x88, 0x03, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x17, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x72, 0x4f, 0x72, 0x64,
	0x65, 0x6e, 0x50, 0x79, 0x6d, 0x65, 0x12, 0x14, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x61, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x50, 0x79, 0x6d, 0x65, 0x1a, 0x1a, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69,
	0x65, 0x6e, 0x74, 0x6f, 0x50, 0x79, 0x6d, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x12, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x16, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x6e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x1c, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x61, 0x2e, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f,
	0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x10, 0x53, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x61, 0x72, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x1a,
	0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x50, 0x61, 0x71, 0x75,
	0x65, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x14, 0x53,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x72, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65,
	0x6e, 0x74, 0x6f, 0x12, 0x1a, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e,
	0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x50, 0x79, 0x6d, 0x65, 0x1a,
	0x27, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x53, 0x65, 0x67, 0x75,
	0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x53, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_logistica_proto_rawDescOnce sync.Once
	file_logistica_proto_rawDescData = file_logistica_proto_rawDesc
)

func file_logistica_proto_rawDescGZIP() []byte {
	file_logistica_proto_rawDescOnce.Do(func() {
		file_logistica_proto_rawDescData = protoimpl.X.CompressGZIP(file_logistica_proto_rawDescData)
	})
	return file_logistica_proto_rawDescData
}

var file_logistica_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_logistica_proto_goTypes = []interface{}{
	(*OrdenPyme)(nil),                    // 0: logistica.OrdenPyme
	(*SeguimientoPaqueteSolicitado)(nil), // 1: logistica.SeguimientoPaqueteSolicitado
	(*SeguimientoPyme)(nil),              // 2: logistica.SeguimientoPyme
	(*OrdenRetail)(nil),                  // 3: logistica.OrdenRetail
	(*SeguimientoRetail)(nil),            // 4: logistica.SeguimientoRetail
	(*Camion)(nil),                       // 5: logistica.Camion
	(*PaqueteCamion)(nil),                // 6: logistica.PaqueteCamion
	(*HelloRequest)(nil),                 // 7: logistica.HelloRequest
	(*HelloReply)(nil),                   // 8: logistica.HelloReply
}
var file_logistica_proto_depIdxs = []int32{
	7, // 0: logistica.LogisticaService.SayHello:input_type -> logistica.HelloRequest
	0, // 1: logistica.LogisticaService.GenerarOrdenPyme:input_type -> logistica.OrdenPyme
	3, // 2: logistica.LogisticaService.GenerarOrdenRetail:input_type -> logistica.OrdenRetail
	5, // 3: logistica.LogisticaService.SolicitarPaquete:input_type -> logistica.Camion
	2, // 4: logistica.LogisticaService.SolicitarSeguimiento:input_type -> logistica.SeguimientoPyme
	8, // 5: logistica.LogisticaService.SayHello:output_type -> logistica.HelloReply
	2, // 6: logistica.LogisticaService.GenerarOrdenPyme:output_type -> logistica.SeguimientoPyme
	4, // 7: logistica.LogisticaService.GenerarOrdenRetail:output_type -> logistica.SeguimientoRetail
	6, // 8: logistica.LogisticaService.SolicitarPaquete:output_type -> logistica.PaqueteCamion
	1, // 9: logistica.LogisticaService.SolicitarSeguimiento:output_type -> logistica.SeguimientoPaqueteSolicitado
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_logistica_proto_init() }
func file_logistica_proto_init() {
	if File_logistica_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logistica_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrdenPyme); i {
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
		file_logistica_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeguimientoPaqueteSolicitado); i {
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
		file_logistica_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeguimientoPyme); i {
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
		file_logistica_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrdenRetail); i {
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
		file_logistica_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeguimientoRetail); i {
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
		file_logistica_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Camion); i {
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
		file_logistica_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaqueteCamion); i {
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
		file_logistica_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_logistica_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
			RawDescriptor: file_logistica_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logistica_proto_goTypes,
		DependencyIndexes: file_logistica_proto_depIdxs,
		MessageInfos:      file_logistica_proto_msgTypes,
	}.Build()
	File_logistica_proto = out.File
	file_logistica_proto_rawDesc = nil
	file_logistica_proto_goTypes = nil
	file_logistica_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogisticaServiceClient is the client API for LogisticaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogisticaServiceClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	GenerarOrdenPyme(ctx context.Context, in *OrdenPyme, opts ...grpc.CallOption) (*SeguimientoPyme, error)
	GenerarOrdenRetail(ctx context.Context, in *OrdenRetail, opts ...grpc.CallOption) (*SeguimientoRetail, error)
	SolicitarPaquete(ctx context.Context, in *Camion, opts ...grpc.CallOption) (*PaqueteCamion, error)
	SolicitarSeguimiento(ctx context.Context, in *SeguimientoPyme, opts ...grpc.CallOption) (*SeguimientoPaqueteSolicitado, error)
}

type logisticaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogisticaServiceClient(cc grpc.ClientConnInterface) LogisticaServiceClient {
	return &logisticaServiceClient{cc}
}

func (c *logisticaServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/logistica.LogisticaService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticaServiceClient) GenerarOrdenPyme(ctx context.Context, in *OrdenPyme, opts ...grpc.CallOption) (*SeguimientoPyme, error) {
	out := new(SeguimientoPyme)
	err := c.cc.Invoke(ctx, "/logistica.LogisticaService/GenerarOrdenPyme", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticaServiceClient) GenerarOrdenRetail(ctx context.Context, in *OrdenRetail, opts ...grpc.CallOption) (*SeguimientoRetail, error) {
	out := new(SeguimientoRetail)
	err := c.cc.Invoke(ctx, "/logistica.LogisticaService/GenerarOrdenRetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticaServiceClient) SolicitarPaquete(ctx context.Context, in *Camion, opts ...grpc.CallOption) (*PaqueteCamion, error) {
	out := new(PaqueteCamion)
	err := c.cc.Invoke(ctx, "/logistica.LogisticaService/SolicitarPaquete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logisticaServiceClient) SolicitarSeguimiento(ctx context.Context, in *SeguimientoPyme, opts ...grpc.CallOption) (*SeguimientoPaqueteSolicitado, error) {
	out := new(SeguimientoPaqueteSolicitado)
	err := c.cc.Invoke(ctx, "/logistica.LogisticaService/SolicitarSeguimiento", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogisticaServiceServer is the server API for LogisticaService service.
type LogisticaServiceServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	GenerarOrdenPyme(context.Context, *OrdenPyme) (*SeguimientoPyme, error)
	GenerarOrdenRetail(context.Context, *OrdenRetail) (*SeguimientoRetail, error)
	SolicitarPaquete(context.Context, *Camion) (*PaqueteCamion, error)
	SolicitarSeguimiento(context.Context, *SeguimientoPyme) (*SeguimientoPaqueteSolicitado, error)
}

// UnimplementedLogisticaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLogisticaServiceServer struct {
}

func (*UnimplementedLogisticaServiceServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedLogisticaServiceServer) GenerarOrdenPyme(context.Context, *OrdenPyme) (*SeguimientoPyme, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerarOrdenPyme not implemented")
}
func (*UnimplementedLogisticaServiceServer) GenerarOrdenRetail(context.Context, *OrdenRetail) (*SeguimientoRetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerarOrdenRetail not implemented")
}
func (*UnimplementedLogisticaServiceServer) SolicitarPaquete(context.Context, *Camion) (*PaqueteCamion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarPaquete not implemented")
}
func (*UnimplementedLogisticaServiceServer) SolicitarSeguimiento(context.Context, *SeguimientoPyme) (*SeguimientoPaqueteSolicitado, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarSeguimiento not implemented")
}

func RegisterLogisticaServiceServer(s *grpc.Server, srv LogisticaServiceServer) {
	s.RegisterService(&_LogisticaService_serviceDesc, srv)
}

func _LogisticaService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticaServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistica.LogisticaService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticaServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogisticaService_GenerarOrdenPyme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdenPyme)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticaServiceServer).GenerarOrdenPyme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistica.LogisticaService/GenerarOrdenPyme",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticaServiceServer).GenerarOrdenPyme(ctx, req.(*OrdenPyme))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogisticaService_GenerarOrdenRetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdenRetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticaServiceServer).GenerarOrdenRetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistica.LogisticaService/GenerarOrdenRetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticaServiceServer).GenerarOrdenRetail(ctx, req.(*OrdenRetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogisticaService_SolicitarPaquete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Camion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticaServiceServer).SolicitarPaquete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistica.LogisticaService/SolicitarPaquete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticaServiceServer).SolicitarPaquete(ctx, req.(*Camion))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogisticaService_SolicitarSeguimiento_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeguimientoPyme)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogisticaServiceServer).SolicitarSeguimiento(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistica.LogisticaService/SolicitarSeguimiento",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogisticaServiceServer).SolicitarSeguimiento(ctx, req.(*SeguimientoPyme))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogisticaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logistica.LogisticaService",
	HandlerType: (*LogisticaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _LogisticaService_SayHello_Handler,
		},
		{
			MethodName: "GenerarOrdenPyme",
			Handler:    _LogisticaService_GenerarOrdenPyme_Handler,
		},
		{
			MethodName: "GenerarOrdenRetail",
			Handler:    _LogisticaService_GenerarOrdenRetail_Handler,
		},
		{
			MethodName: "SolicitarPaquete",
			Handler:    _LogisticaService_SolicitarPaquete_Handler,
		},
		{
			MethodName: "SolicitarSeguimiento",
			Handler:    _LogisticaService_SolicitarSeguimiento_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logistica.proto",
}
