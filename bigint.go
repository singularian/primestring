package main

import (
    "fmt"
    "math/big"
    "os"
)

func main() {

      var argsNum int = len(os.Args)

      if argsNum < 2 { 
		fmt.Println("Primestring Bigint Usage: \n")
                fmt.Println(os.Args[0],  "[Number]")
                fmt.Println("Bigint 3391050601613190613378084052422250933492987908325377" )
                fmt.Println("Bigint hex 9104040040010000400004000004000000100000001" )
                fmt.Println("Bigint bin 1001000100000100000001000000000001000000000000010000000000000000010000000000000000000100000000000000000000000100000000000000000000000000000100000000000000000000000000000001" )
		os.Exit(1)
      }

      n := new(big.Int)
      n, ok := n.SetString(os.Args[1], 10)
      if !ok {
        fmt.Println("SetString: error")
        return
      }

      hex := fmt.Sprintf("%x", n)
      bin := fmt.Sprintf("%b", n)

      fmt.Println("bigint ", n)
      fmt.Println("bigint hex ", hex)
      fmt.Println("bigint bin ", bin)
}
