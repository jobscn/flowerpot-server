package dto

type UserLoginIn struct {
	//Username string `json:"username" binding:"required"`
	//Password string `json:"password" binding:"required"`
}

type UserLoginOut struct {
}

type UserRegisterByPhoneIn struct {
	Phone    string
	Password string
}

type UserRegisterByPhoneOut struct {
}
