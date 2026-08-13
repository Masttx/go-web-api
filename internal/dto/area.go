package dto

type ReadAreaResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateAreaReq struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateAreaReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AddUserReq struct {
	AreaID int64 `param:"area_id"`
	UserID int64 `param:"user_id"`
}

type ListUsersByAreaReq struct {
	AreaID int64 `param:"area_id"`
}

type DeleteUserReq struct {
	AreaID int64 `param:"area_id"`
	UserID int64 `param:"user_id"`
}
