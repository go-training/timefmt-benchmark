package strftime_test

import (
	"fmt"
	"time"

	appleboy "github.com/appleboy/timefmt-go"
	fastly "github.com/fastly/go-utils/strftime"
	imperfectgo "github.com/imperfectgo/go-strftime"
	itchyny "github.com/itchyny/timefmt-go"
)

func ExampleFormat() {
	t := time.Date(2008, 9, 3, 20, 4, 26, 654321000, time.FixedZone("CST", 8*3600))

	fmt.Println(t.Format(layoutString))
	fmt.Println(appleboy.Format(t, layout))
	fmt.Println(fastly.Strftime(layout, t))
	fmt.Println(imperfectgo.Format(t, layout))
	fmt.Println(itchyny.Format(t, layout))

	// Output:
	// 2008-09-03
	// 2008-09-03
	// 2008-09-03
	// 2008-09-03
	// 2008-09-03
}
