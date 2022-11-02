package main

import (
	"fmt"
	"os"
)

type Args struct {
	Subcommand string
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
	} else if args.Subcommand == "cleanup" {
		CleanupSchedule()
	} else {
		ShowTodaysSchedule()
	}
}
