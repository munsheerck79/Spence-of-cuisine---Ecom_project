package usecase

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"github.com/munsheerck79/Ecom_project.git/helpers"
	"github.com/munsheerck79/Ecom_project.git/pkg/config"
	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/pkg/repository/interfaces"
	"github.com/munsheerck79/Ecom_project.git/pkg/usecase/interfacess"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/munsheerck79/Ecom_project.git/util/response"
	"github.com/razorpay/razorpay-go"
)

type OrderUsecase struct {
	orderRepository   interfaces.OrderRepository
	productRepository interfaces.ProductRepository
	userRepository    interfaces.UserRepository
	paymentRepository interfaces.PaymentRepository
	userService       interfacess.UserService
}

func NewOrderService(repo interfaces.OrderRepository, prodRepo interfaces.ProductRepository, userRepo interfaces.UserRepository, PayRepo interfaces.PaymentRepository, userserv interfacess.UserService) interfacess.OrderService {
	return &OrderUsecase{orderRepository: repo,
		productRepository: prodRepo,
		userRepository:    userRepo,
		paymentRepository: PayRepo,
		userService:       userserv,
	}
}

func (p *OrderUsecase) BuyProduct(c context.Context, userId uint) (domain.Address, error) {
	address, err1 := p.userRepository.FindAddress(c, userId)

	if err1 != nil {
		return domain.Address{}, err1
	}

	return address, nil
}

func (p *OrderUsecase) OrderCartProducts(c context.Context, userId uint, body request.Order) (response.Order, string, error) {

	address, err5 := p.userRepository.FindAddress(c, userId)

	if err5 != nil {
		return response.Order{}, "", fmt.Errorf("address is not avalable or error at find address")
	}
	if address.ID == 0{
		return response.Order{}, "", fmt.Errorf("address is not avalable please add address")
	}
	var Order response.Order
	var Total float32
	var DiscPrice float32
	productList, total, discprice, err3 := p.userService.ListCart(c, userId)
	if err3 != nil {
		return Order, "", err3
	}
	DiscPrice = discprice
	Total = total
	fmt.Println(DiscPrice)
	fmt.Println(Total)

	for i := 0; i < len(productList); i++ {
		fmt.Println("for")
		qty := int(productList[i].QtyInStock) - productList[i].Quantity
		if qty < 0 {

			return Order, "", fmt.Errorf("%v product is not avalable or less stock", productList[i].ProductName)
		}
	}
	fmt.Println("dis", DiscPrice)
	NetAmount := DiscPrice

	/////couponcode verification==================

	if body.CouponCode != "" {
		coupon, err := p.orderRepository.GetCouponByCode(c, body.CouponCode)
		if err != nil {
			return Order, "", fmt.Errorf("%v coupon is not avalable", body.CouponCode)
		}
		if coupon.ID == 0 {
			return Order, "", fmt.Errorf("%v coupon is not avalable", body.CouponCode)
		}
		n := time.Now().Unix()
		if coupon.ValidTill < n {
			return Order, "", fmt.Errorf("%v coupon is expired", body.CouponCode)
		}
		if coupon.MinOrderValue > float64(DiscPrice) {
			return Order, "", fmt.Errorf("%v minimum order value is higher than your net amount", body.CouponCode)
		}

		discountamound := (int(DiscPrice) * coupon.DiscountPercent / 100)

		if coupon.DiscountMaxAmount < float64(discountamound) {
			NetAmount = DiscPrice - float32(coupon.DiscountMaxAmount)
		} else {
			NetAmount = DiscPrice - float32(discountamound)
		}
		body.CouponId = coupon.ID
	} else {
		fmt.Println("nil coupon")
		body.CouponId = 0
	}

	var NewOrderID uint
	choice := body.PaymentMethod
	// Switch statement

	switch choice {
	case "online":

		fmt.Println("online payment")
		///razorpay immplimentation====================
		fmt.Println(NetAmount)

		RazorPayOrderId, RazorPayKey, err := MakeRazorPayPaymentId(int(NetAmount * 100))
		if err != nil {
			fmt.Println("444444444")
			return Order, "", err
		}
		Order.PaymentMethod = "online"
		Order.RazorPayOrderId = RazorPayOrderId
		Order.NetAmount = NetAmount
		Order.UsersID = userId
		Order.OrderStatus = "product"

		xstatus := "Product"
		orderIdx, err := p.orderRepository.OrderProductsTemp(c, userId, body, NetAmount, Total, DiscPrice, xstatus, RazorPayOrderId)
		if err != nil {
			return Order, "", err
		}
		for j := 0; j < len(productList); j++ {
			err := p.orderRepository.AddItemsToOrderItemsTemp(c, productList[j], orderIdx)
			if err != nil {
				return Order, "", err
			}
		}
		return Order, RazorPayKey, nil

	case "online Payment":

		fmt.Println("digital payment success")
		TempOrder, err := p.orderRepository.FindTempDataByPaymentId(body.RazorPayPaymentId)
		if err != nil {
			return Order, "", err
		}

		var b request.Order
		b.CouponId = TempOrder.CouponID
		b.PaymentMethod = "online"
		b.RazorPayPaymentId = TempOrder.PaymentId

		order, orderId, err := p.orderRepository.OrderCartProducts(c, TempOrder.UsersID, body, TempOrder.NetAmount, TempOrder.ActualPrice, TempOrder.DiscountPrice)
		Order = order
		if err != nil {
			return Order, "", err
		}
		NewOrderID = orderId

		tempList, err := p.orderRepository.OrderTempItems(TempOrder.ID)
		if err != nil {
			return Order, "", err
		}
		productList = tempList
		///delete function

	case "cod":
		fmt.Println("COD selected")
		body.RazorPayPaymentId = "cash on delivery"
		fmt.Println(NetAmount)
		fmt.Println(Total)
		fmt.Println(DiscPrice)
		order, orderId, err := p.orderRepository.OrderCartProducts(c, userId, body, NetAmount, Total, DiscPrice)
		Order = order
		if err != nil {
			return Order, "", err
		}
		NewOrderID = orderId
	case "Wallet":
		fmt.Println("wallet selected")
		//vallet implimentation======================

		wallet, err := p.userRepository.GetWalletx(c, userId)
		if err != nil {
			return Order, "", err
		}
		if wallet.Balence < NetAmount {
			return Order, "", fmt.Errorf("insufficent balence")
		} else {

			body.RazorPayPaymentId = "Wallet"
			order, orderId, err := p.orderRepository.OrderCartProducts(c, userId, body, NetAmount, Total, DiscPrice)
			Order = order
			if err != nil {
				return Order, "", err
			}
			NewOrderID = orderId
			str := strconv.FormatUint(uint64(Order.ID), 10)
			paymentId := ("payment of %s" + str)
			debitamound := -NetAmount
			err1 := p.paymentRepository.AddMonyToWallet(c, userId, paymentId, debitamound)
			if err1 != nil {
				return order, "", err1
			}
		}

	default:
		fmt.Println("Invalid choice")
		return Order, "", fmt.Errorf("invalid method ")
	}

	var orderstatusU request.UpDateOrderStatus

	orderstatusU.OrderID = NewOrderID
	orderstatusU.UsersID = Order.UsersID
	orderstatusU.Status = "order placed"
	orderStatus, err2 := p.orderRepository.GetOrderStatusId(c, orderstatusU)
	if err2 != nil {
		return Order, "", err2
	}
	err1 := p.orderRepository.UpDateOrderStatus(c, orderStatus)
	if err1 != nil {
		return Order, "", err1
	}

	orderD, err3 := p.orderRepository.OrderDetails(c, userId, Order.ID)
	if err3 != nil {
		return Order, "", err3
	}

	for z := 0; z < len(productList); z++ {
		qty := int(productList[z].QtyInStock) - productList[z].Quantity
		err4 := p.productRepository.UpdateProductStock(c, productList[z].ProductId, qty)
		if err4 != nil {
			return Order, "", err4
		}
	}

	for j := 0; j < len(productList); j++ {
		err := p.orderRepository.AddItemsToOrderItems(c, productList[j], NewOrderID)
		if err != nil {
			return Order, "", err
		}
	}

	for k := 0; k < len(productList); k++ {
		err := p.userRepository.DeleteFromCart(c, productList[k].ID, userId)
		if err != nil {
			return Order, "", err
		}
	}
	fmt.Println("sccess usecase")
	return orderD, "", nil
}

func (p *OrderUsecase) UpDateOrderStatus(ctx context.Context, body request.UpDateOrderStatus) error {

	orderStatus, err := p.orderRepository.GetOrderStatusId(ctx, body)
	if err != nil {
		return err
	}
	if orderStatus.OrderStatusID == 0 {
		return fmt.Errorf("%v status is not avalable", orderStatus.Status)
	} else {
		if body.Status == "return refund" {
			orderD, err3 := p.orderRepository.OrderDetails(ctx, body.UsersID, body.OrderID)
			if err3 != nil {
				return err3
			}
			if orderD.OrderStatus == "money has refunded" {
				return fmt.Errorf("the amount allredy send")
			} else {
				if orderD.OrderStatus == "return request accepted" {
					str := strconv.FormatUint(uint64(orderD.ID), 10)
					paymentId := ("return of %s" + str)
					err := p.paymentRepository.AddMonyToWallet(ctx, body.UsersID, paymentId, orderD.NetAmount)
					if err != nil {
						return err
					}
					orderStatus.Status = "money has refunded"
				}
			}
		}
		orderStatus1, err := p.orderRepository.GetOrderStatusId(ctx, orderStatus)
		if err != nil {
			return err
		}
		err2 := p.orderRepository.UpDateOrderStatus(ctx, orderStatus1)
		if err2 != nil {
			return err
		}
	}
	return nil
}

func (p *OrderUsecase) GetCoupon(c context.Context) ([]domain.Coupon, error) {

	coupons, err := p.orderRepository.GetCoupon(c)
	if err != nil {
		return coupons, err
	}
	return coupons, nil

}
func (p *OrderUsecase) AddCoupon(c context.Context, body domain.Coupon) error {
	err := p.orderRepository.AddCoupon(c, body)
	if err != nil {
		return err
	}
	return nil

}
func (p *OrderUsecase) EditCoupon(c context.Context, body request.EditCoupon) error {

	coupon, err := p.orderRepository.GetCouponByCode(c, body.Code)
	if err != nil {
		return fmt.Errorf("%v coupon is not avalable err", body.Code)
	}
	if coupon.ID == 0 {
		return fmt.Errorf("%v coupon is not avalable", body.Code)
	}
	if body.Description == "" {
		body.Description = coupon.Description
	}
	if body.DiscountMaxAmount == 0 {
		body.DiscountMaxAmount = coupon.DiscountMaxAmount
	}
	if body.DiscountPercent == 0 {
		body.DiscountPercent = coupon.DiscountPercent
	}
	if body.MinOrderValue == 0 {
		body.MinOrderValue = coupon.MinOrderValue
	}
	if body.ValidTill == 0 {
		body.ValidTill = coupon.ValidTill
	} else {
		body.ValidTill = body.ValidTill * (time.Now().Add(24 * time.Hour).Unix())
	}
	err1 := p.orderRepository.EditCoupon(c, body)
	if err1 != nil {
		return err1
	}
	return nil

}

func (o *OrderUsecase) ReturnOrder(c context.Context, userId uint, orderID uint) (response.Order, error) {

	order, err := o.orderRepository.OrderDetails(c, userId, orderID)
	if err != nil {
		return order, fmt.Errorf("err @ oder getting,order is not avalable")
	}
	if order.ID == 0 {

		return order, fmt.Errorf("%v order is not avalable", orderID)
	}

	newDate := order.OrderDate.AddDate(0, 0, 15)
	// Get today's date
	today := time.Now().UTC()
	// Check if the new date is before today
	if newDate.Before(today) {
		fmt.Println("return requst date is before today:", newDate.Format("2006-01-02"))
		return order, fmt.Errorf("return facility is not avleble ,time out")
	}

	if order.OrderStatus != "order Delivered" {

		return order, fmt.Errorf("return facility is not avleble ,order not delivered")
	}

	if order.OrderStatus == "order Delivered" {
		var body request.UpDateOrderStatus
		body.OrderID = orderID
		body.UsersID = userId
		body.Status = "return request has placed"
		err := o.orderRepository.UpDateOrderStatus(c, body)

		if err != nil {
			return order, err
		}
	}
	order1, err := o.orderRepository.OrderDetails(c, userId, orderID)
	if err != nil {
		return order1, err
	}
	return order1, nil
}

// DBProduct, err := p.orderRepository.CheckStock(ctx, order)
// if err != nil {
// 	return err
// }
// fmt.Println(DBProduct.QtyInStock)
// fmt.Println(order.Qty)
// if DBProduct.QtyInStock >= order.Qty {

// 	err = p.orderRepository.PlaceOrder(ctx, order, DBProduct)
// 	if err != nil {
// 		return err
// 	}

// } else {
// 	return fmt.Errorf("%v product is not avalable or less stock", DBProduct.Name)
// }

// return nil
func (p *OrderUsecase) MakeRazorPayPayment(c context.Context, amount float32) (string, string, error) {
	RazorPayOrder, RazorPayKey, err := MakeRazorPayPaymentId(int(amount * 100))
	if err != nil {
		return RazorPayOrder, "", err
	}
	return RazorPayOrder, RazorPayKey, nil
}

func MakeRazorPayPaymentId(amount int) (string, string, error) {
	fmt.Println("55555")
	razorPayKey := config.GetConfig().RAZORPAYKEY
	razorPaySecret := config.GetConfig().RAZORPAYSECRET

	client := razorpay.NewClient(razorPayKey, razorPaySecret)

	data := map[string]interface{}{
		"amount":   amount,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)

	if err != nil {
		return "", "", err
	}
	//save order id from the body
	RazorpayOrderId := body["id"].(string)
	RazorPayKey := razorPayKey
	return RazorpayOrderId, RazorPayKey, nil
}

// func (o *OrderUsecase) ListCartord(ctx context.Context, userId uint) ([]response.Cart, float32, float32, error) {

// 	TotalPrice := float32(0)
// 	TotalDiscPrice := float32(0)
// 	cartList, err := o.orderRepository.CartListord(ctx, userId)
// 	for i := 0; i < len(cartList); i++ {
// 		TotalPrice += cartList[i].ActualPrice * float32(cartList[i].Quantity)
// 		TotalDiscPrice += cartList[i].DiscountPrice * float32(cartList[i].Quantity)
// 	}
// 	if err != nil {
// 		return []response.Cart{}, TotalPrice, TotalDiscPrice, err
// 	}
// 	return cartList, TotalPrice, TotalDiscPrice, nil

// }

func (o *OrderUsecase) CreateInvoice(c context.Context, userId uint, orderID uint) ([]byte, error) {
	println("creatinvoice", orderID)
	orderData, err1 := o.orderRepository.OrderDetails(c, userId, orderID)
	if err1 != nil {
		println("creatinvoiceerr")
		return nil, err1
	}
	println("amount=", int(orderData.NetAmount))
	if orderData.ID == 0 {
		return nil, err1
	}
	DBAddress, err2 := o.userRepository.FindAddress(c, userId)
	if err2 != nil {
		return nil, err2
	}
	println("address", DBAddress.Address)
	if DBAddress.ID == 0 {
		return nil, err2
	}
	productList, err3 := o.orderRepository.GetOrderItemsById(orderID)
	if err3 != nil {
		return nil, err3
	}
	if productList == nil {
		return nil, err3
	}

	var productsData []map[string]interface{}
	for _, product := range productList {
		productData := map[string]interface{}{
			"--------------": "----------------------------",
			"Product Name":   product.ProductName,
			"Variation Name": product.VariationName,
			"Quantity":       product.Quantity,
			"Actual Price":   product.ActualPrice,
			"Discount Price": product.DiscountPrice,
		}
		productsData = append(productsData, productData)
	}

	invoiceData := map[string]interface{}{

		"Order Date": orderData.OrderDate.String(),
		"Order ID":   fmt.Sprint(orderID),

		"Delivery Address": map[string]string{
			"Address":     DBAddress.Address,
			"Muncipality": DBAddress.Muncipality,
			"LandMark":    DBAddress.LandMark,
			"District":    DBAddress.District,
			"State":       DBAddress.State,
			"PhoneNumber": DBAddress.PhoneNumber,
			"PinCode":     DBAddress.PinCode,
		},

		"Product name": productsData,

		"Payment method": orderData.PaymentMethod,

		"Total Amount": fmt.Sprint(orderData.NetAmount),
	}
	println("get helper")
	return helpers.GenerateInvoicePDF(invoiceData), nil
}

func (p *OrderUsecase) GetOrderDetails(c context.Context, userId uint, orderID uint) (response.Orders1, error) {
	var orderDetails response.Orders1

	orderData, err := p.orderRepository.OrderDetails(c, userId, orderID)
	if err != nil {
		return response.Orders1{}, err
	}
	if orderData.ID == 0 {
		return response.Orders1{}, err
	}
	productList, err := p.orderRepository.GetOrderItemsById(orderID)
	if err != nil {
		return response.Orders1{}, err
	}
	if productList == nil {
		return response.Orders1{}, err
	}

	if err := copier.Copy(&orderDetails, orderData); err != nil {
		fmt.Println("Copy failed")
	}
	if err := copier.Copy(&orderDetails.Items, productList); err != nil {
		fmt.Println("Copy failed")
	}

	return orderDetails, nil
}
