package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/google-run-code/Delivery/Controllers"
	infrastructure "github.com/google-run-code/Infrastructure"
	repository "github.com/google-run-code/Repository"
	usecases "github.com/google-run-code/Usecases"
	"github.com/google-run-code/config"
)

func NewUserRouter(env config.Env, router *gin.RouterGroup , dbConfig *config.PostgresConfig) {

	userRepo := repository.NewUserRepository(dbConfig)
	roleRepo := repository.NewRoleRepository(dbConfig)
	groupRepo := repository.NewGroupRepository(dbConfig)
	emailService := infrastructure.NewEmailService(env)

	userUseCase := usecases.NewUserUseCase(userRepo, emailService, roleRepo, groupRepo)
	userHandler := controllers.NewUserController(userUseCase)

	router.GET("/users", userHandler.GetUsers)
	router.GET("/users/:id", userHandler.GetUserById)
	router.GET("/users/:id/groups", userHandler.GetUsersGroup)

	router.POST("/users", userHandler.CreateUser)
	router.PATCH("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	router.POST("/users/:id/groups", userHandler.AddUserToGroup)
	router.DELETE("/users/:id/groups", userHandler.DeletetUserFromGroup)

}
