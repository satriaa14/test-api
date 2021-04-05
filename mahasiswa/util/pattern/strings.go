package pattern

import (
	"fmt"
	"strings"
)

func DataURL(path string) []string {
	return strings.Split(fmt.Sprintf("%s", path), "/")
}
