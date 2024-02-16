package comment_pg

import (
	"final-project-h8-mygram-Reza/src/models"
	"final-project-h8-mygram-Reza/src/pkg/errs"
	"final-project-h8-mygram-Reza/src/repository/comment_repository"
	"fmt"

	"gorm.io/gorm"
)

type commentPG struct {
	db *gorm.DB
}

func NewCommentPG(db *gorm.DB) comment_repository.CommentRepository {
	return &commentPG{db: db}
}

func (c *commentPG) PostComment(commentPayload *models.Comment) (*models.Comment, errs.MessageErr) {
	comment := models.Comment{}
	comment.UserID = commentPayload.UserID

	if err := c.db.Model(&comment).Create(&commentPayload).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	if err := c.db.Last(&comment).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &comment, nil
}

func (c *commentPG) GetAllComments() ([]*models.Comment, errs.MessageErr) {
	comments := []*models.Comment{}

	err := c.db.Preload("User").Preload("Photo").Find(&comments).Error
	if err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return comments, nil
}

func (c *commentPG) GetCommentByID(commentID uint) (*models.Comment, errs.MessageErr) {
	comment := models.Comment{}

	if err := c.db.Where("id = ?", commentID).First(&comment).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &comment, nil
}

func (c *commentPG) EditCommentData(commentID uint, commentPayload *models.Comment) (*models.Comment, errs.MessageErr) {
	comment := models.Comment{}
	fmt.Println("Melihat payload", commentPayload)

	err := c.db.Model(&comment).Where("id = ?", commentID).Updates(&commentPayload).Take(&comment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("Comment not found")
		}
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &comment, nil
}

func (c *commentPG) DeleteComment(commentID uint) errs.MessageErr {
	comment := models.Comment{}

	if err := c.db.Where("id = ?", commentID).Delete(&comment).Error; err != nil {
		return errs.NewInternalServerErrorr("Something went wrong")
	}

	return nil
}
