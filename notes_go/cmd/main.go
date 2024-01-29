package main

import (
	"net/http"
	"notes_go/config"
	"notes_go/utils"
)

func main() {
	//连接数据库
	utils.ConnectToDB()

	//请求路径映射
	config.InitHandler()

	//监听80端口
	port := ":8080"
	http.ListenAndServe(port, config.Router)
}
