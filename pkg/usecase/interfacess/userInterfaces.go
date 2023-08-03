package interfacess

import (
	"context"

	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/munsheerck79/Ecom_project.git/util/response"
)

type UserService interface {
	SignUp(ctx context.Context, user domain.Users) error
	Login(ctx context.Context, user domain.Users) (domain.Users, error)
	OTPLogin(ctx context.Context, user domain.Users) (domain.Users, error)

	GetAddress(ctx context.Context, userId uint) (domain.Address, error)
	ListCart(ctx context.Context, userId uint) ([]response.Cart, float32, float32, error)
	ListWishList(ctx context.Context, userId uint) ([]response.ProductDetails, error)

	AddAddress(ctx context.Context, address domain.Address) error
	EditAddress(ctx context.Context, address domain.Address) error
	AddToCart(ctx context.Context, item request.Cart) error
	EditCartProduct(ctx context.Context, item domain.Cart) error


	DeleteFromCart(c context.Context, cartId uint, userId uint) error

	AddToWishList(ctx context.Context, item domain.WishList) error
	DeleteFromWishLIst(c context.Context, Id uint, userId uint) error
	OrderHistory(ctx context.Context, userId uint,page request.ReqPagination) ([]response.Order, error)

	//GetWallet(ctx context.Context, userId uint) (domain.Wallet, error)
}
