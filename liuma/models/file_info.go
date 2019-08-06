package models

type FileInfo struct {
	Size int64 `json:"size" bson:"size"`
	Name string `json:"name" bson:"name"`
	Hash string `json:"hash" bson:"hash"`
	CreateTime int64 `json:"create_time" bson:"create_time"`
	ServerIp []string `json:"server_ip" yaml:"server_ip"`
}
