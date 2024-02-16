package service

import (
	"fmt"

	"final-project-h8-mygram-Reza/src/models"
	"final-project-h8-mygram-Reza/src/pkg/errs"
	"final-project-h8-mygram-Reza/src/pkg/helpers"
	"final-project-h8-mygram-Reza/src/repository/user_repository"
	"final-project-h8-mygram-Reza/src/structs"
)

type UserService interface {
	Login(userPayload *structs.LoginRequest) (*structs.LoginResponse, errs.MessageErr)
	Register(userPayload *structs.RegisterRequest) (*structs.RegisterResponse, errs.MessageErr)
	UpdateUserData(userId uint, userPayload *structs.UpdateUserDataRequest) (*structs.UpdateUserDataResponse, errs.MessageErr)
	DeleteUser(userId uint) (*structs.DeleteUserResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Login(userPayload *structs.LoginRequest) (*structs.LoginResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(userPayload)
	if err != nil {
		return nil, err
	}

	payload := &models.User{
		Email: userPayload.Email,
	}

	user, err := u.userRepo.Login(payload)
	if err != nil {
		return nil, err
	}

	validPassword := user.VerifyPassword(userPayload.Password)
	if !validPassword {
		return nil, errs.NewUnAuthorized("Invalid email or password")
	}

	token := user.GenerateToken()

	response := &structs.LoginResponse{
		AccessToken: token,
	}

	return response, nil
}

func (u *userService) Register(userPayload *structs.RegisterRequest) (*structs.RegisterResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(userPayload)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: userPayload.Username,
		Email:    userPayload.Email,
		Password: userPayload.Password,
		Age:      userPayload.Age,
	}

	err = user.HashPass()
	if err != nil {
		return nil, err
	}

	user, err = u.userRepo.Register(user)
	if err != nil {
		return nil, err
	}

	response := &structs.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}

	return response, nil
}

func (u *userService) UpdateUserData(userId uint, userPayload *structs.UpdateUserDataRequest) (*structs.UpdateUserDataResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(userPayload)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    userPayload.Email,
		Username: userPayload.Username,
	}

	user, err = u.userRepo.UpdateUserData(userId, user)
	if err != nil {
		return nil, err
	}
	fmt.Println("Data user", user.Email)

	response := &structs.UpdateUserDataResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (u *userService) DeleteUser(userId uint) (*structs.DeleteUserResponse, errs.MessageErr) {
	err := u.userRepo.DeleteUser(userId)
	if err != nil {
		return nil, err
	}

	response := &structs.DeleteUserResponse{
		Message: "User data has been successfully deleted",
	}

	return response, nil
}
