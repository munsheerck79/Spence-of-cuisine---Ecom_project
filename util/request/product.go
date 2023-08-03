package request

type AddCatogory struct {
	CategoryName string `json:"category_name"  binding:"required,min=3,max=15"`
}

type AddVarient struct {
	Name string `json:"name" binding:"required,min=3,max=15"`
}

type AddProduct struct {
	Code        uint   `json:"Code" binding:"required"`
	Name        string `json:"product_name" binding:"required,min=3,max=15"`
	Description string `json:"description" binding:"required,min=2,max=400"`
	CategoryID  uint   `json:"category_id" binding:"required"`
	QtyInStock  int   `json:"qty_in_stock" binding:"required"`
	Image       string `json:"image" binding:"required"`
}

type AddPrice struct {
	ProductID     uint    `json:"product_ID" binding:"required"`
	VariationID   uint    `json:"variation_id" binding:"required"`
	ActualPrice   float32 `json:"price" binding:"required"`
	DiscountPrice float32 `json:"discount_price" binding:"required"`
}
type EditPrice struct {
	ProductID     uint    `json:"product_ID" binding:"required"`
	VariationID   uint    `json:"variation_id" binding:"required"`
	ActualPrice   float32 `json:"price" binding:"required"`
	DiscountPrice float32 `json:"discount_price" binding:"required"`
}

type EditProduct struct {
	Code        uint   `json:"Code" binding:"required"`
	Name        string `json:"product_name,omitempty"`
	Description string `json:"description,omitempty"`
	CategoryID  uint   `json:"category_id,omitempty"`
	QtyInStock  int   `json:"qty_in_stock,omitempty"`
	StockStatus bool   `json:"stock_status" gorm:"default:true;type:boolean"`
	Image       string `json:"image,omitempty"`
}

type EditCoupon struct {
    ID                uint    `json:"id"`
    Code              string  `json:"code,omitempty"`
    MinOrderValue     float64 `json:"min_order_value,omitempty"`
    DiscountPercent   int     `json:"discount_percent,omitempty"`
    DiscountMaxAmount float64 `json:"discount_max_amount,omitempty"`
    Description       string  `json:"description,omitempty"`
    ValidTill         int64   `json:"valid_days,omitempty"`
}
