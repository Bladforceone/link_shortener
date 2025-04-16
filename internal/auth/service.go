package auth

import (
	"errors"
	"go_pro_api/internal/user"
	"go_pro_api/pkg/di"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository di.IUserRepository
}

// NewAuthService создает новый экземпляр AuthService с инъекцией зависимостей
func NewAuthService(userRepository di.IUserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

// Register регистрирует нового пользователя в системе.
// Принимает email, password и name.
// Возвращает email зарегистрированного пользователя или ошибку.
func (s *AuthService) Register(email, password, name string) (string, error) {
	existedUser, _ := s.UserRepository.FindByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrUserExist)
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	u := &user.User{
		Email:    email,
		Password: string(hashPassword),
		Name:     name,
	}
	_, err = s.UserRepository.Create(u)
	if err != nil {
		return "", err
	}
	return u.Email, nil
}

// Login аутентифицирует пользователя в системе.
// Принимает email и password.
// Возвращает email аутентифицированного пользователя или ошибку.
func (s *AuthService) Login(email, password string) (string, error) {
	existedUser, _ := s.UserRepository.FindByEmail(email)
	if existedUser == nil {
		return "", errors.New(ErrWrongCredentials)
	}
	err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrWrongCredentials)
	}
	return email, nil
}
