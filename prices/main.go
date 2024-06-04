package prices

import (
	"section9.go/types"
)

type TaxIncludedPriceJob struct {
	ProdName string
	Prices   float64
	GSTRate  float64
	Total    float64
}

func (job *TaxIncludedPriceJob) CalculateTotal(index int) types.GSTIncludedProductsData {
	textRate := job.GSTRate / 100
	taxRate := textRate * job.Prices
	total := job.Prices + taxRate
	job.Total = total

	return types.GSTIncludedProductsData{
		ProdName:      job.ProdName,
		OriginalPrice: job.Prices,
		GSTRate:       job.GSTRate,
		Total:         total,
	}
}

// constructor
func NewTaxIncludedPriceJob(price float64, GSTRate float64, name string) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{Prices: price, GSTRate: GSTRate, ProdName: name}
}
