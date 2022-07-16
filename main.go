package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

// gum add: prompt for username and email input

// gum list

// gum delete <user>

// gum set <user>

var users map[string]user

var cfgFilePath string

func init() {
	dirname, _ := os.UserHomeDir()
	cfgFilePath = filepath.Join(dirname, ".gum")

	f, err := os.OpenFile(cfgFilePath, os.O_CREATE|os.O_RDWR, 0o666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	users, err = readUsers(f)
	if err != nil {
		fmt.Println(err)
	}
}

func set() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("ERROR: Getwd failed: %v", err)
		os.Exit(1)
	}
	repo, err := git.PlainOpenWithOptions(pwd, &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		panic(err)
	}

	cfg, err := repo.Config()
	if err != nil {
		panic(err)
	}

	cfg.User.Email = "xxxx@yy.com"
	fmt.Println(cfg.User.Name)
	fmt.Println(cfg.User.Email)

	if err := repo.SetConfig(cfg); err != nil {
		panic(err)
	}
}

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func list() (map[string]user, error) {
	f, err := os.Open(cfgFilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readUsers(f)
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

func addUser(u user) error {
	f, err := os.OpenFile(cfgFilePath, os.O_RDWR, 0o666)
	if err != nil {
		return err
	}
	defer f.Close()
	users, err := readUsers(f)
	if err != nil {
		return err
	}

	if _, ok := users[u.Name]; ok {
		return fmt.Errorf("user %s already exists", u.Name)
	}
	users[u.Name] = u
	b, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	return err
}

func deleteUser(u string) error {
	f, err := os.OpenFile(cfgFilePath, os.O_RDWR, 0o666)
	if err != nil {
		return err
	}
	defer f.Close()
	users, err := readUsers(f)
	if err != nil {
		return err
	}

	delete(users, u)
	b, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	return err
}

func main() {
	// fmt.Println(list())
	add(user{Name: "chenluxin", Email: "clx@qq.com"})
}
