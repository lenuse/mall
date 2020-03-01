package transport

// AdminLogin 用户登录参数映射
type AdminLogin struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
