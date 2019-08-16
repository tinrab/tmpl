# tmpl

A tool for generating parameterized files from templates.

## Install

Binary downloads can be found on [releases](https://github.com/tinrab/tmpl/releases) page.

Or install it manually:

```
go get github.com/tinrab/tmpl/cmd/tmpl
```

## Usage

Example:

```
$ tmpl -p "./test/*.params.*" -s ./test/deployment.tmpl.yaml -t
```

Help:

```
A tool for generating parameterized files from templates

Usage:
  tmpl [flags]
  tmpl [command]

Available Commands:
  help        Help about any command
  version     Print version

Flags:
  -d, --delimiter string    template delimiter (default "\n")
  -h, --help                help for tmpl
  -o, --output string       output file path (default "tmpl.out")
  -p, --parameters string   path to parameters
  -s, --source string       path to sources
  -t, --test                print result to stdout

Use "tmpl [command] --help" for more information about a command.
```
