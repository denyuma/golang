package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var M int
	var A []int
	if sc.Scan() {
		arr := strings.Split(sc.Text(), " ")
		M, _ = strconv.Atoi(arr[1])
	}
	if sc.Scan() {
		arr := strings.Split(sc.Text(), " ")
		for _, v := range arr {
			nv, _ := strconv.Atoi(v)
			A = append(A, nv)
		}
	}

	var ans []int
	for ki := 0; ki < M; ki++ {
		var t []int
		for ai := 0; ai < len(A); ai++ {
			// fmt.Println(ai, ki, gcd(ai, ki))
			if gcd(A[ai], ki) == 1 {
				t = append(t, 1)
			} else {
				t = append(t, 0)
			}
		}
		var check []int
		for i := 0; i < len(A); i++ {
			check = append(check, 1)
		}

		if reflect.DeepEqual(t, check) {
			ans = append(ans, ki)
		}
	}

	fmt.Println(len(ans))
	for _, v := range ans {
		fmt.Println(v)
	}
}

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func join(ins []rune, c rune) (result []string) {
// 	for i := 0; i <= len(ins); i++ {
// 		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
// 	}
// 	return
// }

// func permutations(testStr string) []string {
// 	var n func(testStr []rune, p []string) []string
// 	n = func(testStr []rune, p []string) []string {
// 		if len(testStr) == 0 {
// 			return p
// 		} else {
// 			result := []string{}
// 			for _, e := range p {
// 				result = append(result, join([]rune(e), testStr[0])...)
// 			}
// 			return n(testStr[1:], result)
// 		}
// 	}

// 	output := []rune(testStr)
// 	return n(output[1:], []string{string(output[0])})
// }

// func main() {
// 	var S string
// 	var K int
// 	fmt.Scan(&S, &K)
// 	d := permutations(S)
// 	sort.Strings(d)
// 	m := make(map[string]bool)
// 	uniq := []string{}

// 	for _, ele := range d {
// 		if !m[ele] {
// 			m[ele] = true
// 			uniq = append(uniq, ele)
// 		}
// 	}
// 	fmt.Print(uniq[K-1])
// }
