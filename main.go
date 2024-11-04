package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func main() {
	fmt.Println("Hello")
	somme := add(3, 4)
	println(somme)
	fmt.Printf("multiplication: %v\n", multiplication(3, 3))
	fmt.Printf("La division: %v\n", division(6.5, 0))
}
