package commands

import (
	"errors"
	"fmt"
	"internal/utils"
	"os"
)

const PRETTY_QUICK_PATH = "node_modules/tsdev/node_modules/pretty-quick/bin/pretty-quick.js"

func HandlePrettierCommand() error {

	if _, err := os.Stat(PRETTY_QUICK_PATH); errors.Is(err, os.ErrNotExist) {

		return fmt.Errorf("cannot use prettier without installing tsdev as a dependency")
	}

	cwd, err := os.Getwd()

	utils.CheckErr(err)

	return utils.ExecWithOutput(cwd, "node", PRETTY_QUICK_PATH, "--staged")

}
