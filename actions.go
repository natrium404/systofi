package main

import (
	"fmt"
	"log"
)

type action struct {
	name string
	icon string
}

func getActions(svc serviceInfo) []action {
	var actions []action

	log.Printf("[DEBUG] getActions: name=%s activeState=%q enabled=%q canStart=%v canStop=%v canReload=%v",
		svc.name, svc.activeState, svc.enabled, svc.canStart, svc.canStop, svc.canReload)

	switch svc.activeState {
	case "active":
		if svc.canStop {
			actions = append(actions, action{"Stop", "◼"})
		}
		if svc.canReload {
			actions = append(actions, action{"Reload", "↻"})
		} else if svc.canStop {
			actions = append(actions, action{"Restart", "↻"})
		}
		if svc.enabled == "disabled" {
			actions = append(actions, action{"Enable", "✓"})
		}
		if svc.enabled == "enabled" {
			actions = append(actions, action{"Disable", "✗"})
		}
	case "failed":
		if svc.canStart {
			actions = append(actions, action{"Start", "▶"})
		}
		if svc.canReload {
			actions = append(actions, action{"Reload", "↻"})
		} else if svc.canStart {
			actions = append(actions, action{"Restart", "↻"})
		}
		if svc.enabled == "disabled" {
			actions = append(actions, action{"Enable", "✓"})
		}
		if svc.enabled == "enabled" {
			actions = append(actions, action{"Disable", "✗"})
		}
	default:
		if svc.canStart {
			actions = append(actions, action{"Start", "▶"})
		}
		if svc.enabled == "disabled" {
			actions = append(actions, action{"Enable", "✓"})
		}
		if svc.enabled == "enabled" {
			actions = append(actions, action{"Disable", "✗"})
		}
	}

	actions = append(actions, action{"Status", "ℹ"})

	log.Printf("[DEBUG] Actions for %s: %v", svc.name, actions)

	return actions
}

func showActionMenu(svc serviceInfo) string {
	actions := getActions(svc)
	actionStrings := make([]string, len(actions))
	for i, a := range actions {
		actionStrings[i] = fmt.Sprintf("%s %s", a.icon, a.name)
	}
	return runRofi(actionStrings, "Action")
}
