package adventure

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type AdventureHandler struct {
	Arc     string
	Title   string
	Story   []string
	Options []map[string]string
}

func (handler AdventureHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	storyTemplate, err := template.ParseFiles("templates/template.html")
	if err != nil {
		log.Fatal(err)
	}
	storyTemplate.Execute(writer, handler)
}

func ParseJSON(JSONFilename string) map[string]interface{} {
	var JSONHolder map[string]interface{}
	JSONFile, err := os.ReadFile(JSONFilename)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(JSONFile, &JSONHolder); err != nil {
		log.Fatal(err)
	}
	return JSONHolder
}

func convertInterfaceToStringSlice(interfaceToConvert []interface{}) []string {
	interfaceHolder := make([]string, 0, len(interfaceToConvert))
	for _, element := range interfaceToConvert {
		interfaceHolder = append(interfaceHolder, fmt.Sprintf("%v", element))
	}
	return interfaceHolder
}

func convertInterfaceToStringMap(interfaceToConvert []interface{}) []map[string]string {
	interfaceHolder := make([]map[string]string, len(interfaceToConvert))
	for index, element := range interfaceToConvert {
		interfaceHolder[index] = make(map[string]string)
		for key, value := range element.(map[string]interface{}) {
			interfaceHolder[index][key] = value.(string)
		}
	}
	return interfaceHolder
}

func HandlerFromJSON(arc string, json map[string]interface{}) AdventureHandler {
	newHandler := AdventureHandler{}
	newHandler.Arc = arc
	newHandler.Title = json["title"].(string)
	newHandler.Story = convertInterfaceToStringSlice(json["story"].([]interface{}))
	newHandler.Options = convertInterfaceToStringMap(json["options"].([]interface{}))
	return newHandler
}
