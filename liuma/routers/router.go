package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"liuma/controllers"
	"liuma/controllers/file"
)

func init() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		//AllowOrigins: 	  []string{"http://localhost:80", "https://r-share.cn"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"token", "Cookie", "system_token", "systemToken", "Cookies",  "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "cookie", "Cookie", "Cookies", "Set-Cookie", "Set-Cookies", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
    // 文件模块
	beego.Include(&file.FileSystemController{})
    // 系统模块-
    beego.Include(&controllers.ServerControllers{})
}
