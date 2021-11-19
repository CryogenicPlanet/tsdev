package utils

import (
	"encoding/json"
	"internal/types"
	"io/ioutil"
)

func ReadPackageJson(path string) (types.PackageJSON, error) {

	plain, _ := ioutil.ReadFile(path)

	packageJson := types.PackageJSON{}
	err := json.Unmarshal(plain, &packageJson)

	if err != nil {
		return packageJson, err
	}
	return packageJson, nil
}
