package main

import "fmt"

func main() {

	total := 0

	for i := 1; i < 100; i++ {

		if !(i%3 == 0) {
			continue
		}

		fmt.Println(i)

		total += 1

	}

	fmt.Printf("Total de números: %d", total)

}
