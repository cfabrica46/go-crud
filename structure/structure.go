package structure

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ResponseHTTP struct {
	Content   interface{} `json:"content"`
	Code      int         `json:"code"`
	ErrorText string      `json:"err"`
}

type Token struct {
	Token string `json:"Authorization-header"`
}
