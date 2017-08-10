package atimeago_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/douville/atimeago"
)

func ExampleFrame() {
	ta := atimeago.NewByType(atimeago.Frame)
	fmt.Println(ta.Format(time.Now()))
}

func TestFramePast(t *testing.T) {
	testCases := []struct {
		t    time.Time
		want string
	}{
		{time.Now(), "Same time"},
		{time.Now().Add(-1 * time.Second), "1 second before"},
		{time.Now().Add(-30 * time.Second), "30 seconds before"},
		{time.Now().Add(-1 * time.Minute), "1 minute before"},
		{time.Now().Add(-2 * time.Hour), "2 hours before"},
		{time.Now().Add(-24 * 2 * time.Hour), "2 days before"},
		{time.Now().Add(-24 * 7 * time.Hour), "1 week before"},
		{time.Now().Add(-24 * 60 * time.Hour), "2 months before"},
		{time.Now().Add(-24 * 365 * 2 * time.Hour), "2 years before"},
	}

	ta := atimeago.NewByType(atimeago.Frame)
	for _, testCase := range testCases {
		have := ta.Format(testCase.t)
		if have != testCase.want {
			t.Fatalf("Want: %v ... Have: %v", testCase.want, have)
		}
	}
}

func TestFrameFuture(t *testing.T) {
	testCases := []struct {
		t    time.Time
		want string
	}{
		{time.Now(), "Same time"},
		{time.Now().Add(2 * time.Second), "1 second after"},
		{time.Now().Add(31 * time.Second), "30 seconds after"},
		{time.Now().Add(61 * time.Second), "1 minute after"},
		{time.Now().Add(2 * time.Hour), "1 hour after"},
		{time.Now().Add(24 * 2 * time.Hour), "1 day after"},
		{time.Now().Add(25 * 7 * time.Hour), "1 week after"},
		{time.Now().Add(24 * 61 * time.Hour), "2 months after"},
		{time.Now().Add(24 * 366 * 2 * time.Hour), "2 years after"},
	}

	ta := atimeago.NewByType(atimeago.Frame)
	for _, testCase := range testCases {
		have := ta.Format(testCase.t)
		if have != testCase.want {
			t.Fatalf("Want: %v ... Have: %v", testCase.want, have)
		}
	}
}
