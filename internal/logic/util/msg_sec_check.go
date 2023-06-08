package util

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
	"github.com/xh-polaris/meowchat-bff/internal/errorx"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	pb2 "github.com/xh-polaris/sts-rpc/pb"
)

const (
	TextSecUrl = "https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%s"
)

func MsgSecCheck(ctx context.Context, svc *svc.ServiceContext, content string, openId string, scene int) (err error) {

	var j map[string]any
	j = make(map[string]any)
	j["content"] = content
	j["version"] = 2
	j["scene"] = scene
	j["openid"] = openId
	data, err := json.Marshal(j)
	if err != nil {
		return err
	}

	r, err := svc.StsRPC.GetAccessToken(ctx, &pb2.GetAccessTokenReq{
		App: "meowchat",
	})
	if err != nil {
		return err
	}

	res, err := util.HTTPPostContext(ctx, fmt.Sprintf(TextSecUrl, r.AccessToken), data, nil)
	if err != nil {
		return err
	}
	var i map[string]any
	if err = json.Unmarshal(res, &i); err != nil {
		return err
	}
	errcode := int(i["errcode"].(float64))
	if errcode != 0 {
		return errorx.ErrWechatApiWrong
	}
	result := i["result"].(map[string]any)
	suggest := result["suggest"].(string)
	if suggest != "pass" {
		return errorx.ErrMsgNotSec
	}
	return nil
}
