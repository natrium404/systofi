package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type serviceInfo struct {
	name        string
	activeState string
	subState    string
	enabled     string
}

func getIcon(activeState string) string {
	switch activeState {
	case "active":
		return "●"
	case "failed":
		return "✗"
	default:
		return "○"
	}
}

func getActiveState(service string) (string, string) {
	cmd := exec.Command("systemctl", "is-active", service)
	active, _ := cmd.Output()
	activeStr := strings.TrimSpace(string(active))

	cmd = exec.Command("systemctl", "show", service, "--property=SubState", "--value")
	sub, _ := cmd.Output()
	subStr := strings.TrimSpace(string(sub))

	return activeStr, subStr
}

func fetchServices() ([]serviceInfo, error) {
	cmd := exec.Command("systemctl", "list-unit-files", "--type=service", "--no-pager", "--no-legend")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error fetching services: %v", err)
	}

	var services []serviceInfo
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			service := fields[0]
			enabled := fields[1]

			activeState, subState := getActiveState(service)

			services = append(services, serviceInfo{
				name:        service,
				activeState: activeState,
				subState:    subState,
				enabled:     enabled,
			})
		}
	}
	return services, nil
}

func formatServices(services []serviceInfo) []string {
	var result []string
	maxNameLen := 0

	for _, s := range services {
		if len(s.name) > maxNameLen {
			maxNameLen = len(s.name)
		}
	}

	for _, s := range services {
		icon := getIcon(s.activeState)
		paddedName := fmt.Sprintf("%-*s", maxNameLen, s.name)
		result = append(result, fmt.Sprintf("[%s] %s %s %s %s",
			icon, paddedName, s.activeState, s.subState, s.enabled))
	}
	return result
}

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
	}

	rofiCmd := exec.Command("rofi", rofiArgs...)
	rofiCmd.Stdin = strings.NewReader(strings.Join(items, "\n"))

	result, err := rofiCmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(result))
}

var actions = []string{
	"Start",
	"Stop",
	"Restart",
	"Enable",
	"Disable",
	"Status",
}

func showActionMenu() string {
	return runRofi(actions, "Action")
}

func executeAction(service, action string) {
	log.Printf("Executing: %s %s\n", action, service)

	var cmd *exec.Cmd

	switch action {
	case "Status":
		cmd = exec.Command("systemctl", "status", service)
	case "Start":
		cmd = exec.Command("pkexec", "systemctl", "start", service)
	case "Stop":
		cmd = exec.Command("pkexec", "systemctl", "stop", service)
	case "Restart":
		cmd = exec.Command("pkexec", "systemctl", "restart", service)
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
	fields := strings.Fields(selection)
	if len(fields) >= 2 {
		return fields[1]
	}
	return ""
}

func main() {
	services, err := fetchServices()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	formatted := formatServices(services)

	selection := runRofi(formatted, "Services")
	if selection == "" {
		log.Println("No service selected")
		return
	}
	log.Printf("Selection: %s\n", selection)

	service := extractServiceName(selection)
	if service == "" {
		log.Printf("Failed to extract service name from: %s\n", selection)
		return
	}
	log.Printf("Selected service: %s\n", service)

	action := showActionMenu()
	if action == "" {
		log.Println("No action selected")
		return
	}
	log.Printf("Selected action: %s\n", action)

	executeAction(service, action)
}
