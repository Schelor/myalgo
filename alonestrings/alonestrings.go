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

func ReplaceSpace2(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	newRunes := make([]rune, len(runes)*3)
	size := 0
	for i := 0; i < len(runes); i++ {
		// 用字符比较也OK
		if runes[i] == ' ' {
			newRunes[size] = '%'
			newRunes[size+1] = '2'
			newRunes[size+2] = '0'
			size = size + 3
		} else {
			newRunes[size] = runes[i]
			size = size + 1
		}
	}
	return string(newRunes[0:size])
}

func ReverseWords(s string) string {
	if s == "" {
		return s
	}
	// 1.去除多余的空格
	chars := removeSpace(s)
	// 2.反转字符数组
	reverseChars(&chars, 0, len(chars)-1)
	// 3.反转单词(按空格拆分)
	reverseEachWord(&chars, len(chars))

	return string(chars)
}

// 去除字符串多余的空格
// 传地址，减少空间申请
func removeSpace(s string) []rune {
	var chars = []rune(s)
	slow, fast, end := 0, 0, len(s)-1
	// 去重首尾多余的空格
	for fast < end && chars[fast] == ' ' {
		fast++
	}
	for end > fast && chars[end] == ' ' {
		end--
	}
	for fast <= end {
		// 查找连续的空格,如果找到,则fast继续
		if fast > 0 && chars[fast] == chars[fast-1] && chars[fast] == ' ' {
			fast++
			continue
		}
		// slow始终表示下一个OK的字符：即正常的字母,单个空格
		// 如果找到非连续的空格,则将fast位置的字符放到slow位置,慢指针slow++
		chars[slow] = chars[fast]
		slow++
		fast++
	}
	return chars[0:slow] // 读取0~slow直接的切片
}

// 传地址/引用
func reverseChars(chars *[]rune, start int, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		(*chars)[i], (*chars)[j] = (*chars)[j], (*chars)[i]
	}
}

// 切片是引用传递，也可以使用下面的方式
func reverseChars2(chars []rune, start int, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
}

func reverseEachWord(chars *[]rune, length int) {
	start, end := 0, 0
	for start < length {
		for end < length && (*chars)[end] != ' ' {
			end++
		}
		reverseChars(chars, start, end-1)
		start = end + 1
		end = start
	}
}

// 字符串左旋
func ReverseLeftWords(s string, n int) string {
	builder := strings.Builder{}
	for i := n; i < len(s); i++ {
		builder.WriteRune(rune(s[i]))
	}
	for i := 0; i < n; i++ {
		builder.WriteRune(rune(s[i]))
	}
	return builder.String()
}

func ReverseLeftWords2(s string, n int) string {

	return s[n:] + s[0:n]
}

func ReverseLeftWords3(s string, n int) string {
	chars := []rune(s)
	reverseChars2(chars, 0, n-1)
	reverseChars2(chars, n, len(s)-1)
	reverseChars2(chars, 0, len(s)-1)
	return string(chars)
}

// StrStr 字符串匹配,main为主串,长度为n,pattern为模式串,长度为m
// 在长度为n的主串中查找长度为m个模式串,存在n-m+1个子串
func StrStr(main string, pattern string) int {
	n := len(main)
	m := len(pattern)
	for i := 0; i <= n-m; i++ {
		// 依次从每一个i匹配m个字符,如果遇到不匹配的字符，直接终止匹配
		// 每一个i取对应的,i,i+1,i+2,i+m-1与模式串的第0个,第1个,第m-1个字符匹配
		matched := true
		for j := 0; j < m; j++ {
			if main[i+j] != pattern[j] {
				matched = false
				break
			}
		}
		// 返回首次匹配
		if matched {
			return i
		}
	}
	return -1
}

func StrStr2(main string, pattern string) int {
	return strings.Index(main, pattern)
}
