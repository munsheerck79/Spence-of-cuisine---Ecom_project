package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	interfaces "github.com/munsheerck79/Ecom_project.git/pkg/repository/interfaces"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/munsheerck79/Ecom_project.git/util/response"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB: DB}
}

func (i *adminDatabase) FindAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error) {
	// Check any of the user details matching with db user list

	query := `SELECT * FROM Admins WHERE id = ? OR email = ? OR phone = ? OR user_name = ?`
	if err := i.DB.Raw(query, admin.ID, admin.Email, admin.Phone, admin.UserName).Scan(&admin).Error; err != nil {
		return admin, errors.New("failed to get user")
	}

	fmt.Println("get user")

	fmt.Println("from DB", admin)
	return admin, nil
}

func (u *adminDatabase) GetUserlist(ctx context.Context, page request.ReqPagination) (userList []domain.Users, err error) {

	limit := page.Count
	offset := (page.PageNumber - 1) * limit

	query := `SELECT * FROM users
	ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	if err := u.DB.Raw(query, limit, offset).Scan(&userList).Error; err != nil {
		return userList, errors.New("failed to get user")
	}
	return userList, nil
}

func (a *adminDatabase) BlockUser(ctx context.Context, body request.BlockUser) (string, error) {
	var user domain.Users
	query := `SELECT * FROM users WHERE id=? OR user_name=?`
	a.DB.Raw(query, body.UserID, body.UserName).Scan(&user)
	if user.Email == "" { // check user email with user ID
		msg := "user not exist"
		return msg, nil
	}
	query = `UPDATE users SET block_status = $1 WHERE id = $2`
	if a.DB.Exec(query, !user.BlockStatus, body.UserID).Error != nil {
		msg := "failed to update user block_status to %"
		return msg, nil
	}
	msg := "succesfully blocked or unblocked"
	return msg, nil
}

func (a *adminDatabase) UserDetails(ctx context.Context, body request.UserDetails) (response.UserDetails, error) {

	var UserDetails response.UserDetails
	query := `SELECT 
	users.ID,users.user_name,user.first_name,user.last_name,age,email,phone,block_status,address,muncipality,district
	FROM users
	INNER JOIN addresses ON users.id = addresses.users_id
	WHERE id = ? AND user_name = ? ;`
	err := a.DB.Raw(query, body.UserID, body.UserName).Scan(&UserDetails)

	query2 := `SELECT id,items,coupon_code,actual_price,disc_price,net_amount,order_statuses.status,payment_method,order_date
	FROM orders
	INNER JOIN order_statuses ON orders.order_status_id = order_statuses.id ;`
	//WHERE user_id = ? AND user_name = ?
	err2 := a.DB.Raw(query2).Scan(&UserDetails.OrdersList)
	if err != nil {
		return UserDetails, errors.New("failed to get user details")
	}
	if err2 != nil {
		return UserDetails, errors.New("failed to get user order details")
	}
	return UserDetails, nil

}

func (u *adminDatabase) GetOrderlist(ctx context.Context, page request.ReqPagination) (orderList []response.AdminOrderList, err error) {

	limit := page.Count
	offset := (page.PageNumber - 1) * limit
	query := `SELECT  order_statuses.status 
	FROM orders
	LEFT JOIN order_statuses ON orders.order_status_id = order_statuses.id
	ORDER BY orders.order_date DESC LIMIT $1 OFFSET $2`

	if err := u.DB.Raw(query, limit, offset).Scan(&orderList).Error; err != nil {
		return orderList, errors.New("failed to get order list")
	}
	return orderList, nil
}

func (a *adminDatabase) GenerateSalesReport(c context.Context, dateRange request.DateRange) (salesData []response.SalesReport, err error) {

	// query := `SELECT orders.id AS Order_ID,users.user_name,users.first_name AS name,order.Actual_price As total_amound,orders.discount_price AS discount,
	// orders.net_amount,coupon.coupon_code,order_statuses.name AS order_status,orders.payment,orders.order_date
	// FROM orders
	// LEFT JOIN users ON users.Id = orders.users_id
	// LEFT JOIN coupons ON coupon.Id = orders.coupon_ID
	// LEFT JOIN order_statuses ON order_statuses.id = orders.order_status_id
	// WHERE orders.order_date BETWEEN $1 AND $2 `

	query := `SELECT
    orders.id,
    users.user_name,
    users.first_name AS name,
    orders.Actual_price AS total_amount,
    orders.discount_price AS discount,
    orders.net_amount,
    coupons.code AS coupon_code,
    order_statuses.status AS order_status,
    orders.payment_method,
    orders.order_date
FROM
    orders
LEFT JOIN
    users ON users.Id = orders.users_id
LEFT JOIN
    coupons ON coupons.Id = orders.coupon_ID
LEFT JOIN
    order_statuses ON order_statuses.id = orders.order_status_id
WHERE
    orders.order_date BETWEEN $1 AND $2`

	if err := a.DB.Raw(query, dateRange.StartDate, dateRange.EndDate).Scan(&salesData).Error; err != nil {
		return salesData, errors.New("failed to get sales list")
	}
	return salesData, nil
}

func (a *adminDatabase) CancelOrder(ctx context.Context, body request.CancelOrder) (string, error) {
	var order domain.Orders
	query := `SELECT * FROM orders WHERE id=? AND users_id=?`
	a.DB.Raw(query, body.ID, body.UsersID).Scan(&order)
	if order.ID == 0 {
		msg := "user/order not exist"
		return msg, errors.New("user/order not exist")
	}
	query1 := `DELETE FROM orders_items WHERE orders_id = $1`
	if a.DB.Exec(query1, body.ID).Error != nil {
		msg := "failed to delete order items"
		return msg, errors.New("failed to delete order")
	}

	query2 := `DELETE FROM orders WHERE id = $1 AND users_id = $2`
	if a.DB.Exec(query2, body.ID, body.UsersID).Error != nil {
		msg := "failed to delete order"
		return msg, errors.New("failed to delete order")
	}
	msg := "succesfully canceled"
	return msg, nil
}
