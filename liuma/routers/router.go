package routers

import (
	"github.com/astaxie/beego"
	"liuma/controllers"
	"liuma/controllers/file"
)

func init() {

    // 文件模块
	beego.Include(&file.FileSystemController{})
    // 系统模块-
    beego.Include(&controllers.ServerControllers{})


}
