package model

// User модель пользователя в репо слое
type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     int32  `db:"role"`
}

// UserInfo модель пользователя для Авторизации
type UserInfo struct {
	Username string `db:"username"`
	Password string `db:"password"`
	Role     int32  `db:"role"`
}
