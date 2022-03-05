package main

import "fmt"

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("Hello,Go!")

	fmt.Println("Fizzbuzz() test")
	fmt.Printf("Fizzbuzz(%v) = %v\n", 10, Fizzbuzz(12))
	fmt.Printf("Fizzbuzz(%v) = %v\n", 15, Fizzbuzz(25))
	fmt.Printf("Fizzbuzz(%v) = %v\n", 16, Fizzbuzz(30))

	fmt.Println("IsPrime() test")
	fmt.Printf("IsPrime(%v) = %v\n", 5, IsPrime(5))
	fmt.Printf("IsPrime(%v) = %v\n", 6, IsPrime(6))
	fmt.Printf("IsPrime(%v) = %v\n", 11, IsPrime(11))
	fmt.Printf("IsPrime(%v) = %v\n", 32, IsPrime(32))

	fmt.Println("IsPalindrome() test")
	fmt.Printf("IsPalindrome(%v) = %v\n", "258852", IsPalindrome("258852"))
	fmt.Printf("IsPalindrome(%v) = %v\n", "258855", IsPalindrome("258855"))
	fmt.Printf("IsPalindrome(%v) = %v\n", "258552", IsPalindrome("258552"))
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	var s string

	if n%3 == 0 {
		s += "Fizz"
	}

	if n%5 == 0 {
		s += "Buzz"
	}

	return s
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	if n <= 2 {
		return true
	}

	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}