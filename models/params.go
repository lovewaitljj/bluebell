package models

// RegisterReq 注册请求
type RegisterReq struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Gender     int8   `json:"gender" binding:"required,oneof=0 1"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// LoginReq 登录请求
type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
