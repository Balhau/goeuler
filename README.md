# Project euler exercises and Solutions

## Instructions
To compile these programs you need to install first the go programming language environment. You can find the instructions on the language official web page. Then you need to setup a environment variable

    export GOPATH=$HOME/work/goeuler

And then you need to complie the programs by executing, for example for the exercise one , the following command

    go build ex1

Hopefully it will create a ex1 executable file on the *goeuler* folder.

## Exercises

#### Multiples of 5 or 3 less than 1000

To find the sum of multiples of 5 or less than 1000 you need to iterate over all the first 1000 elements and check for the rest of the integer division so it is basically to check the following condition

    (i % 3 == 0  || i % 5 ==0)

If this condition holds then we need to add the *i* value to the *sum* counter, so in a nutshell the procedure is as follows

    sum := 0
    for i:=0; i< limit; i++ {
      if(i % 3 == 0  || i % 5 ==0){
        sum += i
      }
    }

#### Even Fibonacci

Here we need to find the sum of even Fibonacci numbers that are lower than 4000000. To do the job all we need is to compute the Fibonacci sequence and sum only those even elements until the condition given stop being verified.

The function follows

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

#### Large prime factor

The problem here is to find the biggest prime factor of the number 600851475143. From number theory we know that the biggest factor of a number is always smaller than the square root of that same number so in this case 600851475143 ^ 1/2. To tackle the problem we split this in two subproblems. First we need to find all the prime numbers until a given condition and the use these set of prime numbers to find which of them divide the number given. The first part is done here

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

The last part of the problem is as follows

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


#### Palindrome number

  Here we were challenged to find the biggest [palindrome](https://en.wikipedia.org/wiki/Palindrome) number that is a result of a multiplication between number of three digits. First we need to find a way to check if a number is a palindrome. For this we do two things, first we need to compute the length of the number and then we pick the first number and the last and we compare each other, after some time we found that could be done like this

      func isPalindrome(v int) bool {
        if (v < 0) {
              return false;
        }
        left:=0
        right:=0
        div := 1;
        for v / div >= 10 {
            div * = 10;
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


Then we use an auxiliary method to compute all the *n* digit number product and check for the biggest palindrome

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


#### Smallest multiples

The challenge here is to find the smallest number that is divisible for all the numbers until 20. Here we don't need a program to iterate over all the solutions. We can use prime number properties and a fundamental theorem that says that any number is a product of prime numbers to find the smallest product. First we need to note that at least the product is bigger than the product of all the prime numbers until 20 so

    value > 1 x 2 x 3 x 5 x 7 x 11 x 13 x 17 x 19

Than we need to find all the prime factorization for all the non prime numbers between 0 and 20, so

    4 = 2 x 2
    6 = 2 x 3
    8 = 2 x 2 x 2
    9 = 3 x 3
    10 = 2 x 5
    12 = 3 x 2 x 2
    14 = 2 x 7
    15 = 3 x 5
    16 = 2 x 2 x 2 x 2
    18 = 2 x 3 x 3
    20 = 5 x 2 x 2

So the smallest number would be

    v = 1 x 2⁴ x 3² x 5 x 7 x 11 x 13 x 17 x 19 = 232792560
