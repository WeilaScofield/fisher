/*给你一份『词汇表』（字符串数组） words 和一张『字母表』（字符串） chars。
假如你可以用 chars 中的『字母』（字符）拼写出 words 中的某个『单词』
（字符串），那么我们就认为你掌握了这个单词。
注意：每次拼写时，chars 中的每个字母都只能用一次。
返回词汇表 words 中你掌握的所有单词的 长度之和。*/

package leetcode

import (
	"strings"
)

func CountCharacters(words []string, chars string) int {
	if len(words) == 0 || chars == "" {
		return 0
	}

	totalWords := 0
	unMaterWords := 0
	for _, v := range words {
		if len(v) > len(chars) {
			continue
		}
		temp := chars
		totalWords += len(v)
		for _, r := range v {
			if !strings.ContainsRune(temp, r) {
				unMaterWords += len(v)
				break
			}
			temp = strings.Replace(temp, string(r), "", 1)
		}
	}

	return totalWords - unMaterWords
}
