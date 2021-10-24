package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/singhp1069/spunky/x/spunky/types"
)

func CmdCreateNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-nft [name] [symbol] [description] [uri] [uri-hash]",
		Short: "Create a new NFT",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argSymbol := args[1]
			argDescription := args[2]
			argUri := args[3]
			argUriHash := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateNFT(clientCtx.GetFromAddress().String(), argName, argSymbol, argDescription, argUri, argUriHash)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-nft [id] [name] [symbol] [description] [uri] [uri-hash]",
		Short: "Update a NFT",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argName := args[1]

			argSymbol := args[2]

			argDescription := args[3]

			argUri := args[4]

			argUriHash := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateNFT(clientCtx.GetFromAddress().String(), id, argName, argSymbol, argDescription, argUri, argUriHash)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
