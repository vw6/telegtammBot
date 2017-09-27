package main

import (
	"os"
	"encoding/json"
	"fmt"
	"io/ioutil"
)


//Создание массива
type Config struct {
	Listbase []Listbase
}
type Listbase struct {
	Name string `json:"name"`
	List_name string `json:"list_name"`
	List_text string `json:"list_text"`
	List_status string `json:"list_status"`
}


func loadConfiguration(filename string)(Config, error)  {
	var config Config
	configFile, err := os.Open(filename)
	defer  configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser:=json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config,err
}
func main() {

	var jsonBlob = []byte(`
        {"Name":"vw666", "List_name":"123", "List_text":"123", "List_status":"321"}
    `)
	rankings := Listbase{}
	err := json.Unmarshal(jsonBlob, &rankings)
	if err != nil {

	}
	rankingsJson, _ := json.Marshal(rankings)
	err = ioutil.WriteFile("vw666.json", rankingsJson, 0644)
	fmt.Printf("%+v", rankings)
}

