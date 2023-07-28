# Inspectors compiler

`inspc` is a tool to parse and compile inpectors for Go types.

## Usage

In general `inspc` require only path to source. There are three possible sources:
* package (recommended)
* directory
* single file

### Package

Packages is the better source since it provide the whole information about source code. But it has a restriction: path
should be relative to `$GOPATH/src` directory.

Example of usage:
```bash
inspc -package github.com/koykov/inspector/testobj
```
By default, `inspc` will write output to `$GOPATH/src/<package_name>_ins`, as you may see here [github.com/koykov/inspector/testobj_ins](https://github.com/koykov/inspector/tree/master/testobj_ins).
But you may specify specific destination by using `-destination` option (`-dst` alias available):
```bash
inspc -package github.com/koykov/inspector/testobj -destination /home/<user_name>/inspectors/testobj
```

### Directory

Example of usage:
```bash
inspc -directory path/to/directory -import <import_path>
```
`inspc` will write output to `path/to/directory_ins` by default, but you may specify destination directory by using
`-destination` option.

Option `-import` is required due to there is no way to define import path for Go types described in arbitrary directory.
This is a restriction of using `-directory` option. Also, directory must contain at least one .go file.

Option `-destination` is also available if you want to generate output to certain directory.

### File

File way is similar to directory, but single Go file must be specified as a source.

Example of usage:
```bash
inspc -file path/to/file.go -import <import_path>
```
`inspc` will write output to directory `path/to/file_ins` or any directory you specified using `-destination` option.

## Installation

Installation is pretty simple:
```bash
go install github.com/koykov/inspector/inspc@latest
```
As result in `$GOPATH/bin` will build `inspc` executable binary.

## Blacklist

If you need to skip compiling inspectors for certain types in the input packages, just list them in the text file, each
type name in the separate line and specify the path to file in `-black-list` argument.
