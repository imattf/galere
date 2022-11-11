```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	Demo(1)
	Demo()
	Demo(1, 2, 3)
	numbers := []int{4, 5, 6}
	fmt.Println(Demo(numbers...))
}

func Demo(numbers ...int) int {
	sum := 0
	fmt.Println("Hello, the number(s):", numbers)

	for _, number := range numbers {
		sum += number
	}
	return sum
}
```