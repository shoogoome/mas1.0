package rs

import (
	"os"
	"strconv"
)

type RSConfig struct {
	// 数据分片数量
	DataShards int
	// 校验分片
	ParityShards int
	// 总分片
	AllShards int
}


// 启动系统时通过环境变量初始化对象（保持动态）
var RsConfig RSConfig


func InitRsConfig() {
	// rs配置
	RsConfig.DataShards, _ = strconv.Atoi(os.Getenv("DataShards"))
	RsConfig.ParityShards, _ = strconv.Atoi(os.Getenv("ParityShards"))
	RsConfig.AllShards = RsConfig.DataShards + RsConfig.ParityShards
}