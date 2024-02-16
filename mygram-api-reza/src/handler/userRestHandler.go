package handler

import (
	"final-project-h8-mygram-Reza/src/models"
	"final-project-h8-mygram-Reza/src/pkg/helpers"
	"final-project-h8-mygram-Reza/src/service"
	"final-project-h8-mygram-Reza/src/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userRestHandler struct {
	userService service.UserService
}

func NewUserRestHandler(userService service.UserService) *userRestHandler {
	return &userRestHandler{userService: userService}
}

// Login godoc
// @Summary Login into existing account
// @Tags users
// @Description Login into your user account and get access token with jwt
// @ID login-users
// @Accept  json
// @Produce json
// @Param RequestBody body structs.LoginRequest true "Login request body json"
// @Success 200 {object} structs.LoginResponse
// @Router /user/login [post]
func (u *userRestHandler) Login(c *gin.Context) {
	var userRequest structs.LoginRequest
	var err error

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		err = c.ShouldBindJSON(&userRequest)
	} else {
		err = c.ShouldBind(&userRequest)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	token, err2 := u.userService.Login(&userRequest)

	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}
	c.JSON(http.StatusCreated, token)
}

// Register godoc
// @Summary Register new user account
// @Tags users
// @Description Register a new user
// @ID register-users
// @Accept  json
// @Produce json
// @Param RequestBody body structs.RegisterRequest true "Register request body json"
// @Success 201 {object} structs.RegisterResponse
// @Router /user/register [post]
func (u *userRestHandler) Register(c *gin.Context) {
	var userRequest structs.RegisterRequest
	var err error

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		err = c.ShouldBindJSON(&userRequest)
	} else {
		err = c.ShouldBind(&userRequest)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	fmt.Println("User =>", &userRequest)
	successMessage, err2 := u.userService.Register(&userRequest)

	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Message(),
			"message": err2.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, successMessage)
}

// UpdateUserData godoc
// @Summary Update user's email and username
// @Tags users
// @Description Update user data
// @ID update-user-data
// @Accept  json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body structs.UpdateUserDataRequest true "Update user request body json"
// @Success 200 {object} structs.UpdateUserDataResponse
// @Router /user [put]
func (u *userRestHandler) UpdateUserData(c *gin.Context) {
	var updateUserDataRequest structs.UpdateUserDataRequest
	var err error
	var userData models.User

	contentType := helpers.GetContentType(c)
	if contentType == helpers.AppJSON {
		err = c.ShouldBindJSON(&updateUserDataRequest)
	} else {
		err = c.ShouldBind(&updateUserDataRequest)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	if value, ok := c.MustGet("userData").(models.User); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}

	response, err2 := u.userService.UpdateUserData(userData.ID, &updateUserDataRequest)
	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteUserData godoc
// @Summary Delete user's account
// @Tags users
// @Description Delete user data
// @ID delete-user
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} structs.DeleteUserResponse
// @Router /user [delete]
func (u *userRestHandler) DeleteUser(c *gin.Context) {
	var userData models.User
	if value, ok := c.MustGet("userData").(models.User); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}

	response, err2 := u.userService.DeleteUser(userData.ID)
	if err2 != nil {
		c.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
