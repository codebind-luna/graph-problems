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

		for i := 0; i < len(currWord); i++ {
			var newTf string
			for ch := 'a'; ch <= 'z'; ch++ {
				var replace string
				if currWord[i] != byte(ch) {
					replace = string(ch)
				}

				if replace != "" {
					if i == 0 {
						newTf = replace + currWord[i+1:]
					} else if i == len(currWord)-1 {
						newTf = currWord[:len(currWord)-1] + replace
					} else {
						newTf = currWord[:i] + replace + currWord[i+1:]
					}
					if wordsMap[newTf] {
						q = append(q, pair{newTf, currlen + 1})
						wordsMap[newTf] = false
					}
				}
			}
		}
	}

	return 0
}
