package commands

import (
	"fmt"
	"internal/utils"
	"os"
	"sync"

	"github.com/evanw/esbuild/pkg/api"
)

func nodeJsRunner(cwd string) {
	utils.ExecWithOutput(cwd, "node", "dist/index.js")
}

var devWg sync.WaitGroup

func HandleDevCommand(entryPoint string) error {

	devWg.Add(1)

	if utils.CommandPassThrough() {
		fmt.Println("Run yarn dev instead")
		return nil
	}

	cwd, err := os.Getwd()

	utils.CheckErr(err)

	if entryPoint == "" {
		// Default no args
		entryPoint = "src/index.ts"
	}

	result := api.Build(api.BuildOptions{
		EntryPoints: []string{entryPoint},
		// EntryPoints: []string{"test/monorepo/packages/package-b/src/App.tsx"},
		Target: api.ESNext,
		Bundle: true,
		Write:  true,
		Watch: &api.WatchMode{OnRebuild: func(result api.BuildResult) {
			if len(result.Errors) > 0 {
				for _, err := range result.Errors {
					fmt.Println("Error", err.Text, err.Location.File, err.Location.Line)
				}
			} else {
				nodeJsRunner(cwd)
			}
		}},
		Format:  api.FormatCommonJS,
		Outfile: cwd + "/dist/index.js",
		Define: map[string]string{
			"process.env.NODE_ENV": "\"development\"",
		},
		Platform: api.PlatformNode, // Both browser templates and vite have their own runners
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
	nodeJsRunner(cwd)

	fmt.Println("Watching for changes ...")
	devWg.Wait()
	return nil
}
