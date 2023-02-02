package system

import (
	"context"
	"github.com/xh-polaris/meowchat-bff/internal/errorx"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/constant"
	"github.com/xh-polaris/meowchat-system-rpc/pb"
)

// checkNewsPermission
//  @Description: 验证用户是否有权限操作news
//  @param ctx 需要包含userId
//  @param svcCtx 包含SystemRPC
//  @param newsId
//  @return error 如果没有权限则返回error
//
func checkNewsPermission(ctx context.Context, svcCtx *svc.ServiceContext, newsId string) error {
	news, err := svcCtx.SystemRPC.RetrieveNews(ctx, &pb.RetrieveNewsReq{
		Id: newsId,
	})
	if err != nil {
		return err
	}
	communityId := news.News.CommunityId
	return checkCommunityPermission(ctx, svcCtx, communityId)
}

// checkUserPermission
//  @Description: 验证用户是否有权限操作notice
//  @param ctx 需要包含userId
//  @param svcCtx 包含SystemRPC
//  @param communityId
//  @return error 如果没有权限则返回error
func checkNoticePermission(ctx context.Context, svcCtx *svc.ServiceContext, noticeId string) error {
	notice, err := svcCtx.SystemRPC.RetrieveNotice(ctx, &pb.RetrieveNoticeReq{
		Id: noticeId,
	})
	if err != nil {
		return err
	}
	communityId := notice.Notice.CommunityId
	return checkCommunityPermission(ctx, svcCtx, communityId)
}

// checkCommunityPermission
//  @Description: 验证用户是否是某个社区的社区管理员
//  @param ctx 需要包含userId
//  @param svcCtx 包含SystemRPC
//  @param communityId
//  @return error 如果没有权限则返回error
func checkCommunityPermission(ctx context.Context, svcCtx *svc.ServiceContext, communityId string) error {
	userId := ctx.Value("userId").(string)
	resp, err := svcCtx.SystemRPC.ContainsRole(ctx, &pb.ContainsRoleReq{
		UserId: userId,
		Role: &pb.Role{
			Type:        constant.RoleCommunityAdmin,
			CommunityId: communityId,
		},
		AllowSuperAdmin: true,
	})
	if err != nil || !resp.Contains {
		return errorx.NewForbiddenError("您没有权限进行此操作")
	}
	return nil
}
