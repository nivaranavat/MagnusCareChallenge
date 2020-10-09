package handler

import (
	"Project1/platform/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductGet(products *product.ProductData) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := products.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
