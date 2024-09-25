package models

// 定义请求参数结构体
type RegisterReq struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Gender     int8   `json:"gender" binding:"required,oneof=0 1"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}
