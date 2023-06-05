// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	auth "github.com/xh-polaris/meowchat-bff/internal/handler/auth"
	collection "github.com/xh-polaris/meowchat-bff/internal/handler/collection"
	comment "github.com/xh-polaris/meowchat-bff/internal/handler/comment"
	like "github.com/xh-polaris/meowchat-bff/internal/handler/like"
	moment "github.com/xh-polaris/meowchat-bff/internal/handler/moment"
	post "github.com/xh-polaris/meowchat-bff/internal/handler/post"
	sts "github.com/xh-polaris/meowchat-bff/internal/handler/sts"
	system "github.com/xh-polaris/meowchat-bff/internal/handler/system"
	user "github.com/xh-polaris/meowchat-bff/internal/handler/user"
	"github.com/xh-polaris/meowchat-bff/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/sign_in",
				Handler: auth.SignInHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/send_verify_code",
				Handler: auth.SendVerifyCodeHandler(serverCtx),
			},
		},
		rest.WithPrefix("/auth"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/set_password",
				Handler: auth.SetPasswordHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/auth"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/notice/get_admins",
				Handler: system.GetAdminsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notice/new_admin",
				Handler: system.NewAdminHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notice/delete_admin",
				Handler: system.DeleteAdminHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notice/list_apply",
				Handler: system.ListApplyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notice/handle_apply",
				Handler: system.HandleApplyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/notice/get_news",
				Handler: system.GetNewsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notice/new_news",
				Handler: system.NewNewsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notice/remove_news",
				Handler: system.DeleteNewsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/notice/get_notices",
				Handler: system.GetNoticesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notice/new_notice",
				Handler: system.NewNoticeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notice/remove_notice",
				Handler: system.DeleteNoticeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/community/list_community",
				Handler: system.ListCommunityHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/community/new_community",
				Handler: system.NewCommunityHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/community/delete_community",
				Handler: system.DeleteCommunityHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/role/get_user_roles",
				Handler: system.GetUserRolesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/role/update_community_admin",
				Handler: system.UpdateCommunityAdminHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/role/update_super_admin",
				Handler: system.UpdateSuperAdminHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/role/get_user_by_role",
				Handler: system.GetUserByRoleHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/do_like",
				Handler: like.DoLikeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get_user_liked",
				Handler: like.GetUserLikedHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get_count",
				Handler: like.GetLikedCountHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get_liked_users",
				Handler: like.GetLikedUsersHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get_user_likes",
				Handler: like.GetUserLikesHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/like"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/get_post_previews",
				Handler: post.GetPostPreviewsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get_post_detail",
				Handler: post.GetPostDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/new_post",
				Handler: post.NewPostHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete_post",
				Handler: post.DeletePostHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/set_official",
				Handler: post.SetOfficialHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/post"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/get_cat_previews",
				Handler: collection.GetCatPreviewsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get_cat_detail",
				Handler: collection.GetCatDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/new_cat",
				Handler: collection.NewCatHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete_cat",
				Handler: collection.DeleteCatHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/search_cat",
				Handler: collection.SearchCatHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/create_image",
				Handler: collection.CreateImageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete_image",
				Handler: collection.DeleteImageHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get_image_by_cat",
				Handler: collection.GetImageByCatHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/collection"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/get_moment_previews",
				Handler: moment.GetMomentPreviewsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get_moment_detail",
				Handler: moment.GetMomentDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/new_moment",
				Handler: moment.NewMomentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete_moment",
				Handler: moment.DeleteMomentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/search_moment",
				Handler: moment.SearchMomentHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/moment"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/get_comments",
				Handler: comment.GetCommentsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/new_comment",
				Handler: comment.NewCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete_comment",
				Handler: comment.DeleteCommentHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/comment"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/get_user_info",
				Handler: user.GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update_user_info",
				Handler: user.UpdateUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/search_user",
				Handler: user.SearchUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/search_user_for_admin",
				Handler: user.SearchUserForAdminHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/create_apply",
				Handler: user.CreateApplyHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/apply_signed_url",
				Handler: sts.ApplySignedUrlHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/apply_signed_url_as_community",
				Handler: sts.ApplySignedUrlAsCommunityHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/sts"),
	)
}
