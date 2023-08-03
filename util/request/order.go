package request

type Order struct {
	PaymentMethod     string `json:"payment_method"`
	CouponCode        string `json:"coupon_code,omitempty"`
	CouponId          uint   `json:"coupon_id,omitempty"`
	RazorPayPaymentId string `json:"razorpay_payment_id,omitempty"`
}

type UpDateOrderStatus struct {
	OrderID       uint   `json:"order_id" binding:"required"`
	UsersID       uint   `json:"users_id" binding:"required"`
	OrderStatusID uint   `json:"order_status_id,omitempty"`
	Status        string `json:"status" binding:"required"`
}
type RazorPayCheckoutWallet struct {
	Amount float32 `json:"amount" binding:"required"`
}
