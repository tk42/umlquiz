package infra

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"github.com/tk42/umlquiz/backend/gen/sqlc"
	"github.com/tk42/victolinux/env"
)

type SqlHandler struct {
	DB      *sql.DB
	Queries *sqlc.Queries
}

func NewSqlHandler() SqlHandler {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			env.GetString("DB_HOST", "db"),
			env.GetString("DB_PORT", "5432"),
			env.GetString("DB_USER", "postgres"),
			env.GetString("DB_PASS", "e8a48653851e28c69d0506508fb27fc5"),
			env.GetString("DB_NAME", "umlquiz"),
		),
	)
	// defer db.Close() // 'database is closed' occurs
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	return SqlHandler{
		DB:      db,
		Queries: sqlc.New(db),
	}
}

func (s *SqlHandler) Migrate(embedMigrations embed.FS) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	goose.SetBaseFS(embedMigrations)

	if _, err := fs.Stat(embedMigrations, "api"); err != nil {
		return err
	}
	if err := goose.Up(s.DB, "api"); err != nil {
		return err
	}
	if err := goose.Version(s.DB, "api"); err != nil {
		return err
	}
	return nil
}
