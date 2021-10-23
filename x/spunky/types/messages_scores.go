package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateScores{}

func NewMsgCreateScores(creator string, highscore string) *MsgCreateScores {
	return &MsgCreateScores{
		Creator:   creator,
		Highscore: highscore,
	}
}

func (msg *MsgCreateScores) Route() string {
	return RouterKey
}

func (msg *MsgCreateScores) Type() string {
	return "CreateScores"
}

func (msg *MsgCreateScores) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateScores) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateScores) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateScores{}

func NewMsgUpdateScores(creator string, id uint64, highscore string) *MsgUpdateScores {
	return &MsgUpdateScores{
		Id:        id,
		Creator:   creator,
		Highscore: highscore,
	}
}

func (msg *MsgUpdateScores) Route() string {
	return RouterKey
}

func (msg *MsgUpdateScores) Type() string {
	return "UpdateScores"
}

func (msg *MsgUpdateScores) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateScores) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateScores) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteScores{}

func NewMsgDeleteScores(creator string, id uint64) *MsgDeleteScores {
	return &MsgDeleteScores{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteScores) Route() string {
	return RouterKey
}

func (msg *MsgDeleteScores) Type() string {
	return "DeleteScores"
}

func (msg *MsgDeleteScores) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteScores) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteScores) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
