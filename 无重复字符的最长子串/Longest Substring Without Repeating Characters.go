package longestSubstringWithoutRepeatingCharacters

// 3. 无重复字符的最长子串
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

func LengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring(s)
}

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	var (
		strMap = make(map[int32]int) // 记录字符上一次出现的位置
		result = 0                   // 最终的结果
		start  = 0                   // 不重复子串的开始位置
	)

	if len(s) == 0 {
		return 0
	}

	for ind, v := range s {
		val, ok := strMap[v]
		if ok && start <= val { // 开始位置的区间已经包含了这个字符，说明重复了呗
			start = val + 1 // 更新新的开始位置
		}

		// 求出新的子区间长度
		tmp := ind - start + 1
		if result < tmp {
			result = tmp
		}

		// 记录这个字符的出现的位置
		strMap[v] = ind
	}

	return result
}
