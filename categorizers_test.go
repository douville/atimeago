package atimeago

import (
	"errors"
	"testing"
	"time"
)

func TestTruncatingCategorizer(t *testing.T) {
	testCases := []struct {
		d            time.Duration
		wantCategory Category
		wantCount    int
	}{
		{time.Second, Second, 1},
		{time.Second * 59, Second, 59},
		{time.Minute - time.Nanosecond, Second, 59},
		{time.Minute, Minute, 1},
		{time.Hour, Hour, 1},
		{time.Hour * 24, Day, 1},
		{time.Hour * 24 * 7, Week, 1},
		{time.Hour * 24 * 30, Month, 1},
		{time.Hour * 24 * 365, Year, 1},
	}

	for _, testCase := range testCases {
		cat, count := truncatingCategorizer(testCase.d)
		if cat != testCase.wantCategory {
			t.Fatalf("(bad category) Want: %v ... Have: %v", testCase.wantCategory, cat)
		}
		if count != testCase.wantCount {
			t.Fatalf("(bad count) Want: %v ... Have: %v", testCase.wantCount, count)
		}
	}
}

func TestRoundingCategorizer(t *testing.T) {
	testCases := []struct {
		d            time.Duration
		wantCategory Category
		wantCount    int
	}{
		{time.Second, Second, 1},
		{time.Second * 59, Second, 59},
		{time.Minute - time.Nanosecond, Minute, 1},
		{time.Minute, Minute, 1},
		{time.Minute, Minute, 1},
		{time.Hour, Hour, 1},
		{time.Hour * 24, Day, 1},
		{time.Hour * 24 * 7, Week, 1},
		{time.Hour * 24 * 30, Month, 1},
		{time.Hour * 24 * 365, Year, 1},
	}

	for _, testCase := range testCases {
		cat, count := roundingCategorizer(testCase.d)
		if cat != testCase.wantCategory {
			t.Fatalf("(bad category) Want: %v ... Have: %v (duration=%v)", testCase.wantCategory, cat, testCase.d)
		}
		if count != testCase.wantCount {
			t.Fatalf("(bad count) Want: %v ... Have: %v", testCase.wantCount, count)
		}
	}
}

func TestBadCategory(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			errors.New("panic not detected")
		}
	}()
	c := Categorizer(func(_ time.Duration) (Category, int) {
		return -1, 0
	})
	ta := New(c)
	_ = ta.Format(time.Now())
}
