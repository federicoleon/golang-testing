package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSortIncreasingOrder(t *testing.T) {
	elements := GetElements(10)

	BubbleSort(elements)

	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])
}

func TestSortIncreasingOrder(t *testing.T) {
	elements := GetElements(10)

	Sort(elements)

	if elements[0] != 0 {
		t.Error("first element should be 0")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("last elements should be 9")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := GetElements(10000)

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := GetElements(10000)

	for i :=0; i < b.N; i++ {
		Sort(elements)
	}
}