
/**
* Check more info on http://www.maths.surrey.ac.uk/hosted-sites/R.Knott/Fibonacci/fibFormula.html#logs
*/

package main

import (
  "fmt"
  "math"
)

func digitsFibonacci(n int) int{
  phi:=(1+math.Sqrt(5))/2
  return int(math.Ceil(float64(n)*math.Log10(phi) - math.Log10(math.Sqrt(5))))
}

func main(){
  fmt.Println("Digitos Fibonacci: ")
  digitos:=0
  for n:=1;digitos<1000;n++{
    digitos=digitsFibonacci(n);
    fmt.Println(digitos,"-",n)
  }
}
