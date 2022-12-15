package service

import (
	"bcg/ecommerce/api/response"
	entities "bcg/ecommerce/data/entities"
	"fmt"
)

type ItemsService interface {
	Checkout(ids []string) response.Checkout
	Validate(ids []string) error
}

type itemsService struct {
	itemsMap map[string]entities.ItemEntity
}

func NewItemsService(itemsMap map[string]entities.ItemEntity) ItemsService {
	return itemsService{itemsMap: itemsMap}
}

func (is itemsService) Checkout(ids []string) response.Checkout {
	itemCountMap := map[string]int32{}
	for _, id := range ids {
		if val, ok := itemCountMap[id]; ok {
			itemCountMap[id] = val + 1
		} else {
			itemCountMap[id] = 1
		}
	}

	var totalAmount int32
	for id, count := range itemCountMap {
		if item, ok := is.itemsMap[id]; ok {
			if item.Discount.MinimumUnits != 0 {
				totalAmount += count / item.Discount.MinimumUnits * item.Discount.TotalPrice
				totalAmount += count % item.Discount.MinimumUnits * item.UnitPrice
			} else {
				totalAmount += count * item.UnitPrice
			}
		}
	}

	return response.Checkout{Price: totalAmount}
}

func (is itemsService) Validate(ids []string) error {
	for _, id := range ids {
		if _, ok := is.itemsMap[id]; !ok {
			return fmt.Errorf("invalid id: %s", id)
		}
	}
	return nil
}
