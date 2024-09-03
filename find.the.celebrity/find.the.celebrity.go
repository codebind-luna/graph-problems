package main

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
// func solution(knows func(a int, b int) bool) func(n int) int {
// 	return func(n int) int {
// 		// Compare and eliminate one that is not celebrity. Compare with the next number until reaching the last number.
// 		candidate := 0
// 		for j := 1; j < n; j++ {
// 			if knows(candidate, j) {
// 				candidate = j // We get a final candidate by linear comparison
// 			}
// 		}

// 		// Check if the final candidate is the celebrity
// 		for k := 0; k < n; k++ {
// 			if candidate != k && (knows(candidate, k) || !knows(k, candidate)) { // If it knows someone, or someone doesn't know it, it's not a celebrity
// 				return -1
// 			}
// 		}
// 		return candidate
// 	}
// }
