package helper

import (
	"github.com/greeeds/one-api/relay/channel"
	"github.com/greeeds/one-api/relay/channel/aiproxy"
	"github.com/greeeds/one-api/relay/channel/ali"
	"github.com/greeeds/one-api/relay/channel/anthropic"
	"github.com/greeeds/one-api/relay/channel/baidu"
	"github.com/greeeds/one-api/relay/channel/gemini"
	"github.com/greeeds/one-api/relay/channel/openai"
	"github.com/greeeds/one-api/relay/channel/palm"
	"github.com/greeeds/one-api/relay/channel/tencent"
	"github.com/greeeds/one-api/relay/channel/xunfei"
	"github.com/greeeds/one-api/relay/channel/zhipu"
	"github.com/greeeds/one-api/relay/constant"
)

func GetAdaptor(apiType int) channel.Adaptor {
	switch apiType {
	case constant.APITypeAIProxyLibrary:
		return &aiproxy.Adaptor{}
	case constant.APITypeAli:
		return &ali.Adaptor{}
	case constant.APITypeAnthropic:
		return &anthropic.Adaptor{}
	case constant.APITypeBaidu:
		return &baidu.Adaptor{}
	case constant.APITypeGemini:
		return &gemini.Adaptor{}
	case constant.APITypeOpenAI:
		return &openai.Adaptor{}
	case constant.APITypePaLM:
		return &palm.Adaptor{}
	case constant.APITypeTencent:
		return &tencent.Adaptor{}
	case constant.APITypeXunfei:
		return &xunfei.Adaptor{}
	case constant.APITypeZhipu:
		return &zhipu.Adaptor{}
	}
	return nil
}
