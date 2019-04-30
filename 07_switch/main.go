package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write ", i, " as ")

	// Here’s a basic switch.
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("threee")
	}

	// You can use commas to separate multiple expressions in the same case statement.
	//  We use the optional default case in this example as well.

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its Weekend")
	default:
		fmt.Println("Its weekday")
	}

	// switch without an expression is an alternate way to express if/else logic.
	//  Here we also show how the case expressions can be non-constants.

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("Its after noon")
	}

	// A type switch compares types instead of values. You can use this to discover the type of an interface value.
	// In this example, the variable t will have the type corresponding to its clause.

	whatAmI := func(i interface{}) {
		switch t := i.(type) {

		case bool:
			fmt.Println("I am a boolean")
		case int:
			fmt.Println("I am an integer")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(45)
	whatAmI("hey")

}