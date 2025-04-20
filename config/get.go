package config

import (
	"os/user"
	"path/filepath"
)

func getUserHome(username string) (string, error) {
	usr, err := user.Lookup(username)
	if err != nil {
		return "", err
	}

	return usr.HomeDir, nil
}

func getCfgPath(user string) (string, error) {
	home, err := getUserHome(user)
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config"), nil
}

func getDefaultMnt(user string) (string, error) {
	if user == "root" {
		return "/mtp/", nil
	} else {
		home, err := getUserHome(user)
		if err != nil {
			return "", err
		}
		return filepath.Join(home, "mtp"), nil
	}
}
