package nlp

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestWeightTF(t *testing.T) {
	Convey("Given an array of words", t, func() {
		words := []string{"a", "b", "c", "a", "b", "d", "a", "a", "d"}

		Convey("when applying WeightTF", func() {
			w := WeightTF(words)

			Convey("it should have a correct lenght", func() {
				So(len(w), ShouldEqual, 4)
			})

			Convey("it should have correct counts", func() {
				So(w["a"], ShouldEqual, 4)
				So(w["b"], ShouldEqual, 2)
				So(w["c"], ShouldEqual, 1)
				So(w["d"], ShouldEqual, 2)
			})
		})

		Convey("when applying WeightBinary", func() {
			w := WeightBinary(words)

			Convey("it should have a correct lenght", func() {
				So(len(w), ShouldEqual, 4)
			})

			Convey("it should have correct counts", func() {
				So(w["a"], ShouldEqual, 1)
				So(w["b"], ShouldEqual, 1)
				So(w["c"], ShouldEqual, 1)
				So(w["d"], ShouldEqual, 1)
			})
		})
	})
}
