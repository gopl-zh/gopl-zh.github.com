# Copyright 2015 Golang-China. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# install gitbook
# npm install gitbook-cli -g

# https://github.com/GitbookIO/gitbook
# https://github.com/wastemobile/gitbook

default:
	gitbook build

zh2tw:
	go run zh2tw.go . .md$$

tw2zh:
	go run zh2tw.go . .md$$ tw2zh
