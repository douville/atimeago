package atimeago_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/douville/atimeago"
)

func ExampleShort() {
	ta := atimeago.NewByType(atimeago.Short)
	fmt.Println(ta.Format(time.Now()))
}

func TestPastShort(t *testing.T) {
	testCases := []struct {
		t    time.Time
		want string
	}{
		{time.Now(), "0s"},
		{time.Now().Add(-1 * time.Second), "1s"},
		{time.Now().Add(-30 * time.Second), "30s"},
		{time.Now().Add(-1 * time.Minute), "1m"},
		{time.Now().Add(-2 * time.Hour), "2h"},
		{time.Now().Add(-24 * 2 * time.Hour), "2d"},
		{time.Now().Add(-24 * 7 * time.Hour), "1w"},
		{time.Now().Add(-24 * 60 * time.Hour), "2M"},
	}

	ta := atimeago.NewByType(atimeago.Short)
	for _, testCase := range testCases {
		have := ta.Format(testCase.t)
		if have != testCase.want {
			t.Fatalf("Want: %v ... Have: %v", testCase.want, have)
		}
	}
}

func TestFutureShort(t *testing.T) {
	testCases := []struct {
		t    time.Time
		want string
	}{
		{time.Now(), "0s"},
		{time.Now().Add(1 * time.Second), "0s"},
		{time.Now().Add(2 * time.Second), "-1s"},
		{time.Now().Add(30 * time.Second), "-29s"},
		{time.Now().Add(61 * time.Second), "-1m"},
		{time.Now().Add(2 * time.Hour), "-1h"},
		{time.Now().Add(24 * 2 * time.Hour), "-1d"},
		{time.Now().Add(25 * 7 * time.Hour), "-1w"},
		{time.Now().Add(24 * 61 * time.Hour), "-2M"},
		{time.Now().Add(24 * 366 * time.Hour), "-1Y"},
	}

	ta := atimeago.NewByType(atimeago.Short)

	for _, testCase := range testCases {
		have := ta.Format(testCase.t)
		if have != testCase.want {
			t.Fatalf("Want: %v ... Have: %v", testCase.want, have)
		}
	}

}
