package main

import "math"
import "fmt"

/* Problem 7
     By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
     What is the 10 001st prime number?
*/

func isPrime(nb int) bool {
	i := 1
	limit := int(math.Sqrt(float64(nb))) // the square of the number (int : can't % on float...)
	for i < limit {
		i++
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var rank int

	fmt.Print("The rank of the prime number you want : ")
	fmt.Scan(&rank)

	primeNb := 1
	for cpt := 1; cpt <= rank; cpt++ {
		primeNb++
		for !isPrime(primeNb) {
			primeNb++
		}
	}

	fmt.Println("\nThe", rank, "th prime number is ", primeNb)
}
