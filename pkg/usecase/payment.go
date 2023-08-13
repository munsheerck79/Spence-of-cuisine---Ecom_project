package usecase

import (
	"context"
	"fmt"

	"github.com/munsheerck79/Ecom_project.git/pkg/repository/interfaces"
	"github.com/munsheerck79/Ecom_project.git/pkg/usecase/interfacess"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/munsheerck79/Ecom_project.git/util/response"
)

type PaymentUsecase struct {
	paymentRepository interfaces.PaymentRepository
	orderRepository   interfaces.OrderRepository
}

func NewPaymentService(repo interfaces.PaymentRepository, order interfaces.OrderRepository) interfacess.PaymentService {

	return &PaymentUsecase{paymentRepository: repo,
		orderRepository: order,
	}
}

func (p *PaymentUsecase) AddMonyToWallet(ctx context.Context, userId uint, RazorPayPaymentId string, money float32) error {

	err := p.paymentRepository.AddMonyToWallet(ctx, userId, RazorPayPaymentId, money)
	if err != nil {
		return err
	}
	return nil
}

func (p *PaymentUsecase) VerifyRazorPayPayment(c context.Context, razorpayOrderId string, status string) (response.Order, error) {
	Data, err := p.paymentRepository.FindTempData(c, razorpayOrderId)
	if err != nil {
		return Data, err
	}
	if status != "paid" {
		return Data, err
	}
	if razorpayOrderId != Data.RazorPayOrderId {
		return Data, err
	}
	return Data, nil
}

func (p *PaymentUsecase) AddTempData(c context.Context, data response.Order) error {

	userId := data.UsersID
	var b request.Order
	b.PaymentMethod = "wallet"
	b.CouponId = 0
	orderIdx, err := p.orderRepository.OrderProductsTemp(c, userId, b, data.NetAmount, data.ActualPrice, data.DiscountPrice, data.OrderStatus, data.RazorPayOrderId)
	if err != nil {
		return err
	}
	fmt.Println(orderIdx)
	return nil
}

func (p *PaymentUsecase) GetWallet(c context.Context, userid uint) (response.WalletRes, error) {

	wallet, err := p.paymentRepository.GetWallet(c, userid)
	if err != nil {
		return wallet, err
	}
	return wallet, nil

}
func (p *PaymentUsecase) AddPaymentId(paymentID string, RazorPayOrderId string) error {
	err := p.paymentRepository.AddPaymentId(paymentID, RazorPayOrderId)
	if err != nil {
		return err
	}
	return nil
}
