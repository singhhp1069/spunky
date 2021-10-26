package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateRewards{}

func NewMsgCreateRewards(creator string, milestone string, description string, reward string, spunker string) *MsgCreateRewards {
	return &MsgCreateRewards{
		Creator:     creator,
		Milestone:   milestone,
		Description: description,
		Reward:      reward,
		Spunker:     spunker,
	}
}

func (msg *MsgCreateRewards) Route() string {
	return RouterKey
}

func (msg *MsgCreateRewards) Type() string {
	return "CreateRewards"
}

func (msg *MsgCreateRewards) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRewards) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteRewards{}

func NewMsgDeleteRewards(creator string, id uint64) *MsgDeleteRewards {
	return &MsgDeleteRewards{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteRewards) Route() string {
	return RouterKey
}

func (msg *MsgDeleteRewards) Type() string {
	return "DeleteRewards"
}

func (msg *MsgDeleteRewards) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteRewards) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
