package kata

import (
	"fmt"
	"strings"
)

/*
* Build a function that returns an array of integers from n to 1 where n>0.
* Example : n=5 --> [5,4,3,2,1]
 */
func ReverseSeq(n int) []int {
	var n2 []int
	if n > 0 {
		for i := n; i > 0; i-- {
			n2 = append(n2, i)
		}

	}
	return n2
}

/*
* Write a function that checks if a given string (case insensitive) is a palindrome.
* Example:
* "banana" --> false
* "anana" --> true
*
 */
func IsPalindrome(str string) bool {
	reversed := Reverse(str)
	if strings.ToUpper(str) == strings.ToUpper(reversed) {
		return true
	}
	return false
}

/* this funciton is called by the IsPalindrome function */
func Reverse(str string) string {
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	return reversed
}

/*
* Given an array of integers, return an array, where the first element
* is the count of positives numbers and the second element is sum of
* negative *numbers. 0 is neither positive nor negative.
* If the input is an empty array or is null, return an empty array.
* Example:
* For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15]
* you should return [10, -65].
 */
func CountPositivesSumNegatives(numbers []int) []int {
	var res = []int{0, 0}
	count := 0

	for i := 0; i < len(numbers); i++ {
		negative := 0
		if numbers[i] > 0 {
			count++
			res[0] = count
		} else if numbers[i] < 0 {
			negative += numbers[i]
			res[1] += negative
		}
	}
	return res

}

/*
* Your goal is to return multiplication table for number that is always an integer from 1 to 10.
* For example, a multiplication table (string) for number == 5 looks like below:
*
* 1 * 5 = 5
* 2 * 5 = 10
* 3 * 5 = 15
* 4 * 5 = 20
* 5 * 5 = 25
* 6 * 5 = 30
* 7 * 5 = 35
* 8 * 5 = 40
* 9 * 5 = 45
* 10 * 5 = 50
*
* P. S. You can use \n in string to jump to the next line.
*
Note: newlines should be added between rows, but there should be no trailing newline at the end. If you're unsure about * the format, look at the sample tests.
*/
func MultiTable(number int) string {
	table := ""
	for i := 1; i < 11; i++ {
		table += fmt.Sprintf("%d * %d = %d", i, number, number*i)
		//check if there is a last element
		if i != 10 {
			table += "\n"
		}
	}
	return table
}

/*
* You are given two interior angles (in degrees) of a triangle.
* Write a function to return the 3rd.
* Note: only positive integers will be tested.
 */
func OtherAngle(a int, b int) int {
	return (-a - b + 180)
}

/*
* Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.
* The output should be two capital letters with a dot separating them.
*
* It should look like this:
* Sam Harris => S.H
* patrick feeney => P.F
 */
func AbbrevName(name string) string {
	splits := strings.Split(name, " ")
	newstring := string(splits[0][0])
	newstring += "."
	newstring += string(splits[1][0])

	return strings.ToUpper(newstring)
}

/*
* Terminal game move function
* In this game, the hero moves from left to right. The player rolls the dice and moves
* the number of spaces indicated by the dice two times.
* Create a function for the terminal game that takes the current position of the hero and
* the roll (1-6) and return the new position.
* Example:move(3, 6) should equal 15
 */
func Move(position int, roll int) int {
	return position + roll*2
}

/*
* Return the number (count) of vowels in the given string.
* We will consider a, e, i, o, u as vowels for this Kata (but not y).
* The input string will only consist of lower case letters and/or spaces.
 */
func GetCount(str string) (count int) {
	sumA := strings.Count(str, "a")
	sumE := strings.Count(str, "e")
	sumI := strings.Count(str, "i")
	sumO := strings.Count(str, "o")
	sumU := strings.Count(str, "u")

	totals := sumA + sumE + sumI + sumO + sumU
	return totals
}

/* Count the number of divisors of a positive integer n.
 *
 * Random tests go up to n = 500000.
 * Examples (input --> output)***
 *
 * 4 --> 3 (1, 2, 4)
 * 5 --> 2 (1, 5)
 * 12 --> 6 (1, 2, 3, 4, 6, 12)
 * 30 --> 8 (1, 2, 3, 5, 6, 10, 15, 30)
 */
func Divisors(n int) int {
	totals := 1
	for i := 1; i < n; i++ {
		if n%i == 0 {
			totals++
		}
	}
	return totals
}

/*
 * Simple, remove the spaces from the string, then return the resultant string.
 */
func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

/*
 * Create a function that takes an integer as an argument and returns
 * "Even" for even numbers or "Odd" for odd numbers.
 */
func EvenOrOdd(number int) string {
	message := ""
	if number%2 == 0 {
		message = "Even"
	} else {
		message = "Odd"
	}
	return message
}
