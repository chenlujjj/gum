package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add user to gum config.",
	Example: "gum add john john@doe.com",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Create(cfgFilePath)
		if err != nil {
			return err
		}
		defer f.Close()

		name, email := args[0], args[1]
		if _, ok := users[name]; ok {
			fmt.Printf("User %s already in gum config, will be overwritten\n", name)
		}
		users[name] = user{Name: name, Email: email}
		b, err := json.MarshalIndent(users, "", "    ")
		if err != nil {
			return err
		}
		_, err = f.Write(b)
		return err
	},
}
