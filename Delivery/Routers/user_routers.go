package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/google-run-code/Delivery/Controllers"
	infrastructure "github.com/google-run-code/Infrastructure"
	repository "github.com/google-run-code/Repository"
	usecases "github.com/google-run-code/Usecases"
	"github.com/google-run-code/config"
	"gorm.io/gorm"
)

func NewUserRouter(db *gorm.DB, env config.Env, router *gin.Engine) {
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	groupRepo := repository.NewGroupRepository(db)
	passwordService := infrastructure.NewPasswordService()
	emailService := infrastructure.NewEmailService(env)

	userUseCase := usecases.NewUserUseCase(userRepo, passwordService, emailService, roleRepo, groupRepo)
	userHandler := controllers.NewUserController(userUseCase)

	router.GET("/users", userHandler.GetUsers)
	router.GET("/users/:id", userHandler.GetUserById)
	router.GET("/users/:id/groups", userHandler.GetUsersGroup)

	router.POST("/users", userHandler.CreateUser)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	router.POST("/users/:id/groups", userHandler.AddUserToGroup)
	router.POST("/users/:id/roles", userHandler.AddUserToRole)

}
