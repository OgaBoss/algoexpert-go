package bubble_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sorted := BubbleSort(nums)
	assert.Equal(t, expected, sorted)
}
