package interfaces

import (
	"context"

	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/munsheerck79/Ecom_project.git/util/response"
)

type ProductRepository interface {
	GetCategory(ctx context.Context) ([]domain.Category, error)
	FindCategory(ctx context.Context, category domain.Category) (domain.Category, error)
	SaveCategory(ctc context.Context, category domain.Category) error
	GetProduct(ctx context.Context, Id uint) (response.ProductDetails, error)

	Getvariations(ctx context.Context) ([]domain.Variation, error)
	FindVarient(ctx context.Context, varient domain.Variation) (domain.Variation, error)
	SaveVarient(ctc context.Context, varient domain.Variation) error

	FindProduct(ctx context.Context, product domain.Product) (domain.Product, error)
	SaveProduct(ctc context.Context, product domain.Product) error
	EditProduct(ctx context.Context, product domain.Product) error

	AddPrice(ctx context.Context, price domain.Price) (domain.Price, error)
	SaveProductPrice(ctx context.Context, price domain.Price) error
	UpdateProductPrice(ctx context.Context, price domain.Price) error

	GetProductList(ctx context.Context, page request.ReqPagination) ([]response.ProductDetails, error)

	GetProductsByCategoryName(ctc context.Context, CID uint) ([]response.ProductDetails, error)

	GetOrderStatus(ctx context.Context) ([]domain.OrderStatus, error)
	FindOrderStatus(ctx context.Context, body domain.OrderStatus) (domain.OrderStatus, error)
	SaveOrderStatus(ctc context.Context, body domain.OrderStatus) error

	UpdateProductStock(ctx context.Context, p uint, qty int) error
}
