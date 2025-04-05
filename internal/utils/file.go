package utils

import (
	"errors"
	"os"
)

func CheckWorkspaceExist(workspace string) error {
	if workspace == "" {
		return nil
	}

	if _, err := os.Stat(workspace); os.IsNotExist(err) {
		return errors.New("output directory does not exist: " + workspace)
	} else if err != nil {
		return err
	}

	return nil
}
