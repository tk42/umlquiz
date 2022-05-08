package infra

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/tk42/umlquiz/backend/domain"
	autogen "github.com/tk42/umlquiz/backend/gen/proto/golang/github.com/tk42/umlquiz"
	"github.com/tk42/umlquiz/backend/gen/sqlc"
	datetime "google.golang.org/genproto/googleapis/type/datetime"
)

type UserRepository struct {
	domain.IUserRepository
	*SqlHandler
}

func (r *UserRepository) Create(ctx context.Context, req *autogen.AddUserRequest) (*autogen.User, error) {
	user, err := r.SqlHandler.Queries.AddUser(ctx, sqlc.AddUserParams{uuid.New().String(), req.GetUsername(), req.GetPassword(), req.GetEmail(), time.Now(), time.Now()})
	var m autogen.MemgerShip
	user.Membership.Scan(&m)
	return &autogen.User{
		UserId:       user.UserID,
		Username:     user.Username,
		Password:     user.Password,
		Email:        user.Email,
		Profile:      user.Profile.String,
		CreatedAt:    user.CreatedAt.(*datetime.DateTime),
		UpdatedAt:    user.UpdatedAt.(*datetime.DateTime),
		Membership:   m,
		LikedQuizIds: strings.Split(user.LikedQuizIds.String, ","),
		QuizHistory:  strings.Split(user.QuizHistory.String, ","),
	}, err
}
func (r *UserRepository) Update(ctx context.Context, req *autogen.UpdateUserRequest) (*autogen.User, error) {
	user, err := r.SqlHandler.Queries.UpdateUser(ctx, sqlc.UpdateUserParams{req.GetUserId(), req.GetUsername(), req.GetPassword(), req.GetEmail(), sql.NullString{String: req.GetProfile()}, sql.NullInt32{Int32: int32(req.GetMembership())}, time.Now()})
	var m autogen.MemgerShip
	user.Membership.Scan(&m)
	return &autogen.User{
		UserId:       user.UserID,
		Username:     user.Username,
		Password:     user.Password,
		Email:        user.Email,
		Profile:      user.Profile.String,
		CreatedAt:    user.CreatedAt.(*datetime.DateTime),
		UpdatedAt:    user.UpdatedAt.(*datetime.DateTime),
		Membership:   m,
		LikedQuizIds: strings.Split(user.LikedQuizIds.String, ","),
		QuizHistory:  strings.Split(user.QuizHistory.String, ","),
	}, err
}

// TODO: append quiz_id to User.QuizHistory
func (r *UserRepository) AppendQuizHistory() {

}

func (r *UserRepository) Find(ctx context.Context, req *autogen.UserRequest) (*autogen.User, error) {
	user, err := r.SqlHandler.Queries.FindUser(ctx, req.GetUserId())
	var m autogen.MemgerShip
	user.Membership.Scan(&m)
	return &autogen.User{
		UserId:       user.UserID,
		Username:     user.Username,
		Password:     user.Password,
		Email:        user.Email,
		Profile:      user.Profile.String,
		CreatedAt:    user.CreatedAt.(*datetime.DateTime),
		UpdatedAt:    user.UpdatedAt.(*datetime.DateTime),
		Membership:   m,
		LikedQuizIds: strings.Split(user.LikedQuizIds.String, ","),
		QuizHistory:  strings.Split(user.QuizHistory.String, ","),
	}, err
}

func (r *UserRepository) Delete(ctx context.Context, req *autogen.UserRequest) error {
	return r.SqlHandler.Queries.DeleteUser(ctx, req.GetUserId())
}

func NewUserRepository(sqlhandler SqlHandler) UserRepository {
	return UserRepository{
		SqlHandler: &sqlhandler,
	}
}

type QuizRepository struct {
	domain.IQuizRepository
	*SqlHandler
}

func (r *QuizRepository) Create(ctx context.Context, req *autogen.AddQuizRequest) (*autogen.Quiz, error) {
	quiz, err := r.SqlHandler.Queries.AddQuiz(ctx,
		sqlc.AddQuizParams{
			req.GetLanguage(),
			int32(autogen.QuizStatus_DRAFT),
			int32(req.GetDiagramType()),
			sql.NullString{String: req.GetLevel()},
			sql.NullString{String: req.GetTitle()},
			sql.NullString{String: req.GetText()},
			sql.NullString{String: req.GetDiagram()},
			sql.NullString{String: req.GetAuthorId()},
			time.Now(),
			time.Now(),
		})
	if err != nil {
		return &autogen.Quiz{}, err
	}
	quiz, err = r.SqlHandler.Queries.UpdateQuizID(ctx,
		sqlc.UpdateQuizIDParams{
			uuid.New().String(),
			req.GetLanguage(),
			sql.NullString{String: req.GetTitle()},
			quiz.CreatedAt,
		},
	)
	if err != nil {
		return &autogen.Quiz{}, err
	}
	return &autogen.Quiz{
		QuizId:      quiz.QuizID,
		Language:    quiz.Language,
		Status:      autogen.QuizStatus(quiz.Status),
		DiagramType: autogen.DiagramType(quiz.DiagramType),
		Level:       quiz.Level.String,
		Title:       quiz.Title.String,
		Text:        quiz.Text.String,
		Diagram:     quiz.Diagram.String,
		Likes:       quiz.Likes.Int32,
		AuthorId:    quiz.AuthorID.String,
		CreatedAt:   quiz.CreatedAt.(*datetime.DateTime),
		UpdatedAt:   quiz.UpdatedAt.(*datetime.DateTime),
	}, err
}

func (r *QuizRepository) UpdateUpdate(ctx context.Context, req *autogen.UpdateQuizRequest) (*autogen.Quiz, error) {
	quiz, err := r.SqlHandler.Queries.UpdateQuiz(ctx,
		sqlc.UpdateQuizParams{
			req.GetQuizId(),
			req.GetLanguage(),
			int32(req.GetStatus()),
			int32(req.GetDiagramType()),
			sql.NullString{String: req.GetLevel()},
			sql.NullString{String: req.GetTitle()},
			sql.NullString{String: req.GetText()},
			sql.NullString{String: req.GetDiagram()},
			sql.NullInt32{Int32: req.GetLikes()},
			time.Now(),
		},
	)
	return &autogen.Quiz{
		QuizId:      quiz.QuizID,
		Language:    quiz.Language,
		Status:      autogen.QuizStatus(quiz.Status),
		DiagramType: autogen.DiagramType(quiz.DiagramType),
		Level:       quiz.Level.String,
		Title:       quiz.Title.String,
		Text:        quiz.Text.String,
		Diagram:     quiz.Diagram.String,
		Likes:       quiz.Likes.Int32,
		AuthorId:    quiz.AuthorID.String,
		CreatedAt:   quiz.CreatedAt.(*datetime.DateTime),
		UpdatedAt:   quiz.UpdatedAt.(*datetime.DateTime),
	}, err
}

func (r *QuizRepository) UpdateLike(ctx context.Context, req *autogen.LikeQuizRequest) error {
	return r.SqlHandler.Queries.UpdateQuizLike(ctx,
		sqlc.UpdateQuizLikeParams{
			req.GetQuizId(),
			req.GetLanguage(),
			sql.NullInt32{Int32: req.GetDiffLike()},
		},
	)
}

func (r *QuizRepository) Find(ctx context.Context, req *autogen.FindQuizRequest) (*autogen.Quiz, error) {
	quiz, err := r.SqlHandler.Queries.FindQuiz(ctx,
		sqlc.FindQuizParams{
			req.GetQuizId(),
			req.GetLanguage(),
		},
	)
	if err != nil {
		return &autogen.Quiz{}, err
	}
	return &autogen.Quiz{
		QuizId:      quiz.QuizID,
		Language:    quiz.Language,
		Status:      autogen.QuizStatus(quiz.Status),
		DiagramType: autogen.DiagramType(quiz.DiagramType),
		Level:       quiz.Level.String,
		Title:       quiz.Title.String,
		Text:        quiz.Text.String,
		Diagram:     quiz.Diagram.String,
		Likes:       quiz.Likes.Int32,
		AuthorId:    quiz.AuthorID.String,
		CreatedAt:   quiz.CreatedAt.(*datetime.DateTime),
		UpdatedAt:   quiz.UpdatedAt.(*datetime.DateTime),
	}, err
}

func (r *QuizRepository) Delete(ctx context.Context, req *autogen.DeleteQuizRequest) error {
	return r.SqlHandler.Queries.DeleteQuiz(ctx,
		sqlc.DeleteQuizParams{
			req.GetQuizId(),
			req.GetLanguage(),
		},
	)
}

func NewQuizRepository(sqlhandler SqlHandler) QuizRepository {
	return QuizRepository{
		SqlHandler: &sqlhandler,
	}
}

type ReportRepository struct {
	domain.IReportRepository
	*SqlHandler
}

func (r *ReportRepository) Create(ctx context.Context, req *autogen.AddReportRequest) (*autogen.Report, error) {
	report, err := r.SqlHandler.Queries.AddReport(ctx,
		sqlc.AddReportParams{
			uuid.New().String(),
			sql.NullString{String: req.GetQuizId()},
			sql.NullString{String: req.GetLanguage()},
			sql.NullString{String: req.GetUserId()},
			sql.NullString{String: req.GetTitle()},
			sql.NullString{String: req.GetText()},
			sql.NullString{String: req.GetDiagram()},
			sql.NullString{String: req.GetComment()},
			time.Now(),
		},
	)
	if err != nil {
		return &autogen.Report{}, err
	}
	return &autogen.Report{
		ReportId: report.ReportID,
		QuizId:   report.QuizID.String,
		Language: report.Language.String,
		AuthorId: report.AuthorID.String,
		Title:    report.Title.String,
		Text:     report.Text.String,
		Diagram:  report.Diagram.String,
		Comment:  report.Comment.String,
	}, err
}

func (r *ReportRepository) Update(ctx context.Context, req *autogen.UpdateReportRequest) (*autogen.Report, error) {
	report, err := r.SqlHandler.Queries.UpdateReport(ctx,
		sqlc.UpdateReportParams{
			req.GetReportId(),
			sql.NullString{String: req.GetTitle()},
			sql.NullString{String: req.GetText()},
			sql.NullString{String: req.GetDiagram()},
			sql.NullString{String: req.GetComment()},
		},
	)
	if err != nil {
		return &autogen.Report{}, err
	}
	return &autogen.Report{
		ReportId: report.ReportID,
		QuizId:   report.QuizID.String,
		Language: report.Language.String,
		AuthorId: report.AuthorID.String,
		Title:    report.Title.String,
		Text:     report.Text.String,
		Diagram:  report.Diagram.String,
		Comment:  report.Comment.String,
	}, err
}

func (r *ReportRepository) Find(ctx context.Context, req *autogen.FindReportsRequest) ([]*autogen.Report, error) {
	reports, err := r.SqlHandler.Queries.FindReport(ctx,
		sqlc.FindReportParams{
			sql.NullString{String: req.GetQuizId()},
			sql.NullString{String: req.GetLanguage()},
		},
	)
	if err != nil {
		return []*autogen.Report{}, err
	}
	var result []*autogen.Report
	for _, report := range reports {
		result = append(result,
			&autogen.Report{
				ReportId: report.ReportID,
				QuizId:   report.QuizID.String,
				Language: report.Language.String,
				AuthorId: report.AuthorID.String,
				Title:    report.Title.String,
				Text:     report.Text.String,
				Diagram:  report.Diagram.String,
				Comment:  report.Comment.String,
			},
		)
	}
	return result, err
}

func (r *ReportRepository) Delete(ctx context.Context, req *autogen.DeleteReportRequest) error {
	return r.SqlHandler.Queries.DeleteReport(ctx, req.GetReportId())
}

func NewReportRepository(sqlhandler SqlHandler) ReportRepository {
	return ReportRepository{
		SqlHandler: &sqlhandler,
	}
}
