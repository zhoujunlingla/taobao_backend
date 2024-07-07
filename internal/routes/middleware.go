package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"taobao_backend/internal/utils"
)

func cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method

		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Allow-Headers", "*")
		context.Set("content-type", "application/json")

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}
		//处理请求
		context.Next()
	}
}

func CheckId() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Request.URL.Query().Get("id")
		id1 := context.Request.URL.Query().Get("id1")
		if id == "" {
			context.JSON(200, utils.SendResult(400, "参数错误: id", nil))
			context.Abort()
			return
		}
		idInt, _ := strconv.Atoi(id)
		idInt1, _ := strconv.Atoi(id1)
		context.Set("Id", idInt)
		context.Set("Id1", idInt1)
		context.Next()
	}
}

func CheckToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token")
		if token == "" {
			context.JSON(200, utils.SendResult(400, "无权限", nil))
			context.Abort()
			return
		}
		code, msg, id := utils.Jwt_verify(token)
		IdFromUrl, _ := context.Get("Id")
		idFromUrlN := IdFromUrl.(int)

		if code != 0 {
			context.JSON(200, utils.SendResult(400, msg, nil))
			context.Abort()
			return
		} else if id != idFromUrlN {
			context.JSON(200, utils.SendResult(400, "无权限", nil))
			context.Abort()
			return
		} else {
			context.Next()
		}
	}
}
