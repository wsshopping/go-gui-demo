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

https://github.com/upx/upx/releases/tag/v5.0.1

-==-----------------
wget https://download.copr.fedorainfracloud.org/results/alonid/mingw-epel7/epel-7-x86_64/01298207-mingw-gcc/mingw64-cpp-4.9.3-1.el7.x86_64.rpm
 1059  yum install  -y  mingw64-cpp-4.9.3-1.el7.x86_64.rpm 
 1060  wget "https://download.copr.fedorainfracloud.org/results/alonid/mingw-epel7/epel-7-x86_64/*"
 1061  wget https://download.copr.fedorainfracloud.org/results/alonid/mingw-epel7/epel-7-x86_64/01298185-mingw-binutils/mingw64-binutils-2.25-1.el7.x86_64.rpm
 1062  yum install  -y  mingw64-binutils-2.25-1.el7.x86_64.rpm 
 1063  wget  "https://download.copr.fedorainfracloud.org/results/alonid/mingw-epel7/epel-7-x86_64/01298184-mingw-filesystem/mingw64-filesystem-110-2.el7.noarch.rpm"
 1064  yum install -y  mingw64-filesystem-110-2.el7.noarch.rpm 
 1065  wget  "https://download.copr.fedorainfracloud.org/results/alonid/mingw-epel7/epel-7-x86_64/01298184-mingw-filesystem/mingw-filesystem-base-110-2.el7.noarch.rpm"
 1066  yum install -y  mingw-filesystem-base-110-2.el7.noarch.rpm 
 1067  yum install -y  mingw64-filesystem-110-2.el7.noarch.rpm 
 1068  yum install  -y  mingw64-binutils-2.25-1.el7.x86_64.rpm 
 1069  yum install  -y  mingw64-gcc-4.9.3-1.el7.x86_64.rpm 
 1070  yum install  -y  mingw64-binutils-2.25-1.el7.x86_64.rpm 
 1071  wget  "https://download.copr.fedorainfracloud.org/results/alonid/mingw-epel7/epel-7-x86_64/01298184-mingw-filesystem/mingw-binutils-generic-2.25-1.el7.x86_64.rpm"
 1072  wget  https://download.copr.fedorainfracloud.org/results/alonid/mingw-epel7/epel-7-x86_64/01298185-mingw-binutils//mingw-binutils-generic-2.25-1.el7.x86_64.rpm
 1073  yum install  -y  mingw-binutils-generic-2.25-1.el7.x86_64.rpm 
 1074  yum install  -y  mingw64-binutils-2.25-1.el7.x86_64.rpm 
 1075  yum install  -y  mingw64-gcc-4.9.3-1.el7.x86_64.rpm 
 1076  wget  https://download.copr.fedorainfracloud.org/results/alonid/mingw-epel7/epel-7-x86_64/01298205-mingw-headers/mingw64-headers-4.0.4-5.el7.noarch.rpm
 1077  wget  https://download.copr.fedorainfracloud.org/results/alonid/mingw-epel7/epel-7-x86_64/01298205-mingw-headers/mingw-headers-4.0.4-5.el7.src.rpm
 1078  yum install  -y  mingw64-headers-4.0.4-5.el7.noarch.rpm 
 1079  wget  https://download.copr.fedorainfracloud.org/results/alonid/mingw-epel7/epel-7-x86_64/01298205-mingw-headers/mingw-headers-4.0.4-5.el7.src.rpm
 1080  ll
 1081  yum install  -y  mingw-headers-4.0.4-5.el7.src.rpm
 1082  yum install  -y  mingw64-headers-4.0.4-5.el7.noarch.rpm 
 1083  wget "https://download.copr.fedorainfracloud.org/results/alonid/mingw-epel7/epel-7-x86_64/01298220-mingw-winpthreads/"
 1084  wget "https://download.copr.fedorainfracloud.org/results/alonid/mingw-epel7/epel-7-x86_64/01298220-mingw-winpthreads/mingw64-winpthreads-4.0.4-1.el7.noarch.rpm"
 1085  yum install  -y  mingw64-winpthreads-4.0.4-1.el7.noarch.rpm 
 1086  wget https://download.copr.fedorainfracloud.org/results/alonid/mingw-epel7/epel-7-x86_64/01298198-mingw-crt/mingw64-crt-4.0.4-3.el7.noarch.rpm
 1087  yum install  -y  mingw64-crt-4.0.4-3.el7.noarch.rpm 
 1088  yum install  -y  mingw64-winpthreads-4.0.4-1.el7.noarch.rpm 
 1089  yum install  -y  mingw64-headers-4.0.4-5.el7.noarch.rpm 
 1090  yum install  -y  mingw64-gcc-4.9.3-1.el7.x86_64.rpm 
 1091  wget  https://download.copr.fedorainfracloud.org/results/alonid/mingw-epel7/epel-7-x86_64/01298207-mingw-gcc/mingw64-cpp-4.9.3-1.el7.x86_64.rpm
 1092  yum install  -y  mingw64-gcc-4.9.3-1.el7.x86_64.rpm 
 1093  yum install  -y  mingw64-cpp-4.9.3-1.el7.x86_64.rpm 
 1094  yum install  -y  mingw64-gcc-4.9.3-1.el7.x86_64.rpm 
 1095  wget  https://download.copr.fedorainfracloud.org/results/alonid/mingw-epel7/epel-7-x86_64/01298207-mingw-gcc/mingw64-gcc-c++-4.9.3-1.el7.x86_64.rpm
 1096  yum install  -y  mingw64-gcc-c++-4.9.3-1.el7.x86_64.rpm 
 1097  wget  https://download.copr.fedorainfracloud.org/results/alonid/mingw-epel7/epel-7-x86_64/01298207-mingw-gcc/mingw64-gcc-4.9.3-1.el7.x86_64.rpm
 1098  yum install  -y  mingw64-gcc-4.9.3-1.el7.x86_64.rpm.1
 1099  ll
 1100  rm -rf    mingw64-gcc-4.9.3-1.el7.x86_64.rpm*
 1101  ll
 1102  wget  https://download.copr.fedorainfracloud.org/results/alonid/mingw-epel7/epel-7-x86_64/01298207-mingw-gcc/mingw64-gcc-4.9.3-1.el7.x86_64.rpm
 1103  yum install  -y  mingw64-gcc-4.9.3-1.el7.x86_64.rpm
 1104  ll
 1105  yum install  -y  mingw64-gcc-c++-4.9.3-1.el7.x86_64.rpm 
 1106  mingw-make
 1107  mingw64-make 
 1108  x86_64-w64-mingw32-gcc --version
 1109  echo '#include <stdio.h>\nint main() { printf("Hello"); return 0; }' > test.c
 1110  x86_64-w64-mingw32-gcc test.c -o test.exe
 1111  file test.exe  
 1112  # 检查编译器版本
 1113  x86_64-w64-mingw32-gcc --version
 1114  # 测试编译（生成 Windows EXE）
 1115  echo '#include <stdio.h>\nint main() { printf("Hello"); return 0; }' > test.c
 1116  x86_64-w64-mingw32-gcc test.c -o test.exe
 1117  file test.exe  # 应显示 "PE32+ executable for MS Windows"
 1118  ll
 1119  cd go
 1120  ll
 1121  cd wd-des/
 1122  ll
 1123  GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 go build -o win64.exe

 
# GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 go build -o win64.exe
生成windows环境

