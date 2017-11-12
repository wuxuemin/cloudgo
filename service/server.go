package service

import (
	"github.com/go-martini/martini"
)

// NewServer defined
func NewServer(port string) {
	server := martini.Classic()
	server.Get("/hello/:name", func(params martini.Params) string {
		return "hello " + params["name"]
	})
	//运行程序监听端口
	server.RunOnAddr(":" + port)
}
