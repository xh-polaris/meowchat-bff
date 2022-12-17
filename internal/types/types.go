// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
}

type SignInReq struct {
	AuthType string   `json:"authType,options=phone|email|wechat"`
	AuthId   string   `json:"authId"`
	Password string   `json:"password,optional"`
	Params   []string `json:"params,optional"`
}

type SignInResp struct {
	Status
	UserId       string `json:"userId"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

type SetPasswordReq struct {
	Password string `json:"password"`
}

type SetPasswordResp struct {
	Status
}

type SendVerifyCodeReq struct {
	AuthType string `json:"authType,options=phone|email"`
	AuthId   string `json:"authId"`
}

type SendVerifyCodeResp struct {
	Status
}

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type UserPreview struct {
	Id        string `json:"id"`
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatarUrl"`
}

type Tag struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Post struct {
	Id          string      `json:"id"`
	CreateAt    int64       `json:"createAt"`
	IsAnonymous bool        `json:"isAnonymous"`
	Title       string      `json:"title"`
	Text        string      `json:"text"`
	CoverUrl    string      `json:"coverUrl"`
	Tags        []Tag       `json:"tags"`
	Likes       int64       `json:"likes"`
	Comments    int64       `json:"comments"`
	User        UserPreview `json:"user"`
	Status      int64       `json:"status"`
}

type GetPostPreviewsReq struct {
	Page int64 `form:"page"`
}

type GetPostPreviewsResp struct {
	Status
	Posts []Post `json:"posts"`
}

type GetPostDetailReq struct {
	PostId string `form:"postId"`
}

type GetPostDetailResp struct {
	Post Post `json:"post"`
	Status
}

type NewPostReq struct {
	Id          string `json:"id,optional"`
	IsAnonymous bool   `json:"isAnonymous"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	CoverUrl    string `json:"coverUrl,optional"`
	Tags        []Tag  `json:"tags"`
}

type NewPostResp struct {
	Status
}

type DeletePostReq struct {
	Id string `json:"id"`
}

type DeletePostResp struct {
	Status
}

type Cat struct {
	Id           string   `json:"id"`
	CreateAt     int64    `json:"createAt"`
	Age          string   `json:"age"`
	CommunityId  string   `json:"communityId"`
	Color        string   `json:"color"`
	Details      string   `json:"details"`
	Name         string   `json:"name"`
	Popularity   int64    `json:"popularity"`
	Sex          string   `json:"sex"`
	Status       int32    `json:"status"`
	Area         string   `json:"area"`
	IsSnipped    bool     `json:"isSnipped"`
	IsSterilized bool     `json:"isSterilized"`
	Avatars      []string `json:"avatars"`
}

type CatPreview struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Area        string `json:"area"`
	Color       string `json:"color"`
	AvatarUrl   string `json:"avatarUrl"`
	IsCollected bool   `json:"isCollected"`
}

type GetCatPreviewsReq struct {
	CommunityId string `form:"communityId"`
	Page        int64  `form:"page"`
}

type GetCatPreviewsResp struct {
	Status
	Cats []CatPreview `json:"cats"`
}

type GetCatDetailReq struct {
	CatId string `form:"catId"`
}

type GetCatDetailResp struct {
	Status
	Cat Cat `json:"cat"`
}

type DeleteCatReq struct {
	CatId string `json:"catId"`
}

type DeleteCatResp struct {
	Status
}

type NewCatReq struct {
	Id           string   `json:"id,optional"`
	Age          string   `json:"age"`
	CommunityId  string   `json:"communityId"`
	Color        string   `json:"color"`
	Details      string   `json:"details"`
	Name         string   `json:"name"`
	Popularity   int64    `json:"popularity"`
	Sex          string   `json:"sex"`
	Area         string   `json:"area"`
	IsSnipped    bool     `json:"isSnipped"`
	IsSterilized bool     `json:"isSterilized"`
	Avatars      []string `json:"avatars"`
}

type NewCatResp struct {
	Status
	CatId string `json:"catId"`
}

type Moment struct {
	Id          string      `json:"id"`
	CreateAt    int64       `json:"createAt"`
	CatId       string      `json:"catId,optional"`
	Photos      []string    `json:"photos"`
	Title       string      `json:"title"`
	Text        string      `json:"text"`
	User        UserPreview `json:"user"`
	CommunityId string      `json:"communityId"`
}

type GetMomentPreviewsReq struct {
	CommunityId string `form:"communityId"`
	Page        int64  `form:"page"`
}

type GetMomentPreviewsResp struct {
	Status
	Moments []Moment `json:"moments"`
}

type GetMomentDetailReq struct {
	MomentId string `form:"momentId"`
}

type GetMomentDetailResp struct {
	Status
	Moment Moment `json:"moment"`
}

type DeleteMomentReq struct {
	MomentId string `json:"momentId"`
}

type DeleteMomentResp struct {
	Status
}

type NewMomentReq struct {
	Id          string   `json:"id,optional"`
	Title       string   `json:"title"`
	CatId       string   `json:"catId,optional"`
	Text        string   `json:"text"`
	Photos      []string `json:"photos"`
	CommunityId string   `json:"communityId"`
}

type NewMomentResp struct {
	MomentId string `json:"momentId"`
	Status
}

type Comment struct {
	Id       string      `json:"id"`
	CreateAt int64       `json:"createAt"`
	Text     string      `json:"text"`
	User     UserPreview `json:"user"`
	Comments int64       `json:"comments"`
	ReplyId  string      `json:"replyName,optional"`
}

type NewCommentReq struct {
	Text    string `json:"text"`
	OwnerId string `json:"ownerId"`
	Scope   string `json:"scope"`
}

type NewCommentResp struct {
	Status
}

type GetCommentsReq struct {
	Scope string `form:"scope"`
	Page  int64  `form:"page"`
	Id    string `form:"id"`
}

type GetCommentsResp struct {
	Status
	Comments []Comment `json:"comments"`
}

type News struct {
	Id          string `json:"id"`
	CreateAt    int64  `json:"createAt"`
	CommunityId string `json:"communityId"`
	ImageUrl    string `json:"imageUrl"`
	LinkUrl     string `json:"linkUrl"`
	Type        string `json:"type"`
}

type Admin struct {
	Id          string      `json:"id"`
	CreateAt    int64       `json:"createAt"`
	CommunityId string      `json:"communityId"`
	Name        string      `json:"name"`
	Phone       string      `json:"phone"`
	User        UserPreview `json:"user"`
	Wechat      string      `json:"wechat"`
}

type Notice struct {
	Id       string `json:"id"`
	Text     string `json:"text"`
	CreateAt int64  `json:"createAt"`
}

type GetNewsReq struct {
	CommunityId string `form:"communityId"`
}

type GetNewsResp struct {
	Status
	News []News `json:"news"`
}

type GetAdminsReq struct {
	CommunityId string `form:"communityId"`
}

type GetAdminsResp struct {
	Status
	Admins []Admin `json:"admins"`
}

type GetNoticesReq struct {
	CommunityId string `form:"communityId"`
}

type GetNoticesResp struct {
	Status
	Notices []Notice `json:"notices"`
}

type NewNoticeReq struct {
	Id          string `json:"id,optional"`
	CommunityId string `json:"communityId,optional"`
	Text        string `json:"text"`
}

type NewNoticeResp struct {
	Status
	NoticeId string `json:"noticeId"`
}

type NewNewsReq struct {
	Id          string `json:"id,optional"`
	CommunityId string `json:"communityId,optional"`
	ImageUrl    string `json:"imageUrl"`
	LinkUrl     string `json:"linkUrl"`
	Type        string `json:"type"`
}

type NewNewsResp struct {
	Status
	NewId string `json:"newId"`
}

type DeleteNoticeReq struct {
	Id string `json:"id"`
}

type DeleteNoticeResp struct {
	Status
}

type DeleteNewsReq struct {
	Id string `json:"id"`
}

type DeleteNewsResp struct {
	Status
}

type DoLikeReq struct {
	TargetId   string `json:"targetId"`
	TargetType int64  `json:"type"`
}

type DoLikeResp struct {
	Status
}

type GetUserLikedReq struct {
	TargetId   string `form:"targetId"`
	TargetType int64  `form:"type"`
}

type GetUserLikedResp struct {
	Status
	Liked bool `json:"liked"`
}

type GetLikedCountReq struct {
	TargetId   string `form:"targetId"`
	TargetType int64  `form:"type"`
}

type GetLikedCountResp struct {
	Status
	Count int64 `json:"count"`
}

type GetUserInfoReq struct {
}

type GetUserInfoResp struct {
	Status
	User UserPreview `json:"user"`
}

type UpdateUserInfoReq struct {
	AvatarUrl string `json:"avatarUrl,optional"`
	Nickname  string `json:"nickname,optional"`
}

type UpdateUserInfoResp struct {
	Status
}

type ApplyTokenReq struct {
	Path string `form:"path"`
}

type ApplyTokenResp struct {
	Status
	SecretId     string `json:"secretId"`
	SecretKey    string `json:"secretKey"`
	SessionToken string `json:"sessionToken"`
}
