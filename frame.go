package atimeago

import (
	"fmt"
	"time"
)

const (
	frameParse = time.RFC1123Z

	frameInstantMsg = "Same time"

	frameSecondsPastMsg = "%v second%s before"
	frameMinutesPastMsg = "%v minute%s before"
	frameHoursPastMsg   = "%v hour%s before"
	frameDaysPastMsg    = "%v day%s before"
	frameWeeksPastMsg   = "%v week%s before"
	frameMonthsPastMsg  = "%v month%s before"
	frameYearsPastMsg   = "%v year%s before"

	frameSecondsFutureMsg = "%v second%s after"
	frameMinutesFutureMsg = "%v minute%s after"
	frameHoursFutureMsg   = "%v hour%s after"
	frameDaysFutureMsg    = "%v day%s after"
	frameWeeksFutureMsg   = "%v week%s after"
	frameMonthsFutureMsg  = "%v month%s after"
	frameYearsFutureMsg   = "%v year%s after"
)

func newFrame(opts ...Option) Formatter {
	return NewFormatter(
		newFramePast(),
		newFrameFuture(),
		truncatingCategorizer,
		time.RFC1123Z,
		opts...,
	)
}

func newFramePast() *Translator {
	return &Translator{
		frameSecondsPast,
		frameMinutesPast,
		frameHoursPast,
		frameDaysPast,
		frameWeeksPast,
		frameMonthsPast,
		frameYearsPast,
	}
}

func newFrameFuture() *Translator {
	return &Translator{
		frameSecondsFuture,
		frameMinutesFuture,
		frameHoursFuture,
		frameDaysFuture,
		frameWeeksFuture,
		frameMonthsFuture,
		frameYearsFuture,
	}
}

func frameSecondsPast(c int, _ time.Time) string {
	switch c {
	case 0:
		return fmt.Sprint(frameInstantMsg)
	case 1:
		return fmt.Sprintf(frameSecondsPastMsg, c, "")
	}
	return fmt.Sprintf(frameSecondsPastMsg, c, "s")
}

func frameMinutesPast(c int, _ time.Time) string {
	if c == 1 {
		return fmt.Sprintf(frameMinutesPastMsg, c, "")
	}
	return fmt.Sprintf(frameMinutesPastMsg, c, "s")
}

func frameHoursPast(c int, _ time.Time) string {
	if c == 1 {
		return fmt.Sprintf(frameHoursPastMsg, c, "")
	}
	return fmt.Sprintf(frameHoursPastMsg, c, "s")
}

func frameDaysPast(c int, _ time.Time) string {
	if c == 1 {
		return fmt.Sprintf(frameDaysPastMsg, c, "")
	}
	return fmt.Sprintf(frameDaysPastMsg, c, "s")
}

func frameWeeksPast(c int, _ time.Time) string {
	if c == 1 {
		return fmt.Sprintf(frameWeeksPastMsg, c, "")
	}
	return fmt.Sprintf(frameWeeksPastMsg, c, "s")
}

func frameMonthsPast(c int, _ time.Time) string {
	if c == 1 {
		return fmt.Sprintf(frameMonthsPastMsg, c, "")
	}
	return fmt.Sprintf(frameMonthsPastMsg, c, "s")
}

func frameYearsPast(c int, _ time.Time) string {
	if c == 1 {
		return fmt.Sprintf(frameYearsPastMsg, c, "")
	}
	return fmt.Sprintf(frameYearsPastMsg, c, "s")
}

func frameSecondsFuture(c int, _ time.Time) string {
	if c == 1 {
		return fmt.Sprintf(frameSecondsFutureMsg, c, "")
	}
	return fmt.Sprintf(frameSecondsFutureMsg, c, "s")
}

func frameMinutesFuture(c int, _ time.Time) string {
	if c == 1 {
		return fmt.Sprintf(frameMinutesFutureMsg, c, "")
	}
	return fmt.Sprintf(frameMinutesFutureMsg, c, "s")
}

func frameHoursFuture(c int, _ time.Time) string {
	if c == 1 {
		return fmt.Sprintf(frameHoursFutureMsg, c, "")
	}
	return fmt.Sprintf(frameHoursFutureMsg, c, "s")
}

func frameDaysFuture(c int, _ time.Time) string {
	if c == 1 {
		return fmt.Sprintf(frameDaysFutureMsg, c, "")
	}
	return fmt.Sprintf(frameDaysFutureMsg, c, "s")
}

func frameWeeksFuture(c int, t time.Time) string {
	if c == 1 {
		return fmt.Sprintf(frameWeeksFutureMsg, c, "")
	}
	return fmt.Sprintf(frameWeeksFutureMsg, c, "s")
}

func frameMonthsFuture(c int, t time.Time) string {
	if c == 1 {
		return fmt.Sprintf(frameMonthsFutureMsg, c, "")
	}
	return fmt.Sprintf(frameMonthsFutureMsg, c, "s")
}

func frameYearsFuture(c int, _ time.Time) string {
	if c == 1 {
		return fmt.Sprintf(frameYearsFutureMsg, c, "")
	}
	return fmt.Sprintf(frameYearsFutureMsg, c, "s")
}
