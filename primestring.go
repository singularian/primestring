package main

import (
      "fmt"
      "bytes"
      "math/big"
      "os"
      "strconv"
)

func SieveOfEratosthenes(n int) []int {
      // Create a boolean array "prime[0..n]" and initialize
      // all entries it as true. A value in prime[i] will
      // finally be false if i is Not a prime, else true.
      integers := make([]bool, n+1)
      for i := 2; i < n+1; i++ {
              integers[i] = true
      }

      for p := 2; p*p <= n; p++ {
              // If integers[p] is not changed, then it is a prime
              if integers[p] == true {
                      // Update all multiples of p
                      for i := p * 2; i <= n; i += p {
                              integers[i] = false
                      }
              }
      }

      // return all prime numbers <= n
      var primes []int
      for p := 2; p <= n; p++ {
              if integers[p] == true {
                      primes = append(primes, p)
              }
      }

      return primes
}

func primesToBin(arr []int) int {
   // var i int
   //var j int
   var avg int
   var size int
   //var prime int

   size = len(arr)

   var buffer bytes.Buffer

   for i := 0; i < size; i++ {
      prime := arr[i]

      buffer.WriteString("1")
      for j := 0; j < prime; j++ {
         buffer.WriteString("0")
      }

   }

   buffer.WriteString("1")

   fmt.Println("prime bitstring buffer ", buffer.String())

   n := new(big.Int)
   n, ok := n.SetString(buffer.String(), 2)
   if !ok {
        fmt.Println("SetString: error")
        return avg
   }


   fmt.Println("prime bitstring bigint", n)

   return avg


}

func main() {

      // argCount := len(os.Args[1:])
      for i := range os.Args {
         arg := os.Args[i]

         // fmt.Println("args ", i, arg)
         p, _ := strconv.Atoi(arg)
         if p != 0 {
             primesToBin(SieveOfEratosthenes(p))
         }
      }
}
