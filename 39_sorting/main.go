package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings are ", strs)

	ints := []int{5, 4, 25}
	sort.Ints(ints)
	fmt.Println("Ints ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted ", s)
}
