package main

import (
	"fmt"
	"time"

	"github.com/douville/atimeago"
)

func main() {
	var events = []struct{ t time.Time }{
		{time.Now().Add(24 * time.Hour * 366)},
		{time.Now().Add(24 * time.Hour * 31)},
		{time.Now().Add(24 * time.Hour * 21)},
		{time.Now().Add(24 * time.Hour * 4)},
		{time.Now().Add(121 * time.Minute)},
		{time.Now().Add(61 * time.Second)},
		{time.Now()},
		{time.Now().Add(-1 * time.Second)},
		{time.Now().Add(-2 * time.Second)},
		{time.Now().Add(-51 * time.Minute)},
		{time.Now().Add(-45 * time.Hour)},
		{time.Now().Add(-24 * time.Hour * 22)},
		{time.Now().Add(-24 * time.Hour * 30 * 5)},
		{time.Now().Add(-24 * time.Hour * 365)},
	}

	fmt.Println("\nDefault")
	for _, e := range events {
		fmt.Printf("%40s\n", atimeago.Format(e.t))
	}

	fmt.Println("\nShort")
	ta := atimeago.NewByType(atimeago.Short)
	for _, e := range events {
		fmt.Printf("%40s\n", ta.Format(e.t))
	}

	fmt.Println("\nFrame")
	ta = atimeago.NewByType(atimeago.Frame)
	for _, e := range events {
		fmt.Printf("%40s\n", ta.Format(e.t))
	}
}
