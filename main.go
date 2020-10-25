package main

import (
	"fmt"
	"strconv"
)

func main() {
	value := 1100000
	for value < 1100002 {
		valueConvert := strconv.Itoa(value)
		num := valueConvert
		oneDigit := [10]string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
		specialDigit := [...]string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
		twoDigit := [...]string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
		otherDigit := [...]string{"", "", "", "Hundred", "Thousand"}

		var word string
		for i := 0; i < len(num); i++ {
			numValues := num[i]
			numValue := convertNumToInt(string(numValues))
			position := len(num) - i

			if position == 2 || position == 5 {
				if numValue < 1 {
					nextNum := convertNumToInt(string(num[i+1]))
					if nextNum < 1 {
						continue
					}
					word += "and "
					continue
				} else if numValue < 2 {
					nextNum := convertNumToInt(string(num[i+1]))
					if position == 2 {
						word += "and " + specialDigit[nextNum]
						break
					} else {
						word += specialDigit[nextNum] + " " + otherDigit[position-1] + " "
						i++
						continue
					}
				} else {
					if position == 5 {
						nextNum := convertNumToInt(string(num[i+1]))
						if nextNum < 1 {
							word += twoDigit[numValue]
							continue
						}
						word += twoDigit[numValue] + "-"
						continue
					}
					nextNum := convertNumToInt(string(num[i+1]))
					if nextNum == 0 {
						word += "and " + twoDigit[numValue]
						continue
					}
					word += "and " + twoDigit[numValue] + "-"
					continue
				}
			}

			if position == 3 {
				if numValue < 1 {
					// word += " "
					continue
				}
			}
			if position == 6 {
				nextNum := convertNumToInt(string(num[i+1]))
				nextNum2 := convertNumToInt(string(num[i+2]))
				if nextNum == 0 {
					if nextNum2 == 0 {
						word += oneDigit[numValue] + " " + otherDigit[3]
						i++
						continue
					}
					word += oneDigit[numValue] + " " + otherDigit[3] + " "
					i++
					continue
				}
				word += oneDigit[numValue] + " " + otherDigit[3] + " "
				continue
			}

			if position == 7 {
				word += oneDigit[numValue] + " " + "million" + " "
				continue
			}
			word += oneDigit[numValue] + " " + otherDigit[position] + " "

		}
		fmt.Println(word)
		value++
	}
}

func convertNumToInt(num string) int {
	numConverted, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println(err)
	}

	return numConverted
}

// var noOfDigits int

// for num != 0 {
// 	remainder := num % 10
// 	num = num / 10
// 	noOfDigits++
// 	if remainder > 0 && noOfDigits == 1 {
// 		fmt.Println(oneDigitInWord[remainder])
// 	} else if remainder > 0 && noOfDigits == 2 {
// 		fmt.Println(twoDigitInWord[remainder])
// 	} else if noOfDigits == 3 {
// 		fmt.Println(oneDigitInWord[remainder], restDigitInWord[0])
// 	}
// }

// fmt.Println(len(num))
// fmt.Printf("%T\n", num[0])
// fmt.Println(oneDigitInWord[string(num[0])])
// slice := []string{}
