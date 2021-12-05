package main

import (
	"encoding/json"
	"flag"
	"fmt"
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
var output = ""

func loadConfig() {
	var (
		ex  string
		err error
	)
	c := flag.String("c", "config.json", "config file name")
	a := flag.Bool("a", true, "absolute file path")
	o := flag.String("o", "special-day", "output file name with (.ics)")
	flag.Parse()
	filename := ""
	if !*a {
		ex, err = os.Executable()
		if err != nil {
			panic(err)
		}
		exPath = filepath.Dir(ex)
		filename = path.Join(exPath, *c)
	} else {
		filename = *c
	}

	if _, err = os.Stat(filename); err != nil {
		panic(err)
	}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, &as)
	output = *o
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
	ioutil.WriteFile(path.Join(exPath, fmt.Sprintf("%s.ics", output)), []byte(res), 0644)
}
