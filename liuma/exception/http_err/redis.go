package http_err

func RedisConnectExcept () (ctx LiumaExceptBase){
	ctx.Status = false
	ctx.Code = 5511
	ctx.Msg = "redis连接错误"
	return ctx
}

func RedisVerificationError () (ctx LiumaExceptBase) {
	ctx.Status = false
	ctx.Code = 5512
	ctx.Msg = "redis验证错误"
	return ctx
}

