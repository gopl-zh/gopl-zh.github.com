// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// 更新版本信息
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {
	version := getGitCommitVersion()
	data := makeVersionMarkdown(version)

	err := ioutil.WriteFile("./version.md", []byte(data), 0666)
	if err != nil {
		log.Fatalf("ioutil.WriteFile: err = %v", err)
	}

	fmt.Println("build version", version)
}

// 生成版本文件
func makeVersionMarkdown(version string) string {
	buildTime := time.Now().Format("2006-01-02")

	return fmt.Sprintf(`
<!-- 版本号文件，用于被其它md文件包含 -->

### 版本信息

- 仓库版本：[%s](gopl-zh-%s.zip)
- 构建时间：%s
`,
		version, version,
		buildTime,
	)
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
		return "master"
	}
	for _, line := range strings.Split(string(cmdOut), "\n") {
		line := strings.TrimSpace(line)
		if strings.HasPrefix(line, "commit") {
			version = strings.TrimSpace(line[len("commit"):])
			return
		}
	}
	return "master"
}
