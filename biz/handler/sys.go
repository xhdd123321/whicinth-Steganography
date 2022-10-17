package handler

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/redis"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/utils"
)

// GetApiStatistic 获取API调用统计数据
func GetApiStatistic(ctx context.Context, c *app.RequestContext) {
	encodeImageNum, err := strconv.Atoi(redis.GetValue(ctx, ENCODE_IMAGE_KEY))
	if err != nil {
		hlog.CtxErrorf(ctx, "Get encodeImageNum failed, err: %v", err)
	}
	encodeDocNum, err := strconv.Atoi(redis.GetValue(ctx, ENCODE_DOC_KEY))
	if err != nil {
		hlog.CtxErrorf(ctx, "Get encodeDocNum failed, err: %v", err)
	}
	decodeImageNum, err := strconv.Atoi(redis.GetValue(ctx, DECODE_IMAGE_KEY))
	if err != nil {
		hlog.CtxErrorf(ctx, "Get encodeDocNum failed, err: %v", err)
	}
	decodeDocNum, err := strconv.Atoi(redis.GetValue(ctx, DECODE_DOC_KEY))
	if err != nil {
		hlog.CtxErrorf(ctx, "Get encodeDocNum failed, err: %v", err)
	}
	decodeIntelligentNum, err := strconv.Atoi(redis.GetValue(ctx, DECODE_INTELLIGENT_KEY))
	if err != nil {
		hlog.CtxErrorf(ctx, "Get encodeDocNum failed, err: %v", err)
	}
	driftCount := redis.GetDriftCount(ctx)
	// Response
	resp := map[string]interface{}{
		"encode_image_num":       encodeImageNum,
		"encode_doc_num":         encodeDocNum,
		"decode_image_num":       decodeImageNum,
		"decode_doc_num":         decodeDocNum,
		"decode_intelligent_num": decodeIntelligentNum,
		"drift_count":            driftCount,
	}
	utils.ResponseOK(c, "GetApiStatistic Success", resp)
}
