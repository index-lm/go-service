package rpc

import (
	"context"
	"go-service/pkg/pb/transfer"
)

type Transfer interface {
	PasswordLogin(ctx context.Context, req transfer.LoginReq) (res *transfer.LoginRes, err error)
}
