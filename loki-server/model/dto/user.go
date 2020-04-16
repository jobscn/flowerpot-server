package dto

type UserLoginIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInfo struct {
	UID      int64  `json:"uid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
}
