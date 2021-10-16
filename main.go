package main

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
	fmt.Println(r)
	return
}

var cap1 = 20.

func main() {

	testItems1 := []Booking{
		{name: "Lolapaluza", value: 289.0, time: 7200., stars: 10.},
		{name: "NotreDame", value: 190.0, time: 2880., stars: 3.},
		{name: "Shaluna", value: 195.0, time: 4320., stars: 4.},
		{name: "LasNoches", value: 300.0, time: 1440., stars: 2.},
		{name: "Bienvenue", value: 130.0, time: 5760., stars: 1.},
		{name: "Cielo", value: 279.0, time: 8640., stars: 1.},
		{name: "Amigos", value: 350.0, time: 10080., stars: 1.},
		{name: "Donatello", value: 110.0, time: 2160., stars: 1.},
	}

	f := func(i, j int) bool {
		return testItems1[i].value/testItems1[i].stars > testItems1[j].value/testItems1[j].stars
	}

	KnapSack(testItems1, cap1, f)
}
