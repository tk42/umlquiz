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
	"github.com/tk42/umlquiz/backend/utility"
)

type UserRepository struct {
	domain.IUserRepository
	*SqlHandler
}

func (r *UserRepository) Create(ctx context.Context, req *autogen.AddUserRequest) (*autogen.User, error) {
	user, err := r.SqlHandler.Queries.AddUser(ctx,
		sqlc.AddUserParams{
			UserID:    uuid.New().String(),
			Username:  req.GetUsername(),
			Password:  req.GetPassword(),
			Email:     req.GetEmail(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	)
	if err != nil {
		return nil, err
	}
	var m autogen.MemgerShip
	user.Membership.Scan(&m)
	return &autogen.User{
		UserId:       user.UserID,
		Username:     user.Username,
		Password:     user.Password,
		Email:        user.Email,
		Profile:      user.Profile.String,
		CreatedAt:    utility.ToDatetime(user.CreatedAt),
		UpdatedAt:    utility.ToDatetime(user.UpdatedAt),
		Membership:   m,
		LikedQuizIds: strings.Split(user.LikedQuizIds.String, ","),
		QuizHistory:  strings.Split(user.QuizHistory.String, ","),
	}, nil
}
func (r *UserRepository) Update(ctx context.Context, req *autogen.UpdateUserRequest) (*autogen.User, error) {
	user, err := r.SqlHandler.Queries.UpdateUser(ctx, sqlc.UpdateUserParams{UserID: req.GetUserId(), Username: req.GetUsername(), Password: req.GetPassword(), Email: req.GetEmail(), Profile: sql.NullString{String: req.GetProfile()}, Membership: sql.NullInt32{Int32: int32(req.GetMembership())}, UpdatedAt: time.Now()})
	if err != nil {
		return nil, err
	}
	var m autogen.MemgerShip
	user.Membership.Scan(&m)
	return &autogen.User{
		UserId:       user.UserID,
		Username:     user.Username,
		Password:     user.Password,
		Email:        user.Email,
		Profile:      user.Profile.String,
		CreatedAt:    utility.ToDatetime(user.CreatedAt),
		UpdatedAt:    utility.ToDatetime(user.UpdatedAt),
		Membership:   m,
		LikedQuizIds: strings.Split(user.LikedQuizIds.String, ","),
		QuizHistory:  strings.Split(user.QuizHistory.String, ","),
	}, nil
}

// TODO: append quiz_id to User.QuizHistory
func (r *UserRepository) AppendQuizHistory() {

}

func (r *UserRepository) Find(ctx context.Context, req *autogen.UserRequest) (*autogen.User, error) {
	user, err := r.SqlHandler.Queries.FindUser(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	var m autogen.MemgerShip
	user.Membership.Scan(&m)
	return &autogen.User{
		UserId:       user.UserID,
		Username:     user.Username,
		Password:     user.Password,
		Email:        user.Email,
		Profile:      user.Profile.String,
		CreatedAt:    utility.ToDatetime(user.CreatedAt),
		UpdatedAt:    utility.ToDatetime(user.UpdatedAt),
		Membership:   m,
		LikedQuizIds: strings.Split(user.LikedQuizIds.String, ","),
		QuizHistory:  strings.Split(user.QuizHistory.String, ","),
	}, nil
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
			Language:    req.GetLanguage(),
			Status:      int32(autogen.QuizStatus_DRAFT),
			DiagramType: int32(req.GetDiagramType()),
			Level:       sql.NullString{String: req.GetLevel()},
			Title:       sql.NullString{String: req.GetTitle()},
			Text:        sql.NullString{String: req.GetText()},
			Diagram:     sql.NullString{String: req.GetDiagram()},
			AuthorID:    sql.NullString{String: req.GetAuthorId()},
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
	if err != nil {
		return nil, err
	}
	quiz, err = r.SqlHandler.Queries.UpdateQuizID(ctx,
		sqlc.UpdateQuizIDParams{
			QuizID:    uuid.New().String(),
			Language:  req.GetLanguage(),
			Title:     sql.NullString{String: req.GetTitle()},
			CreatedAt: quiz.CreatedAt,
		},
	)
	if err != nil {
		return nil, err
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
		CreatedAt:   utility.ToDatetime(quiz.CreatedAt),
		UpdatedAt:   utility.ToDatetime(quiz.UpdatedAt),
	}, nil
}

func (r *QuizRepository) UpdateUpdate(ctx context.Context, req *autogen.UpdateQuizRequest) (*autogen.Quiz, error) {
	quiz, err := r.SqlHandler.Queries.UpdateQuiz(ctx,
		sqlc.UpdateQuizParams{
			QuizID:      req.GetQuizId(),
			Language:    req.GetLanguage(),
			Status:      int32(req.GetStatus()),
			DiagramType: int32(req.GetDiagramType()),
			Level:       sql.NullString{String: req.GetLevel()},
			Title:       sql.NullString{String: req.GetTitle()},
			Text:        sql.NullString{String: req.GetText()},
			Diagram:     sql.NullString{String: req.GetDiagram()},
			Likes:       sql.NullInt32{Int32: req.GetLikes()},
			UpdatedAt:   time.Now(),
		},
	)
	if err != nil {
		return nil, err
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
		CreatedAt:   utility.ToDatetime(quiz.CreatedAt),
		UpdatedAt:   utility.ToDatetime(quiz.UpdatedAt),
	}, nil
}

func (r *QuizRepository) UpdateLike(ctx context.Context, req *autogen.LikeQuizRequest) error {
	return r.SqlHandler.Queries.UpdateQuizLike(ctx,
		sqlc.UpdateQuizLikeParams{
			QuizID:   req.GetQuizId(),
			Language: req.GetLanguage(),
			Likes:    sql.NullInt32{Int32: req.GetDiffLike()},
		},
	)
}

func (r *QuizRepository) Find(ctx context.Context, req *autogen.FindQuizRequest) (*autogen.Quiz, error) {
	quiz, err := r.SqlHandler.Queries.FindQuiz(ctx,
		sqlc.FindQuizParams{
			QuizID:   req.GetQuizId(),
			Language: req.GetLanguage(),
		},
	)
	if err != nil {
		return nil, err
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
		CreatedAt:   utility.ToDatetime(quiz.CreatedAt),
		UpdatedAt:   utility.ToDatetime(quiz.UpdatedAt),
	}, nil
}

func (r *QuizRepository) Delete(ctx context.Context, req *autogen.DeleteQuizRequest) error {
	return r.SqlHandler.Queries.DeleteQuiz(ctx,
		sqlc.DeleteQuizParams{
			QuizID:   req.GetQuizId(),
			Language: req.GetLanguage(),
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
			ReportID:  uuid.New().String(),
			QuizID:    sql.NullString{String: req.GetQuizId()},
			Language:  sql.NullString{String: req.GetLanguage()},
			AuthorID:  sql.NullString{String: req.GetUserId()},
			Title:     sql.NullString{String: req.GetTitle()},
			Text:      sql.NullString{String: req.GetText()},
			Diagram:   sql.NullString{String: req.GetDiagram()},
			Comment:   sql.NullString{String: req.GetComment()},
			CreatedAt: time.Now(),
		},
	)
	if err != nil {
		return nil, err
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
	}, nil
}

func (r *ReportRepository) Update(ctx context.Context, req *autogen.UpdateReportRequest) (*autogen.Report, error) {
	report, err := r.SqlHandler.Queries.UpdateReport(ctx,
		sqlc.UpdateReportParams{
			ReportID: req.GetReportId(),
			Title:    sql.NullString{String: req.GetTitle()},
			Text:     sql.NullString{String: req.GetText()},
			Diagram:  sql.NullString{String: req.GetDiagram()},
			Comment:  sql.NullString{String: req.GetComment()},
		},
	)
	if err != nil {
		return nil, err
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
	}, nil
}

func (r *ReportRepository) Find(ctx context.Context, req *autogen.FindReportsRequest) ([]*autogen.Report, error) {
	reports, err := r.SqlHandler.Queries.FindReport(ctx,
		sqlc.FindReportParams{
			QuizID:   sql.NullString{String: req.GetQuizId()},
			Language: sql.NullString{String: req.GetLanguage()},
		},
	)
	if err != nil {
		return nil, err
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
	return result, nil
}

func (r *ReportRepository) Delete(ctx context.Context, req *autogen.DeleteReportRequest) error {
	return r.SqlHandler.Queries.DeleteReport(ctx, req.GetReportId())
}

func NewReportRepository(sqlhandler SqlHandler) ReportRepository {
	return ReportRepository{
		SqlHandler: &sqlhandler,
	}
}
