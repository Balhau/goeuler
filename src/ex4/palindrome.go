package main

import (
  "fmt"
)

func isPalindrome(x int) bool {
  if (x < 0) {
        return false;
  }
  l:=0
  r:=0
  div := 1;
  for x / div >= 10 {
      div *= 10;
  }
  for (x != 0) {
      l = x / div;
      r = x % 10;
      if (l != r){
        return false;
      }
      x = (x % div) / 10;
      div /= 100;
  }
  return true;
}

func computeBiggestPalindrome(dim int) int {
  palindrome:=0
  aux := 0
  for i:=0;i<dim;i++ {
    for j:=0;j<dim;j++ {
      aux = i*j
      if(isPalindrome(aux) && aux > palindrome){
        palindrome=aux
      }
    }
  }
  return palindrome
}

func main(){
  fmt.Print("É palindrome: ")
  fmt.Println(computeBiggestPalindrome(1000))
}
