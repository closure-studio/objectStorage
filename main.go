package main

import (
	"github.com/closure-studio/objectStorage/server"
	"github.com/closure-studio/objectStorage/utils/env"
	"github.com/closure-studio/objectStorage/utils/storage"
)

func main() {

	// 初始化环境变量
	env.InstanceInit()

	storage.S3ClientInit()
	// 手动创建 Server 实例
	srv := server.NewServer()

	// 启动服务器
	srv.Start()
}
