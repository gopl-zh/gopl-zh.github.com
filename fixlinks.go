// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

//
// 给特殊模式标定的单词增加链接.
//
// Example:
//	fixlinks
//	fixlinks dir "\.go$"
//
// Help:
//	fixlinks -h
//
package main

import (
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
Usage: fixlinks dir [nameFilter]
       fixlinks -h

Example:
  fixlinks
  fixlinks dir "\.go$"
    
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

	newData := append([]byte{}, oldData...)
	for re, v := range _RegexpLinksTable {
		newData = re.ReplaceAll(newData, []byte(v))
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

var _RegexpLinksTable = func() map[*regexp.Regexp]string {
	const (
		reHttp     = `(https?://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|])?`
		reWikiTail = `(\([-A-Za-z0-9+~_]+\))?`
	)

	m := make(map[*regexp.Regexp]string)
	for k, v := range _LinkTable {
		reKey := regexp.MustCompile(regexp.QuoteMeta(`[`+k+`]`) + `\(` + reHttp + reWikiTail + `\)`)
		m[reKey] = fmt.Sprintf("[%s](%s)", k, v)
	}
	return m
}()

var _LinkTable = map[string]string{

	// 人名
	"Alan Donovan":       "https://github.com/adonovan",
	"Brian Kernighan":    "http://www.cs.princeton.edu/~bwk/",
	"Alan A. A. Donovan": "https://github.com/adonovan",
	"Brian W. Kernighan": "http://www.cs.princeton.edu/~bwk/",
	"Dennis M. Ritchie":  "http://genius.cat-v.org/dennis-ritchie/",
	"Robert Griesemer":   "http://research.google.com/pubs/author96.html",
	"Rob Pike":           "http://genius.cat-v.org/rob-pike/",
	"Ken Thompson":       "http://genius.cat-v.org/ken-thompson/",
	"Russ Cox":           "http://research.swtch.com/",
	"Niklaus Wirth":      "https://en.wikipedia.org/wiki/Niklaus_Wirth",
	"Tony Hoare":         "https://en.wikipedia.org/wiki/Tony_Hoare",
	"Fred Brooks":        "http://www.cs.unc.edu/~brooks/",

	// 图书
	"The C Programming Language":  "http://s3-us-west-2.amazonaws.com/belllabs-microsite-dritchie/cbook/index.html",
	"The Practice of Programming": "https://en.wikipedia.org/wiki/The_Practice_of_Programming",

	// Go语言
	"Go":              "https://golang.org/",
	"Google’s Go":     "https://golang.org/",
	"oracle":          "https://godoc.org/golang.org/x/tools/oracle",
	"godoc -analysis": "https://godoc.org/golang.org/x/tools/cmd/godoc",
	"gorename":        "https://godoc.org/golang.org/x/tools/cmd/gorename",

	// 其他语言
	"Alef":      "http://doc.cat-v.org/plan_9/2nd_edition/papers/alef/",
	"APL":       "https://en.wikipedia.org/wiki/APL_(programming_language)",
	"Limbo":     "http://doc.cat-v.org/inferno/4th_edition/limbo_language/",
	"Modula-2":  "https://en.wikipedia.org/wiki/Modula-2",
	"Newsqueak": "http://doc.cat-v.org/bell_labs/squeak/",
	"Oberon":    "https://en.wikipedia.org/wiki/Oberon_(programming_language)",
	"Oberon-2":  "https://en.wikipedia.org/wiki/Oberon-2_(programming_language)",
	"Pascal":    "https://en.wikipedia.org/wiki/Pascal_(programming_language)",
	"Scheme":    "https://en.wikipedia.org/wiki/Scheme_(programming_language)",
	"Squeak":    "http://doc.cat-v.org/bell_labs/squeak/",

	// 系统
	"Unix":              "http://doc.cat-v.org/unix/",
	"UNIX":              "http://doc.cat-v.org/unix/",
	"Linux":             "http://www.linux.org/",
	"FreeBSD":           "https://www.freebsd.org/",
	"OpenBSD":           "http://www.openbsd.org/",
	"Mac OSX":           "http://www.apple.com/cn/osx/",
	"Mac OS X":          "http://www.apple.com/cn/osx/",
	"Plan9":             "http://plan9.bell-labs.com/plan9/",
	"Microsoft Windows": "https://www.microsoft.com/zh-cn/windows/",

	// 其他
	"Bell Labs":                          "http://www.cs.bell-labs.com/",
	"communicating sequential processes": "https://en.wikipedia.org/wiki/Communicating_sequential_processes",
	"CSP": "https://en.wikipedia.org/wiki/Communicating_sequential_processes",
	"K&R": "https://en.wikipedia.org/wiki/K%26R",
}
