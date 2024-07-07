package main

import (
	"taobao_backend/internal/app"
	"taobao_backend/internal/utils"
)

func main() {
	utils.Version = "1.0.1"
	app.Run()
}
