// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id             string `json:"id"`
	CreateAt       string `json:"createAt"`
	Nickname       string `json:"nickname"`
	AvatarUrl      string `json:"avatarUrl"`
	PrivilegeLevel int32  `json:"privilegeLevel"`
}

type GetUserDetailReq struct {
	Id string `json:"id"`
}

type GetUserDetailResp struct {
	Status
	User User `json:"user"`
}

type UpdateUserDetailReq struct {
	Nickname  string `json:"nickname,optional"`
	AvatarUrl string `json:"avatarUrl,optional"`
}

type UpdateUserDetailResp struct {
	Status
}

type UpdateUserPrivilegeReq struct {
	Id             string `json:"id"`
	PrivilegeLevel int32  `json:"privilegeLevel"`
}

type UpdateUserPrivilegeResp struct {
	Status
}

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type UserVO struct {
	UserId    string `json:"userId"`
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatarUrl"`
}

type Cat struct {
	Id            string `json:"id"`
	CreateAt      string `json:"createAt"`
	Age           string `json:"age"`
	CampusId      int32  `json:"campusId"`
	ColorId       int32  `json:"colorId"`
	Details       string `json:"details"`
	Name          string `json:"name"`
	Popularity    int32  `json:"popularity"`
	Sex           string `json:"sex"`
	StatusId      int32  `json:"statusId"`
	Area          string `json:"area"`
	SnipEar       bool   `json:"snipEar"`
	Sterilization bool   `json:"sterilization"`
	AvatarUrl     string `json:"avatarUrl"`
}

type GetCatDetailReq struct {
	Id string `json:"id"`
}

type GetCatDetailResp struct {
	Status
	Cat Cat `json:"cat"`
}

type QueryCatReq struct {
	CampusId int32 `json:"campusId"`
	Page     int32 `json:"page"`
	Size     int32 `json:"size"`
}

type QueryCatResp struct {
	Status
	Cats []Cat `json:"cats"`
}

type UploadCatReq struct {
	Cat Cat `json:"cat"`
}

type UploadCatResp struct {
	Status
}

type Moment struct {
	Id        string   `json:"id"`
	CreateAt  string   `json:"createAt"`
	CatId     string   `json:"catId"`
	Photos    []string `json:"photos"`
	Status    int32    `json:"status"`
	Text      string   `json:"text"`
	User      UserVO   `json:"user"`
	UserLikes int32    `json:"userLikes"`
}

type GetMomentDetailReq struct {
	Id string `json:"id"`
}

type GetMomentDetailResp struct {
	Status
	Moment Moment `json:"moment"`
}

type QueryMomentReq struct {
	CampusId int32 `json:"campusId"`
	Page     int32 `json:"page"`
	Size     int32 `json:"size"`
}

type QueryMomentResp struct {
	Status
	Moments []Moment `json:"moments"`
}

type UploadMomentReq struct {
	CatId  string   `json:"catId"`
	Text   string   `json:"text"`
	Photos []string `json:"photos"`
}

type UploadMomentResp struct {
	Status
}

type Comment struct {
	Id       string `json:"id"`
	CreateAt string `json:"createAt"`
	Comment  string `json:"comment"`
	User     UserVO `json:"user"`
}

type ListCommentReq struct {
	Id string `json:"id"`
}

type ListCommentResp struct {
	Status
	Comments []Comment `json:"comments"`
}

type PostCommentReq struct {
	Comment string `json:"comment"`
}

type PostCommentResp struct {
	Status
}

type Admin struct {
	Id       string `json:"id"`
	CreateAt string `json:"createAt"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	User     UserVO `json:"user"`
	Wechat   string `json:"wechat"`
}

type GetAdminDetailReq struct {
	Id string `json:"id"`
}

type GetAdminDetailResp struct {
	Status
	Admin Admin `json:"admin"`
}

type ListAdminReq struct {
}

type ListAdminResp struct {
	Status
	Admins []Admin `json:"admins"`
}

type Carousel struct {
	Id       string `json:"id"`
	CreateAt string `json:"createAt"`
	ImgUrl   string `json:"imgUrl"`
	Order    int32  `json:"order"`
	SrcUrl   string `json:"srcUrl"`
	TypeId   int32  `json:"typeId"`
}

type ListCarouselReq struct {
}

type ListCarouselResp struct {
	Status
	CarouselList []Carousel `json:"carouselList"`
}