package utils

var Version string = "unknown"

type M map[string]interface{}

func SendResult(code int, msg string, data interface{}) M {
	if data != nil {
		return M{
			"code": code,
			"msg":  msg,
			"data": data,
		}
	} else {
		return M{
			"code": code,
			"msg":  msg,
		}
	}
}
