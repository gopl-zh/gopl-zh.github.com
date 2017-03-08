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
	version, commitTime := getGitCommitVersion()
	buildTime := time.Now()

	data := makeVersionMarkdown(version, commitTime, buildTime)

	err := ioutil.WriteFile("./version.md", []byte(data), 0666)
	if err != nil {
		log.Fatalf("ioutil.WriteFile: err = %v", err)
	}

	fmt.Println("build version", version)
	fmt.Println(commitTime.Format("2006-01-02 15:04:05"))
	fmt.Println(buildTime.Format("2006-01-02 15:04:05"))
}

// 生成版本文件
func makeVersionMarkdown(version string, commitTime, buildTime time.Time) string {
	return fmt.Sprintf(`
<!-- 版本号文件，用于被其它md文件包含 -->

### 版本信息

- 仓库版本：[%s](gopl-zh-%s.zip)
- 更新时间：%s
- 构建时间：%s
`,
		version, version,
		commitTime.Format("2006-01-02 15:04:05"),
		buildTime.Format("2006-01-02 15:04:05"),
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
func getGitCommitVersion() (version string, date time.Time) {
	cmdOut, err := exec.Command(`git`, `log`, `-1`).CombinedOutput()
	if err != nil {
		return "unknown", time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC) // 第一版发布时间
	}
	for _, line := range strings.Split(string(cmdOut), "\n") {
		line := strings.TrimSpace(line)
		if strings.HasPrefix(line, "commit") {
			version = strings.TrimSpace(line[len("commit"):])
		}
		if strings.HasPrefix(line, "Date") {
			const longForm = "Mon Jan 2 15:04:05 2006 -0700"
			date, _ = time.Parse(longForm, strings.TrimSpace(line[len("Date:"):]))
		}
	}
	if version == "" || date.IsZero() {
		return "unknown", time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC) // 第一版发布时间
	}
	return
}
