package routers

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	middleware "github.com/google-run-code/Delivery/Middlewares"
	models "github.com/google-run-code/Domain/Models"
	infrastructure "github.com/google-run-code/Infrastructure"
	"github.com/google-run-code/config"
)

func getDatabasesFromEnv(env config.Env) []string {
	dbs := env.DB_NAMES
	return strings.Split(dbs, ",")
}

func SetUp() {
	env := config.NewEnv()
	dbConfig := config.NewPostgresConfig(*env)

	dbNames := getDatabasesFromEnv(*env)
	if len(dbNames) == 0 {
		log.Fatalf("No database names provided")
	}

	dbConfig.InitializeConnections(dbNames)

	for _, dbname := range dbNames {
		dbConfig.Migrate(dbname, &models.User{}, &models.Role{}, &models.Group{})
	}

	log.Println(dbNames, "dbname")

	jwtService := infrastructure.NewJwtService(env)
	middleware := middleware.DatabaseMiddleware(env, jwtService)

	router := gin.Default()

	public := router.Group("")
	protected := public.Group("")
	protected.Use(middleware)
	NewUserRouter(*env, protected, dbConfig)
	NewGroupRouter(*env, protected, dbConfig)
	NewRoleRouter(*env, protected, dbConfig)
	NewGenerateTokenRouter(public)

	router.Run(":8081")
}
