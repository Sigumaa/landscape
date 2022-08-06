package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generates text.",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		arg := ""
		if len(args) != 0 {
			arg = args[0]
		}
		return genText(arg)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}

func genText(arg string) (err error) {
	fmt.Print(arg)
	return nil
}
