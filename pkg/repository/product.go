package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	interfaces "github.com/munsheerck79/Ecom_project.git/pkg/repository/interfaces"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/munsheerck79/Ecom_project.git/util/response"
	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB: DB}
}
func (p *productDatabase) GetProductList(ctx context.Context, page request.ReqPagination) ([]response.ProductRes, []response.VariationR, error) {
	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	var Products []response.ProductRes
	var Vari []response.VariationR
	query := `SELECT products.id,products.code,products.name,products.description,products.qty_in_stock,products.image ,categories.category_name 
	FROM products
	INNER JOIN categories ON products.category_id=categories.id
	ORDER BY products.created_at DESC LIMIT $1 OFFSET $2`
	if err := p.DB.Raw(query, limit, offset).Scan(&Products).Error; err != nil {
		return Products, Vari, errors.New("faild to show products1111")
	}
	query2 := `SELECT * FROM variations`
	if err := p.DB.Raw(query2).Scan(&Vari).Error; err != nil {
		return Products, Vari, errors.New("faild to show products")
	}

	return Products, Vari, nil

}

func (p *productDatabase) GetCategory(ctx context.Context) ([]domain.Category, error) {

	var category []domain.Category
	query := `SELECT * FROM categories`
	if err := p.DB.Raw(query).Scan(&category).Error; err != nil {
		return category, errors.New("failed to get catogory list")
	}
	fmt.Println("get catogorylist")
	return category, nil
}

func (p *productDatabase) Getvariations(ctx context.Context) ([]domain.Variation, error) {

	var variations []domain.Variation
	query := `SELECT * FROM variations`
	if err := p.DB.Raw(query).Scan(&variations).Error; err != nil {
		return variations, errors.New("failed to get variations list")
	}
	fmt.Println("get variationslist")
	return variations, nil
}

func (p *productDatabase) FindCategory(ctx context.Context, category domain.Category) (domain.Category, error) {

	fmt.Println(category.CategoryName)
	query := `SELECT * FROM categories WHERE category_name = ?`
	if err := p.DB.Raw(query, category.CategoryName).Scan(&category).Error; err != nil {
		return category, errors.New("failed to get catogory")
	}
	fmt.Println("get catogory")
	return category, nil
}

func (p *productDatabase) SaveCategory(ctc context.Context, category domain.Category) error {
	query := `INSERT INTO categories (category_name,created_at) 
			  VALUES ($1, $2)`
	createdAt := time.Now()
	// fmt.Printf("==/%v/==== %v=/ ==%v=== ", query, category.CategoryName, createdAt)
	err := p.DB.Exec(query, category.CategoryName, createdAt).Error
	if err != nil {
		return fmt.Errorf("failed to save category %s", category.CategoryName)
	}
	return nil
}

/////////////////////////////////////////////

func (p *productDatabase) FindVarient(ctx context.Context, varient domain.Variation) (domain.Variation, error) {

	fmt.Println(varient.Name)
	query := `SELECT * FROM variations WHERE name = ?`
	if err := p.DB.Raw(query, varient.Name).Scan(&varient).Error; err != nil {
		return varient, errors.New("failed to get varient")
	}
	fmt.Println("get varient")
	return varient, nil
}

func (p *productDatabase) SaveVarient(ctc context.Context, varient domain.Variation) error {
	query := `INSERT INTO variations (name,created_at) 
			  VALUES ($1, $2)`
	createdAt := time.Now()
	err := p.DB.Exec(query, varient.Name, createdAt).Error
	if err != nil {
		return fmt.Errorf("failed to save varient %s", varient.Name)
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (p *productDatabase) FindProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	fmt.Println(product.Code)
	query := `SELECT * FROM products WHERE code = ? `
	if err := p.DB.Raw(query, product.Code).Scan(&product).Error; err != nil {
		fmt.Println("fail to get")
		return product, errors.New("failed to get product")
	}
	fmt.Println("=====get product")
	return product, nil
}

func (p *productDatabase) SaveProduct(ctx context.Context, product domain.Product) error {
	query := `INSERT INTO products (code,name,description,category_id,qty_in_stock,image,created_At) 
			  VALUES ($1,$2,$3,$4,$5,$6,$7)`
	createdAt := time.Now()
	err := p.DB.Exec(query, product.Code, product.Name, product.Description, product.CategoryID, product.QtyInStock, product.Image, createdAt).Error
	if err != nil {
		return fmt.Errorf("failed to save product %s", product.Name)
	}
	return nil
}

func (p *productDatabase) EditProduct(ctx context.Context, product domain.Product) error {
	query := `UPDATE products SET name = $1, description = $2, category_id = $3,
	 qty_in_stock= $4, image = $5,stock_status = $8, updated_at = $6 WHERE code = $7`

	updatedAt := time.Now()

	if p.DB.Exec(query, product.Name, product.Description, product.CategoryID,
		product.QtyInStock, product.Image, updatedAt, product.Code, product.StockStatus).Error != nil {
		return errors.New("failed to update product")
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////

func (p *productDatabase) AddPrice(ctx context.Context, price domain.Price) (domain.Price, error) {
	query := `SELECT * FROM prices WHERE product_id = ? AND variation_id = ?`
	if err := p.DB.Raw(query, price.ProductID, price.VariationID).Scan(&price).Error; err != nil {
		return price, errors.New("failed to get product price")
	}
	fmt.Println("get product")
	return price, nil
}
func (p *productDatabase) SaveProductPrice(ctc context.Context, price domain.Price) error {
	fmt.Println("getin saveprdctprice")
	query := `INSERT INTO prices (product_id,variation_id,actual_price,discount_price,created_at) 
			  VALUES ($1,$2,$3,$4,$5)`
	createdAt := time.Now()
	err := p.DB.Exec(query, price.ProductID, price.VariationID, price.ActualPrice, price.DiscountPrice, createdAt).Error
	if err != nil {
		return fmt.Errorf("failed to save product %d", price.ProductID)
	}
	return nil
}

func (p *productDatabase) UpdateProductPrice(ctc context.Context, price domain.Price) error {

	var getid domain.Price
	query1 := `SELECT * FROM prices WHERE product_id = ? AND variation_id = ?`
	if err := p.DB.Raw(query1, price.ProductID, price.VariationID).Scan(&getid).Error; err != nil {
		return errors.New("failed to get product price")
	}
	query := `UPDATE prices SET actual_price = $3,discount_price = $4,updated_at = $2 WHERE id = $1 ;`
	createdAt := time.Now()
	err := p.DB.Exec(query, getid.ID, createdAt, price.ActualPrice, price.DiscountPrice).Error
	if err != nil {
		return fmt.Errorf("failed to save price %d", price.ProductID)
	}
	return nil
}

func (p *productDatabase) UpdateProductStock(ctx context.Context, id uint, qty int) error {
	query := `UPDATE products SET qty_in_stock = $1,updated_at = $2 WHERE id = $3`
	createdAt := time.Now()
	err := p.DB.Exec(query, qty, createdAt, id).Error
	if err != nil {
		return fmt.Errorf("failed to save price %d", id)
	}
	return nil

}

func (p *productDatabase) GetOrderStatus(ctx context.Context) ([]domain.OrderStatus, error) {

	var status []domain.OrderStatus
	query := `SELECT * FROM order_statuses`
	if err := p.DB.Raw(query).Scan(&status).Error; err != nil {
		return status, errors.New("failed to get variations list")
	}
	fmt.Println("get variationslist")
	return status, nil
}

func (p *productDatabase) FindOrderStatus(ctx context.Context, status domain.OrderStatus) (domain.OrderStatus, error) {

	var dbstatus domain.OrderStatus
	query := `SELECT * FROM order_statuses WHERE status = ?`
	if err := p.DB.Raw(query, status.Status).Scan(&dbstatus).Error; err != nil {
		return dbstatus, errors.New("failed to get status")
	}
	fmt.Println("get catogory")
	return dbstatus, nil
}

func (p *productDatabase) SaveOrderStatus(ctc context.Context, status domain.OrderStatus) error {
	query := `INSERT INTO order_statuses (status) 
			  VALUES ($1)`
	err := p.DB.Exec(query, status.Status).Error
	if err != nil {
		return fmt.Errorf("failed to save status %s", status.Status)
	}
	return nil
}

func (p *productDatabase) GetProduct(ctx context.Context, Id uint) (response.ProductDetails, error) {

	var product response.ProductDetails
	query := `SELECT products.id,products.code,products.name,products.description,products.qty_in_stock,products.image ,categories.category_name 
	FROM products
	LEFT JOIN categories ON products.category_id=categories.id
	WHERE Products.id =? `
	fmt.Println("db", Id)

	if err := p.DB.Raw(query, Id).Scan(&product).Error; err != nil {
		return product, errors.New("faild to show products")
	}
	query2 := `SELECT variations.name,prices.actual_price,prices.discount_price 
	FROM prices
	INNER JOIN variations ON variations.Id = prices.variation_id
	INNER JOIN products ON products.Id = prices.product_id
	WHERE  prices.product_id =? `
	if err := p.DB.Raw(query2, Id).Scan(&product.PriceList).Error; err != nil {
		return product, errors.New("faild to show products")
	}

	return product, nil
}

func (p *productDatabase) GetProductsByCategoryName(ctc context.Context, CID uint) ([]response.ProductDetails, error) {

	var Products []response.ProductDetails
	query := `SELECT products.id AS id,products.code,products.name AS product_name,products.description,products.qty_in_stock,products.image,category_Name
	FROM products 
	INNER JOIN categories ON products.category_id =categories.id 
	WHERE products.category_id = ? `
	if err := p.DB.Raw(query, CID).Scan(&Products).Error; err != nil {
		return Products, errors.New("faild to show products")
	}

	for k := 0; k < len(Products); k++ {
		query2 := `SELECT variations.name,prices.actual_price,prices.discount_price 
		FROM prices
		INNER JOIN variations ON variations.Id = prices.variation_id
		INNER JOIN products ON products.Id = prices.product_id
		WHERE  prices.product_id =? `
		if err := p.DB.Raw(query2, Products[k].ID).Scan(&Products[k].PriceList).Error; err != nil {
			return Products, errors.New("faild to show products")
		}
	}

	return Products, nil
}
