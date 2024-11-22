package main

type pair struct {
	// transformation string
	tf string
	// transform length
	l int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// create a set from the list of valid words to ease the search to O(1)
	wordsMap := make(map[string]bool)

	for _, v := range wordList {
		wordsMap[v] = true
	}

	var q []pair
	q = append(q, pair{beginWord, 1})

	for len(q) > 0 {
		e := q[0]
		currWord, currlen := e.tf, e.l
		q = q[1:]

		if currWord == endWord {
			return currlen
		}

		n := len(currWord)
		for i := 0; i < n; i++ {
			var nWord string
			for ch := 'a'; ch <= 'z'; ch++ {
				if i == 0 {
					nWord = string(ch) + currWord[i+1:]
				} else if i == n-1 {
					nWord = currWord[:n-1] + string(ch)
				} else {
					nWord = currWord[:i] + string(ch) + currWord[i+1:]
				}
				if wordsMap[nWord] {
					q = append(q, pair{nWord, currlen + 1})
					wordsMap[nWord] = false
				}
			}
		}
	}

	return 0
}
