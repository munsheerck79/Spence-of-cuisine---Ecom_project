package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/pkg/repository/interfaces"
	"github.com/munsheerck79/Ecom_project.git/pkg/usecase/interfacess"
	"github.com/munsheerck79/Ecom_project.git/util/request"
)

type CouponUsecase struct {
	couponRepository interfaces.CouponRepository
}

func NewCouponService(repo interfaces.CouponRepository) interfacess.CouponService {
	return &CouponUsecase{couponRepository: repo}
}


func (p *CouponUsecase) GetCoupon(c context.Context) ([]domain.Coupon, error) {

	coupons, err := p.couponRepository.GetCoupon(c)
	if err != nil {
		return coupons, err
	}
	return coupons, nil

}
func (p *CouponUsecase) AddCoupon(c context.Context, body domain.Coupon) error {
	err := p.couponRepository.AddCoupon(c, body)
	if err != nil {
		return err
	}
	return nil

}
func (p *CouponUsecase) EditCoupon(c context.Context, body request.EditCoupon) error {

	coupon, err := p.couponRepository.GetCouponByCode(c, body.Code)
	if err != nil {
		return fmt.Errorf("%v coupon is not avalable err", body.Code)
	}
	if coupon.ID == 0 {
		return fmt.Errorf("%v coupon is not avalable", body.Code)
	}
	if body.Description == "" {
		body.Description = coupon.Description
	}
	if body.DiscountMaxAmount == 0 {
		body.DiscountMaxAmount = coupon.DiscountMaxAmount
	}
	if body.DiscountPercent == 0 {
		body.DiscountPercent = coupon.DiscountPercent
	}
	if body.MinOrderValue == 0 {
		body.MinOrderValue = coupon.MinOrderValue
	}
	if body.ValidTill == 0 {
		body.ValidTill = coupon.ValidTill
	} else {
		body.ValidTill = body.ValidTill * (time.Now().Add(24 * time.Hour).Unix())
	}
	err1 := p.couponRepository.EditCoupon(c, body)
	if err1 != nil {
		return err1
	}
	return nil

}
