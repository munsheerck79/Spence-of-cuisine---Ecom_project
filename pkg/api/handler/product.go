package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/pkg/usecase/interfacess"
	"github.com/munsheerck79/Ecom_project.git/util/request"
)

type ProductHandler struct {
	productService interfacess.ProductService
}

func NewProductHandler(productUsecase interfacess.ProductService) *ProductHandler {
	return &ProductHandler{productService: productUsecase}
}

//=======================================================================================================================

// ProductList godoc
// @Summary API for admin or user to list all products
// @security ApiKeyAuth
// @tags User.Product
// @id ProductList
// @Produce json
// @Param page_number query int false "Page Number"
// @Param count query int false "Count Of Order"
// @Router /user/product/ [get]
// @Success 200 {object} response.ProductDetails{}
// @Success 204 "No products to show"
// @Failure 500 "Failed to get all products"
func (p *ProductHandler) ProductList(c *gin.Context) {

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

	products, err := p.productService.GetProductList(c, page)
	fmt.Println(products)
	fmt.Println(err)
	if err != nil {
		respone := "failed to get all products"
		c.JSON(http.StatusInternalServerError, respone)
		return
	}
	// check there is no products
	if len(products) == 0 {
		response := "No products to show"
		c.JSON(http.StatusNoContent, response)
		return
	}
	data := gin.H{
		"Message": "List product successful",
		"Data":    products,
	}

	c.JSON(http.StatusOK, data)

}

// ProductListAdmin godoc
// @Summary API for admin or user to list all products
// @security ApiKeyAuth
// @tags Admin.ProductDash
// @id ProductListAdmin
// @Produce json
// @Param page_number query int false "Page Number"
// @Param count query int false "Count Of Order"
// @Router /admin/product/ [get]
// @Success 200 {object} response.ProductDetails{}
// @Success 204 "No products to show"
// @Failure 500 "Failed to get all products"
func (p *ProductHandler) ProductListAdmin(c *gin.Context) {

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

	products, err := p.productService.GetProductList(c, page)
	fmt.Println(products)
	fmt.Println(err)
	if err != nil {
		respone := "failed to get all products"
		c.JSON(http.StatusInternalServerError, respone)
		return
	}
	// check there is no products
	if len(products) == 0 {
		response := "No products to show"
		c.JSON(http.StatusNoContent, response)
		return
	}
	data := gin.H{
		"Message": "List product successful",
		"Data":    products,
	}

	c.JSON(http.StatusOK, data)

}

// GetCategory godoc
// @summary API for get category list
// @description get category list for admin and user
// @security ApiKeyAuth
// @id GetCategory
// @tags User.Product
// @Produce json
// @Router /user/product/category [get]
// @Success 200 {object} []domain.Category
// @Failure 400 "string "Invalid input"
func (p *ProductHandler) GetCategory(c *gin.Context) {

	categoryList, err := p.productService.GetCategory(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	Response := gin.H{
		"Message": "List category successful",
		"Data":    categoryList,
	}

	c.JSON(http.StatusOK, Response)

}

// GetCategoryAdmin godoc
// @summary API for get category list
// @description get category list for admin and user
// @security ApiKeyAuth
// @id GetCategoryAdmin
// @tags Admin.ProductDash
// @Produce json
// @Router /admin/product/category [get]
// @Success 200 {object} []domain.Category
// @Failure 400 "string "Invalid input"
func (p *ProductHandler) GetCategoryAdmin(c *gin.Context) {

	categoryList, err := p.productService.GetCategory(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	Response := gin.H{
		"Message": "List category successful",
		"Data":    categoryList,
	}

	c.JSON(http.StatusOK, Response)

}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

// AddCatogory godoc
// @summary api for admin to add catogory
// @security ApiKeyAuth
// @id AddCategory
// @tags Admin.ProductDash
// @Param input body  request.AddCatogory{} true "inputs"
// @Router /admin/product/addcategory [post]
// @Success 200 "category added successfully"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "Something went wrong !"
func (p *ProductHandler) AddCategory(c *gin.Context) {

	var body request.AddCatogory
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var category domain.Category

	if err := copier.Copy(&category, body); err != nil {
		fmt.Println("Copy failed")
	}

	if err := p.productService.AddCategory(c, category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	c.JSON(http.StatusOK, gin.H{"success": "category created successfuly"})

}

// Getvariations godoc
// @summary API for get variation list
// @description get variation list for admin and user
// @security ApiKeyAuth
// @id Getvariations
// @tags Admin.ProductDash
// @Produce json
// @Router /admin/product/variation [get]
// @Success 200 {object} []domain.Variation "Successfully grt variation"
// @Failure 500 "Something went wrong !"
func (p *ProductHandler) Getvariations(c *gin.Context) {

	variationList, err := p.productService.Getvariations(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	Response := gin.H{
		"Message": "List variations successful",
		"Data":    variationList,
	}

	c.JSON(http.StatusOK, Response)

}

// ///////////////////////////////////////////////////////////////////////////////////////////

// AddVarient godoc
// @summary api for admin to add varient
// @security ApiKeyAuth
// @id AddVarient
// @tags Admin.ProductDash
// @Param input body  request.AddVarient{} true "inputs"
// @Router /admin/product/addvariation [post]
// @Success 200 "varient added successfully"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "Something went wrong !"
func (p *ProductHandler) AddVarient(c *gin.Context) {

	var body request.AddVarient
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var varient domain.Variation

	if err := copier.Copy(&varient, body); err != nil {
		fmt.Println("Copy failed")
	}

	if err := p.productService.AddVarient(c, varient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	c.JSON(http.StatusOK, gin.H{"success": "varient created successfuly"})

}

// ////////////////////////////////////////////////////////////////////////////////////////////

// AddProduct godoc
// @summary api for admin to add product
// @security ApiKeyAuth
// @id AddProduct
// @tags Admin.ProductDash
// @Param input body  request.AddProduct{} true "inputs"
// @Router /admin/product/addproduct [post]
// @Success 200 "product added successfully"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "Something went wrong !"
func (p *ProductHandler) AddProduct(c *gin.Context) {

	var body request.AddProduct
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var product domain.Product

	if err := copier.Copy(&product, body); err != nil {
		fmt.Println("Copy failed")
	}

	if err := p.productService.AddProduct(c, product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	c.JSON(http.StatusOK, gin.H{"success": "product added successfuly"})

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// EditProdct godoc
// @summary api for admin to edit product
// @security ApiKeyAuth
// @id EditProduct
// @tags Admin.ProductDash
// @Param input body  request.EditProduct{} true "inputs"
// @Router /admin/product/EditProduct [put]
// @Success 200 "edited successfully"
// @Failure 400 "Missing or invalid entry"
func (p *ProductHandler) EditProduct(c *gin.Context) {

	var body request.EditProduct
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var product domain.Product

	if err := copier.Copy(&product, body); err != nil {
		fmt.Println("Copy failed")
	}

	if err := p.productService.EditProduct(c, product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	c.JSON(http.StatusOK, gin.H{"success": "product edited successfuly"})

}

// AddPrice godoc
// @Summary api for admin to add price
// @security ApiKeyAuth
// @id AddPrice
// @tags Admin.ProductDash
// @Param input body  request.AddPrice{} true "inputs"
// @Router /admin/product/addprice [post]
// @Success 200 "price added successfully"
// @Failure 400 "Missing or invalid entry"
func (p *ProductHandler) AddPrice(c *gin.Context) {

	var body request.AddPrice
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var price domain.Price

	if err := copier.Copy(&price, body); err != nil {
		fmt.Println("Copy failed")
	}

	if err := p.productService.AddPrice(c, price); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	c.JSON(http.StatusOK, gin.H{"success": "price added successfuly"})

}

// EditPrice godoc
// @summary api for admin to edit price
// @security ApiKeyAuth
// @id EditPrice
// @tags Admin.ProductDash
// @Param input body  request.EditPrice{} true "inputs"
// @Router /admin/product/editprice [put]
// @Success 200 "edited successfully"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "Something went wrong !"
func (p *ProductHandler) EditPrice(c *gin.Context) {

	var body request.EditPrice
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var price domain.Price

	if err := copier.Copy(&price, body); err != nil {
		fmt.Println("Copy failed")
	}

	if err := p.productService.AddPrice(c, price); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	c.JSON(http.StatusOK, gin.H{"success": "price edited successfuly"})

}

// GetOrderStatus godoc
// @summary API for get order status list
// @description get order status list for admin
// @security ApiKeyAuth
// @id GetOrderStatus
// @tags Admin.OrderDash
// @Produce json
// @Router /admin/product/orderstatus [get]
// @Success 200 {object} []domain.OrderStatus
// @Failure 500 "Something went wrong !"
func (p *ProductHandler) GetOrderStatus(c *gin.Context) {

	statusList, err := p.productService.GetOrderStatus(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	Response := gin.H{
		"Message": "List status successfully",
		"Data":    statusList,
	}

	c.JSON(http.StatusOK, Response)

}

// AddOrderStatas godoc
// @summary api for admin to add status
// @security ApiKeyAuth
// @id AddOrderStatus
// @tags Admin.OrderDash
// @Param input body  domain.OrderStatus{} true "inputs"
// @Router /admin/product/addorderstatus [post]
// @Success 200 "status added successfully"
// @Failure 400 "Missing or invalid entry"
// @Failure 500 "Something went wrong !"
func (p *ProductHandler) AddOrderStatus(c *gin.Context) {

	var body domain.OrderStatus
	if err := c.ShouldBindJSON(&body); err != nil {
		response := "invalid input"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := p.productService.AddOrderStatus(c, body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// success response
	c.JSON(http.StatusOK, gin.H{"success": "order status created successfully"})

}

// GetProduct godoc
// @Summary Get a product by ID
// @Description GetProduct list for admin and user
// @Security ApiKeyAuth
// @ID GetProduct
// @Tags User.Product
// @Produce json
// @Param ID query uint false "productID"
// @Router /user/product/listproductbyid [get]
// @Success 200 {object} response.ProductDetails
// @Failure 400 {string} string "can't get product"
func (p *ProductHandler) GetProduct(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Query("ID"))
	fmt.Println("Product ID:", ID)
	product, err := p.productService.GetProduct(c, uint(ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	Response := gin.H{
		"Message": "List category successful",
		"Data":    product,
	}

	c.JSON(http.StatusOK, Response)
}

// GetProductAdmin godoc
// @Summary Get a product by ID
// @Description GetProduct list for admin and user
// @Security ApiKeyAuth
// @ID GetProductAdmin
// @Tags User.Product
// @Produce json
// @Param ID query uint false "productID"
// @Router /admin/product/getproductbyid [get]
// @Success 200 {object} response.ProductDetails
// @Failure 400 {string} string "can't get product"
func (p *ProductHandler) GetProductAdmin(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Query("ID"))
	product, err := p.productService.GetProduct(c, uint(ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	Response := gin.H{
		"Message": "List category successful",
		"Data":    product,
	}

	c.JSON(http.StatusOK, Response)
}

// GetProductsByCategoryName godoc
// @Summary Get a product by name
// @Description GetProduct list for admin and user by using name
// @Security ApiKeyAuth
// @ID GetProductsByCategoryName
// @Tags User.Product
// @Produce json
// @Param Name query string false "category"
// @Router /user/product/listproductsbycatogory [get]
// @Success 200 {object} response.ProductDetails{}
// @Success 204 "didnt get catogory name"
// @Failure 400 {string} string "can't get product"
func (p *ProductHandler) GetProductsByCategoryName(c *gin.Context) {
	name := c.Query("Name")
	fmt.Println(name)

	categoryList, err := p.productService.GetCategory(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var CID uint
	for i := 0; i < len(categoryList); i++ {
		if categoryList[i].CategoryName == name {
			CID = categoryList[i].ID
		}
	}
	if CID == 0 {
		Response := gin.H{
			"Message": "catogory is not get",
			"Data":    name,
		}
		c.JSON(http.StatusNoContent, Response)
		return
	}

	products, err := p.productService.GetProductsByCategoryName(c, CID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	Response := gin.H{
		"Message": "List category successful",
		"Data":    products,
	}

	c.JSON(http.StatusOK, Response)

}

// GetProductsByCategoryNameAdmin godoc
// @Summary Get a product by name
// @Description GetProduct list for admin and user by using name
// @Security ApiKeyAuth
// @ID GetProductsByCategoryNameAdmin
// @Tags Admin.ProductDash
// @Produce json
// @Param Name query string false "category"
// @Router /admin/product/listproductsbycatogory [get]
// @Success 200 {object} response.ProductDetails{}
// @Success 204 "didnt get catogory name"
// @Failure 400 {string} string "can't get product"
func (p *ProductHandler) GetProductsByCategoryNameAdmin(c *gin.Context) {
	name := c.Query("Name")
	fmt.Println(name)

	categoryList, err := p.productService.GetCategory(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var CID uint
	for i := 0; i < len(categoryList); i++ {
		if categoryList[i].CategoryName == name {
			CID = categoryList[i].ID
		}
	}
	if CID == 0 {
		Response := gin.H{
			"Message": "catogory is not get",
			"Data":    name,
		}
		c.JSON(http.StatusNoContent, Response)
		return
	}

	products, err := p.productService.GetProductsByCategoryName(c, CID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// success response
	Response := gin.H{
		"Message": "List category successful",
		"Data":    products,
	}

	c.JSON(http.StatusOK, Response)

}
