// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: query.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const addQuiz = `-- name: AddQuiz :one
INSERT INTO "quiz" (
  language, status, diagram_type, level, title, text, diagram, author_id, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING quiz_id, language, status, diagram_type, level, title, text, diagram, likes, author_id, created_at, updated_at
`

type AddQuizParams struct {
	Language    string         `json:"language"`
	Status      int32          `json:"status"`
	DiagramType int32          `json:"diagram_type"`
	Level       sql.NullString `json:"level"`
	Title       sql.NullString `json:"title"`
	Text        sql.NullString `json:"text"`
	Diagram     sql.NullString `json:"diagram"`
	AuthorID    sql.NullString `json:"author_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func (q *Queries) AddQuiz(ctx context.Context, arg AddQuizParams) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, addQuiz,
		arg.Language,
		arg.Status,
		arg.DiagramType,
		arg.Level,
		arg.Title,
		arg.Text,
		arg.Diagram,
		arg.AuthorID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Quiz
	err := row.Scan(
		&i.QuizID,
		&i.Language,
		&i.Status,
		&i.DiagramType,
		&i.Level,
		&i.Title,
		&i.Text,
		&i.Diagram,
		&i.Likes,
		&i.AuthorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const addReport = `-- name: AddReport :one
INSERT INTO "report" (
  report_id, quiz_id, language, author_id, title, text, diagram, comment, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING report_id, quiz_id, language, author_id, title, text, diagram, comment, created_at
`

type AddReportParams struct {
	ReportID  string         `json:"report_id"`
	QuizID    sql.NullString `json:"quiz_id"`
	Language  sql.NullString `json:"language"`
	AuthorID  sql.NullString `json:"author_id"`
	Title     sql.NullString `json:"title"`
	Text      sql.NullString `json:"text"`
	Diagram   sql.NullString `json:"diagram"`
	Comment   sql.NullString `json:"comment"`
	CreatedAt time.Time      `json:"created_at"`
}

func (q *Queries) AddReport(ctx context.Context, arg AddReportParams) (Report, error) {
	row := q.db.QueryRowContext(ctx, addReport,
		arg.ReportID,
		arg.QuizID,
		arg.Language,
		arg.AuthorID,
		arg.Title,
		arg.Text,
		arg.Diagram,
		arg.Comment,
		arg.CreatedAt,
	)
	var i Report
	err := row.Scan(
		&i.ReportID,
		&i.QuizID,
		&i.Language,
		&i.AuthorID,
		&i.Title,
		&i.Text,
		&i.Diagram,
		&i.Comment,
		&i.CreatedAt,
	)
	return i, err
}

const addUser = `-- name: AddUser :one
INSERT INTO "user" (
  user_id, username, password, email, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING user_id, username, password, email, profile, created_at, updated_at, membership, liked_quiz_ids, quiz_history
`

type AddUserParams struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) AddUser(ctx context.Context, arg AddUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, addUser,
		arg.UserID,
		arg.Username,
		arg.Password,
		arg.Email,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Profile,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Membership,
		&i.LikedQuizIds,
		&i.QuizHistory,
	)
	return i, err
}

const deleteQuiz = `-- name: DeleteQuiz :exec
DELETE FROM "quiz"
WHERE quiz_id = $1 and language = $2
`

type DeleteQuizParams struct {
	QuizID   string `json:"quiz_id"`
	Language string `json:"language"`
}

func (q *Queries) DeleteQuiz(ctx context.Context, arg DeleteQuizParams) error {
	_, err := q.db.ExecContext(ctx, deleteQuiz, arg.QuizID, arg.Language)
	return err
}

const deleteReport = `-- name: DeleteReport :exec
DELETE FROM "report"
WHERE report_id = $1
`

func (q *Queries) DeleteReport(ctx context.Context, reportID string) error {
	_, err := q.db.ExecContext(ctx, deleteReport, reportID)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "user"
WHERE user_id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, userID string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, userID)
	return err
}

const findQuiz = `-- name: FindQuiz :one
SELECT quiz_id, language, status, diagram_type, level, title, text, diagram, likes, author_id, created_at, updated_at FROM "quiz"
WHERE quiz_id = $1 and language = $2
`

type FindQuizParams struct {
	QuizID   string `json:"quiz_id"`
	Language string `json:"language"`
}

func (q *Queries) FindQuiz(ctx context.Context, arg FindQuizParams) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, findQuiz, arg.QuizID, arg.Language)
	var i Quiz
	err := row.Scan(
		&i.QuizID,
		&i.Language,
		&i.Status,
		&i.DiagramType,
		&i.Level,
		&i.Title,
		&i.Text,
		&i.Diagram,
		&i.Likes,
		&i.AuthorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findReport = `-- name: FindReport :many
SELECT report_id, quiz_id, language, author_id, title, text, diagram, comment, created_at FROM "report"
WHERE quiz_id = $1 and language = $2
`

type FindReportParams struct {
	QuizID   sql.NullString `json:"quiz_id"`
	Language sql.NullString `json:"language"`
}

func (q *Queries) FindReport(ctx context.Context, arg FindReportParams) ([]Report, error) {
	rows, err := q.db.QueryContext(ctx, findReport, arg.QuizID, arg.Language)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Report{}
	for rows.Next() {
		var i Report
		if err := rows.Scan(
			&i.ReportID,
			&i.QuizID,
			&i.Language,
			&i.AuthorID,
			&i.Title,
			&i.Text,
			&i.Diagram,
			&i.Comment,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findUser = `-- name: FindUser :one
SELECT user_id, username, password, email, profile, created_at, updated_at, membership, liked_quiz_ids, quiz_history FROM "user"
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) FindUser(ctx context.Context, userID string) (User, error) {
	row := q.db.QueryRowContext(ctx, findUser, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Profile,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Membership,
		&i.LikedQuizIds,
		&i.QuizHistory,
	)
	return i, err
}

const listQuizzesAll = `-- name: ListQuizzesAll :many
SELECT quiz_id, language, status, diagram_type, level, title, text, diagram, likes, author_id, created_at, updated_at FROM "quiz"
WHERE language = $1 and status = $2
`

type ListQuizzesAllParams struct {
	Language string `json:"language"`
	Status   int32  `json:"status"`
}

func (q *Queries) ListQuizzesAll(ctx context.Context, arg ListQuizzesAllParams) ([]Quiz, error) {
	rows, err := q.db.QueryContext(ctx, listQuizzesAll, arg.Language, arg.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Quiz{}
	for rows.Next() {
		var i Quiz
		if err := rows.Scan(
			&i.QuizID,
			&i.Language,
			&i.Status,
			&i.DiagramType,
			&i.Level,
			&i.Title,
			&i.Text,
			&i.Diagram,
			&i.Likes,
			&i.AuthorID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listQuizzesByUser = `-- name: ListQuizzesByUser :many
SELECT quiz_id, language, status, diagram_type, level, title, text, diagram, likes, author_id, created_at, updated_at FROM "quiz"
WHERE author_id = $1 and language = $2 and status = $3
`

type ListQuizzesByUserParams struct {
	AuthorID sql.NullString `json:"author_id"`
	Language string         `json:"language"`
	Status   int32          `json:"status"`
}

func (q *Queries) ListQuizzesByUser(ctx context.Context, arg ListQuizzesByUserParams) ([]Quiz, error) {
	rows, err := q.db.QueryContext(ctx, listQuizzesByUser, arg.AuthorID, arg.Language, arg.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Quiz{}
	for rows.Next() {
		var i Quiz
		if err := rows.Scan(
			&i.QuizID,
			&i.Language,
			&i.Status,
			&i.DiagramType,
			&i.Level,
			&i.Title,
			&i.Text,
			&i.Diagram,
			&i.Likes,
			&i.AuthorID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateQuiz = `-- name: UpdateQuiz :one
UPDATE "quiz" SET 
  language = $2, status = $3, diagram_type = $4, level = $5, title = $6, text = $7, diagram = $8, likes = $9, updated_at = $10
WHERE quiz_id = $1
RETURNING quiz_id, language, status, diagram_type, level, title, text, diagram, likes, author_id, created_at, updated_at
`

type UpdateQuizParams struct {
	QuizID      string         `json:"quiz_id"`
	Language    string         `json:"language"`
	Status      int32          `json:"status"`
	DiagramType int32          `json:"diagram_type"`
	Level       sql.NullString `json:"level"`
	Title       sql.NullString `json:"title"`
	Text        sql.NullString `json:"text"`
	Diagram     sql.NullString `json:"diagram"`
	Likes       sql.NullInt32  `json:"likes"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func (q *Queries) UpdateQuiz(ctx context.Context, arg UpdateQuizParams) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, updateQuiz,
		arg.QuizID,
		arg.Language,
		arg.Status,
		arg.DiagramType,
		arg.Level,
		arg.Title,
		arg.Text,
		arg.Diagram,
		arg.Likes,
		arg.UpdatedAt,
	)
	var i Quiz
	err := row.Scan(
		&i.QuizID,
		&i.Language,
		&i.Status,
		&i.DiagramType,
		&i.Level,
		&i.Title,
		&i.Text,
		&i.Diagram,
		&i.Likes,
		&i.AuthorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateQuizID = `-- name: UpdateQuizID :one
UPDATE "quiz" SET 
  quiz_id = $1
WHERE language = $2 and title = $3 and created_at = $4
RETURNING quiz_id, language, status, diagram_type, level, title, text, diagram, likes, author_id, created_at, updated_at
`

type UpdateQuizIDParams struct {
	QuizID    string         `json:"quiz_id"`
	Language  string         `json:"language"`
	Title     sql.NullString `json:"title"`
	CreatedAt time.Time      `json:"created_at"`
}

func (q *Queries) UpdateQuizID(ctx context.Context, arg UpdateQuizIDParams) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, updateQuizID,
		arg.QuizID,
		arg.Language,
		arg.Title,
		arg.CreatedAt,
	)
	var i Quiz
	err := row.Scan(
		&i.QuizID,
		&i.Language,
		&i.Status,
		&i.DiagramType,
		&i.Level,
		&i.Title,
		&i.Text,
		&i.Diagram,
		&i.Likes,
		&i.AuthorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateQuizLike = `-- name: UpdateQuizLike :exec
UPDATE "quiz" SET 
  likes = likes + $3
WHERE quiz_id = $1 and language = $2
RETURNING quiz_id, language, status, diagram_type, level, title, text, diagram, likes, author_id, created_at, updated_at
`

type UpdateQuizLikeParams struct {
	QuizID   string        `json:"quiz_id"`
	Language string        `json:"language"`
	Likes    sql.NullInt32 `json:"likes"`
}

func (q *Queries) UpdateQuizLike(ctx context.Context, arg UpdateQuizLikeParams) error {
	_, err := q.db.ExecContext(ctx, updateQuizLike, arg.QuizID, arg.Language, arg.Likes)
	return err
}

const updateReport = `-- name: UpdateReport :one
UPDATE "report" SET 
  title = $2, text = $3, diagram = $4, comment = $5
WHERE report_id = $1
RETURNING report_id, quiz_id, language, author_id, title, text, diagram, comment, created_at
`

type UpdateReportParams struct {
	ReportID string         `json:"report_id"`
	Title    sql.NullString `json:"title"`
	Text     sql.NullString `json:"text"`
	Diagram  sql.NullString `json:"diagram"`
	Comment  sql.NullString `json:"comment"`
}

func (q *Queries) UpdateReport(ctx context.Context, arg UpdateReportParams) (Report, error) {
	row := q.db.QueryRowContext(ctx, updateReport,
		arg.ReportID,
		arg.Title,
		arg.Text,
		arg.Diagram,
		arg.Comment,
	)
	var i Report
	err := row.Scan(
		&i.ReportID,
		&i.QuizID,
		&i.Language,
		&i.AuthorID,
		&i.Title,
		&i.Text,
		&i.Diagram,
		&i.Comment,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE "user" SET 
  username = $2, password = $3, email = $4, profile = $5, membership = $6, updated_at = $7
WHERE user_id = $1
RETURNING user_id, username, password, email, profile, created_at, updated_at, membership, liked_quiz_ids, quiz_history
`

type UpdateUserParams struct {
	UserID     string         `json:"user_id"`
	Username   string         `json:"username"`
	Password   string         `json:"password"`
	Email      string         `json:"email"`
	Profile    sql.NullString `json:"profile"`
	Membership sql.NullInt32  `json:"membership"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.UserID,
		arg.Username,
		arg.Password,
		arg.Email,
		arg.Profile,
		arg.Membership,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Profile,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Membership,
		&i.LikedQuizIds,
		&i.QuizHistory,
	)
	return i, err
}
