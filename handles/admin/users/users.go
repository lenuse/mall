package users

type createUser struct {
	username string `json:"username" binding:"required"`
	mobile   string `json:"mobile" binding:"required"`
}

func create() {

}
