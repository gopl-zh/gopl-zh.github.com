# Copyright 2015 Golang-China. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# install gitbook
# npm install gitbook-cli -g

# https://github.com/GitbookIO
# https://github.com/GitbookIO/gitbook
# https://github.com/GitbookIO/plugin-katex
# https://github.com/wastemobile/gitbook

# http://www.imagemagick.org/

default:
	go run update_version.go
	gitbook build
	go run fix-data-revision.go
	go run builder.go

zh2tw:
	go run zh2tw.go . .md$$

tw2zh:
	go run zh2tw.go . .md$$ tw2zh

loop:
	go run zh2tw.go . .md$$ tw2zh
	go run zh2tw.go . .md$$ zh2tw

cover:
	composite  cover_patch.png cover_bgd.png cover.jpg
	convert    -resize 1800x2360! cover.jpg  cover.jpg
	convert    -resize 200x262!   cover.jpg  cover_small.jpg
	convert    -resize 400x524!   cover.jpg  cover_middle.jpg
	convert    -quality 75% cover.jpg        cover.jpg
	convert    -quality 75% cover_small.jpg  cover_small.jpg
	convert    -quality 75% cover_middle.jpg cover_middle.jpg
	convert    -strip       cover.jpg        cover.jpg
	convert    -strip       cover_small.jpg  cover_small.jpg
	convert    -strip       cover_middle.jpg cover_middle.jpg
