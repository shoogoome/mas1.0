package models


type RedisChucks struct {
	ChuckInfo map[string]string `json:"chuck_info"`
}

type ShardsStatus struct {
	Ip     string
	Status bool
	Index  int
}
