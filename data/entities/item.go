package data

type ItemEntity struct {
	ID        string
	BrandName string
	UnitPrice int32
	Discount  DiscountEntity
}

type DiscountEntity struct {
	MinimumUnits int32
	TotalPrice   int32
}
