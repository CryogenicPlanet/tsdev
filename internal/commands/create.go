package commands

import (
	"encoding/json"
	"fmt"
	"internal/types"
	"internal/utils"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/AlecAivazis/survey/v2"
)

var projectConfig types.ProjectConfig

// Prompts to get project config
func getProjectConfig() {

	template := "basic"
	packageManger := "yarn"

	prompt := &survey.Select{
		Message: "Choose a template",
		Options: []string{"basic", "react"},
	}

	survey.AskOne(prompt, &template)

	if template == "react" {
		nextJs := false // default vite
		libraryMode := false

		// React template

		projectConfig.Template = types.ReactTemplate

		prompt := &survey.Confirm{
			Message: "Do you want to use nextjs",
			Default: false,
		}

		survey.AskOne(prompt, &nextJs)

		if nextJs {
			projectConfig.Template = types.NextTemplate
		} else {
			prompt := &survey.Confirm{
				Message: "Do you want to publish this package?",
				Default: false,
			}
			survey.AskOne(prompt, &libraryMode)

			if libraryMode {
				projectConfig.Template = types.ViteLibraryModeTemplate
			}
		}

		prompt = &survey.Confirm{
			Message: "Do you want to use tailwindcss",
			Default: false,
		}

		survey.AskOne(prompt, &projectConfig.TailwindCss)

	} else {
		express := true

		confirmPrompt := &survey.Confirm{
			Message: "Do you want to use express",
			Default: true,
		}

		survey.AskOne(confirmPrompt, &express)

		if express {
			projectConfig.Template = types.ExpressTemplate
		}
	}

	prompt = &survey.Select{
		Message: "Choose a package manager",
		Options: []string{"yarn", "pnpm", "npm"},
	}

	survey.AskOne(prompt, &packageManger)

	if packageManger == "yarn" {
		projectConfig.PackageManager = types.Yarn
	} else if packageManger == "pnpm" {
		projectConfig.PackageManager = types.Pnpm
	} else {
		projectConfig.PackageManager = types.Npm
	}

	fmt.Println("Chosen template is", projectConfig.Template)

}

// Will create project dir
func createDir(name string) error {
	// Check if directory already exists

	if _, err := os.Stat(name); os.IsNotExist(err) {

		// Create directory
		err := os.Mkdir(name, 0755)
		if err != nil {
			return err
		}
	} else {
		log.Fatalln("Directory " + name + " already exits, cannot create a project with this name")
	}

	return nil
}

// Will generate package json
func generatePackageJson(name string) {

	var prepareScript string

	switch projectConfig.PackageManager {
	case types.Pnpm:
		{
			prepareScript = "pnpm build && pnpm dts"
		}
	case types.Yarn:
		{
			prepareScript = "yarn build && yarn dts"
		}
	case types.Npm:
		{
			prepareScript = "npm run build && npm run dts"
		}
	default:
		{
			prepareScript = "npm run build && npm run dts"
		}
	}

	// Default package json
	packageJson := types.PackageJSON{
		Name:            name,
		License:         "MIT",
		Version:         "0.1.0",
		Main:            "dist/index.js",
		Typings:         "dist/index.d.ts",
		Module:          "dist/index.es.js",
		Files:           []string{"dist"},
		Scripts:         map[string]string{"start": "tsdev start", "build": "tsdev build", "dev": "tsdev dev", "lint": "tsdev lint", "dts": "tsdev dts", "prepare": prepareScript},
		DevDependencies: map[string]string{"typescript": "latest", "husky": "latest", "prettier": "latest", "prettier-config-standard": "latest", "@cryogenicplanet/tsdev": "latest"},
		Husky:           map[string]map[string]string{"hooks": {"pre-commit": "tsdev prettier", "pre-push": "tsdev lint"}},
		Engines:         map[string]string{"node": ">12"},
		Prettier:        "prettier-config-standard",
	}

	reactDeps := map[string]string{"react": "^17.0.0", "react-dom": "^17.0.0"}
	reactDevDeps := map[string]string{"@types/react": "^17.0.0", "@types/react-dom": "^17.0.0"}
	nodeDevDeps := map[string]string{"@types/node": "^14"}

	switch projectConfig.Template {
	case types.ReactTemplate, types.ViteLibraryModeTemplate:
		viteScripts := map[string]string{"start": "vite preview", "build": "vite build", "dev": "vite"}
		viteDevDeps := map[string]string{"vite": "latest", "@vitejs/plugin-react-refresh": "latest", "vite-tsconfig-paths": "latest", "dts-bundle": "latest"}

		// Pass on scripts directly to vite
		// This saves us work of running a node instance from go
		packageJson.Scripts = utils.MergeStringMaps(packageJson.Scripts, viteScripts)
		packageJson.Dependencies = reactDeps
		packageJson.DevDependencies = utils.MergeStringMaps(packageJson.DevDependencies, viteDevDeps, reactDevDeps)

		if projectConfig.Template == types.ViteLibraryModeTemplate {
			// Vite library mode exports https://vitejs.dev/guide/build.html#library-mode
			packageJson.Exports = map[string]map[string]string{".": {"imports": "./dist/index.es.js", "exports": "./dist/index.umd.js"}}
		}

	case types.NextTemplate:
		// Pass on scripts directly to next
		// This saves us work of running a node instance from go
		nextScripts := map[string]string{"start": "next start", "build": "next build", "dev": "next"}
		packageJson.Dependencies = utils.MergeStringMaps(reactDeps, map[string]string{"next": "latest"})
		packageJson.DevDependencies = utils.MergeStringMaps(packageJson.DevDependencies, reactDevDeps)
		packageJson.Scripts = utils.MergeStringMaps(packageJson.Scripts, nextScripts)
	case types.ExpressTemplate:
		packageJson.Dependencies = map[string]string{"express": "latest", "cors": "latest", "morgan": "latest"}
		packageJson.DevDependencies = utils.MergeStringMaps(packageJson.DevDependencies, nodeDevDeps, map[string]string{"@types/express": "latest", "@types/cors": "latest", "@types/morgan": "latest"})
	case types.BasicTemplate:
		packageJson.DevDependencies = utils.MergeStringMaps(packageJson.DevDependencies, nodeDevDeps)
	default:
	}

	if projectConfig.TailwindCss {
		packageJson.DevDependencies = utils.MergeStringMaps(packageJson.DevDependencies, setupTailwindPackages(projectConfig.Template == types.ViteLibraryModeTemplate))
	}

	packageJson.TSDEV = projectConfig

	configJson, _ := json.Marshal(packageJson)
	err := ioutil.WriteFile(dirPath(name)+"/package.json", configJson, 0644)

	utils.CheckErr(err)

}

var setupWg sync.WaitGroup

func setupTailwindPackages(libraryMode bool) map[string]string {

	defaultTailwind := map[string]string{"tailwindcss": "latest", "autoprefixer": "latest", "postcss": "latest", "@tailwindcss/forms": "latest"}

	if libraryMode {
		// setup twind with tailwind

		defaultTailwind = utils.MergeStringMaps(defaultTailwind, map[string]string{"twind": "latest", "@twind/forms": "latest"})
	}

	return defaultTailwind
}

func downloadTemplate(dirName string) {

	switch projectConfig.Template {
	case types.BasicTemplate:
		utils.DownloadArchive("https://tsdev.vercel.app/templates/basic.zip", dirName)
	case types.ExpressTemplate:
		utils.DownloadArchive("https://tsdev.vercel.app/templates/express.zip", dirName)
	case types.ReactTemplate:
		utils.DownloadArchive("https://tsdev.vercel.app/templates/vite.zip", dirName)
	case types.NextTemplate:
		utils.DownloadArchive("https://tsdev.vercel.app/templates/next.zip", dirName)
	case types.ViteLibraryModeTemplate:
		utils.DownloadArchive("https://tsdev.vercel.app/templates/viteLib.zip", dirName)
	default:
	}
	setupWg.Done()
}

// This will download the tailwind config files
func setupTailwind(dirName string, libraryMode bool) {

	utils.DownloadFile("https://tsdev.vercel.app/config/tailwind.config.js", dirName+"/tailwind.config.js")
	utils.DownloadFile("https://tsdev.vercel.app/config/postcss.config.js", dirName+"/postcss.config.js")
	os.MkdirAll(dirName+"/styles", 0777)
	utils.DownloadFile("https://tsdev.vercel.app/config/tailwind.css", dirName+"/styles/tailwind.css")

	if libraryMode {
		utils.DownloadFile("https://tsdev.vercel.app/config/twindSetup.ts", dirName+"/twindSetup.ts")
	}
	setupWg.Done()
}

func gitInit(path string) error {

	utils.ExecWithOutput(path, "git", "init")
	utils.ExecWithOutput(path, "git", "add", ".")
	utils.ExecWithOutput(path, "git", "commit", "-m", "🚀 Tsdev setup")

	return nil
}

// Will install the packages
func installPackages(path string) {

	fmt.Println("Installing Packages", path)

	switch projectConfig.PackageManager {
	case types.Yarn:
		//
		utils.ExecWithOutput(path, "yarn")
	case types.Pnpm:
		//
		utils.ExecWithOutput(path, "pnpm", "i")
	case types.Npm:
		//
		utils.ExecWithOutput(path, "npm", "i")
	}

	setupWg.Done()
}

func dirPath(name string) string {
	cwd, err := os.Getwd()
	utils.CheckErr(err)
	return cwd + "/" + name
}

// Will Handle the Create Command
func HandleCreateCommand(name string) error {

	err := createDir(name)

	utils.CheckErr(err)

	getProjectConfig()

	generatePackageJson(name)

	setupWg.Add(1)

	go downloadTemplate(dirPath(name))

	setupWg.Add(1)

	go installPackages(dirPath(name))

	setupWg.Wait()

	// Tailwind would overwrite the dummy files
	if projectConfig.TailwindCss {
		setupWg.Add(1)
		go setupTailwind(dirPath(name), projectConfig.Template == types.ViteLibraryModeTemplate)
	}

	setupWg.Wait()

	gitInit(dirPath(name))

	return nil
}
