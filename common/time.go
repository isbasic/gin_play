package common

import (
	"fmt"
	"time"
)

func FormatTime(t time.Time) string {
	resTmpl := "%-4d%-2d%-2d"
	y, m, d := t.Date()
	return fmt.Sprintf(resTmpl, y, m, d)
}
