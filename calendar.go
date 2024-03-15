package iranian_tools

import (
	"time"

	"github.com/YasinSaee/iranian_tools/util/calendar"
)

// تغییر تاریخ میلادی به تاریخ جلالی
func ChangeToJalali(time time.Time) calendar.Time {
	return calendar.FormatIranian(time)
}
