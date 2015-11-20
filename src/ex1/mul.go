package main

import "fmt"

func sumFiveOrThreeTo(limit int) int {
  sum := 0
  for i:=0; i< limit; i++ {
    if(i % 3 == 0  || i % 5 ==0){
      sum += i
    }
  }
  return sum
}

func main(){
  //var ola string = "Ola Euler?"
  fmt.Println(sumFiveOrThreeTo(1000))
}
