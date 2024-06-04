package main

import (
	"fmt"

	"section9.go/filemanager"
	"section9.go/prices"
	"section9.go/types"
)

/* I'm creating a price calculator which will calculate prices of equivalent items which their taxes and returns a GTS calculated price. */
func main() {
	products, err := filemanager.LoadJsonData[types.PricesData]("prices.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	gstIncludedProductPrices := make([]types.GSTIncludedProductsData, 0)

	for index, value := range products.Products {
		count := index
		count++

		productPrice := prices.NewTaxIncludedPriceJob(value.Price, products.CommonGSTRate, value.Name)
		price := productPrice.CalculateTotal(index)
		// separate invoice_details file
		filemanager.WriteJsonData(fmt.Sprintf("invoice_%v.json", count), price)
		gstIncludedProductPrices = append(gstIncludedProductPrices, price)
	}

	type eager_result struct {
		products []types.GSTIncludedProductsData
	}
	result := eager_result{
		products: gstIncludedProductPrices,
	}

	filemanager.WriteJsonData("eager_result.json", result)
}
