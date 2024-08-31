package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/google-run-code/Delivery/Controllers"
	repository "github.com/google-run-code/Repository"
	usecases "github.com/google-run-code/Usecases"
	"github.com/google-run-code/config"
	"gorm.io/gorm"
)

func NewRoleRouter(db *gorm.DB, env config.Env, router *gin.Engine) {

	roleRepo := repository.NewRoleRepository(db)
	userRepo := repository.NewUserRepository(db)
	roleUseCase := usecases.NewRoleUseCase(roleRepo, userRepo)
	roleHandler := controllers.NewRoleController(roleUseCase)

	router.GET("/roles", roleHandler.GetAllRoles)
	router.GET("/roles/:id", roleHandler.GetRoleById)
	router.GET("/roles/:id/users", roleHandler.GetRoleUsers)
	router.POST("/roles", roleHandler.CreateRole)
	router.PUT("/roles/:id", roleHandler.UpdateRole)
	router.DELETE("/roles/:id", roleHandler.DeleteRole)
}
