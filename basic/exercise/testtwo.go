package exercise

import "fmt"

func LengthOfLongestSubstring() int {
	s := "abcabcbb"
	m := map[byte]int{}
	n := len(s)

	rk, ans := 0, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}

		for rk < n && m[s[rk]] == 0 {
			m[s[rk]]++
			rk++
		}

		ans = max(ans, rk-i)
	}
	fmt.Println(ans)
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}
