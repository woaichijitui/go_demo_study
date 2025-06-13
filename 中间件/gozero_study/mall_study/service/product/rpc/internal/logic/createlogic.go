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

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *productlclient.CreateRequest) (*productlclient.CreateResponse, error) {
	newProduct := model.Product{
		Name:   in.Name,
		Desc:   in.Desc,
		Stock:  uint64(in.Stock),
		Amount: uint64(in.Amount),
		Status: uint64(in.Status),
	}

	res, err := l.svcCtx.ProductModel.Insert(l.ctx, &newProduct)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	id, err := res.LastInsertId()
	newProduct.Id = uint64(id)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &product.CreateResponse{
		Id: int64(newProduct.Id),
	}, nil
}
