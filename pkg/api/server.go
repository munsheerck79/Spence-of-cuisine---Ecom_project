package http

import (
	"github.com/gin-gonic/gin"
	_ "github.com/munsheerck79/Ecom_project.git/cmd/api/docs"
	"github.com/munsheerck79/Ecom_project.git/pkg/api/handler"
	"github.com/munsheerck79/Ecom_project.git/pkg/api/middleware"
	swaggoFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler,
) *ServerHTTP {

	engine := gin.New()

	engine.LoadHTMLGlob("templates/*.html") //  loading html for razorpay payment

	// to load views
	// Serve static files
	// engine.Static("/assets", "./views/static/assets")

	// Add the Gin Logger middleware.
	engine.Use(gin.Logger())

	// Get swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggoFiles.Handler))

	//===========================================================  USER Side     ====================================================
	//  Calling user api
	engine.POST("user/order/payment/razorpay/process-order", paymentHandler.ProccessRazorpayOrder)

	apiuser := engine.Group("/user")

	apiuser.POST("/signup", userHandler.UserSignup)
	apiuser.GET("/home", userHandler.UserHome)
	apiuser.POST("/login", userHandler.UserLogin)
	// OTP verfication
	apiuser.POST("/otp-verify", userHandler.UserOTPVerify)

	apiuser.Use(middleware.UserAuthentication)

	{

		//productdashbord := apiuser.Group("/product")
		apiuser.GET("/login", userHandler.UserHome)
		apiuser.GET("/logout", userHandler.LogoutUser)
		apiuser.GET("/addaddress", userHandler.GetAddress)
		apiuser.POST("/addaddress", userHandler.AddAddress)
		apiuser.GET("/editaddress", userHandler.GetAddress)
		apiuser.PUT("/editaddress", userHandler.EditAddress)
		apiuser.GET("/getwallet", paymentHandler.GetWallet)

		userproductdashbord := apiuser.Group("/product")
		{
			userproductdashbord.GET("/", productHandler.ProductList)
			userproductdashbord.GET("/listproductbyid", productHandler.GetProduct)
			userproductdashbord.GET("/listproductsbycatogory", productHandler.GetProductsByCategoryName)
			userproductdashbord.GET("/category", productHandler.GetCategory)

		}

		usercartdashbord := apiuser.Group("/cart")
		{
			usercartdashbord.GET("/list", userHandler.ListCart)
			usercartdashbord.POST("/addtocart", userHandler.AddToCart)
			usercartdashbord.PUT("/editcart", userHandler.EditCartProduct)
			usercartdashbord.DELETE("/deletefromcart/:ID", userHandler.DeleteFromCart)
		}
		userwishlistdashbord := apiuser.Group("/wishlist")
		{
			userwishlistdashbord.GET("/show", userHandler.ListWishList)
			userwishlistdashbord.POST("/addtowishlist", userHandler.AddToWishList)
			userwishlistdashbord.DELETE("/deletefromwishlist/:WishlistID", userHandler.DeleteFromWishLIst)

		}

		apiuser.GET("/coupon", orderHandler.GetCoupon)

		userorderdashbord := apiuser.Group("/order")
		{
			userorderdashbord.GET("/confirm-details", orderHandler.BuyProduct)
			//userorderdashbord.POST("/checkout", orderHandler.OrderCartProducts)
			userorderdashbord.POST("/checkout/wallet", orderHandler.OrderCartProductsByWallet)
			userorderdashbord.POST("/checkout/cod", orderHandler.OrderCartProductsByCOD)

			userorderdashbord.DELETE("/cancelorder", adminHandler.CancelOrder)
			userorderdashbord.GET("/orderhistory", userHandler.OrderHistory)
			userorderdashbord.GET("/returnorder", orderHandler.ReturnOrder)
			userorderdashbord.GET("/orderdetails", orderHandler.GetOrderDetails)
			userorderdashbord.GET("/invoice/:orderID", orderHandler.DownloadInvoice)
			paymentdashbord := userorderdashbord.Group("/payment")
			{
				// Razorpay
				paymentdashbord.GET("/checkout/razorpay", paymentHandler.RazorPayCheckout)

				//addmoney  to wllet
				paymentdashbord.GET("/addmonytowallet", paymentHandler.RazorPayCheckoutWallet)
				//paymentdashbord.POST("/razorpay/process-wallet", paymentHandler.ProccessRazorpayWallet)

			}

		}

	}
	//========================================================== admin side  ===========================================================

	apiadmin := engine.Group("/admin")
	{
		//apiadmin.POST("signup", userHandler.UserSignup)//seller

		apiadmin.POST("/login", adminHandler.AdminLogin)
		// OTP verfication
		apiadmin.POST("/otp-verify", adminHandler.AdminOTPVerify)

		apiadmin.Use(middleware.AdminAuthentication)
		{
			apiadmin.GET("/login", adminHandler.AdminHome)
			apiadmin.GET("/home", adminHandler.AdminHome)
			apiadmin.GET("/logout", adminHandler.LogoutAdmin)
			apiadmin.GET("/salesreport", adminHandler.SalesReport)

			userdashbord := apiadmin.Group("/users")
			{
				userdashbord.GET("/", adminHandler.ListUsers)
				userdashbord.PATCH("/blockuser", adminHandler.BlockUser)
				userdashbord.GET("/userdetails", adminHandler.UserDetails)

			}

			productdashbord := apiadmin.Group("/product")
			{
				productdashbord.GET("/", productHandler.ProductListAdmin)
				productdashbord.GET("/getproductbyid", productHandler.GetProductAdmin)
				productdashbord.GET("/listproductsbycatogory", productHandler.GetProductsByCategoryNameAdmin)
				productdashbord.POST("/addproduct", productHandler.AddProduct)
				productdashbord.PUT("/editproduct", productHandler.EditProduct)

				//productdashbord.GET("/productprice")
				productdashbord.POST("/addprice", productHandler.AddPrice)
				productdashbord.PUT("/editprice", productHandler.EditPrice)

				productdashbord.GET("/category", productHandler.GetCategoryAdmin)
				productdashbord.POST("/addcategory", productHandler.AddCategory)
				productdashbord.GET("/variation", productHandler.Getvariations)
				productdashbord.POST("/addvariation", productHandler.AddVarient)
				productdashbord.GET("/orderstatus", productHandler.GetOrderStatus)
				productdashbord.POST("/addorderstatus", productHandler.AddOrderStatus)

			}

			orderdashbord := apiadmin.Group("/order")
			{
				orderdashbord.GET("/", adminHandler.ListOrder)
				orderdashbord.GET("/orderdetailsbyid", orderHandler.GetOrderDetailsAdmin)
				orderdashbord.DELETE("/cancelorder", adminHandler.CancelOrderAdmin)
				orderdashbord.GET("/coupons", orderHandler.GetCouponAdmin)
				orderdashbord.POST("/addcoupon", orderHandler.AddCoupon)
				orderdashbord.PUT("/editcoupon", orderHandler.EditCoupon)
				orderdashbord.PATCH("/updateorderstatus", orderHandler.UpDateOrderStatus)
				orderUpdateDashbord := orderdashbord.Group("/updateorderstatus")
				{
					orderUpdateDashbord.PATCH("/shipped", orderHandler.OrderShiped)
					orderUpdateDashbord.PATCH("/delivered", orderHandler.OrderDelivered)
					orderUpdateDashbord.PATCH("/refund", orderHandler.ReturnRefund)
					orderUpdateDashbord.PATCH("/returnrequestaccept", orderHandler.RetreturnRequestAccept)
				}
			}

		}

		return &ServerHTTP{engine: engine}
	}
	//local host
}
func (s *ServerHTTP) Run() {
	s.engine.Run(":3000")
}
