package utility

import (
	"github.com/tk42/umlquiz/backend/adapter"
	"github.com/tk42/umlquiz/backend/application"
	"github.com/tk42/umlquiz/backend/infra"
)

func InjectDB() infra.SqlHandler {
	return infra.NewSqlHandler()
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
