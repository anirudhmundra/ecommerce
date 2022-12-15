package data

import (
	dto "bcg/ecommerce/data/dto"
	entities "bcg/ecommerce/data/entities"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadData(fileName string) (map[string]entities.ItemEntity, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var items []dto.Item
	err = json.Unmarshal(bytes, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, errors.New("no data to load")
	}
	return transformItemsArrayToMap(items)
}

func transformItemsArrayToMap(items []dto.Item) (map[string]entities.ItemEntity, error) {
	result := map[string]entities.ItemEntity{}
	for _, item := range items {
		minUnits, totalPrice, err := getDiscount(item.Discount)
		if err != nil {
			return nil, fmt.Errorf("error for item id %s: %-v", item.ID, err)
		}
		result[item.ID] = entities.ItemEntity{
			ID:        item.ID,
			BrandName: item.BrandName,
			UnitPrice: item.UnitPrice,
			Discount: entities.DiscountEntity{
				MinimumUnits: minUnits,
				TotalPrice:   totalPrice,
			},
		}
	}
	return result, nil
}

func getDiscount(discountStr string) (int32, int32, error) {
	if discountStr == "" {
		return 0, 0, nil
	}
	discount := strings.Split(strings.ToLower(discountStr), " for ")
	minUnits, err := strconv.Atoi(discount[0])
	if err != nil {
		return 0, 0, err
	}
	totalPrice, err := strconv.Atoi(discount[1])
	if err != nil {
		return 0, 0, err
	}
	return int32(minUnits), int32(totalPrice), nil
}
