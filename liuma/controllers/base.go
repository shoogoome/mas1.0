package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"liuma/exception/http_err"
	"liuma/models"
	"liuma/utils"
	"net/http"
)

// MAS控制器基类
type ControllerBase struct {
	beego.Controller
}

func (this *ControllerBase) Verification() {
	// 验证系统token
	token := this.Ctx.Input.Header("systemToken")

	if utils.SystemConfig.Server.Token != token{
		this.Exception(http_err.SystemTokenVerificationFail())
	}
}

// 从token载入hash
func (this *ControllerBase) LoadHash(tokenType int) string {

	token := this.Ctx.Input.Header("token")
	hash, except := utils.VerificationToken(token, tokenType); if except != nil {
		this.Exception(except)
	}
	return hash
}

// 转换Body中数据
func (this *ControllerBase) requestJSON(ob interface{}) {
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob); if err != nil {
		this.Exception(http_err.UnmarshalBodyError())
	}
}

// 输出
func (this * ControllerBase) ReturnJSON(ob interface{}) {

	this.Ctx.Output.Status = http.StatusOK
	this.Data["json"] = models.RestfulApi{
		Status: true,
		Code: 200,
		Data: ob,
	}
	this.ServeJSON()
}

// 返回JSON格式信息
func (this *ControllerBase) Finish() {
	this.ServeJSON()
}

// 异常处理
func (this *ControllerBase) Exception (error interface{}) {

	this.Ctx.Output.Status = http.StatusInternalServerError
	switch err := error.(type) {
	case http_err.LiumaExceptBase:
		this.Data["json"] = err
	default:
		this.Data["json"] = models.RestfulApi{
			Status: false,
			Code: 5500,
			Msg: "系统异常",
		}
	}
	this.ServeJSON()
	this.StopRun()
}



// redis数据库连接
func (this *ControllerBase) RedisConn() redis.Conn{

	connectString := fmt.Sprintf(
		"%s:%d",
			utils.SystemConfig.Redis.Host,
			utils.SystemConfig.Redis.Port,
		)
	conn, err := redis.Dial("tcp", connectString)
	if err != nil {
		this.Exception(http_err.RedisConnectExcept())
	}
	if _, err := conn.Do("AUTH", utils.SystemConfig.Redis.Password); err != nil {
		_ = conn.Close()
		this.Exception(http_err.RedisVerificationError())
	}
	return conn
}

