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

func getContextId1(c *gin.Context) int {
	idInterface, _ := c.Get("Id1")
	return idInterface.(int)
}

func Login(c *gin.Context) {
	id := getContextId(c)
	fmt.Println(id)
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

func Register(c *gin.Context) {
	id := getContextId(c)
	body := loginForm{}
	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusOK, utils.SendResult(400, fmt.Sprintf("系统错误：%v", err), nil))
		return
	}

	code, msg := database.Register(id, body.User, body.Pass)
	if code != 0 {
		c.JSON(http.StatusOK, utils.SendResult(400, msg, nil))
		return
	}
	c.JSON(http.StatusOK, utils.SendResult(200, msg, nil))
}

func ClothShow(c *gin.Context) {
	id := getContextId(c)
	var cloth database.Cloths
	if database.CheckBoy(id) {
		cloth = database.ShowBoy(id)
	} else {
		cloth = database.ShowGirl(id)
	}
	c.JSON(200, utils.SendResult(200, "展示衣服成功", cloth))
}

func UserShow(c *gin.Context) {
	id := getContextId(c)
	user := database.ShowUser(id)
	c.JSON(200, utils.SendResult(200, "展示用户成功", user))
}

func Buy(c *gin.Context) {
	id := getContextId(c)
	id1 := getContextId1(c)
	body := buyForm{}
	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusOK, utils.SendResult(400, fmt.Sprintf("系统错误：%v", err), nil))
		return
	}
	number := body.Number
	money := database.GetPeopleMoney(id)
	totalScale := database.GetClothMoney(id1) * number

	if money-totalScale >= 0 && number <= database.TotalNumber(id1) {
		database.SaleMoney(id, totalScale)
		database.SaleCloth(id1, number)
		c.JSON(http.StatusOK, utils.SendResult(200, "购买成功", nil))
	} else {
		c.JSON(http.StatusOK, utils.SendResult(400, "购买失败", nil))
	}
}

func Spend(c *gin.Context) {
	id := getContextId(c)
	body := SpendForm{}
	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusOK, utils.SendResult(400, fmt.Sprintf("系统错误：%v", err), nil))
		return
	}
	database.Spend(id, body.Money)
	c.JSON(http.StatusOK, utils.SendResult(200, "消费成功", nil))
}
