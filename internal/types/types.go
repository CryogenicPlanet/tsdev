package types

type PackageJSON struct {
	Name            string                       `json:"name"`
	Dependencies    map[string]string            `json:"dependencies,omitempty"`
	DevDependencies map[string]string            `json:"devDependencies,omitempty"`
	Scripts         map[string]string            `json:"scripts,omitempty"`
	Files           []string                     `json:"files,omitempty"`
	Engines         map[string]string            `json:"engines,omitempty"`
	Main            string                       `json:"main,omitempty"`
	Typings         string                       `json:"typings,omitempty"`
	Husky           map[string]map[string]string `json:"husky,omitempty"`
	Version         string                       `json:"version,omitempty"`
	License         string                       `json:"license,omitempty"`
	Module          string                       `json:"module,omitempty"`
	Exports         map[string]map[string]string `json:"exports,omitempty"`
	Prettier        string                       `json:"prettier,omitempty"`
	TSDEV           ProjectConfig                `json:"tsdev,omitempty"`
}

type ProjectConfig struct {
	Template       TemplateType       `json:"template,omitempty"`
	PackageManager PackageManagerType `json:"packageManager,omitempty"`
	TailwindCss    bool               `json:"tailwindcss,omitempty"`
}
