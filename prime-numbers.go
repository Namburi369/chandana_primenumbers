package main

import "fmt"

func main() {
	sum := 0
	for num := 2; num <= 10; num++ {
		if isPrime(num) {
			sum += num
		}
	}
	fmt.Println("Sum of prime numbers between 1 and 10:", sum)
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
