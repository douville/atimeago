package atimeago_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/douville/atimeago"
)

func TestSecondsFuture(t *testing.T) {
	var have, want string
	secondsFuture := atimeago.SecondsFuture(func(c int, _ time.Time) string { return "custom" })
	ta := atimeago.New(secondsFuture)
	have = ta.Format(time.Now().Add(10 * time.Second))
	want = "custom"
	if have != want {
		t.Fatalf("Want: %v -- Have: %v", want, have)
	}
}

func ExampleSecondsFuture() {
	secondsFuture := atimeago.SecondsFuture(func(c int, _ time.Time) string { return fmt.Sprintf("in %d seconds...", c) })
	ta := atimeago.New(secondsFuture)
	fmt.Println(ta.Format(time.Now().Add(10 * time.Second)))
}

func TestSecondsPast(t *testing.T) {
	var have, want string
	secondsPast := atimeago.SecondsPast(func(c int, _ time.Time) string { return "custom" })
	ta := atimeago.New(secondsPast)
	have = ta.Format(time.Now().Add(-10 * time.Second))
	want = "custom"
	if have != want {
		t.Fatalf("Want: %v -- Have: %v", want, have)
	}
}

func ExampleSecondsPast() {
	secondsPast := atimeago.SecondsPast(func(c int, t time.Time) string {
		switch c {
		case 0:
			return fmt.Sprint("Less than a second ago")
		case 1:
			return fmt.Sprint("One second ago")
		}
		return fmt.Sprintf("%d secs", c)

	})
	ta := atimeago.New(secondsPast)
	fmt.Println(ta.Format(time.Now().Add(-1 * time.Second)))
}

func ExampleMaxPast() {
	weeksPast := atimeago.WeeksPast(func(c int, t time.Time) string { return t.Format(time.RFC850) })
	monthsPast := atimeago.MonthsPast(func(c int, t time.Time) string { return t.Format(time.RFC850) })
	yearsPast := atimeago.YearsPast(func(c int, t time.Time) string { return t.Format(time.RFC850) })
	ta := atimeago.New(weeksPast, monthsPast, yearsPast)
	fmt.Println(ta.Format(time.Now().Add(-1 * 24 * 364 * time.Hour)))
}
