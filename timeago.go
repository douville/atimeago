/*
Package atimeago is for formatting times as custom strings.

Sample output:

	Just now
	2 seconds ago
	51 minutes ago
	A day ago
	Wed, 12 Jul 2017 08:46:33 -0700
*/
package atimeago

import (
	"fmt"
	"time"
)

// TranslatorType specifies built-in translators.
type TranslatorType int

const (
	// Default
	// 2 seconds ago, In 2 seconds
	Default TranslatorType = iota

	// Short
	// 2s, -2s
	Short

	// Frame
	// 2 seconds before, 2 seconds after
	Frame
)

// Category is Second, Minute, ..., Year
type Category int

const (
	Second Category = iota
	Minute
	Hour
	Day
	Week
	Month
	Year
)

// Formatter provides formatting functions for time objects or strings.
type Formatter interface {
	Format(time.Time) string
	FormatWithRef(time.Time, time.Time) string
	Parse(string) (string, error)
	ParseWithRef(string, string) (string, error)
}

// FormatFunc is function that returns a formatted time string.
type FormatFunc func(int, time.Time) string

// Option allows functional options.
type Option func(*timeAgo) *timeAgo

// CategorizerFunc is a func that takes a duration and tells how many of a
// Category (Second, Minute, ..., Year) it is away from the configured reference time.
type CategorizerFunc func(d time.Duration) (Category, int)

// timeAgo is the core struct.
type timeAgo struct {
	past        *Translator
	future      *Translator
	categorizer CategorizerFunc
	parseFormat string
}

// Format is a helper for printing of the default format.
func Format(t time.Time) string {
	return New().Format(t)
}

// New returns a formatter with the default translator.
func New(opts ...Option) Formatter {
	return NewByType(Default, opts...)
}

// NewByType returns a new of formatter of the given type.
func NewByType(t TranslatorType, opts ...Option) Formatter {
	switch t {
	case Default:
		return newDefault(opts...)
	case Short:
		return newShort(opts...)
	case Frame:
		return newFrame(opts...)
	}
	panic("bad translator type")
}

// NewFormatter allows creation of any type of Formatter.
func NewFormatter(
	past,
	future *Translator,
	categorizer CategorizerFunc,
	parseFormat string,
	opts ...Option,
) Formatter {
	ta := &timeAgo{
		past:        past,
		future:      future,
		categorizer: categorizer,
		parseFormat: parseFormat,
	}
	for _, opt := range opts {
		ta = opt(ta)
	}
	return ta
}

// Format returns a string for the given time using the current time as a
// frame of reference.
func (ta *timeAgo) Format(t time.Time) string {
	return ta.FormatWithRef(t, time.Now())
}

// FormatWithRef allows a time reference to be given.
func (ta *timeAgo) FormatWithRef(t time.Time, ref time.Time) string {
	d := ref.Sub(t)
	cat, count := ta.categorizer(d)
	if d < 0 {
		return ta.future.translate(cat, count, t)
	} else {
		return ta.past.translate(cat, count, t)
	}
}

// Translator is used to format a past or future time as a string.
type Translator struct {
	seconds FormatFunc
	minutes FormatFunc
	hours   FormatFunc
	days    FormatFunc
	weeks   FormatFunc
	months  FormatFunc
	years   FormatFunc
}

// translate takes a Translator with the duration between time t and the reference time,
// and returns the formatted string. Time t is usually not used.
func (tr *Translator) translate(cat Category, count int, t time.Time) string {
	switch cat {
	case Second:
		return tr.seconds(count, t)
	case Minute:
		return tr.minutes(count, t)
	case Hour:
		return tr.hours(count, t)
	case Day:
		return tr.days(count, t)
	case Week:
		return tr.weeks(count, t)
	case Month:
		return tr.months(count, t)
	case Year:
		return tr.years(count, t)
	}
	panic(fmt.Sprintf("bad category: %d", cat))
}

// Parse parses a time from the given string using the configure parse time format.
func (ta *timeAgo) Parse(s string) (string, error) {
	t, err := time.Parse(ta.parseFormat, s)
	if err != nil {
		return "", err
	}
	return ta.FormatWithRef(t, time.Now()), nil
}

// ParseWithRef allows parsing with a reference time (which is usually time.Now()).
func (ta *timeAgo) ParseWithRef(s string, ref string) (string, error) {
	t, err := time.Parse(ta.parseFormat, s)
	if err != nil {
		return "", err
	}
	rt, err := time.Parse(ta.parseFormat, ref)
	if err != nil {
		return "", err
	}
	return ta.FormatWithRef(t, rt), nil
}

func timePrinter(fmt string) FormatFunc {
	return func(_ int, t time.Time) string {
		return t.Format(fmt)
	}
}
