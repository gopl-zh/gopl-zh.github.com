// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

//
// 给特殊模式标定的单词增加链接.
//
// Example:
//	addlinks
//	addlinks dir "\.go$"
//
// Help:
//	addlinks -h
//
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"unicode/utf8"
)

const MaxFileSize = 8 << 20 // 8MB

const usage = `
Usage: addlinks dir [nameFilter]
       addlinks -h

Example:
  addlinks
  addlinks dir "\.go$"
    
Report bugs to <chaishushan{AT}gmail.com>.
`

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" {
		fmt.Fprintln(os.Stderr, usage[1:len(usage)-1])
		os.Exit(0)
	}
	dir, nameFilter := os.Args[1], ".*"
	if len(os.Args) > 2 {
		nameFilter = os.Args[2]
	}

	total := 0
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("filepath.Walk: ", err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		relpath, err := filepath.Rel(dir, path)
		if err != nil {
			log.Fatal("filepath.Rel: ", err)
			return err
		}
		mathed, err := regexp.MatchString(nameFilter, relpath)
		if err != nil {
			log.Fatal("regexp.MatchString: ", err)
		}
		if mathed {
			if changed := convertFile(path); changed {
				fmt.Printf("%s\n", relpath)
				total++
			}
		}
		return nil
	})
	fmt.Printf("total %d\n", total)
}

func convertFile(path string) (changed bool) {
	abspath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal("convertFile: filepath.Abs:", err)
	}

	fi, err := os.Lstat(abspath)
	if err != nil {
		log.Fatal("convertFile: os.Lstat:", err)
	}
	if fi.Size() > MaxFileSize {
		return false
	}

	oldData, err := ioutil.ReadFile(abspath)
	if err != nil {
		log.Fatal("convertFile: ioutil.ReadFile:", err)
	}
	if !utf8.Valid(oldData) {
		return false
	}

	newData := oldData
	for k, v := range _LinksTable {
		newData = bytes.Replace(oldData, []byte(k), []byte(v), -1)
	}

	if string(newData) == string(oldData) {
		return false
	}

	err = ioutil.WriteFile(abspath, newData, 0666)
	if err != nil {
		log.Fatal("convertFile: ioutil.WriteFile:", err)
	}
	return true
}

var _LinksTable = map[string]string{
	"[Alan A. A. Donovan]()": "[Alan A. A. Donovan](https://github.com/adonovan)",
}
