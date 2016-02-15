# Go语言圣经（中文版）

Go语言圣经 [《The Go Programming Language》](http://gopl.io) 中文版本，仅供学习交流之用。

[![](cover_middle.jpg)](http://golang-china.github.io/gopl-zh)

- 在线版本：http://golang-china.github.io/gopl-zh
- 离线版本：http://github.com/golang-china/gopl-zh/archive/gh-pages.zip
- 项目主页：http://github.com/golang-china/gopl-zh
- 原版官网：http://gopl.io


### 从源文件构建

先安装NodeJS和GitBook命令行工具(`npm install gitbook-cli -g`命令)。

1. 运行`go get github.com/golang-china/gopl-zh`，获取 [源文件](https://github.com/golang-china/gopl-zh/archive/master.zip)。
2. 切换到 `gopl-zh` 目录，运行 `gitbook install`，安装GitBook插件。
3. 运行`make`，生成`_book`目录。
4. 打开`_book/index.html`文件。

### 简体/繁体转换

切片到 `gopl-zh` 目录：

- `make zh2tw` 或 `go run zh2tw.go . "\.md$" zh2tw`，转繁体。
- `make tw2zh` 或 `go run zh2tw.go . "\.md$" tw2zh`，转简体。

# 版权声明

<a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/">Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License</a>。

<a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="./images/by-nc-sa-4.0-88x31.png"></img></a>

严禁任何商业行为使用或引用该文档的全部或部分内容！

欢迎大家提供建议！
