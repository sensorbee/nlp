package nlp

import (
	"strings"
)

// NGram returns UTF-8 character n-grams created from the given text. This
// function assumes that s only contains valid UTF-8 letters. It returns an
// empty array when n isn't greater than 0.
func NGram(n int, s string) []string {
	if n <= 0 {
		return []string{}
	}
	resLen := len(s) - n + 1
	if resLen <= 0 { // Too short
		return []string{s}
	}

	res := make([]string, resLen) // This may be too long due to UTF-8.
	r := 0

	idx := make([]int, n)
	idx[n-1] = -1 // for unigram
	k := 0
	for i := range s {
		if idx[n-1] >= 0 {
			// This function intentionally avoid using append because it's
			// slower than the direct assignment.
			res[r] = s[idx[k]:i]
			r++
		}
		idx[k] = i
		k = (k + 1) % n
	}
	if r == 0 { // Because s contains UTF-8, len(s) > n and r == 0 could be true.
		return []string{s}
	}

	res[r] = s[idx[k]:]
	return res[:r+1]
}

// WordNGram creates word n-grams from the given array of words. A separator
// sep can be any string.It returns an empty array when n isn't greater than 0.
func WordNGram(n int, words []string, sep string) []string {
	if n <= 0 {
		return []string{}
	}
	if len(words) <= n {
		return []string{strings.Join(words, sep)}
	}

	res := make([]string, len(words)-n+1)
	for i := 0; i <= len(words)-n; i++ {
		res[i] = strings.Join(words[i:i+n], sep)
	}
	return res
}

// RemoveEmptyWord removes an empty string from an array of strings.
func RemoveEmptyWord(a []string) []string {
	i := 0
	res := make([]string, len(a))
	for _, s := range a {
		if len(s) != 0 {
			res[i] = s
			i++
		}
	}
	return res[:i]
}
