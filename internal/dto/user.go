package dto

type ReadUserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserReq struct {
	ID    int64  `param:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserAreaReq struct {
	ID     int64 `param:"id"`
	AreaID int64 `json:"area_id"`
}

type CreateUserReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type RegisterReq struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ListAreaByUsersReq struct {
	AreaID int64 `param:"area_id"`
}
