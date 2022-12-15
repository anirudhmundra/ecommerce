package data

type Item struct {
	ID        string `json:"id"`
	BrandName string `json:"brandName"`
	UnitPrice int32  `json:"unitPrice"`
	Discount  string `json:"discount"`
}
