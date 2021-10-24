package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateNFT{}

func NewMsgCreateNFT(creator string, name string, symbol string, description string, uri string, uriHash string) *MsgCreateNFT {
	return &MsgCreateNFT{
		Creator:     creator,
		Name:        name,
		Symbol:      symbol,
		Description: description,
		Uri:         uri,
		UriHash:     uriHash,
	}
}

func (msg *MsgCreateNFT) Route() string {
	return RouterKey
}

func (msg *MsgCreateNFT) Type() string {
	return "CreateNFT"
}

func (msg *MsgCreateNFT) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateNFT{}

func NewMsgUpdateNFT(creator string, id uint64, name string, symbol string, description string, uri string, uriHash string) *MsgUpdateNFT {
	return &MsgUpdateNFT{
		Id:          id,
		Creator:     creator,
		Name:        name,
		Symbol:      symbol,
		Description: description,
		Uri:         uri,
		UriHash:     uriHash,
	}
}

func (msg *MsgUpdateNFT) Route() string {
	return RouterKey
}

func (msg *MsgUpdateNFT) Type() string {
	return "UpdateNFT"
}

func (msg *MsgUpdateNFT) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
