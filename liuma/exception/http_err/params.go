package http_err

import "fmt"

func UnmarshalBodyError () (ctx LiumaExceptBase){
	ctx.Status = false
	ctx.Code = 5521
	ctx.Msg = "读取body数据错误"
	return ctx
}

func MarshalFail() (ctx LiumaExceptBase){
	ctx.Status = false
	ctx.Code = 5522
	ctx.Msg = "序列化失败"
	return ctx
}

func LengthIsNotAllow(ob string, min int, max int) (ctx LiumaExceptBase) {

	msg := fmt.Sprintf("%s参数长度", ob)
	if min != -1 {
		msg += fmt.Sprintf("不得小于%d", min)
	}
	if max != -1 {
		msg += fmt.Sprintf("不得大于%d", max)
	}

	ctx.Status = false
	ctx.Code = 5523
	ctx.Msg = msg
	return ctx
}

func LackParams(par string) (ctx LiumaExceptBase){
	ctx.Status = false
	ctx.Code = 5524
	ctx.Msg = fmt.Sprintf("缺少%s参数", par)
	return ctx
}