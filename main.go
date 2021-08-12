package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var N, M int
	var dial [][]string
	var strs []string
	if sc.Scan() {
		arr := strings.Fields(sc.Text())
		N, _ = strconv.Atoi(arr[0])
		M, _ = strconv.Atoi(arr[1])
	}

	for i := 0; i < N; i++ {
		if sc.Scan() {
			text := sc.Text()
			slice := strings.Split(text, "")
			dial = append(dial, slice)
		}
	}

	for i := 0; i < M; i++ {
		if sc.Scan() {
			strs = append(strs, sc.Text())
		}
	}

	var count = 0
	var str = ""

	fmt.Println(dial)

	for a := 0; a < N; a++ {
		for b := 0; b < N; b++ {
			for i := 0; i < N-b-a+1; i++ {
				if i+a > 5 && i+b > 5 {
					break
				}
				str += dial[i+a][i+b]
				fmt.Println(str)
				for _, v := range strs {
					if v == str {
						count++
					}
				}
			}
			str = ""
		}
	}

}
