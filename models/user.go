package models

type CreateUser struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phone_number"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginRes struct {
	Token string `json:"token"`
}
type User struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Age         int    `json:"age"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type RequestByID struct {
	ID string
}
type RequestByUsername struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type ChangePassword struct {
	Id          string
	NewPassword string
	OldPassword string
}
type GetAllUsersResponse struct {
	Users []User `json:"users"`
	Count int32  `json:"count"`
}
type GetAllUsersRequest struct {
	Search string
	Job    string
	Age    int
	Page   int
	Limit  int
}
