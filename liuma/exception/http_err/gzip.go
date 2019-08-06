package http_err

func GzipFail() (ctx LiumaExceptBase) {
	ctx.Status = false
	ctx.Code = 5561
	ctx.Msg = "压缩文件失败"
	return ctx
}

func GunzipFail() (ctx LiumaExceptBase) {
	ctx.Status = false
	ctx.Code = 5562
	ctx.Msg = "解压缩文件失败"
	return ctx
}
