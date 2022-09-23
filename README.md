# primestring

A program to generate integers or hex numbers whose bits or numbers form the prime number sequence.

# Usage Primestring

This primestring variant encodes a bit string where the consecutive 0 numbers form a prime number.
It could also be the inverse

Primestring Usage:

./primestring [Number]

```
server$ ./primestring 11
prime bitstring buffer 1001000100000100000001000000000001000000000000010000000000000000010000000000000000000100000000000000000000000100000000000000000000000000000100000000000000000000000000000001  
prime bitstring bigint 3391050601613190613378084052422250933492987908325377
```
# Ussage

This variant creates a bigint whose hex value is a string of hex encoded primes.

 ./primestringHex 100
hex  2  
hex  3  
hex  5  
hex  7  
hex  b  
hex  d  
hex  11  
hex  13  
hex  17  
hex  1d  
hex  1f  
hex  25  
hex  29  
hex  2b  
hex  2f  
hex  35  
hex  3b  
hex  3d  
hex  43  
hex  47  
hex  49  
hex  4f  
hex  53  
hex  59  
hex  61  
prime bitstring buffer  20305070b0d01101301701d01f02502902b02f03503b03d04304704904f0530590610  
prime bitstring bigint 15266611191208962936317731265925196392934567571098155773419187168253252600258430480  
