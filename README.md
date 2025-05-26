# Fyne Demo

This repo is the new home of the Fyne demo application, migrated from `cmd/fyne_demo` in the main repository.

#1） windows安装 goland； 
#2）go env -w CGO_ENABLED=1 # 强制启用 CGO6,11 3）
#3）优化大小

# 步骤 1：编译优化
go build -ldflags="-s -w" main.go

# 步骤 2：strip 处理（Linux/macOS）
strip --strip-unneeded main

# 步骤 3：UPX 压缩
upx --best --lzma main
