package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"
)

var version = "0.0.1"

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use: "Tiger Operating System",
	Version: version,
	Short: "Tiger OS PoC CLI. Enter the commands -- ",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error - '%s' ",err)
		os.Exit(1)
	}
}