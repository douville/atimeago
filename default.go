package atimeago

import (
	"fmt"
	"time"
)

const (
	defaultParse = time.RFC1123Z
	defaultPrint = time.RFC1123Z

	defaultInstantMsg = "Just now"
	// past
	defaultSecondsPastMsg = "%v second%s ago"
	defaultMinutesPastMsg = "%v minute%s ago"
	defaultHoursPastMsg   = "%v hour%s ago"
	defaultDaysPastMsg    = "%v day%s ago"
	// future
	defaultDayFutureMsg     = "Tomorrow"
	defaultSecondsFutureMsg = "In %d second%s"
	defaultMinutesFutureMsg = "In %d minute%s"
	defaultHoursFutureMsg   = "In %d hour%s"
	defaultDaysFutureMsg    = "In %d day%s"
)

// NewDefault returns the standard Formatter.
func newDefault(opts ...Option) Formatter {
	return NewFormatter(
		newDefaultPast(),
		newDefaultFuture(),
		truncatingCategorizer,
		defaultParse,
		opts...,
	)
}

func newDefaultPast() *Translator {
	formatTime := timePrinter(defaultPrint)
	return &Translator{
		defaultSecondsPast,
		defaultMinutesPast,
		defaultHoursPast,
		defaultDaysPast,
		formatTime,
		formatTime,
		formatTime,
	}
}

func newDefaultFuture() *Translator {
	formatTime := timePrinter(defaultPrint)
	return &Translator{
		defaultSecondsFuture,
		defaultMinutesFuture,
		defaultHoursFuture,
		defaultDaysFuture,
		formatTime,
		formatTime,
		formatTime,
	}
}

func defaultSecondsPast(c int, _ time.Time) string {
	switch c {
	case 0:
		return fmt.Sprint(defaultInstantMsg)
	case 1:
		return fmt.Sprintf(defaultSecondsPastMsg, "A", "")
	}
	return fmt.Sprintf(defaultSecondsPastMsg, c, "s")
}

func defaultMinutesPast(c int, _ time.Time) string {
	if c == 1 {
		return fmt.Sprintf(defaultMinutesPastMsg, "A", "")
	}
	return fmt.Sprintf(defaultMinutesPastMsg, c, "s")
}

func defaultHoursPast(c int, _ time.Time) string {
	if c == 1 {
		return fmt.Sprintf(defaultHoursPastMsg, "An", "")
	}
	return fmt.Sprintf(defaultHoursPastMsg, c, "s")
}

func defaultDaysPast(c int, _ time.Time) string {
	if c == 1 {
		return fmt.Sprintf(defaultDaysPastMsg, "A", "")
	}
	return fmt.Sprintf(defaultDaysPastMsg, c, "s")
}

func defaultSecondsFuture(c int, _ time.Time) string {
	switch c {
	case 0:
		return fmt.Sprint(defaultInstantMsg)
	case 1:
		return fmt.Sprintf(defaultSecondsFutureMsg, c, "")
	}
	return fmt.Sprintf(defaultSecondsFutureMsg, c, "s")
}

func defaultMinutesFuture(c int, _ time.Time) string {
	if c == 1 {
		return fmt.Sprintf(defaultMinutesFutureMsg, c, "")
	}
	return fmt.Sprintf(defaultMinutesFutureMsg, c, "s")
}

func defaultHoursFuture(c int, _ time.Time) string {
	if c == 1 {
		return fmt.Sprintf(defaultHoursFutureMsg, c, "")
	}
	return fmt.Sprintf(defaultHoursFutureMsg, c, "s")
}

func defaultDaysFuture(c int, _ time.Time) string {
	if c == 1 {
		return fmt.Sprintf(defaultDaysFutureMsg, c, "")
	}
	return fmt.Sprintf(defaultDaysFutureMsg, c, "s")
}
