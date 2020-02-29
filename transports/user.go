package transports

type UserLoginTransport struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
