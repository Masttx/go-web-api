package api

import (
	"net/http"
	"projetoinfiel/internal/dto"
	"projetoinfiel/internal/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AreaAPI struct {
	areaRepository     repositories.AreaRepository
	areaUserRepository repositories.AreaUserRepository
}

func NewAreaAPI(areaRepository repositories.AreaRepository, areaUserRepository repositories.AreaUserRepository) *AreaAPI {
	return &AreaAPI{
		areaRepository:     areaRepository,
		areaUserRepository: areaUserRepository,
	}
}

func (r *AreaAPI) Create(c echo.Context) error {
	req := new(dto.CreateAreaReq)

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// if err := c.Validate(req); err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	// }

	areaCreated, err := r.areaRepository.Create(req.Name, req.Description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, areaCreated)
}

func (r *AreaAPI) Update(c echo.Context) error {
	req := new(dto.UpdateAreaReq)

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid area id"})
	}
	req.ID = id

	err = r.areaRepository.Update(req.ID, req.Name, req.Description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, true)
}

func (r *AreaAPI) Read(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid area id"})
	}

	user, err := r.areaRepository.Read(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := dto.ReadAreaResponse{
		ID:          user.ID,
		Name:        user.Name,
		Description: user.Description,
	}

	return c.JSON(http.StatusCreated, response)
}

func (r *AreaAPI) AddUser(c echo.Context) error {
	req := new(dto.AddUserReq)

	// Bind parameters from URL and JSON body
	_ = c.Bind(req)

	// Retrieve area_id from URL path fallback
	if req.AreaID == 0 {
		areaIDStr := c.Param("area_id")
		if areaIDStr == "" {
			areaIDStr = c.Param("id")
		}
		if areaIDStr != "" {
			areaID, err := strconv.ParseInt(areaIDStr, 10, 64)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid area id in URL"})
			}
			req.AreaID = areaID
		}
	}

	// Retrieve user_id from URL path fallback
	if req.UserID == 0 {
		userIDStr := c.Param("user_id")
		if userIDStr != "" {
			userID, err := strconv.ParseInt(userIDStr, 10, 64)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id in URL"})
			}
			req.UserID = userID
		}
	}

	// Validate inputs
	if req.AreaID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing area id"})
	}
	if req.UserID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing user id"})
	}

	result, err := r.areaUserRepository.Create(req.AreaID, req.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func (r *AreaAPI) ListUsersByArea(c echo.Context) error {
	idStr := c.Param("area_id")

	area_id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid area id"})
	}

	users, err := r.areaUserRepository.ListUsersByArea(area_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, users)
}

func (r *AreaAPI) DeleteUser(c echo.Context) error {
	req := new(dto.DeleteUserReq)

	// Bind parameters from URL and JSON body
	_ = c.Bind(req)

	// Retrieve area_id from URL path fallback
	if req.AreaID == 0 {
		areaIDStr := c.Param("area_id")
		if areaIDStr == "" {
			areaIDStr = c.Param("id")
		}
		if areaIDStr != "" {
			areaID, err := strconv.ParseInt(areaIDStr, 10, 64)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid area id in URL"})
			}
			req.AreaID = areaID
		}
	}

	// Retrieve user_id from URL path fallback
	if req.UserID == 0 {
		userIDStr := c.Param("user_id")
		if userIDStr != "" {
			userID, err := strconv.ParseInt(userIDStr, 10, 64)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id in URL"})
			}
			req.UserID = userID
		}
	}

	// Validate inputs
	if req.AreaID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing area id"})
	}
	if req.UserID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing user id"})
	}

	err := r.areaUserRepository.Delete(req.AreaID, req.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, true)
}
