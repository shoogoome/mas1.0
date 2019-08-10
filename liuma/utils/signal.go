package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type body struct {
	Status bool `json:"status"`
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data bodyIp `json:"data"`
}

type bodyIp struct {
	Ip string `json:"ip"`
}


// 信号触发
// 信号触发不采用rabbitmq是因为怕同时发起请求的话会串
func SignalTrigger() chan string{

	var serverIps = make(chan string, SystemConfig.Server.ServerNum)
	for _, ip := range SystemConfig.Server.ServerIp {
		go signal(ip, serverIps)
	}
	// 等待响应
	time.Sleep(1 * time.Second)
	return serverIps
}


func signal (nip string, ips chan string) {

	client := http.Client{}
	request, _ := http.NewRequest("GET",
		fmt.Sprintf("http://%s%s", nip, SystemConfig.Server.SignalUrl),
		nil)

	request.Header.Add("systemToken", SystemConfig.Server.Token)
	response, err := client.Do(request); if err != nil {
		return
	}

	var info body
	bb, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(bb, &info); if err != nil {
		return
	}
	ips <- info.Data.Ip
}

