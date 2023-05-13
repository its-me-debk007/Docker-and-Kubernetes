package model

type User struct {
	FirstName string `json:"first_name"    binding:"required"`
	LastName  string `json:"last_name"    binding:"required"`
	Email     string `json:"email"    binding:"required"    gorm:"unique"`
	Age       int8   `json:"age"    binding:"required"`
}
