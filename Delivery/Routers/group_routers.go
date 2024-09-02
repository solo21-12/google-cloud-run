package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/google-run-code/Delivery/Controllers"
	repository "github.com/google-run-code/Repository"
	usecases "github.com/google-run-code/Usecases"
	"github.com/google-run-code/config"
)

func NewGroupRouter(env config.Env, router *gin.RouterGroup) {

	groupRepo := repository.NewGroupRepository()
	groupUseCase := usecases.NewGroupUseCase(groupRepo)
	groupHandler := controllers.NewGroupController(groupUseCase)

	router.GET("/groups", groupHandler.GetAllGroups)
	router.GET("/groups/:id", groupHandler.GetGroupById)
	router.GET("/groups/:id/users", groupHandler.GetGroupUsers)
	router.POST("/groups", groupHandler.CreateGroup)
	router.PATCH("/groups/:id", groupHandler.UpdateGroup)
	router.DELETE("/groups/:id", groupHandler.DeleteGroup)

}
