package auth_test

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"go_pro_api/configs"
	"go_pro_api/internal/auth"
	"go_pro_api/internal/user"
	"go_pro_api/pkg/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func bootsrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return nil, nil, err
	}
	userRepo := user.NewRepository(&db.DB{
		DB: gormDB,
	})
	handler := auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: auth.NewAuthService(userRepo),
	}
	return &handler, mock, nil
}

func TestLoginSuccess(t *testing.T) {
	handler, mock, err := bootsrap()
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("mail3@mail.ru", "$2a$10$Q39e0MmOvb4M6e6dFU5WHetEYx76nGwOLYomONSU71FwZ/6Cic/tm")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	if err != nil {
		t.Fatal(err)
		return
	}
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "m@mail.ru",
		Password: "5242",
	})
	reader := bytes.NewReader(data)
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Login()(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got: %d, expected %d", wr.Code, 200)
	}
}

func TestRegisterSuccess(t *testing.T) {
	handler, mock, err := bootsrap()
	rows := sqlmock.NewRows([]string{"email", "password", "name"})
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()
	if err != nil {
		t.Fatal(err)
		return
	}
	data, _ := json.Marshal(&auth.RegisterRequest{
		Email:    "m@mail.ru",
		Password: "5242",
		Name:     "Вова",
	})
	reader := bytes.NewReader(data)
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/register", reader)
	handler.Register()(wr, req)
	if wr.Code != http.StatusCreated {
		t.Errorf("got: %d, expected %d", wr.Code, 201)
	}
}
