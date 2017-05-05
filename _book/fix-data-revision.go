// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// 修复Gitbook生成html的时间戳.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"
)

// 输出目录
const dir = "_book"

var (
	// data-revision="Mon Feb 01 2016 10:18:48 GMT+0800 (中国标准时间)"
	reDataRevision     = regexp.MustCompile(`data\-revision\=\"[^"]+\"`)
	goldenDataRevision = `data-revision="Mon Jan 2 15:04:05 -0700 MST 2006"`
)

func main() {
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

		if !strings.HasSuffix(relpath, ".html") {
			return nil
		}
		if changed := convertFile(path); changed {
			fmt.Printf("%s\n", relpath)
			total++
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

	oldData, err := ioutil.ReadFile(abspath)
	if err != nil {
		log.Fatal("convertFile: ioutil.ReadFile:", err)
	}
	if !utf8.Valid(oldData) {
		return false
	}

	newData := reDataRevision.ReplaceAll(oldData, []byte(goldenDataRevision))
	if string(newData) == string(oldData) {
		return false
	}

	err = ioutil.WriteFile(abspath, newData, 0666)
	if err != nil {
		log.Fatal("convertFile: ioutil.WriteFile:", err)
	}
	return true
}
