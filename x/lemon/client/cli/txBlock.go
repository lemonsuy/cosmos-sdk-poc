package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/lemonsuy/lemon/x/lemon/types"
)

func CmdCreateBlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-block [blockNumber] [blockHash] [blockHeader] [blockData]",
		Short: "Create a new block",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsBlockNumber, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argsBlockHash, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsBlockHeader, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}
			argsBlockData, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateBlock(clientCtx.GetFromAddress().String(), argsBlockNumber, argsBlockHash, argsBlockHeader, argsBlockData)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateBlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-block [id] [blockNumber] [blockHash] [blockHeader] [blockData]",
		Short: "Update a block",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsBlockNumber, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			argsBlockHash, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			argsBlockHeader, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}

			argsBlockData, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateBlock(clientCtx.GetFromAddress().String(), id, argsBlockNumber, argsBlockHash, argsBlockHeader, argsBlockData)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteBlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-block [id]",
		Short: "Delete a block by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteBlock(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
