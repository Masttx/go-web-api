package api

import (
	"net/http"
	"projetoinfiel/internal/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserAPI struct {
	userRepository     repositories.UserRepository
	areaUserRepository repositories.AreaUserRepository
}

func NewUserAPI(userRepository repositories.UserRepository, areaUserRepository repositories.AreaUserRepository) *UserAPI {
	return &UserAPI{
		userRepository:     userRepository,
		areaUserRepository: areaUserRepository,
	}
}

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

type ListAreaByUsersReq struct {
	AreaID int64 `param:"area_id"`
}

func (r *UserAPI) Create(c echo.Context) error {
	req := new(CreateUserReq)

	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	userCreated, err := r.userRepository.Create(req.Name, req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, userCreated)
}

func (r *UserAPI) Update(c echo.Context) error {
	req := new(UpdateUserReq)

	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id"})
	}
	req.ID = id

	err = r.userRepository.Update(req.ID, req.Name, req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, true)
}

func (r *UserAPI) Read(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id"})
	}

	user, err := r.userRepository.Read(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := ReadUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return c.JSON(http.StatusCreated, response)
}

func (r *UserAPI) ListAreaByUsers(c echo.Context) error {
	idStr := c.Param("user_id")

	user_id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid users id"})
	}

	users, err := r.areaUserRepository.ListAreasByUser(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}
