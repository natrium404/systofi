package main

import (
	"os/exec"
	"strings"
)

func runRofi(items []string, prompt string) string {
	if len(items) == 0 {
		return ""
	}

	rofiArgs := []string{
		"-dmenu",
		"-p", prompt,
		"-i",
		"-format", "s",
		"-width", "60",
		"-lines", "15",
		"-markup-rows",
		"-theme", "/home/natrium/.config/rofi/main.rasi",
	}

	rofiCmd := exec.Command("rofi", rofiArgs...)
	rofiCmd.Stdin = strings.NewReader(strings.Join(items, "\n"))

	result, err := rofiCmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(result))
}
