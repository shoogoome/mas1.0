package utils

import (
	"bytes"
	"mime/multipart"
)

// 构建form-data buffer流
func CreateFileBuffer(fileByte []byte) (bytes.Buffer, string, error){

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	fw, err := w.CreateFormFile("file", "file"); if err != nil {
		return buf, "", err
	}

	_, err = fw.Write(fileByte)
	if err != nil {
		return buf, "", err
	}
	_ = w.Close()
	return buf, w.FormDataContentType(), nil
}

