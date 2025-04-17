package model

// User модель пользователя сервисного слоя
type User struct {
	UserID   int64  `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     int32  `json:"role"`
}

// UserUpdate модель обновления пользователя сервисного слоя
type UserUpdate struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Role     int32  `json:"role"`
	Email    string `json:"email"`
}

// Client модель клиента сервисного слоя
type Client struct {
	ClientID         int    `json:"client_id"`
	OrganizationName string `json:"organization_name"`
	ContactPerson    string `json:"contact_person"`
	Address          string `json:"address"`
	Phone            string `json:"phone"`
	Requisites       string `json:"requisites"`
}
