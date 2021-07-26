package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

type Taxi struct {
	a     int
	b     int
	c     int
	d     int
	money int
}

func main() {
	var N, X int
	var Tarr []Taxi
	fmt.Scan(&N, &X)

	for i := 0; i < N; i++ {
		var s []string
		var narr []int
		if sc.Scan() {
			s = strings.Fields(sc.Text())
		}

		for _, v := range s {
			a, _ := strconv.Atoi(v)
			narr = append(narr, a)
		}

		t := Taxi{
			a:     narr[0],
			b:     narr[1],
			c:     narr[2],
			d:     narr[3],
			money: narr[1],
		}

		Tarr = append(Tarr, t)
	}

	for i, v := range Tarr {
		remainLen := X - v.a

		fmt.Println(remainLen)
		fmt.Println(Tarr[i].money)

		for remainLen >= 0 {
			Tarr[i].money += v.d
			fmt.Println(Tarr[i].money)
			remainLen = remainLen - v.c
		}

	}

	var moneys []int
	for _, v := range Tarr {
		moneys = append(moneys, v.money)
	}

	sort.Ints(moneys)

	fmt.Println(moneys[0], moneys[len(moneys)-1])
}
