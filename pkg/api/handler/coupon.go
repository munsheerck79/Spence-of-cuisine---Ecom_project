package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/pkg/usecase/interfacess"
	"github.com/munsheerck79/Ecom_project.git/util/request"
)

type CouponHandler struct {
	couponService interfacess.CouponService
}

func NewCouponHandler(couponUsecase interfacess.CouponService) *CouponHandler {
	return &CouponHandler{couponService: couponUsecase}
}

//====================================================== get coupon list ===================================================

// GetCoupon godoc
// @summary api for get coupons for admin and user
// @id GetCoupon
// @security ApiKeyAuth
// @Produce json
// @tags User
// @Router /user/coupon [get]
// @Success 200 {object} domain.Coupon{}
// @Failure 500 "Something went wrong !"
func (p *CouponHandler) GetCoupon(c *gin.Context) {

	coupons, err := p.couponService.GetCoupon(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	Response := gin.H{
		"Message": "List coupons successful",
		"Data":    coupons,
	}

	c.JSON(http.StatusOK, Response)

}

//====================================================== get coupon list ===================================================

// GetCouponAdmin godoc
// @summary api for get coupons for admin and user
// @id GetCouponAdmin
// @security ApiKeyAuth
// @Produce json
// @tags Admin.Coupon
// @Router /admin/order/coupon [get]
// @Success 200 {object} domain.Coupon{}
// @Failure 500 "Something went wrong !"
func (p *CouponHandler) GetCouponAdmin(c *gin.Context) {

	coupons, err := p.couponService.GetCoupon(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	Response := gin.H{
		"Message": "List coupons successful",
		"Data":    coupons,
	}

	c.JSON(http.StatusOK, Response)

}

//====================================================== Add coupons ===================================================

// AddCoupon godoc
// @summary api for admin to add coupon
// @security ApiKeyAuth
// @id AddCoupon
// @tags Admin.Coupon
// @Param input body  request.AddCoupon{} true "inputs"
// @Router /admin/order/addcoupon [post]
// @Success 200 "coupon added successfully"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "Something went wrong !"
func (p *CouponHandler) AddCoupon(c *gin.Context) {

	var body request.AddCoupon
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var coupon domain.Coupon
	copier.Copy(&coupon, body)

	if err := p.couponService.AddCoupon(c, coupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	c.JSON(http.StatusOK, gin.H{"success": "coupon created successfuly"})

}

//====================================================== edit coupon ===================================================

// EditCoupon godoc
// @summary API for edit coupon by admin
// @description edit coupon
// @security ApiKeyAuth
// @id EditCoupon
// @tags Admin.Coupon
// @Accept json
// @Param inputs body request.EditCoupon{} true "Input Field"
// @Router /admin/order/editcoupon [put]
// @Success 200 "Successfully address added"
// @Failure 400 "string "Invalid input"
func (p *CouponHandler) EditCoupon(c *gin.Context) {
	var body request.EditCoupon
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := p.couponService.EditCoupon(c, body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	c.JSON(http.StatusOK, gin.H{"success": "coupon edited successfuly"})

}
