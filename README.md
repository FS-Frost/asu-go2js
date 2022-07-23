# asu-go2js

[![Run tests](https://github.com/FS-Frost/asu-go2js/actions/workflows/main.yml/badge.svg)](https://github.com/FS-Frost/asu-go2js/actions/workflows/main.yml)

Asu is a library to work with subtitles on [ASS format](https://en.wikipedia.org/wiki/SubStation_Alpha).

asu-go2js is a port of [Asu](https://github.com/FS-Frost/Asu.Utilidades) (originally for .NET) written in Go and compiled to JavaScript with TypeScript support.

## Requeriments

-   [Go](https://go.dev) 1.17.9
-   [Node.js](https://nodejs.org) 16.13.1
-   [Yarn](https://classic.yarnpkg.com) 1.22.5
-   [GopherJS](https://github.com/gopherjs/gopherjs) 1.17.0

    ```shell
    go install github.com/gopherjs/gopherjs@latest
    ```

-   [GNU Make](https://www.gnu.org/software/make) 4.2.1

## Main make commands

```shell
# Compile Go to JavaScript
make build

# Run both Go and JavaScript tests
make test
```

## Folder structure

`go`: Go source code.

`go/exports`: GopherJS exports.

`js`: JavaScript source code.

`js/dist`: JavaScript files produced by GopherJS targeting [Node.js](https://nodejs.org) and [Deno](https://deno.land) runtimes.

`js/dist/asu.d.ts`: manually written type definitions.

`js/dist/browser`: JavaScript files compiled by [esbuild](https://esbuild.github.io) targeting browsers.

`js/tests`: JavaScript tests targeting `js/dist`.
