package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var m = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func intToRoman(num int) string {
	var result []string

	var base int
	for n := 1000; n > 0; n /= 10 {
		if r := num / n; r > 0 {
			base = n
			break
		}
	}

	for reminder := num; reminder > 0; base /= 10 {
		f := reminder / base
		reminder = reminder % base
		if f == 9 {
			result = append(result, m[9*base])
		} else if f > 5 {
			result = append(result, m[5*base])
			for i := 0; i < f-5; i++ {
				result = append(result, m[base])
			}
		} else if f == 5 || f == 4 {
			result = append(result, m[f*base])
		} else if f < 4 {
			for i := 0; i < f; i++ {
				result = append(result, m[base])
			}
		}
	}

	return strings.Join(result, "")
}

func TestIntoToRoman(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("MM", intToRoman(2000))
	assert.Equal("MMCM", intToRoman(2900))
	assert.Equal("LVIII", intToRoman(58))
	assert.Equal("MCMXCIV", intToRoman(1994))
}
