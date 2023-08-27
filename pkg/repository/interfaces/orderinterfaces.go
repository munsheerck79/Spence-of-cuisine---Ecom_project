package interfaces

import (
	"context"

	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/munsheerck79/Ecom_project.git/util/response"
)

type OrderRepository interface {
	OrderCartProducts(c context.Context, userId uint, body request.Order, NetAmount float32, total float32, discPrice float32) (response.Order, uint, error)
	OrderDetails(c context.Context, userId uint, ID uint) (response.Order, error)
	GetOrderStatusId(c context.Context, body request.UpDateOrderStatus) (request.UpDateOrderStatus, error)
	UpDateOrderStatus(c context.Context, orderStatus request.UpDateOrderStatus) error

	AddItemsToOrderItems(c context.Context, productList response.Cart, NewOrderID uint) error
	GetOrderItemsById(orderId uint) ([]response.Cart, error)
	
	OrderProductsTemp(c context.Context, userId uint, body request.Order, NetAmount float32,
		total float32, discPrice float32, status string, orderId string) (uint, error)
	AddItemsToOrderItemsTemp(c context.Context, productList response.Cart, NewOrderID uint) error

	FindTempDataByPaymentId(RazorPayPaymentId string) (domain.OrdersTemp, error)
	OrderTempItems(ID uint)([]response.Cart,error)
}
//CartListord(ctx context.Context, userId uint) ([]response.Cart, error)

	//FindOrder(c context.Context,userId uint,orderID uint)(domain.Orders,error)

	//CheckStock(ctx context.Context, order domain.Orders) (domain.Product, error)
	//PlaceOrder(ctx context.Context, order domain.Orders,product domain.Product) error