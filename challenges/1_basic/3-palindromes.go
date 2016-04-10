package main

import (
	"fmt"
)

func main() {
	/* PALINDROMES
	Return `true` if the given string is a palindrome (https://en.wikipedia.org/wiki/Palindrome).
	Otherwise, return `false`.
	*/

	palindrome("eye")               // => true
	palindrome("race car")          // => true
	palindrome("not a palindrome")  // => false
	palindrome("never odd or even") // => true
	palindrome("nope")              // => false

}
