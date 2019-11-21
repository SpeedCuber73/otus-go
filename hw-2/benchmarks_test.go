package unpack

import (
	"testing"
)

var benchCases = []string{
	"е4rr",
	"a4bc2d5a4bc2d5a4bc2d5a4bc2d5a4bc2d5a4bc2d5a4bc2d5a4bc2d5a4bc2d5a4bc2d5a4bc2d5a4bc2d5",
	"45",
	`tw\`,
	`qwe\4\5`,
	`qwe\\54`,
}

func BenchmarkUnpack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, str := range benchCases {
			Unpack(str)
		}
	}
}
