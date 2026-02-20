package main

import (
	"fmt"
	"log"
)

func main() {
	services, err := fetchServices()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	formatted, serviceMap := formatServices(services)

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

	idx := 0
	for i, s := range services {
		if s.name == service {
			idx = i
			break
		}
	}
	svc := serviceMap[idx]

	actionSelection := showActionMenu(svc)
	if actionSelection == "" {
		log.Println("No action selected")
		return
	}
	action := extractAction(actionSelection)
	log.Printf("Selected action: %s\n", action)

	executeAction(service, action)
}
