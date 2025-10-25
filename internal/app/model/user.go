package model

import "time"

// User 用户模型
// @Description 用户基本信息
// @Property ID uint64 用户ID
// @Property Username string 用户名
// @Property Password string 密码
// @Property Email string 邮箱
type User struct {
	ID       uint64 `gorm:"primaryKey autoIncrement"`
	Username string `gorm:"unique"`
	Password string `gorm:"password"`
	Email    string `gorm:"unique"`
}

func (u *User) TableName() string {
	return "t_user"
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string    `json:"username"`
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expire_at"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
