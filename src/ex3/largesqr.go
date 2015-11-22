package main

import (
  "fmt"
  "math"
)

func getFirstPrimes(less int) (primes []int , size int) {
  size = 1
  primes = make([]int,size)
  primes[0]=1
  flag := true
  for candidate:=2; candidate < less ;candidate++ {
    flag=true
    for i:=1; i<size ;i++{
      if(candidate % primes[i] == 0){
        flag=false
      }
    }
    if(flag){
      primes = append(primes,candidate)
      size++
    }
  }
  return
}

func largePrimeFactor(value int) int {
  largePrime:=1
  sqrt:=int(math.Sqrt(float64(value)))
  primes,size := getFirstPrimes(sqrt)
  for i := 0;i<size;i++ {
    if(value % primes[i] == 0){
      largePrime=primes[i]
    }
  }
  return largePrime
}

func main(){
  fmt.Println(getFirstPrimes(10))
  v:=600851475143
  fmt.Println(largePrimeFactor(v))
  //var ola string = "Ola Euler?"
  fmt.Println(math.Sqrt(float64(v)))
}
