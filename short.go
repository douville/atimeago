package atimeago

import (
	"fmt"
	"time"
)

const (
	shortParse = time.RFC1123Z

	shortSecondsMsg = "%ds"
	shortMinutesMsg = "%dm"
	shortHoursMsg   = "%dh"
	shortDaysMsg    = "%dd"
	shortWeeksMsg   = "%dw"
	shortMonthsMsg  = "%dM"
	shortYearsMsg   = "%dY"
)

func newShort(opts ...Option) Formatter {
	return NewFormatter(
		newShortPast(),
		newShortFuture(),
		truncatingCategorizer,
		shortParse,
		opts...,
	)
}

func newShortPast() *Translator {
	return &Translator{
		shortSecondsPast,
		shortMinutesPast,
		shortHoursPast,
		shortDaysPast,
		shortWeeksPast,
		shortMonthsPast,
		shortYearsPast,
	}
}

func newShortFuture() *Translator {
	return &Translator{
		shortSecondsFuture,
		shortMinutesFuture,
		shortHoursFuture,
		shortDaysFuture,
		shortWeeksFuture,
		shortMonthsFuture,
		shortYearsFuture,
	}
}

func shortSecondsPast(c int, _ time.Time) string {
	return fmt.Sprintf(shortSecondsMsg, c)
}

func shortMinutesPast(c int, _ time.Time) string {
	return fmt.Sprintf(shortMinutesMsg, c)
}

func shortHoursPast(c int, _ time.Time) string {
	return fmt.Sprintf(shortHoursMsg, c)
}

func shortDaysPast(c int, _ time.Time) string {
	return fmt.Sprintf(shortDaysMsg, c)
}

func shortWeeksPast(c int, _ time.Time) string {
	return fmt.Sprintf(shortWeeksMsg, c)
}

func shortMonthsPast(c int, _ time.Time) string {
	return fmt.Sprintf(shortMonthsMsg, c)
}

func shortYearsPast(c int, _ time.Time) string {
	return fmt.Sprintf(shortYearsMsg, c)
}

func shortSecondsFuture(c int, _ time.Time) string {
	return fmt.Sprintf(shortSecondsMsg, -c)
}

func shortMinutesFuture(c int, _ time.Time) string {
	return fmt.Sprintf(shortMinutesMsg, -c)
}

func shortHoursFuture(c int, _ time.Time) string {
	return fmt.Sprintf(shortHoursMsg, -c)
}

func shortDaysFuture(c int, _ time.Time) string {
	return fmt.Sprintf(shortDaysMsg, -c)
}

func shortWeeksFuture(c int, t time.Time) string {
	return fmt.Sprintf(shortWeeksMsg, -c)
}

func shortMonthsFuture(c int, t time.Time) string {
	return fmt.Sprintf(shortMonthsMsg, -c)
}

func shortYearsFuture(c int, _ time.Time) string {
	return fmt.Sprintf(shortYearsMsg, -c)
}
