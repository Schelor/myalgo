package integers

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Divide(a int, b int) int {

	return divide(a, b)

}

// 由于入参存在多种情况，防止溢出，需要处理一些边界
func divide(a int, b int) int {
	// 边界1，a=MinInt32
	if a == math.MinInt32 {
		if b == 1 {
			return math.MinInt32
		}
		if b == -1 { // a/b => -2147483648 / -1 => 2147483648 超过int32最大值
			return math.MaxInt32 // 根据题意返回MaxInt32
		}
	}
	// 边界2，b=MinInt32
	if b == math.MinInt32 {
		if a == math.MinInt32 {
			return 1
		}
		return 0
	}
	// 边界3，a=0
	if a == 0 {
		return 0 // 被除数=0
	}
	// 由于a,b存在符号不同,当a=MinInt32,转换为正数会越界,所以统一转换为负数
	var sym = 2
	if a > 0 {
		a = -a
		sym -= 1
	}
	if b > 0 {
		b = -b
		sym -= 1
	}
	// 如果sym=0,表示a,b都是正数，因此结果为正
	// 如果sym=2,表示a,b都是负数，因此结果为正
	// 其他情况则结果为负,即a,b符号不同
	var result = div2(a, b)
	if sym == 0 || sym == 2 {
		return result
	} else {
		return -result
	}
}

func div(a int, b int) int {
	var result = 0
	for a <= b {
		result += 1
		a = a - b
	}
	return result
}

func div2(a int, b int) int {
	var result = 0
	for a <= b {
		temp := b
		count := 1           //a的绝对值大于b的 那么肯定a能减一次b
		for a <= temp+temp { //减数不越界方便控制后一个条件：且a的绝对值比两倍的减数还大
			count += count //可以减的次数翻倍
			temp += temp   //减数也翻倍
		}
		result += count
		a -= temp
	}
	return result
}

func AddBinary(a string, b string) string {
	return addBinary(a, b)
}

func addBinary(a string, b string) string {
	var ra = Reverse(a)
	var rb = Reverse(b)
	var result strings.Builder
	var carry int = 0 // 进位
	// a,b字符串的长度可能不同,通过同一个循环,需要考虑a或b是否遍历到末尾
	// a或b任何一个字符串遍历完成,另一个字符串对应的位置的数字相当于为0,这样可以统一执行对应位数相加
	for i, j := 0, 0; i < len(a) || j < len(b); i, j = i+1, j+1 {
		var ca, cb int = 0, 0
		if i < len(a) {
			ca = int(ra[i]) - int('0')
		}
		if j < len(b) {
			cb = int(rb[j]) - int('0')
		}
		sum := ca + cb + carry
		carry = sum / 2
		currentCh := sum % 2
		result.WriteString(strconv.Itoa(currentCh))
	}
	if carry == 1 {
		result.WriteString("1")
	}
	s := result.String()
	fmt.Println("before=", s)

	return Reverse(s)
}

// Reverse 反转,内部工具函数
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func SingleNumber(nums []int) int {
	return singleNumber2(nums)
}

func singleNumber(nums []int) int {
	// 把所有元素转为map
	var numCnt = map[int]int{}
	for _, n := range nums {
		// 因为map的value为int,默认会返回int的零值,所以可以简化if-else的逻辑为： numCnt[n]++
		cnt, ok := numCnt[n]
		if !ok { // key不在map中,初始化
			numCnt[n] = 1
		} else {
			cnt++
			numCnt[n] = cnt
		}
	}
	// 查找map中count=1
	for k, v := range numCnt {
		if v == 1 {
			return k
		}
	}
	return 0
}

func singleNumber2(nums []int) int {
	// int最大为32位,依次查看数组中每个数的第i位二进制1的个数
	var found = int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0) // total为所有元素第i位二进制=1的个数
		for _, n := range nums {
			oneOrZero := (int32(n) >> i) & 1 // n右移i个位置,可以把第i位移到最右边,然后与1.可得到第i位是1还是0
			total += oneOrZero
		}
		// 遍历所有的元素后,即知道了当前第i位总共有多少个1.
		// total对3求余数,如果=0,说明要查找的数第i位为0,如果不为0,说明要查找的数的第i位为1,需要标记对应i位置的二进制为1
		if total%3 != 0 {
			found |= 1 << i // 第i位为1,把1放到对应的位置上,然后对32个位置执行或操作(|),即可得到完整的二进制（即查找数的底层表示）
		}
	}
	return int(found)
}
