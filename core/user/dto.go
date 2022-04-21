package user

type RegisterCredentials struct {
	Name     string `json:"name"`
	Login    string `json:"username"`
	Password string `json:"password"`
}

type SignInCredentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
