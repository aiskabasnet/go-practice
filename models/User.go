package models

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address"`
}

func (b *User) TableName() string {
	return "user"
}
