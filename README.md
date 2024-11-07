[![Go](https://github.com/lil5/typex2/actions/workflows/go.yml/badge.svg)](https://github.com/lil5/typex2/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/lil5/typex2)](https://goreportcard.com/report/github.com/lil5/typex2)
[![Coverage Status](https://coveralls.io/repos/github/lil5/typex2/badge.svg?branch=master)](https://coveralls.io/github/lil5/typex2?branch=master)

# TypeX 2

Reads a go package's types and export them to typescript interfaces and basic types.

This is the follow up of https://github.com/dtgorski/typex

## Installation

```
go install github.com/lil5/typex2@1.1.1
```

## Usage

```
$ typex2 ./examples
```

This will do the following;

1. Read all go files inside the path specified (must use one [package name](https://blog.golang.org/package-names)).
2. Generate typescript [types](https://www.typescriptlang.org/docs/handbook/basic-types.html) and [interfaces](https://www.typescriptlang.org/docs/handbook/interfaces.html) from said go files.
3. Write generated content into `./examples/index.ts`.

## TypeScript type mapping

> copy from Typex readme
>
> https://github.com/dtgorski/typex#typescript-type-mapping

TypeScript (resp. JavaScript aka ECMAScript) lacks a native integer number type.
The numeric type provided there is inherently a 64 bit float.
You should keep this in mind when working with exported numeric types - this includes `byte` and `rune` type aliases as well.    

|Go native type|TypeScript type
| --- | ---
|`bool`|`boolean`
|`string`|`string`
|`map`|`Record<K, V>`
|`interface`|`Record<string, any>`
|`struct` `(named)`|`T`
|`struct` `(anonymous)`|`{}`
|`array` `(slice)`|`T[]`
|`complex`[`64`&vert;`128`]|`any`
|`chan`, `func`, `interface`|`any`
|`int`[`8`&vert;`16`&vert;`32`&vert;`64`]|`number`
|`uint`[`8`&vert;`16`&vert;`32`&vert;`64`]|`number`
|`byte`(=`uint8`)|`number`
|`rune`(=`int32`)|`number`
|`float`[`32`&vert;`64`]|`number`
|`uintptr`|`any`
|`*`|`T \| null`

## Differences between typex2 and typex

1. Code legibility.
   - Typex2 uses go's strengths in functional programming.
   - It also improves seperation of concerns, the reading of the go structs and types is seperated from the generation of the types in said language.
2. Generated code is instantly written to that same path instead of out putting it to the console.
3. Pointers are possibly nil in go, thus implemented in Typex2.
