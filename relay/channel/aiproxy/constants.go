package aiproxy

import "github.com/greeeds/one-api/relay/channel/openai"

var ModelList = []string{""}

func init() {
	ModelList = openai.ModelList
}
