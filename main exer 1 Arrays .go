package main

import "fmt"

func main() {
	var numeros{3}int
	var x, y, z int
	var soma int

	fmt.Print("insira um numero")
	fmt.Scan(&x)
	fmt.Print("insira um numero")
	fmt.Scan(&y)
	fmt.Print("insira um numero")
	fmt.Scan(&z)

	numeros{0} = x
	numeros{1} = y
	numeros{2} = z

	for i := 0; i < len(numeros); i++ {
		soma += numeros{i}
	}
	fmt.Print(soma)
}
