package api

import (
	"encoding/json"
	"net/http"
	"projetoinfiel/internal/repositories"
	"strconv"
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

func (r *AreaAPI) Create(writer http.ResponseWriter, request *http.Request) {
	req := new(CreateAreaReq)

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}

	areaCreated, err := r.areaRepository.Create(req.Name, req.Description)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(areaCreated)
}

func (r *AreaAPI) Update(writer http.ResponseWriter, request *http.Request) {
	req := new(UpdateAreaReq)

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}

	idStr := request.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": "invalid area id"})

		return
	}
	req.ID = id

	err = r.areaRepository.Update(req.ID, req.Name, req.Description)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(true)
}

func (r *AreaAPI) Read(writer http.ResponseWriter, request *http.Request) {
	idStr := request.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": "invalid area id"})

		return
	}

	user, err := r.areaRepository.Read(id)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}

	response := ReadAreaResponse{
		ID:          user.ID,
		Name:        user.Name,
		Description: user.Description,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(response)
}

func (r *AreaAPI) AddUser(writer http.ResponseWriter, request *http.Request) {
	req := new(AddUserReq)

	// Try to decode the request body if present
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil && err.Error() != "EOF" {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": "invalid json body: " + err.Error()})
		return
	}

	// Retrieve area_id from URL path
	areaIDStr := request.PathValue("area_id")
	if areaIDStr == "" {
		areaIDStr = request.PathValue("id")
	}
	if areaIDStr != "" {
		areaID, err := strconv.ParseInt(areaIDStr, 10, 64)
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(map[string]string{"error": "invalid area id in URL"})
			return
		}
		req.AreaID = areaID
	}

	// Retrieve user_id from URL path if present
	userIDStr := request.PathValue("user_id")
	if userIDStr != "" {
		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(map[string]string{"error": "invalid user id in URL"})
			return
		}
		req.UserID = userID
	}

	// Validate inputs
	if req.AreaID == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": "missing area id"})
		return
	}
	if req.UserID == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": "missing user id"})
		return
	}

	result, err := r.areaUserRepository.Create(req.AreaID, req.UserID)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(result)
}

func (r *AreaAPI) ListUsersByArea(writer http.ResponseWriter, request *http.Request) {
	idStr := request.PathValue("area_id")

	area_id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": "invalid area id"})

		return
	}

	users, err := r.areaUserRepository.ListUsersByArea(area_id)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(users)
}

func (r *AreaAPI) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	req := new(DeleteUserReq)

	// Try to decode the request body if present
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil && err.Error() != "EOF" {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": "invalid json body: " + err.Error()})
		return
	}

	// Retrieve area_id from URL path
	areaIDStr := request.PathValue("area_id")
	if areaIDStr == "" {
		areaIDStr = request.PathValue("id")
	}
	if areaIDStr != "" {
		areaID, err := strconv.ParseInt(areaIDStr, 10, 64)
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(map[string]string{"error": "invalid area id in URL"})
			return
		}
		req.AreaID = areaID
	}

	// Retrieve user_id from URL path if present
	userIDStr := request.PathValue("user_id")
	if userIDStr != "" {
		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(map[string]string{"error": "invalid user id in URL"})
			return
		}
		req.UserID = userID
	}

	// Validate inputs
	if req.AreaID == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": "missing area id"})
		return
	}
	if req.UserID == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": "missing user id"})
		return
	}

	err = r.areaUserRepository.Delete(req.AreaID, req.UserID)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(true)
}
