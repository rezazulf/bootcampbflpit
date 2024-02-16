package comment_repository

import (
	"final-project-h8-mygram-Reza/src/models"
	"final-project-h8-mygram-Reza/src/pkg/errs"
)

type CommentRepository interface {
	PostComment(comment *models.Comment) (*models.Comment, errs.MessageErr)
	GetAllComments() ([]*models.Comment, errs.MessageErr)
	GetCommentByID(commentID uint) (*models.Comment, errs.MessageErr)
	EditCommentData(commentID uint, comment *models.Comment) (*models.Comment, errs.MessageErr)
	DeleteComment(commentID uint) errs.MessageErr
}