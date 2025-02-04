package types

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	gogotypes "github.com/gogo/protobuf/types"

	"github.com/irisnet/irismod/modules/mt/exported"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

// RegisterLegacyAminoCodec concrete types on codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgIssueDenom{}, "furymod/mt/MsgIssueDenom", nil)
	cdc.RegisterConcrete(&MsgTransferMT{}, "furymod/mt/MsgTransferMT", nil)
	cdc.RegisterConcrete(&MsgEditMT{}, "furymod/mt/MsgEditMT", nil)
	cdc.RegisterConcrete(&MsgMintMT{}, "furymod/mt/MsgMintMT", nil)
	cdc.RegisterConcrete(&MsgBurnMT{}, "furymod/mt/MsgBurnMT", nil)
	cdc.RegisterConcrete(&MsgTransferDenom{}, "furymod/mt/MsgTransferDenom", nil)

	cdc.RegisterInterface((*exported.MT)(nil), nil)
	cdc.RegisterConcrete(&MT{}, "furymod/mt/MT", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgIssueDenom{},
		&MsgTransferMT{},
		&MsgEditMT{},
		&MsgMintMT{},
		&MsgBurnMT{},
		&MsgTransferDenom{},
	)

	registry.RegisterImplementations(
		(*exported.MT)(nil),
		&MT{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// return supply protobuf code
func MustMarshalSupply(cdc codec.Codec, supply uint64) []byte {
	supplyWrap := gogotypes.UInt64Value{Value: supply}
	return cdc.MustMarshal(&supplyWrap)
}

// return th supply
func MustUnMarshalSupply(cdc codec.Codec, value []byte) uint64 {
	var supplyWrap gogotypes.UInt64Value
	cdc.MustUnmarshal(value, &supplyWrap)
	return supplyWrap.Value
}

// return the mtID protobuf code
func MustMarshalMTID(cdc codec.Codec, mtID string) []byte {
	mtIDWrap := gogotypes.StringValue{Value: mtID}
	return cdc.MustMarshal(&mtIDWrap)
}

// return th mtID
func MustUnMarshalMTID(cdc codec.Codec, value []byte) string {
	var mtIDWrap gogotypes.StringValue
	cdc.MustUnmarshal(value, &mtIDWrap)
	return mtIDWrap.Value
}

func MustMarshalAmount(cdc codec.Codec, amount uint64) []byte {
	amountWrap := gogotypes.UInt64Value{Value: amount}
	return cdc.MustMarshal(&amountWrap)
}

func MustUnMarshalAmount(cdc codec.Codec, value []byte) uint64 {
	var amount gogotypes.UInt64Value
	cdc.MustUnmarshal(value, &amount)
	return amount.Value
}
