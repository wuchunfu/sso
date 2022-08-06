package auth

import (
	"context"
	"net/http"

	"sso/auth/api/internal/svc"
	"sso/auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignOutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignOutLogic {
	return &SignOutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignOutLogic) SignOut(r *http.Request) (resp *types.SignOutReply, err error) {
	// todo: add your logic here and delete this line

	return
}
