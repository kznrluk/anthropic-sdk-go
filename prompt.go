package anthropic

import (
	"fmt"
	"time"

	"github.com/3JoB/ulib/litefmt"

	"github.com/3JoB/anthropic-sdk-go/data"
)

type CacheItem struct {
	Key        string
	Value      string
	Type       string
	Compress   bool
	ExpireTime time.Time
}

func _Set(human, assistant string) (string, error) {
	if human == "" {
		return "", data.ErrPromptHumanEmpty
	}
	if assistant == "" {
		return litefmt.Sprint("\n\nHuman: ", human, "\n\nAssistant:"), nil
		// return fmt.Sprintf("\n\nHuman: %v\n\nAssistant:", human), nil
	}
	return litefmt.Sprint(human, assistant), nil
	// return fmt.Sprintf("%v%v", human, assistant), nil
}

// Commit: The loop overhead here is too large, caching may need to be set up
func (c *Context) build(module any) (string, error) {
	if c.cache {
	}

	switch r := module.(type) {
	case data.MessageModule:
		return _Set(r.Human, r.Assistant)
	case []data.MessageModule:
		var prompts string
		for _, d := range r {
			if d.Human == "" {
				return "", data.ErrPromptHumanEmpty
			}
			if d.Assistant == "" {
				return litefmt.Sprint(prompts, "\n\nHuman: ", d.Human, "\n\nAssistant:"), nil
				// return fmt.Sprintf("%v\n\nHuman: %v\n\nAssistant:", prompts, d.Human), nil
			}
			prompts = litefmt.Sprint(prompts, "\n\nHuman: ", d.Human, "\n\nAssistant:", d.Assistant)
			// prompts = fmt.Sprintf("%v\n\nHuman: %v\n\nAssistant:%v", prompts, d.Human, d.Assistant)
		}
		return prompts, nil
	default:
		return "", fmt.Errorf("unknown module type: %T", module)
	}
}

/*func Add(context, human string) (string, error) {
	if human == "" {
		return "", data.ErrPromptHumanEmpty
	}
	if context == "" {
		return "", data.ErrPromptCtxEmpty
	}
	return fmt.Sprintf("%v\n\nHuman: %v\n\nAssistant:", context, human), nil
}
*/
