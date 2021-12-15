package service

import (
	"context"
	"go-service/pkg/log"
	"go-service/pkg/pb/transfer"
)

type Transfer struct{}

func (a Transfer) PasswordLogin(ctx context.Context, req *transfer.LoginReq) (res *transfer.LoginRes, err error) {
	log.Info("test", "sss")
	return &transfer.LoginRes{Code: 0, Success: req.Password}, nil
}
