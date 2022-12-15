package data

import (
	"testing"

	entities "bcg/ecommerce/data/entities"

	"github.com/stretchr/testify/assert"
)

func TestLoadData_Successfully(t *testing.T) {
	itemsMap, err := LoadData("testdata/testItems.json")

	assert.Equal(t, map[string]entities.ItemEntity{
		"001": {ID: "001", BrandName: "Rolex", UnitPrice: 100, Discount: entities.DiscountEntity{MinimumUnits: 3, TotalPrice: 200}},
		"004": {ID: "004", BrandName: "Casio", UnitPrice: 30, Discount: entities.DiscountEntity{MinimumUnits: 0, TotalPrice: 0}}}, itemsMap)
	assert.Nil(t, err)
}

func TestLoadData_NoFileFound(t *testing.T) {
	itemsMap, err := LoadData("testdata/testItems1.json")

	assert.Nil(t, itemsMap)
	assert.Equal(t, "open testdata/testItems1.json: The system cannot find the file specified.", err.Error())
}

func TestLoadData_InvalidDiscountedUnits(t *testing.T) {
	itemsMap, err := LoadData("testdata/invalidTestItems_discountedUnits.json")

	assert.Nil(t, itemsMap)
	assert.Equal(t, "error for item id 001: strconv.Atoi: parsing \"3a\": invalid syntax", err.Error())
}


func TestLoadData_InvalidDiscountedPrice(t *testing.T) {
	itemsMap, err := LoadData("testdata/invalidTestItems_discountedPrice.json")

	assert.Nil(t, itemsMap)
	assert.Equal(t, "error for item id 001: strconv.Atoi: parsing \"2x0\": invalid syntax", err.Error())
}

func TestLoadData_EmptyFile(t *testing.T) {
	itemsMap, err := LoadData("testdata/emptyTestItems.json")

	assert.Nil(t, itemsMap)
	assert.Equal(t, "no data to load", err.Error())
}
