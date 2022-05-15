// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package sqlc

import (
	"database/sql"
	"time"
)

type Quiz struct {
	QuizID      string         `json:"quiz_id"`
	Language    string         `json:"language"`
	Status      int32          `json:"status"`
	DiagramType int32          `json:"diagram_type"`
	Level       sql.NullString `json:"level"`
	Title       sql.NullString `json:"title"`
	Text        sql.NullString `json:"text"`
	Diagram     sql.NullString `json:"diagram"`
	Likes       sql.NullInt32  `json:"likes"`
	AuthorID    sql.NullString `json:"author_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type Report struct {
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

type User struct {
	UserID       string         `json:"user_id"`
	Username     string         `json:"username"`
	Password     string         `json:"password"`
	Email        string         `json:"email"`
	Profile      sql.NullString `json:"profile"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	Membership   sql.NullInt32  `json:"membership"`
	LikedQuizIds sql.NullString `json:"liked_quiz_ids"`
	QuizHistory  sql.NullString `json:"quiz_history"`
}
