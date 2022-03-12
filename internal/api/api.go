package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"main/internal/api/middleware"
	"main/internal/conf"
	"main/internal/conf/db/mongo"
	"main/internal/conf/db/mysql"
	"time"
)

var authMid *middleware.AuthMiddleware

type Router struct {
	engine  *gin.Engine
	app     *conf.App
	mysqlDB *mysql.DB
	mongoDB *mongo.DB
}

func New(app *conf.App, mysqlDB *mysql.DB, mongoDB *mongo.DB) *Router {
	e := getEngine()
	authMid = middleware.NewAuthMiddleware(mysqlDB)
	r := Router{
		engine:  e,
		app:     app,
		mysqlDB: mysqlDB,
		mongoDB: mongoDB,
	}
	r.setup()
	return &r
}

func (r *Router) Run() {
	if err := r.engine.Run(":" + r.app.InsecurePort); err != nil {
		log.Panicf("error while running app: %v", err)
	}
}

func (r *Router) getGroup() *gin.RouterGroup {
	return r.engine.Group("/v1")
}

func (r *Router) getGroupWithAuth() *gin.RouterGroup {
	return r.engine.Group("/v1", authMid.Authenticate())
}

func getEngine() *gin.Engine {
	e := gin.Default()
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://www.economicus.kr"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	return e
}
