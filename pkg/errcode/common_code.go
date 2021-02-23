package errcode

var (
	Success                  = NewError(0, "成功")
	ServerError              = NewError(10000000, "服务内部错误")
	InvalidParams            = NewError(10000001, "入参错误")
	NotFound                 = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist = NewError(10000003, "签权失败，找不到对应AppKey和AppServer")
	UnauthorizedTokenError   = NewError(10000004, "签权失败，Token错误")
	UnauthorizedTokenTimeout = NewError(10000003, "签权失败，Token超时")
	UnauthorizedTokeGenerate = NewError(10000003, "签权失败，Token生成失败")
	TooManySequests          = NewError(10000007, "请求过多")
)
