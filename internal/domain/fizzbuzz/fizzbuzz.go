package fizzbuzz

import (
	"fmt"
	"strings"
)

type FizzBuzz struct {
}

// NewFizzBuzz creates a new FizzBuzz instance
func NewFizzBuzz() *FizzBuzz {
	return &FizzBuzz{}
}

func (fb *FizzBuzz) Calculate(int1, int2, limit int, str1, str2 string) (string, error) {
	if int1 <= 0 || int2 <= 0 || limit <= 0 {
		return "", fmt.Errorf("int1, int2, and limit must be greater than zero")
	}
	product := int1 * int2
	if int1 == int2 {
		product = int1
	}

	str := strings.Builder{}
	for i := 1; i <= limit; i++ {
		if i%product == 0 {
			str.WriteString(str1 + str2)
		} else if i%int1 == 0 {
			str.WriteString(str1)
		} else if i%int2 == 0 {
			str.WriteString(str2)
		} else {
			str.WriteString(fmt.Sprintf("%d", i))
		}
		str.WriteString(",")
	}
	return str.String()[:str.Len()-1], nil
}
