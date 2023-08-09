package response



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