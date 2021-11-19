package commands

import (
	"errors"
	"fmt"
	"internal/utils"
	"os"
)

const ESLINT_PATH = "node_modules/tsdev/node_modules/eslint/bin/eslint.js"

const ESLINT_CONFIG_PATH = "node_modules/tsdev/static/config/.eslintrc"

// Change this to use https://rome.tools later

func HandleLintCommand(fix bool) error {

	cwd, err := os.Getwd()

	utils.CheckErr(err)

	if _, err := os.Stat(ESLINT_PATH); errors.Is(err, os.ErrNotExist) {

		return fmt.Errorf("cannot use lint without installing tsdev as a dependency")
	}

	if fix {
		utils.ExecWithOutput(cwd, "node", ESLINT_PATH, "--config", ESLINT_CONFIG_PATH, "src/*", "--fix")
	} else {
		utils.ExecWithOutput(cwd, "node", ESLINT_PATH, "--config", ESLINT_CONFIG_PATH, "src/*")
	}

	return nil
}
