package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"gozero_mall_study/service/user/model"
	"gozero_mall_study/service/user/rpc/user"

	"gozero_mall_study/service/user/rpc/internal/svc"
	"gozero_mall_study/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *userclient.UserInfoRequest) (*userclient.UserInfoResponse, error) {
	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &user.UserInfoResponse{
		Id:     int64(res.Id),
		Name:   res.Name,
		Gender: int64(res.Gender),
		Mobile: res.Mobile,
	}, nil
}
