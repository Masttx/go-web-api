package api

import (
	"encoding/json"
	"net/http"
	"projetoinfiel/internal/repositories"
	"strconv"
)

type AreaAPI struct {
	areaRepository repositories.AreaRepository
}

func NewAreaAPI(areaRepository repositories.AreaRepository) *AreaAPI {
	return &AreaAPI{
		areaRepository: areaRepository,
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

func (r *AreaAPI) Create(writer http.ResponseWriter, request *http.Request) {
	req := new(CreateAreaReq)

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(err)

		return
	}

	areaCreated, err := r.areaRepository.Create(req.Name, req.Description)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(err)

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
		json.NewEncoder(writer).Encode(err)

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
		json.NewEncoder(writer).Encode(err)

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
		json.NewEncoder(writer).Encode(err)

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
