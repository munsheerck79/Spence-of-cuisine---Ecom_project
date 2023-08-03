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
	fmt.Println("status use", Data)
	return Data, nil
}

func (p *PaymentUsecase) AddTempData(c context.Context, data response.Order) error {

	userId := data.UsersID
	var b request.Order
	b.PaymentMethod = "wallet"
	b.CouponId = 6
	orderIdx, err := p.orderRepository.OrderProductsTemp(c, userId, b, data.NetAmount, data.ActualPrice, data.DiscountPrice, data.OrderStatus, data.RazorPayOrderId)
	if err != nil {
		return err
	}
	fmt.Println(orderIdx)
	return nil
}

// func VerifyRazorPayPayment1(signature, orderId, paymentId string) error {
// 	// Get razor pay api config
// 	fmt.Println("96655449")

// 	razorPayKey := config.GetConfig().RAZORPAYKEY
// 	razorPaySecret := config.GetConfig().RAZORPAYSECRET

// 	// Verify signature
// 	data := orderId + "|" + paymentId
// 	h := hmac.New(sha256.New, []byte(razorPaySecret))
// 	_, err := h.Write([]byte(data))
// 	if err != nil {
// 		return err
// 	}
// 	sha := hex.EncodeToString(h.Sum(nil))
// 	if subtle.ConstantTimeCompare([]byte(sha), []byte(signature)) != 1 {
// 		return err
// 	}
// 	// verify payment
// 	razorpayClient := razorpay.NewClient(razorPayKey, razorPaySecret)

// 	// fetch payment and verify
// 	payment, err := razorpayClient.Payment.Fetch(paymentId, nil, nil)
// 	if err != nil {
// 		return err
// 	}
////////////////////////////////////////////////////////////////////////////////////
//////amount chekkingggg

// 	// check payment status
// 	if payment["status"] != "captured" {
// 		return fmt.Errorf("failed to verify payment \n razor pay payment with payment_id %v", paymentId)
// 	}
// 	return nil
// }

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
