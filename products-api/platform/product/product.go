package product

import (
	"errors"
)

type Item struct {
	ProductId    string "json:productId"
	Title        string "json:title"
	BrandId      string "json:brandId"
	BrandName    string "json:brandName"
	CategoryId   string "json:categoryId"
	CategoryName string "json:categoryName"
}

type DS map[string][]string

type BrandIndex struct {
	BrandIds map[string]string
}

type CategoryIndex struct {
	CategoryIds map[string]string
}

type ProductIndex struct {
	Titles        DS
	BrandNames    DS
	CategoryNames DS
	BrandIndex    BrandIndex
	CategoryIndex CategoryIndex
}

type ProductData struct {
	Products     map[string]Item
	ProductIndex ProductIndex
}

func New() *ProductData {
	return &ProductData{
		Products: map[string]Item{},
		ProductIndex: ProductIndex{
			Titles:        DS{},
			BrandNames:    DS{},
			CategoryNames: DS{},
			BrandIndex: BrandIndex{
				BrandIds: map[string]string{},
			},
			CategoryIndex: CategoryIndex{
				CategoryIds: map[string]string{},
			},
		},
	}

}

func (pd *ProductData) Add(item Item) {
	pd.Products[item.ProductId] = item
	product_id := item.ProductId
	brand_id := item.BrandId
	category_id := item.CategoryId

	pd.Products[product_id] = item

	pd.ProductIndex.Titles[item.Title] = append(pd.ProductIndex.Titles[item.Title], product_id)

	pd.ProductIndex.BrandIndex.BrandIds[brand_id] = item.BrandName
	pd.ProductIndex.BrandNames[item.BrandName] = append(pd.ProductIndex.BrandNames[item.BrandName], product_id)

	pd.ProductIndex.CategoryIndex.CategoryIds[category_id] = item.CategoryName
	pd.ProductIndex.CategoryNames[item.CategoryName] = append(pd.ProductIndex.CategoryNames[item.CategoryName], product_id)

}

func (pd *ProductData) GetAll() map[string]Item {
	return pd.Products
}

func (pd *ProductData) Find(group string, values []string) ([]Item, error) {

	product_ids := map[string]bool{}

	if group == "productId" {
		for _, value := range values {
			product_ids[value] = true
		}
	} else if group == "title" {
		for _, value := range values {
			for _, id := range pd.ProductIndex.Titles[value] {
				product_ids[id] = true
			}
		}
	} else if group == "brandName" {
		for _, value := range values {
			for _, id := range pd.ProductIndex.BrandNames[value] {
				product_ids[id] = true
			}
		}
	} else if group == "brandId" {
		for _, value := range values {
			brand_name := pd.ProductIndex.BrandIndex.BrandIds[value]
			for _, id := range pd.ProductIndex.BrandNames[brand_name] {
				product_ids[id] = true
			}
		}
	} else if group == "categoryName" {
		for _, value := range values {
			for _, id := range pd.ProductIndex.CategoryNames[value] {
				product_ids[id] = true
			}
		}
	} else if group == "categoryId" {
		for _, value := range values {
			category_name := pd.ProductIndex.CategoryIndex.CategoryIds[value]
			for _, id := range pd.ProductIndex.CategoryNames[category_name] {
				product_ids[id] = true
			}
		}
	} else {
		//there has been a incorrect input in the query that doesn't exist in our database
		return nil, errors.New("InvalidQueryException")
	}

	results := []Item{}

	for id, _ := range product_ids {
		results = append(results, pd.Products[id])
	}

	return results, nil

}

func (pd *ProductData) ReturnPage(results []Item, from int, size int) ([]Item, error) {

	starting_index := from * size

	if len(results) <= size || starting_index >= len(results) {
		//there has been a error the pagination inputted is out of bounds of the data given or resulted
		return nil, errors.New("OutofBoundsException")
	} else {
		return results[starting_index : starting_index+size], nil
	}

}
