package main

import (
	"main/internal/api"
	"main/internal/conf"
	"main/internal/conf/db/mongo"
	"main/internal/conf/db/mysql"
)

type App struct {
	config *conf.App
	mysql  *mysql.DB
	router *api.Router
}

func New() *App {
	config := conf.New()
	mysqlDB := mysql.New()
	mongoDB := mongo.New()
	router := api.New(config, mysqlDB, mongoDB)

	return &App{
		config: config,
		mysql:  mysqlDB,
		router: router,
	}
}

func (app *App) Run() {
	app.router.Run()
}
