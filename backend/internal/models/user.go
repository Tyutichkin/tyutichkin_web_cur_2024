package models

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type SearchUserRequest struct {
	FullName string `json:"fullname"`
	Login    string `json:"login"`
}
