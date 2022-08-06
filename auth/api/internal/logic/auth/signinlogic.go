package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"sso/auth/api/internal/svc"
	"sso/auth/api/internal/types"
	"sso/auth/model"

	"github.com/ginvcom/util"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type SignInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignInLogic) SignIn(req *types.SignInReq) (resp *types.SignInReply, err error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Mobile)
	if err != nil {
		if err == sqlc.ErrNotFound {
			err = errors.New("mobile or password error")
		}

		return
	}

	password := util.MD5(req.Password + user.Salt)

	if password != user.Password {
		err = errors.New("mobile or password error")
		return
	}

	objectArgs := &model.ObjectFindOneArgs{
		TopKey: req.ServiceCode,
	}

	object, err := l.svcCtx.ObjectModel.FindOne(l.ctx, objectArgs)
	if err != nil {
		fmt.Println(err)
		err = fmt.Errorf("service \"%s\" not exits", req.ServiceCode)
		return
	}

	now := time.Now().Unix()
	seconds := l.svcCtx.Config.Auth.AccessExpire

	if req.Remember == "on" {
		seconds = l.svcCtx.Config.Auth.RememberAccessExpire
	}
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, seconds, user.Id)
	if err != nil {
		return
	}

	resp = &types.SignInReply{
		Redirect:    object.Identifier,
		AccessToken: jwtToken,
		Name:        user.Name,
		Mobile:      user.Mobile,
		Avatar:      user.Avatar,
		Gender:      user.Gender,
	}
	return
}

func (l *SignInLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds // jwt的过期时间，等于签发时间加上配置的过期时间
	claims["iat"] = iat           // jwt的签发时间
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
