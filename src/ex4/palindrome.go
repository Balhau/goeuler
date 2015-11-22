package main

import (
  "fmt"
)

func isPalindrome(v int) bool {
  if (v < 0) {
        return false;
  }
  left:=0
  right:=0
  div := 1;
  for v / div >= 10 {
      div *= 10;
  }
  for (v != 0) {
      left = v / div;
      right = v % 10;
      if (left != right){
        return false;
      }
      v = (v % div) / 10;
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
