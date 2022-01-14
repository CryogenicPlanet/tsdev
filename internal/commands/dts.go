package commands

import (
	"errors"
	"fmt"
	"internal/types"
	"internal/utils"
	"os"
)

func EmitDts(cwd string, name string) error {

	packageJson, err := utils.ReadPackageJson(cwd + "/package.json")

	if err != nil {
		fmt.Println("Could not read package.json", cwd+"./package.json")
		buildWg.Done()
		return err
	}

	if _, err := os.Stat(types.BUNDLE_DTS_PATH); errors.Is(err, os.ErrNotExist) {

		if _, err := os.Stat(types.BUNDLE_BACKUP_DTS_PATH); errors.Is(err, os.ErrNotExist) {
			fmt.Println("[WARN] You can only use --dts flag if you have installed tsdev as a dependency. Error:", err)
			buildWg.Done()
			return errors.New("cannot find bundle-dts path")
		}
	}
	fmt.Println("Package manager is ", packageJson.TSDEV.PackageManager)
	err = utils.ExecWithOutput(cwd, utils.GetPackageManager(packageJson.TSDEV.PackageManager), "tsc", "--outDir", "dist/src/")
	if err != nil {
		buildWg.Done()
		return err
	}
	return bundleDts(cwd, name)
}

func bundleDts(cwd string, name string) error {

	if _, err := os.Stat(types.BUNDLE_DTS_PATH); errors.Is(err, os.ErrNotExist) {

		if _, err := os.Stat("./" + types.BUNDLE_BACKUP_DTS_PATH); errors.Is(err, os.ErrNotExist) {
			fmt.Println("[WARN] You can only use --dts flag if you have installed tsdev as a dependency. Error:", err)
			return errors.New("cannot find bundle-dts path")
		}
		utils.ExecWithOutput(cwd, "node", types.BUNDLE_BACKUP_DTS_PATH, "--name", name, "--main", "dist/src/index.d.ts", "--out", "../index.d.ts")
		buildWg.Done()
		return nil
	}

	utils.ExecWithOutput(cwd, "node", types.BUNDLE_DTS_PATH, "--name", name, "--main", "dist/src/index.d.ts", "--out", "../index.d.ts")
	buildWg.Done()
	return nil
}

func RunDts() error {

	cwd, err := os.Getwd()

	utils.CheckErr(err)
	name, err := utils.GetName()
	utils.CheckErr(err)
	buildWg.Add(1)

	go EmitDts(cwd, name)

	buildWg.Wait()
	return nil
}
