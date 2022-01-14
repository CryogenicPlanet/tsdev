# Tsdev (Zero Config Modern Typescript Projects)


## Motivation

I really like the idea behind https://github.com/jaredpalmer/tsdx but I find it uses older tools and it ends up taking time for me to setup my typescript packages anyways. So I wanted to make something like it that fit my stack and preferred tools better, something that obfuscates all the configuration needed to run modern typescript applications and allows use to get started instantly.


## Get started

```bash
npm install -g tsdev-installer

tsdev create {name}
```


## Features

- Instantly bootstrap typescript apps for `express`, `react` and `next`
- Automatically handles all `eslint` and `prettier` config
- Comes with built in `dev` mode with in built watcher
- Use very fast built tools, like `esbuild` and `vite`
- Inbuilt config and setup for `tailwind`
- Has tooling for bundling `.dts` with `dts-bundle`
- Has defaults for publishing `react` packages with `vite-library-mode`
    - Adds support for `twind` (tailwind as css-in-js) so you can publish react packages while using tailwind with tiny bundles. This also has better support when use in `next`


## Base commands


```
➜ tsdev --help
NAME:
   tsdev - Zero config modern typescript tooling

USAGE:
   tsdev [global options] command [command options] Run a .ts file with zero config directly

COMMANDS:
   create    Create a new application
   dev       This run the app in dev mode with file watching
   build     This builds the app for production.
   prettier  Will run pretty-quick
   lint      Will lint the application
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --watch     Run in watch mode (default: false)
```

## Templates

Templates are not stored here, this just allows us to keep this repo clean, all the templates are here https://github.com/CryogenicPlanet/tsdev-templates/ and are hosted at https://tsdev.vercel.app/

## TODOs

Can find all TODOs on https://github.com/CryogenicPlanet/tsdev/issues/5