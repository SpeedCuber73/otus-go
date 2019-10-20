package envdir

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	envVars = []string{"hi", "greeting"}
	envDir  = "testdata/env"
	prog    = "testdata/prog"
)

// ATTENTION!!! testdata/prog must be executable
func TestEvndir(t *testing.T) {
	assert := assert.New(t)

	for _, envVar := range envVars {
		command := []string{prog, envVar}
		out, err := Run(envDir, command)
		assert.Nil(err)
		expected, _ := ioutil.ReadFile(path.Join(envDir, envVar))
		assert.Equal(string(expected), out)
	}
}
