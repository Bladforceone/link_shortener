package auth

import (
	"go_pro_api/configs"
	"go_pro_api/pkg/jwt"
	"go_pro_api/pkg/request"
	"go_pro_api/pkg/response"
	"net/http"
)

// AuthHandlerDeps содержит зависимости для AuthHandler
type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

// AuthHandler обрабатывает HTTP запросы для аутентификации
type AuthHandler struct {
	*configs.Config
	*AuthService
}

// NewAuthHandler создает новый обработчик аутентификации и регистрирует его маршруты
func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{deps.Config, deps.AuthService}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

// Login возвращает обработчик для входа пользователя
func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}

		email, err := h.AuthService.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(h.Auth.Secret).Create(jwt.JWTData{Email: email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := LoginResponse{
			Token: token,
		}

		response.JSON(w, data, http.StatusOK)
	}
}

// Register возвращает обработчик для регистрации нового пользователя
func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}

		email, err := h.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			return
		}

		token, err := jwt.NewJWT(h.Auth.Secret).Create(jwt.JWTData{Email: email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := RegisterResponse{
			Token: token,
		}

		response.JSON(w, data, http.StatusCreated)
	}
}
