package auth

// RegisterRequest - структура запроса на регистрацию нового пользователя.
// Содержит поля, необходимые для создания нового аккаунта
type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// RegisterResponse - структура ответа после успешной регистрации.
// Содержит JWT токен для аутентификации пользователя
type RegisterResponse struct {
	Token string `json:"token"`
}

// LoginRequest - структура запроса на вход в систему.
// Содержит учетные данные пользователя
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse - структура ответа после успешного входа.
// Содержит JWT токен для аутентификации пользователя
type LoginResponse struct {
	Token string `json:"token"`
}
