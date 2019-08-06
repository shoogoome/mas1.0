package http_err

import "fmt"

func SaveFileInfoError(err error) (ctx LiumaExceptBase) {
	ctx.Status = false
	ctx.Code = 5551
	ctx.Msg = fmt.Sprintf("保存文件信息错误: %v", err)
	return ctx
}
