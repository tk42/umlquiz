package application

import (
	"context"
	"database/sql"

	"github.com/kylelemons/godebug/diff"
	autogen "github.com/tk42/umlquiz/backend/gen/proto/golang/github.com/tk42/umlquiz"
	"github.com/tk42/umlquiz/backend/gen/sqlc"
	"github.com/tk42/umlquiz/backend/infra"
	"github.com/tk42/umlquiz/backend/utility"
)

type UserUsecase struct {
	infra.UserRepository
}

func NewUserUsecase(userRepo infra.UserRepository) UserUsecase {
	return UserUsecase{
		UserRepository: userRepo,
	}
}

func (u *UserUsecase) AddUser(ctx context.Context, req *autogen.AddUserRequest) (*autogen.User, error) {
	return u.UserRepository.Create(ctx, req)
}
func (u *UserUsecase) UpdateUser(ctx context.Context, req *autogen.UpdateUserRequest) (*autogen.User, error) {
	return u.UserRepository.Update(ctx, req)
}
func (u *UserUsecase) FindUser(ctx context.Context, req *autogen.UserRequest) (*autogen.User, error) {
	return u.UserRepository.Find(ctx, req)
}
func (u *UserUsecase) DeleteUser(ctx context.Context, req *autogen.UserRequest) error {
	return u.UserRepository.Delete(ctx, req)
}

type QuizUsecase struct {
	infra.UserRepository
	infra.QuizRepository
	infra.ReportRepository
}

func NewQuizUsecase(userRepo infra.UserRepository, quizRepo infra.QuizRepository, repoRepo infra.ReportRepository) QuizUsecase {
	return QuizUsecase{
		UserRepository:   userRepo,
		QuizRepository:   quizRepo,
		ReportRepository: repoRepo,
	}
}

func (u *QuizUsecase) AddQuiz(ctx context.Context, req *autogen.AddQuizRequest) (*autogen.Quiz, error) {
	return u.QuizRepository.Create(ctx, req)
}
func (u *QuizUsecase) FindQuiz(ctx context.Context, req *autogen.FindQuizRequest) (*autogen.Quiz, error) {
	return u.QuizRepository.Find(ctx, req)
}
func (u *QuizUsecase) UpdateQuiz(ctx context.Context, req *autogen.UpdateQuizRequest) (*autogen.Quiz, error) {
	return u.QuizRepository.Update(ctx, req)
}
func (u *QuizUsecase) DeleteQuiz(ctx context.Context, req *autogen.DeleteQuizRequest) error {
	return u.QuizRepository.Delete(ctx, req)
}
func (u *QuizUsecase) ListQuizzesAll(ctx context.Context, req *autogen.ListQuizzesAllRequest) ([]*autogen.Quiz, error) {
	quizzez, err := u.QuizRepository.SqlHandler.Queries.ListQuizzesAll(ctx, sqlc.ListQuizzesAllParams{
		Language: req.GetLanguage(),
		Status:   int32(req.GetStatus()),
	})
	var result []*autogen.Quiz
	for _, quiz := range quizzez {
		result = append(result, &autogen.Quiz{
			QuizId:      quiz.QuizID,
			Language:    quiz.Language,
			Status:      autogen.QuizStatus_DRAFT,
			DiagramType: autogen.DiagramType(quiz.DiagramType),
			Level:       quiz.Level.String,
			Title:       quiz.Title.String,
			Text:        quiz.Text.String,
			Diagram:     quiz.Diagram.String,
			Likes:       quiz.Likes.Int32,
			AuthorId:    quiz.AuthorID.String,
			CreatedAt:   utility.ToDatetime(quiz.CreatedAt),
			UpdatedAt:   utility.ToDatetime(quiz.UpdatedAt),
		})
	}
	return result, err
}
func (u *QuizUsecase) ListQuizzesByUser(ctx context.Context, req *autogen.ListQuizzesByUserRequest) ([]*autogen.Quiz, error) {
	quizzez, err := u.QuizRepository.SqlHandler.Queries.ListQuizzesByUser(ctx,
		sqlc.ListQuizzesByUserParams{
			AuthorID: sql.NullString{String: req.GetUserId()},
			Language: req.GetLanguage(),
			Status:   int32(req.GetStatus()),
		})
	var result []*autogen.Quiz
	for _, quiz := range quizzez {
		result = append(result,
			&autogen.Quiz{
				QuizId:      quiz.QuizID,
				Language:    quiz.Language,
				Status:      autogen.QuizStatus_DRAFT,
				DiagramType: autogen.DiagramType(quiz.DiagramType),
				Level:       quiz.Level.String,
				Title:       quiz.Title.String,
				Text:        quiz.Text.String,
				Diagram:     quiz.Diagram.String,
				Likes:       quiz.Likes.Int32,
				AuthorId:    quiz.AuthorID.String,
				CreatedAt:   utility.ToDatetime(quiz.CreatedAt),
				UpdatedAt:   utility.ToDatetime(quiz.UpdatedAt),
			},
		)
	}
	return result, err
}
func (u *QuizUsecase) SolveQuiz(ctx context.Context, req *autogen.SolveQuizRequest) (string, error) {
	quiz, err := u.FindQuiz(ctx, &autogen.FindQuizRequest{
		QuizId:   req.GetQuizId(),
		Language: req.GetLanguage(),
	})
	if err != nil {
		return "", err
	}
	// TODO: append quiz_id to User.QuizHistory
	// user, err := u.UserRepository.Find(user_id)
	// if err != nil {
	// 	return "", err
	// }
	// _, err = u.UserRepository.Update(user_id, user.Username, user.Password, user.Email, user.Profile, user.Membership)
	return diff.Diff(quiz.Diagram, req.GetDiagram()), err
}

func (u *QuizUsecase) LikeQuiz(ctx context.Context, req *autogen.LikeQuizRequest) error {
	err := u.QuizRepository.UpdateLike(ctx, req)
	if err != nil {
		return err
	}
	// TODO: append quiz_id to User.LikedQuizIds
	// user, err := u.UserRepository.Find(user_id)
	return nil
}

func (u *QuizUsecase) AddReport(ctx context.Context, req *autogen.AddReportRequest) (*autogen.Report, error) {
	return u.ReportRepository.Create(ctx, req)
}

func (u *QuizUsecase) FindReports(ctx context.Context, req *autogen.FindReportsRequest) ([]*autogen.Report, error) {
	return u.ReportRepository.Find(ctx, req)
}

func (u *QuizUsecase) UpdateReport(ctx context.Context, req *autogen.UpdateReportRequest) (*autogen.Report, error) {
	return u.ReportRepository.Update(ctx, req)
}

func (u *QuizUsecase) DeleteReport(ctx context.Context, req *autogen.DeleteReportRequest) error {
	return u.ReportRepository.Delete(ctx, req)
}
