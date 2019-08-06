package utils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"liuma/exception/http_err"
	"os"
	"path"
	"time"
)

// 压缩文件
// 过程出错则丢回原数据和报错信息
func GzipFile(file []byte, name string) ([]byte, interface{}) {

	fileName := fmt.Sprintf("%d", time.Now().Unix()) + ".tar.gz"
	// 构造文件存储路径
	filePath := path.Join(
		SystemConfig.Server.FileTempPath,
		fileName,
	)
	// 创建gzip对象
	d, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
	// 删除临时文件
	defer func() {
		_ = os.Remove(filePath)
	}()
	gw := gzip.NewWriter(d)
	tw := tar.NewWriter(gw)
	header := tar.Header{
		Name: name,
		Size: int64(len(file)),
	}
	// 写入文件header
	err := tw.WriteHeader(&header); if err != nil {
		return file, http_err.GzipFail()
	}
	a := bytes.NewReader(file)
	_, err = io.Copy(tw, a); if err != nil {
		return file, http_err.GzipFail()
	}
	// 这里需提前close流 且顺序不能交换
	// ！不能使用defer
	// 因为要继续读取所以要先关闭才能进行下次读取
	_ = tw.Close()
	_ = gw.Close()
	d.Close()
	// 读取压缩文件流
	gzipFileByte, _ := ioutil.ReadFile(filePath)

	return gzipFileByte, nil
}

// 解压缩文件
// 过程出错则丢回原数据和报错信息
func GunzipFile(file []byte) ([]byte, interface{}) {

	srcFile := bytes.NewReader(file)
	gr, err := gzip.NewReader(srcFile); if err != nil {
		return file, http_err.GunzipFail()
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	// 读取文件
	_, err = tr.Next(); if err != nil {
		return file, http_err.GunzipFail()
	}
	// 转换byte类型
	b, err := ioutil.ReadAll(tr); if err != nil {
		return file, http_err.GunzipFail()
	}
	return b, nil
}
