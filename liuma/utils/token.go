package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"liuma/exception/http_err"
	"liuma/models"
	"strings"
	"time"
)

const (

	Upload = 1
	Download = 2
	All = 3
)

/**
生成token
 */
func GenerateToken(info models.FileToken) (string,  interface{}){

	key := SystemConfig.Server.Key
	key = strings.Replace(key, " ", "", -1); if key == "" {
		return "", http_err.GetEnvKeyFail()
	}
	infoString, err := json.Marshal(info); if err != nil {
		return "", http_err.MarshalFail()
	}
	payloadString := base64.StdEncoding.EncodeToString(infoString)
	return fmt.Sprintf("%s.%s", payloadString, encryptionString(payloadString, key)), nil
}

/**
验证token
 */
func VerificationToken(token string, tokenType int) (string, interface{}) {

	var fileToken models.FileToken
	// 环境变量异常
	key := SystemConfig.Server.Key
	key = strings.Replace(key, " ", "", -1); if key == "" {
		return "", http_err.GetEnvKeyFail()
	}
	// token格式错误
	tokenList := strings.Split(token, "."); if len(tokenList) != 2 {
		return "", http_err.TokenVerificationFail()
	}

	verificationCode := tokenList[1]
	payloadString := tokenList[0]
	// token验证失败
	if encryptionString(payloadString, key) != verificationCode {
		return "", http_err.TokenVerificationFail()
	}
	// base64解码
	infoString, err := base64.StdEncoding.DecodeString(payloadString); if err != nil {
		return "", http_err.TokenVerificationFail()
	}
	// 反序列化
	err = json.Unmarshal(infoString, &fileToken); if err != nil {
		return "", http_err.TokenVerificationFail()
	}
	// 验证token使用类型
	if fileToken.TokenType & tokenType == 0 {
		return "", http_err.TokenVerificationFail()
	}
	// 验证token是否过期
	if time.Now().Unix() >= fileToken.ExpireAt {
		return "", http_err.TokenVerificationFail()
	}
	return fileToken.Hash, nil
}

/**
加密字符串
 */
func encryptionString(payloadString string, key string) string {
	ghmac := hmac.New(sha256.New, []byte(key))
	ghmac.Write([]byte(payloadString))
	return hex.EncodeToString(ghmac.Sum([]byte(nil)))
}


