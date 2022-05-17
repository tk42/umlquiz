package adapter

import (
	"context"
	"fmt"

	auth "github.com/tk42/jwt-go-auth"
	"github.com/tk42/umlquiz/backend/application"
	autogen "github.com/tk42/umlquiz/backend/gen/proto/golang/github.com/tk42/umlquiz"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Presentation struct {
	autogen.UnimplementedUMLQuizHelloServiceServer
	autogen.UnimplementedUMLQuizUserServiceServer
	autogen.UnimplementedUMLQuizQuizServiceServer
	autogen.UnimplementedUMLQuizReportServiceServer

	application.UserUsecase
	application.QuizUsecase
}

func (p *Presentation) authenticate(ctx context.Context) error {
	// this check can be seen as optional but its recommended to double check that the context was authenticated
	authCrossCheck, ok := ctx.Value(auth.Authenticated).(bool)
	if !ok && !authCrossCheck {
		return status.Errorf(codes.Unauthenticated, "context not authenticated")
	}
	return nil
}

func (p *Presentation) Hello(ctx context.Context, req *autogen.HelloRequest) (*autogen.HelloResponse, error) {
	return &autogen.HelloResponse{
		Response: fmt.Sprintf("requested: %s", req.GetRequest()),
	}, nil
}

func (p *Presentation) AddUser(ctx context.Context, req *autogen.AddUserRequest) (*autogen.UserResponse, error) {
	// if err := p.authenticate(ctx); err != nil {
	// 	return nil, err
	// }
	user, err := p.UserUsecase.AddUser(ctx, req)
	return &autogen.UserResponse{
		User:  user,
		Error: fmt.Sprint(err),
	}, err
}
func (p *Presentation) UpdateUser(ctx context.Context, req *autogen.UpdateUserRequest) (*autogen.UserResponse, error) {
	// if err := p.authenticate(ctx); err != nil {
	// 	return nil, err
	// }
	user, err := p.UserUsecase.UpdateUser(ctx, req)
	return &autogen.UserResponse{
		User:  user,
		Error: fmt.Sprint(err),
	}, err
}
func (p *Presentation) FindUser(ctx context.Context, req *autogen.UserRequest) (*autogen.UserResponse, error) {
	// if err := p.authenticate(ctx); err != nil {
	// 	return nil, err
	// }
	user, err := p.UserUsecase.FindUser(ctx, req)
	return &autogen.UserResponse{
		User:  user,
		Error: fmt.Sprint(err),
	}, err
}
func (p *Presentation) DeleteUser(ctx context.Context, req *autogen.UserRequest) (*autogen.ErrorResponse, error) {
	// if err := p.authenticate(ctx); err != nil {
	// 	return nil, err
	// }
	err := p.UserUsecase.DeleteUser(ctx, req)
	return &autogen.ErrorResponse{
		Error: fmt.Sprint(err),
	}, err
}
func (p *Presentation) AddQuiz(ctx context.Context, req *autogen.AddQuizRequest) (*autogen.QuizResponse, error) {
	// if err := p.authenticate(ctx); err != nil {
	// 	return nil, err
	// }
	quiz, err := p.QuizUsecase.AddQuiz(ctx, req)
	return &autogen.QuizResponse{
		Quiz:  quiz,
		Error: fmt.Sprint(err),
	}, err
}
func (p *Presentation) FindQuiz(ctx context.Context, req *autogen.FindQuizRequest) (*autogen.QuizResponse, error) {
	// if err := p.authenticate(ctx); err != nil {
	// 	return nil, err
	// }
	// TODO: check the status or should be passed from args?
	quiz, err := p.QuizUsecase.FindQuiz(ctx, req)
	return &autogen.QuizResponse{
		Quiz:  quiz,
		Error: fmt.Sprint(err),
	}, err
}
func (p *Presentation) UpdateQuiz(ctx context.Context, req *autogen.UpdateQuizRequest) (*autogen.QuizResponse, error) {
	// if err := p.authenticate(ctx); err != nil {
	// 	return nil, err
	// }
	quiz, err := p.QuizUsecase.UpdateQuiz(ctx, req)
	return &autogen.QuizResponse{
		Quiz:  quiz,
		Error: fmt.Sprint(err),
	}, err
}
func (p *Presentation) DeleteQuiz(ctx context.Context, req *autogen.DeleteQuizRequest) (*autogen.ErrorResponse, error) {
	// if err := p.authenticate(ctx); err != nil {
	// 	return nil, err
	// }
	err := p.QuizUsecase.DeleteQuiz(ctx, req)
	return &autogen.ErrorResponse{
		Error: fmt.Sprint(err),
	}, err
}
func (p *Presentation) ListQuizzesAll(ctx context.Context, req *autogen.ListQuizzesAllRequest) (*autogen.QuizzesResponse, error) {
	// if err := p.authenticate(ctx); err != nil {
	// 	return nil, err
	// }
	quizzes, err := p.QuizUsecase.ListQuizzesAll(ctx, req)
	var array []*autogen.Quiz
	for _, quiz := range quizzes {
		array = append(array, quiz)
	}
	return &autogen.QuizzesResponse{
		Quiz:  array,
		Error: fmt.Sprint(err),
	}, err
}
func (p *Presentation) ListQuizzesByUser(ctx context.Context, req *autogen.ListQuizzesByUserRequest) (*autogen.QuizzesResponse, error) {
	// if err := p.authenticate(ctx); err != nil {
	// 	return nil, err
	// }
	quizzes, err := p.QuizUsecase.ListQuizzesByUser(ctx, req)
	var array []*autogen.Quiz
	for _, quiz := range quizzes {
		array = append(array, quiz)
	}
	return &autogen.QuizzesResponse{
		Quiz:  array,
		Error: fmt.Sprint(err),
	}, err
}
func (p *Presentation) SolveQuiz(ctx context.Context, req *autogen.SolveQuizRequest) (*autogen.SolveResponse, error) {
	// if err := p.authenticate(ctx); err != nil {
	// 	return nil, err
	// }
	diff, err := p.QuizUsecase.SolveQuiz(ctx, req)
	return &autogen.SolveResponse{
		Diff:  diff,
		Error: fmt.Sprint(err),
	}, err
}
func (p *Presentation) LikeQuiz(ctx context.Context, req *autogen.LikeQuizRequest) (*autogen.ErrorResponse, error) {
	// if err := p.authenticate(ctx); err != nil {
	// 	return nil, err
	// }
	err := p.QuizUsecase.LikeQuiz(ctx, req)
	return &autogen.ErrorResponse{
		Error: fmt.Sprint(err),
	}, err
}

func (p *Presentation) AddReport(ctx context.Context, req *autogen.AddReportRequest) (*autogen.ReportResponse, error) {
	// if err := p.authenticate(ctx); err != nil {
	// 	return nil, err
	// }
	report, err := p.QuizUsecase.AddReport(ctx, req)
	return &autogen.ReportResponse{
		Report: report,
		Error:  fmt.Sprint(err),
	}, err
}
func (p *Presentation) FindReports(ctx context.Context, req *autogen.FindReportsRequest) (*autogen.ReportsResponse, error) {
	// if err := p.authenticate(ctx); err != nil {
	// 	return nil, err
	// }
	reports, err := p.QuizUsecase.FindReports(ctx, req)
	var array []*autogen.Report
	for _, report := range reports {
		array = append(array, report)
	}
	return &autogen.ReportsResponse{
		Report: array,
		Error:  fmt.Sprint(err),
	}, err
}
func (p *Presentation) UpdateReport(ctx context.Context, req *autogen.UpdateReportRequest) (*autogen.ReportResponse, error) {
	// if err := p.authenticate(ctx); err != nil {
	// 	return nil, err
	// }
	report, err := p.QuizUsecase.UpdateReport(ctx, req)
	return &autogen.ReportResponse{
		Report: report,
		Error:  fmt.Sprint(err),
	}, err
}
func (p *Presentation) DeleteReport(ctx context.Context, req *autogen.DeleteReportRequest) (*autogen.ErrorResponse, error) {
	// if err := p.authenticate(ctx); err != nil {
	// 	return nil, err
	// }
	err := p.QuizUsecase.DeleteReport(ctx, req)
	return &autogen.ErrorResponse{
		Error: fmt.Sprint(err),
	}, err
}

func NewPresentation(userUsecase application.UserUsecase, quizUsecase application.QuizUsecase) Presentation {
	return Presentation{
		UserUsecase: userUsecase,
		QuizUsecase: quizUsecase,
	}
}
