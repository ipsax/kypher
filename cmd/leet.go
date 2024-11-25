package cmd

import (
	"fmt"

	"github.com/ipsax/kypher/internal"
	"github.com/spf13/cobra"
)

var encodeArg string

func init() {
	leetCmd.PersistentFlags().StringVarP(&encodeArg, "encode", "e", "", "Leet message value to be encoded")
	rootCmd.AddCommand(leetCmd)
}

var leetCmd = &cobra.Command{
	Use:   "leet [OPTIONS]",
	Short: "use the leet cipher",
	Long:  `This is kypher's leet cipher`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		leet := internal.Leet{}
		if encodeArg != "" {
			fmt.Println(leet.Encode(encodeArg))
		}
	},
}
