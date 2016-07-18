package main

import "fmt"
import "io/ioutil"
import "log"
import "encoding/json"

func main() {
	fileContents, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	fileString := string(fileContents)

	fmt.Println(fileString)

	var inputJSON map[string]interface{}
	json.Unmarshal(fileContents, &inputJSON)

	printType(inputJSON["dataset_name"])
}
