// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

//
// List files, support file/header regexp.
//
// Example:
//	lsdir dir
//	lsdir dir "\.go$"
//	lsdir dir "\.go$" "chaishushan"
//	lsdir dir "\.tiff?|jpg|jpeg$"
//
// Help:
//	lsdir -h
//
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

const usage = `
Usage: lsdir dir [nameFilter [dataFilter]]
       lsdir -h

Example:
  lsdir dir
  lsdir dir "\.go$"
  lsdir dir "\.go$" "chaishushan"
  lsdir dir "\.tiff?|jpg|jpeg$"
    
Report bugs to <chaishushan{AT}gmail.com>.
`

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" {
		fmt.Fprintln(os.Stderr, usage[1:len(usage)-1])
		os.Exit(0)
	}
	dir, nameFilter, dataFilter := os.Args[1], ".*", ""
	if len(os.Args) > 2 {
		nameFilter = os.Args[2]
	}
	if len(os.Args) > 3 {
		dataFilter = os.Args[3]
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
			if dataFilter != "" {
				data, err := ioutil.ReadFile(path)
				if err != nil {
					fmt.Printf("ioutil.ReadFile: %s\n", path)
					log.Fatal("ioutil.ReadFile: ", err)
				}
				mathed, err := regexp.MatchString(dataFilter, string(data))
				if err != nil {
					log.Fatal("regexp.MatchString: ", err)
				}
				if mathed {
					fmt.Printf("%s\n", relpath)
					total++
				}
			} else {
				fmt.Printf("%s\n", relpath)
				total++
			}
		}
		return nil
	})
	fmt.Printf("total %d\n", total)
}
