package routers

import (
	"github.com/gin-gonic/gin"
	models "github.com/google-run-code/Domain/Models"
	"github.com/google-run-code/config"
)

func SetUp() {
	env := config.NewEnv()
	db := config.NewPostgresConfig(*env)
	client := db.Client(env.DB_NAME)
	db.Migrate(&models.User{}, &models.Group{}, &models.Role{})

	router := gin.Default()

	NewUserRouter(client, *env, router)
	NewGroupRouter(client, *env, router)
	NewRoleRouter(client, *env, router)

	router.Run(":8081")
}
