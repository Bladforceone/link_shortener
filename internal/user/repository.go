package user

import "go_pro_api/pkg/db"

//Create
//FindByEmail

type Repository struct {
	DB *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Create(user *User) (*User, error) {
	result := r.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	var user User
	result := r.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
