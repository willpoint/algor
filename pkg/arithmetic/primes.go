package arithmetic

// SieveOfEratosthenes is performed following
// an idea of a sift, where all numbers not prime
// fall-through and the primes remain at the end
// it starts with a list of all candidate numbers and
// then crosses out the ones known not to be primes
// since prime numbers cannot be even numbers
// we start with odd number candidates with m as
// the maximum number to perform the sieve on
