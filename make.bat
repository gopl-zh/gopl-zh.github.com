@echo off
cls
setlocal EnableDelayedExpansion
rem gitbook install
rem &^ 批处理运行gitbook会中断命令 所以用&链接,用^处理换行
go run update_version.go
gitbook build &^
go run fix-data-revision.go &^
go run builder.go