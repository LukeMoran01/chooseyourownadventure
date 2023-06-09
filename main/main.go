package main

import (
	"net/http"

	adventure "github.com/lukemoran01/chooseyourownadventure"
)

func main() {

	parsedJSON := adventure.ParseJSON("adventures/gopher.json")
	handlerCount := len(parsedJSON)
	storyArcHandlers := make([]adventure.AdventureHandler, 0, handlerCount)
	for arc, arcFeatures := range parsedJSON {
		storyArcHandlers = append(storyArcHandlers, adventure.HandlerFromJSON(arc, arcFeatures.(map[string]interface{})))
	}
	for _, arcHandler := range storyArcHandlers {
		http.Handle("/"+arcHandler.Arc, arcHandler)
	}
	http.ListenAndServe(":8080", nil)
}
