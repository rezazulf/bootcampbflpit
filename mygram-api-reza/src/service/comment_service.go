package service

import (
	"final-project-h8-mygram-Reza/src/models"
	"final-project-h8-mygram-Reza/src/pkg/errs"
	"final-project-h8-mygram-Reza/src/pkg/helpers"
	"final-project-h8-mygram-Reza/src/repository/comment_repository"
	"final-project-h8-mygram-Reza/src/repository/photo_repository"
	"final-project-h8-mygram-Reza/src/structs"
)

type CommentService interface {
	PostComment(userID uint, commentPayload *structs.CommentRequest) (*structs.CommentResponse, errs.MessageErr)
	GetAllComments() ([]*structs.GetCommentResponse, errs.MessageErr)
	EditCommentData(commentID uint, commentPayload *structs.UpdateCommentRequest) (*structs.UpdateCommentResponse, errs.MessageErr)
	DeleteComment(commentID uint) (*structs.DeleteCommentResponse, errs.MessageErr)
}

type commentService struct {
	commentRepository comment_repository.CommentRepository
	photoRepository   photo_repository.PhotoRepository
}

func NewCommentService(commentRepository comment_repository.CommentRepository, photoRepository photo_repository.PhotoRepository) CommentService {
	return &commentService{commentRepository: commentRepository, photoRepository: photoRepository}
}

func (c *commentService) PostComment(userID uint, commentPayload *structs.CommentRequest) (*structs.CommentResponse, errs.MessageErr) {
	if err := helpers.ValidateStruct(commentPayload); err != nil {
		return nil, err
	}

	models := &models.Comment{
		Message: commentPayload.Message,
		PhotoID: commentPayload.PhotoID,
		UserID:  userID,
	}

	_, err := c.photoRepository.GetPhotoByID(commentPayload.PhotoID)
	if err != nil {
		return nil, err
	}

	comment, err := c.commentRepository.PostComment(models)
	if err != nil {
		return nil, err
	}

	response := &structs.CommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
	}

	return response, nil
}

func (c *commentService) GetAllComments() ([]*structs.GetCommentResponse, errs.MessageErr) {
	comments, err := c.commentRepository.GetAllComments()
	if err != nil {
		return nil, err
	}

	response := make([]*structs.GetCommentResponse, 0)
	for _, comment := range comments {
		response = append(response, comment.ToGetCommentResponseStructs())
	}

	return response, nil
}

func (c *commentService) EditCommentData(commentID uint, commentPayload *structs.UpdateCommentRequest) (*structs.UpdateCommentResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(commentPayload)
	if err != nil {
		return nil, err
	}

	models := &models.Comment{
		Message: commentPayload.Message,
	}

	comment, err := c.commentRepository.EditCommentData(commentID, models)
	if err != nil {
		return nil, err
	}

	response := &structs.UpdateCommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}

	return response, nil
}

func (c *commentService) DeleteComment(commentID uint) (*structs.DeleteCommentResponse, errs.MessageErr) {
	if err := c.commentRepository.DeleteComment(commentID); err != nil {
		return nil, err
	}

	response := &structs.DeleteCommentResponse{
		Message: "Your comment has been deleted",
	}

	return response, nil
}
