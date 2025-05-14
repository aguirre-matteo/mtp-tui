package config

import (
	"github.com/aguirre-matteo/mtp-tui/errors"
	"os"
	"strings"
)

func assertConfigPath(path string) error {
	if !strings.HasPrefix(path, "/") {
		return errors.ConfigPathRelativeError(path)
	}

	_, err := os.Stat(path)
	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		return errors.ConfigFileNotFoundError(path)
	}

	return err
}
