package web

type AdminReponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
}

type AdminLoginResponse struct {
    Email    string `json:"email"`
    Password string `json:"password"`
    Token    string `json:"token"`
}