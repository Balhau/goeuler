package main

import (
  "fmt"
)

func squareSum(val int) int {
  sum:=0
  for i:=1;i<=val;i++ {
    sum += i
  }
  return sum*sum
}

func sumSquare(val int) int{
  sum:=0
  for i:=1;i<=val;i++ {
    sum += i*i
  }
  return sum
}

func main(){
  fmt.Println("Squares: ")
  fmt.Println(squareSum(100))
  fmt.Println(sumSquare(100))
  fmt.Println((squareSum(100)-sumSquare(100)))
}
