package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"liuma/exception/http_err"
	"liuma/models"
)

/**
数据库中获取file信息
 */
func GetFileInfo(hash string) (fileInfo *models.FileInfo) {

	collection := MongoConn.Collection("fileserver")
	record := collection.FindOne(context.Background(),&bson.D{{
		"hash", hash,
	}})

	err := record.Decode(&fileInfo); if err != nil {
		return nil
	}
	return fileInfo
}

/**
写入file信息
 */
func SaveFileInfo(fileInfo models.FileInfo) interface{} {
	// 写入到数据库
	collection := MongoConn.Collection("fileserver")
	_, err := collection.InsertOne(
		context.Background(),
		//&bson.M {
		//	"hash": fileInfo.Hash,
		//	"size": fileInfo.Size,
		//	"server_ip": fileInfo.ServerIp,
		//}
		fileInfo,
	); if err != nil {
		return http_err.SaveFileInfoError(err)
	}
	return nil
}