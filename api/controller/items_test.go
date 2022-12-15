package controller

import (
	"bcg/ecommerce/api/mocks"
	"bcg/ecommerce/api/response"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type ItemsControllerTestSuite struct {
	suite.Suite
	mockCtrl   *gomock.Controller
	service    *mocks.MockItemsService
	controller ItemsController
}

func (suite *ItemsControllerTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.service = mocks.NewMockItemsService(suite.mockCtrl)
	suite.controller = NewItemsController(suite.service)
}

func TestItemsControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ItemsControllerTestSuite))
}

func (suite *ItemsControllerTestSuite) TestCheckout_Successfully() {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodPost, "/checkout", strings.NewReader(`["001","002"]`))
	ctx.Request = req

	suite.service.EXPECT().Validate([]string{"001", "002"}).Return(nil)
	suite.service.EXPECT().Checkout([]string{"001", "002"}).Return(response.Checkout{Price: 200})

	suite.controller.Checkout(ctx)

	suite.Assert().Equal(http.StatusOK, w.Result().StatusCode)
	suite.Assert().Equal(`{"price":200}`, w.Body.String())
}

func (suite *ItemsControllerTestSuite) TestCheckout_ValidationError() {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodPost, "/checkout", strings.NewReader(`["001","002"]`))
	ctx.Request = req

	suite.service.EXPECT().Validate([]string{"001", "002"}).Return(errors.New("error occurred"))

	suite.controller.Checkout(ctx)

	suite.Assert().Equal(http.StatusBadRequest, w.Result().StatusCode)
	suite.Assert().Equal(`{"message":"error occurred"}`, w.Body.String())
}

func (suite *ItemsControllerTestSuite) TestCheckout_InvalidJSONRequest() {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(http.MethodPost, "/checkout", strings.NewReader(`["001"002"]`))
	ctx.Request = req

	suite.controller.Checkout(ctx)

	suite.Assert().Equal(http.StatusBadRequest, w.Result().StatusCode)
	suite.Assert().Equal(`{"message":"invalid character '0' after array element"}`, w.Body.String())
}
