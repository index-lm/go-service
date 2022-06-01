package service

import (
	"context"
	"go-service/pkg/pb/transfer"
)

type Auth struct{}

func (a Auth) PasswordLogin(ctx context.Context, req *transfer.LoginReq) (res *transfer.LoginRes, err error) {
	return &transfer.LoginRes{Code: 0, Success: req.Password}, nil
}
