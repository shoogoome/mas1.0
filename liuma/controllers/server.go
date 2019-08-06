package controllers

import (
	"fmt"
	"liuma/exception/http_err"
	"liuma/utils"
	"sync"
)

type ServerControllers struct {
	ControllerBase
}

type Token struct {
	token string `json:"token" yaml:"token"`
}

// 修改系统token
// @router /server/token [put]
func (this *ServerControllers) ChangeToken() {
	this.Verification()
	var params Token
	this.requestJSON(params)
	// 过滤格式
	if 20 > len(params.token) || len(params.token) > 200 {
		this.Exception(http_err.LengthIsNotAllow("token", 20, 200))
	}
	var mu sync.RWMutex
	mu.Lock()
	defer mu.Unlock()
	utils.SystemConfig.Server.Token = params.token
	this.ReturnJSON(map[string]string {
		"status": "success",
	})
}


// 活跃信号反馈
// @router /server/signal [get]
func (this *ServerControllers) Signal() {
	this.Verification()
	// 返回当前服务ip
	this.ReturnJSON(map[string]string {
		"ip": utils.SystemConfig.Server.Server,
	})
}

// 活跃信号反馈
// @router /server/active [get]
func (this *ServerControllers) Active() {
	this.Verification()
	ips := utils.SignalTrigger()

	// 返回当前服务ip
	this.ReturnJSON(map[string]string{
		"status": fmt.Sprintf("当前活跃服务数量: %d", len(ips)),
	})
}

// 阿里云ssl验证
// @router /.well-known/pki-validation/fileauth.txt [get]
func (this *ServerControllers) SSL() {
	this.Ctx.WriteString("201908050000002m8nba7ne2mgf4nf5baugpivczr7yol7zgqd9sjvofq9gr3zm0")
}
