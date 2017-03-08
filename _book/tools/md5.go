// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

//
// Cacl dir or file MD5Sum, support regexp.
//
// Example:
//	md5 file
//	md5 dir "\.go$"
//	md5 dir "\.tiff?$"
//	md5 dir "\.tiff?|jpg|jpeg$"
//
// Help:
//	cpdir -h
//
package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

const usage = `
Usage: md5 [dir|file [filter]]
       md5 -h

Example:
  md5 file
  md5 dir "\.go$"
  md5 dir "\.go$"
  md5 dir "\.tiff?$"
  md5 dir "\.tiff?|jpg|jpeg$"
    
Report bugs to <chaishushan{AT}gmail.com>.
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage[1:len(usage)-1])
		os.Exit(0)
	}
	filter := ".*"
	if len(os.Args) > 2 {
		filter = os.Args[2]
	}

	if name := os.Args[1]; IsDir(name) {
		m, err := MD5Dir(name, filter)
		if err != nil {
			log.Fatalf("%s: %v", name, err)
		}
		var paths []string
		for path := range m {
			paths = append(paths, path)
		}
		sort.Strings(paths)
		for _, path := range paths {
			fmt.Printf("%x *%s\n", m[path], path)
		}
	} else {
		sum, err := MD5File(name)
		if err != nil {
			log.Fatalf("%s: %v", name, err)
		}
		fmt.Printf("%x *%s\n", sum, name)
	}
}

func IsDir(name string) bool {
	fi, err := os.Lstat(name)
	if err != nil {
		log.Fatal(err)
	}
	return fi.IsDir()
}

func MD5Dir(root string, filter string) (map[string][md5.Size]byte, error) {
	m := make(map[string][md5.Size]byte)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		mathed, err := regexp.MatchString(filter, path)
		if err != nil {
			log.Fatal("regexp.MatchString: ", err)
		}
		if mathed {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			m[path] = md5.Sum(data)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return m, nil
}

func MD5File(filename string) (sum [md5.Size]byte, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	sum = md5.Sum(data)
	return
}
