package http_err

func TokenVerificationFail() (ctx LiumaExceptBase) {
	ctx.Status = false
	ctx.Code = 5531
	ctx.Msg = "token验证失败"
	return ctx
}

func ModifyTokenFail() (ctx LiumaExceptBase){
	ctx.Status = false
	ctx.Code = 5532
	ctx.Msg = "修改token失败"
	return ctx
}

func SystemTokenVerificationFail() (ctx LiumaExceptBase) {
	ctx.Status = false
	ctx.Code = 5533
	ctx.Msg = "系统token验证失败"
	return ctx
}
