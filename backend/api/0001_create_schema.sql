-- +goose Up
CREATE TABLE "user" (
    user_id        varchar(36)       PRIMARY KEY  NOT NULL,
    username       varchar(255)      NOT NULL,
    password       varchar(255)      NOT NULL,
    email          varchar(255)      NOT NULL,
    profile        text,
    created_at     timestamp      NOT NULL,
    updated_at     timestamp      NOT NULL,
    membership     integer,
    liked_quiz_ids text,
    quiz_history   text
);
CREATE TABLE "quiz" (
    quiz_id        varchar(36) ,
    language       varchar(8)   NOT NULL,
    status         integer      NOT NULL,
    diagram_type   integer      NOT NULL,
    level          varchar(8),
    title          text,
    text           text,
    diagram        text,
    likes          integer,
    author_id      varchar(36),
    created_at     timestamp      NOT NULL,
    updated_at     timestamp      NOT NULL,
    PRIMARY KEY (quiz_id, language)
);
CREATE TABLE "report" (
    report_id      varchar(36)  PRIMARY KEY  NOT NULL,
    quiz_id        varchar(36),
    language       varchar(4),
    author_id      varchar(36),
    title          text,
    text           text,
    diagram        text,
    comment        text,
    created_at     timestamp       NOT NULL
);

-- +goose Down
DROP TABLE "user";
DROP TABLE "quiz";
DROP TABLE "report";
