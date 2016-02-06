package nlp

// NGram returns UTF-8 character n-grams created from the given text. This
// function assumes that s only contains valid UTF-8 letters. It returns an
// empty array when n isn't greater than 0.
func NGram(n int, s string) []string {
	if n <= 0 {
		return []string{}
	}

	res := make([]string, len(s)-n+1) // This may be too much
	r := 0

	idx := make([]int, n)
	idx[n-1] = -1 // for unigram
	k := 0
	for i := range s {
		if idx[n-1] >= 0 {
			res[r] = s[idx[k]:i]
			r++
		}
		idx[k] = i
		k = (k + 1) % n
	}
	res[r] = s[idx[k]:]
	return res[:r+1]
}
