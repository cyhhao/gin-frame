package res

var ErrorCode = map[int]string{
	0:     "ok",
	10001: "缺少参数",
	10002: "参数错误",

	40100: "未登录",

	40000: "自定义错误类型",
	50000: "服务器内部错误",
}
