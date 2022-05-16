package main

import (
	"embed"

	"github.com/tk42/umlquiz/backend/adapter"
	"github.com/tk42/umlquiz/backend/application"
	"github.com/tk42/umlquiz/backend/infra"
)

//go:embed api/*_*.sql
var embedMigrations embed.FS

func InjectDB() infra.SqlHandler {
	sqlHandler := infra.NewSqlHandler()
	if err := sqlHandler.Migrate(embedMigrations); err != nil {
		panic(err)
	}
	return sqlHandler
}

func InjectUserRepository() infra.UserRepository {
	sqlHandler := InjectDB()
	return infra.NewUserRepository(sqlHandler)
}

func InjectQuizRepository() infra.QuizRepository {
	sqlHandler := InjectDB()
	return infra.NewQuizRepository(sqlHandler)
}

func InjectReportRepository() infra.ReportRepository {
	sqlHandler := InjectDB()
	return infra.NewReportRepository(sqlHandler)
}

func InjectUserUsecase() application.UserUsecase {
	UserRepo := InjectUserRepository()
	return application.NewUserUsecase(UserRepo)
}

func InjectQuizUsecase() application.QuizUsecase {
	UserRepo := InjectUserRepository()
	QuizRepo := InjectQuizRepository()
	ReportRepo := InjectReportRepository()

	return application.NewQuizUsecase(UserRepo, QuizRepo, ReportRepo)
}

func InjectPresentation() adapter.Presentation {
	return adapter.NewPresentation(InjectUserUsecase(), InjectQuizUsecase())
}

// func InjectServerUnauthenticated() adapter.ServerUnauthenticated {
// 	return adapter.NewServerUnauthenticated()
// }
