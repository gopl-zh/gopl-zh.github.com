// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// Lookpath is a simple which.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf(`Usage: lookpath COMMAND [...]
Write the full path of COMMAND(s) to standard output.

Report bugs to <chaishushan@gmail.com>.
`)
		os.Exit(0)
	}
	for i := 1; i < len(os.Args); i++ {
		path, err := exec.LookPath(os.Args[i])
		if err != nil {
			fmt.Printf("lookpath: no %s in (%v)\n", os.Args[i], GetEnv("PATH"))
			os.Exit(0)
		}
		fmt.Println(path)
	}
}

func GetEnv(key string) string {
	key = strings.ToUpper(key) + "="
	for _, env := range os.Environ() {
		if strings.HasPrefix(strings.ToUpper(env), key) {
			return env[len(key):]
		}
	}
	return ""
}
