package main

import (
	"fmt"
)

func Add(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func Substract(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1 - num2
	} else {
		result = num2 - num1
	}
	return result
}

func main() {
	fmt.Println("****Unit Testing****")
	numbers := []int{1, 2, 3, 4, 5}
	ans := Add(numbers)
	fmt.Println("Sum is", ans)

	ans2 := Substract(65, 45)
	fmt.Println("Result after substraction is", ans2)
}
