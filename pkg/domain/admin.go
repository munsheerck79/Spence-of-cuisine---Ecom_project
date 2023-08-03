package domain

//admin model
type Admin struct {
	ID        uint   `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	UserName  string `json:"user_name" gorm:"not null" binding:"required,min=3,max=15"`
	FirstName string `json:"first_name" gorm:"not null" binding:"required,min=2,max=40"`
	LastName  string `json:"last_name" gorm:"not null" binding:"required,min=1,max=40"`
	Age       uint   `json:"age" gorm:"not null" binding:"required,numeric"`
	Email     string `json:"email" gorm:"unique;not null" binding:"required,email"`
	Phone     string `json:"phone" gorm:"unique;not null" binding:"required,min=10,max=10"`
	Password  string `json:"password" gorm:"not null" binding:"required"`
}
