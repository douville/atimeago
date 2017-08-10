package atimeago_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/douville/atimeago"
)

func TestFormat(t *testing.T) {
	have := atimeago.Format(time.Now().Add(-2 * time.Hour))
	want := "2 hours ago"
	if have != want {
		t.Fatalf("Want: %v ... Have: %v", want, have)
	}
}

func ExampleFormat() {
	fmt.Println(atimeago.Format(time.Now()))
}
