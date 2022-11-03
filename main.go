package main

import (
	"fmt"
	"os"
	"time"
)

type Args struct {
	Subcommand string
	Param      string
}

func printUsage() {
	fmt.Println("Usage: schedule [subcommand]")
	fmt.Println("Subcommands:")
	fmt.Println("	edit		Edit the schedule")
	fmt.Println("	cleanup		Remove schedule for days that have already passed")
}

func parseArgs() (result Args, success bool) {
	args := os.Args[1:]

	if len(args) < 1 {
		return Args{Subcommand: ""}, true
	}

	result = Args{
		Subcommand: args[0],
		Param:      "",
	}

	if len(args) == 2 {
		result.Param = args[1]
	}

	return result, true
}

func main() {
	args, success := parseArgs()

	if !success {
		printUsage()
		return
	}

	if args.Subcommand == "edit" {
		EditSchedule()
	} else if args.Subcommand == "on" {
		if args.Param == "" {
			fmt.Println("Error: date is not specified")
			printUsage()
			return
		}

		ShowSchedule(args.Param)
	} else if args.Subcommand == "cleanup" {
		CleanupSchedule()
	} else {
		ShowSchedule(time.Now().Format("2006-01-02"))
	}
}
