package main

import "fmt"

func getIcon(activeState string) string {
	switch activeState {
	case "active":
		return "<span color='#00FF88'>●</span>"
	case "failed":
		return "<span color='#FF5555'>✗</span>"
	default:
		return "<span color='#bbbbbb'>○</span>"
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
		result = append(result, fmt.Sprintf("[%s] %s <span color='#646464'>%s %s %s</span>",
			icon, paddedName, s.activeState, s.subState, s.enabled))

		serviceMap[i] = s
	}
	return result, serviceMap
}
