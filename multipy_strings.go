package leetcode

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// so that len(num1) >= len(num2)
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	var result string

	for c2, times := len(num2)-1, 0; c2 >= 0; times++ {
		padding := ""
		for i := 0; i < times; i++ {
			padding += "0"
		}

		tmp := fmt.Sprintf("%v%v", multiplyOneDigit(num1, rune(num2[c2])), padding)
		fmt.Println(tmp)
		result = addStrings(result, tmp)
		c2--
	}
	return fmt.Sprint(result)
}

func multiplyOneDigit(num string, digit rune) string {
	advancer := 0
	reminder := 0

	d := int(digit - rune('0'))

	var result = make([]byte, len(num)+1)
	for i := len(num) - 1; i >= 0; i-- {
		// calculate next
		n := int(rune(num[i]) - rune('0'))
		// set for the current slot
		reminder = (n * d) % 10
		if (reminder + advancer) >= 10 {
			result[i+1] = byte((reminder+advancer)%10 + '0')
			advancer = (n*d)/10 + 1
		} else {
			result[i+1] = byte(reminder + advancer + '0')
			advancer = (n * d) / 10
		}
	}

	if advancer > 0 {
		result[0] = byte(advancer + '0')
	} else {
		result = result[1:]
	}
	return string(result)
}

func addStrings(s1, s2 string) string {
	advancer := 0
	reminder := 0

	max := len(s1)
	if max < len(s2) {
		max = len(s2)
	}

	var result = make([]byte, max+1)

	var n1, n2 int
	for i, c1, c2 := max, len(s1)-1, len(s2)-1; c1 >= 0 || c2 >= 0; i, c1, c2 = i-1, c1-1, c2-1 {
		if c1 >= 0 {
			n1 = int(rune(s1[c1]) - rune('0'))
		} else {
			n1 = 0
		}

		if c2 >= 0 {
			n2 = int(rune(s2[c2]) - rune('0'))
		} else {
			n2 = 0
		}

		reminder = (n1 + n2) % 10
		if (reminder + advancer) >= 10 {
			result[i] = byte((reminder+advancer)%10 + '0')
			advancer = (n1+n2)/10 + 1
		} else {
			result[i] = byte((reminder + advancer) + '0')
			advancer = (n1 + n2) / 10
		}
	}

	if advancer > 0 {
		result[0] = byte(advancer + '0')
	} else {
		result = result[1:]
	}
	return string(result)
}

func multiplyFastest(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	// The product can be stored in the corresponding position.
	// Note: []int instead of []rune since in Go, rune is int32. To consider for overflow in case of carry, int is chosen
	temp := make([]int, len(num1)+len(num2))
	// rune represents a Unicode code point
	n1 := []rune(num1)
	n2 := []rune(num2)
	for i := 0; i < len(n1); i++ {
		for j := 0; j < len(n2); j++ {
			temp[i+j+1] += int(n1[i]-'0') * int(n2[j]-'0')
		}
	}
	for i := len(temp) - 1; i > 0; i-- {
		temp[i-1] += temp[i] / 10
		temp[i] = temp[i] % 10
	}
	// temp[0] can be 0 is num1 and num2 are small and no carry
	// We will trim in case there is a zero in the first position
	if temp[0] == 0 {
		temp = temp[1:]
	}
	result := make([]byte, len(temp))
	for i := 0; i < len(temp); i++ {
		result[i] = byte(temp[i]) + '0'
	}
	return string(result)
}
