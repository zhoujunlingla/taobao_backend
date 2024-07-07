package database

import (
	"taobao_backend/internal/utils"
)

func checkPassWord(id int, username string, password string) (int, string) {
	var user users
	result := instance.Select("id").Where("username = ? and password = ?",
		username, password).First(&user)
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

func Register(id int, username string, password string) (int, string) {
	newUser := users{
		Username: username,
		Password: utils.MD5(password),
		Money:    0,
		Address:  "杭州",
	}
	result := instance.Create(&newUser)
	if result.Error != nil {
		return -1, "注册失败"
	}
	return 0, "注册成功"
}

func ShowBoy(id int) Cloths {
	var ShowCloth Cloths
	instance.Select("price", "brand_name", "desc", "color", "size", "numb", "pic_url", "status").Where("id = ?", id).First(&ShowCloth)
	return ShowCloth
}
func ShowGirl(id int) Cloths {
	var ShowCloth Cloths
	instance.Select("price", "brand_name", "desc", "color", "size", "numb", "pic_url", "status").Where("id = ?", id).First(&ShowCloth)
	return ShowCloth
}

func CheckBoy(id int) bool {
	var cloth Cloths
	instance.Select(id).First(&cloth)
	if cloth.Gender == 0 {
		return true
	} else {
		return false
	}
}

func GetPeopleMoney(id int) int {
	var user users
	instance.Select("Money").Where("id = ?", id).First(&user)
	return user.Money
}

func GetClothMoney(id int) int {
	var cloth Cloths
	instance.Select("Price").Where("id = ?", id).First(&cloth)
	return cloth.Price
}

func GetClothTotal(id int) int {
	var cloth Cloths
	instance.Select("Price", "Numb").Where("id = ?", id).First(&cloth)
	return cloth.Price * cloth.Numb
}

func TotalNumber(id int) int {
	var cloth Cloths
	instance.Select("Numb").Where("id = ?", id).First(&cloth)
	return cloth.Numb
}
func SaleCloth(id int, number int) {
	var cloth Cloths
	instance.Select("Numb").Where("id = ?", id).First(&cloth)
	cloth.Numb = cloth.Numb - number
	instance.Model(&cloth).Where("id = ?", id).Update("Numb", cloth.Numb)
}

func SaleMoney(id int, money int) {
	var user users
	instance.Select("Money").Where("id = ?", id).First(&user)
	user.Money -= money
	instance.Model(&user).Where("id = ?", id).Update("Money", user.Money)
}
