# Zero Config Modern Typescript Projects

I really like the idea behind https://github.com/jaredpalmer/tsdx but find in practise it uses slower, older tools. So I wanted to make something that plays a similar role with faster modern tooling. Think `esbuild`, `vite`, `next@12`, or even things like `bun` and `rome` in the future

The tool will also come with correct configuration for publishing packages using `dts-bundle`, `vite-library-mode` and finally, the tool will support monorepos too

## Get started

```
npm install -g tsdev

tsdev create {name}
```

## Base commands

- `tsdev` - Will run any `.ts` or `.tsx` file
- `tsdev dev`
- `tsdev build`
    - `--dts` Will emit `.d.ts` files
    - `--sourcemap` Will emit source maps
    - `--dist` Set output directory by default `dist`
- `tsdev lint lint:fix`