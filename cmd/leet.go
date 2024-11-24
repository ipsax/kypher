package cmd

import (
	"github.com/ipsax/kypher/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(leetCmd)
}

var leetCmd = &cobra.Command{
	Use:   "leet",
	Short: "use the 1337 cipher",
	Long:  `This is kypher's leet cipher`,
	Run: func(cmd *cobra.Command, args []string) {
		leet := internal.Leet{}
		leet.Encode("todo, take cli args for encode or decode, maybe add autodetect for autodecode eventually?")
	},
}
