package models

type Product struct {
	ID
	ProductID string `gorm:"primary_key"`
	Name      string
	//CategoryID    int
	//Title   string
	Info string `gorm:"size:1000"`
	//ImgPath string
	Price string
	//DiscountPrice string
	OnSale string
	Num    string
	Timestamps
	SoftDeletes
	//BossID     int
	//BossName   string
	//BossAvatar string
}
