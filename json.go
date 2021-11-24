//go:build !jsoniter
// +build !jsoniter

package gopkg

import (
	"encoding/json"
)

var (
	Marshal       = json.Marshal
	Unmarshal     = json.Unmarshal
	MarshalIndent = json.MarshalIndent
	NewDecoder    = json.NewDecoder
	NewEncoder    = json.NewEncoder
)
