package alonestrings

import (
	"strings"
	"unicode"
)

// 字符串回文
func IsPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j { // 用小于即ok
		s1 := rune(s[i])
		s2 := rune(s[j])
		if !isLetterOrNumber(s1) {
			i++
			continue
		}
		if !isLetterOrNumber(s2) {
			j--
			continue
		}
		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func isLetterOrNumber(s rune) bool {
	return unicode.IsLetter(s) || unicode.IsNumber(s)
}

func ReverseString(s []byte) string {
	i, j := 0, len(s)-1
	for ; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func ReverseStr(s string, k int) string {
	runes := []rune(s)
	n := len(s)
	for i := 0; i < n; i += 2 * k {
		// 1. 每隔 2k 个字符的前 k 个字符进行反转
		// 2. 剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符
		if i+k <= n {
			reverseRangeString(runes, i, i+k-1)
			continue
		}
		// 3. 剩余字符少于 k 个，则将剩余字符全部反转。
		reverseRangeString(runes, i, n-1)
	}
	return string(runes)
}

func ReverseStr2(s string, k int) string {
	runes := []rune(s)
	n := len(s)
	i := 0
	for i < n {
		// 剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符
		if i+k <= n {
			reverseRangeString(runes, i, i+k-1)
		} else { // 剩余字符少于 k 个，则将剩余字符全部反转。
			reverseRangeString(runes, i, n-1)
		}
		// 每隔 2k 个字符的前 k 个字符进行反转
		i += 2 * k
	}
	return string(runes)
}

func reverseRangeString(r []rune, start int, end int) {

	for i, j := start, end; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return
}

func ReplaceSpace(s string) string {
	if s == "" {
		return s
	}
	builder := strings.Builder{}
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		//if runes[i] == ' ' { // 用字符比较也OK
		//
		//}
		if unicode.IsSpace(runes[i]) {
			builder.WriteString("%20")
		} else {
			builder.WriteRune(runes[i])
		}
	}
	return builder.String()
}
