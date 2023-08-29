package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	interfaces "github.com/munsheerck79/Ecom_project.git/pkg/repository/interfaces"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"gorm.io/gorm"
)

type couponDatabase struct {
	DB *gorm.DB
}

func NewCouponRepository(DB *gorm.DB) interfaces.CouponRepository {
	return &couponDatabase{DB: DB}
}

func (p *couponDatabase) GetCoupon(c context.Context) ([]domain.Coupon, error) {
	var coupons []domain.Coupon
	query := `SELECT * FROM coupons`
	if err := p.DB.Raw(query).Scan(&coupons).Error; err != nil {
		return coupons, fmt.Errorf("failed to show coupns ")
	}
	return coupons, nil
}

func (p *couponDatabase) GetCouponByCode(c context.Context, code string) (domain.Coupon, error) {
	var coupon domain.Coupon
	query := `SELECT * FROM coupons WHERE code = ?`
	if err := p.DB.Raw(query, code).Scan(&coupon).Error; err != nil {
		return coupon, fmt.Errorf("failed to show coupns ")
	}
	return coupon, nil
}

func (p *couponDatabase) AddCoupon(c context.Context, body domain.Coupon) error {

	query := `INSERT INTO coupons (Code,description,min_order_value,discount_percent,discount_max_amount,valid_till)
						  VALUES ($1,$2,$3,$4,$5,$6)`
	validAt := body.ValidTill * (time.Now().Add(24 * time.Hour).Unix())

	err := p.DB.Exec(query, body.Code, body.Description, body.MinOrderValue, body.DiscountPercent, body.DiscountMaxAmount, validAt).Error
	if err != nil {
		return fmt.Errorf("failed to add coupon ")
	}
	return nil
}

func (p *couponDatabase) EditCoupon(c context.Context, body request.EditCoupon) error {

	query := `UPDATE coupons SET min_order_value = $1,discount_percent =$2,discount_max_amount = $3,description =$4,valid_till = $5 WHERE code = $6`

	err := p.DB.Exec(query, body.MinOrderValue, body.DiscountPercent, body.DiscountMaxAmount, body.Description, body.ValidTill, body.Code).Error
	if err != nil {
		return fmt.Errorf("failed to save coupon id")
	}
	return nil
}
