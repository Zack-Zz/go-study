package main

/*
 * 前缀树节点
 */
type TrieNode struct {
	//树子节点
	children map[rune]*TrieNode
	//是否是单词的结尾
	isEnd bool
}

func longestCommonPrefix_trie(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 构造root节点
	root := &TrieNode{children: make(map[rune]*TrieNode)}
	for _, str := range strs {
		cur := root
		for _, ch := range str {
			if cur.children[ch] == nil {
				cur.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
			}
			cur = cur.children[ch]
		}
		cur.isEnd = true
	}

	// 遍历树,查找最长公共前缀
	var prefix []rune
	var cur *TrieNode = root
	for {
		if cur.isEnd || len(cur.children) != 1 {
			break
		}
		for ch, next := range cur.children {
			prefix = append(prefix, ch)
			cur = next
			break
		}
	}
	return string(prefix)
}
