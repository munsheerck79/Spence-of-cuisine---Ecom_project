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
	Quantity    int       `json:"Quantity" gorm:"not null"`
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
