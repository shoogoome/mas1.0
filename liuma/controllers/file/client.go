package file

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"io"
	"io/ioutil"
	"liuma/exception/http_err"
	"liuma/models"
	"liuma/utils"
	"liuma/utils/rs"
	"net/http"
	"os"
	"path"
	"strconv"
	"sync"
	"time"
)

// 生成上传令牌
// @router /upload/token [get]
func (this *FileSystemController) GenerateUploadToken() {
	this.generateToken(utils.Upload)
}

// 生成下载令牌
// @router /download/token [get]
func (this *FileSystemController) GenerateDownloadToken() {
	this.generateToken(utils.Download)
}

// 单文件上传
// @router /upload/single [post]
func (this *FileSystemController) SingleUpload() {
	// 从token载入hash
	hash := this.LoadHash(utils.Upload)
	// 查看文件是否存在
	if this.SearchFile(hash) {
		this.ReturnJSON(map[string]string{
			"status": "success",
		})
		return
	}
	// 获取file
	file, headers, err := this.GetFile("file"); if err != nil {
		this.Exception(http_err.GetFileFail())
	}
	// 计算真实文件hash
	var dd bytes.Buffer
	reader := io.TeeReader(file, &dd)
	fileHash, except := utils.CalculateHash(reader); if except != nil {
		this.Exception(except)
	}
	// hash不匹配则报token不匹配错误
	if fileHash != hash {
		this.Exception(http_err.TokenFail())
	}
	// 构建文件基础信息
	fileInfo := models.FileInfo{
		Name:       headers.Filename,
		CreateTime: time.Now().Unix(),
		Size:       int64(dd.Len()),
		Hash:       hash,
	}
	// 保存文件
	this.saveFile(dd.Bytes(), fileInfo)
	this.ReturnJSON(map[string]string{
		"status": "success",
	})

}

// 分片上传
// @router /upload/chuck [post]
func (this *FileSystemController) ChunkUpload() {

	hash := this.LoadHash(utils.Upload)
	chuck := this.GetString("chuck")

	redisConn := this.RedisConn()
	defer redisConn.Close()
	chuckInfoString, err := redis.String(redisConn.Do("get", hash))
	// 查询是否已存储
	var chuckInfo models.RedisChucks
	if err == nil {
		err = json.Unmarshal([]byte(chuckInfoString), &chuckInfo); if err != nil {
			this.Exception(http_err.UploadFail())
		}
		_, ok := chuckInfo.ChuckInfo[chuck]; if ok {
			this.Exception(http_err.ChuckExists())
		}
	} else {
		chuckInfo = models.RedisChucks{
			ChuckInfo: map[string]string{},
		}
	}
	file, _, err := this.GetFile("file"); if err != nil {
		this.Exception(http_err.UploadFail())
	}

	fileName := path.Join(
		utils.SystemConfig.Server.FileTempPath,
		fmt.Sprintf("%s.%s", hash, chuck),
	)
	chuckWrite, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0664); if err != nil {
		this.Exception(http_err.UploadFail())
	}
	defer chuckWrite.Close()
	// 存储分片数据
	_, err = io.Copy(chuckWrite, file); if err != nil {
		_ = os.Remove(fileName)
		this.Exception(http_err.UploadFail())
	}
	chuckInfo.ChuckInfo[chuck] = utils.SystemConfig.Server.Server
	chuckInfoByte, _ := json.Marshal(chuckInfo)
	_, _ = redisConn.Do("set", hash, string(chuckInfoByte[:]))
	this.ReturnJSON(map[string]string{
		"status": "success",
	})
}

// 完成上传
// @router /upload/finish [get]
func (this *FileSystemController) Finish() {

	hash := this.LoadHash(utils.Upload)
	token := this.Ctx.Input.Header("token")
	fileName := this.GetString("file_name")
	if fileName == "" {
		this.Exception(http_err.LackParams("file_name"))
	}

	chuckNum, err := this.GetInt("chuck_num"); if err != nil {
		this.Exception(http_err.LackParams("chuck_num"))
	}

	redisConn := this.RedisConn()
	defer redisConn.Close()
	chuckInfoString, err := redis.String(redisConn.Do("get", hash)); if err != nil {
		this.Exception(http_err.UploadFail())
	}
	// 查询是否已存储
	var chuckInfo models.RedisChucks
	if len(chuckInfoString) == 0 {
		this.Exception(http_err.UploadFail())
	}
	err = json.Unmarshal([]byte(chuckInfoString), &chuckInfo); if err != nil {
		this.Exception(http_err.UploadFail())
	}
	// 读取文件
	var mu sync.RWMutex
	lock := make(chan int)
	var chucks = make([][]byte, chuckNum)
	for chuck, ip := range chuckInfo.ChuckInfo {
		go func(c string, nip string, lock chan int) {
			client := http.Client{}
			request, _ := http.NewRequest(
				"GET",
				fmt.Sprintf("http://%s%s?index=%s", nip, utils.SystemConfig.Server.StorageChuckUrl, c),
				nil,
			)
			request.Header.Add("system_token", utils.SystemConfig.Server.Token)
			request.Header.Add("token", token)

			response, err := client.Do(request); if err != nil{
				lock <- 0
			}
			chuckInt, err := strconv.Atoi(c); if err != nil {
				lock <- 0
			}
			dd, err := ioutil.ReadAll(response.Body); if err != nil {
				lock <- 0
			}
			mu.Lock()
			chucks[chuckInt] = dd
			mu.Unlock()
			lock <- 1
		}(chuck, ip, lock)
	}
	// 读取所有分片次数
	for i := 0; i < chuckNum; i++ {
		<- lock
	}
	// 合并所有分块
	fileByte := []byte("")
	allFileByte := bytes.Join(chucks, fileByte)
	// 计算真实文件hash
	var dd bytes.Buffer
	reader := io.TeeReader(bytes.NewBuffer(allFileByte), &dd)
	fileHash, except := utils.CalculateHash(reader)
	if except != nil {
		this.Exception(except)
	}
	// hash不匹配则报token不匹配错误
	if fileHash != hash {
		// 删除分片
		utils.DeleteFileChuck(chuckInfo, token)
		// 清除redis记录
		_, _ = redisConn.Do("del", hash)
		this.Exception(http_err.TokenFail())
	}
	ddByte := dd.Bytes()
	// 构建文件基础信息
	fileInfo := models.FileInfo{
		Name:       fileName,
		CreateTime: time.Now().Unix(),
		Size:       int64(len(ddByte)),
		Hash:       hash,
	}
	// 保存文件
	this.saveFile(ddByte, fileInfo)
	// 删除临时分片
	utils.DeleteFileChuck(chuckInfo, token)
	this.ReturnJSON(map[string]string {
		"status": "success",
	})
}

// 文件下载
// @router /download [get]
func (this *FileSystemController) Download() {

	hash := this.LoadHash(utils.Download)
	token := this.Ctx.Input.Header("token")
	// 查询文件信息
	fileInfo := utils.GetFileInfo(hash)
	if fileInfo == nil {
		this.Exception(http_err.FileIsNotExists())
	}
	shards := make([][]byte, rs.RsConfig.AllShards)
	// 获取分片数据
	var lock = make(chan int)
	for index, ip := range fileInfo.ServerIp {

		go func(ip string, index int, lock chan int) {
			client := http.Client{}
			request, _ := http.NewRequest("GET",
				fmt.Sprintf("http://%s%s?index=%d", ip, utils.SystemConfig.Server.StorageUrl, index),
				nil)
			request.Header.Add("system_token", utils.SystemConfig.Server.Token)
			request.Header.Add("token", token)
			response, err := client.Do(request); if err != nil {
				lock <- 0
				return
			}
			dd, _ := ioutil.ReadAll(response.Body)
			var mu sync.RWMutex
			mu.Lock()
			shards[index] = dd
			mu.Unlock()
			lock <- 1
		}(ip, index, lock)
	}
	// 读取所有分片次数
	for i := 0; i < rs.RsConfig.AllShards; i++ {
		<- lock
	}
	// 获取原文件
	var file io.ReadSeeker
	decode := rs.NewDecoder(shards, fileInfo.ServerIp)
	dd, except := decode.Decode(token); if except != nil {
		utils.DeleteFileShard(fileInfo.ServerIp, token)
		this.Exception(except)
	}
	// gunzip
	if utils.SystemConfig.Server.Gzip {
		dd, except = utils.GunzipFile(dd); if except != nil {
			utils.DeleteFileShard(fileInfo.ServerIp, token)
			this.Exception(except)
		}
	}
	file = bytes.NewReader(dd)
	// 输出文件
	this.Ctx.Output.Header("Content-Disposition", "attachment; filename="+fileInfo.Name)
	//this.Ctx.Output.Header("Content-Length", strconv.FormatFloat(fileInfo.Size, 'E', 1, 64))
	this.Ctx.Output.Header("Content-Length", fmt.Sprintf("%d", len(dd)))
	http.ServeContent(this.Ctx.Output.Context.ResponseWriter, this.Ctx.Output.Context.Request, fileInfo.Name, time.Now(), file)
}

// 查询文件存在情况
func (this *FileSystemController) SearchFile(hash string) bool {
	fileInfo := utils.GetFileInfo(hash)
	if fileInfo != nil {
		return true
	}
	return false
}

// 生成token
func (this *FileSystemController) generateToken(tokenType int) {
	this.Verification()
	hash := this.GetString("hash")
	fileToken := models.FileToken{
		Hash:       hash,
		TokenType:  tokenType,
		CreateTime: time.Now().Unix(),
		ExpireAt:   time.Now().Unix() + 86400,
	}
	token, except := utils.GenerateToken(fileToken)

	if token == "" {
		this.Exception(except)
	}
	this.ReturnJSON(map[string]string{
		"token": token,
	})
}


// 保存文件
func (this *FileSystemController)saveFile(ddbyte []byte, fileInfo models.FileInfo) {
	// gzip压缩
	if utils.SystemConfig.Server.Gzip {
		ddbyte, _ = utils.GzipFile(ddbyte, fileInfo.Name)
	}
	// server
	serverIP := utils.SignalTrigger()
	// 文件服务器数量少于分片总数量则报错
	if len(serverIP) < rs.RsConfig.AllShards {
		this.Exception(http_err.ServerNumLess(len(serverIP)))
	}
	// 文件数据切片
	encode := rs.NewEncoder(ddbyte)
	shards, except := encode.Encode(); if except != nil {
		this.Exception(except)
	}
	// 发送数据分片存储数据
	var statusMap = make(chan models.ShardsStatus, utils.SystemConfig.Server.ServerNum)
	for index, shard := range shards {
		// 填充服务数据
		fileInfo.ServerIp = append(fileInfo.ServerIp, <-serverIP)
		go utils.SendFileShard(fileInfo.ServerIp[len(fileInfo.ServerIp)-1], shard, index, this.Ctx.Input.Header("token"), statusMap)
	}
	// 读取结果 如果有允许损坏分片数量之内的分片数量损坏时
	// 重新修复分片并再次上传

	// 分片计数
	count := rs.RsConfig.AllShards
	// 允许单一分片重发次数
	resend := make([]int, count)
	for {
		// 读取分片传输数据
		var status = <- statusMap
		count -= 1
		if !status.Status {
			// 分片重发
			if resend[status.Index] < utils.SystemConfig.Server.Resend {
				go utils.SendFileShard(<-serverIP,
					shards[status.Index],
					status.Index,
					this.Ctx.Input.Header("token"), statusMap)
				resend[status.Index] += 1
			} else {
				// 重发次数超出设定 认定失败
				// 删除数据分片
				utils.DeleteFileShard(fileInfo.ServerIp, this.Ctx.Input.Header("token"))
				this.Exception(http_err.ResendOver())
			}
		}
		// 所有分片处理完毕
		if count <= 0 {
			break
		}
	}
	// 文件信息存入数据库
	except = utils.SaveFileInfo(fileInfo)
	if except != nil {
		this.Exception(except)
	}
}
