package application

import (
	"context"

	autogen "github.com/tk42/umlquiz/backend/gen/proto/golang/github.com/tk42/umlquiz"
)

type IUserUsecase interface {
	AddUser(context.Context, *autogen.AddUserRequest) (*autogen.User, error)
	UpdateUser(context.Context, *autogen.UpdateUserRequest) (*autogen.User, error)
	FindUser(context.Context, *autogen.UserRequest) (*autogen.User, error)
	DeleteUser(context.Context, *autogen.UserRequest) error
}

type IQuizUsecase interface {
	AddQuiz(context.Context, *autogen.AddQuizRequest) (*autogen.Quiz, error)
	FindQuiz(context.Context, *autogen.FindQuizRequest) (*autogen.Quiz, error)
	UpdateQuiz(context.Context, *autogen.UpdateQuizRequest) (*autogen.Quiz, error)
	DeleteQuiz(context.Context, *autogen.DeleteQuizRequest) error
	ListQuizzesAll(context.Context, *autogen.ListQuizzesAllRequest) ([]*autogen.Quiz, error)
	ListQuizzesByUser(context.Context, *autogen.ListQuizzesByUserRequest) ([]*autogen.Quiz, error)
	SolveQuiz(context.Context, *autogen.SolveQuizRequest) (string, error)
	LikeQuiz(context.Context, *autogen.LikeQuizRequest) error
	AddReport(context.Context, *autogen.AddReportRequest) (*autogen.Report, error)
	FindReports(context.Context, *autogen.FindReportsRequest) ([]*autogen.Report, error)
	UpdateReport(context.Context, *autogen.UpdateReportRequest) (*autogen.Report, error)
	DeleteReport(context.Context, *autogen.DeleteReportRequest) error
}
