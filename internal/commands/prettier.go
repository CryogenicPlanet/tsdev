package commands

import (
	"errors"
	"fmt"
	"internal/types"
	"internal/utils"
	"os"
)

func HandlePrettierCommand() error {

	if _, err := os.Stat(types.PRETTY_QUICK_PATH); errors.Is(err, os.ErrNotExist) {

		return fmt.Errorf("cannot use prettier without installing tsdev as a dependency")
	}

	cwd, err := os.Getwd()

	utils.CheckErr(err)

	return utils.ExecWithOutput(cwd, "node", types.PRETTY_QUICK_PATH, "--staged")

}
