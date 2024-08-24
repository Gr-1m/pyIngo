
# Convert to Pythoner designed Go library

> This is a Golang library that allows you to use Python-like syntax in your Golang code.
> It provides a set of data types and built-in functions commonly found in Python, …

## Install:

For Linux: 
```bash
git clone https://github.com/Gr-1m/pyIngo
cd pyIngo
sudo ./install.sh

[+] OK
```

For Window:
```powershell
go env GOROOT

# 查看golang安装位置, 找到该路径下的src目录
# Check the installation location of golang and find the src directory under that path

# 手动将本库目录复制进src或者创建一个软链接
# Manually copy the directory of this library into src or create a soft link
```

For file.go:
```golang
// go get github.com/Gr-1m/pyIngo

import (
        "github.com/Gr-1m/pyIngo"
)
```


## Import In Use:


### net/http

r,err:= Get(url, proxy , headers, verify)
r,err:= Post(url, proxy , headers, body, verify)


### hacktools/brute
