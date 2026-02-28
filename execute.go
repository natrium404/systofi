package main

import (
	"log"
	"os/exec"
	"strings"
)

func executeAction(service, action string) {
	log.Printf("Executing: %s %s\n", action, service)

	var cmd *exec.Cmd

	switch action {
	case "Status":
		cmd = exec.Command("systemctl", "status", service)
	case "Start", "Reload", "Restart":
		cmd = exec.Command("pkexec", "systemctl", strings.ToLower(action), service)
	case "Stop":
		cmd = exec.Command("pkexec", "systemctl", "stop", service)
	case "Enable":
		cmd = exec.Command("pkexec", "systemctl", "enable", service)
	case "Disable":
		cmd = exec.Command("pkexec", "systemctl", "disable", service)
	}

	output, err := cmd.CombinedOutput()
	log.Printf("Output: %s\n", output)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func extractServiceName(selection string) string {
	if idx := strings.Index(selection, ".service"); idx != -1 {
		start := strings.LastIndex(selection[:idx], " ") + 1
		return selection[start : idx+8]
	}
	return ""
}

func extractAction(selection string) string {
	fields := strings.Fields(selection)
	if len(fields) >= 2 {
		return fields[1]
	}
	return ""
}
