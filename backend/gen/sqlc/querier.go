// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package sqlc

import (
	"context"
)

type Querier interface {
	AddQuiz(ctx context.Context, arg AddQuizParams) (Quiz, error)
	AddReport(ctx context.Context, arg AddReportParams) (Report, error)
	AddUser(ctx context.Context, arg AddUserParams) (User, error)
	DeleteQuiz(ctx context.Context, arg DeleteQuizParams) error
	DeleteReport(ctx context.Context, reportID string) error
	DeleteUser(ctx context.Context, userID string) error
	FindQuiz(ctx context.Context, arg FindQuizParams) (Quiz, error)
	FindReport(ctx context.Context, arg FindReportParams) ([]Report, error)
	FindUser(ctx context.Context, userID string) (User, error)
	ListQuizzesAll(ctx context.Context, arg ListQuizzesAllParams) ([]Quiz, error)
	ListQuizzesByUser(ctx context.Context, arg ListQuizzesByUserParams) ([]Quiz, error)
	UpdateQuiz(ctx context.Context, arg UpdateQuizParams) (Quiz, error)
	UpdateQuizID(ctx context.Context, arg UpdateQuizIDParams) (Quiz, error)
	UpdateQuizLike(ctx context.Context, arg UpdateQuizLikeParams) error
	UpdateReport(ctx context.Context, arg UpdateReportParams) (Report, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
