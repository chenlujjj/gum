package cmd

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:       "set",
	Short:     "Set local git user.",
	Example:   "gum set john",
	Args:      cobra.ExactArgs(1),
	ValidArgs: userNames,
	RunE: func(cmd *cobra.Command, args []string) error {
		user, ok := users[args[0]]
		if !ok {
			return fmt.Errorf("No user %s in gum config", args[0])
		}

		pwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("Getwd failed: %v", err)
		}
		repo, err := git.PlainOpenWithOptions(pwd, &git.PlainOpenOptions{DetectDotGit: true})
		if err != nil {
			return fmt.Errorf("Open current git repository failed: %v", err)
		}

		cfg, err := repo.Config()
		if err != nil {
			return fmt.Errorf("Get current git repository config failed: %v", err)
		}

		cfg.User.Name = user.Name
		cfg.User.Email = user.Email

		if err := repo.SetConfig(cfg); err != nil {
			return fmt.Errorf("Set user to current git repository config failed: %v", err)
		}
		fmt.Printf("Set user %s to current git repository config success\n", user.Name)
		return nil
	},
}
