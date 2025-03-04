package model

import "github.com/dgrijalva/jwt-go"

// RegisterRequest структура запроса на регистрацию
type RegisterRequest struct {
	OrganizationName string `json:"organization_name"`
	ContactPerson    string `json:"contact_person" binding:"required"`
	Address          string `json:"address" binding:"required"`
	Phone            string `json:"phone" binding:"required,e164"`
	Requisites       string `json:"requisites" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	Password         string `json:"password" binding:"required,min=1"`
	ConfirmPassword  string `json:"confirm_password" binding:"required,min=1,eqfield=Password"`
}

// AuthRequest структура запроса на аутентификацию
type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=1"`
}

// AuthResponse структура ответа с токенами
type AuthResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

// UserClaims структура claims jwt-токена
type UserClaims struct {
	jwt.StandardClaims
	Email string `json:"username"`
	Role  int64  `json:"role"`
}
