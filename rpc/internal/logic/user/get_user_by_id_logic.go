package user

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent"
	"github.com/suyuan32/simple-admin-core/pkg/ent/user"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/pkg/uuidx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *core.UUIDReq) (*core.UserInfo, error) {
	result, err := l.svcCtx.DB.User.Query().Where(user.IDEQ(uuidx.ParseUUIDString(in.Id))).WithRoles().First(l.ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
		case ent.IsConstraintError(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.UpdateFailed)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	if err != nil {
		return nil, err
	}

	return &core.UserInfo{
		Nickname:     result.Nickname,
		Avatar:       result.Avatar,
		RoleIds:      GetRoleIds(result.Edges.Roles),
		Mobile:       result.Mobile,
		Email:        result.Email,
		Status:       uint32(result.Status),
		Id:           result.ID.String(),
		Username:     result.Username,
		HomePath:     result.HomePath,
		Description:  result.Description,
		DepartmentId: result.DepartmentID,
		CreatedAt:    result.CreatedAt.Unix(),
		UpdatedAt:    result.UpdatedAt.Unix(),
	}, nil
}

func GetRoleIds(data []*ent.Role) []uint64 {
	var ids []uint64
	for _, v := range data {
		ids = append(ids, v.ID)
	}
	return ids
}

func GetRoleCodes(data []*ent.Role) []string {
	var codes []string
	for _, v := range data {
		codes = append(codes, v.Code)
	}
	return codes
}
