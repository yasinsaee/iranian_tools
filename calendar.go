package iranian_tools

import (
	"time"

	"github.com/YasinSaee/iranian_tools/util/calendar"
)

func ChangeToJalali(time time.Time) calendar.Time {
	return calendar.FormatIranian(time)
}