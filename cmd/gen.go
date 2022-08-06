package cmd

import (
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generates text.",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return genText()
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}

func genText() (err error) {
	return nil
}
