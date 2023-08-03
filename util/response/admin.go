package response

import "time"

type AdminOrderList struct {
	ID            uint      `json:"id"`
	UsersID       uint      `json:"users_id"`
	ActualPrice   float32   `json:"actual_price"`
	DiscountPrice float32   `json:"disc_price"`
	NetAmount     float32   `json:"net_amount"`
	OrderStatusID uint      `json:"order_status_id"`
	PaymentMethod string    `json:"payment_method"`
	OrderDate     time.Time `json:"order_date"`
}

type UserDetails struct {
	ID          uint      `json:"id"`
	UserName    string    `json:"user_name"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Age         uint      `json:"age"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	BlockStatus bool      `json:"block_status"`
	Address     string    `json:"address"`
	Muncipality string    `json:"muncipality"`
	District    string    `json:"district"`
	OrdersList  []Orders1 `json:"Orders_list" gorm:"type:jsonb"`
}

type Orders1 struct {
	ID              uint      `json:"id"`
	Items           []Cart    `json:"items" gorm:"type:jsonb"` // Use jsonb type for JSON data
	CouponID        string    `json:"coupon_code"`
	ActualPrice     float32   `json:"actual_price"`
	DiscountPrice   float32   `json:"disc_price"`
	NetAmount       float32   `json:"net_amount"`
	OrderStatusName string    `json:"order_status_name"`
	PaymentMethod   string    `json:"payment_method"`
	OrderDate       time.Time `json:"order_date"`
}

type SalesReport struct {
	ID            uint      `json:"order_id"`
	UserName      string    `json:"user_name"`
	Name          string    `json:"name"`
	Total_amound  float32   `json:"total_amount"`
	Discount      float32   `json:"discount"`
	NetAmount     float32   `json:"net_amount"`
	CouponCode    string    `json:"coupon_code"`
	OrderStatus   string    `json:"order_status"`
	PaymentMethod string    `json:"payment_method"`
	OrderDate     time.Time `json:"order_date"`
}

type Prices struct {
	Name          string  `json:"Variation_Name"`
	ActualPrice   float32 `json:"price" gorm:"not null"`
	DiscountPrice float32 `json:"discount_price" gorm:"default:null"`
}

type ProductDetails struct {
	ID           uint     `json:"id"`
	Code         uint     `json:"code"`
	Name         string   `json:"product_name"`
	Description  string   `json:"description"`
	CategoryName string   `json:"category_Name"`
	QtyInStock   int     `json:"qty_in_stock"`
	StockStatus  bool     `json:"stock_status"`
	Image        string   `json:"image"`
	PriceList    []Prices `json:"Price_list" gorm:"type:jsonb"`
}
