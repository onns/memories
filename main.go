package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	"mem/memories"
)

/*
@Time : 2021/11/27 23:34
@Author : onns
@File : /main.go
*/

var as = make([]*memories.Anniversary, 0)
var exPath = ""

func loadConfig() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath = filepath.Dir(ex)
	filename := path.Join(exPath, "config.json")
	if _, err = os.Stat(filename); err != nil {
		panic(err)
	}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, &as)
}

func init() {
	loadConfig()
	location, _ := time.LoadLocation("UTC")
	for _, a := range as {
		a.Date, _ = time.ParseInLocation("2006-01-02", a.DateRaw, location)
	}
}

func main() {
	res := memories.GenerateIcs(memories.GenerateDays(as))
	ioutil.WriteFile(path.Join(exPath, "special-day.ics"), []byte(res), 0644)
}
