package dto

type SessionLoginParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SessionToken struct {
	Token        string `json:"token"`
	ExpiresIn    int    `json:"expireIn"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenParam struct {
	Token string `json:"token" bindingL"required"`
}
