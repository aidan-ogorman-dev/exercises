package main

/*
Given a positive number n > 1 find the prime factor decomposition of n. The result will be a string with the following form :

 "(p1**n1)(p2**n2)...(pk**nk)"

 with the p(i) in increasing order and n(i) empty if n(i) is 1.

Example: n = 86240 should return "(2**5)(5)(7**2)(11)"
*/

import (
	"fmt"
	"math"
	"sort"
)

// prime number can be checked by dividing by all prime numbers < sqrt(n)
func getNextPrime(n int) int {
	for {
		n = n + 1
		sqrtN := int(math.Floor(math.Sqrt(float64(n))))
		isPrime := true
		for i := 2; i <= sqrtN; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			return n
		}
	}
}

func PrimeFactors(n int) string {
	// starting with 2, divide the number 'n' by prime numbers.
	// if remainder is 0, prime is a factor - store it map[int]int e.g.
	// 2 | 3 <-- 2 divides into n 3 times
	// 3 | 2 <-- 3 divides into n 2 times
	// ...
	prime := 2
	result := n
	numsMap := make(map[int]int)
	for result > 1 {
		// if number is divisible by prime, log in the map
		remainder := result % prime
		if remainder == 0 {
			numsMap[prime] += 1
			result = result / prime
		} else {
			prime = getNextPrime(prime)
		}
	}
	primeFactorDecomp := ""

	mapKeys := []int{}
	for k, _ := range numsMap {
		mapKeys = append(mapKeys, k)
	}
	sort.Ints(mapKeys)

	for _, k := range mapKeys {
		mapValue := numsMap[k]
		if mapValue == 1 {
			primeFactorDecomp = primeFactorDecomp + fmt.Sprintf("(%d)", k)
		} else {
			primeFactorDecomp = primeFactorDecomp + fmt.Sprintf("(%d**%d)", k, mapValue)
		}
	}
	return primeFactorDecomp
}
