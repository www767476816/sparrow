package config

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

func ReadConfig(file string,config interface{}) bool {
	path,pathErr :=os.Getwd()
	if pathErr!=nil{
		panic(pathErr)
	}
	path =path+"\\"+file
	fileContent,readErr :=ioutil.ReadFile(path)
	if readErr!=nil{
		panic(readErr)
	}
	if parseErr :=xml.Unmarshal(fileContent,config);parseErr!=nil{
		panic(parseErr)
	}
	return true
}
