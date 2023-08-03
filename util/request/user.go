package request

type SignupUserData struct {
	UserName        string `json:"user_name"  binding:"required,min=3,max=15"`
	FirstName       string `json:"first_name"  binding:"required,min=2,max=50"`
	LastName        string `json:"last_name"  binding:"required,min=1,max=50"`
	Age             uint   `json:"age"  binding:"required,numeric"`
	Email           string `json:"email" binding:"required,email"`
	Phone           string `json:"phone" binding:"required,min=10,max=10"`
	Password        string `json:"password"  binding:"required,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type LoignUserData struct {
	UserName string `json:"user_name"  binding:"required,min=3,max=15"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"  binding:"required"`
}
type OTPVerify struct {
	OTP    string `json:"otp" binding:"required,min=4,max=8"`
	UserID uint   `json:"user_id" binding:"required,numeric"`
}
type Address struct {
	UsersID     uint   `json:"-"`
	Address     string `json:"address" binding:"required,min=5,max=40"`
	Muncipality string `json:"muncipality" binding:"required,min=4,max=20"`
	LandMark    string `json:"land_mark" binding:"required,min=4,max=20"`
	District    string `json:"district" binding:"required,min=4,max=20"`
	State       string `json:"state" binding:"required,min=4,max=20"`
	PhoneNumber string `json:"phone_number" binding:"required,min=10,max=10"`
	PinCode     string `json:"pin_code" binding:"required,min=6,max=10"`
}

type EditAddress struct {
	UsersID     uint   `json:"-"`
	Address     string `json:"address,omitempty"`
	Muncipality string `json:"muncipality,omitempty"`
	LandMark    string `json:"land_mark,omitempty"`
	District    string `json:"district,omitempty"`
	State       string `json:"state,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	PinCode     string `json:"pin_code,omitempty"`
}

type Cart struct {
	UsersID     uint `json:"-"`
	ProductID   uint `json:"product_id" binding:"required"`
	VariationID uint `json:"Variation_id" binding:"required"`
	Quantity    int  `json:"Quantity" gorm:"not null"`
}

type WishList struct {
	UsersID   uint `json:"-"`
	ProductID uint `json:"product_id" binding:"required"`
}

type VerifyPayment struct {
	Signature         string `json:"razorpay_signature" binding:"required"`
	RazorpayOrderId   string `json:"razorpay_order_id" binding:"required"`
	RazorPayPaymentId string `json:"razorpay_payment_id" binding:"required"`
}

type EditCart struct {
	UsersID     uint `json:"-"`
	ProductID   uint `json:"product_id" binding:"required"`
	VariationID uint `json:"Variation_id" binding:"omitempty"`
	Quantity    int  `json:"Quantity,omitempty"`
}
type ReqPagination struct {
	Count      uint `json:"count"`
	PageNumber uint `json:"page_number"`
}
