package transport

// AdminSignIn 用户登录参数映射
type AdminSignIn struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// AdminCreate 创建后台用户映射
type AdminCreate struct {
	Username    string `form:"username" json:"username" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
	RepeatInput string `form:"repeat_input" json:"repeat_input" binding:"required,eqcsfield=Password"`
	NickName    string `form:"nickname" json:"nickname" binding:"required"`
	Email       string `form:"email" json:"email" binding:"required,email"`
	Note        string `form:"note" json:"note"`
}
