package array

import (
	"strconv"
	"strings"
)

// CustomEncoder https://leetcode.com/problems/encode-and-decode-strings/description/
type CustomEncoder struct {
	Delimiter rune
}

type Encoder interface {
	Encode(arr []string) string
	Decode(s string) []string
}

func (c *CustomEncoder) Encode(arr []string) string {
	var builder strings.Builder
	for _, v := range arr {
		builder.WriteString(strconv.Itoa(len(v)))
		builder.WriteRune(c.Delimiter)
		builder.WriteString(v)
	}
	return builder.String()
}

func (c *CustomEncoder) Decode(s string) []string {
	var results []string
	if len(s) == 0 {
		return []string{}
	}
	if !strings.Contains(s, string(c.Delimiter)) {
		return []string{s}
	}
	i := 0
	for i < len(s) {
		j := i
		for s[j] != uint8(c.Delimiter) {
			j++
		}
		length, _ := strconv.Atoi(s[i:j])
		i = j + 1
		results = append(results, s[i:i+length])
		i += length
	}
	return results
}
