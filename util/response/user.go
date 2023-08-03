package response

import "time"

type Cart struct {
	ID            uint    `json:"id"`
	ProductId     uint    `json:"product_id"`
	ProductName   string  `json:"product_name"`
	VariationName string  `json:"Variation_name"`
	Quantity      int    `json:"Quantity"`
	ActualPrice   float32 `json:"actual_price"`
	DiscountPrice float32 `json:"discount_price"`
	QtyInStock    int    `json:"qty_in_stock"`
}

type Order struct {
	ID              uint      `json:"id"`
	UsersID         uint      `json:"users_id"`
	ActualPrice     float32   `json:"actual_price"`
	DiscountPrice   float32   `json:"disc_price"`
	NetAmount       float32   `json:"net_amount"`
	OrderStatusID   uint      `json:"order_status_id"`
	OrderStatus     string    `json:"status"`
	PaymentMethod   string    `json:"payment_method"`
	RazorPayOrderId string    `json:"Razor_pay_order_id"`
	OrderDate       time.Time `json:"order_date"`
	PayPaymentId    string    `json:"payment_id"`
}

type VariationR struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
type ProductRes struct {
	ID           uint   `json:"id"`
	Code         uint   `json:"code"`
	Name         string `json:"product_name"`
	Description  string `json:"description"`
	CategoryName string `json:"category_name"`
	QtyInStock   int   `json:"qty_in_stock"`
	Image        string `json:"image"`
}
type WalletRes struct {
	Balence float32    `json:"balence"`
	History []WHistoty `json:"historys" gorm:"type:jsonb"`
}
type WHistoty struct {
	PaymentId string    `json:"razorpay_payment_id"`
	Amount    float32   `json:"amount"`
	Date      time.Time `json:"added_at"`
}
