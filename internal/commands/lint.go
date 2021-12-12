package commands

import (
	"errors"
	"fmt"
	"internal/types"
	"internal/utils"
	"os"
)

// Change this to use https://rome.tools later

func HandleLintCommand(fix bool) error {

	cwd, err := os.Getwd()

	utils.CheckErr(err)

	if _, err := os.Stat(types.ESLINT_PATH); errors.Is(err, os.ErrNotExist) {

		if _, err := os.Stat(types.ESLINT_BACKUP_PATH); errors.Is(err, os.ErrNotExist) {

			return fmt.Errorf("cannot use lint without installing tsdev as a dependency")

		}

		if fix {
			utils.ExecWithOutput(cwd, "node", types.ESLINT_BACKUP_PATH, "--config", types.ESLINT_CONFIG_PATH, "src/*", "--fix")
		} else {
			utils.ExecWithOutput(cwd, "node", types.ESLINT_BACKUP_PATH, "--config", types.ESLINT_CONFIG_PATH, "src/*")
		}
		return nil
	}

	if fix {
		utils.ExecWithOutput(cwd, "node", types.ESLINT_PATH, "--config", types.ESLINT_CONFIG_PATH, "src/*", "--fix")
	} else {
		utils.ExecWithOutput(cwd, "node", types.ESLINT_PATH, "--config", types.ESLINT_CONFIG_PATH, "src/*")
	}

	return nil
}
