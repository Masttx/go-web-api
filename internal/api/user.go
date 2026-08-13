package api

import (
	"log"
	"net/http"
	"projetoinfiel/internal/dto"
	"projetoinfiel/internal/pkg/security"
	"projetoinfiel/internal/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

func (r *UserAPI) Register(c echo.Context) error {
	req := new(dto.RegisterReq)

	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// VALIDAR EMAIL DE USUARIO
	exists, err := r.userRepository.ExistsUserByEmail(req.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": `Erro ao validar email`})
	}
	if exists {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": `Email ja utilizado`})
	}

	// SE PASSWORD EH DIFERENTE DE PASSWORD CONFIRM RETORNA ERRO
	if req.Password != req.ConfirmPassword {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": `Senhas diferentes`})
	}
	// A SENHA TEM QUE TER NO MINIMO 8 CARACTERES
	if len(req.Password) < 8 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": `Senha muito curta, a senha precisa ter 8 ou mais caracteres`})
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao criptografar senha do usuario"})
	}

	userCreated, err := r.userRepository.Create(req.Name, req.Email, string(passwordHash))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, userCreated)
}

func (r *UserAPI) Login(c echo.Context) error {
	// Pra logar precisa do email e senha - criar um dto de login, tem que conter email e a senha
	req := new(dto.LoginReq)

	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Procurar um usuario com esse email, se nao existir retorna erro - Criar uma query pra buscar user por e-mail
	user, err := r.userRepository.FindUserByEmail(req.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": `Email ou senha invalido`})
	}

	// Pegar a senha do usuario, criptografar ela - OK - e ver se essa senha criptografa -
	// - eh igual a do banco desse usuario do email, se for certo retorna `true` - Usar linha 51 como ref
	log.Println("userpassword: ", user.PasswordHash)

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Email ou senha invalido"})
	}

	token, err := security.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, token)
}

func (r *UserAPI) Update(c echo.Context) error {
	req := new(dto.UpdateUserReq)

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

	response := dto.ReadUserResponse{
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
