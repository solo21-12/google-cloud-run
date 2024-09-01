package routers

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/google-run-code/Delivery/Middlewares"
	infrastructure "github.com/google-run-code/Infrastructure"
	"github.com/google-run-code/config"
)

func SetUp() {
	env := config.NewEnv()

	jwtService := infrastructure.NewJwtService(env)
	middleware := middleware.DatabaseMiddleware(env, jwtService)

	router := gin.Default()

	public := router.Group("")
	protected := public.Group("")
	protected.Use(middleware)
	NewUserRouter(*env, protected)
	NewGroupRouter(*env, protected)
	NewRoleRouter(*env, protected)
	NewGenerateTokenRouter(public)

	router.Run(":8081")
}
