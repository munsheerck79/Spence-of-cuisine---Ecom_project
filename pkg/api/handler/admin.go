package handler

import (
	"encoding/csv"
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
)

type AdminHandler struct {
	adminService interfacess.AdminService
}

func NewAdminHandler(adminUsecase interfacess.AdminService) *AdminHandler {
	return &AdminHandler{adminService: adminUsecase}
}

//======================================================== Admin home page =========================================================

// AdminHome godoc
// @summary api for Admin home page
// @security ApiKeyAuth
// @id AdminHome
// @tags Admin
// @Param input body request.LoginAdminData{} true "Input Fields"
// @Router /admin/home [get]
// @Success 200 "Login successful,welcome to home"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "Something went wrong !"
func (u *AdminHandler) AdminHome(c *gin.Context) {

	response := "Welcome to admin side !"
	c.JSON(http.StatusOK, response)
}

//======================================================admin login page ===================================================

// LoginSubmit godoc
// @summary api for Admin login
// @security ApiKeyAuth
// @id AdminLogin
// @tags Admin
// @Param input body request.LoginAdminData{} true "Input Fields"
// @Router /admin/login [post]
// @Success 200 "Login successful"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "Something went wrong !"
func (u *AdminHandler) AdminLogin(c *gin.Context) {
	var body request.LoginAdminData
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
	var admin domain.Admin
	copier.Copy(&admin, body)

	dbAdmin, err := u.adminService.LoginAdmin(c, admin)
	if err != nil {
		response := "Something went wrong !"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := gin.H{"Successfuly send OTP to registered mobile number": dbAdmin.ID}
	c.JSON(http.StatusOK, response)
}

// ======================================================= admin logout ======================================================
// AdminLogout godoc
// @summary api for Admin logout
// @security ApiKeyAuth
// @id LogOutAdmin
// @tags Admin
// @Router /admin/logout [get]
// @Success 200 "Logot successful"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "failed to send OTP"
func (a *AdminHandler) LogoutAdmin(c *gin.Context) {
	c.SetCookie("admin-auth", "", -1, "", "", false, true)
	response := "Log out successful"
	c.JSON(http.StatusOK, response)
}

// ====================================================otp verification ==================================================
// OTPVerification godoc
// @summary api for Admin otp verification
// @security ApiKeyAuth
// @id OtpVerify
// @tags Admin
// @Param input body request.OTPVerify{} true "Input Fields"
// @Router /admin/otp-verify [post]
// @Success 200 "Login successful"
// @Failure 400 "Missing or invalid entry"
// @Failure 500  "failed to send OTP"
func (u *AdminHandler) AdminOTPVerify(c *gin.Context) {

	var body request.OTPVerify
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "Missing or invalid entry"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var admin = domain.Admin{
		ID: body.UserID,
	}

	usr, err := u.adminService.OTPLogin(c, admin)
	if err != nil {
		response := "user not registered"
		c.JSON(http.StatusBadRequest, response)
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
	ok := auth.JwtCookieSetup(c, "admin-auth", usr.ID)
	fmt.Println("jwt set")
	if !ok {
		response := "failed to login"
		c.JSON(http.StatusInternalServerError, response)
		return

	}
	response := "Successfuly logged in by admin!"
	c.JSON(http.StatusOK, response)
}

//========================================================  list users  =======================================================

// Listusers godoc
// @summary api for getuserlist
// @security ApiKeyAuth
// @id ListUsers
// @tags Admin.UserDash
// @Param page_number query int false "Page Number"
// @Param count query int false "Count Of Order"
// @Router /admin/users/ [get]
// @Success 200 {object} domain.Users{} "Login successful"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "Something went wrong !"
func (a *AdminHandler) ListUsers(c *gin.Context) {

	var page request.ReqPagination
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

	userList, err := a.adminService.GetUserlist(c,page)
	if err != nil {
		respone := "failed to get all users"
		c.JSON(http.StatusInternalServerError, respone)
		return
	}

	// check there is no user
	if len(userList) == 0 {
		response := "No user to show"
		c.JSON(http.StatusOK, response)
		return
	}

	data := gin.H{
		"Message": "List user successful",
		"Data":    userList,
	}

	c.JSON(http.StatusOK, data)

}

//==============================================================  block user =======================================================

// BlockUser godoc
// @summary api for admin to block or unblock user
// @id BlockUser
// @tags Admin.UserDash
// @Param input body request.BlockUser{} true "inputs"
// @Router /admin/users/blockuser [patch]
// @Success 200 {object} request.BlockUser{} "Successfully changed user block_status"
// @Failure 400 "invalid input"
func (a *AdminHandler) BlockUser(c *gin.Context) {

	var body request.BlockUser
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "Missing or invalid entry"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if body.UserName == "" && body.UserID == 0 {
		response := "Field should not be empty"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msg, err := a.adminService.BlockUser(c, body)

	if err != nil {
		respone := "failed to block user"
		c.JSON(http.StatusInternalServerError, respone)
		return
	}
	data := gin.H{
		"Message": msg,
		"Data":    body.UserID,
	}
	c.JSON(http.StatusOK, data)

}

// UserDetails godoc
// @summary api for get user details for admin
// @security ApiKeyAuth
// @tags Admin.UserDash
// @Param input body request.UserDetails{} true "Input Fields"
// @Router /admin/users/userdetails [get]
// @Success 200 {object} response.UserDetails{} "Successfully get user details"
// @Failure 400 "Missing or invalid entry"
func (a *AdminHandler) UserDetails(c *gin.Context) {
	var body request.UserDetails
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "Missing or invalid entry"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userDetails, err := a.adminService.UserDetails(c, body)
	if err != nil {
		return
	}
	data := gin.H{
		"Data": userDetails,
	}
	c.JSON(http.StatusOK, data)
}

// ListOrder godoc
// @summary api for get order list
// @security ApiKeyAuth
// @id ListOrder
// @Accept json
// @Produce json
// @tags Admin.OrderDash
// @Router /admin/order/ [get]
// @Param page_number query int false "Page Number"
// @Param count query int false "Count Of Order"
// @Success 200 {object} []response.AdminOrderList{} " get order list"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "Something went wrong !"
func (a *AdminHandler) ListOrder(c *gin.Context) {

var page request.ReqPagination
co:= c.Query("count")
pa:=c.Query("page_number")
count, err0 := strconv.Atoi(co)
page_number, err1 := strconv.Atoi(pa)
err0 = errors.Join(err0, err1)
if err0 != nil {
	response :="Missing or invalid inputs"
	c.JSON(http.StatusBadRequest, response)
	return
}
page.PageNumber = uint(page_number)
page.Count = uint(count)

	orderList, err := a.adminService.GetOrderlist(c,page)
	if err != nil {
		respone := "failed to get all users"
		c.JSON(http.StatusInternalServerError, respone)
		return
	}

	// check there is no user
	if len(orderList) == 0 {
		response := "No user to show"
		c.JSON(http.StatusOK, response)
		return
	}

	data := gin.H{
		"Message": "List order successful",
		"Data":    orderList,
	}

	c.JSON(http.StatusOK, data)

}

// CancelOrder godoc
// @summary API for canceling an order
// @security ApiKeyAuth
// @ID CancelOrder
// @tags Admin.OrderDash
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

// SalesReport godoc
// @summary api for admin to download sales report as csv format
// @id SalesReport
// @tags Admin
// @Router /admin/sales-report [get]
// @Success 500 "success"
// @Failure 500 "Something went wrong! failed to generate sales report"
// @Failure 400 "Missing or Invalid inputs"
func (a *AdminHandler) SalesReport(c *gin.Context) {
	var body request.DateRange
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "Missing or invalid entry"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	salesReport, err := a.adminService.SalesReport(c, body)
	if err != nil {
		response := "error in get sales report"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// set header for downloading browser
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment;filename= Spences_of_spices_sales_report.csv")
	wr := csv.NewWriter(c.Writer)

	headers := []string{"Order ID", "User Name", "Name", "Total Amount", "Discount", "Net Amount", "Coupon Code", "Payment Method", "Order Status", "Order Date"}
	if err := wr.Write(headers); err != nil {
		response := "error in get sales report"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	for _, sales := range salesReport {
		var row = []string{
			fmt.Sprintf("%v", sales.ID),
			fmt.Sprintf("%v", sales.UserName),
			fmt.Sprintf("%v", sales.Name),
			fmt.Sprintf("%v", sales.Total_amound),
			fmt.Sprintf("%v", sales.Discount),
			fmt.Sprintf("%v", sales.NetAmount),

			sales.CouponCode,
			sales.PaymentMethod,
			sales.OrderStatus,
			sales.OrderDate.Format("2006-01-02 15:04:05")}

		if err := wr.Write(row); err != nil {

			response := "Something went wrong! failed to generate sales report"
			c.JSON(http.StatusBadRequest, response)
			return
		}

	}
	// Flush the writer's buffer to ensure all data is written to the client
	wr.Flush()
}
