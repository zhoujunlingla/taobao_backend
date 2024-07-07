package routes

import (
	"fmt"
	"log"
	"net/http"
	"taobao_backend/config"
	"taobao_backend/internal/routes/common"
	"taobao_backend/internal/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(config.Cfg.AppMode)
	if config.Cfg.AppMode == "release" {
		log.Println("You are in production now... Be careful!!!")
	}
	router := gin.Default()
	router.Use(cors())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, utils.SendResult(200, fmt.Sprintf("service is runing @%s", utils.Version), nil))
	})
	router.Use(CheckId())
	router.POST("/login", common.Login)
	router.Use(CheckToken())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.SendResult(404, "Not Found", nil))
	})
	router.Run(config.Cfg.Port)
}
