package handler

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Respone struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserReponse struct {
	Token string `json:"token"`
}
