package service

import (
	"bcg/ecommerce/api/response"
	entities "bcg/ecommerce/data/entities"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ItemServiceTestSuite struct {
	suite.Suite
	itemsMap map[string]entities.ItemEntity
	service  ItemsService
}

func (suite *ItemServiceTestSuite) SetupTest() {
	suite.itemsMap = map[string]entities.ItemEntity{
		"001": {
			ID:        "001",
			BrandName: "BrandA",
			UnitPrice: 100,
			Discount: entities.DiscountEntity{
				MinimumUnits: 3,
				TotalPrice:   200,
			},
		},
		"002": {
			ID:        "002",
			BrandName: "BrandB",
			UnitPrice: 50,
		},
	}
	suite.service = NewItemsService(suite.itemsMap)
}

func TestItemServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ItemServiceTestSuite))
}

func (suite *ItemServiceTestSuite) TestValidate_Successfully() {
	suite.Nil(suite.service.Validate([]string{"001", "002"}))
}

func (suite *ItemServiceTestSuite) TestValidate_InvalidID() {
	suite.Equal("invalid id: 003", suite.service.Validate([]string{"001", "003"}).Error())
}

func (suite *ItemServiceTestSuite) TestCheckout_WithSinlgeItem() {
	suite.Equal(response.Checkout{Price: 50}, suite.service.Checkout([]string{"002"}))
}

func (suite *ItemServiceTestSuite) TestCheckout_WithItemsHavingDiscount() {
	suite.Equal(response.Checkout{Price: 300}, suite.service.Checkout([]string{"001", "001", "001", "001"}))
}

func (suite *ItemServiceTestSuite) TestCheckout_WithMultipleItems() {
	suite.Equal(response.Checkout{Price: 350}, suite.service.Checkout([]string{"001", "001", "001", "001", "002"}))
}
