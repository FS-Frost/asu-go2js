# asu-go2js

[![Run tests](https://github.com/FS-Frost/asu-go2js/actions/workflows/main.yml/badge.svg)](https://github.com/FS-Frost/asu-go2js/actions/workflows/main.yml)

Asu is a library to work with subtitles on ASS format.

asu-go2js is a port of [Asu](https://github.com/FS-Frost/Asu.Utilidades) (originally for .NET) written in Go and compiled to JavaScript with TypeScript support.

## Requeriments

-   Go 1.17.5
-   Node.js 16.13.1
-   Yarn 1.22.5
-   GopherJS 1.17.0

```shell
go install github.com/gopherjs/gopherjs@latest
```

-   GNU Make 4.2.1

## Main make commands

```shell
# Compile Go to JavaScript
make build

# Run both Go and JavaScript tests
make test
```

## Folder structure

`(root)`: Go source code.

`exports`: GopherJS exports.

`js`: JavaScript source code.

`js/dist`: JavaScript files produced by GopherJS targeting Node and Deno runtimes.

`js/dist/asu.d.ts`: manually written type definitions.

`js/dist/browser`: JavaScript files compiled by esbuild targeting browsers.

`js/tests`: JavaScript tests targeting `js/dist`.
