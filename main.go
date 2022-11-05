package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

type Args struct {
	Subcommand string
	Param      string
}

func printUsage() {
	fmt.Println("Usage: schedule [subcommand]")
	fmt.Println("Subcommands:")
	fmt.Println("	edit				Edit the schedule")
	fmt.Println("	tomorrow			Show schedule for tomorrow")
	fmt.Println("	for <YYYY-MM-DD>	Show schedule for a specific date")
	fmt.Println("	cleanup				Remove schedule for days that have already passed")
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
	} else if args.Subcommand == "for" {
		if args.Param == "" {
			color.Red("Error: date is not specified\n\n")
			printUsage()
			return
		}

		ShowSchedule(args.Param)
	} else if args.Subcommand == "cleanup" {
		CleanupSchedule()
	} else if args.Subcommand == "tomorrow" {
		ShowSchedule(time.Now().AddDate(0, 0, 1).Format("2006-01-02"))
	} else {
		ShowSchedule(time.Now().Format("2006-01-02"))
	}
}
