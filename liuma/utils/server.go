package utils

import (
	"fmt"
	"liuma/models"
	"net/http"
)

// 远程发送数据分片
func SendFileShard(ip string, shard []byte, index int, token string, statusMap chan models.ShardsStatus) {

	buf, contentType, err := CreateFileBuffer(shard); if err != nil {
		statusMap <- models.ShardsStatus{
			Ip:     ip,
			Status: false,
			Index:  index,
		}
		return
	}

	client := http.Client{}
	request, err := http.NewRequest("POST",
		fmt.Sprintf("http://%s%s?index=%d", ip, SystemConfig.Server.StorageUrl, index),
		&buf)
	if err != nil {
		statusMap <- models.ShardsStatus{
			Ip:     ip,
			Status: false,
			Index:  index,
		}
		return
	}
	request.Header.Add("systemToken", SystemConfig.Server.Token)
	request.Header.Add("token", token)
	request.Header.Set("Content-Type", contentType)
	response, err := client.Do(request)
	if err != nil || response.StatusCode != http.StatusOK {
		statusMap <- models.ShardsStatus{
			Ip:     ip,
			Status: false,
			Index:  index,
		}
		return
	}
	statusMap <- models.ShardsStatus{
		Ip:     ip,
		Status: true,
		Index:  index,
	}
	return
}

// 发送删除数据分片
func DeleteFileShard(ips []string, token string) {

	// TODO: 后续添加日志 该接口的报错信息将放在日志

	for index, ip := range ips {
		go func(i string, ind int) {
			client := http.Client{}
			request, _ := http.NewRequest("DELETE",
				fmt.Sprintf("http://%s%s?index=%d", i, SystemConfig.Server.StorageUrl, ind),
				nil)
			request.Header.Add("systemToken", SystemConfig.Server.Token)
			request.Header.Add("token", token)
			_, _ = client.Do(request)
		}(ip, index)
	}
}

// 发送删除数据分片
func DeleteFileChuck(chuckInfo models.RedisChucks, token string) {

	// TODO: 后续添加日志 该接口的报错信息将放在日志

	for index, ip := range chuckInfo.ChuckInfo {
		go func(i string, ind string) {
			client := http.Client{}
			request, _ := http.NewRequest("DELETE",
				fmt.Sprintf("http://%s%s?index=%s", i, SystemConfig.Server.StorageChuckUrl, ind),
				nil)
			request.Header.Add("systemToken", SystemConfig.Server.Token)
			request.Header.Add("token", token)
			_, _ = client.Do(request)
		}(ip, index)
	}
}
