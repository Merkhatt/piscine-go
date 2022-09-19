package main

import "fmt"

func main() {

	var a int
	var b int
	var sum int

	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if b > 9 && 100 > b && b%8 == 0 {
			sum += b
		}
	}
	fmt.Print(sum)
}
