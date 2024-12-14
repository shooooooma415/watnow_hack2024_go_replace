package auth

type AuthResponse struct {
	Id int `json:"id"`
}

type SuccessResponse struct {
	IsSuccess bool `json:"is_success"`
}

type SignUp struct {
	UserName string `json:"user_name"`
	AuthId   string `json:"auth_id"`
	Token    string `json:"token"`
}

type SignIn struct {
	AuthId string `json:"auth_id"`
}
