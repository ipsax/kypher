package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kypher",
	Long:  `All software has versions this is kypher's version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Amateur Kypher Cipher Punk Ciphers ~ v0.0.1")
	},
}
