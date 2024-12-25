package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = cobra.Command{
	Use: "treeman",
	Run: runRootCmd,
}

func runRootCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Hello, Golang!")
}
