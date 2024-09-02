package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/google-run-code/Delivery/Controllers"
	repository "github.com/google-run-code/Repository"
	usecases "github.com/google-run-code/Usecases"
	"github.com/google-run-code/config"
)

func NewRoleRouter(env config.Env, router *gin.RouterGroup) {

	roleRepo := repository.NewRoleRepository()
	userRepo := repository.NewUserRepository()
	roleUseCase := usecases.NewRoleUseCase(roleRepo, userRepo)
	roleHandler := controllers.NewRoleController(roleUseCase)

	router.GET("/roles", roleHandler.GetAllRoles)
	router.GET("/roles/:id", roleHandler.GetRoleById)
	router.GET("/roles/:id/users", roleHandler.GetRoleUsers)
	router.POST("/roles", roleHandler.CreateRole)
	router.PATCH("/roles/:id", roleHandler.UpdateRole)
	router.DELETE("/roles/:id", roleHandler.DeleteRole)
}
