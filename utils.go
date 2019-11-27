package recordjar

import (
	"strings"
)

func isComment(s string) bool {
	return strings.HasPrefix(s, "%%")
}
