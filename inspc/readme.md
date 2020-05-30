# Inspectors compiler

`inspc` is a tool to parse and compile inpectors for Go types.

## Usage

The tool requires input arg `-pkg` that should be a path to the Go package. The path should be relative to `$GOPATH/src`
directory, for example:
```bash
inspc -pkg github.com/koykov/inspector/testobj
```
You may see the result of execution of that command here [github.com/koykov/inspector/testobj_ins](https://github.com/koykov/inspector/tree/master/testobj_ins)

The tool by default makes the output around the input package with suffix `_ins`. You may specify destination directory
by input arg `-dst`, example:
```bash
inspc -pkg github.com/koykov/inspector/testobj -dst github.com/koykov/my_inspectors/testobj
```

Also supports blacklist files by passing the erg `-bl`. It's just a path to the blacklist file (see the section "Blacklist").

## Installation

Build the tool with the following commands:
```bash
go get github.com/koykov/inspector/inspc
go build -o $GOPATH/bin/inspc github.com/koykov/inspector/inspc
```
Make sure `$GOPATH/bin` is in the `$PATH` variable.

## Blacklist

If you need to skip compiling inspectors for certain types in the input packages, just list them in the text file, each
type name in the separate line and specify the path to file in `-bl` argument.
