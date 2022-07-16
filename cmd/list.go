package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users in gum config.",
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, user := range users {
			fmt.Printf("name: %s, email: %s\n", user.Name, user.Email)
		}
		return nil
	},
}
