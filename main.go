package main

import (
	"encoding/json"
	"flag"
	"gopkg.in/Masterminds/sprig.v2"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {

	inputTemplate := flag.String("i", "", "Path to gotemplate")
	outputFile := flag.String("o", "", "Output path")
	jsonFile := flag.String("d", "", "Path to JSON file with data")

	flag.Parse()

	if inputTemplate == nil || *inputTemplate == "" {
		println("You must pass -i argument")
		return
	}
	if outputFile == nil || *outputFile == "" {
		println("You must pass -o argument")
		return
	}
	if jsonFile == nil || *jsonFile == "" {
		println("You must pass -d argument")
		return
	}

	templateBytes, err := ioutil.ReadFile(*inputTemplate)
	if err != nil {
		panic(err)
	}

	jsonBytes, err := ioutil.ReadFile(*jsonFile)
	if err != nil {
		panic(err)
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(jsonBytes, &jsonMap)
	if err != nil {
		panic(err)
	}

	parsedTemplate, err := template.New(*inputTemplate).
		Funcs(sprig.TxtFuncMap()).
		Parse(string(templateBytes))
	if err != nil {
		panic(err)
	}


	interpolated, err := os.OpenFile(*outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}

	var templateData = make(map[string]interface{})
	templateData["data"] = jsonMap

	err = parsedTemplate.Execute(interpolated, templateData)
	if err != nil {
		panic(err)
	}

}
