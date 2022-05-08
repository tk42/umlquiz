package infra

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/tk42/umlquiz/backend/gen/sqlc"
	"github.com/tk42/victolinux/env"
)

type SqlHandler struct {
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
	defer db.Close()
	if err != nil {
		panic(err)
	}

	return SqlHandler{
		Queries: sqlc.New(db),
	}
}
