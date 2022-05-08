package adapter

import (
	"context"

	"github.com/tk42/umlquiz/backend/application"
	autogen "github.com/tk42/umlquiz/backend/gen/proto/golang/github.com/tk42/umlquiz"
)

type Presentation struct {
	autogen.UMLQuizServiceServer

	application.UserUsecase
	application.QuizUsecase
}

func (p *Presentation) AddUser(ctx context.Context, req *autogen.AddUserRequest) (*autogen.UserResponse, error) {
	user, err := p.UserUsecase.AddUser(ctx, req)
	return &autogen.UserResponse{
		User:  user,
		Error: err.Error(),
	}, err
}
func (p *Presentation) UpdateUser(ctx context.Context, req *autogen.UpdateUserRequest) (*autogen.UserResponse, error) {
	user, err := p.UserUsecase.UpdateUser(ctx, req)
	return &autogen.UserResponse{
		User:  user,
		Error: err.Error(),
	}, err
}
func (p *Presentation) FindUser(ctx context.Context, req *autogen.UserRequest) (*autogen.UserResponse, error) {
	user, err := p.UserUsecase.FindUser(ctx, req)
	return &autogen.UserResponse{
		User:  user,
		Error: err.Error(),
	}, err
}
func (p *Presentation) DeleteUser(ctx context.Context, req *autogen.UserRequest) (*autogen.ErrorResponse, error) {
	err := p.UserUsecase.DeleteUser(ctx, req)
	return &autogen.ErrorResponse{
		Error: err.Error(),
	}, err
}
func (p *Presentation) AddQuiz(ctx context.Context, req *autogen.AddQuizRequest) (*autogen.QuizResponse, error) {
	quiz, err := p.QuizUsecase.AddQuiz(ctx, req)
	return &autogen.QuizResponse{
		Quiz:  quiz,
		Error: err.Error(),
	}, err
}
func (p *Presentation) FindQuiz(ctx context.Context, req *autogen.FindQuizRequest) (*autogen.QuizResponse, error) {
	// TODO: check the status or should be passed from args?
	quiz, err := p.QuizUsecase.FindQuiz(ctx, req)
	return &autogen.QuizResponse{
		Quiz:  quiz,
		Error: err.Error(),
	}, err
}
func (p *Presentation) UpdateQuiz(ctx context.Context, req *autogen.UpdateQuizRequest) (*autogen.QuizResponse, error) {
	quiz, err := p.QuizUsecase.UpdateQuiz(ctx, req)
	return &autogen.QuizResponse{
		Quiz:  quiz,
		Error: err.Error(),
	}, err
}
func (p *Presentation) DeleteQuiz(ctx context.Context, req *autogen.DeleteQuizRequest) (*autogen.ErrorResponse, error) {
	err := p.QuizUsecase.DeleteQuiz(ctx, req)
	return &autogen.ErrorResponse{
		Error: err.Error(),
	}, err
}
func (p *Presentation) ListQuizzesAll(ctx context.Context, req *autogen.ListQuizzesAllRequest) (*autogen.QuizzesResponse, error) {
	quizzes, err := p.QuizUsecase.ListQuizzesAll(ctx, req)
	var array []*autogen.Quiz
	for _, quiz := range quizzes {
		array = append(array, quiz)
	}
	return &autogen.QuizzesResponse{
		Quiz:  array,
		Error: err.Error(),
	}, err
}
func (p *Presentation) ListQuizzesByUser(ctx context.Context, req *autogen.ListQuizzesByUserRequest) (*autogen.QuizzesResponse, error) {
	quizzes, err := p.QuizUsecase.ListQuizzesByUser(ctx, req)
	var array []*autogen.Quiz
	for _, quiz := range quizzes {
		array = append(array, quiz)
	}
	return &autogen.QuizzesResponse{
		Quiz:  array,
		Error: err.Error(),
	}, err
}
func (p *Presentation) SolveQuiz(ctx context.Context, req *autogen.SolveQuizRequest) (*autogen.SolveResponse, error) {
	diff, err := p.QuizUsecase.SolveQuiz(ctx, req)
	return &autogen.SolveResponse{
		Diff:  diff,
		Error: err.Error(),
	}, err
}
func (p *Presentation) LikeQuiz(ctx context.Context, req *autogen.LikeQuizRequest) (*autogen.ErrorResponse, error) {
	err := p.QuizUsecase.LikeQuiz(ctx, req)
	return &autogen.ErrorResponse{
		Error: err.Error(),
	}, err
}

func (p *Presentation) AddReport(ctx context.Context, req *autogen.AddReportRequest) (*autogen.ReportResponse, error) {
	report, err := p.QuizUsecase.AddReport(ctx, req)
	return &autogen.ReportResponse{
		Report: report,
		Error:  err.Error(),
	}, err
}
func (p *Presentation) FindReports(ctx context.Context, req *autogen.FindReportsRequest) (*autogen.ReportsResponse, error) {
	reports, err := p.QuizUsecase.FindReports(ctx, req)
	var array []*autogen.Report
	for _, report := range reports {
		array = append(array, report)
	}
	return &autogen.ReportsResponse{
		Report: array,
		Error:  err.Error(),
	}, err
}
func (p *Presentation) UpdateReport(ctx context.Context, req *autogen.UpdateReportRequest) (*autogen.ReportResponse, error) {
	report, err := p.QuizUsecase.UpdateReport(ctx, req)
	return &autogen.ReportResponse{
		Report: report,
		Error:  err.Error(),
	}, err
}
func (p *Presentation) DeleteReport(ctx context.Context, req *autogen.DeleteReportRequest) (*autogen.ErrorResponse, error) {
	err := p.QuizUsecase.DeleteReport(ctx, req)
	return &autogen.ErrorResponse{
		Error: err.Error(),
	}, err
}

func NewPresentation(userUsecase application.UserUsecase, quizUsecase application.QuizUsecase) Presentation {
	return Presentation{
		UserUsecase: userUsecase,
		QuizUsecase: quizUsecase,
	}
}
