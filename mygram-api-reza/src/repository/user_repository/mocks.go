package user_repository

import (
	"final-project-h8-mygram-Reza/src/models"
	"final-project-h8-mygram-Reza/src/pkg/errs"
)

// * MockUserRepository is a mock of UserRepository interface
var (
	Login               func(user *models.User) (*models.User, errs.MessageErr)
	Register            func(user *models.User) (*models.User, errs.MessageErr)
	GetUserByIDAndEmail func(user *models.User) (*models.User, errs.MessageErr)
	UpdateUserData      func(userId uint, user *models.User) (*models.User, errs.MessageErr)
	DeleteUser          func(userId uint) errs.MessageErr
)

type mockUserRepository struct{}

func NewMockUserRepository() UserRepository {
	return &mockUserRepository{}
}

func (m *mockUserRepository) Login(user *models.User) (*models.User, errs.MessageErr) {
	return Login(user)
}

func (m *mockUserRepository) Register(user *models.User) (*models.User, errs.MessageErr) {
	return Register(user)
}

func (m *mockUserRepository) GetUserByIDAndEmail(user *models.User) (*models.User, errs.MessageErr) {
	return nil, nil
}

func (m *mockUserRepository) UpdateUserData(userId uint, user *models.User) (*models.User, errs.MessageErr) {
	return UpdateUserData(userId, user)
}

func (m *mockUserRepository) DeleteUser(userId uint) errs.MessageErr {
	return DeleteUser(userId)
}
