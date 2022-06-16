package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
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
		ex       string
		err      error
		fileInfo os.FileInfo
		files    []fs.FileInfo
		bs       []byte
	)
	c := flag.String("c", "config", "config file name or config folder")
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

	if fileInfo, err = os.Stat(filename); err != nil {
		panic(err)
	}
	if fileInfo.IsDir() {
		files, err = ioutil.ReadDir(filename)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			tas := make([]*memories.Anniversary, 0)
			tempFilename := path.Join(filename, f.Name())
			bs, err = ioutil.ReadFile(tempFilename)
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(bs, &tas)
			if err != nil {
				panic(err)
			}
			as = append(as, tas...)
		}

	} else {
		bs, err = ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bs, &as)
		if err != nil {
			panic(err)
		}
	}
	output = *o
}

func init() {
	loadConfig()
	allDayLocation, _ := time.LoadLocation("UTC")
	location := time.Local
	for _, a := range as {
		a.AllDay = true
		a.Date, _ = time.ParseInLocation("2006-01-02", a.DateRaw, allDayLocation)
		if a.StartRaw != "" {
			a.Start, _ = time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s %s", a.DateRaw, a.StartRaw), location)
			a.AllDay = false
		}
		if a.EndRaw != "" {
			a.End, _ = time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s %s", a.DateRaw, a.EndRaw), location)
			a.AllDay = false
		}
	}
}

func main() {
	res := memories.GenerateIcs(output, memories.GenerateDays(as))
	ioutil.WriteFile(path.Join(exPath, fmt.Sprintf("%s.ics", output)), []byte(res), 0644)
}
