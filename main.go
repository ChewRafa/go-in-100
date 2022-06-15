package main

import (
	"fmt"

	"github.com/ChewRafa/go100days/kata"
)

func main() {

	fmt.Print("Hello World!", "\n\n")
	fmt.Print(kata.ReverseSeq(75), "\n\n")
	fmt.Print(kata.IsPalindrome("banana"), "\n\n")
	fmt.Print(kata.Reverse("Programming"), "\n\n")
	fmt.Print(kata.CountPositivesSumNegatives([]int{1, -2, 3, -4, 5, -6, 7, -8, 9}), "\n\n")
	fmt.Print(kata.MultiTable(7), "\n\n")
	fmt.Print(kata.OtherAngle(30, 45), "\n\n")
	fmt.Print(kata.AbbrevName("John Johnson"), "\n\n")
	fmt.Print(kata.Move(8, 14), "\n\n")
	fmt.Print(kata.GetCount("unexpected"), "\n\n")
	fmt.Print(kata.Divisors(83), "\n\n")
	fmt.Print(kata.NoSpace("this is fine"), "\n\n")
	fmt.Print(kata.EvenOrOdd(94), "\n\n")
}
