package gocopy

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGocopy(t *testing.T) {
	assert := assert.New(t)

	from := "testdata/source"
	to := "testdata/dest"
	offset := int64(10)
	limit := int64(100000)

	err := Gocopy(from, to, offset, limit)
	assert.Nil(err)

	src, _ := ioutil.ReadFile(from)
	dst, _ := ioutil.ReadFile(to)
	assert.Equal(0, bytes.Compare(src[10:100010], dst))
}
