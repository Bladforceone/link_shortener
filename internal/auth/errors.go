package auth

const (
	// ErrUserExist - ошибка, возникающая при попытке регистрации уже существующего пользователя
	ErrUserExist = "user exist"
	// ErrWrongCredentials - ошибка, возникающая при вводе неверных учетных данных (email/пароль)
	ErrWrongCredentials = "wrong email or password"
)
