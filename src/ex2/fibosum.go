package main

import "fmt"

func fiboEven(limit int) int {
  sum := 0
  prev := 1
  now := 2
  aux := 0
  for now < limit {
    if(now % 2 == 0){
      sum += now
    }
    aux = prev
    prev = now
    now = now + aux
  }
  return sum
}

func main(){
  //var ola string = "Ola Euler?"
  fmt.Println(fiboEven(4000000))
}
