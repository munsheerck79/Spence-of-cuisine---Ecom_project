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

type orderDatabase struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepository {
	return &orderDatabase{DB: DB}
}
func (p *orderDatabase) OrderCartProducts(c context.Context, userId uint, body request.Order, NetAmount float32, total float32, discPrice float32,
) (response.Order, uint, error) {
	var order response.Order
	orderAt := time.Now()
	var orderId uint
	query := `INSERT INTO orders (users_id, payment_method, coupon_id, actual_price, discount_price, net_amount, order_date, payment_id )
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?) RETURNING ID`

	result := p.DB.Raw(query, userId, body.PaymentMethod, body.CouponId, total, discPrice, NetAmount, orderAt, body.RazorPayPaymentId).Scan(&orderId)
	if result.Error != nil {
		return order, orderId, fmt.Errorf("failed to order products: %v", result.Error)
	}

	fmt.Println("enter in err 1111")
	query2 := `SELECT id,users_id,actual_price,discount_price,net_amount,order_status_id,payment_method,order_date,payment_id
	 FROM orders WHERE id = ?`
	if err := p.DB.Raw(query2, orderId).Scan(&order).Error; err != nil {
		return order, orderId, fmt.Errorf("failed to show order products ")
	}
	return order, orderId, nil
}

func (p *orderDatabase) OrderProductsTemp(c context.Context, userId uint, body request.Order, NetAmount float32, total float32,
	discPrice float32, status string, RazorPayorderId string) (uint, error) {
	orderAt := time.Now()
	var orderId uint
	query := `INSERT INTO orders_temps (users_id, payment_method, coupon_id, actual_price, discount_price, net_amount, order_date, payment_id,razor_pay_order_id,status )
				  VALUES (?, ?, ?, ?, ?, ?, ?, ?,?,?) RETURNING ID`

	result := p.DB.Raw(query, userId, body.PaymentMethod, body.CouponId, total, discPrice, NetAmount, orderAt, body.RazorPayPaymentId, RazorPayorderId, status).Scan(&orderId)
	if result.Error != nil {
		return orderId, fmt.Errorf("failed to order products: %v", result.Error)
	}
	return orderId, nil
}
func (p *orderDatabase) AddItemsToOrderItemsTemp(c context.Context, productList response.Cart, NewOrderID uint) error {

	query3 := `INSERT INTO orders_items_temps (order_id, product_id, product_name, Variation_name, Quantity, actual_price, discount_price)
        VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Use Exec method to execute the raw SQL query
	result := p.DB.Exec(query3, NewOrderID, productList.ProductId, productList.ProductName, productList.VariationName, productList.Quantity, productList.ActualPrice, productList.DiscountPrice)

	if result.Error != nil {
		// Handle the error if any
		fmt.Println("Error executing SQL query:", result.Error)
		return result.Error
	}

	return nil
}

func (p *orderDatabase) OrderTempItems(ID uint) ([]response.Cart, error) {
	var data []response.Cart
	query := `SELECT * FROM orders_items_temps WHERE order_id = ?`
	if err := p.DB.Raw(query, ID).Scan(&data).Error; err != nil {
		return data, fmt.Errorf("failed to show order products ")
	}
	return data, nil

}

func (p *orderDatabase) FindTempDataByPaymentId(RazorPayPaymentId string) (domain.OrdersTemp, error) {
	var tempData domain.OrdersTemp

	query := `SELECT * FROM orders_temps WHERE payment_id = ?`
	if err := p.DB.Raw(query, RazorPayPaymentId).Scan(&tempData).Error; err != nil {
		return tempData, fmt.Errorf("failed to show order products ")
	}
	return tempData, nil
}

func (p *orderDatabase) AddItemsToOrderItems(c context.Context, productList response.Cart, NewOrderID uint) error {

	query3 := `INSERT INTO orders_items (orders_id, product_id, product_name, Variation_name, Quantity, actual_price, discount_price)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Use Exec method to execute the raw SQL query
	result := p.DB.Exec(query3, NewOrderID, productList.ProductId, productList.ProductName, productList.VariationName, productList.Quantity, productList.ActualPrice, productList.DiscountPrice)

	if result.Error != nil {
		// Handle the error if any
		return result.Error
	}
	//fmt.Printf("... %v....%v..%v..%v...%v.....%v....%v..", NewOrderID, productList.ProductId, productList.ProductName, productList.VariationName, productList.Quantity, productList.ActualPrice, productList.DiscountPrice)

	return nil
}
func (p *orderDatabase) GetOrderItemsById(orderId uint) ([]response.Cart, error) {

	var Orderlist []response.Cart
	query := `SELECT id,product_id,product_name,variation_name,quantity,actual_price ,discount_price 
	FROM orders_items
	WHERE orders_items.orders_id = ?`

	if err := p.DB.Raw(query, orderId).Scan(&Orderlist).Error; err != nil {
		return Orderlist, errors.New("failed to get user")
	}
	return Orderlist, nil
}

func (p *orderDatabase) OrderDetails(c context.Context, userId uint, ID uint) (response.Order, error) {

	var orderDetails response.Order
	query2 := `SELECT orders.id,orders.users_id,orders.actual_price,orders.discount_price,orders.net_amount,
	order_statuses.status AS order_status,orders.payment_method,orders.order_date,orders.payment_id
	FROM orders 
	LEFT JOIN order_statuses ON orders.order_status_id = order_statuses.id
	WHERE orders.Id = ?`

	if err := p.DB.Raw(query2, ID).Scan(&orderDetails).Error; err != nil {
		return orderDetails, fmt.Errorf("failed to show order products ")
	}
	return orderDetails, nil
}

func (p *orderDatabase) GetOrderStatusId(c context.Context, body request.UpDateOrderStatus) (request.UpDateOrderStatus, error) {
	query2 := `SELECT id FROM order_statuses WHERE status = ?`
	if err := p.DB.Raw(query2, body.Status).Scan(&body.OrderStatusID).Error; err != nil {
		return body, fmt.Errorf("failed to get order status id ")
	}
	return body, nil
}
func (p *orderDatabase) UpDateOrderStatus(c context.Context, orderstatus request.UpDateOrderStatus) error {

	query := `UPDATE orders SET order_status_id = $1 WHERE id = $2 AND users_id =$3`

	err := p.DB.Exec(query, orderstatus.OrderStatusID, orderstatus.OrderID, orderstatus.UsersID).Error
	if err != nil {
		return fmt.Errorf("failed to save ordercstatus id")
	}
	return nil

}
