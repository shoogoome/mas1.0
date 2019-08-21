package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"liuma/exception"
	"liuma/models"
	"os"
	"strconv"
	"strings"
)

var SystemConfig models.SystemConfig

// 初始化读取配置
func InitSystemConfig() {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		exception.OutputException("fail to load config.yaml", err)
	}
	err = yaml.Unmarshal(yamlFile, &SystemConfig)
	if err != nil {
		exception.OutputException("fail to unmarshal config.yaml", err)
	}
}

// 从环境变量读取配置
func InitEnvConfig() {
	// 系统配置
	SystemConfig.Server.FileRootPath = "/root"
	SystemConfig.Server.FileTempPath = "/tmp"
	SystemConfig.Server.SignalUrl = os.Getenv("SignalUrl")
	SystemConfig.Server.StorageUrl = os.Getenv("StorageUrl")
	SystemConfig.Server.StorageChuckUrl = os.Getenv("StorageChuckUrl")
	SystemConfig.Server.Token = os.Getenv("Token")
	SystemConfig.Server.Key = os.Getenv("Key")
	SystemConfig.Server.Resend, _ = strconv.Atoi(os.Getenv("Resend"))
	if os.Getenv("Gzip") == "true" {
		SystemConfig.Server.Gzip = true
	} else {
		SystemConfig.Server.Gzip = false
	}
	serverIP := os.Getenv("ServerIp")
	SystemConfig.Server.ServerIp = strings.Split(serverIP, ",")
	SystemConfig.Server.ServerNum = len(SystemConfig.Server.ServerIp)
	// 获取无头服务名构造dns域名
	headlessServiceName := strings.Split(SystemConfig.Server.ServerIp[0], ".")[1]
	SystemConfig.Server.Server = strings.Join([]string{os.Getenv("HOSTNAME"), headlessServiceName}, ".")
	//SystemConfig.Server.Server = os.Getenv("Server")
}