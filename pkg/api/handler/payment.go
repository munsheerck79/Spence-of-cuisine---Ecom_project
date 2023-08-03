package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/munsheerck79/Ecom_project.git/pkg/api/auth"
	"github.com/munsheerck79/Ecom_project.git/pkg/usecase/interfacess"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/munsheerck79/Ecom_project.git/util/response"
)

type PaymentHandler struct {
	paymentService interfacess.PaymentService
	orderService   interfacess.OrderService
	userService    interfacess.UserService
}

func NewPaymentHandler(paymentServ interfacess.PaymentService, orderServ interfacess.OrderService, userServ interfacess.UserService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentServ,
		userService:  userServ,
		orderService: orderServ}
}

// RazorPayCheckout godoc
// @summary api for user pay section
// @security ApiKeyAuth
// @id RazorPayCheckout
// @tags User.Payment
// @Param coupon query string false "coupon"
// @Router /user/order/payment/checkout/razorpay [get]
// @Success 200 "html added successfully"
// @Failure 404 "invalid "
// @Failure 400 "Missing or invalid entry"
func (r *PaymentHandler) RazorPayCheckout(c *gin.Context) {

	userId := auth.GetUserIdFromContext(c)
	fmt.Println("============000", userId)
	var body request.Order

	body.CouponCode = c.Query("coupon")
	// if err := c.ShouldBindJSON(&body); err != nil {
	// 	response := "invalid input"
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	body.PaymentMethod = "online"

	order, RazorPayKey, err := r.orderService.OrderCartProducts(c, uint(userId), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "error1"})
		return
	}

	c.HTML(200, "Razorpay.html", gin.H{
		"userId":            order.UsersID,
		"razorpay_order_id": order.RazorPayOrderId,
		"amount":            (int(order.NetAmount) * 100),
		"razor_pay_key":     RazorPayKey,
	})
}

//////////////////////////////////////////////////////////////////////////////

// MakePaymentRazorpay is the handler function for pay using razorpay.
//
//	@Summary		Make payment razorpay
//	@Description	Make payment using razorpay page .
//	@Tags			User.Payment
//
// @Param Amount query string false "Amount"
//
//	@Success		200
//	@Failure		500	"faild"
//	@Router			/user/order/payment/addmonytowallet [get]
func (r *PaymentHandler) RazorPayCheckoutWallet(c *gin.Context) {
	fmt.Println("qwe")
	userId := auth.GetUserIdFromContext(c)
	var Amount float64
	str := c.Query("Amount")
	fmt.Println("////", str)
	Amount, err := strconv.ParseFloat(str, 64)
	fmt.Println("././", Amount)
	if err != nil {
		fmt.Println("112233")
		// Handle the error if the conversion fails
		fmt.Println("Error converting string to float:", err)
		return
	}
	//var body request.RazorPayCheckoutWallet
	// if err := c.ShouldBindJSON(&body); err != nil {
	// 	response := "invalid input"
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }
	RazorPayOrderId, RazorPayKey, err := r.orderService.MakeRazorPayPayment(c, float32(Amount))
	if err != nil {
		fmt.Println("112244")
		return
	}

	fmt.Println("dadaaasasas")
	fmt.Println("22ww222", RazorPayKey)
	fmt.Println("dddwwd", RazorPayOrderId)
	var Order response.Order
	Order.UsersID = userId
	Order.OrderStatus = "wallet"
	Order.NetAmount = float32(Amount)
	Order.ActualPrice = float32(Amount)
	Order.DiscountPrice = float32(Amount)
	Order.RazorPayOrderId = RazorPayOrderId

	err2 := r.paymentService.AddTempData(c, Order)
	if err2 != nil {
		return
	}
	fmt.Println("22222", RazorPayKey)
	fmt.Println("html", RazorPayOrderId)

	c.HTML(200, "Razorpay.html", gin.H{
		"userId":            userId,
		"razorpay_order_id": RazorPayOrderId,
		"amount":            (int(Amount) * 100),
		"razor_pay_key":     RazorPayKey,
	})
}

///////////////////////////////////////////////////////////////////////////////////

// proccessRazorPayOrder godoc
// @summary api for user pay section
// @security ApiKeyAuth
// @id ProccessRazorpayOrder
// @tags User
//
//	@Description	Verify razorpay payment using razorpay credentials .
//	@Router	/user/order/payment/razorpay/process-order [post]
//	@Success 200	"success"
//	@Failure 400	"faild"
//	@Failure 406	"invalid input"
//
// @failiure 404 "not found enything"
func (r *PaymentHandler) ProccessRazorpayOrder(c *gin.Context) {
	fmt.Println("post rzr")
	var data map[string]interface{}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		fmt.Println("errin prc rzr")
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	// Now you can access the JSON data as needed using the map keys.
	// For example:
	event := data["event"].(string)
	entity := data["entity"].(string)
	orderID, orderIDOk := data["payload"].(map[string]interface{})["order"].(map[string]interface{})["entity"].(map[string]interface{})["id"].(string)
	paymentID, paymentIDOk := data["payload"].(map[string]interface{})["payment"].(map[string]interface{})["entity"].(map[string]interface{})["id"].(string)

	// Check if the order_id and payment_id fields exist and are of type string
	if orderIDOk {
		fmt.Println("Order ID:", orderID)
	} else {
		fmt.Println("Order ID field not found or not a string.")
	}

	if paymentIDOk {
		fmt.Println("Payment ID:", paymentID)
	} else {
		fmt.Println("Payment ID field not found or not a string.")
	}
	status, ok := data["payload"].(map[string]interface{})["payment"].(map[string]interface{})["entity"].(map[string]interface{})["status"].(string)

	// Check if the status field exists and is of type string
	if ok {
		fmt.Println("Status:", status)
	} else {
		fmt.Println("Status field not found or not a string.")
	}

	fmt.Println("Received event:", event)
	fmt.Println("Received entity:", entity)
	fmt.Println("Razorpay data - status:", status)
	fmt.Println("Razorpay data - order_id:", orderID)
	fmt.Println("Razorpay data - bank_transaction_id:", paymentID) // Convert float64 to int

	Data, err1 := r.paymentService.VerifyRazorPayPayment(c, orderID, status)
	fmt.Println("vrfy in")
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	fmt.Println("data", Data.OrderStatus)
	if Data.OrderStatus == "Product" {
		fmt.Println("payment successfull")
		userId := Data.UsersID
		var body1 request.Order

		body1.PaymentMethod = "online Payment"
		body1.RazorPayPaymentId = paymentID

		err := r.paymentService.AddPaymentId(paymentID, Data.RazorPayOrderId)
		if err != nil {
			return
		}
		order, _, err2 := r.orderService.OrderCartProducts(c, userId, body1)
		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
			return
		}
		// success response
		response := gin.H{
			"":     "order details",
			"data": order,
		}
		c.JSON(http.StatusOK, response)
		return
	}
	if Data.OrderStatus == "wallet" {
		fmt.Println("wallet payment successfull")
		userId := Data.UsersID

		err2 := r.paymentService.AddMonyToWallet(c, userId, paymentID, Data.NetAmount)
		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
			return
		}
		fmt.Println("money added to in your wallet")
		// success response
		response := gin.H{
			"": "money addded to wallet",
		}
		c.JSON(http.StatusOK, response)
		return
	}
	response := gin.H{
		"": "faild in all condition",
	}
	c.JSON(http.StatusNotFound, response)
}

// GetWallet godoc
// @Summary Get wallet history
// @Description Get wallet history of user
// @Security ApiKeyAuth
// @ID GetWallet
// @Tags User.Payment
// @Produce json
// @Router /user/getwallet [get]
// @Success 200 {object} response.WalletRes
// @Failure 400 "can't get product"
func (p *PaymentHandler) GetWallet(c *gin.Context) {

	userId := auth.GetUserIdFromContext(c)

	wallet, err := p.paymentService.GetWallet(c, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, "faild to get wallet")
		return
	}

	response := gin.H{
		"":     "wallet details",
		"data": wallet,
	}
	c.JSON(http.StatusOK, response)

}

// var body request.VerifyPayment
// err := c.BindJSON(&body)
// if err != nil {
// 	c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
// 	return
// }
// fmt.Println(body.RazorPayPaymentId)
//////get amount from success masg
// err1 := r.paymentService.VerifyRazorPayPayment(, body.RazorpayOrderId, body.RazorPayPaymentId)
// fmt.Println(err1)
// if err1 != nil {
// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	return
// }

// ProccessRazorpayOrder is the handler function for verify  razorpay payment.
//
//	@Summary		Verify razorpay payment
//	@Description	Verify razorpay payment using razorpay credentials .
//	@Tags			user
//	@Produce		json
//	@Param			body	body		request.VerifyPayment	true	"inputs"
//	@Success		200	"success"
//	@Failure		400	"faild"
//	@Failure		403	"faild"
//	@Failure		500	"faild"
//	@Router			/user/razorpay/process-wallet [post]

// func (r *PaymentHandler) ProccessRazorpayWallet(c *gin.Context) {
// 	var body request.VerifyPayment
// 	err := c.BindJSON(&body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	var money float32
// 	//////get amount from success masg
// 	err1 := r.paymentService.VerifyRazorPayPayment(body.Signature, body.RazorpayOrderId, body.RazorPayPaymentId)
// 	if err1 != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	fmt.Println("payment successfull")
// 	userId := auth.GetUserIdFromContext(c)

// 	err2 := r.paymentService.AddMonyToWallet(c, userId, body.RazorPayPaymentId, money)
// 	if err2 != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// success response
// 	response := gin.H{
// 		"": "money addded to wallet",
// 	}
// 	c.JSON(http.StatusOK, response)

// }
