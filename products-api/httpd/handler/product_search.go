package handler

import (
	"fmt"
	"net/http"
	product "products-api/platform/product"
	query "products-api/platform/search_query"

	"github.com/gin-gonic/gin"
)

type productSearchPostRequest struct {
	Conditions []query.Condition "json:conditions"
	Pagination query.Pagination  "json:pagination"
}

func ProductSearch(products *product.ProductData) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := productSearchPostRequest{}
		c.ShouldBind(&requestBody)

		q := query.Query{
			Conditions: requestBody.Conditions,
			Pagination: requestBody.Pagination,
		}

		results := []product.Item{}
		conditions := q.GetConditions()
		pagination := q.GetPagination()

		fmt.Println("Conditions:", conditions)
		fmt.Println("Pagination:", pagination)
		flag := false
		//loop through the different conditions and get results
		for _, cond := range conditions {
			group := cond.Type
			values := cond.Values
			search_result, err1 := products.Find(group, values)
			if err1 == nil {
				results = append(results, search_result...)
			} else {
				flag = true
				fmt.Println(err1)
				c.JSON(http.StatusConflict, "The Inputted Query is invalid")
				break
			}
		}

		//cut the answer based on pagination given or not given
		results, err2 := products.ReturnPage(results, pagination.From, pagination.Size)
		if !flag {
			if err2 != nil {
				fmt.Println(err2)
				c.JSON(http.StatusConflict, "The Inputted Pagination is invalid")
			} else {
				fmt.Println("Results for search query", results)
				c.JSON(http.StatusOK, results)
			}
		}

		//fmt.Println("All Products", products.GetAll())

	}
}
