package types

type TemplateType int64

const (
	BasicTemplate TemplateType = iota
	ReactTemplate              // This is a vite template
	NextTemplate
	ViteLibraryModeTemplate
	ExpressTemplate
)

type Template interface {
	Template() TemplateType
}

type PackageManagerType int64

const (
	Npm  PackageManagerType = iota
	Yarn                    // This is a vite template
	Pnpm
)

type PackageManager interface {
	PackageManager() PackageManagerType
}
