package problem_0003

func HasDuplicate(s string) bool {
	rs := []rune(s)
	encountered := map[rune]bool{}
	for _, r := range rs {
		_, found := encountered[r]
		if found {
			return true
		}
		encountered[r] = true
	}
	return false
}

func LengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring(s)
}

func lengthOfLongestSubstring(s string) int {
	max_length := 0
	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			if !HasDuplicate(s[i:j]) {
				if len(s[i:j]) > max_length {
					max_length = len(s[i:j])
				}
			}
		}
	}
	return max_length
}
