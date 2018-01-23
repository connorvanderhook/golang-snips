package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type timeSlice []time.Time

func (s timeSlice) Less(i, j int) bool { return s[i].Before(s[j]) }
func (s timeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s timeSlice) Len() int           { return len(s) }

func DatesInPast(i int) time.Time {
	x := -rand.Intn(8)
	y := -rand.Intn(60)
	return time.Now().AddDate(0, 0, i).Add(time.Duration(x) * time.Hour).Add(time.Duration(y) * time.Minute)
	// .Format(time.RFC1123Z)
}

func SundaysInPast() timeSlice {
	sundays := []time.Time{}
	for i := 90; i < 360; i++ {
		if i%3 == 0 && !(i%4 == 0) && !(i%9 == 0) {
			sundays = append(sundays, DatesInPast(-i))
		}
		if !(i%2 == 0) && !(i%5 == 0) {
			sundays = append(sundays, DatesInPast(-i))
		}
	}
	return sundays
}

// Sun 11 Nov 2019 19:20:25 PST

func main() {
	sundays := SundaysInPast()
	sort.Sort(sundays)
	for _, sunday := range sundays {
		fmt.Println(sunday.Format(time.RFC1123Z))
	}
}
