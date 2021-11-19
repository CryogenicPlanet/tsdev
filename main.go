package main

import (
	"internal/utils"
	"os"
)

func main() {

	app, err := SetupCliApp()
	utils.CheckErr(err)

	err = app.Run(os.Args)
	utils.CheckErr(err)
}
