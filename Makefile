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
	gitbook build

zh2tw:
	go run zh2tw.go . .md$$

tw2zh:
	go run zh2tw.go . .md$$ tw2zh

loop:
	go run zh2tw.go . .md$$ tw2zh
	go run zh2tw.go . .md$$ zh2tw

review:
	go run zh2tw.go . .md$$ tw2zh
	gitbook build
	go run zh2tw.go . .md$$ zh2tw

fixlink:
	go run fixlinks.go . .md$$

cover:
	composite  cover_patch.png cover_bgd.png cover.jpg
	convert    -resize 1800x2360! cover.jpg  cover.jpg
	convert    -resize 200x262!   cover.jpg  cover_small.jpg
	convert    -quality 75% cover.jpg        cover.jpg
	convert    -quality 75% cover_small.jpg  cover_small.jpg
	convert    -strip       cover.jpg        cover.jpg
	convert    -strip       cover_small.jpg  cover_small.jpg

