package integration_tests

import (
	entities "bcg/ecommerce/data/entities"
	"bcg/ecommerce/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckout(t *testing.T) {
	router := router.SetupRouter(map[string]entities.ItemEntity{
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
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/checkout", strings.NewReader(`["001","001","001","001","002"]`))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"price":350}`, w.Body.String())
}
