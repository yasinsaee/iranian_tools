package calendar

import (
	"fmt"
	"time"
)

// A Month specifies a month of the year starting from Farvardin = 1.
type Month int

// A Weekday specifies a day of the week starting from Shanbe = 0.
type Weekday int

// List of months in Iranian calendar.
const (
	Farvardin Month = 1 + iota
	Ordibehesht
	Khordad
	Tir
	Mordad
	Shahrivar
	Mehr
	Aban
	Azar
	Dey
	Bahman
	Esfand
)

// List of days in a week.
const (
	Shanbeh Weekday = iota
	Yekshanbeh
	Doshanbeh
	Seshanbeh
	Charshanbeh
	Panjshanbeh
	Jomeh
)

var months = [12]string{
	"فروردین",
	"اردیبهشت",
	"خرداد",
	"تیر",
	"مرداد",
	"شهریور",
	"مهر",
	"آبان",
	"آذر",
	"دی",
	"بهمن",
	"اسفند",
}

var days = [7]string{
	"شنبه",
	"یک\u200cشنبه",
	"دوشنبه",
	"سه\u200cشنبه",
	"چهارشنبه",
	"پنج\u200cشنبه",
	"جمعه",
}

type Time struct {
	Year    int
	Month   Month
	Day     int
	Hour    int
	Min     int
	Sec     int
	Weekday Weekday
}

// برگرداندن تاریخ به صورت متنی
func (m Month) String() string {
	switch {
	case m < 1:
		return months[0]
	case m > 11:
		return months[11]
	default:
		return months[m-1]
	}
}

// برگرداندن روز به صورت متنی
func (d Weekday) String() string {
	switch {
	case d < 0:
		return days[0]
	case d > 6:
		return days[6]
	default:
		return days[d]
	}
}

// چاپ کردن تاریخ به صورت عدد
func (t Time) Print() {
	fmt.Printf("%d/%02d/%02d", t.Year, t.Month, t.Day)
}

// چاپ کردن تاریخ به صورت متنی
func (t Time) PrintText() {
	fmt.Printf("%d/%s/%s", t.Year, t.Month.String(), t.Weekday.String())
}

// تغییر تاریخ میلادی به تاریخ جلالی
// یا همان تاریخ ایرانیان
func FormatIranian(ti time.Time) (t Time) {
	gDM := [13]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}

	year := ti.Year()
	month := ti.Month()
	day := ti.Day()

	if year > 1600 {
		t.Year = 979
		year -= 1600
	} else {
		t.Year = 0
		year -= 621
	}

	year2 := year
	if month > 2 {
		year2++
	}

	days := (365 * year) + (year2+3)/4 - (year2+99)/100 + (year2+399)/400 - 80 + day + gDM[month-1]
	t.Year += 33 * (days / 12053)
	days %= 12053
	t.Year += 4 * (days / 1461)
	days %= 1461

	if days > 365 {
		t.Year += (days - 1) / 365
		days = (days - 1) % 365
	}

	if days < 186 {
		t.Month = Month(1 + days/31)
	} else {
		t.Month = Month(7 + (days-186)/30)
	}

	if days < 186 {
		t.Day = 1 + days%31
	} else {
		t.Day = 1 + (days-186)%30
	}

	t = Time{
		Year:    t.Year,
		Month:   t.Month,
		Day:     t.Day,
		Hour:    ti.Hour(),
		Min:     ti.Minute(),
		Sec:     ti.Second(),
		Weekday: GetWeekday(ti.Weekday()),
	}

	return
}

// Iran returns a pointer to time.Location of Asia/Tehran
func Iran() *time.Location {
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		loc = time.FixedZone("Asia/Tehran", 12600) // UTC + 03:30
	}
	return loc
}

func GetWeekday(wd time.Weekday) Weekday {
	switch wd {
	case time.Saturday:
		return Shanbeh
	case time.Sunday:
		return Yekshanbeh
	case time.Monday:
		return Doshanbeh
	case time.Tuesday:
		return Seshanbeh
	case time.Wednesday:
		return Charshanbeh
	case time.Thursday:
		return Panjshanbeh
	case time.Friday:
		return Jomeh
	}
	return 0
}

// نمایش اختلاف بین دو تاریخ به صورت خلاصه وار و متنی
func DifferenceDate(inputDatetime time.Time) string {
	systemDatetime := time.Now()

	differenceInDays := int(systemDatetime.Sub(inputDatetime).Hours() / 24)
	if differenceInDays > 7 {
		if differenceInDays > 7 && differenceInDays < 14 {
			return "هفته پیش"
		} else if differenceInDays < 7*4 {
			return "چند هفته پیش"
		} else if differenceInDays >= 30 && differenceInDays < 60 {
			return "ماه پیش"
		} else if differenceInDays < 365 {
			return "چند ماه پیش"
		} else if differenceInDays >= 365 && differenceInDays < 730 {
			return "سال پیش"
		} else {
			return "چند سال پیش"
		}
	} else {
		differenceInSeconds := int(systemDatetime.Sub(inputDatetime).Seconds())
		if differenceInSeconds < 60 {
			return "لحظاتی پیش"
		} else if differenceInSeconds < 60*60 {
			return "دقایقی پیش"
		} else if differenceInSeconds < 3600*24 {
			return "ساعاتی پیش"
		} else if differenceInDays == 1 {
			return "دیروز"
		} else {
			return "چند روز پیش"
		}
	}
}
