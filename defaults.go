package stringify

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Setup defaults functionality for most common types
var (
	maxByteLen        int            = 10000
	defaultJsonParser OptionFunction = func(v any) string {
		j, _ := v.(json.RawMessage)
		if j == nil {
			return "nil"
		}
		return string(j)
	}
	defaultTimeParser  OptionFunction = func(v any) string { return `"` + v.(time.Time).Format(time.RFC3339) + `"` }
	defaultUUIDParser  OptionFunction = func(v any) string { return `"` + v.(uuid.UUID).String() + `"` }
	defaultBytesParser OptionFunction = func(v any) string {
		b, _ := v.([]uint8)
		bl := len(b)
		if bl >= maxByteLen {
			return `"` + fmt.Sprintf("<[]byte len %d>", bl) + `"`
		}

		return `"` + fmt.Sprintf("0x%x", v) + `"`
	}
)

func init() {
	Options = make(map[string]OptionFunction)
	Options["json.RawMessage"] = defaultJsonParser
	Options["time.Time"] = defaultTimeParser
	Options["uuid.UUID"] = defaultUUIDParser
	Options["[]uint8"] = defaultBytesParser
}
