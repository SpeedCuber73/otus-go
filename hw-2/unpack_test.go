package unpack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	src string
	res string
	err error
}

func new(src, res string, err error) *testCase {
	return &testCase{
		src: src,
		res: res,
		err: err,
	}
}

var testCases = [...]*testCase{
	new("4rr", "", ErrStartsWithDigit),
	new("", "", nil),
	new("w2q", "wwq", nil),
	new("a4bc2d5e", "aaaabccddddde", nil),
	new("abcd", "abcd", nil),
	new("45", "", ErrStartsWithDigit),
	new(`tw\`, "", ErrEndsWithEscape),
	new(`qwe\4\5`, `qwe45`, nil),
	new(`qwe\45`, `qwe44444`, nil),
	new(`qwe\\5`, `qwe\\\\\`, nil),
}

func TestUnpack(t *testing.T) {
	assert := assert.New(t)
	for _, testCase := range testCases {
		res, err := Unpack(testCase.src)
		assert.Equal(testCase.res, res, fmt.Sprintf("Failed on res of %s", testCase.src))
		assert.Equal(testCase.err, err, fmt.Sprintf("Failed on err of %s", testCase.src))
	}
}
