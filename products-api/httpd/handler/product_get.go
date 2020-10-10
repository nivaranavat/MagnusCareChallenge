package handler

import (
	"net/http"
	product "products-api/platform/product"

	"github.com/gin-gonic/gin"
)

func ProductGet(products *product.ProductData) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := products.GetAll()
		jsonResponse := map[string][]product.Item{}
		jsonResponse["Products"] = results
		c.JSON(http.StatusOK, jsonResponse)
	}
}
