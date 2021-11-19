module commands

go 1.17

require (
	github.com/AlecAivazis/survey/v2 v2.3.2
	internal/types v1.0.0
	internal/utils v1.0.0

)

require (
	github.com/evanw/esbuild v0.13.13 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/mattn/go-isatty v0.0.8 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	golang.org/x/sys v0.0.0-20211113001501-0c823b97ae02 // indirect
	golang.org/x/term v0.0.0-20210503060354-a79de5458b56 // indirect
	golang.org/x/text v0.3.3 // indirect
)

replace internal/utils => ../utils

replace internal/types => ../types
