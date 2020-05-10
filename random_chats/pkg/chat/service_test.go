package chat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitGroups(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e"}
	res := splitGroups(slice, 3)
	assert.Len(t, res, 2)
	assert.Equal(t, []string{"a", "b", "c"}, res[0])
	assert.Equal(t, []string{"d", "e"}, res[1])

	slice = []string{"a", "b", "c", "d"}
	res = splitGroups(slice, 3)
	assert.Len(t, res, 1)
	assert.Equal(t, []string{"a", "b", "c", "d"}, res[0])

	slice = []string{"a", "b", "c", "d", "e"}
	res = splitGroups(slice, 2)
	assert.Len(t, res, 2)
	assert.Equal(t, []string{"a", "b"}, res[0])
	assert.Equal(t, []string{"c", "d", "e"}, res[1])

	slice = []string{"a", "b", "c", "d", "e"}
	res = splitGroups(slice, 5)
	assert.Len(t, res, 1)
	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, res[0])

	slice = []string{"a", "b", "c", "d", "e"}
	res = splitGroups(slice, 6)
	assert.Len(t, res, 1)
	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, res[0])
}

func TestComposeMessage(t *testing.T) {
	testData := [][]string{{"a", "b", "c"}, {"d", "e"}}
	msg := composeMessage(testData)
	expected := "Good day team:roller_coaster:. The random chat roster of this week :scroll::\\n@a :blob-wine-gif: @b :blob-wine-gif: @c\\n@d :blob-wine-gif: @e"
	assert.Equal(t, expected, msg)
}

func TestShuffle(t *testing.T) {
	testData := []string{"a", "b", "c", "d", "e"}
	randomize := shuffle(testData)
	assert.Len(t, randomize, 5)
	assert.NotEqual(t, randomize, testData)
}
