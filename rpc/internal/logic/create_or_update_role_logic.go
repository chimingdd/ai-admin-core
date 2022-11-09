package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/msg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateRoleLogic {
	return &CreateOrUpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// role service
func (l *CreateOrUpdateRoleLogic) CreateOrUpdateRole(in *core.RoleInfo) (*core.BaseResp, error) {
	if in.Id == 0 {
		err := l.svcCtx.DB.Role.Create().
			SetName(in.Name).
			SetValue(in.Value).
			SetDefaultRouter(in.DefaultRouter).
			SetStatus(uint8(in.Status)).
			SetOrderNo(in.OrderNo).
			SetRemark(in.Remark).
			Exec(l.ctx)

		if err != nil {
			switch {
			case ent.IsConstraintError(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(i18n.DuplicateRoleValue)
			default:
				logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(errorx.DatabaseError)
			}
		}

		err = l.UpdateRoleInfoInRedis()
		if err != nil {
			logx.Errorw("fail to update the role info in Redis", logx.Field("detail", err.Error()))
			return nil, err
		}

		return &core.BaseResp{Msg: errorx.CreateSuccess}, nil
	} else {
		err := l.svcCtx.DB.Role.UpdateOneID(in.Id).
			SetName(in.Name).
			SetValue(in.Value).
			SetDefaultRouter(in.DefaultRouter).
			SetStatus(uint8(in.Status)).
			SetOrderNo(in.OrderNo).
			SetRemark(in.Remark).
			Exec(l.ctx)

		if err != nil {
			switch {
			case ent.IsNotFound(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(errorx.TargetNotExist)
			case ent.IsConstraintError(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(i18n.DuplicateRoleValue)
			default:
				logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(errorx.DatabaseError)
			}
		}

		err = l.UpdateRoleInfoInRedis()
		if err != nil {
			logx.Errorw("fail to update the role info in Redis", logx.Field("detail", err.Error()))
			return nil, err
		}

		return &core.BaseResp{Msg: errorx.UpdateSuccess}, nil
	}
}

func (l *CreateOrUpdateRoleLogic) UpdateRoleInfoInRedis() error {
	roles, err := l.svcCtx.DB.Role.Query().All(l.ctx)

	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Error(err.Error())
			return statuserr.NewInvalidArgumentError(errorx.TargetNotExist)
		case ent.IsConstraintError(err):
			logx.Error(err.Error())
			return statuserr.NewInvalidArgumentError(errorx.UpdateFailed)
		default:
			logx.Errorw(errorx.DatabaseError, logx.Field("detail", err.Error()))
			return statuserr.NewInternalError(errorx.DatabaseError)
		}
	}

	for _, v := range roles {
		err := l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d", v.ID), v.Name)
		err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d_value", v.ID), v.Value)
		err = l.svcCtx.Redis.Hset("roleData", fmt.Sprintf("%d_status", v.ID), strconv.Itoa(int(v.Status)))
		if err != nil {
			return statuserr.NewInternalError(errorx.RedisError)
		}
	}
	return nil
}