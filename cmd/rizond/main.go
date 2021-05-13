package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	app "github.com/rizon-world/rizon/app"
	"github.com/rizon-world/rizon/cmd/rizond/cmd"
	"github.com/rizon-world/rizon/types"
)

func main() {
	// Set address prefix and bip44 cointype
	types.SetConfig()

	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
