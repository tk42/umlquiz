package domain

import (
	"context"

	autogen "github.com/tk42/umlquiz/backend/gen/proto/golang/github.com/tk42/umlquiz"
)

type IUserRepository interface {
	Create(context.Context, *autogen.AddUserRequest) (*autogen.User, error)
	Update(context.Context, *autogen.UpdateUserRequest) (*autogen.User, error)
	Find(context.Context, *autogen.UserRequest) (*autogen.User, error)
	Delete(context.Context, *autogen.UserRequest) error
}

type IQuizRepository interface {
	Create(context.Context, *autogen.AddQuizRequest) (*autogen.Quiz, error)
	Update(context.Context, *autogen.UpdateQuizRequest) (*autogen.Quiz, error)
	UpdateLike(context.Context, *autogen.LikeQuizRequest) error
	Find(context.Context, *autogen.FindQuizRequest) (*autogen.Quiz, error)
	Delete(context.Context, *autogen.DeleteQuizRequest) error
}

type IReportRepository interface {
	Create(context.Context, *autogen.AddReportRequest) (*autogen.Report, error)
	Update(context.Context, *autogen.UpdateReportRequest) (*autogen.Report, error)
	Find(context.Context, *autogen.FindReportsRequest) ([]*autogen.Report, error)
	Delete(context.Context, *autogen.DeleteReportRequest) error
}
