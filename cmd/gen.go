package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generates text.",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("file name argument required")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return genText(args[0])
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}

func genText(arg string) (err error) {
	fmt.Print(arg)
	err = fileEdit(arg)
	if err != nil {
		return err
	}
	return nil
}
