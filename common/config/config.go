package config

import (
	"encoding/xml"
	"io/ioutil"
)

func ReadConfig(file string,config interface{}) bool {
	fileContent,readErr :=ioutil.ReadFile(file)
	if readErr!=nil{
		panic(readErr)
	}
	if parseErr :=xml.Unmarshal(fileContent,config);parseErr!=nil{
		panic(parseErr)
	}
	return true
}
