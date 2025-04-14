[![Go](https://github.com/lil5/typex2/actions/workflows/go.yml/badge.svg)](https://github.com/lil5/typex2/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/lil5/typex2)](https://goreportcard.com/report/github.com/lil5/typex2)

# TypeX 2

Reads a go package's types and export them to TypeScript interfaces and basic types.

🆕 Now with support for Dart, Kotlin and Swift!

Great for keeping frontend types the same as your go backend.

This is the follow up of https://github.com/dtgorski/typex

## Installation

```
go install github.com/lil5/typex2@1.3.0
```

## Usage

```
$ typex2 -l typescript -i ./examples -o ./examples/typex2.ts
```

This will do the following;

1. Read all go files inside the path specified (must use one [package name](https://blog.golang.org/package-names)).
2. Generate typescript [types](https://www.typescriptlang.org/docs/handbook/basic-types.html) and [interfaces](https://www.typescriptlang.org/docs/handbook/interfaces.html) from said go files.
3. Write generated content into `./examples/typex2.ts`.

## Help

```
$ typex2 -h
NAME:
   typex2 - Convert go structs to other language types

USAGE:
   typex2 [global options] command [command options] path

DESCRIPTION:
   Useful for generating types from golang json RestAPI projects for a frontend to ingest.

   Example:
   typex2 -l kotlin -i . -o ./classes.kotlin

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --lang value, -l value, --language value  Language to generate to [typescript, dart, kotlin, swift] (default: typescript)
   --input value, -i value                   Input directory to read go types from (default: ".")
   --output value, -o value                  Output path & file to write translated types to
   --help, -h                                show help
```

## Language type mapping

> copy from Typex readme
>
> https://github.com/dtgorski/typex#typescript-type-mapping

TypeScript (resp. JavaScript aka ECMAScript) lacks a native integer number type.
The numeric type provided there is inherently a 64 bit float.
You should keep this in mind when working with exported numeric types - this includes `byte` and `rune` type aliases as well.    

|Go native types|TypeScript|Dart 🆕|Kotlin 🆕|Swift 🆕|Rust 🆕|
| --- | --- | --- | --- | --- | --- |
|`bool`|`boolean`|`bool`|`Boolean`|`Bool`|`bool`|
|`string`|`string`|`String`|`String`|`String`|`String`|
|`map`|`Record<K, V>`|`Map<K, V>`|`Map<K, V>`|`Dictionary<K, V>`|`Map<K, V>`|
|`interface`|`Record<string, any>`|`Map<String, dynamic>`|`Map<String, Any>`|`Dictionary<String, Any>`|`Map<K, V>`|
|`struct` `(named)`|`T`|`dynamic`|`Any`|`Any`|`<T1>`|
|`struct` `(anonymous)`|`{}`|`dynamic`|`Any`|`Any`|`<T2>`|
|`array` `(slice)`|`T[]`|`List<T>`|`Array<T>`|`Array<T>`|`Vec<T>`|
|`complex`[`64`&vert;`128`]|`any`|`dynamic`|`Float`|`Float`|`f64`|
|`chan`, `func`, `interface`|`any`|`dynamic`|`Any`|`Any`|`T3`|
|`int`[`8`&vert;`16`&vert;`32`&vert;`64`]|`number`|`int`|`Int`|`Int`|`i8,i16,i32,i64`|
|`uint`[`8`&vert;`16`&vert;`32`&vert;`64`]|`number`|`int`|`Int`|`Int`|`u8,u16,u32,u64`|
|`byte`(=`uint8`)|`number`|`int`|`Int`|`Int`|`i8`|
|`rune`(=`int32`)|`number`|`dynamic`|`Any`|`Any`|`i32`|
|`float`[`32`&vert;`64`]|`number`|`double`|`Double`|`Double`|`f32,f64`|
|`uintptr`|`any`|`dynamic`|`Any`|`Any`|`T4`|
|`*`|`T \| null`|`T?`|`T?`|`T?`|`Option<T5>`|

## Differences between typex2 and typex

1. Code legibility.
   - Typex2 uses go's strengths in functional programming.
   - It also improves separation of concerns, the reading of the go structs and types is separated from the generation of the types in said language.
2. Generated code is instantly written to that same path instead of out putting it to the console.
3. Pointers are possibly nil in go, thus implemented in Typex2.
4. Ability to generate for other languages than TypeScript: Dart, Kotlin, Swift (it's only 300 lines to add a new language)
