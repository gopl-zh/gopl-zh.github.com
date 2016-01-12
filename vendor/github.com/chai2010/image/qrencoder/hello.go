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

func main() {
	c, err := qr.Encode("hello, world", qr.L)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("zz_qrout.png", c.PNG(), 0666)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("output: zz_qrout.png\n")
}
