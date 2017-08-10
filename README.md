## [Another](#similar-to) TimeAgo package 

 [![GoDoc](https://godoc.org/github.com/douville/atimeago?status.svg)](https://godoc.org/github.com/douville/atimeago) [![Build Status](https://travis-ci.org/douville/atimeago.svg?branch=master)](https://travis-ci.org/douville/atimeago)

### Install
```
go get github.com/douville/atimeago
```

### Usage
```
package main

import (
	"fmt"
	"time"

	"github.com/douville/atimeago"
)

func main() {
	fmt.Println(atimeago.Format(time.Now().Add(-5 * time.Second)))
}
```
```
>>> go run whats_above.go
5 seconds ago
```

### [Examples](/examples/example.go)
```
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
```
```
>>> go run examples/example.go

Default
         Sat, 11 Aug 2018 11:16:58 -0500
         Sun, 10 Sep 2017 11:16:58 -0500
         Thu, 31 Aug 2017 11:16:58 -0500
                               In 3 days
                              In 2 hours
                             In 1 minute
                                Just now
                            A second ago
                           2 seconds ago
                          51 minutes ago
                               A day ago
         Wed, 19 Jul 2017 11:16:58 -0500
         Mon, 13 Mar 2017 11:16:58 -0500
         Wed, 10 Aug 2016 11:16:58 -0500

Short
                                     -1Y
                                     -1M
                                     -2w
                                     -3d
                                     -2h
                                     -1m
                                      0s
                                      1s
                                      2s
                                     51m
                                      1d
                                      3w
                                      5M
                                      1Y

Frame
                            1 year after
                           1 month after
                           2 weeks after
                            3 days after
                           2 hours after
                          1 minute after
                               Same time
                         1 second before
                        2 seconds before
                       51 minutes before
                            1 day before
                          3 weeks before
                         5 months before
                           1 year before
```

### Similar To
  - [xeonx/timeago](https://github.com/xeonx/timeago)
  - [justincampbell/timeago](https://github.com/justincampbell/timeago)
  - [ararog/timeago](https://github.com/ararog/timeago)
  - [victorjonsson/timeagoo](https://github.com/victorjonsson/timeagoo)

### License
[BSD](/LICENSE)
