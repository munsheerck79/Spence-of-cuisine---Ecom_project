package interfacess

import (
	"context"

	"github.com/munsheerck79/Ecom_project.git/util/response"
)

type PaymentService interface {
	VerifyRazorPayPayment(c context.Context, razorpayOrderId string, status string) (response.Order, error)
	AddMonyToWallet(ctx context.Context, userId uint, RazorPayPaymentId string, money float32) error
	GetWallet(c context.Context,userid uint)(response.WalletRes,error)
	AddPaymentId(paymentID string,RazorPayOrderId string)error
	AddTempData(c context.Context,data response.Order)error
}
