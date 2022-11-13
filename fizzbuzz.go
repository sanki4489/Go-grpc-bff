package main

import (
	"fmt"
)

func fizzbuzz() {
	for num := 0; num <= 20; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Println("fizz Buzz")
		} else if num%5 == 0 {
			fmt.Println("Buzz")
		} else if num%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(num)
		}
	}
}
