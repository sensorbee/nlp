package nlp

import (
	"gopkg.in/sensorbee/sensorbee.v0/data"
	"math"
)

// WeightTF creates a map having a word as a key and its count (i.e. tf) as
// a value.
func WeightTF(a []string) data.Map {
	m := map[string]int{}
	for _, s := range a {
		c := m[s]
		m[s] = c + 1
	}

	res := data.Map{}
	for k, v := range m {
		res[k] = data.Float(v)
	}
	return res
}

// WeightBinary creates a map having weights of each word. The weight is 1 if
// there's at least one word, or 0 otherwise. Because feature vectors created
// by this function is sparse, all values in resulting maps are 1. In other
// words, instead of having 0 as a value, a key doesn't exist for a word that
// is not in the given array.
func WeightBinary(a []string) data.Map {
	res := data.Map{}
	for _, s := range a {
		res[s] = data.Float(1)
	}
	return res
}

// WeightLogTF creates a map having a word as a key and its log(1 + tf) as
// a value. This function is useful when some words appear too much but
// binary weight isn't sufficient.
func WeightLogTF(a []string) data.Map {
	m := map[string]int{}
	for _, s := range a {
		c := m[s]
		m[s] = c + 1
	}

	res := data.Map{}
	for k, v := range m {
		res[k] = data.Float(math.Log(1 + float64(v)))
	}
	return res
}
