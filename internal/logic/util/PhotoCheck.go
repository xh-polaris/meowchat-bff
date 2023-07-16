package util

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/xh-polaris/meowchat-bff/internal/errorx"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
)

func PhotoCheck(ctx context.Context, svc *svc.ServiceContext, urls []string) (err error) {

	var input []cos.ImageAuditingInputOptions
	for key, url := range urls {
		input = append(input, cos.ImageAuditingInputOptions{
			DataId: string(key),
			Url:    url,
		})
	}
	opt := &cos.BatchImageAuditingOptions{
		Input: input,
		Conf:  &cos.ImageAuditingJobConf{},
	}
	res, _, err := svc.CiClient.CI.BatchImageAuditing(ctx, opt)
	if err != nil {
		return err
	}

	for key := range urls {
		if res.JobsDetail[key].Result != 0 {
			host := "https://" + svc.Config.CdnHost
			for _, url := range urls {
				name := url[len(host)+1:]
				_, err = svc.CiClient.Object.Delete(ctx, name)
				if err != nil {
					return err
				}
			}
			return errorx.ErrPhotoNotSec
		}
	}
	return nil
}
