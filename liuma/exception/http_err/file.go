package http_err

import "fmt"

func GetFileFail() (ctx LiumaExceptBase){
	ctx.Status = false
	ctx.Code = 5541
	ctx.Msg = "获取文件失败"
	return ctx
}

func CalculateHashError() (ctx LiumaExceptBase) {
	ctx.Status = false
	ctx.Code = 5542
	ctx.Msg = "计算文件hash失败"
	return ctx
}

func TokenFail() (ctx LiumaExceptBase) {
	ctx.Status = false
	ctx.Code = 5543
	ctx.Msg = "token无效或过期，请重新获取token"
	return ctx
}

func DamageToRawData() (ctx LiumaExceptBase) {
	ctx.Status = false
	ctx.Code = 5544
	ctx.Msg = "原始文件损坏，请重新上传"
	return ctx
}

func UploadFail() (ctx LiumaExceptBase) {
	ctx.Status = false
	ctx.Code = 5545
	ctx.Msg = "上传失败"
	return ctx
}

func ChuckExists() (ctx LiumaExceptBase) {
	ctx.Status = false
	ctx.Code = 5545
	ctx.Msg = "分片存在"
	return ctx
}

func StorageUnexpectedTermination(err error) (ctx LiumaExceptBase) {
	ctx.Status = false
	ctx.Code = 5546
	ctx.Msg = fmt.Sprintf("存储意外终止: %v", err)
	return ctx
}

func FileIsNotExists() (ctx LiumaExceptBase) {
	ctx.Status = false
	ctx.Code = 5547
	ctx.Msg = "文件不存在"
	return ctx
}

func ResendOver() (ctx LiumaExceptBase) {
	ctx.Status = false
	ctx.Code = 5548
	ctx.Msg = "分片重发超出设定"
	return ctx
}

func DownloadFail() (ctx LiumaExceptBase) {
	ctx.Status = false
	ctx.Code = 5549
	ctx.Msg = "下载失败"
	return ctx
}

