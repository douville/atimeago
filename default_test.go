package atimeago_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/douville/atimeago"
)

func ExampleDefault() {
	fmt.Println(atimeago.Format(time.Now()))
}

func TestDefaultPast(t *testing.T) {
	testCases := []struct {
		t    time.Time
		want string
	}{
		{time.Now(), "Just now"},
		{time.Now().Add(-1 * time.Second), "A second ago"},
		{time.Now().Add(-30 * time.Second), "30 seconds ago"},
		{time.Now().Add(-1 * time.Minute), "A minute ago"},
		{time.Now().Add(-2 * time.Hour), "2 hours ago"},
		{time.Now().Add(-24 * 2 * time.Hour), "2 days ago"},
	}

	ta := atimeago.New()
	for _, testCase := range testCases {
		have := ta.Format(testCase.t)
		if have != testCase.want {
			t.Fatalf("Want: %s ... Have: %s", testCase.want, have)
		}
	}
}

func TestDefaultFuture(t *testing.T) {
	testCases := []struct {
		t    time.Time
		want string
	}{
		{time.Now(), "Just now"},
		{time.Now().Add(2 * time.Second), "In 1 second"},
		{time.Now().Add(31 * time.Second), "In 30 seconds"},
		{time.Now().Add(2 * time.Minute), "In 1 minute"},
		{time.Now().Add(3 * time.Hour), "In 2 hours"},
		{time.Now().Add(24 * 3 * time.Hour), "In 2 days"},
	}

	ta := atimeago.New()
	for _, testCase := range testCases {
		have := ta.Format(testCase.t)
		if have != testCase.want {
			t.Fatalf("Want: %s ... Have: %s", testCase.want, have)
		}
	}
}

func TestDefaultMax(t *testing.T) {
	testCases := []struct {
		t time.Time
	}{
		{time.Now().Add(-24 * 365 * 2 * time.Hour)},
		{time.Now().Add(-24 * 60 * time.Hour)},
		{time.Now().Add(-24 * 7 * time.Hour)},
		{time.Now().Add(-24 * 7 * time.Hour)},
		{time.Now().Add(-24 * 60 * time.Hour)},
		{time.Now().Add(-24 * 365 * 2 * time.Hour)},
	}
	ta := atimeago.New()
	for _, testCase := range testCases {
		have := ta.Format(testCase.t)
		_, err := time.Parse(time.RFC1123Z, have)
		if err != nil {
			t.Fatal(err)
		}
	}
}
