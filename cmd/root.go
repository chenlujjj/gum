package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	users       map[string]user
	cfgFilePath string

	rootCmd = &cobra.Command{
		Use:   "gum",
		Short: "gum is your Git User Manager",
	}
)

func init() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	cfgFilePath = filepath.Join(home, ".gum")

	f, err := os.OpenFile(cfgFilePath, os.O_RDWR, 0o666)
	if errors.Is(err, os.ErrNotExist) {
		users = map[string]user{}
		return
	}

	defer f.Close()
	users, err = readUsers(f)
	cobra.CheckErr(err)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func readUsers(r io.Reader) (map[string]user, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var users map[string]user
	err = json.Unmarshal(b, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
