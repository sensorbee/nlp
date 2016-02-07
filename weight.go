package nlp

import (
	"pfi/sensorbee/sensorbee/data"
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
