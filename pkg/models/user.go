package models

import (
	"github.com/sinameshkini/auth-service/pkg/enums"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string         `json:"username"`
	Mobile   string         `json:"mobile"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Role     enums.UserRole `json:"role"`
}
