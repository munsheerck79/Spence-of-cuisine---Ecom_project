package domain

import (
	"time"
)

// User model
type Users struct {
	ID          uint      `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	UserName    string    `json:"user_name" gorm:"not null" binding:"required,min=3,max=15"`
	FirstName   string    `json:"first_name" gorm:"not null" binding:"required,min=2,max=40"`
	LastName    string    `json:"last_name" gorm:"not null" binding:"required,min=1,max=40"`
	Age         uint      `json:"age" gorm:"not null" binding:"required,numeric"`
	Email       string    `json:"email" gorm:"unique;not null" binding:"required,email"`
	Phone       string    `json:"phone" gorm:"not null" binding:"required,min=10,max=10"`
	Password    string    `json:"password" gorm:"not null" binding:"required"`
	BlockStatus bool      `json:"block_status" gorm:"not null;default:false"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Address struct {
	ID          uint      `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	UsersID     uint      `json:"users_id" gorm:"not null"`
	Users       Users     `json:"-" gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Address     string    `json:"address" gorm:"not null"`
	Muncipality string    `json:"muncipality" gorm:"not null"`
	LandMark    string    `json:"land_mark" gorm:"not null"`
	District    string    `json:"district" gorm:"not null"`
	State       string    `json:"state" gorm:"not null"`
	PhoneNumber string    `json:"phone_number" gorm:"not null"`
	PinCode     string    `json:"pin_code" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Cart struct {
	ID          uint      `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	UsersID     uint      `json:"users_id" gorm:"not null"`
	Users       Users     `json:"-" gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProductID   uint      `json:"product_id" gorm:"not null"`
	Product     Product   `json:"-" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	VariationID uint      `json:"Variation_id" gorm:"not null"`
	Variation   Variation `json:"-" gorm:"foreignKey:VariationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Quantity    int      `json:"Quantity" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WishList struct {
	ID        uint      `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	UsersID   uint      `json:"users_id" gorm:"not null"`
	Users     Users     `json:"-" gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	Product   Product   `json:"-" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
}

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
	Quantity      int    `json:"Quantity" gorm:"not null"`
	ActualPrice   float32 `json:"actual_price" gorm:"not null"`
	DiscountPrice float32 `json:"discount_price" gorm:"not null"`
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
type OrderStatus struct {
	Id     uint   `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	Status string `json:"status" gorm:"not null" binding:"required,min=3,max=50"`
}

type Wallet struct {
	ID        uint      `gorm:"primaryKey;unique;autoIncrement" json:"id"`
	UsersID   uint      `json:"users_id" gorm:"not null"`
	Users     Users     `json:"-" gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Balence   float32   `json:"amount" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at"`
}
type WalletHistory struct {
	ID        uint      `gorm:"primaryKey;unique;autoIncrement" json:"id"`
	UsersID   uint      `json:"users_id" gorm:"not null"`
	Users     Users     `json:"-" gorm:"foreignKey:UsersID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	PaymentId string    `json:"razorpay_payment_id,omitempty"`
	Amount    float32   `json:"amount" gorm:"not null"`
	Date      time.Time `json:"added_at" gorm:"not null"`
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
