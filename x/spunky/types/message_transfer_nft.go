package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransferNFT{}

func NewMsgTransferNFT(creator string, sender string, recipient string, nft string) *MsgTransferNFT {
	return &MsgTransferNFT{
		Creator:   creator,
		Sender:    sender,
		Recipient: recipient,
		Nft:       nft,
	}
}

func (msg *MsgTransferNFT) Route() string {
	return RouterKey
}

func (msg *MsgTransferNFT) Type() string {
	return "TransferNFT"
}

func (msg *MsgTransferNFT) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransferNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
