package sets_test

import (
	"testing"

	"github.com/hienduyph/collections/sets"
	"github.com/stretchr/testify/assert"
)

func Test_Set(t *testing.T) {
	art := assert.New(t)
	setStr := sets.NewWithCap[string](10)
	setStr.Add("my")
	art.True(setStr.Has("my"))
	art.Equal(1, setStr.Length())

	setStr.Add("hello")
	art.Equal(2, setStr.Length())
	setStr.Pop("hello")
	art.Equal(1, setStr.Length())
}
