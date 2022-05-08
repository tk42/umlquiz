-- name: AddUser :one
INSERT INTO "user" (
  user_id, username, password, email, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: FindUser :one
SELECT * FROM "user"
WHERE user_id = $1 LIMIT 1;


-- name: UpdateUser :one
UPDATE "user" SET 
  username = $2, password = $3, email = $4, profile = $5, membership = $6, updated_at = $7
WHERE user_id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE user_id = $1;

-- name: AddQuiz :one
INSERT INTO "quiz" (
  language, status, diagram_type, level, title, text, diagram, author_id, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: FindQuiz :one
SELECT * FROM "quiz"
WHERE quiz_id = $1 and language = $2;

-- name: UpdateQuizID :one
UPDATE "quiz" SET 
  quiz_id = $1
WHERE language = $2 and title = $3 and created_at = $4
RETURNING *;

-- name: UpdateQuiz :one
UPDATE "quiz" SET 
  language = $2, status = $3, diagram_type = $4, level = $5, title = $6, text = $7, diagram = $8, likes = $9, updated_at = $10
WHERE quiz_id = $1
RETURNING *;

-- name: UpdateQuizLike :exec
UPDATE "quiz" SET 
  likes = likes + $3
WHERE quiz_id = $1 and language = $2
RETURNING *;


-- name: DeleteQuiz :exec
DELETE FROM "quiz"
WHERE quiz_id = $1 and language = $2;


-- name: ListQuizzesAll :many
SELECT * FROM "quiz"
WHERE language = $1 and status = $2;


-- name: ListQuizzesByUser :many
SELECT * FROM "quiz"
WHERE author_id = $1 and language = $2 and status = $3;

-- name: AddReport :one
INSERT INTO "report" (
  report_id, quiz_id, language, author_id, title, text, diagram, comment, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: FindReport :many
SELECT * FROM "report"
WHERE quiz_id = $1 and language = $2;

-- name: UpdateReport :one
UPDATE "report" SET 
  title = $2, text = $3, diagram = $4, comment = $5
WHERE report_id = $1
RETURNING *;

-- name: DeleteReport :exec
DELETE FROM "report"
WHERE report_id = $1;
