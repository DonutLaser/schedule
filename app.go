package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/skratchdot/open-golang/open"
)

type Day struct {
	Value string
	Times []string
}

func ReadFile(path string) (result string, success bool) {
	bytes, err := os.ReadFile((path))
	if err != nil {
		color.Red("Error: %s\n", err)
		return "", false
	}

	return string(bytes), true
}

func WriteFile(path string, content string) bool {
	file, err := os.Create(path)
	if err != nil {
		color.Red("Error: %s\n", err)
		return false
	}
	defer file.Close()

	file.WriteString((content))
	file.Sync()

	return true
}

func GetSchedulePath() string {
	exePath, _ := os.Executable()
	exeDir := path.Dir(strings.ReplaceAll(exePath, "\\", "/"))

	return fmt.Sprintf("%s/schedule.txt", exeDir)
}

func ShowSchedule(date string) {
	file, success := ReadFile(GetSchedulePath())
	if !success {
		return
	}

	schedule := make(map[string]string)
	var currentDay string
	var currentSchedule []string

	lines := strings.Split(file, "\n")

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "@") {
			currentDay = strings.TrimLeft(trimmed, "@")
		} else if len(trimmed) == 0 {
			schedule[currentDay] = strings.Join(currentSchedule, "\n")
			currentDay = ""
			currentSchedule = nil
		} else {
			currentSchedule = append(currentSchedule, trimmed)
		}
	}

	if val, ok := schedule[date]; ok {
		color.Green(date)
		fmt.Println(val)
	} else {
		color.Yellow("No schedule for %s\n", date)
	}
}

func CleanupSchedule() {
	schedulePath := GetSchedulePath()
	file, success := ReadFile(schedulePath)
	if !success {
		return
	}

	schedule := make([]Day, 0)
	var currentDay Day

	lines := strings.Split(file, "\n")

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "@") {
			currentDay.Value = strings.TrimLeft(trimmed, "@")
		} else if len(trimmed) == 0 {
			schedule = append(schedule, currentDay)
			currentDay = Day{}
		} else {
			currentDay.Times = append(currentDay.Times, trimmed)
		}
	}

	currentTime := time.Now()
	currentTimeFormatted := currentTime.Format("2006-01-02")

	var sb strings.Builder
	for _, day := range schedule {
		dayTime, err := time.Parse("2006-01-02", day.Value)
		if err != nil {
			color.Red("Error: %s\n", err)
			return
		}

		if currentTimeFormatted == day.Value || dayTime.After(currentTime) {
			sb.WriteString(fmt.Sprintf("@%s\n", day.Value))
			sb.WriteString(strings.Join(day.Times, "\n"))
			sb.WriteString("\n\n")
		}
	}

	_ = WriteFile(schedulePath, sb.String())
}

func EditSchedule() {
	open.Start(GetSchedulePath())
}
