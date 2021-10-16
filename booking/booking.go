package booking

import (
	"fmt"
	"sort"
)

type Booking struct {
	name  string
	value float64
	time  float64
	stars float64
}

func KnapSack(bookings []Booking, maxStars float64, metric func(i, j int) bool) (r []Booking) {
	sort.Slice(bookings, metric)

	s := 0.

	for _, i := range bookings {
		if s+i.stars <= maxStars {
			r = append(r, i)
			s += i.stars
		}
	}
	fmt.Println(s)
	return
}
