package handler

import (
	"final-project-h8-mygram-Reza/src/models"
	"final-project-h8-mygram-Reza/src/pkg/helpers"
	"final-project-h8-mygram-Reza/src/service"
	"final-project-h8-mygram-Reza/src/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentRestHandler struct {
	commentService service.CommentService
}

func NewCommentRestHandler(commentService service.CommentService) *commentRestHandler {
	return &commentRestHandler{commentService: commentService}
}

// CODED BY MURE
// PostComment godoc
// @Summary Post new comment on photo
// @Tags comments
// @Description Post new comment on photo
// @ID post-comment
// @Accept  json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body structs.CommentRequest true "Post comment request body json"
// @Success 201 {object} structs.CommentResponse
// @Router /comment [post]
func (c *commentRestHandler) PostComment(ctx *gin.Context) {
	var commentRequest structs.CommentRequest
	var err error
	var userData models.User

	contentType := helpers.GetContentType(ctx)
	if contentType == helpers.AppJSON {
		err = ctx.ShouldBindJSON(&commentRequest)
	} else {
		err = ctx.ShouldBind(&commentRequest)
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	if value, ok := ctx.MustGet("userData").(models.User); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}

	comment, err2 := c.commentService.PostComment(userData.ID, &commentRequest)
	if err2 != nil {
		ctx.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

// GetAllComments godoc
// @Summary Get all comments
// @Tags comments
// @Description Get all comments
// @ID get-all-comments
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} structs.GetCommentResponse
// @Router /comment [get]
func (c *commentRestHandler) GetAllComments(ctx *gin.Context) {
	var userData models.User
	if value, ok := ctx.MustGet("userData").(models.User); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}
	_ = userData

	comments, err := c.commentService.GetAllComments()
	if err != nil {
		ctx.JSON(err.Status(), gin.H{
			"error":   err.Error(),
			"message": err.Message(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

// UpdateComment godoc
// @Summary Update existing comment
// @Tags comments
// @Description Update comment
// @ID update-comment
// @Accept  json
// @Produce json
// @Param commentID path uint true "comments's id"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body structs.UpdateCommentRequest true "Edit photo request body json"
// @Success 200 {object} structs.UpdateCommentResponse
// @Router /comment/{commentID} [put]
func (c *commentRestHandler) UpdateComment(ctx *gin.Context) {
	var commentRequest structs.UpdateCommentRequest
	var userData models.User
	var err error

	contentType := helpers.GetContentType(ctx)
	if contentType == helpers.AppJSON {
		err = ctx.ShouldBindJSON(&commentRequest)
	} else {
		err = ctx.ShouldBind(&commentRequest)
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": err.Error(),
		})
		return
	}

	if value, ok := ctx.MustGet("userData").(models.User); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}
	_ = userData

	commentIdParam, err := helpers.GetParamId(ctx, "commentID")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err_message": "invalid params",
		})
		return
	}

	comment, err2 := c.commentService.EditCommentData(commentIdParam, &commentRequest)
	if err2 != nil {
		ctx.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// DeleteComment godoc
// @Summary Delete existing comment
// @Tags comments
// @Description Delete comment
// @ID delete-comment
// @Produce json
// @Param commentID path uint true "comments's id"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} structs.DeleteCommentResponse
// @Router /comment/{commentID} [delete]
func (c *commentRestHandler) DeleteComment(ctx *gin.Context) {
	var userData models.User
	if value, ok := ctx.MustGet("userData").(models.User); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"err_message": "unauthorized",
		})
		return
	} else {
		userData = value
	}
	_ = userData

	commentIdParam, err := helpers.GetParamId(ctx, "commentID")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err_message": "invalid params",
		})
		return
	}

	res, err2 := c.commentService.DeleteComment(commentIdParam)
	if err2 != nil {
		ctx.JSON(err2.Status(), gin.H{
			"error":   err2.Error(),
			"message": err2.Message(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
