package model

type RegisterReq struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Password string `password:"password"`
}

type RegisterResp struct {
	AccessToken string `json:"access_token"`
}

type CreateTokenReq struct {
	UserId    string `json:"user_id"`
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

type GetTokenResp struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetUserResp struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}
