package models

type UserPrimaryKey struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
}

type CreateUser struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type User struct {
	Id        string `json:"id"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateUser struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type UserGetListResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"users"`
}
