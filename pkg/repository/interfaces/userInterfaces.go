package interfaces

import (
	"context"

	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/munsheerck79/Ecom_project.git/util/response"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user domain.Users) error
	FindUser(ctx context.Context, user domain.Users) (domain.Users, error)
	FindAddress(ctx context.Context, id uint) (domain.Address, error)
	SaveAddress(ctx context.Context, address domain.Address) error
	UpdateAddress(ctx context.Context, address domain.Address) error
	AddToCart(ctx context.Context, item request.Cart) error
	AddToWishList(ctx context.Context, item domain.WishList) error
	CartList(ctx context.Context, userId uint) ([]response.Cart, error)
	ListWishList(ctx context.Context, userId uint) ([]response.ProductDetails, error)
	OrderHistory(ctx context.Context, userId uint,page request.ReqPagination) ([]response.Order, error)
	DeleteFromCart(c context.Context, cartId uint, userId uint) error
	DeleteFromWishLIst(c context.Context, Id uint, userId uint) error
	GetWalletx(ctx context.Context, userId uint) (domain.Wallet, error)
	EditCartProduct(ctx context.Context, item domain.Cart) error	
}
