package request

type Create struct {
	Name      string `form:"name" json:"name" binding:"required"`
	ProductID string `form:"productid" json:"productid" binding:"required"`
	Price     string `form:"price" json:"price" binding:"required"`
	Num       int    `form:"num" json:"num" binding:"required"`
}

func (create Create) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":  "商品名称不能为空",
		"price.required": "商品价格不能为空",
		"num.required":   "商品数量不能为空",
	}
}
