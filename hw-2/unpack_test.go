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

var testCases = [...]testCase{
	{"4rr", "", ErrStartsWithDigit},
	{"", "", nil},
	{"w2q", "wwq", nil},
	{"a4bc2d5e", "aaaabccddddde", nil},
	{"abcd", "abcd", nil},
	{"45", "", ErrStartsWithDigit},
	{`tw\`, "", ErrEndsWithEscape},
	{`qwe\4\5`, `qwe45`, nil},
	{`qwe\45`, `qwe44444`, nil},
	{`qwe\\5`, `qwe\\\\\`, nil},
	{`qwe\\54`, ``, ErrWrongString},
}

func TestUnpack(t *testing.T) {
	assert := assert.New(t)
	for _, testCase := range testCases {
		res, err := Unpack(testCase.src)
		assert.Equal(testCase.res, res, fmt.Sprintf("Failed on res of %s", testCase.src))
		assert.Equal(testCase.err, err, fmt.Sprintf("Failed on err of %s", testCase.src))
	}
}
