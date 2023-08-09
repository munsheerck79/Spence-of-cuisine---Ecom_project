package interfaces

import (
	"context"

	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/munsheerck79/Ecom_project.git/util/response"
)

type AdminRepository interface {
	FindAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error)
	GetUserlist(ctx context.Context, page request.ReqPagination) (userList []domain.Users, err error)
	BlockUser(ctx context.Context, body request.BlockUser) (string, error)
	GetOrderlist(ctx context.Context,page request.ReqPagination) (orderList []response.AdminOrderList, err error)
	CancelOrder(ctx context.Context, body request.CancelOrder) (string, error)
	UserDetails(ctx context.Context,body request.UserDetails)(response.UserDetails,error)
	GenerateSalesReport(c context.Context, dateRange request.DateRange) (salesData []response.SalesReport, err error)
}
