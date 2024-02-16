package router

import (
	"final-project-h8-mygram-Reza/docs"
	"final-project-h8-mygram-Reza/src/database"
	"final-project-h8-mygram-Reza/src/handler"
	"final-project-h8-mygram-Reza/src/repository/comment_repository/comment_pg"
	"final-project-h8-mygram-Reza/src/repository/photo_repository/photo_pg"
	"final-project-h8-mygram-Reza/src/repository/social_media_repository/social_media_pg"
	"final-project-h8-mygram-Reza/src/repository/user_repository/user_pg"
	"final-project-h8-mygram-Reza/src/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var (
	PORT = os.Getenv("PORT")
)

func StartApp() {
	database.StartDB()

	// ! Inject all the dependencies here
	db := database.GetDB()
	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userRestHandler := handler.NewUserRestHandler(userService)

	photoRepo := photo_pg.NewPhotoPG(db)
	photoService := service.NewPhotoService(photoRepo)
	photoRestHandler := handler.NewPhotoRestHandler(photoService)

	commentRepo := comment_pg.NewCommentPG(db)
	commentService := service.NewCommentService(commentRepo, photoRepo)
	commentRestHandler := handler.NewCommentRestHandler(commentService)

	socialMediaRepo := social_media_pg.NewSocialMediaPG(db)
	socialMediaService := service.NewSocialMediaService(socialMediaRepo)
	socialMediaRestHandler := handler.NewSocialMediaRestHandler(socialMediaService)

	authService := service.NewAuthService(userRepo, photoRepo, commentRepo, socialMediaRepo)

	// ! Routing
	route := gin.Default()

	userRoute := route.Group("/user")
	{
		userRoute.POST("/login", userRestHandler.Login)
		userRoute.POST("/register", userRestHandler.Register)
		userRoute.PUT("/", authService.Authentication(), userRestHandler.UpdateUserData)
		userRoute.DELETE("/", authService.Authentication(), userRestHandler.DeleteUser)
	}

	photoRoute := route.Group("/photo")
	{
		photoRoute.Use(authService.Authentication())
		photoRoute.POST("/", photoRestHandler.PostPhoto)
		photoRoute.GET("/", photoRestHandler.GetAllPhotos)
		photoRoute.PUT("/:photoID", authService.PhotoAuthorization(), photoRestHandler.UpdatePhoto)
		photoRoute.DELETE("/:photoID", authService.PhotoAuthorization(), photoRestHandler.DeletePhoto)
	}

	commentRoute := route.Group("/comment")
	{
		commentRoute.Use(authService.Authentication())
		commentRoute.POST("/", commentRestHandler.PostComment)
		commentRoute.GET("/", commentRestHandler.GetAllComments)
		commentRoute.PUT("/:commentID", authService.CommentAuthorization(), commentRestHandler.UpdateComment)
		commentRoute.DELETE("/:commentID", authService.CommentAuthorization(), commentRestHandler.DeleteComment)
	}

	socialMediaRoute := route.Group("/socialmedia")
	{
		socialMediaRoute.Use(authService.Authentication())
		socialMediaRoute.POST("/", socialMediaRestHandler.AddSocialMedia)
		socialMediaRoute.GET("/", socialMediaRestHandler.GetAllSocialMedias)
		socialMediaRoute.PUT("/:socialMediaID", authService.SocialMediaAuthorization(), socialMediaRestHandler.EditSocialMediaData)
		socialMediaRoute.DELETE("/:socialMediaID", authService.SocialMediaAuthorization(), socialMediaRestHandler.DeleteSocialMedia)
	}

	// ! Docs
	docs.SwaggerInfo.Title = "MyGram API"
	docs.SwaggerInfo.Description = "My Gram API is an API that allows users to create, read, update, and delete photos."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "mygram-api-reza-production.up.railway.app"
	docs.SwaggerInfo.Schemes = []string{"https"}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	fmt.Println("Server running on PORT =>", PORT)
	route.Run(":" + PORT)
}
