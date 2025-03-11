package req

type User struct {
	Email    string `json:"email" example:"test@test.com"`
	Password string `json:"password" example:"123456"`
}

type Token struct {
	Token string `json:"token"`
}
