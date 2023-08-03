package interfacess

import (
	"context"

	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/munsheerck79/Ecom_project.git/util/response"
)

type ProductService interface {
	GetCategory(ctx context.Context) ([]domain.Category, error)
	Getvariations(ctx context.Context) ([]domain.Variation, error)
	AddCategory(ctx context.Context, category domain.Category) error
	AddVarient(ctx context.Context, varient domain.Variation) error
	AddProduct(ctx context.Context, product domain.Product) error
	AddPrice(ctx context.Context, price domain.Price) error
	GetProductList(ctx context.Context, page request.ReqPagination) ([]response.ProductRes, []response.VariationR, error)
	GetProductsByCategoryName(ctc context.Context, CID uint) ([]response.ProductDetails, error)

	EditProduct(ctx context.Context, product domain.Product) error

	GetOrderStatus(ctx context.Context) ([]domain.OrderStatus, error)
	AddOrderStatus(ctx context.Context, body domain.OrderStatus) error
	GetProduct(ctx context.Context, Id uint) (response.ProductDetails, error)
	//========-------======--=============---------GetProductList(ctx context.Context) (productList []domain.Users, err error)
}
