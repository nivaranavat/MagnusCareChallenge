package handler

import (
	"net/http"
	product "products-api/platform/product"

	"github.com/gin-gonic/gin"
)

type productPostRequest struct {
	ProductId    string "json:productId"
	Title        string "json:title"
	BrandId      string "json:brandId"
	BrandName    string "json:brandName"
	CategoryId   string "json:categoryId"
	CategoryName string "json:categoryName"
}

func ProductPost(products *product.ProductData) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := productPostRequest{}
		c.Bind(&requestBody)

		item := product.Item{
			ProductId:    requestBody.ProductId,
			Title:        requestBody.Title,
			BrandId:      requestBody.BrandId,
			BrandName:    requestBody.BrandName,
			CategoryId:   requestBody.CategoryId,
			CategoryName: requestBody.CategoryName,
		}
		products.Add(item)
		c.Status(http.StatusNoContent)
	}
}
