package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBubbleSortIncreasingOrder(t *testing.T) {
	elements := GetElements(10)

	timeoutChan := make(chan bool, 1)
	defer close(timeoutChan)

	go func() {
		BubbleSort(elements)
		timeoutChan <- false
	}()

	// Timeout settings
	go func() {
		time.Sleep(50 * time.Millisecond)
		timeoutChan <- true
	}()

	if <- timeoutChan {
		assert.Fail(t, "bubble sort took more than 50 ms")
		return
	}

	BubbleSort(elements)

	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])
}

func TestSortIncreasingOrder(t *testing.T) {
	elements := GetElements(10)

	Sort(elements)

	assert.EqualValues(t, 0, elements[0], "first element should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "last elements should be 9")
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