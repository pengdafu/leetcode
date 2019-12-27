package leetcode

import "testing"

func TestConvert(t *testing.T) {
	str := "LEETCODEISHIRING"
	numRows := 3
	if ans := Convert(str, numRows); ans != "LCIRETOESIIGEDHN" {
		t.Error(ans)
	}
}
