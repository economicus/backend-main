package api

import (
	"main/internal/api/handler"
	"main/internal/api/repo"
	"main/internal/api/route"
	"main/internal/api/service"
	"main/internal/conf/aws"
	"main/internal/core/model/quant"
	"main/internal/core/model/table"
)

func (r *Router) migrate() {
	ms := []interface{}{
		&table.User{},
		&table.Profile{},
		&table.Quant{},
		&table.QuantOption{},
		&quant.MainSector{},
		&table.Reply{},
		&table.Comment{},
		&table.Reply{},
	}
	r.mysqlDB.Migrate(ms)
}

func (r *Router) setup() {
	r.migrate()

	a := aws.New()

	authRepo := repo.NewAuthRepo(r.mysqlDB.DB)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)
	route.SetAuth(r.getGroup(), authHandler)

	quantRepo := repo.NewQuantRepo(r.mysqlDB.DB, r.mongoDB.DB)
	quantService := service.NewQuantService(quantRepo)
	quantHandler := handler.NewQuantHandler(quantService)
	route.SetQuant(r.getGroupWithAuth(), quantHandler)

	userRepo := repo.NewUser(r.mysqlDB.DB, a)
	userService := service.NewUserService(userRepo, a)
	userHandler := handler.NewUserHandler(userService, quantService)
	route.SetUser(r.getGroup(), userHandler, authMid)

	commentRepo := repo.NewCommentRepo(r.mysqlDB.DB)
	commentService := service.NewCommentService(commentRepo)
	commentHandler := handler.NewCommentHandler(commentService)
	route.SetComment(r.getGroupWithAuth(), commentHandler)

	replyRepo := repo.NewReplyRepository(r.mysqlDB.DB)
	replyService := service.NewReplyService(replyRepo)
	replyHandler := handler.NewReplyHandler(replyService)
	route.SetReply(r.getGroupWithAuth(), replyHandler)
}
