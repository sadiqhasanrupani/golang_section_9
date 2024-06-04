package types

type PricesData struct {
	CommonGSTRate float64 `json:"commonGSTRate"` // this `json:"commonGSTRate"` is indeed a struct tag.
	Products      []struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	} `json:"products"`
}

type GSTIncludedProductsData struct {
	ProdName      string  `json:"name"`
	OriginalPrice float64 `json:"original_price"`
	GSTRate       float64 `json:"gst_rate"`
	Total         float64 `json:"total"`
}
