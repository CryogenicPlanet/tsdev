package utils

import "internal/types"

func GetPackageManager(packageManger types.PackageManagerType, run bool) string {

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
			// npm run cannot run node modules, the way yarn or pnpm can
			if run {
				return "yarn"
			}
			return "npm"
		}
	default:
		{
			// npm run cannot run node modules, the way yarn or pnpm can
			if run {
				return "yarn"
			}
			return "npm"
		}
	}
}
