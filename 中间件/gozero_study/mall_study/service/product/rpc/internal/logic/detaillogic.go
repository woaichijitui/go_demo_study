package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"gozero_mall_study/service/product/model"
	"gozero_mall_study/service/product/rpc/product"

	"gozero_mall_study/service/product/rpc/internal/svc"
	"gozero_mall_study/service/product/rpc/productlclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *productlclient.DetailRequest) (*productlclient.DetailResponse, error) {
	// 查询产品是否存在
	res, err := l.svcCtx.ProductModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &product.DetailResponse{
		Id:     int64(res.Id),
		Name:   res.Name,
		Desc:   res.Desc,
		Stock:  int64(res.Stock),
		Amount: int64(res.Amount),
		Status: int64(res.Status),
	}, nil
}
