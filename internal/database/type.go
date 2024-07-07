package database

type users struct {
	Id       int    `gorm:"column:id;AUTO_INCREMENT"`
	Username string `json:"username"`
	Password string `json:"password"`
	Money    int    `json:"money"`
	Address  string `json:"address"`
}

type Cloths struct {
	Id        int    `gorm:"column:id;AUTO_INCREMENT"`
	Price     int    `json:"price"`
	BrandName string `json:"brand_name"`
	Desc      string `json:"desc"`
	Color     string `json:"color"`
	Size      int    `json:"size"`
	Numb      int    `json:"numb"`
	Gender    int    `json:"gender"`
	PicURL    string `json:"pic_url"`
	Status    string `json:"status"`
}
