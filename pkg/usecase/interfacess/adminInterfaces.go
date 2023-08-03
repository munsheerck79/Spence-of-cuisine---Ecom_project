package interfacess

import (
	"context"

	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/munsheerck79/Ecom_project.git/util/response"
)

type AdminService interface {
	//SignUp(ctx context.Context, user domain.Users) error
	LoginAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error)
	OTPLogin(ctx context.Context, admin domain.Admin) (domain.Admin, error)
	GetUserlist(ctx context.Context) (userList []domain.Users, err error)
	BlockUser(ctx context.Context, body request.BlockUser) (string, error)
	GetOrderlist(ctx context.Context) (orderList []response.AdminOrderList, err error)
	CancelOrder(ctx context.Context, body request.CancelOrder) (string, error)
	UserDetails(ctx context.Context,body request.UserDetails)(response.UserDetails,error)	
	SalesReport(c context.Context, daterange request.DateRange) (salesReport []response.SalesReport, err error)

}
