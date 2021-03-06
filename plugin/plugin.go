package plugin

import (
	"github.com/sensorbee/nlp"
	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
	"strings"
)

func init() {
	udf.MustRegisterGlobalUDF("nlp_ngram", udf.MustConvertGeneric(nlp.NGram))
	udf.MustRegisterGlobalUDF("nlp_word_ngram",
		udf.MustConvertGeneric(nlp.WordNGram))

	udf.MustRegisterGlobalUDF("nlp_split", udf.MustConvertGeneric(strings.Split))
	udf.MustRegisterGlobalUDF("nlp_remove_empty_word",
		udf.MustConvertGeneric(nlp.RemoveEmptyWord))

	udf.MustRegisterGlobalUDF("nlp_weight_tf",
		udf.MustConvertGeneric(nlp.WeightTF))
	udf.MustRegisterGlobalUDF("nlp_weight_binary",
		udf.MustConvertGeneric(nlp.WeightBinary))
	udf.MustRegisterGlobalUDF("nlp_weight_log_tf",
		udf.MustConvertGeneric(nlp.WeightLogTF))

	udf.MustRegisterGlobalUDF("nlp_to_lower",
		udf.MustConvertGeneric(strings.ToLower))
	udf.MustRegisterGlobalUDF("nlp_to_upper",
		udf.MustConvertGeneric(strings.ToUpper))
}
