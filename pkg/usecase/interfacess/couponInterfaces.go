package interfacess

import (
	"context"

	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/util/request"
)

type CouponService interface {
	GetCoupon(c context.Context) ([]domain.Coupon, error)
	AddCoupon(c context.Context, body domain.Coupon) error
	EditCoupon(c context.Context, body request.EditCoupon) error
}
