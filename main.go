package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Ingredient struct {
	name   string
	number int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var n int
	if sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
	}

	var recipes []Ingredient

	for i := 0; i < n; i++ {
		if sc.Scan() {
			strs := strings.Split(sc.Text(), " ")
			name := strs[0]
			number, _ := strconv.Atoi(strs[1])
			ingredient := Ingredient{
				name:   name,
				number: number,
			}
			recipes = append(recipes, ingredient)
		}
	}

	var m int
	if sc.Scan() {
		m, _ = strconv.Atoi(sc.Text())
	}

	var possessions []Ingredient

	for i := 0; i < m; i++ {
		if sc.Scan() {
			strs := strings.Split(sc.Text(), " ")
			name := strs[0]
			number, _ := strconv.Atoi(strs[1])
			ingredient := Ingredient{
				name:   name,
				number: number,
			}
			possessions = append(possessions, ingredient)
		}
	}

	var count []int
	var check []int

	for _, recipe := range recipes {
		for _, possession := range possessions {
			if recipe.name == possession.name {
				count = append(count, possession.number/recipe.number)
			}
		}
	}

	if reflect.DeepEqual(count, check) {
		count = append(count, 0)
	}

	fmt.Println(count, check)
	fmt.Println(minInt(count))

}

func minInt(a []int) int {
	sort.Sort(sort.IntSlice(a))
	return a[0]
}
