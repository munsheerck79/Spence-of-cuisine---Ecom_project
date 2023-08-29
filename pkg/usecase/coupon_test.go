package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/pkg/mocks"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/stretchr/testify/assert"
)

func TestGetCoupon(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCouponRepository(ctrl)
	couponService := NewCouponService(mockRepo)

	ctx := context.Background()
	expectedCoupons := []domain.Coupon{
		{ID: 1, Code: "12345", Description: "Coupon 1", MinOrderValue: 100.00, DiscountMaxAmount: 200.00, DiscountPercent: 10, ValidTill: 10},
		{ID: 2, Code: "67890", Description: "Coupon 2", MinOrderValue: 200.00, DiscountMaxAmount: 300.00, DiscountPercent: 15, ValidTill: 10},
	}

	mockRepo.EXPECT().GetCoupon(gomock.Any()).Return(expectedCoupons, nil)

	coupons, err := couponService.GetCoupon(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedCoupons, coupons)

}

func TestAddCoupon(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCouponRepository(ctrl)
	couponService := NewCouponService(mockRepo)

	mockRepo.EXPECT().AddCoupon(gomock.Any(), gomock.Any()).Return(nil)

	err := couponService.AddCoupon(context.Background(), domain.Coupon{
		Code:              "112233",
		MinOrderValue:     100.00,
		DiscountMaxAmount: 200.00,
		DiscountPercent:   10,
		Description:       "addcoupon of 100min and 200max",
		ValidTill:         10,
	})

	assert.NoError(t, err)
}

func TestEditCoupon(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCouponRepository(ctrl)
	couponService := NewCouponService(mockRepo)

	mockRepo.EXPECT().GetCouponByCode(gomock.Any(), "112233").Return(domain.Coupon{
		ID:                1,
		Code:              "112233",
		MinOrderValue:     100.00,
		DiscountMaxAmount: 200.00,
		DiscountPercent:   10,
		Description:       "addcoupon of 100min and 200max",
		ValidTill:         10,
	}, nil)

	mockRepo.EXPECT().EditCoupon(gomock.Any(), gomock.Any()).Return(nil)

	err := couponService.EditCoupon(context.Background(), request.EditCoupon{
		Code:              "112233",
		MinOrderValue:     100.00,
		DiscountMaxAmount: 200.00,
		DiscountPercent:   10,
		Description:       "addcoupon of 100min and 200max",
		ValidTill:         10,
	})

	assert.NoError(t, err)
}
