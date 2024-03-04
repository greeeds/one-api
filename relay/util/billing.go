package util

import (
	"context"
	"github.com/greeeds/one-api/common/logger"
	"github.com/greeeds/one-api/model"
)

func ReturnPreConsumedQuota(ctx context.Context, preConsumedQuota int, tokenId int) {
	if preConsumedQuota != 0 {
		go func(ctx context.Context) {
			// return pre-consumed quota
			err := model.PostConsumeTokenQuota(tokenId, -preConsumedQuota)
			if err != nil {
				logger.Error(ctx, "error return pre-consumed quota: "+err.Error())
			}
		}(ctx)
	}
}
