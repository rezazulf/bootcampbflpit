package service

import (
	"final-project-h8-mygram-Reza/src/models"
	"final-project-h8-mygram-Reza/src/pkg/errs"
	"final-project-h8-mygram-Reza/src/repository/user_repository"
	"final-project-h8-mygram-Reza/src/structs"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_Register_Success(t *testing.T) {
	// Arrange
	userRepositoryMock := user_repository.NewMockUserRepository()
	userService := NewUserService(userRepositoryMock)

	requestBody := &structs.RegisterRequest{
		Email:    "test@gmail.com",
		Password: "123456",
		Username: "test",
		Age:      20,
	}

	user_repository.Register = func(userPayload *models.User) (*models.User, errs.MessageErr) {
		response := &models.User{
			Username: userPayload.Username,
			Email:    userPayload.Email,
			Password: userPayload.Password,
			Age:      userPayload.Age,
		}

		return response, nil
	}

	// Act
	user, err := userService.Register(requestBody)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.Equal(t, requestBody.Email, user.Email)
	assert.Equal(t, requestBody.Username, user.Username)
	assert.Equal(t, requestBody.Age, user.Age)
}

func TestUserService_Register_BadRequest(t *testing.T) {
	// Arrange
	userRepositoryMock := user_repository.NewMockUserRepository()
	userService := NewUserService(userRepositoryMock)

	tests := []struct {
		name        string
		requestBody *structs.RegisterRequest
		errMsg      string
		status      int
		errError    string
	}{
		{
			name: "Empty email",
			requestBody: &structs.RegisterRequest{
				Email:    "",
				Password: "123456",
				Username: "test",
				Age:      20,
			},
			errMsg:   "Email is required",
			status:   http.StatusBadRequest,
			errError: "bad_request",
		},
		{
			name: "Empty password",
			requestBody: &structs.RegisterRequest{
				Email:    "test@gmail.com",
				Password: "",
				Username: "test",
				Age:      20,
			},
			errMsg:   "Password is required",
			status:   http.StatusBadRequest,
			errError: "bad_request",
		},
		{
			name: "Empty username",
			requestBody: &structs.RegisterRequest{
				Email:    "test@gmail.com",
				Password: "123456",
				Username: "",
				Age:      20,
			},
			errMsg:   "Username is required",
			status:   http.StatusBadRequest,
			errError: "bad_request",
		},
		{
			name: "Empty age",
			requestBody: &structs.RegisterRequest{
				Email:    "test@gmail.com",
				Password: "123456",
				Username: "ssss",
				Age:      0,
			},
			errMsg:   "Age is required",
			status:   http.StatusBadRequest,
			errError: "bad_request",
		},
		{
			name: "Invalid email",
			requestBody: &structs.RegisterRequest{
				Email:    "test",
				Password: "123456",
				Username: "ssss",
				Age:      20,
			},
			errMsg:   "Email is not valid",
			status:   http.StatusBadRequest,
			errError: "bad_request",
		},
		{
			name: "Minimum age invalid",
			requestBody: &structs.RegisterRequest{
				Email:    "test@gmail.com",
				Password: "123456",
				Username: "ssss",
				Age:      5,
			},
			errMsg:   "Age must be between 8 and 100",
			status:   http.StatusBadRequest,
			errError: "bad_request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			user, err := userService.Register(tt.requestBody)

			// Assert
			assert.Nil(t, user)
			assert.NotNil(t, err)

			assert.Equal(t, tt.errMsg, err.Message())
			assert.Equal(t, tt.status, err.Status())
			assert.Equal(t, tt.errError, err.Error())
		})
	}
}

func TestUserService_Register_InternalServerError(t *testing.T) {
	// Arrange
	userRepositoryMock := user_repository.NewMockUserRepository()
	userService := NewUserService(userRepositoryMock)

	expectedErr := errs.NewInternalServerErrorr("something went wrong")
	user_repository.Register = func(user *models.User) (*models.User, errs.MessageErr) {
		return nil, expectedErr
	}

	requestBody := &structs.RegisterRequest{
		Email:    "test@mail.com",
		Password: "123456",
		Username: "ssss",
		Age:      20,
	}

	// Act
	user, err := userService.Register(requestBody)
	assert.NotNil(t, err)
	assert.Nil(t, user)

	assert.Equal(t, expectedErr.Error(), err.Error())
}
