package file

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"liuma/exception/http_err"
	"liuma/utils"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

// 保存数据分片数据
// @router /data/shard [post]
func (this *FileSystemController) SaveShard() {

	this.Verification()
	hash := this.LoadHash(utils.All)
	index := this.GetString("index")

	file, _, err := this.GetFile("file"); if err != nil {
		this.Exception(http_err.UploadFail())
	}
	fileName := path.Join(
		utils.SystemConfig.Server.FileRootPath,
		fmt.Sprintf("%s.%s", hash, index),
	)
	fileWrite, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0664); if err != nil {
		this.Exception(http_err.UploadFail())
	}
	defer fileWrite.Close()
	_, err = io.Copy(fileWrite, file); if err != nil {
		this.Exception(http_err.UploadFail())
	}
	this.ReturnJSON(map[string]string {
		"status": "success",
	})
}

// 删除数据分片数据
// @router /data/shard [delete]
func (this *FileSystemController) DeleteShard() {
	this.DeleteLocalFile(utils.SystemConfig.Server.FileRootPath)
}

// 删除分片数据
// @router /data/chuck [delete]
func (this *FileSystemController) DeleteChuck() {
	this.DeleteLocalFile(utils.SystemConfig.Server.FileTempPath)
}

// 获取数据分片数据
// @router /data/shard [get]
func (this *FileSystemController) GetShard() {
	this.GetLocalFile(utils.SystemConfig.Server.FileRootPath)
}

// 获取分片数据
// @router /data/chuck [get]
func (this *FileSystemController) GetChuck() {
	this.GetLocalFile(utils.SystemConfig.Server.FileTempPath)
}

// 获取文件
func (this *FileSystemController) GetLocalFile(pathString string) {
	this.Verification()
	hash := this.LoadHash(utils.All)

	index := this.GetString("index")
	fileName := fmt.Sprintf("%s.%s", hash, index)
	filePath := path.Join(
		pathString,
		fileName,
	)
	fileReader, err := ioutil.ReadFile(filePath); if err != nil {
		this.Exception(http_err.DownloadFail())
	}
	var aa io.ReadSeeker
	aa = bytes.NewReader(fileReader)

	this.Ctx.Output.Header("Content-Disposition", "attachment; filename="+fileName)
	this.Ctx.Output.Header("Content-Length", strconv.FormatFloat(float64(len(fileReader)), 'E', 1, 64))
	http.ServeContent(this.Ctx.Output.Context.ResponseWriter, this.Ctx.Output.Context.Request, fileName, time.Now(), aa)
}

// 删除文件
func (this *FileSystemController) DeleteLocalFile(pathString string) {
	this.Verification()
	hash := this.LoadHash(utils.All)
	index := this.GetString("index")

	fileName := path.Join(
		pathString,
		fmt.Sprintf("%s.%s", hash, index))
	_ = os.Remove(fileName)
	this.ReturnJSON(map[string]string {
		"status": "success",
	})
}




