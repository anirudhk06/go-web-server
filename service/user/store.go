package user

import (
	"github.com/anirudhk06/go-web-server/types"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		DB: db,
	}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	var user types.User
	err := s.DB.First(&user, "email = ?", email).Error
	return &user, err
}

func (s *Store) CreateUser(user types.User) error {
	err := s.DB.Create(&user).Error
	return err

}

func (s *Store) GetUserByID(ID int) (*types.User, error) {
	return nil, nil
}

func (s *Store) FindUsers() ([]types.User, error) {

	var users []types.User

	err := s.DB.Find(&users).Error
	return users, err

}
