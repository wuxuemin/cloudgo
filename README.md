# cloudgo
### 框架:martini
`go get github.com/go-martini/martini`
### 测试

1.启动服务器
`go run main.go -p9090`

![启动服务器:](https://github.com/wuxuemin/cloudgo/blob/master/image/1.png)

2.另外打开一个终端进行curl测试
`curl -v http://localhost:9090/hello/wu`
##### curl测试结果

![curl测试:](https://github.com/wuxuemin/cloudgo/blob/master/image/4.png)

3.进行ab测试
`  ab -n 1000 -c 100 http://localhost:9090/hello/wu`

##### ab测试结果

![ab测试:](https://github.com/wuxuemin/cloudgo/blob/master/image/5.png)

![ab测试:](https://github.com/wuxuemin/cloudgo/blob/master/image/6.png)

##### ab测试结果说明：
+ Server Software 表示服务器使用的软件
+ Server Hostname 表示服务器主机名
+ Server Port 表示服务器端口
+ Document Path 表示文档路径
+ Document Length 表示文档长度
+ Concurrency Level 表示并发等级，即并发数
+ Time taken for tests 表示测试花费的时间
+ Complete requests 表示完成的请求数
+ Failed requests 表示失败的请求数
+ Total transferred 表示传输的字节数
+ HTML transferred 表示传输的 HTML 报文体的字节数
+ Requests per second 表示平均每秒的请求数
+ Time per request 表示平均每个请求花费的时间
+ Time per request 表示考虑并发，平均每个请求花费的时间
+ Transfer rate 表示平均每秒传输的千字节数
+ Connection Times (ms) 传输时间统计
百分之五十的请求可以在43毫秒内得到响应，百分之九十的请求可以在53毫秒内得到响应，最长的一个请求得到响应的时间为244毫秒。
