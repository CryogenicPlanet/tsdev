# Zero Config Modern Typescript Projects

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
   tsdev [global options] command [command options] [arguments...]

COMMANDS:
   create    Create a new application
   dev       This run the app in dev mode with file watching
   build     This builds the app for production.
   prettier  Will run pretty-quick
   lint      Will lint the application
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## Templates

Templates are not stored here, this just allows us to keep this repo clean, all the templates are here https://github.com/CryogenicPlanet/tsdev-templates/ and are hosted at https://tsdev.vercel.app/

## TODOs

Some of these might happen sooner than others. If you want something to be prioritized make an issue

- [] Monorepo support, automatically bootstrap monorepos
- [] Support for `publishConfig` beyond just `pnpm`. This is super useful for typescript packages in monorepos
- [] Add basic CI setup for github actions (with all three package managers)
- [] Add defaults for publishing normal non-react packages
- [] Allow overwriting or extending eslint configs
- [] Add `graphql` template using [tsgql](https://github.com/modfy/tsgql). This will be `graphql` without having to write any `graphql` code at all, just typescript.
- [] Add `prisma` batteries
- [] Add support for all nextjs examples from https://github.com/vercel/next.js/tree/canary/examples
- [] `tsdev filename` Automatically run any `.ts` or `.tsx` file with zero config
    - The `.ts` part of this is easy, it is basically what `tsdev dev filename` is
    - The `.tsx` part is a bit more complicated, and will require making a custom version of `vite` that just runs file without a config file or a `.html` file or having to use `react-dom` yourself
- [] Make `vite` default template use filesystem routing 
- [] Clone [`bun run` feature](https://twitter.com/jarredsumner/status/1454218996983623685?s=20) to allow really fast `npm run commands` 
- [] Add support for [bun](https://bun.sh) once it becomes more stable 