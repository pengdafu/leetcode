package leetcode

const (
	int32Max = 1<<31 - 1
	int32Min = -1 << 31
)

/**
1、先去除字符串前面的空格
2、找到有效字符串后停下
3、记录正负
*/
func myAtoi(str string) int {
	index := 0
	reg := 1
	flag := 0
	for i, c := range str {
		index = i
		if c != ' ' {
			break
		}
	}
	str = str[index:]
	for i, c := range str {
		index = i
		if flag > 1 {
			return 0
		}
		if c >= '0' && c <= '9' {
			break
		} else if c == '-' {
			flag++
			reg = -1
		} else if c == '+' {
			flag++
			reg = 1
		} else if c < '0' || c > '9' {
			return 0
		}
	}
	result := 0
	for _, c := range str[index:] {
		if c < '0' || c > '9' {
			return reg * result
		}
		result = result*10 + int(c-'0')
		if tmp := reg * result; reg == -1 && tmp < int32Min {
			return int32Min
		} else if reg == 1 && tmp > int32Max {
			return int32Max
		}
	}
	return reg * result
}
