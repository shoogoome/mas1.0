package main

import (
	"github.com/astaxie/beego"
	_ "liuma/routers"
	"liuma/utils"
	"liuma/utils/rs"
)

func init() {
	// 初始化配置
	utils.InitSystemConfig()
	utils.InitEnvConfig()
	// 初始化连接mongodb
	utils.InitMongoClient()
	// 初始化rs配置
	rs.InitRsConfig()
}

func main() {
	beego.Run()
}
