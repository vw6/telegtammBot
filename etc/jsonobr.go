package main

import (
	"os"
	"encoding/json"
	"fmt"
)


//Создание массива
type Config struct {
	Listbase []Listbase
}
type Listbase struct {
	Name string `json:"name"`
	List1 List1 `json:"list1"`
	List2 List2 `json:"list2"`
	List3 List3 `json:"list3"`
}
type List1 struct {
	List_name string `json:"list_name"`
	List_text string `json:"list_text"`
	List_status string `json:"list_status"`
}
type List2 struct {
	List_name string `json:"list_name"`
	List_text string `json:"list_text"`
	List_status string `json:"list_status"`
}
type List3 struct {
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
	config,_:=loadConfiguration("config.json")
	fmt.Println(config.Listbase[0].Name)
	fmt.Println(config.Listbase[0].List1.List_name)
	fmt.Println(config.Listbase[0].List1.List_text)
	fmt.Println(config.Listbase[0].List1.List_status)
}