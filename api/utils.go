package api

import (
	"encoding/json"
	"fmt"
)

// @TODO
func toJSON(v any) string {
	j, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("failed to toJSON: %#v", v)
	}

	return string(j)
}
