// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"fmt"
	"io/ioutil"
	"log"

	qr "github.com/chai2010/image/qrencoder"
)

const (
	gopl_zh_url = "https://github.com/golang-china/gopl-zh"
	output      = "gopl-zh-qrcode.png"
)

func main() {
	c, err := qr.Encode(gopl_zh_url, qr.H)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(output, c.PNG(), 0666)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("output:", output)
}
