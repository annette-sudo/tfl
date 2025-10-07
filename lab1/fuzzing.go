package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var Alphabet = []rune{'a', 'b'}

var T = []RewritingRule{
	{left: "abbaba", right: "aabbbaa"},
	{left: "aaa", right: "abab"},
	{left: "abba", right: "baaab"},
	{left: "bbbb", right: "ba"},
}

var T_1 = []RewritingRule{
	{left: "bba", right: "bab"},
	{left: "aaaa", right: "aaa"},
	{left: "aaab", right: "aaa"},
	{left: "abaa", right: "aaa"},
	{left: "abab", right: "aaa"},
	{left: "baaa", right: "aaa"},
	{left: "baba", right: "baab"},
	{left: "bbbb", right: "ba"},
	{left: "baaba", right: "aaa"},
	{left: "baabb", right: "aaa"},
	{left: "babbb", right: "baa"},
}

const maxWordLen = 10
const maxSteps = 100

type RewritingRule struct {
	left, right string
}

func ReplaceN(str, s, t string, n int) string {
	index := 0
	for i := 0; i < n-1; i++ {
		index += strings.Index(str[index:], s) + len(s)
	}
	index += strings.Index(str[index:], s)
	return str[:index] + strings.Replace(str[index:], s, t, 1)
}

func GenerateWords() string {
	var word strings.Builder
	lenght := rand.Intn(maxWordLen)
	for range lenght {
		word.WriteRune(Alphabet[rand.Intn(2)])
	}
	return word.String()
}

func GenerateChain(word string, srs []RewritingRule) ([]string, string) {
	var chain []string
	currentWord := word
	chain = append(chain, currentWord)

	for step := 0; step < maxSteps; step++ {
		orderRules := rand.Perm(len(srs))

		applied := false
		for _, k := range orderRules {
			entryCount := strings.Count(currentWord, srs[k].left)
			if entryCount != 0 {
				changes := rand.Intn(entryCount)
				currentWord = ReplaceN(currentWord, srs[k].left, srs[k].right, changes)
				chain = append(chain, currentWord)
				applied = true
				break
			}
		}
		if !applied {
			break
		}
	}

	return chain, currentWord
}

func CheckEqualance(word, word_1 string, srs_1 []RewritingRule) {
	if word == word_1 {
		chain_0, NF_0 := GenerateChain(word, srs_1)
		fmt.Printf("«%s» and «%s» are equal.\n", word, word_1)
		fmt.Printf("The normal form of this word: %s\n", NF_0)
		fmt.Printf("The chain of transformations: %s\n", chain_0)
		return
	}
	_, NF := GenerateChain(word, srs_1)
	_, NF_1 := GenerateChain(word_1, srs_1)

	if NF == NF_1 {
		fmt.Printf("Testing for words «%s» and «%s». Fuzzing is successful.\n", word, word_1)
		fmt.Printf("The normal form of this words: %s\n", NF)
	} else {
		fmt.Printf("The SRS T and SRS T_1 are not equivalent.")
		fmt.Printf("The normal form of word «%s»: %s\n", word, NF)
		fmt.Printf("The normal form of word «%s»: %s\n", word_1, NF_1)
	}

	return
}

func main() {
	s := GenerateWords()
	_, w := GenerateChain(s, T)
	CheckEqualance(s, w, T_1)
}
