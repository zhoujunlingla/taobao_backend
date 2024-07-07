package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"taobao_backend/internal/database"
	"taobao_backend/internal/utils"
)

func getContextId(c *gin.Context) int {
	idInterface, _ := c.Get("Id")
	return idInterface.(int)
}

func Login(c *gin.Context) {
	id := getContextId(c)
	body := loginForm{}
	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusOK, utils.SendResult(400, fmt.Sprintf("系统错误：%v", err), nil))
		return
	}

	body.Pass = utils.MD5(body.Pass)
	code, msg, token := database.Login(id, body.User, body.Pass)
	if code != 0 {
		c.JSON(http.StatusOK, utils.SendResult(400, msg, nil))
		return
	}
	c.JSON(http.StatusOK, utils.SendResult(200, msg, tokenRes{Token: token}))
}
