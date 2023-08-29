package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/munsheerck79/Ecom_project.git/pkg/api/handler"
	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/pkg/mocks"
)

func TestGetCoupon(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	couponServiceMock := mocks.NewMockCouponService(ctrl)
	couponHandler := handler.NewCouponHandler(couponServiceMock)

	router := gin.Default()
	router.GET("/user/coupon", couponHandler.GetCoupon)

	req, _ := http.NewRequest("GET", "/user/coupon", nil)
	resp := httptest.NewRecorder()

	expectedCoupons := []domain.Coupon{
		{
			ID:                1,
			Code:              "112233",
			MinOrderValue:     100.00,
			DiscountMaxAmount: 200.00,
			DiscountPercent:   10,
			Description:       "addcoupon of 100min and 200max",
			ValidTill:         10,
		},
		{
			ID:                2,
			Code:              "112234",
			MinOrderValue:     104.00,
			DiscountMaxAmount: 204.00,
			DiscountPercent:   10,
			Description:       "addcoupon2 of 100min and 200max",
			ValidTill:         10,
		},
	}
	couponServiceMock.EXPECT().GetCoupon(gomock.Any()).Return(expectedCoupons, nil)

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response struct {
		Message string          `json:"Message"`
		Data    []domain.Coupon `json:"Data"`
	}
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "List coupons successful", response.Message)
	assert.Equal(t, expectedCoupons, response.Data)
}
