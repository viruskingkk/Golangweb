<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
</head>

<body>
<p>下载&解壓縮: </p>
<table border="0" cellspacing="0" cellpadding="0" width="0">
  <tr>
   <td><pre>
      1、下载二进制文件：: <br />
      wget https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz <br />
      2、解压并创建工作目录： <br />
      tar -zxf go1.7.3.linux-amd64.tar.gz -C /usr/local/ <br />
      </pre></td>
  </tr>
</table>

<div>
  <h3>设置环境变量： </h3>
</div>
<table border="0" cellspacing="0" cellpadding="0" width="0">
  <tr>
    <td><pre>在 /etc/profile 添加：
          export GOROOT=/usr/local/go <br />
          export GOBIN=$GOROOT/bin   <br />
          export GOPKG=$GOROOT/pkg/tool/linux_amd64   <br />
          export GOARCH=amd64   <br />
          export GOOS=linux   <br />
          export GOPATH=/Golang  <br />
          export PATH=$PATH:$GOBIN:$GOPKG:$GOPATH/bin <br />
      </pre></td>
  </tr>
</table>

<p>使profile配置立即生效： </p>
<table border="0" cellspacing="0" cellpadding="0" width="0">
  <tr>
   <td><pre>
      执行 source /etc/profile 使之生效或者重新登录Linux也可。
      </pre></td>
  </tr>
</table>   
