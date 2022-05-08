CREATE TABLE "user" (
    user_id        varchar(36)       PRIMARY KEY  NOT NULL,
    username       varchar(255)      NOT NULL,
    password       varchar(255)      NOT NULL,
    email          varchar(255)      NOT NULL,
    profile        text,
    created_at     datetime,
    updated_at     datetime,
    membership     integer,
    liked_quiz_ids text,
    quiz_history   text
);

CREATE TABLE "quiz" (
    quiz_id        varchar(36)  PRIMARY KEY,
    language       varchar(8)   PRIMARY KEY  NOT NULL,
    status         integer      NOT NULL,
    diagram_type   integer      NOT NULL,
    level          varchar(8),
    title          text,
    text           text,
    diagram        text,
    likes          integer,
    author_id      varchar(36),
    created_at     datetime,
    updated_at     datetime
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
    created_at     datetime
);