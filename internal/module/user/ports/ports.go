package ports

import (
	"codebase-app/internal/module/user/entity"
	"context"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*entity.UserResult, error)
	FindById(ctx context.Context, id string) (*entity.UserResult, error)

	GetCouriers(ctx context.Context, req *entity.GetCouriersReq) (*entity.GetCouriersRes, error)
}

type UserService interface {
	Login(ctx context.Context, req *entity.LoginReq) (*entity.LoginResp, error)
	GetCouriers(ctx context.Context, req *entity.GetCouriersReq) (*entity.GetCouriersRes, error)
}
