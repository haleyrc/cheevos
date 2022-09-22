package fake

import (
	"fmt"
	"time"
)

func uniqify(s string) string {
	return fmt.Sprintf("%s%d", s, time.Now().UnixNano())
}
