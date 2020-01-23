package main

import (
      "fmt"
)

// Prime factorization of a given number
//
// A map is returned.
//
//   key of map: prime
//   value of map: prime exponents
//
func PrimeFactorization(n int) (pfs map[int]int) {
      pfs = make(map[int]int)

      // Get the number of 2s that divide n
      for n%2 == 0 {
              if _, ok := pfs[2]; ok {
                      pfs[2] += 1
              } else {
                      pfs[2] = 1
              }
              n = n / 2
      }

      // n must be odd at this point. so we can skip one element
      // (note i = i + 2)
      for i := 3; i*i <= n; i = i + 2 {
              // while i divides n, append i and divide n
              for n%i == 0 {
                      if _, ok := pfs[i]; ok {
                              pfs[i] += 1
                      } else {
                              pfs[i] = 1
                      }
                      n = n / i
              }
      }

      // This condition is to handle the case when n is a prime number
      // greater than 2
      if n > 2 {
              pfs[n] = 1
      }

      return
}

// Calculate number of divisors of a given number
func NumberOfDivisors(n int) int {
      pfs := PrimeFactorization(n)

      num := 1
      for _, exponents := range pfs {
              num *= (exponents + 1)
      }

      return num
}

// return n-th triangular number
func TriangularNumber(n int) int {
      return n * (n + 1) / 2
}

func main() {
      for i := 1; i < 100000; i++ {
              n := TriangularNumber(i)
              nn := NumberOfDivisors(n)
              if nn >= 500 {
                      fmt.Print(i)
                      fmt.Print("-th ")

                      fmt.Print(n)
                      fmt.Print(": ")
                      fmt.Println(nn)
                      break
              }
      }
}
