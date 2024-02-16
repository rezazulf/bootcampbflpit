package user_pg

import (
	"final-project-h8-mygram-Reza/src/models"
	"final-project-h8-mygram-Reza/src/pkg/errs"
	"final-project-h8-mygram-Reza/src/repository/user_repository"

	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) user_repository.UserRepository {
	return &userPG{db: db}
}

func (u *userPG) GetUserByIDAndEmail(userPayload *models.User) (*models.User, errs.MessageErr) {
	user := models.User{}

	err := u.db.Where("email = ? AND id = ?", userPayload.Email, userPayload.ID).Take(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("User not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &user, nil
}

func (u *userPG) Login(userPayload *models.User) (*models.User, errs.MessageErr) {
	user := models.User{}

	err := u.db.Where("email = ?", userPayload.Email).Take(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("User not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &user, nil
}

func (u *userPG) Register(userPayload *models.User) (*models.User, errs.MessageErr) {
	user := models.User{}

	err := u.db.Create(userPayload).Error
	if err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	err = u.db.Where("email = ?", userPayload.Email).Take(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("User not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &user, nil
}

func (u *userPG) UpdateUserData(userId uint, userPayload *models.User) (*models.User, errs.MessageErr) {
	user := models.User{}

	err := u.db.Model(user).Where("id = ?", userId).Updates(userPayload).Take(&user).Error
	if err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &user, nil
}

func (u *userPG) DeleteUser(userId uint) errs.MessageErr {
	user := models.User{}

	err := u.db.Where("id = ?", userId).Delete(&user).Error
	if err != nil {
		return errs.NewInternalServerErrorr("Something went wrong")
	}

	return nil
}
