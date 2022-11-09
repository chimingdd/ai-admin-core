package oauth

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProviderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProviderListLogic {
	return &GetProviderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProviderListLogic) GetProviderList(req *types.PageInfo) (resp *types.ProviderListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetProviderList(l.ctx,
		&core.PageInfoReq{
			Page:     req.Page,
			PageSize: req.PageSize,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.ProviderListResp{}
	resp.Total = data.GetTotal()
	for _, v := range data.Data {
		resp.Data = append(resp.Data,
			types.ProviderInfo{
				BaseInfo: types.BaseInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Name:         v.Name,
				ClientId:     v.ClientId,
				ClientSecret: v.ClientSecret,
				RedirectURL:  v.RedirectUrl,
				Scopes:       v.Scopes,
				AuthURL:      v.AuthUrl,
				TokenURL:     v.TokenUrl,
				InfoURL:      v.InfoUrl,
				AuthStyle:    int(v.AuthStyle),
			})
	}
	return resp, nil
}