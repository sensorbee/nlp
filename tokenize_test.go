package nlp

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNGram(t *testing.T) {
	Convey("Given NGram function", t, func() {
		Convey("when n is too small", func() {
			Convey("it should return an empty array with n = 0", func() {
				So(NGram(0, "abc"), ShouldBeEmpty)
			})

			Convey("it should return an empty array with n < 0", func() {
				So(NGram(-1, "abc"), ShouldBeEmpty)
			})
		})

		Convey("when n is 1", func() {
			Convey("it should tokenize 1 ascii letter", func() {
				ns := NGram(1, "a")
				So(len(ns), ShouldEqual, 1)
				So(ns[0], ShouldEqual, "a")
			})

			Convey("it should tokenize 1 UTF-8 letter", func() {
				ns := NGram(1, "あ")
				So(len(ns), ShouldEqual, 1)
				So(ns[0], ShouldEqual, "あ")
			})

			Convey("it should tokenize ascii letters", func() {
				ns := NGram(1, "abc")
				So(len(ns), ShouldEqual, 3)
				So(ns[0], ShouldEqual, "a")
				So(ns[1], ShouldEqual, "b")
				So(ns[2], ShouldEqual, "c")
			})

			Convey("it should tokenize UTF-8 letters", func() {
				ns := NGram(1, "ａbｃ")
				So(len(ns), ShouldEqual, 3)
				So(ns[0], ShouldEqual, "ａ")
				So(ns[1], ShouldEqual, "b")
				So(ns[2], ShouldEqual, "ｃ")
			})
		})

		Convey("When n is 2", func() {
			Convey("it should accept a string having 1 letter", func() {
				ns := NGram(2, "a")
				So(len(ns), ShouldEqual, 1)
				So(ns[0], ShouldEqual, "a")
			})

			Convey("it should tokenize 2 ascii letters", func() {
				ns := NGram(2, "ab")
				So(len(ns), ShouldEqual, 1)
				So(ns[0], ShouldEqual, "ab")
			})

			Convey("it should tokenize 2 UTF-8 letters", func() {
				ns := NGram(2, "aｂ")
				So(len(ns), ShouldEqual, 1)
				So(ns[0], ShouldEqual, "aｂ")
			})

			Convey("it should tokenize ascii letters", func() {
				ns := NGram(2, "abcd")
				So(len(ns), ShouldEqual, 3)
				So(ns[0], ShouldEqual, "ab")
				So(ns[1], ShouldEqual, "bc")
				So(ns[2], ShouldEqual, "cd")
			})

			Convey("it should tokenize UTF-8 letters", func() {
				ns := NGram(2, "ａbｃd")
				So(len(ns), ShouldEqual, 3)
				So(ns[0], ShouldEqual, "ａb")
				So(ns[1], ShouldEqual, "bｃ")
				So(ns[2], ShouldEqual, "ｃd")
			})
		})

		Convey("When n is 3", func() {
			Convey("it should accept a string having 1 letter", func() {
				ns := NGram(3, "a")
				So(len(ns), ShouldEqual, 1)
				So(ns[0], ShouldEqual, "a")
			})

			Convey("it should accept a string having 2 letters", func() {
				ns := NGram(3, "ab")
				So(len(ns), ShouldEqual, 1)
				So(ns[0], ShouldEqual, "ab")
			})

			Convey("it should tokenize 3 ascii letters", func() {
				ns := NGram(3, "abc")
				So(len(ns), ShouldEqual, 1)
				So(ns[0], ShouldEqual, "abc")
			})

			Convey("it should tokenize 2 UTF-8 letters", func() {
				ns := NGram(3, "ａbｃ")
				So(len(ns), ShouldEqual, 1)
				So(ns[0], ShouldEqual, "ａbｃ")
			})

			Convey("it should tokenize ascii letters", func() {
				ns := NGram(3, "abcde")
				So(len(ns), ShouldEqual, 3)
				So(ns[0], ShouldEqual, "abc")
				So(ns[1], ShouldEqual, "bcd")
				So(ns[2], ShouldEqual, "cde")
			})

			Convey("it should tokenize UTF-8 letters", func() {
				ns := NGram(3, "ａbｃdｅ")
				So(len(ns), ShouldEqual, 3)
				So(ns[0], ShouldEqual, "ａbｃ")
				So(ns[1], ShouldEqual, "bｃd")
				So(ns[2], ShouldEqual, "ｃdｅ")
			})
		})
	})
}

func TestWordNGram(t *testing.T) {
	Convey("Given WordNGram function", t, func() {
		Convey("when n is too small", func() {
			Convey("it should return an empty array with n = 0", func() {
				ns := WordNGram(0, []string{"a", "b", "c"}, " ")
				So(len(ns), ShouldEqual, 0)
			})

			Convey("it should return an empty array with n < 0", func() {
				ns := WordNGram(-1, []string{"a", "b", "c"}, " ")
				So(len(ns), ShouldEqual, 0)
			})
		})

		Convey("when n is 1", func() {
			Convey("it should work correctly", func() {
				ns := WordNGram(1, []string{"a", "b", "c"}, " ")
				So(len(ns), ShouldEqual, 3)
				So(ns[0], ShouldEqual, "a")
				So(ns[1], ShouldEqual, "b")
				So(ns[2], ShouldEqual, "c")
			})
		})

		Convey("when n is 2", func() {
			Convey("it should work correctly", func() {
				ns := WordNGram(2, []string{"a", "b", "c", "d"}, " ")
				So(len(ns), ShouldEqual, 3)
				So(ns[0], ShouldEqual, "a b")
				So(ns[1], ShouldEqual, "b c")
				So(ns[2], ShouldEqual, "c d")
			})
		})

		Convey("when n is 3", func() {
			Convey("it should work correctly", func() {
				ns := WordNGram(3, []string{"a", "b", "c", "d", "e"}, " ")
				So(len(ns), ShouldEqual, 3)
				So(ns[0], ShouldEqual, "a b c")
				So(ns[1], ShouldEqual, "b c d")
				So(ns[2], ShouldEqual, "c d e")
			})
		})
	})
}
