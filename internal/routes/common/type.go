package common

type loginForm struct {
	User string `json:"username"`
	Pass string `json:"password"`
}

type tokenRes struct {
	Token string `json:"token"`
}

type buyForm struct {
	Number int `json:"number"`
}

type SpendForm struct {
	Money int `json:"money"`
}
