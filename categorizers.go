package atimeago

import (
	"time"
)

const (
	dayDuration   = 24 * time.Hour
	weekDuration  = 24 * 7 * time.Hour
	monthDuration = 24 * 30 * time.Hour
	yearDuration  = 24 * 365 * time.Hour
)

const (
	minute = 60
	hour   = 60 * 60
	day    = 60 * 60 * 24
	week   = 60 * 60 * 24 * 7
	month  = 60 * 60 * 24 * 30
	year   = 60 * 60 * 24 * 365
)

func truncatingCategorizer(d time.Duration) (Category, int) {
	if d < 0 {
		d = -d
	}
	switch {
	case d < time.Minute:
		return Second, int(d.Seconds())
	case d < time.Hour:
		return Minute, int(d.Seconds() / minute)
	case d < dayDuration:
		return Hour, int(d.Seconds() / hour)
	case d < weekDuration:
		return Day, int(d.Seconds() / day)
	case d < monthDuration:
		return Week, int(d.Seconds() / week)
	case d < yearDuration:
		return Month, int(d.Seconds() / month)
	}
	return Year, int(d.Seconds() / year)
}

func roundingCategorizer(d time.Duration) (Category, int) {
	if d < 0 {
		d = -d
	}
	s := int(d.Seconds() + 0.5)
	switch {
	case s < minute:
		return Second, s
	case s < hour:
		return Minute, s / minute
	case s < day:
		return Hour, s / hour
	case s < week:
		return Day, s / day
	case s < month:
		return Week, s / week
	case s < year:
		return Month, s / month
	}
	return Year, s / year
}
