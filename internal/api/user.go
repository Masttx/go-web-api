package api

import (
	"encoding/json"
	"net/http"
	"projetoinfiel/internal/repositories"
	"strconv"
)

type UserAPI struct {
	userRepository repositories.UserRepository
}

func NewUserAPI(userRepository repositories.UserRepository) *UserAPI {
	return &UserAPI{
		userRepository: userRepository,
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

func (r *UserAPI) Create(writer http.ResponseWriter, request *http.Request) {
	req := new(CreateUserReq)

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(err)

		return
	}

	userCreated, err := r.userRepository.Create(req.Name, req.Email)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(err)

		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(userCreated)
}

func (r *UserAPI) Update(writer http.ResponseWriter, request *http.Request) {
	req := new(UpdateUserReq)

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
		json.NewEncoder(writer).Encode(map[string]string{"error": "invalid user id"})

		return
	}
	req.ID = id

	err = r.userRepository.Update(req.ID, req.Name, req.Email)
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

func (r *UserAPI) Read(writer http.ResponseWriter, request *http.Request) {
	idStr := request.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": "invalid user id"})

		return
	}

	user, err := r.userRepository.Read(id)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(err)

		return
	}

	response := ReadUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(response)
}

func (r *UserAPI) UpdateUserArea(writer http.ResponseWriter, request *http.Request) {
	req := new(UpdateUserAreaReq)

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
		json.NewEncoder(writer).Encode(map[string]string{"error": "invalid user id"})

		return
	}
	req.ID = id

	err = r.userRepository.UpdateUserArea(req.ID, req.AreaID)
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

func (r *UserAPI) ListByArea(writer http.ResponseWriter, request *http.Request) {
	idStr := request.PathValue("area_id")

	areaID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": "invalid area_id"})

		return
	}

	users, err := r.userRepository.ListByArea(areaID)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}

	response := users

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(response)
}
