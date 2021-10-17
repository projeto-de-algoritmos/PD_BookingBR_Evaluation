package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Booking struct {
	name  string
	value float64
	time  float64
	stars float64
}

func main() {

	maxStars := readUser()

	bookValues := []Booking{
		{name: "Lolapaluza", value: 289.0, time: 7200., stars: 10.},
		{name: "NotreDame", value: 190.0, time: 2880., stars: 3.},
		{name: "Shaluna", value: 195.0, time: 4320., stars: 4.},
		{name: "LasNoches", value: 300.0, time: 1440., stars: 2.},
		{name: "Bienvenue", value: 130.0, time: 5760., stars: 1.},
		{name: "Cielo", value: 279.0, time: 8640., stars: 1.},
		{name: "Amigos", value: 350.0, time: 10080., stars: 5.},
		{name: "Donatello", value: 110.0, time: 2160., stars: 1.},

		{name: "SecondHand Hotel", value: 689.0, time: 17200., stars: 7.},
		{name: "Spitfire Hotel", value: 490.0, time: 6580., stars: 8.},
		{name: "Wicked Hotel", value: 905.0, time: 1320., stars: 6.},
		{name: "Hotel Escolhas", value: 515.0, time: 7840., stars: 3.},
		{name: "Foundry Hotel", value: 130.0, time: 5760., stars: 9.},
		{name: "Quantum Hotel", value: 279.0, time: 8640., stars: 1.},
		{name: "Oráculo Hotel", value: 350.0, time: 10080., stars: 5.},
		{name: "Global Hotel", value: 1310.0, time: 12160., stars: 7.},
		{name: "Glorial Hotel", value: 1410.0, time: 12160., stars: 7.},
	}

	f := func(i, j int) bool {
		return bookValues[i].value > bookValues[j].value
	}

	KnapSack(bookValues, maxStars, f)

	v, s := BestCombination(bookValues, maxStars)

	fmt.Print(`
	O valor total das acomodações foi de: 
	`, v)

	fmt.Println(`
	O Hoteis escolhidos foram os seguintes:
	`)

	for _, v := range s {
		fmt.Println(`		`, v.name, v.value)
	}
}

func KnapSack(bookings []Booking, maxStars float64, metric func(i, j int) bool) (r []Booking, r2 []Booking) {
	sort.Slice(bookings, metric)

	s := 0.

	for _, i := range bookings {
		if s+i.stars <= maxStars {
			r = append(r, i)
			s += i.stars
		}
	}
	//fmt.Println(r)
	return
}

func PossibleCombinations(bookings []Booking, ch chan []Booking) {
	defer close(ch)

	p := int(math.Pow(2., float64(len(bookings))))

	for i := 0; i < p; i++ {
		set := []Booking{}
		for j := 0; j < len(bookings); j++ {
			if (i>>uint(j))&1 == 1 {
				set = append(set, bookings[j])
			}
		}
		ch <- set
	}
}

func getSackStars(set []Booking) (r float64) {
	for _, i := range set {
		r += i.stars
	}
	return
}

func getSackValue(set []Booking) (r float64) {
	for _, i := range set {
		r += i.value
	}
	return
}

func BestCombination(bookings []Booking, maxStars float64) (float64, []Booking) {
	bestVal := 0.
	bestSack := []Booking{}

	ch := make(chan []Booking)
	go PossibleCombinations(bookings, ch)

	for sack := range ch {
		if getSackStars(sack) <= maxStars {
			v := getSackValue(sack)
			if v > bestVal {
				bestVal = v
				bestSack = sack
			}
		}
	}
	return bestVal, bestSack
}

func readUser() float64 {
	fmt.Print(`
							BookingBR.com
			
	Para selecionarmos as melhores acomodações para você, precisamos que você selecione o número de estrelas de 1 a 10: 
	`)

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	star := strings.TrimRight(strings.Replace(text, "\n", "", -1), "\r")

	stars, err := strconv.ParseFloat(star, 32)

	if err != nil {
		fmt.Println(err)
	}

	return stars
}
