package commands

import (
	"fmt"
	"internal/utils"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/evanw/esbuild/pkg/api"
)

// This will be a default runner for tsdev
// It will have two states
// 1. run esbuild on file in watch mode and run a node server
//	the compiled node file will be in a tmp folder
// 	This will allow you to directly execute any .ts file, anywhere with zero config
// 2. run vite on a file in watch mode and run a vite server on any .tsx file
//  This will be more challenging but will setup an invisible vite env and just run the .tsx file directly
// 	You can write zero config react code like this

var watchWg sync.WaitGroup

func HandleDefault(watch bool, paths []string) error {

	tsFiles := []string{}

	for _, path := range paths {
		ext := filepath.Ext(path)

		if ext == ".ts" {
			// Run file in node
			tsFiles = append(tsFiles, path)

		} else if ext == ".tsx" {
			// TODO Run file in browser using vite
			// Will run these files parrallely using vite
			fmt.Println("Tsx support coming soon...")
		} else {
			fmt.Println("[WARN] Tsdev can only be used with typescript files (.ts or .tsx)")
		}
	}

	if watch {
		watchWg.Add(1)
	}

	handleTsFiles(tsFiles, watch)

	watchWg.Wait() // Will not finish
	return nil
}

func runJSFiles(tempDir string, jsFile string) error {

	return utils.ExecWithOutput(tempDir, "node", jsFile)
}

func handleTsFiles(tsFiles []string, watch bool) {

	dir, err := os.MkdirTemp("", "tsdev")

	if err != nil {
		// fmt.Println("Could not create temp dir")
		panic(err)
	}

	defer os.RemoveAll(dir)

	result := api.BuildResult{}

	if watch {
		result = api.Build(api.BuildOptions{
			EntryPoints: tsFiles,
			Target:      api.ESNext,
			Bundle:      true,
			Write:       true,
			Format:      api.FormatESModule,
			Outdir:      dir,
			Platform:    api.PlatformNode, // Both browser templates next and vite have their own runners
			Loader: map[string]api.Loader{
				".js": api.LoaderJSX,
			},
			Watch: &api.WatchMode{
				OnRebuild: func(br api.BuildResult) {
					for _, file := range tsFiles {
						jsFile := strings.Replace(file, "ts", "js", -1)
						runJSFiles(dir, jsFile)
					}
					fmt.Println("Watching ...")
				},
			},
		})
	} else {
		result = api.Build(api.BuildOptions{
			EntryPoints: tsFiles,
			Target:      api.ESNext,
			Bundle:      true,
			Write:       true,
			Format:      api.FormatESModule,
			Outdir:      dir,
			Platform:    api.PlatformNode, // Both browser templates next and vite have their own runners
			Loader: map[string]api.Loader{
				".js": api.LoaderJSX,
			},
		})
	}

	if len(result.Errors) > 0 {

		for _, err := range result.Errors {
			fmt.Println("Error", err.Text, err.Location.File, err.Location.Line)
		}

		os.Exit(1)
	}

	for _, file := range tsFiles {
		jsFile := strings.Replace(file, "ts", "js", -1)
		runJSFiles(dir, jsFile)
	}

	if watch {
		fmt.Println("Watching ...")
	}
	watchWg.Wait()
}
