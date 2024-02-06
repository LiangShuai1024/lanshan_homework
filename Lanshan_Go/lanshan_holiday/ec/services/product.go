package services

import (
	"demo/ec/common/request"
	"demo/ec/global"
	"demo/ec/models"
	"errors"
)

type productService struct {
}

var ProductService = new(productService)

// Create 创建商品
func (productService *productService) Create(params request.Create) (err error, product models.Product) {
	var result = global.Ec.DB.Where("ProductID = ?", params.ProductID).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("商品ID已存在")
		return
	}
	product = models.Product{Name: params.Name, ProductID: params.ProductID, Price: params.Price, Num: params.Num}
	err = global.Ec.DB.Create(&product).Error
	return
}

// GetProductInfo 获取商品信息
func (productService *productService) GetProductInfo(id string) (err error, product models.Product) {
	//intId, err := strconv.Atoi(id)
	err = global.Ec.DB.First(&product, id).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}
