package request

import "time"

type LoginAdminData struct {
	UserName string `json:"user_name"  binding:"required,min=3,max=15"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"  binding:"required"`
}
type BlockUser struct {
	UserName string `json:"user_name"  binding:"required,min=3,max=15"`
	UserID   uint   `json:"user_id" binding:"required,numeric"`
}

type UserDetails struct {
	UserName string `json:"user_name"  binding:"required,min=3,max=15"`
	UserID   uint   `json:"user_id" binding:"required,numeric"`
}

type DateRange struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
