package repositories

import (
	"database/sql"
	"log"
	"user_auth/models"
)

type IUserRepository interface {
	FindByUsername(username string) (*models.User, error)
	Create(user models.User) error
	UpdateUserDetails(user models.User) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, username, password FROM users WHERE username = ?"
	err := r.db.QueryRow(query, username).Scan(&user.Id, &user.Username, &user.Password)
	return user, err
}

func (r *UserRepository) Create(user models.User) error {
	query := "INSERT INTO users (username,password) VALUES (?,?)"
	_, err := r.db.Exec(query, user.Username, user.Password)
	return err
}

func (r *UserRepository) UpdateUserDetails(user models.User) error {
	query := "UPDATE users SET name=?, lastname=?, dni=?, address=?, postal_code=?, phone=? WHERE id=?"
	result, err := r.db.Exec(query, user.Name, user.Lastname, user.DNI, user.Address, user.PostalCode, user.Phone, user.Id)
	log.Printf("UpdateUserDetails result: %v", result)
	return err
}
