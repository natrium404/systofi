package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type serviceInfo struct {
	name        string
	activeState string
	subState    string
	enabled     string
	canStart    bool
	canStop     bool
	canReload   bool
}

func getServiceInfo(service string) serviceInfo {
	cmd := exec.Command("systemctl", "is-active", service)
	active, _ := cmd.Output()
	activeStr := strings.TrimSpace(string(active))

	cmd = exec.Command("systemctl", "show", service, "--property=SubState", "--value")
	sub, _ := cmd.Output()
	subStr := strings.TrimSpace(string(sub))

	cmd = exec.Command("systemctl", "show", service, "--property=CanStart,CanStop,CanReload", "--value")
	props, _ := cmd.Output()
	propsStr := strings.TrimSpace(string(props))
	propsLines := strings.Split(propsStr, "\n")

	canStart := true
	canStop := true
	canReload := false

	for _, line := range propsLines {
		if after, ok := strings.CutPrefix(line, "CanStart="); ok {
			canStart = after == "yes"
		}
		if after, ok := strings.CutPrefix(line, "CanStop="); ok {
			canStop = after == "yes"
		}
		if after, ok := strings.CutPrefix(line, "CanReload="); ok {
			canReload = after == "yes"
		}
	}

	cmd = exec.Command("systemctl", "is-enabled", service)
	enabled, _ := cmd.Output()
	enabledStr := strings.TrimSpace(string(enabled))

	return serviceInfo{
		name:        service,
		activeState: activeStr,
		subState:    subStr,
		enabled:     enabledStr,
		canStart:    canStart,
		canStop:     canStop,
		canReload:   canReload,
	}
}

func fetchServices() ([]serviceInfo, error) {
	cmd := exec.Command("systemctl", "list-unit-files", "--type=service", "--no-pager", "--no-legend")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error fetching services: %v", err)
	}

	var services []serviceInfo
	lines := strings.SplitSeq(string(output), "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			service := fields[0]
			svcInfo := getServiceInfo(service)
			services = append(services, svcInfo)
		}
	}
	return services, nil
}
