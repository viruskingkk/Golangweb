方法一：使用二进制文件安装 (推荐)

1、下载二进制文件：

wget https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz
1
2、解压并创建工作目录：

tar -zxf go1.7.3.linux-amd64.tar.gz -C /usr/local/
mkdir /Golang
1
2
3、设置环境变量：

在 /etc/profile 添加：

export GOROOT=/usr/local/go 
export GOBIN=$GOROOT/bin
export GOPKG=$GOROOT/pkg/tool/linux_amd64 
export GOARCH=amd64
export GOOS=linux
export GOPATH=/Golang
export PATH=$PATH:$GOBIN:$GOPKG:$GOPATH/bin
