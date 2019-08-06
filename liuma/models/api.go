package models

// api结构
type RestfulApi struct {
	Status bool `json:"status"`
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}
