package system

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/pb"
	pb2 "github.com/xh-polaris/meowchat-user-rpc/pb"
	"sync"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByRoleLogic {
	return &GetUserByRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByRoleLogic) GetUserByRole(req *types.RetrieveUserPreviewReq) (resp *types.RetrieveUserPreviewResp, err error) {
	resp = new(types.RetrieveUserPreviewResp)

	Userid, err := l.svcCtx.SystemRPC.ListUseridByRole(l.ctx, &pb.ListUseridReq{RoleType: req.RoleType, CommunityId: req.CommunityId})
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	resp.Users = make([]types.UserPreview, len(Userid.UserId))

	wg.Add(len(Userid.UserId))
	var errorChannel = make(chan error, len(Userid.UserId))

	for i, userid := range Userid.UserId {
		go l.GetOneUser(userid, &wg, i, resp.Users, errorChannel)
	}

	wg.Wait()
	if len(errorChannel) != 0 {
		return nil, <-errorChannel
	}

	return
}

func (l *GetUserByRoleLogic) GetOneUser(userid string, wg *sync.WaitGroup, i int, Users []types.UserPreview, chan1 chan error) (err error) {
	defer wg.Done()
	request := &pb2.GetUserReq{
		UserId: userid,
	}
	data, err := l.svcCtx.UserRPC.GetUser(l.ctx, request)
	if err != nil {
		chan1 <- err
		return err
	}
	Users[i] = types.UserPreview{
		Id:        data.User.Id,
		Nickname:  data.User.Nickname,
		AvatarUrl: data.User.AvatarUrl,
	}
	return nil
}
