// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// 打包 gopl-zh 为 zip 文件.
//
// 文件名格式: gopl-zh-20151001-2ae607.zip
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// Git版本号
	gitVersion := getGitCommitVersion()

	// zip文件名
	zipBaseName := fmt.Sprintf("gopl-zh-%s-%s", time.Now().Format("20060102"), gitVersion[:6])

	// 导出Git
	exportGitToZip("./_book/" + "gopl-zh-" + gitVersion + ".zip")

	os.Remove(zipBaseName + ".zip")
	file, err := os.Create(zipBaseName + ".zip")
	if err != nil {
		log.Fatal("os.Create: ", err)
	}
	defer file.Close()

	zipFile := zip.NewWriter(file)
	defer zipFile.Close()

	// create /gopl-zh-20151001-2ae607/
	f, err := zipFile.Create(zipBaseName + "/")
	if err != nil {
		log.Fatal(err)
	}
	if _, err = f.Write([]byte("")); err != nil {
		log.Fatal(err)
	}

	dir := "_book"
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("filepath.Walk: ", err)
		}
		if info.IsDir() {
			return nil
		}
		relpath, err := filepath.Rel(dir, path)
		if err != nil {
			log.Fatal("filepath.Rel: ", err)
		}

		filename := filepath.ToSlash(relpath)
		if isIngoreFile(filename) {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal("ioutil.ReadFile: ", err)
		}

		f, err := zipFile.Create(zipBaseName + "/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		if _, err = f.Write(data); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", filename)
		return nil
	})

	fmt.Printf("Done\n")
}

// 获取Git最新的版本号
//
//	git log -1
//	commit 0460c1b3bb8fbb7e2fc88961e69aa37f4041d6c1
//	Merge: b2d582a e826457
//	Author: chai2010 <chaishushan{AT}gmail.com>
//	Date:   Mon Feb 1 08:04:44 2016 +0800
//
//		Merge pull request #249 from sunclx/patch-3
//
//		fix typo
func getGitCommitVersion() (version string) {
	cmdOut, err := exec.Command(`git`, `log`, `-1`).CombinedOutput()
	if err != nil {
		return "unknown"
	}
	for _, line := range strings.Split(string(cmdOut), "\n") {
		line := strings.TrimSpace(line)
		if strings.HasPrefix(line, "commit") {
			version = strings.TrimSpace(line[len("commit"):])
			return
		}
	}
	return "unknown"
}

// 导出Git到Zip文件
func exportGitToZip(filename string) error {
	if !strings.HasSuffix(filename, ".zip") {
		filename += ".zip"
	}
	if _, err := exec.Command(`git`, `archive`, `--format`, `zip`, `--output`, filename, `master`).CombinedOutput(); err != nil {
		return err
	}
	return nil
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

func isIngoreFile(path string) bool {
	if strings.HasPrefix(path, ".git") {
		return true
	}
	if strings.HasSuffix(path, ".gitignore") {
		return true
	}

	if strings.HasPrefix(path, "vendor") {
		return true
	}
	if strings.HasPrefix(path, "tools") {
		return true
	}
	if strings.HasPrefix(path, "testdata") {
		return true
	}
	if strings.HasSuffix(path, ".go") {
		return true
	}

	return false
}
