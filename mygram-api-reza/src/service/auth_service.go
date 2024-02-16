package service

import (
	"final-project-h8-mygram-Reza/src/models"
	"final-project-h8-mygram-Reza/src/pkg/helpers"
	"final-project-h8-mygram-Reza/src/repository/comment_repository"
	"final-project-h8-mygram-Reza/src/repository/photo_repository"
	"final-project-h8-mygram-Reza/src/repository/social_media_repository"
	"final-project-h8-mygram-Reza/src/repository/user_repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	PhotoAuthorization() gin.HandlerFunc
	CommentAuthorization() gin.HandlerFunc
	SocialMediaAuthorization() gin.HandlerFunc
}

type authService struct {
	userRepository        user_repository.UserRepository
	photoRepository       photo_repository.PhotoRepository
	commentRepository     comment_repository.CommentRepository
	socialMediaRepository social_media_repository.SocialMediaRepository
}

func NewAuthService(userRepository user_repository.UserRepository, photoRepository photo_repository.PhotoRepository, commentRepository comment_repository.CommentRepository, socialMediaRepository social_media_repository.SocialMediaRepository) AuthService {
	return &authService{
		userRepository:        userRepository,
		photoRepository:       photoRepository,
		commentRepository:     commentRepository,
		socialMediaRepository: socialMediaRepository,
	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var user *models.User = &models.User{}

		// Get token from header
		tokenStr := ctx.Request.Header.Get("Authorization")
		err := user.VerifyToken(tokenStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error_message": err.Error(),
			})
			return
		}

		_, err = a.userRepository.GetUserByIDAndEmail(user)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error_message": err.Error(),
			})
			return
		}

		ctx.Set("userData", *user)
		ctx.Next()
	})
}

func (a *authService) PhotoAuthorization() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var userData models.User

		if value, ok := ctx.MustGet("userData").(models.User); !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error_message": "unauthorized",
			})
			return
		} else {
			userData = value
		}

		photoIdParam, err := helpers.GetParamId(ctx, "photoID")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"err_message": "invalid params",
			})
			return
		}

		photo, err := a.photoRepository.GetPhotoByID(photoIdParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"err_message": "photo not found",
			})
			return
		}

		if photo.UserID != userData.ID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"err_message": "forbidden access",
			})
			return
		}

		ctx.Next()
	})
}

func (a *authService) CommentAuthorization() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var userData models.User

		if value, ok := ctx.MustGet("userData").(models.User); !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error_message": "unauthorized",
			})
			return
		} else {
			userData = value
		}

		commentIdParam, err := helpers.GetParamId(ctx, "commentID")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"err_message": "invalid params",
			})
			return
		}

		comment, err := a.commentRepository.GetCommentByID(commentIdParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"err_message": "comment not found",
			})
			return
		}

		if comment.UserID != userData.ID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"err_message": "forbidden access",
			})
			return
		}
	})
}

func (a *authService) SocialMediaAuthorization() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var userData models.User

		if value, ok := ctx.MustGet("userData").(models.User); !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error_message": "unauthorized",
			})
			return
		} else {
			userData = value
		}

		socialMediaIdParam, err := helpers.GetParamId(ctx, "socialMediaID")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"err_message": "invalid params",
			})
			return
		}

		socialMedia, err := a.socialMediaRepository.GetSocialMediaByID(socialMediaIdParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"err_message": "social media not found",
			})
			return
		}

		if socialMedia.UserID != userData.ID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"err_message": "forbidden access",
			})
			return
		}
	})
}
