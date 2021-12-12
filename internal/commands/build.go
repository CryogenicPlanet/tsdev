package commands

import (
	"errors"
	"fmt"
	"internal/types"
	"internal/utils"
	"os"
	"sync"

	"github.com/evanw/esbuild/pkg/api"
)

var buildWg sync.WaitGroup

func emitDts(cwd string, name string) error {
	if _, err := os.Stat(types.BUNDLE_DTS_PATH); errors.Is(err, os.ErrNotExist) {

		fmt.Println("[WARN] You can only use --dts flag if you have installed tsdev as a dependency")
		buildWg.Done()
		return nil
	}
	utils.ExecWithOutput(cwd, "tsc", "--outDir", "dist/src/")
	return bundleDts(cwd, name)
}

func bundleDts(cwd string, name string) error {

	if _, err := os.Stat(types.BUNDLE_DTS_PATH); errors.Is(err, os.ErrNotExist) {

		if _, err := os.Stat(types.BUNDLE_BACKUP_DTS_PATH); errors.Is(err, os.ErrNotExist) {
			fmt.Println("[WARN] You can only use --dts flag if you have installed tsdev as a dependency")
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

func buildCJS(entryPoint string, cwd string) {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{entryPoint},
		Target:      api.ESNext,
		Bundle:      true,
		Write:       true,
		Format:      api.FormatCommonJS,
		Outfile:     cwd + "/dist/index.js",
		Platform:    api.PlatformNode, // Both browser templates next and vite have their own runners
		Loader: map[string]api.Loader{
			".js": api.LoaderJSX,
		},
	})

	if len(result.Errors) > 0 {

		for _, err := range result.Errors {
			fmt.Println("Error", err.Text, err.Location.File, err.Location.Line)
		}

		os.Exit(1)
	}
	buildWg.Done()
}

func buildESM(entryPoint string, cwd string) {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{entryPoint},
		Target:      api.ESNext,
		Bundle:      true,
		Write:       true,
		Format:      api.FormatESModule,
		Outfile:     cwd + "/dist/index.es.js",
		Platform:    api.PlatformNode, // Both browser templates next and vite have their own runners
		Loader: map[string]api.Loader{
			".js": api.LoaderJSX,
		},
	})

	if len(result.Errors) > 0 {

		for _, err := range result.Errors {
			fmt.Println("Error", err.Text, err.Location.File, err.Location.Line)
		}

		os.Exit(1)
	}
	buildWg.Done()
}

func HandleBuildCommand(entryPoint string, dts bool) error {

	if utils.CommandPassThrough() {
		fmt.Println("Run yarn build instead")
		return nil
	}

	cwd, err := os.Getwd()

	utils.CheckErr(err)

	if entryPoint == "" {
		// Default no args
		entryPoint = "src/index.ts"
	}

	buildWg.Add(1)
	go buildCJS(entryPoint, cwd)
	buildWg.Add(1)
	go buildESM(entryPoint, cwd)

	if dts {
		buildWg.Add(1)
		name, err := utils.GetName()
		utils.CheckErr(err)
		go emitDts(cwd, name)
	}
	buildWg.Wait()

	return nil
}
