package utils

import "internal/types"

func GetPackageManager(packageManger types.PackageManagerType) string {

	switch packageManger {
	case types.Pnpm:
		{
			return "pnpm"
		}
	case types.Yarn:
		{
			return "yarn"
		}
	case types.Npm:
		{
			return "npm"
		}
	default:
		{
			return "npm"
		}
	}
}
