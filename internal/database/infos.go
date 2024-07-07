package database

import (
	"taobao_backend/internal/utils"
)

func checkPassWord(id int, username string, password string) (int, string) {
	var user users
	result := instance.Select("id").Where("id = ? and username = ? and password = ?",
		id, username, password).First(&user)
	if result.Error != nil {
		return -1, "账号或密码错误"
	}
	return 0, "登录成功"
}

func Login(id int, username string, password string) (int, string, string) {
	code, msg := checkPassWord(id, username, password)
	if code != 0 {
		return code, msg, ""
	}
	token := utils.Jwt_generate(id)
	return code, msg, token
}
