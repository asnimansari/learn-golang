package main

import "fmt"

func main() {
	/*

		Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).

		To create an empty map, use the builtin make: make(map[key-type]val-type).
	*/
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map :", m)

	v1 := m["k1"]
	fmt.Println("v1", v1)

	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("map after deletion ", m)

	/* The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _.*/
	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	/* Another way of declaring and initizing map */
	n := map[string]int{"foo": 1, "bar": 1}
	fmt.Println("map: ", n)

}
