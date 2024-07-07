package common

type loginForm struct {
	User string `form:"username"`
	Pass string `form:"password"`
}

type tokenRes struct {
	Token string `json:"token"`
}

type buyForm struct {
	Number int `form:"number"`
}
