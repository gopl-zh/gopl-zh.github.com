// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

//
// Copy dir, support regexp.
//
// Example:
//	cpdir src dst
//	cpdir src dst "\.go$"
//	cpdir src dst "\.tiff?$"
//	cpdir src dst "\.tiff?|jpg|jpeg$"
//
// Help:
//	cpdir -h
//
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

const usage = `
Usage: cpdir src dst [filter]
       cpdir -h

Example:
  cpdir src dst
  cpdir src dst "\.go$"
  cpdir src dst "\.tiff?$"
  cpdir src dst "\.tiff?|jpg|jpeg$"
    
Report bugs to <chaishushan{AT}gmail.com>.
`

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, usage[1:len(usage)-1])
		os.Exit(0)
	}
	filter := ".*"
	if len(os.Args) > 3 {
		filter = os.Args[3]
	}
	total := cpDir(os.Args[2], os.Args[1], filter)
	fmt.Printf("total %d\n", total)
}

func cpDir(dst, src, filter string) (total int) {
	entryList, err := ioutil.ReadDir(src)
	if err != nil && !os.IsExist(err) {
		log.Fatal("cpDir: ", err)
	}
	for _, entry := range entryList {
		if entry.IsDir() {
			cpDir(dst+"/"+entry.Name(), src+"/"+entry.Name(), filter)
		} else {
			mathed, err := regexp.MatchString(filter, entry.Name())
			if err != nil {
				log.Fatal("regexp.MatchString: ", err)
			}
			if mathed {
				srcFname := filepath.Clean(src + "/" + entry.Name())
				dstFname := filepath.Clean(dst + "/" + entry.Name())
				fmt.Printf("copy %s\n", srcFname)

				cpFile(dstFname, srcFname)
				total++
			}
		}
	}
	return
}

func cpFile(dst, src string) {
	err := os.MkdirAll(filepath.Dir(dst), 0666)
	if err != nil && !os.IsExist(err) {
		log.Fatal("cpFile: ", err)
	}
	fsrc, err := os.Open(src)
	if err != nil {
		log.Fatal("cpFile: ", err)
	}
	defer fsrc.Close()

	fdst, err := os.Create(dst)
	if err != nil {
		log.Fatal("cpFile: ", err)
	}
	defer fdst.Close()
	if _, err = io.Copy(fdst, fsrc); err != nil {
		log.Fatal("cpFile: ", err)
	}
}
