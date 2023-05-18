package user

import (
	"context"
	"github.com/xh-polaris/meowchat-like-rpc/likerpc"
	pb4 "github.com/xh-polaris/meowchat-like-rpc/pb"
	pb2 "github.com/xh-polaris/meowchat-moment-rpc/pb"
	pb3 "github.com/xh-polaris/meowchat-post-rpc/pb"
	"github.com/xh-polaris/meowchat-user-rpc/pb"
	"sync"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) getLessDependentInfo(user *types.User) {
	wg := &sync.WaitGroup{}
	wg.Add(4)

	go func() {
		defer wg.Done()
		momentCount, err := l.svcCtx.MomentRPC.CountMoment(l.ctx, &pb2.CountMomentReq{
			FilterOptions: &pb2.FilterOptions{OnlyUserId: &user.Id},
		})
		if err != nil {
			logx.Error(err)
		}
		user.Article += momentCount.Total
	}()

	go func() {
		defer wg.Done()
		postCount, err := l.svcCtx.PostRPC.CountPost(l.ctx, &pb3.CountPostReq{
			FilterOptions: &pb3.FilterOptions{
				OnlyUserId: &user.Id,
			},
		})
		if err != nil {
			logx.Error(err)
		}
		// TODO 偶尔有并发问题
		user.Article += postCount.Total
	}()

	go func() {
		defer wg.Done()
		follower, err := l.svcCtx.LikeRPC.GetTargetLikes(l.ctx, &pb4.GetTargetLikesReq{
			TargetId: user.Id,
			Type:     likerpc.TargetTypeUser,
		})
		if err != nil {
			logx.Error(err)
		}
		user.Follower = follower.Count
	}()

	go func() {
		defer wg.Done()
		followee, err := l.svcCtx.LikeRPC.GetUserLikes(l.ctx, &pb4.GetUserLikesReq{
			UserId:     user.Id,
			TargetType: likerpc.TargetTypeUser,
		})
		if err != nil {
			logx.Error(err)
		}
		user.Following = int64(len(followee.Likes))
	}()

	wg.Wait()
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	resp = new(types.GetUserInfoResp)

	var userId string
	if req.UserId != nil {
		userId = *req.UserId
	} else {
		userId = l.ctx.Value("userId").(string)
	}

	data, err := l.svcCtx.UserRPC.GetUserDetail(l.ctx, &pb.GetUserDetailReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	resp.User = types.User{
		Id:        data.User.Id,
		Nickname:  data.User.Nickname,
		AvatarUrl: data.User.AvatarUrl,
		Motto:     data.User.Motto,
	}

	l.getLessDependentInfo(&resp.User)

	return
}
