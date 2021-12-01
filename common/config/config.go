package config

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
	"sparrow/common"
)

func ReadConfig(file string,config interface{}) bool {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(path)

	dir =dir+"\\"+file
	if common.DuildMode=="debug"{
		dir,_=os.Getwd()
		dir =dir+"\\center_server"+"\\"+file
	}
	fileContent,readErr :=ioutil.ReadFile(dir)
	if readErr!=nil{
		panic(readErr)
	}
	if parseErr :=xml.Unmarshal(fileContent,config);parseErr!=nil{
		panic(parseErr)
	}
	return true
}
