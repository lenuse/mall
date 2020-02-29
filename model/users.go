package model

type User struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	EmailVerifiedAt string `json:"email_verified_at"`
	Phone           string `json:"phone"`
	PhoneVerifiedAt string `json:"phone_verified_at"`
	Password        string `json:"password"`
	RememberToken   string `json:"remember_token"`
	State           int    `json:"state"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
}


func (m *User) GetTableName() string {
	return "users"
}

// GetUserByID 通过用户id查询用户
func GetUserByID(id string) User {
	user := User{}
	db.First(&user, id)
	return user
}