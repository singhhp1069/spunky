package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/singhp1069/spunky/x/spunky/types"
)

var _ = strconv.Itoa(0)

func CmdTransferNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-nft [sender] [recipient] [nft]",
		Short: "Broadcast message transferNFT",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSender := args[0]
			argRecipient := args[1]
			argNft := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferNFT(
				clientCtx.GetFromAddress().String(),
				argSender,
				argRecipient,
				argNft,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
