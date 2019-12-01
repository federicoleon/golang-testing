package mysort

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {

	elements := []int{9, 8, 7, 5, 4, 6, 1, 2, 3, 0}
	felement := elements[0]
	lelement := elements[len(elements)-1]
	// if len(elements) != 10 {
	// 	t.Error("slice should contain 10 elements")
	// }

	assert.EqualValues(t, 10, len(elements))

	timeoutChan := make(chan bool, 1)
	defer close(timeoutChan)
	go func() {
		BubbleSort(elements)
		timeoutChan <- false
	}()

	go func() {
		time.Sleep(1000 * time.Millisecond)
		timeoutChan <- true
	}()

	if <-timeoutChan {
		assert.Fail(t, "bubble sort took more than 1000 ms")
		return
	}
	assert.NotNil(t, elements)
	fmt.Println(elements)
	// if elements[len(elements)-1] != 0 {
	// 	t.Error("The lastElement should be 0. it is ", lelement)
	// }
	assert.EqualValues(t, lelement, 0)
	// if elements[0] != 9 {
	// 	t.Error("The first element should be 9, it is ", felement)
	// }
	assert.EqualValues(t, felement, 9)

	if elements[len(elements)-1] != 0 {
		t.Error("The lastElement should be 0. it is ", lelement)
	}

	if elements[0] != 9 {
		t.Error("The first element should be 9, it is ", felement)
	}
	fmt.Println(elements)
}

func TestBubbleSortDesc(t *testing.T) {

	elements := []int{9, 8, 7, 5, 4, 6, 1, 0, 3, 2}
	fmt.Println(elements)

	timeoutChan := make(chan bool, 1)
	defer close(timeoutChan)
	go func() {
		BubbleSortDesc(elements)
		timeoutChan <- false
	}()

	go func() {
		time.Sleep(1000 * time.Millisecond)
		timeoutChan <- true
	}()

	if <-timeoutChan {
		assert.Fail(t, "bubblesortdesc took more than 500 ms")
		return
	}

	lelement := elements[len(elements)-1]
	if elements[len(elements)-1] != 0 {
		t.Error("The lastElement should be 0. it is ", lelement)
	}
	felement := elements[0]
	if elements[0] != 9 {
		t.Error("The first element should be 9, it is ", felement)
	}
	fmt.Println(elements)
}

func TestSort(t *testing.T) {
	elements := []int{9, 8, 7, 5, 4, 6, 1, 0, 3, 2}
	fmt.Println(elements)

	Sort(elements)
	lelement := elements[len(elements)-1]
	if elements[len(elements)-1] != 9 {
		t.Error("The lastElement should be 9. it is ", lelement)
	}
	felement := elements[0]
	if elements[0] != 0 {
		t.Error("The first element should be 0, it is ", felement)
	}
	fmt.Println(elements)
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func getElements(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n - 1; i > 0; i-- {
		result[j] = i
		j++
	}
	return result
}
