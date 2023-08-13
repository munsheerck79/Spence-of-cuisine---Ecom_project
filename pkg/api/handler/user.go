package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/munsheerck79/Ecom_project.git/pkg/api/auth"
	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/pkg/usecase/interfacess"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/munsheerck79/Ecom_project.git/util/response"
)

type UserHandler struct {
	userService interfacess.UserService
}

//======================================================================================================

func NewUserHandler(userUsecase interfacess.UserService) *UserHandler {
	return &UserHandler{userService: userUsecase}
}

//=======================================================user signup===============================================

// UserSignUp godoc
// @summary api for register user
// @security ApiKeyAuth
// @id UserSignUp
// @tags User
// @Param input body request.SignupUserData{} true "Input Fields"
// @Router /user/signup [post]
// @Success 200 "Account created successfuly"
// @Failure 400 "invalid input"
func (u *UserHandler) UserSignup(c *gin.Context) {
	var body request.SignupUserData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//
	var user domain.Users
	// var user domain.Users
	if err := copier.Copy(&user, body); err != nil {
		fmt.Println("Copy failed")
	}

	// Check the user already exist in DB and save user if not
	if err := u.userService.SignUp(c, user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	c.JSON(http.StatusOK, gin.H{"success": "Account created successfuly"})

}

//================================================================ user login ====================================================

// LoginSubmit godoc
// @summary api for user login
// @security ApiKeyAuth
// @id UserLogin
// @tags User
// @Param input body request.LoignUserData{} true "Input Fields"
// @Router /user/login [post]
// @Success 200 "Login successful"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "Something went wrong !"
func (u *UserHandler) UserLogin(c *gin.Context) {
	var body request.LoignUserData
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "Missing or invalid entry"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if body.Email == "" && body.Password == "" && body.UserName == "" {
		response := "Field should not be empty"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// Copying
	var user domain.Users
	copier.Copy(&user, body)

	dbUser, err := u.userService.Login(c, user)
	if err != nil {
		response := gin.H{"blocked by admin1": user.UserName}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := gin.H{"Successfuly send OTP to registered mobile number": dbUser.ID}
	c.JSON(http.StatusOK, response)
}

// ===========================================  logout user ===================================================================
// LogoutUser godoc
// @summary api for logout
// @security ApiKeyAuth
// @id LogoutUser
// @tags User
// @Router /user/logout [get]
// @Success 200 "Logout successful"
// @Failure 500 "Something went wrong !"
func (u *UserHandler) LogoutUser(c *gin.Context) {
	c.SetCookie("user-auth", "", -1, "", "", false, true)
	response := "Log out successful"
	c.JSON(http.StatusOK, response)
}

//===================================================== user otp verification===========================================

// OTPVerification godoc
// @summary api for user otp verification
// @security ApiKeyAuth
// @id UserOtpVerify
// @tags User
// @Param input body request.OTPVerify{} true "Input Fields"
// @Router /user/otp-verify [post]
// @Success 200 "Login successful"
// @Failure 400  "Missing or invalid entry"
// @Failure 401 "failed to send OTP"
// @Failure 500 "detect error"
func (u *UserHandler) UserOTPVerify(c *gin.Context) {

	var body request.OTPVerify
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "Missing or invalid entry"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var user = domain.Users{
		ID: body.UserID,
	}

	usr, err := u.userService.OTPLogin(c, user)
	if err != nil {
		response := "user not registered"
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	// Verify otp
	err = auth.TwilioVerifyOTP("+91"+usr.Phone, body.OTP)
	if err != nil {
		response := gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// setup JWT
	ok := auth.JwtCookieSetup(c, "user-auth", usr.ID)
	if !ok {
		response := "failed to login"
		c.JSON(http.StatusInternalServerError, response)
		return

	}
	response := "Successfuly logged in!"
	c.JSON(http.StatusOK, response)
}

//=====================================================home===========================================

// UserHome godoc
// @summary api for user home page
// @description after user login user will seen this page with user informations using mid.ware
// @security ApiKeyAuth
// @id UserHome
// @tags User
// @Router /user/login [get]
// @Success 200 "Login successfully !"
func (u *UserHandler) UserHome(c *gin.Context) {
	userId := auth.GetUserIdFromContext(c)
	response := gin.H{
		"":   "Welcome to home !",
		"id": userId,
	}

	c.JSON(http.StatusOK, response)
}

// ListCart godoc
// @summary API for get cart list
// @description get cart list pf autherised user
// @security ApiKeyAuth
// @id ListCart
// @tags User.Cart
// @Produce json
// @Router /user/cart/list [get]
// @Success 200 {object} ListCartResponse
// @Failure 400 "Something went wrong !"
func (u *UserHandler) ListCart(c *gin.Context) {

	userId := auth.GetUserIdFromContext(c)

	cart, total, discprice, err := u.userService.ListCart(c, userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saved := total - discprice

	// success response
	response := gin.H{
		"total price": total,
		"Net amount":  discprice,
		"you saved":   saved,
		"":            "get cart sucsessfully",
		"data":        cart,
	}
	c.JSON(http.StatusOK, response)

}

type ListCartResponse struct {
	List          []response.Cart
	TotalPrice    float32
	DiscountPrice float32
}

// GetAddress godoc
// @summary API for user to get address
// @security ApiKeyAuth
// @ID GetAddress
// @tags User.Details
// @Router /user/addaddress [get]
// @Success 200 {object} domain.Address{} "get address sucsessfully"
// @Failure 400 "Missing or invalid entry"
func (u *UserHandler) GetAddress(c *gin.Context) {
	userId := auth.GetUserIdFromContext(c)

	// Check the user already exist in DB and save user if not
	address, err := u.userService.GetAddress(c, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	response := gin.H{
		"":     "get address sucsessfully",
		"data": address,
	}
	c.JSON(http.StatusOK, response)

}

// AddAddress godoc
// @summary API for adding a new address for a user
// @description Get a new address from the user to store in the database
// @security ApiKeyAuth
// @id AddAddress
// @tags User.Details
// @Param inputs body request.Address{} true "Input Field"
// @Router /user/addaddress [post]
// @Success 200 {string} string "Successfully address added"
// @Failure 400 {string} string "Invalid input"
func (u *UserHandler) AddAddress(c *gin.Context) {

	var body request.Address
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	userId := auth.GetUserIdFromContext(c)
	body.UsersID = userId
	var address domain.Address
	if err := copier.Copy(&address, body); err != nil {
		fmt.Println("Copy failed")
	}
	// Check the user already exist in DB and save user if not
	if err := u.userService.AddAddress(c, address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	c.JSON(http.StatusOK, gin.H{"success": "Account created successfuly"})
}

// EditAddress godoc
// @summary API for editing address for a user
// @security ApiKeyAuth
// @ID EditAddress
// @tags User.Details
// @Param inputs body request.EditAddress{} true "Input Field"
// @Router /user/editaddress [put]
// @Success 200 "Successfully address added"
// @Failure 400 "Invalid input"
func (u *UserHandler) EditAddress(c *gin.Context) {
	var editData request.EditAddress
	if err := c.ShouldBindJSON(&editData); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	userId := auth.GetUserIdFromContext(c)
	editData.UsersID = userId
	var address domain.Address
	if err := copier.Copy(&address, editData); err != nil {
		fmt.Println("Copy failed")
	}
	if err := u.userService.EditAddress(c, address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	c.JSON(http.StatusOK, gin.H{"success": "Account created successfuly"})

}

// AddToCart godoc
// @summary api for add product to user cart
// @description user can add a product to cart
// @security ApiKeyAuth
// @id AddToCart
// @tags User.Cart
// @Param input body request.Cart{} true "Input Field"
// @Router /user/cart/addtocart [post]
// @Success 200 "Successfuly added product item to cart "
// @Failure 406 "StatusNotAcceptable,invalid input!! "
// @Failure 500 "Failed to add product item in cart"
func (u *UserHandler) AddToCart(c *gin.Context) {

	var body request.Cart
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusNotAcceptable, response)
		return
	}
	userId := auth.GetUserIdFromContext(c)
	body.UsersID = userId
	if body.Quantity == 0 {
		body.Quantity = 1
	}

	if err := u.userService.AddToCart(c, body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	c.JSON(http.StatusOK, gin.H{"success": "Add to cart successfuly"})
}

// DeleteFromCart godoc
// @summary API for delete cart product
// @description delete cart items
// @security ApiKeyAuth
// @id DeleteFromCart
// @tags User.Cart
// @Produce json
// @Accept json
// @Param	ID   query string  false	"ID"
// @Router /user/cart/deletefromcart/:ID [delete]
// @Success 200  "successfully deleted from cart"
// @Failure 400 "Something went wrong !"
func (u *UserHandler) DeleteFromCart(c *gin.Context) {

	cartID, err := strconv.Atoi(c.Query("ID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cartID"})
		return
	}

	userId := auth.GetUserIdFromContext(c)

	if err := u.userService.DeleteFromCart(c, uint(cartID), userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	c.JSON(http.StatusOK, gin.H{"success": "delete product from cart successfully"})
}

// EditCartProduct godoc
// @summary API for edit cart product
// @description edit cart items
// @security ApiKeyAuth
// @id EditCartProduct
// @tags User.Cart
// @Produce json
// @Accept json
// @param input body request.EditCart{} true "Input Field"
// @Router /user/cart/editcart [put]
// @Success 200  "edited successfully"
// @Failure 400 "Something went wrong !"
// @Failure 406 "StatusNotAcceptable,invalid input!! "
func (u *UserHandler) EditCartProduct(c *gin.Context) {

	var body request.EditCart
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusNotAcceptable, response)
		return
	}
	userId := auth.GetUserIdFromContext(c)
	body.UsersID = userId
	var product domain.Cart
	if err := copier.Copy(&product, body); err != nil {
		fmt.Println("Copy failed")
	}

	if err := u.userService.EditCartProduct(c, product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	c.JSON(http.StatusOK, gin.H{"success": " edit cart successfuly"})
}

// AddToWishList godoc
// @summary api for add product item to user wishlist
// @description user can add a stock in product to wishlist
// @security ApiKeyAuth
// @id AddToWishList
// @tags User.Wishlist
// @Param  ID query string false "ID"
// @Router /user/wishlist/addtowishlist [post]
// @Success 200 "Successfuly added product item to wish "
// @Failure 400 "Failed to add product item in wish"
func (u *UserHandler) AddToWishList(c *gin.Context) {
	str := c.Query("ID")
	ID, err := strconv.Atoi(str)
	if err != nil {
		return
	}
	var body request.WishList
	// if err := c.ShouldBindJSON(&body); err != nil {
	// 	response := "invalid input"
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }
	body.ProductID = uint(ID)
	userId := auth.GetUserIdFromContext(c)
	body.UsersID = userId
	var product domain.WishList
	if err := copier.Copy(&product, body); err != nil {
		fmt.Println("Copy failed")
	}

	if err := u.userService.AddToWishList(c, product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	c.JSON(http.StatusOK, gin.H{"success": "Add to wishlist successfuly"})
}

// DeleteFromWishLIst godoc
// @summary API for delete wishlist product
// @description delete wishlist items
// @security ApiKeyAuth
// @id DeleteFromWishLIst
// @tags User.Wishlist
// @Accept json
// @Param	WishlistID	query	string	 false	"WishlistID"
// @Router /user/wishlist/deletefromwishlist/:WishlistID [delete]
// @Success 200  "success"
// @Failure 400 "string "Invalid input"
// @Failure 500 "Something went wrong !"
func (u *UserHandler) DeleteFromWishLIst(c *gin.Context) {
	cartID := c.Query("WishlistID")
	Id, _ := strconv.Atoi(cartID)
	userId := auth.GetUserIdFromContext(c)

	if err := u.userService.DeleteFromWishLIst(c, uint(Id), userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	c.JSON(http.StatusOK, gin.H{"success": "delete from wishlist successfuly"})
}

// ListWishlist godoc
// @summary api for list product items in the wishlist
// @security ApiKeyAuth
// @id ListWishList
// @tags User.Wishlist
// @Router /user/wishlist/show [get]
// @Success 200 {object} response.ProductDetails{} "get list"
// @Failure 400 "Failed to list product item in wishlist"
func (u *UserHandler) ListWishList(c *gin.Context) {
	userId := auth.GetUserIdFromContext(c)

	wishList, err := u.userService.ListWishList(c, userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	response := gin.H{
		"":     "get wishlist sucsessfully",
		"data": wishList,
	}
	c.JSON(http.StatusOK, response)

}

// OrderHistory godoc
// @summary api for list orderitems
// @security ApiKeyAuth
// @id OrderHistory
// @tags User.Order
// @Param page_number query int false "Page Number"
// @Param count query int false "Count Of Order"
// @Router /user/order/orderhistory [get]
// @Success 200 {object} response.Order{}
// @Failure 400 "Failed to list order history"
func (u *UserHandler) OrderHistory(c *gin.Context) {
	var page request.ReqPagination
	userId := auth.GetUserIdFromContext(c)

	co := c.Query("count")
	pa := c.Query("page_number")
	count, err0 := strconv.Atoi(co)
	page_number, err1 := strconv.Atoi(pa)
	err0 = errors.Join(err0, err1)
	if err0 != nil {
		response := "Missing or invalid inputs"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	page.PageNumber = uint(page_number)
	page.Count = uint(count)

	orderHistory, err := u.userService.OrderHistory(c, userId, page)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	response := gin.H{
		"":     "get order history",
		"data": orderHistory,
	}
	c.JSON(http.StatusOK, response)

}