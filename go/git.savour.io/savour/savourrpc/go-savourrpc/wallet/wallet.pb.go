// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: savourrpc/wallet.proto

package wallet

import (
	common "git.savour.io/savour/savourrpc/go-savourrpc/common"
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

type SupportCoins struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId   string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ChainName string `protobuf:"bytes,2,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty"`
	CoinId    string `protobuf:"bytes,3,opt,name=coin_id,json=coinId,proto3" json:"coin_id,omitempty"`
	CoinName  string `protobuf:"bytes,4,opt,name=coin_name,json=coinName,proto3" json:"coin_name,omitempty"`
	Network   string `protobuf:"bytes,5,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *SupportCoins) Reset() {
	*x = SupportCoins{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupportCoins) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportCoins) ProtoMessage() {}

func (x *SupportCoins) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportCoins.ProtoReflect.Descriptor instead.
func (*SupportCoins) Descriptor() ([]byte, []int) {
	return file_savourrpc_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *SupportCoins) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *SupportCoins) GetChainName() string {
	if x != nil {
		return x.ChainName
	}
	return ""
}

func (x *SupportCoins) GetCoinId() string {
	if x != nil {
		return x.CoinId
	}
	return ""
}

func (x *SupportCoins) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *SupportCoins) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

type SupportCoinsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
}

func (x *SupportCoinsRequest) Reset() {
	*x = SupportCoinsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupportCoinsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportCoinsRequest) ProtoMessage() {}

func (x *SupportCoinsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportCoinsRequest.ProtoReflect.Descriptor instead.
func (*SupportCoinsRequest) Descriptor() ([]byte, []int) {
	return file_savourrpc_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *SupportCoinsRequest) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

type SupportCoinsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error        *common.Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	SupportCoins []*SupportCoins `protobuf:"bytes,2,rep,name=support_coins,json=supportCoins,proto3" json:"support_coins,omitempty"`
}

func (x *SupportCoinsResponse) Reset() {
	*x = SupportCoinsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupportCoinsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportCoinsResponse) ProtoMessage() {}

func (x *SupportCoinsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportCoinsResponse.ProtoReflect.Descriptor instead.
func (*SupportCoinsResponse) Descriptor() ([]byte, []int) {
	return file_savourrpc_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *SupportCoinsResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SupportCoinsResponse) GetSupportCoins() []*SupportCoins {
	if x != nil {
		return x.SupportCoins
	}
	return nil
}

type NonceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	Address       string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *NonceRequest) Reset() {
	*x = NonceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NonceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonceRequest) ProtoMessage() {}

func (x *NonceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonceRequest.ProtoReflect.Descriptor instead.
func (*NonceRequest) Descriptor() ([]byte, []int) {
	return file_savourrpc_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *NonceRequest) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *NonceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type NonceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *common.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Nonce string        `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *NonceResponse) Reset() {
	*x = NonceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_wallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NonceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonceResponse) ProtoMessage() {}

func (x *NonceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_wallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonceResponse.ProtoReflect.Descriptor instead.
func (*NonceResponse) Descriptor() ([]byte, []int) {
	return file_savourrpc_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *NonceResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *NonceResponse) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

type GasPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
}

func (x *GasPriceRequest) Reset() {
	*x = GasPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_wallet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GasPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GasPriceRequest) ProtoMessage() {}

func (x *GasPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_wallet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GasPriceRequest.ProtoReflect.Descriptor instead.
func (*GasPriceRequest) Descriptor() ([]byte, []int) {
	return file_savourrpc_wallet_proto_rawDescGZIP(), []int{5}
}

func (x *GasPriceRequest) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

type GasPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *common.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Gas   string        `protobuf:"bytes,2,opt,name=gas,proto3" json:"gas,omitempty"`
}

func (x *GasPriceResponse) Reset() {
	*x = GasPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_wallet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GasPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GasPriceResponse) ProtoMessage() {}

func (x *GasPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_wallet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GasPriceResponse.ProtoReflect.Descriptor instead.
func (*GasPriceResponse) Descriptor() ([]byte, []int) {
	return file_savourrpc_wallet_proto_rawDescGZIP(), []int{6}
}

func (x *GasPriceResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GasPriceResponse) GetGas() string {
	if x != nil {
		return x.Gas
	}
	return ""
}

type SendTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	CoinId        string `protobuf:"bytes,2,opt,name=coin_id,json=coinId,proto3" json:"coin_id,omitempty"`
	ChainId       string `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Network       string `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	RawTx         string `protobuf:"bytes,5,opt,name=raw_tx,json=rawTx,proto3" json:"raw_tx,omitempty"`
}

func (x *SendTxRequest) Reset() {
	*x = SendTxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_wallet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTxRequest) ProtoMessage() {}

func (x *SendTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_wallet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTxRequest.ProtoReflect.Descriptor instead.
func (*SendTxRequest) Descriptor() ([]byte, []int) {
	return file_savourrpc_wallet_proto_rawDescGZIP(), []int{7}
}

func (x *SendTxRequest) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *SendTxRequest) GetCoinId() string {
	if x != nil {
		return x.CoinId
	}
	return ""
}

func (x *SendTxRequest) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *SendTxRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *SendTxRequest) GetRawTx() string {
	if x != nil {
		return x.RawTx
	}
	return ""
}

type SendTxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *common.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	TxHash string        `protobuf:"bytes,2,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
}

func (x *SendTxResponse) Reset() {
	*x = SendTxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savourrpc_wallet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTxResponse) ProtoMessage() {}

func (x *SendTxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_savourrpc_wallet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTxResponse.ProtoReflect.Descriptor instead.
func (*SendTxResponse) Descriptor() ([]byte, []int) {
	return file_savourrpc_wallet_proto_rawDescGZIP(), []int{8}
}

func (x *SendTxResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SendTxResponse) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

var File_savourrpc_wallet_proto protoreflect.FileDescriptor

var file_savourrpc_wallet_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72,
	0x72, 0x70, 0x63, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x1a, 0x16, 0x73, 0x61, 0x76, 0x6f,
	0x75, 0x72, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x0c, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x69, 0x6e, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x6f, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x3c, 0x0a,
	0x13, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x83, 0x01, 0x0a, 0x14,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x43, 0x0a, 0x0d,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x69, 0x6e, 0x73, 0x52, 0x0c, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x69, 0x6e,
	0x73, 0x22, 0x4f, 0x0a, 0x0c, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x4d, 0x0a, 0x0d, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x22, 0x38, 0x0a, 0x0f, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4c, 0x0a, 0x10, 0x47,
	0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x61, 0x73, 0x22, 0x9b, 0x01, 0x0a, 0x0d, 0x53, 0x65,
	0x6e, 0x64, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x15, 0x0a, 0x06, 0x72, 0x61, 0x77, 0x5f, 0x74, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x61, 0x77, 0x54, 0x78, 0x22, 0x51, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x54,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75,
	0x72, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x32, 0xe9, 0x02, 0x0a, 0x0d, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x0f,
	0x67, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x12,
	0x25, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72,
	0x70, 0x63, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4d, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x73,
	0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73,
	0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x56, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x21,
	0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x2e, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x54,
	0x78, 0x12, 0x1f, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4a, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61,
	0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5a, 0x32,
	0x67, 0x69, 0x74, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x61,
	0x76, 0x6f, 0x75, 0x72, 0x2f, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_savourrpc_wallet_proto_rawDescOnce sync.Once
	file_savourrpc_wallet_proto_rawDescData = file_savourrpc_wallet_proto_rawDesc
)

func file_savourrpc_wallet_proto_rawDescGZIP() []byte {
	file_savourrpc_wallet_proto_rawDescOnce.Do(func() {
		file_savourrpc_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_savourrpc_wallet_proto_rawDescData)
	})
	return file_savourrpc_wallet_proto_rawDescData
}

var file_savourrpc_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_savourrpc_wallet_proto_goTypes = []interface{}{
	(*SupportCoins)(nil),         // 0: savourrpc.wallet.SupportCoins
	(*SupportCoinsRequest)(nil),  // 1: savourrpc.wallet.SupportCoinsRequest
	(*SupportCoinsResponse)(nil), // 2: savourrpc.wallet.SupportCoinsResponse
	(*NonceRequest)(nil),         // 3: savourrpc.wallet.NonceRequest
	(*NonceResponse)(nil),        // 4: savourrpc.wallet.NonceResponse
	(*GasPriceRequest)(nil),      // 5: savourrpc.wallet.GasPriceRequest
	(*GasPriceResponse)(nil),     // 6: savourrpc.wallet.GasPriceResponse
	(*SendTxRequest)(nil),        // 7: savourrpc.wallet.SendTxRequest
	(*SendTxResponse)(nil),       // 8: savourrpc.wallet.SendTxResponse
	(*common.Error)(nil),         // 9: savourrpc.Error
}
var file_savourrpc_wallet_proto_depIdxs = []int32{
	9, // 0: savourrpc.wallet.SupportCoinsResponse.error:type_name -> savourrpc.Error
	0, // 1: savourrpc.wallet.SupportCoinsResponse.support_coins:type_name -> savourrpc.wallet.SupportCoins
	9, // 2: savourrpc.wallet.NonceResponse.error:type_name -> savourrpc.Error
	9, // 3: savourrpc.wallet.GasPriceResponse.error:type_name -> savourrpc.Error
	9, // 4: savourrpc.wallet.SendTxResponse.error:type_name -> savourrpc.Error
	1, // 5: savourrpc.wallet.WalletService.getSupportCoins:input_type -> savourrpc.wallet.SupportCoinsRequest
	3, // 6: savourrpc.wallet.WalletService.getNonce:input_type -> savourrpc.wallet.NonceRequest
	5, // 7: savourrpc.wallet.WalletService.getGasPrice:input_type -> savourrpc.wallet.GasPriceRequest
	7, // 8: savourrpc.wallet.WalletService.SendTx:input_type -> savourrpc.wallet.SendTxRequest
	2, // 9: savourrpc.wallet.WalletService.getSupportCoins:output_type -> savourrpc.wallet.SupportCoinsResponse
	4, // 10: savourrpc.wallet.WalletService.getNonce:output_type -> savourrpc.wallet.NonceResponse
	6, // 11: savourrpc.wallet.WalletService.getGasPrice:output_type -> savourrpc.wallet.GasPriceResponse
	8, // 12: savourrpc.wallet.WalletService.SendTx:output_type -> savourrpc.wallet.SendTxResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_savourrpc_wallet_proto_init() }
func file_savourrpc_wallet_proto_init() {
	if File_savourrpc_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_savourrpc_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupportCoins); i {
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
		file_savourrpc_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupportCoinsRequest); i {
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
		file_savourrpc_wallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupportCoinsResponse); i {
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
		file_savourrpc_wallet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NonceRequest); i {
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
		file_savourrpc_wallet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NonceResponse); i {
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
		file_savourrpc_wallet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GasPriceRequest); i {
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
		file_savourrpc_wallet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GasPriceResponse); i {
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
		file_savourrpc_wallet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTxRequest); i {
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
		file_savourrpc_wallet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTxResponse); i {
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
			RawDescriptor: file_savourrpc_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_savourrpc_wallet_proto_goTypes,
		DependencyIndexes: file_savourrpc_wallet_proto_depIdxs,
		MessageInfos:      file_savourrpc_wallet_proto_msgTypes,
	}.Build()
	File_savourrpc_wallet_proto = out.File
	file_savourrpc_wallet_proto_rawDesc = nil
	file_savourrpc_wallet_proto_goTypes = nil
	file_savourrpc_wallet_proto_depIdxs = nil
}
