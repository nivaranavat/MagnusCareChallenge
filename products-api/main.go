package main

import (
	"encoding/csv"
	"fmt"
	"os"
	handler "products-api/httpd/handler"
	product "products-api/platform/product"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Server Started")
	products := product.New()

	//read the tsv file that was provided and add those values individually
	tsvFile, err := os.Open("sample_product_data.tsv")

	if err != nil {
		fmt.Println(err)
	}

	defer tsvFile.Close()

	reader := csv.NewReader(tsvFile)

	reader.Comma = '\t'      // Use tab-delimited instead of comma <---- here!
	reader.LazyQuotes = true //allow for quotations
	reader.FieldsPerRecord = -1

	tsvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, each := range tsvData {
		item := product.Item{each[0], each[1], each[2], each[3], each[4], each[5]}
		products.Add(item)
	}

	//create a gin object that will used to make the api calls
	r := gin.Default()

	r.GET("api/products", handler.ProductGet(products))
	r.POST("api/products", handler.ProductPost(products))
	r.POST("api/products/search", handler.ProductSearch(products))

	fmt.Println("Server Running")
	r.Run(":8088")

	fmt.Println("Server Stopped")

}
