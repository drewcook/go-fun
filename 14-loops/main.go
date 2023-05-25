package main

import "fmt"

func main() {
	fmt.Println("loops in golang")

	// slices
	days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}
	fmt.Println(days)

	fmt.Println("---")

	// for loop with iterator
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	fmt.Println("---")

	// range loop - useful for maps, slices, arrays
	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("---")

	// for each loop - gives index and value
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}
	fmt.Println("---")
	for _, day := range days {
		fmt.Printf("value is %v\n", day)
	}

	fmt.Println("---")

	// break & continue in a while loop
	rogueValue := 1
	for rogueValue < 10 {
		fmt.Println("Value is: ", rogueValue)

		// if rogueValue == 5 {
		// 	break // exits the loop
		// }

		if rogueValue == 5 {
			rogueValue++
			// skips over a particular phase
			continue
		}

		rogueValue++

		// goto usage
		switch {
		case rogueValue == 1:
			fmt.Println("one is fun")
		case rogueValue == 2:
			fmt.Println("value is deuce")
		case rogueValue == 3:
		case rogueValue == 1:
			goto den
		default:
			goto dco
		}

		if rogueValue == 2 {
			// this jumps out of the loop and goes to the definition of the goto; in this case a log statement
			// goto dco
			goto den
		}
	}

	// goto
dco:
	fmt.Println("dco is learning golang")
den:
	fmt.Println("flying to denver")
}
