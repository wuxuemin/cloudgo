# cloudgo
### 框架:martini
`go get github.com/go-martini/martini`
### 测试

1.启动服务器
`go run main.go -p9090`

2.另外打开一个终端进行curl测试
`curl -v http://localhost:9090/hello/wu`
##### curl测试结果

3.进行ab测试
`  ab -n 1000 -c 100 http://localhost:9090/hello/wu`

##### ab测试结果
##### ab测试结果中的参数：
+ Server Software: 服务器使用的软件，即响应头的 Server 字段
+ Server Hostname: 服务器主机名，即请求头的 Host 字段
+ Server Port: 服务器端口
+ Document Path: 文档路径，即请求头的请求路径
+ Document Length: 文档长度，即响应头的 Content-Length 字段
+ Concurrency Level: 并发等级，即并发数
+ Time taken for tests: 测试花费的时间
+ Complete requests: 完成的请求数
+ Failed requests: 失败的请求数
+ Total transferred: 传输的字节数
+ HTML transferred: 传输的 HTML 报文体的字节数
+ Requests per second: 平均每秒的请求数
+ Time per request: 平均每个请求花费的时间
+ Time per request: 考虑并发，平均每个请求花费的时间
+ Transfer rate: 平均每秒传输的千字节数
+ Connection Times (ms) 传输时间统计
