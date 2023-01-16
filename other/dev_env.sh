#!/bin/sh

# 开发环境搭建


# 1. 下载安装包（经过 gzip 压缩的 tar 文件）
wget https://studygolang.com/dl/golang/go1.17.12.linux-amd64.tar.gz


# 2. 将安装包解压至目标路径
sudo tar -C /usr/local -xzf go1.17.12.linux-amd64.tar.gz 


# 3. 设置 $PATH 系统环境变量（将 '/usr/local/go/bin' 添加到 $PATH 环境变量）
vim ~/.profile
    
    ```
    export PATH=$PATH:/usr/local/go/bin
    
    ```
source ~/.profile


# 4. 设置 go env 

go env -w GO111MODULE="on"
go env -w GOPROXY="https://mirrors.aliyun.com/goproxy/,https://mirrors.cloud.tencent.com/go/,direct"
go env -w GOPATH="/home/xsj/go"

vim ~/.profile
    
    ```   
    # 开启 Go Modules 特性
    export GO111MODULE="on"     
    # go get 安装模块，代理服务器设置
    export GOPROXY="https://mirrors.aliyun.com/goproxy/,https://mirrors.cloud.tencent.com/go/,direct"
    # 设置 GOPATH
    export GOPATH=/home/xsj/go

    # 把通过 go install 安装的二进制文件加入到 $PATH 环境变量
    export PATH=$GOPATH/bin:$PATH

    ```
source ~/.profile


# 5. 验证
go version

