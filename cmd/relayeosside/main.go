package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/keys"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	eosrelay "github.com/blockchain-develop/relay-eos-side/app"
	"github.com/blockchain-develop/relay-eos-side/types"
)

// rootCmd is the entry point for this binary
var (
	rootCmd = &cobra.Command{
		Use:   "relay",
		Short: "king chain and hub communication relay",
	}
)

func main() {
	// disable sorting
	cobra.EnableCommandSorting = false

	// get the codec
	cdc := eosrelay.MakeCodec()

	// TODO: Setup keybase, viper object, etc. to be passed into
	// the below functions and eliminate global vars, like we do
	// with the cdc.

	// add query/post commands (custom to binary)
	rootCmd.AddCommand(
		client.GetCommands(
			authcmd.GetAccountCmd("acc", cdc, types.GetAccountDecoder(cdc)),
		)...)
	
	rootCmd.AddCommand(
		client.PostCommands(
			eosrelay.IBCRelayCmd(cdc),
		)...)

	// add proxy, version and key info
	rootCmd.AddCommand(
		client.LineBreak,
		keys.Commands(),
		version.VersionCmd,
	)

	// prepare and add flags
	executor := cli.PrepareMainCmd(rootCmd, "BC", os.ExpandEnv("$HOME/.eosrelay"))
	err := executor.Execute()
	if err != nil {
		// Note: Handle with #870
		panic(err)
	}
}
