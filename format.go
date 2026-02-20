package main

import "fmt"

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

func formatServices(services []serviceInfo) ([]string, map[int]serviceInfo) {
	var result []string
	serviceMap := make(map[int]serviceInfo)
	maxNameLen := 0

	for _, s := range services {
		if len(s.name) > maxNameLen {
			maxNameLen = len(s.name)
		}
	}

	for i, s := range services {
		icon := getIcon(s.activeState)
		paddedName := fmt.Sprintf("%-*s", maxNameLen, s.name)
		result = append(result, fmt.Sprintf("[%s] %s %s %s %s",
			icon, paddedName, s.activeState, s.subState, s.enabled))
		serviceMap[i] = s
	}
	return result, serviceMap
}
