package sort

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBubbleSortIncreasingOrder(t *testing.T) {
	elements := GetElements(10)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])

	BubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])
}

func TestSortIncreasingOrder(t *testing.T) {
	elements := GetElements(10)

	Sort(elements)

	assert.EqualValues(t, 0, elements[0], "first element should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "last element should be 9")
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := GetElements(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := GetElements(10000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
