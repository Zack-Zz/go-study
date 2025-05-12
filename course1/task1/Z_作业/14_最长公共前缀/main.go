package main

func main() {

}

// 纵向比较
func longestCommonPrefix_ver(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[j-1][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

// 横向寻找
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)

	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}

	return prefix
}

func lcp(prefix, strVal string) string {
	length := min(len(prefix), len(strVal))
	idx := 0
	for idx < length && prefix[idx] == strVal[idx] {
		idx++
	}
	return prefix[:idx]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
