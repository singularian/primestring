package main

import (
      "fmt"
      "bytes"
      "math/big"
      "os"
      "strconv"
)

func IsPrime(n int) bool {
      if n < 2 {
              return false
      }

      for i := 2; i*i <= n; i++ {
              if n%i == 0 {
                      return false
              }
      }
      return true
}

// variations of a prime bit string
// 100100010000010000000100000000000100000000000001
// 
// this is a run of alternating 0's and 1's 
// the first 00 group is one prime number n bits 0
// the following 1 bit group is one prime number delineated by n bits 1
// 00111000001111111000000000001111111111111
//
// break the bigint number into constituent groups of 60 numbers
// then concatenate each block together
// formula with concatenation groups
// 421612065005906503886017970712759923748952792881600196162550115244998622073006303405395430776097230328628476580766070593761091990440640513
// 42161206500590650388601797071275992374895279288160019616255 01152449986220730063034053954307760972303286284765807660705 93761091990440640513
// 
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

      buffer.WriteString("0")
      for j := 0; j < prime; j++ {
         buffer.WriteString("1")
      }

   }

   buffer.WriteString("0")

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

func primesToBin2(primenumber int) int {

   var count int  = 0
   var number int = 2
   var buffer bytes.Buffer
   for {
      if IsPrime(number) {
         count++
         buffer.WriteString("0")
         for j := 0; j < number; j++ {
            buffer.WriteString("1")
         }

      }
      if count == primenumber {
         break
      }
      number++
   }

   buffer.WriteString("0")

   fmt.Println("prime bitstring buffer ", buffer.String())

   n := new(big.Int)
   n, ok := n.SetString(buffer.String(), 2)
   if !ok {
        fmt.Println("SetString: error")
        return 0 
   }


   fmt.Println("prime bitstring bigint", n)

   return 0
}

func main() {

      var argsNum int = len(os.Args)

      if argsNum < 2 { 
		fmt.Println("Primestring Usage: \n")
                fmt.Println(os.Args[0],  "[Number]\n")
                fmt.Println(os.Args[0], "2 3 5 7" )
                fmt.Println("prime bitstring buffer  0110111011111011111110")
                fmt.Println("prime bitstring bigint 1818366")
		os.Exit(1)
      }


      var primes []int
      for i := range os.Args {
         arg := os.Args[i]

         // fmt.Println("args ", i, arg)
         p, _ := strconv.Atoi(arg)
         if p != 0 {
             primes = append(primes, p)
         }
      }
      primesToBin(primes)
}
