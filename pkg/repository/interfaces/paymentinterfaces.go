package interfaces

import (
	"context"

	"github.com/munsheerck79/Ecom_project.git/util/response"
)

type PaymentRepository interface {
	AddMonyToWallet(ctx context.Context, userId uint, RazorPayPaymentId string, money float32) error
	
	FindTempData(ctx context.Context, RazorPayKey string) (response.Order, error)
	//DeleteRazorData(c context.Context, order_id string) error
	GetWallet(c context.Context,userid uint)(response.WalletRes,error)
	AddPaymentId(paymentID string,RazorPayOrderId string)error
}
