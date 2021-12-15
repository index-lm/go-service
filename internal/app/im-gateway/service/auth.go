package service

import (
	"context"
	"go-service/pkg/pb/auth"
)

type Auth struct{}

func (a Auth) PasswordLogin(ctx context.Context, req auth.LoginReq) (res *auth.LoginRes, err error) {
	return &auth.LoginRes{Code: 0, Success: req.Password}, nil
}
