package main

import (
	"flag"
	"fmt"
	"os"
)

var name = "command_name"
var version = "v0.0.0"

func main() {
	os.Exit(run())
}

const globalUsage = `Usage: %s [global options] <subcommand> [options]

subcommands:
  put     Put something
  get     Get something

global options:
`

type globalOptions struct {
	version bool
	verbose bool
}

func run() int {
	flag.Usage = func() {
		fmt.Printf(globalUsage, name)
		flag.PrintDefaults()
	}
	var gOpts globalOptions
	flag.BoolVar(&gOpts.version, "version", false, "show version and exit")
	flag.BoolVar(&gOpts.verbose, "verbose", false, "enable verbose output")
	flag.Parse()
	if gOpts.version {
		fmt.Println(version)
		return 0
	}

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return 1
	}
	switch args[0] {
	case "put":
		return handlePutCommand(gOpts, args[1:])
	case "get":
		return handleGetCommand(gOpts, args[1:])
	default:
		flag.Usage()
		return 1
	}
	return 0
}

const putCommandUsage = `Usage: %s put [options] target ...

options:
`

type putOptions struct {
	recursive   bool
	updatesTime bool
}

func handlePutCommand(gOpts globalOptions, args []string) int {
	fs := flag.NewFlagSet("put", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Printf(putCommandUsage, name)
		fs.PrintDefaults()
	}
	var opts putOptions
	fs.BoolVar(&opts.recursive, "r", false, "put recursive")
	fs.BoolVar(&opts.updatesTime, "t", false, "updates times")
	fs.Parse(args)
	if len(args) == 0 {
		fs.Usage()
		return 1
	}

	fmt.Println("args:", fs.Args())
	fmt.Printf("put options: %+v\n", opts)
	return 0
}

const getCommandUsage = `Usage: %s get [options] target ...

options:
`

type getOptions struct {
	recursive bool
}

func handleGetCommand(gOpts globalOptions, args []string) int {
	fs := flag.NewFlagSet("get", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Printf(getCommandUsage, name)
		fs.PrintDefaults()
	}
	var opts getOptions
	fs.BoolVar(&opts.recursive, "r", false, "get recursive")
	fs.Parse(args)
	if len(args) == 0 {
		fs.Usage()
		return 1
	}

	fmt.Println("args:", fs.Args())
	fmt.Printf("get options: %+v\n", opts)
	return 0
}
