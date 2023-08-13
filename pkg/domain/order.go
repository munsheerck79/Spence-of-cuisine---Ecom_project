package domain

import "time"

type Orders struct {
	ID      uint  `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	UsersID uint  `json:"users_id" gorm:"not null"`
	Users   Users `json:"-" gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	//Items         []response.Cart `json:"-" gorm:"not null;type:jsonb"` // Use jsonb type for JSON data
	CouponID      uint        `json:"coupon_id,omitempty"`
	Coupon        Coupon      `json:"-" gorm:"foreignKey:CouponID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ActualPrice   float32     `json:"actual_price" gorm:"not null"`
	DiscountPrice float32     `json:"disc_price" gorm:"not null"`
	NetAmount     float32     `json:"net_amount" gorm:"not null"`
	OrderStatusID uint        `json:"order_status_id"`
	OrderStatus   OrderStatus `gorm:"foreignKey:OrderStatusID" json:"-"`
	PaymentMethod string      `json:"payment_method" gorm:"not null"`
	PaymentId     string      `json:"razorpay_payment_id,omitempty"`
	OrderDate     time.Time   `json:"order_date" gorm:"not null"`
}
type OrdersItems struct {
	ID            uint    `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	OrdersID      uint    `json:"Orders_id" gorm:"not null"`
	Orders        Orders  `json:"-" gorm:"foreignKey:OrdersID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProductId     uint    `json:"product_id" gorm:"not null"`
	ProductName   string  `json:"product_name" gorm:"not null"`
	VariationName string  `json:"Variation_name" gorm:"not null"`
	Quantity      int     `json:"Quantity" gorm:"not null"`
	ActualPrice   float32 `json:"actual_price" gorm:"not null"`
	DiscountPrice float32 `json:"discount_price" gorm:"not null"`
}
type OrderStatus struct {
	Id     uint   `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	Status string `json:"status" gorm:"not null" binding:"required,min=3,max=50"`
}

type OrdersTemp struct {
	ID      uint  `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	UsersID uint  `json:"users_id" gorm:"not null"`
	Users   Users `json:"-" gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	//Items         []response.Cart `json:"-" gorm:"not null;type:jsonb"` // Use jsonb type for JSON data
	CouponID        uint      `json:"coupon_id,omitempty"`
	Coupon          Coupon    `json:"-" gorm:"foreignKey:CouponID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ActualPrice     float32   `json:"actual_price" gorm:"not null"`
	DiscountPrice   float32   `json:"disc_price" gorm:"not null"`
	NetAmount       float32   `json:"net_amount" gorm:"not null"`
	Status          string    `json:"status"`
	PaymentMethod   string    `json:"payment_method" gorm:"not null"`
	RazorPayOrderId string    `json:"razorpay_Order_id" gorm:"not null"`
	PaymentId       string    `json:"razorpay_payment_id,omitempty"`
	OrderDate       time.Time `json:"order_date" gorm:"not null"`
}

type OrdersItemsTemp struct {
	ID            uint       `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	OrderID       uint       `json:"Orders_id" gorm:"not null"`
	Order         OrdersTemp `json:"-" gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProductId     uint       `json:"product_id" gorm:"not null"`
	ProductName   string     `json:"product_name" gorm:"not null"`
	VariationName string     `json:"Variation_name" gorm:"not null"`
	Quantity      int        `json:"Quantity" gorm:"not null"`
	ActualPrice   float32    `json:"actual_price" gorm:"not null"`
	DiscountPrice float32    `json:"discount_price" gorm:"not null"`
}
type Coupon struct {
	ID                uint    `gorm:"primaryKey;unique;autoIncrement" json:"id"`
	Code              string  `json:"code," gorm:"unique;not null" binding:"required,min=4,max=8"`
	MinOrderValue     float64 `json:"min_order_value" binding:"required"`
	DiscountPercent   int     `json:"discount_percent" binding:"required"`
	DiscountMaxAmount float64 `json:"discount_max_amount" binding:"required"`
	Description       string  `json:"description" gorm:"not null;size:500"`
	ValidTill         int64   `json:"valid_days" binding:"required"`
}
