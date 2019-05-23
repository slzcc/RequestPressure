# Request Pressure

通过 Go 编写的 HTTP 的压力请求 Tools。
参数说明:
```-c  client         用户数量
   -t  request number 用户请求的次数
   -u  request url    请求的 URLs
   -o  request status 请求 HTTP 状态码
   -oo request body   请求的响应体（o 与 oo 参数互斥只能用其中一个）
```

## 使用
如果本地拥有 GOPATH 环境可直接编译：
```
$ go build rq.go
$ ./rq -o -c 1500 -t 1 -u https://www.baidu.com
```

如果本地没有 GOPATH 环境可以使用 Docker 进行编译：
```
$ make build
$ ./rq -o -c 1500 -t 1 -u https://www.baidu.com
```

目前支持 GET 方法，如有 POST 方法，请写入 URL 内容中模拟 GET 执行。