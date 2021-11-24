package gopkg

import (
	"strings"
)

// StrJoin String concatenation.
func StrJoin(variables ...string) string {
	switch len(variables) {
	case 0:
		return ""
	case 1:
		return variables[0]
	default:
		var (
			n int
			b strings.Builder
		)
		for i := 0; i < len(variables); i++ {
			n += len(variables[i])
		}
		b.Grow(n)
		for i := range variables {
			b.WriteString(variables[i])
		}
		return b.String()
	}
}
