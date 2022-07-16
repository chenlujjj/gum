package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete user in gum config.",
	Example: "gum delete john",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Create(cfgFilePath)
		if err != nil {
			return err
		}
		defer f.Close()

		delete(users, args[0])
		b, err := json.MarshalIndent(users, "", "    ")
		if err != nil {
			return err
		}
		_, err = f.Write(b)
		return err
	},
}
