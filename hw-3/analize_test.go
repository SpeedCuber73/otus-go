package analize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type indexTest struct {
	text        string
	expectedTop []string
}

var testCases = [...]indexTest{
	{
		"asdf, d!s this. What's what this is? d",
		[]string{"asdf", "d", "is", "s", "this", "what"},
	},
	{
		"hi hi hi hi Hi hI HI",
		[]string{"hi"},
	},
	{
		`
		I'm sitting here in the boring room
		It's just another rainy Sunday afternoon
		I'm wasting my time, I got nothing to do
		I'm hanging around, I'm waiting for you
		But nothing ever happens and I wonder

		I'm driving around in my car
		I'm driving too fast, I'm driving too far
		I'd like to change my point of view
		I feel so lonely, I'm waiting for you
		But nothing ever happens and I wonder

		I wonder how, I wonder why
		Yesterday you told me 'bout the blue blue sky
		And all that I can see is just a yellow lemon tree
		I'm turning my head up and down
		I'm turning, turning, turning, turning, turning around
		And all that I can see is just another lemon tree

		I'm sitting here, I miss the power
		I'd like to go out, taking a shower
		But there's a heavy cloud inside my head
		I feel so tired, I put myself into bed
		Where nothing ever happens and I wonder

		Isolation is not good for me
		Isolation, I don't want to sit on a lemon tree

		I'm stepping around in a desert of joy
		Maybe anyhow I'll get another toy
		And everything will happen and you'll wonder

		I wonder how, I wonder why
		Yesterday you told me 'bout the blue blue sky
		And all that I can see is just another lemon tree
		I'm turning my head up and down
		Turning, turning, turning, turning, turning around
		And all that I can see is just a yellow lemon tree

		And I wonder, I wonder, I wonder how, I wonder why
		Yesterday you told me 'bout the blue blue sky
		And all that I can see
		And all that I can see
		And all that I can see is just a yellow lemon tree
		`,
		[]string{"i", "and", "m", "wonder", "turning", "can", "that", "a", "see", "all"},
	},
}

func TestMostCommonWords(t *testing.T) {
	for _, testCase := range testCases {
		topWords := MostCommonWords(testCase.text)
		assert.Equal(t, len(testCase.expectedTop), len(topWords), "unexpected length of top words")
		for _, word := range testCase.expectedTop {
			assert.Contains(t, topWords, word)
		}
	}
}
