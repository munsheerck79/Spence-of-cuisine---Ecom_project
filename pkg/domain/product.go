package domain

import (
	"time"
)

// catogory struct
type Category struct {
	ID           uint      `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	CategoryName string    `json:"category_name" gorm:"unique;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:null"`
}

type Variation struct {
	ID uint `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	//VariationIndex uint      `json:"variation_index" gorm:"unique;not null"`
	Name      string    `json:"name" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:null"`
}

// Product struct
type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	Code        uint      `json:"code" gorm:"unique;not null"`
	Name        string    `json:"product_name" gorm:"not null;size:50"`
	Description string    `json:"description" gorm:"not null;size:500"`
	CategoryID  uint      `json:"category_id" gorm:"not null"`
	Category    Category  `json:"-" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	QtyInStock  int      `json:"qty_in_stock" gorm:"not null"`
	StockStatus bool      `json:"stock_status" gorm:"not null;default:true;type:boolean"`
	Image       string    `json:"image" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:null"`
}
type Price struct {
	ID            uint      `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	ProductID     uint      `json:"product_id" gorm:"not null"`
	Product       Product   `json:"-" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	VariationID   uint      `json:"Variation_id" gorm:"not null"`
	Variation     Variation `json:"-" gorm:"foreignKey:VariationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ActualPrice   float32   `json:"price" gorm:"not null"`
	DiscountPrice float32   `json:"discount_price" gorm:"default:null"`
	CreatedAt     time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"default:null"`
}


