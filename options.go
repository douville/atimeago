package atimeago

func SecondsFuture(f FormatFunc) Option {
	return func(ta *timeAgo) *timeAgo {
		ta.future.seconds = f
		return ta
	}
}

func MinutesFuture(f FormatFunc) Option {
	return func(ta *timeAgo) *timeAgo {
		ta.future.minutes = f
		return ta
	}
}

func HoursFuture(f FormatFunc) Option {
	return func(ta *timeAgo) *timeAgo {
		ta.future.hours = f
		return ta
	}
}

func DaysFuture(f FormatFunc) Option {
	return func(ta *timeAgo) *timeAgo {
		ta.future.days = f
		return ta
	}
}

func WeeksFuture(f FormatFunc) Option {
	return func(ta *timeAgo) *timeAgo {
		ta.future.weeks = f
		return ta
	}
}

func MonthsFuture(f FormatFunc) Option {
	return func(ta *timeAgo) *timeAgo {
		ta.future.months = f
		return ta
	}
}

func YearsFuture(f FormatFunc) Option {
	return func(ta *timeAgo) *timeAgo {
		ta.future.years = f
		return ta
	}
}

func SecondsPast(f FormatFunc) Option {
	return func(ta *timeAgo) *timeAgo {
		ta.past.seconds = f
		return ta
	}
}

func MinutesPast(f FormatFunc) Option {
	return func(ta *timeAgo) *timeAgo {
		ta.past.minutes = f
		return ta
	}
}

func HoursPast(f FormatFunc) Option {
	return func(ta *timeAgo) *timeAgo {
		ta.past.hours = f
		return ta
	}
}

func DaysPast(f FormatFunc) Option {
	return func(ta *timeAgo) *timeAgo {
		ta.past.days = f
		return ta
	}
}

func WeeksPast(f FormatFunc) Option {
	return func(ta *timeAgo) *timeAgo {
		ta.past.weeks = f
		return ta
	}
}

func MonthsPast(f FormatFunc) Option {
	return func(ta *timeAgo) *timeAgo {
		ta.past.months = f
		return ta
	}
}

// YearsPast returns an Option that allows modification
// of what will be printed when the time is years before the
// reference time.
func YearsPast(f FormatFunc) Option {
	return func(ta *timeAgo) *timeAgo {
		ta.past.years = f
		return ta
	}
}

// ParseFormat returns an Option that sets the format
// to use when parsing times from strings.
func ParseFormat(format string) Option {
	return func(ta *timeAgo) *timeAgo {
		ta.parseFormat = format
		return ta
	}
}

func Categorizer(c CategorizerFunc) Option {
	return func(ta *timeAgo) *timeAgo {
		ta.categorizer = c
		return ta
	}
}
