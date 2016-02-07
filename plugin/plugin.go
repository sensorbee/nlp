package plugin

import (
	"github.com/sensorbee/nlp"
	"pfi/sensorbee/sensorbee/bql/udf"
)

func init() {
	udf.MustRegisterGlobalUDF("nlp_ngram", udf.MustConvertGeneric(nlp.NGram))
	udf.MustRegisterGlobalUDF("nlp_word_ngram",
		udf.MustConvertGeneric(nlp.WordNGram))
}
