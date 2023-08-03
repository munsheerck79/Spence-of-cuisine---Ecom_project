package interfacess

import (
	"context"

	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/munsheerck79/Ecom_project.git/util/response"
)

type OrderService interface {
	BuyProduct(c context.Context, userId uint) (domain.Address, error)
	OrderCartProducts(c context.Context, userId uint, body request.Order) (response.Order, string, error)
	UpDateOrderStatus(c context.Context, body request.UpDateOrderStatus) error
	GetCoupon(c context.Context) ([]domain.Coupon, error)
	AddCoupon(c context.Context, body domain.Coupon) error

	EditCoupon(c context.Context, body request.EditCoupon) error

	MakeRazorPayPayment(c context.Context, amount float32) (string, string, error)
	ReturnOrder(c context.Context, userId uint, orderID uint) (response.Order, error)
	CreateInvoice(c context.Context, userId uint, orderID uint) ([]byte, error)
	GetOrderDetails(c context.Context, userId uint, orderID uint) (response.Orders1, error)

	//ListCartord(ctx context.Context, userId uint) ([]response.Cart, float32, float32, error)
}
