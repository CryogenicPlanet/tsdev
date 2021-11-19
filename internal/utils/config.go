package utils

import (
	"encoding/json"
	"fmt"
	"internal/types"
	"io/ioutil"
	"os"
)

func cwd() string {
	dir, err := os.Getwd()
	CheckErr(err)
	return dir
}

func readPackageJson() (types.PackageJSON, error) {
	config := types.PackageJSON{}

	configJson, err := os.ReadFile(cwd() + "/package.json")

	if err != nil {
		return config, err
	}

	err = json.Unmarshal(configJson, &config)

	if err != nil {
		return config, err
	}

	return config, nil

}

func GetName() (string, error) {
	config, err := readPackageJson()
	if err != nil {
		return "", err
	}
	return config.Name, nil
}

func ReadConfig() (types.ProjectConfig, error) {

	config, err := readPackageJson()

	if err != nil {
		var tsdev types.ProjectConfig

		return tsdev, err
	}

	return config.TSDEV, nil
}

var cachedConfig *types.ProjectConfig

// This will return if the template has pass through build or dev commands
func CommandPassThrough() bool {

	// Prevents subsequent reads
	if cachedConfig == nil {
		config, err := ReadConfig()
		if err != nil {
			fmt.Println("[WARN] could not read config")
			return false
		}
		cachedConfig = &config
	}

	if cachedConfig.Template == types.NextTemplate {
		return true
	} else if cachedConfig.Template == types.ReactTemplate {
		return true
	} else if cachedConfig.Template == types.ViteLibraryModeTemplate {
		return true
	}

	return false
}

func WriteConfig(config types.ProjectConfig) error {

	packageJson, err := readPackageJson()

	if err != nil {
		return err
	}

	packageJson.TSDEV = config
	configJson, _ := json.Marshal(config)
	err = ioutil.WriteFile(cwd()+"/package.json", configJson, 0644)

	CheckErr(err)

	return nil
}
