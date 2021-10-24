package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateScores{}, "spunky/CreateScores", nil)
	cdc.RegisterConcrete(&MsgUpdateScores{}, "spunky/UpdateScores", nil)
	cdc.RegisterConcrete(&MsgCreateNFT{}, "spunky/CreateNFT", nil)
	cdc.RegisterConcrete(&MsgUpdateNFT{}, "spunky/UpdateNFT", nil)
	cdc.RegisterConcrete(&MsgTransferNFT{}, "spunky/TransferNFT", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateScores{},
		&MsgUpdateScores{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateNFT{},
		&MsgUpdateNFT{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferNFT{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
