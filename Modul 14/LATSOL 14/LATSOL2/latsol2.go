package main

import "fmt"

func main() {
	var b int
	fmt.Scan(&b)
	if b == 1 {
		fmt.Println("Bukan Prima")
		return
	}
	isPrime := true
	for i := 2; i*i <= b; i++ {
		if b%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println("Prima")
	} else {
		fmt.Println("Bukan Prima")
	}
}
