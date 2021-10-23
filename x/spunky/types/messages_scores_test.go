package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/singhp1069/spunky/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateScores_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateScores
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateScores{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateScores{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateScores_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateScores
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateScores{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateScores{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteScores_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteScores
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteScores{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteScores{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
