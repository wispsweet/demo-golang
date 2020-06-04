package main

import (
	"Question/Route"
	_ "github.com/go-sql-driver/mysql"
)

var Array map[string]interface{}

func main() {

	//路由设置
	engine := Route.InitRoute()
	engine.Run()
}
