package main

import "fmt"

func main() {
	var matriz [3][4]int

	for linha:= range matriz {

		for coluna := range matriz[linha] {
			for coluna := range matriz[linha] {
				matriz[linha][coluna] = linha + coluna
			}
		}
		media:= soma / float64(len(numeros))
	fmt.Println(media)
}
