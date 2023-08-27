package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/munsheerck79/Ecom_project.git/pkg/api/auth"
	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/pkg/usecase/interfacess"
	"github.com/munsheerck79/Ecom_project.git/util/request"
)

type OrderHandler struct {
	orderService interfacess.OrderService
	userService  interfacess.UserService
}

func NewOrderHandler(orderUsecase interfacess.OrderService, userSer interfacess.UserService) *OrderHandler {
	return &OrderHandler{orderService: orderUsecase, userService: userSer}
}

//====================================================== check and confirm all details ===================================================

// BuyProduct godoc
// @summary API for checking all details for buying cart products
// @security ApiKeyAuth
// @id BuyProduct
// @tags User.Order
// @Router /user/order/confirm-details [get]
// @Produce json
// @Success 200 {object} CheckoutResponse
// @Failure 400 "Missing or invalid entry"
func (p *OrderHandler) BuyProduct(c *gin.Context) {

	userId := auth.GetUserIdFromContext(c)
	cart, total, discprice, err := p.userService.ListCart(c, userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	address, err := p.orderService.BuyProduct(c, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if address.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please enter the user address"})
		return

	}
	if total == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart is empty"})
		return
	}

	// success response
	response := gin.H{
		"address":       address,
		"":              "products in cart",
		"data":          cart,
		"totalprice":    total,
		"you saved ":    total - discprice,
		"discountprice": discprice,
	}
	c.JSON(http.StatusOK, response)

}

type CheckoutResponse struct {
	Address domain.Address `json:"address"`
	Cart    domain.Cart    `json:"cart"`
}

//====================================================== order cart products ===================================================

// OrderCartProducts godoc
// @summary API for order cart products
// @description place order
// @security ApiKeyAuth
// @id OrderCartProducts
// @tags NotUse
// @Param coupon query string false "coupon"
// @Param payMethod query string false "pay method"
// @Router /user/order/checkout [post]
// @Success 200 {object} response.Order "html added"
// @Failure 406 "string "Invalid input"
// @Failure 400 "somthing wrong!!"
func (p *OrderHandler) OrderCartProducts(c *gin.Context) {

	var body request.Order

	body.CouponCode = c.Query("coupon")
	body.PaymentMethod = c.Query("payMethod")
	userId := auth.GetUserIdFromContext(c)
	order, _, err := p.orderService.OrderCartProducts(c, userId, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	response := gin.H{
		"":     "order details",
		"data": order,
	}
	c.JSON(http.StatusOK, response)

}

//====================================================== order cart products by cod ===================================================

// OrderCartProductsByCOD godoc
// @summary API for order cart products using cod
// @description place order by cod method
// @security ApiKeyAuth
// @id OrderCartProductsByCOD
// @tags User.Order
// @Param coupon query string false "coupon"
// @Router /user/order/checkout/cod [post]
// @Success 200 {object} response.Order "html added"
// @Failure 406 "string "Invalid input"
// @Failure 400 "somthing wrong!!"
func (p *OrderHandler) OrderCartProductsByCOD(c *gin.Context) {

	var body request.Order
	body.CouponCode = c.Query("coupon")
	body.PaymentMethod = "cod"
	userId := auth.GetUserIdFromContext(c)
	order, _, err := p.orderService.OrderCartProducts(c, userId, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	response := gin.H{
		"":     "order details",
		"data": order,
	}
	c.JSON(http.StatusOK, response)

}

//====================================================== order cart products by wallet ===================================================

// OrderCartProductsByWallet godoc
// @summary API for order cart products by using vallet
// @description place order
// @security ApiKeyAuth
// @id OrderCartProductsByWallet
// @tags User.Order
// @Param coupon query string false "coupon"
// @Router /user/order/checkout/wallet [post]
// @Success 200 {object} response.Order "html added"
// @Failure 406 "string "Invalid input"
// @Failure 400 "somthing wrong!!"
func (p *OrderHandler) OrderCartProductsByWallet(c *gin.Context) {

	var body request.Order

	body.CouponCode = c.Query("coupon")
	body.PaymentMethod = "Wallet"
	userId := auth.GetUserIdFromContext(c)
	order, _, err := p.orderService.OrderCartProducts(c, userId, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	response := gin.H{
		"":     "order details",
		"data": order,
	}
	c.JSON(http.StatusOK, response)

}

//====================================================== cancel order ===================================================

// CancelOrder godoc
// @summary API for canceling an order
// @security ApiKeyAuth
// @ID CancelOrder
// @tags User.Order
// @Param input body request.CancelOrder{} true "inputs"
// @Router /user/order/cancelorder [delete]
// @Success 200 "cancel order successful"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "Something went wrong !"
func (a *AdminHandler) CancelOrder(c *gin.Context) {

	var body request.CancelOrder
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "Missing or invalid entry"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if body.UsersID == 0 {
		body.UsersID = auth.GetUserIdFromContext(c)
	}

	msg, err := a.adminService.CancelOrder(c, body)

	if err != nil {
		respone := "failed cancel"
		c.JSON(http.StatusInternalServerError, respone)
		return
	}
	data := gin.H{
		"Message": msg,
	}
	c.JSON(http.StatusOK, data)

}

//====================================================== update order status ===================================================

// UpDateOrderStatus godoc
// @summary api for update order status of a user/order by using id
// @Param input body request.UpDateOrderStatus{} true "Input Fields"
// @security ApiKeyAuth
// @id UpdateOrderStatus
// @Accept json
// @Produce json
// @tags NotUse
// @Router /admin/order/updaterderstatus [patch]
// @Success 200 "Updated order status"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "Something went wrong !"
func (p *OrderHandler) UpDateOrderStatus(c *gin.Context) {

	var body request.UpDateOrderStatus
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err := p.orderService.UpDateOrderStatus(c, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	response := gin.H{
		"": "order status updated",
	}
	c.JSON(http.StatusOK, response)
}

//====================================================== return order request===================================================

// ReturnOrder godoc
// @summary api for return order
// @security ApiKeyAuth
// @id ReturnOrder
// @tags User.Order
//
//	@Produce		json
//
// @Param orderID query uint false "orderID"
// @Router /user/order/returnorder [get]
// @Success 200 "requst to return successfully"
// @Failure 400 "Missing or invalid entry/error"
func (p *OrderHandler) ReturnOrder(c *gin.Context) {

	qID := c.Query("orderID")
	fmt.Println("id=", qID)
	ID, err := strconv.Atoi(qID)
	if err != nil {
		return
	}
	userId := auth.GetUserIdFromContext(c)
	order, err1 := p.orderService.ReturnOrder(c, userId, uint(ID))

	if err1 != nil {
		Response := gin.H{
			"Data": err1.Error(),
		}
		c.JSON(http.StatusBadRequest, Response)
		return
	}
	// success response
	Response := gin.H{
		"Message": "send return reqst successfully",
		"Data":    order,
	}

	c.JSON(http.StatusOK, Response)

}

//	  GetOrderDetails godoc
//		@Summary		order details get by id
//		@Description	order details get by id
//		@Tags			User.Order
//		@Produce		json
//		@Param			orderID	 query		int	true	"Order ID"
//		@Success		200		{object}	 response.Orders1{}
//		@Failure		400		"faild"
//		@Failure		500		"faild"
//		@Router		/user/order/orderdetails [get]
func (p *OrderHandler) GetOrderDetails(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Query("orderID"))
	if err != nil {
		response := "Invalid entry"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	userId := auth.GetUserIdFromContext(c)

	orderDetails, err := p.orderService.GetOrderDetails(c, userId, uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	responce := gin.H{
		"success": "get details successfuly",
		"data":    orderDetails,
	}
	c.JSON(http.StatusOK, responce)

}

// DownloadInvoice godoc
//
//	@Summary		Download invoice
//	@Description	Download the invoice as a PDF file.
//	@Tags			User.Order
//	@Produce		application/pdf
//	@Param			orderID	 query		int	true	"Order ID"
//	@Success		200		{file}		application/pdf
//	@Failure		400		"faild"
//	@Failure		500		"faild"
//	@Router	/user/order/invoice/:orderID [get]
func (p *OrderHandler) DownloadInvoice(c *gin.Context) {
	str := c.Query("orderID")
	orderID, err := strconv.Atoi(str)
	if err != nil {
		response := "Invalid entry"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	userId := auth.GetUserIdFromContext(c)
	pdfData, err := p.orderService.CreateInvoice(c, userId, uint(orderID))
	if err != nil {
		response := "Failed"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// Set the response headers for downloading the file
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=invoice.pdf")

	// Send the PDF data as the response
	c.Data(http.StatusOK, "application/pdf", pdfData)
}

// ReturnRefund godoc
// @summary refund process of return product
// @Discription api for update order status of a user/order by using id
// @security ApiKeyAuth
// @id ReturnRefund
// @Accept json
// @tags Admin.OrderStatus
//
//	@Param	orderID	 query		int	true	"orderid"
//	@Param	userID	 query		int	true	"userid"
//
// @Router /admin/order/updateorderstatus/refund [patch]
// @Success 200 "Updated order status"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "Something went wrong !"
func (p *OrderHandler) ReturnRefund(c *gin.Context) {

	str := c.Query("orderID")
	orderID, err1 := strconv.Atoi(str)
	if err1 != nil {
		response := "Invalid entry"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	u := c.Query("userID")
	userID, err2 := strconv.Atoi(u)
	if err2 != nil {
		response := "Invalid entry"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var body request.UpDateOrderStatus

	body.OrderID = uint(orderID)
	body.UsersID = uint(userID)
	body.Status = "return refund"

	err := p.orderService.UpDateOrderStatus(c, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	response := gin.H{
		"": "order status updated as return refund",
	}
	c.JSON(http.StatusOK, response)
}

// OrderShipped godoc
// @summary order shipped status
// @discription api for update order status of a user/order by using id
//
//	@Param orderID query int true "orderid"
//	@Param userID query int	true "userid"
//
// @security ApiKeyAuth
// @id OrderShipped
// @Accept json
// @tags Admin.OrderStatus
// @Router /admin/order/updateorderstatus/shipped [patch]
// @Success 200 "Updated order status"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "Something went wrong !"
func (p *OrderHandler) OrderShiped(c *gin.Context) {

	str := c.Query("orderID")
	orderID, err1 := strconv.Atoi(str)
	if err1 != nil {
		response := "Invalid entry"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	u := c.Query("userID")
	userID, err2 := strconv.Atoi(u)
	if err2 != nil {
		response := "Invalid entry"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var body request.UpDateOrderStatus

	body.OrderID = uint(orderID)
	body.UsersID = uint(userID)
	body.Status = "order Shipped"

	err := p.orderService.UpDateOrderStatus(c, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	response := gin.H{
		"": "order status updated as order shipped",
	}
	c.JSON(http.StatusOK, response)
}

// OrderDelivered godoc
// @summary order delivered status
// @Discription api for update order status of a user/order by using id
// @security ApiKeyAuth
// @id OrderDelivered
// @Accept json
//
//	@Param	orderID	 query		int	true	"orderid"
//	@Param	userID	 query		int	true	"userid"
//
// @tags Admin.OrderStatus
// @Router /admin/order/updateorderstatus/delivered [patch]
// @Success 200 "Updated order status"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "Something went wrong !"
func (p *OrderHandler) OrderDelivered(c *gin.Context) {

	str := c.Query("orderID")
	orderID, err1 := strconv.Atoi(str)
	if err1 != nil {
		response := "Invalid entry"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	u := c.Query("userID")
	userID, err2 := strconv.Atoi(u)
	if err2 != nil {
		response := "Invalid entry"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var body request.UpDateOrderStatus

	body.OrderID = uint(orderID)
	body.UsersID = uint(userID)
	body.Status = "order Delivered"

	err := p.orderService.UpDateOrderStatus(c, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	response := gin.H{
		"": "order status updated As order delivered",
	}
	c.JSON(http.StatusOK, response)
}

// RetreturnRequestAccept godoc
// @summary refund process of return request verified
// @Discription api for update order status of a user/order by using id
// @security ApiKeyAuth
// @id RetreturnRequestAccept
// @Accept json
// @tags Admin.OrderStatus
//
//	@Param	orderID	 query		int	true	"orderid"
//	@Param	userID	 query		int	true	"userid"
//
// @Router /admin/order/updateorderstatus/returnrequestaccept [patch]
// @Success 200 "Updated order status"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "Something went wrong !"
func (p *OrderHandler) RetreturnRequestAccept(c *gin.Context) {

	str := c.Query("orderID")
	orderID, err1 := strconv.Atoi(str)
	if err1 != nil {
		response := "Invalid entry"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	u := c.Query("userID")
	userID, err2 := strconv.Atoi(u)
	if err2 != nil {
		response := "Invalid entry"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var body request.UpDateOrderStatus

	body.OrderID = uint(orderID)
	body.UsersID = uint(userID)
	body.Status = "return request accepted"

	err := p.orderService.UpDateOrderStatus(c, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	response := gin.H{
		"": "order status updated as requst accepted",
	}
	c.JSON(http.StatusOK, response)
}
